package bizchesed

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"net/http"
	"strconv"
	"sync"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelchesed"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcodec"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/wire"
	"google.golang.org/protobuf/types/known/durationpb"
)

var ProviderSet = wire.NewSet(
	NewChesed,
	NewImageCache,
)

type ChesedRepo interface {
	CreateImage(context.Context, model.InternalID, *modelchesed.Image) error
	ListImages(context.Context, model.InternalID, model.Paging) ([]*modelchesed.Image, int64, error)
	ListImageNeedScan(context.Context) ([]*modelchesed.Image, error)
	SetImageStatus(context.Context, model.InternalID, modelchesed.ImageStatus) error
	GetImage(context.Context, model.InternalID, model.InternalID) (*modelchesed.Image, error)
}

type Chesed struct {
	repo        ChesedRepo
	mapper      mapper.LibrarianMapperServiceClient
	searcher    searcher.LibrarianSearcherServiceClient
	porter      porter.LibrarianPorterServiceClient
	upload      *modelbinah.UploadCallBack
	download    *modelbinah.DownloadCallBack
	imageCache  *libcache.Map[model.InternalID, modelchesed.Image]
	muScanImage sync.Mutex
}

func NewChesed(
	repo ChesedRepo,
	cron *libcron.Cron,
	mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient searcher.LibrarianSearcherServiceClient,
	block *modelbinah.ControlBlock,
	imageCache *libcache.Map[model.InternalID, modelchesed.Image],
) (*Chesed, error) {
	c := &Chesed{
		repo:        repo,
		mapper:      mClient,
		porter:      pClient,
		searcher:    sClient,
		upload:      nil,
		download:    nil,
		imageCache:  imageCache,
		muScanImage: sync.Mutex{},
	}
	c.upload = block.RegisterUploadCallback(
		modelbinah.UploadChesedImage,
		c.UploadImageCallback,
	)
	c.download = block.RegisterDownloadCallback(
		modelbinah.DownloadEmpty,
		nil,
	)
	err := cron.BySeconds(60, c.ScanImage, context.Background()) //nolint:gomnd //TODO
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewImageCache(
	store libcache.Store,
) *libcache.Map[model.InternalID, modelchesed.Image] {
	return libcache.NewMap[model.InternalID, modelchesed.Image](
		store,
		"Images",
		func(k model.InternalID) string {
			return strconv.FormatInt(int64(k), 10)
		},
		nil,
		libcache.WithExpiration(libtime.Day),
	)
}

func (c *Chesed) UploadImage(ctx context.Context, image modelchesed.Image,
	metadata modelbinah.FileMetadata) (string, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return "", pb.ErrorErrorReasonForbidden("no permission")
	}
	if err := metadata.Check(); err != nil {
		return "", pb.ErrorErrorReasonBadRequest("invalid file metadata: %s", err.Error())
	}
	resp, err := c.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	image.ID = converter.ToBizInternalID(resp.Id)
	metadata.ID = converter.ToBizInternalID(resp.Id)
	err = c.imageCache.Set(ctx, image.ID, &image)
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	token, err := c.upload.GenerateUploadToken(ctx, metadata, libtime.HalfDay)
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return token, nil
}

func (c *Chesed) UploadImageCallback(ctx context.Context, id model.InternalID) error {
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	image, err := c.imageCache.Get(ctx, id)
	if err != nil {
		return err
	}
	err = c.repo.CreateImage(ctx, claims.InternalID, image)
	if err != nil {
		return err
	}
	return nil
}

func (c *Chesed) ScanImage(ctx context.Context) { //nolint:funlen,gocognit //TODO
	if c.muScanImage.TryLock() {
		defer c.muScanImage.Unlock()
	} else {
		return
	}
	images, err0 := c.repo.ListImageNeedScan(ctx)
	if err0 != nil {
		return
	}
	if len(images) == 0 {
		return
	}
	httpc := new(http.Client)
	for _, image := range images {
		data, err := c.porter.PresignedPullData(ctx, &porter.PresignedPullDataRequest{
			Source:     porter.DataSource_DATA_SOURCE_INTERNAL_DEFAULT,
			ContentId:  strconv.FormatInt(int64(image.ID), 10),
			ExpireTime: durationpb.New(libtime.Day),
		})
		if err != nil {
			return
		}
		getReq, err := http.NewRequestWithContext(ctx, http.MethodGet, data.PullUrl, bytes.NewBuffer([]byte("")))
		if err != nil {
			return
		}
		getResp, err := httpc.Do(getReq)
		if err != nil {
			return
		}
		imgBytes, err := io.ReadAll(getResp.Body)
		if err != nil {
			return
		}
		if len(imgBytes) == 0 {
			continue
		}
		err = getResp.Body.Close()
		if err != nil {
			return
		}
		ocrReq, err := http.NewRequestWithContext(
			ctx,
			http.MethodPost,
			"http://localhost:9010/predict/ocr_system",
			bytes.NewBuffer([]byte("{\"images\":[\""+base64.StdEncoding.EncodeToString(imgBytes)+"\"]}")),
		)
		if err != nil {
			return
		}
		ocrReq.Header.Set("Content-Type", "application/json")
		ocrResp, err := httpc.Do(ocrReq)
		if err != nil {
			return
		}
		body, err := io.ReadAll(ocrResp.Body)
		if err != nil {
			return
		}
		err = ocrResp.Body.Close()
		if err != nil {
			return
		}
		ocrResponse := modelchesed.OCRResponse{} //nolint:exhaustruct //TODO
		err = libcodec.Unmarshal(libcodec.JSON, body, &ocrResponse)
		if err != nil {
			return
		}
		if len(ocrResponse.Results) != 1 {
			return
		}
		if len(ocrResponse.Results[0]) > 0 {
			var desReq string
			var desReqStr []byte
			for _, r := range ocrResponse.Results[0] {
				desReq += r.Text + " "
			}
			desReqStr, err = libcodec.Marshal(libcodec.JSON, desReq)
			if err != nil {
				return
			}
			_, err = c.searcher.DescribeID(ctx, &searcher.DescribeIDRequest{
				Id:          converter.ToPBInternalID(image.ID),
				Description: string(desReqStr),
				Mode:        searcher.DescribeIDRequest_DESCRIBE_MODE_APPEND,
			})
			if err != nil {
				return
			}
		}
		err = c.repo.SetImageStatus(ctx, image.ID, modelchesed.ImageStatusScanned)
		if err != nil {
			return
		}
	}
}

func (c *Chesed) ListImages(ctx context.Context, paging model.Paging) ([]model.InternalID, int64, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	images, total, err := c.repo.ListImages(ctx, claims.InternalID, paging)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	res := make([]model.InternalID, 0, len(images))
	for _, image := range images {
		res = append(res, image.ID)
	}
	return res, total, nil
}

func (c *Chesed) SearchImages(ctx context.Context, paging model.Paging, keywords string) (
	[]model.InternalID, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	resp, err := c.searcher.SearchID(ctx,
		&searcher.SearchIDRequest{
			Paging: &librarian.PagingRequest{
				PageNum:  int32(paging.PageNum),
				PageSize: int32(paging.PageSize),
			},
			Keyword: keywords,
		},
	)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	res := make([]model.InternalID, 0, len(resp.GetResult()))
	for _, r := range resp.GetResult() {
		res = append(res, converter.ToBizInternalID(r.Id))
	}
	return res, nil
}

func (c *Chesed) DownloadImage(ctx context.Context, id model.InternalID) (string, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return "", pb.ErrorErrorReasonForbidden("no permission")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return "", pb.ErrorErrorReasonUnauthorized("empty token")
	}
	image, err := c.repo.GetImage(ctx, claims.InternalID, id)
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	token, err := c.download.GenerateDownloadToken(ctx, modelbinah.FileMetadata{
		ID:     id,
		Name:   image.Name,
		Size:   0,
		Type:   0,
		Sha256: nil,
	}, libtime.HalfDay)
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return token, nil
}
