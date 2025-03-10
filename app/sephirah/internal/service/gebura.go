package service

import (
	"context"
	"io"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LibrarianSephirahServiceService) CreateApp(ctx context.Context, req *pb.CreateAppRequest) (
	*pb.CreateAppResponse, error,
) {
	app := req.GetApp()
	if app == nil {
		return nil, pb.ErrorErrorReasonBadRequest("app required")
	}
	a, err := s.g.CreateApp(ctx, converter.ToBizApp(req.GetApp()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateAppResponse{
		Id: converter.ToPBInternalID(a.ID),
	}, nil
}
func (s *LibrarianSephirahServiceService) UpdateApp(ctx context.Context, req *pb.UpdateAppRequest) (
	*pb.UpdateAppResponse, error,
) {
	app := req.GetApp()
	if app == nil || app.GetId() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("app and internal_id required")
	}
	err := s.g.UpdateApp(ctx, converter.ToBizApp(req.GetApp()))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAppResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListApps(ctx context.Context, req *pb.ListAppsRequest) (
	*pb.ListAppsResponse, error,
) {
	a, total, err := s.g.ListApps(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter.ToBizAppSourceList(req.GetSourceFilter()),
		converter.ToBizAppTypeList(req.GetTypeFilter()),
		converter.ToBizInternalIDList(req.GetIdFilter()),
		req.GetContainDetails())
	if err != nil {
		return nil, err
	}
	return &pb.ListAppsResponse{
		Paging: &librarian.PagingResponse{TotalSize: total},
		Apps:   converter.ToPBAppList(a),
	}, nil
}
func (s *LibrarianSephirahServiceService) RefreshApp(ctx context.Context, req *pb.RefreshAppRequest) (
	*pb.RefreshAppResponse, error,
) {
	return nil, pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) MergeApps(ctx context.Context, req *pb.MergeAppsRequest) (
	*pb.MergeAppsResponse, error,
) {
	app := converter.ToBizApp(req.Base)
	if app == nil {
		return nil, pb.ErrorErrorReasonBadRequest("base required")
	}
	err := s.g.MergeApps(ctx,
		*app,
		converter.ToBizInternalID(req.GetMerged()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.MergeAppsResponse{}, nil
}
func (s *LibrarianSephirahServiceService) SearchApps(ctx context.Context, req *pb.SearchAppsRequest) (
	*pb.SearchAppsResponse, error,
) {
	apps, total, err := s.g.SearchApps(ctx,
		model.ToBizPaging(req.GetPaging()),
		req.GetKeywords(),
	)
	if err != nil {
		return nil, err
	}
	return &pb.SearchAppsResponse{
		Paging: &librarian.PagingResponse{TotalSize: int64(total)},
		Apps:   converter.ToPBAppList(apps),
	}, nil
}
func (s *LibrarianSephirahServiceService) GetApp(ctx context.Context, req *pb.GetAppRequest) (
	*pb.GetAppResponse, error,
) {
	res, err := s.g.GetApp(ctx, converter.ToBizInternalID(req.GetAppId()))
	if err != nil {
		return nil, err
	}
	return &pb.GetAppResponse{App: converter.ToPBApp(res)}, nil
}
func (s *LibrarianSephirahServiceService) GetBindApps(ctx context.Context, req *pb.GetBindAppsRequest) (
	*pb.GetBindAppsResponse, error,
) {
	al, err := s.g.GetBindApps(ctx, converter.ToBizInternalID(req.GetAppId()))
	if err != nil {
		return nil, err
	}
	return &pb.GetBindAppsResponse{Apps: converter.ToPBAppList(al)}, nil
}
func (s *LibrarianSephirahServiceService) PurchaseApp(ctx context.Context, req *pb.PurchaseAppRequest) (
	*pb.PurchaseAppResponse, error,
) {
	err := s.g.PurchaseApp(ctx, converter.ToBizInternalID(req.GetAppId()))
	if err != nil {
		return nil, err
	}
	return &pb.PurchaseAppResponse{}, nil
}
func (s *LibrarianSephirahServiceService) GetPurchasedApps(ctx context.Context, req *pb.GetPurchasedAppsRequest) (
	*pb.GetPurchasedAppsResponse, error,
) {
	apps, err := s.g.GetPurchasedApps(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetPurchasedAppsResponse{
		Apps: converter.ToPBAppList(apps),
	}, nil
}

func (s *LibrarianSephirahServiceService) CreateAppPackage(
	ctx context.Context,
	req *pb.CreateAppPackageRequest,
) (*pb.CreateAppPackageResponse, error) {
	ap, err := s.g.CreateAppPackage(ctx, converter.ToBizAppPackage(req.GetAppPackage()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateAppPackageResponse{Id: converter.ToPBInternalID(ap.ID)}, nil
}
func (s *LibrarianSephirahServiceService) UpdateAppPackage(
	ctx context.Context,
	req *pb.UpdateAppPackageRequest,
) (*pb.UpdateAppPackageResponse, error) {
	err := s.g.UpdateAppPackage(ctx, converter.ToBizAppPackage(req.GetAppPackage()))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAppPackageResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListAppPackages(
	ctx context.Context,
	req *pb.ListAppPackagesRequest,
) (*pb.ListAppPackagesResponse, error) {
	ap, total, err := s.g.ListAppPackages(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter.ToBizAppPackageSourceList(req.GetSourceFilter()),
		converter.ToBizInternalIDList(req.GetIdFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListAppPackagesResponse{
		Paging:      &librarian.PagingResponse{TotalSize: int64(total)},
		AppPackages: converter.ToPBAppPackageList(ap),
	}, nil
}
func (s *LibrarianSephirahServiceService) AssignAppPackage(
	ctx context.Context,
	req *pb.AssignAppPackageRequest,
) (*pb.AssignAppPackageResponse, error) {
	err := s.g.AssignAppPackage(ctx,
		converter.ToBizInternalID(req.GetAppId()),
		converter.ToBizInternalID(req.GetAppPackageId()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.AssignAppPackageResponse{}, nil
}
func (s *LibrarianSephirahServiceService) UnAssignAppPackage(
	ctx context.Context,
	req *pb.UnAssignAppPackageRequest,
) (*pb.UnAssignAppPackageResponse, error) {
	err := s.g.UnAssignAppPackage(ctx, converter.ToBizInternalID(req.GetAppPackageId()))
	if err != nil {
		return nil, err
	}
	return &pb.UnAssignAppPackageResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ReportAppPackages(
	conn pb.LibrarianSephirahService_ReportAppPackagesServer,
) error {
	ctx, err1 := s.authFunc(conn.Context())
	if err1 != nil {
		return err1
	}
	handler, err2 := s.g.NewReportAppPackageHandler(ctx)
	if err2 != nil {
		return err2
	}
	for {
		var binaries []*modelgebura.AppPackageBinary
		if req, err := conn.Recv(); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		} else {
			binaries = converter.ToBizAppPackageBinaryList(req.GetAppPackageBinaries())
		}
		if err := handler.Handle(conn.Context(), binaries); err != nil {
			return err
		}
		if err := conn.Send(&pb.ReportAppPackagesResponse{}); err != nil {
			return err
		}
	}
}
func (s *LibrarianSephirahServiceService) UploadGameSaveFile(
	ctx context.Context,
	req *pb.UploadGameSaveFileRequest,
) (*pb.UploadGameSaveFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadGameSaveFile not implemented")
}
func (s *LibrarianSephirahServiceService) DownloadGameSaveFile(
	ctx context.Context,
	req *pb.DownloadGameSaveFileRequest,
) (*pb.DownloadGameSaveFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadGameSaveFile not implemented")
}
func (s *LibrarianSephirahServiceService) ListGameSaveFiles(
	ctx context.Context,
	req *pb.ListGameSaveFilesRequest,
) (*pb.ListGameSaveFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGameSaveFile not implemented")
}
