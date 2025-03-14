package bizgebura

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"golang.org/x/exp/slices"
)

func (g *Gebura) CreateAppPackage(
	ctx context.Context,
	a *modelgebura.AppPackage,
) (*modelgebura.AppPackage, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	id, err := g.searcher.NewID(ctx)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	a.ID = id
	a.Source = modelgebura.AppPackageSourceManual
	a.SourceID = 0
	if _, err = g.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{
		VertexList: []*mapper.Vertex{{
			Vid:  int64(a.ID),
			Type: mapper.VertexType_VERTEX_TYPE_ABSTRACT,
		}},
	}); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if err = g.repo.CreateAppPackage(ctx, claims.InternalID, a); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return a, nil
}

func (g *Gebura) UpdateAppPackage(ctx context.Context, a *modelgebura.AppPackage) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return pb.ErrorErrorReasonUnauthorized("empty token")
	}
	a.Source = modelgebura.AppPackageSourceManual
	err := g.repo.UpdateAppPackage(ctx, claims.InternalID, a)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) ListAppPackages(
	ctx context.Context,
	paging model.Paging,
	sources []modelgebura.AppPackageSource,
	ids []model.InternalID,
) ([]*modelgebura.AppPackage, int, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	res, total, err := g.repo.ListAppPackages(ctx, paging, sources, ids)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, total, nil
}

func (g *Gebura) AssignAppPackage(
	ctx context.Context,
	appID model.InternalID,
	appPackageID model.InternalID,
) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	if claims, ok := libauth.FromContext(ctx); !ok {
		return pb.ErrorErrorReasonForbidden("no permission")
	} else {
		err := g.repo.AssignAppPackage(ctx, claims.InternalID, appID, appPackageID)
		if err != nil {
			return pb.ErrorErrorReasonUnspecified("%s", err)
		}
	}
	return nil
}

func (g *Gebura) UnAssignAppPackage(ctx context.Context, appPackageID model.InternalID) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	if claims, ok := libauth.FromContext(ctx); !ok {
		return pb.ErrorErrorReasonForbidden("no permission")
	} else {
		err := g.repo.UnAssignAppPackage(ctx, claims.InternalID, appPackageID)
		if err != nil {
			return pb.ErrorErrorReasonUnspecified("%s", err)
		}
	}
	return nil
}

func (g *Gebura) NewReportAppPackageHandler(ctx context.Context) (ReportAppPackageHandler, *errors.Error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist || claims == nil {
		return nil, pb.ErrorErrorReasonUnauthorized("token required")
	}
	checksums, err := g.repo.ListAppPackageBinaryChecksumOfOneSource(ctx,
		modelgebura.AppPackageSourceSentinel, claims.InternalID)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return &reportAppPackageHandler{
		g:        g,
		sourceID: claims.InternalID,
		sha256:   checksums,
	}, nil
}

type reportAppPackageHandler struct {
	g        *Gebura
	sourceID model.InternalID
	sha256   []string
}

func (r *reportAppPackageHandler) Handle(ctx context.Context, binaries []*modelgebura.AppPackageBinary) *errors.Error {
	var vl []*mapper.Vertex
	packages := make([]*modelgebura.AppPackage, 0, len(binaries))
	ids, err := r.g.searcher.NewBatchIDs(ctx, len(binaries))
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	for i := range binaries {
		packages = append(packages, new(modelgebura.AppPackage))
		if !slices.Contains(r.sha256, string(binaries[i].Sha256)) {
			packages[i].ID = ids[i]
			vl = append(vl, &mapper.Vertex{
				Vid:  int64(ids[i]),
				Type: mapper.VertexType_VERTEX_TYPE_OBJECT,
				Prop: nil,
			})
		}
		packages[i].Source = modelgebura.AppPackageSourceSentinel
		packages[i].SourceID = r.sourceID
	}
	if len(vl) > 0 {
		if _, err = r.g.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: vl}); err != nil {
			return pb.ErrorErrorReasonUnspecified("%s", err.Error())
		}
		if err = r.g.repo.UpsertAppPackages(ctx, r.sourceID, packages); err != nil {
			return pb.ErrorErrorReasonUnspecified("%s", err.Error())
		}
	}
	return nil
}
