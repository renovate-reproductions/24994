package biznetzach

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type NetzachRepo interface {
	CreateNotifyTarget(context.Context, model.InternalID, *modelnetzach.NotifyTarget) error
	UpdateNotifyTarget(context.Context, model.InternalID, *modelnetzach.NotifyTarget) error
	ListNotifyTargets(context.Context, model.Paging, model.InternalID, []model.InternalID,
		[]modelnetzach.NotifyTargetType, []modelnetzach.NotifyTargetStatus) (
		[]*modelnetzach.NotifyTarget, int64, error)
	GetNotifyTarget(context.Context, model.InternalID) (*modelnetzach.NotifyTarget, error)
	CreateNotifyFlow(context.Context, model.InternalID, *modelnetzach.NotifyFlow) error
	UpdateNotifyFlow(context.Context, model.InternalID, *modelnetzach.NotifyFlow) error
	ListNotifyFlows(context.Context, model.Paging, model.InternalID, []model.InternalID) (
		[]*modelnetzach.NotifyFlow, int64, error)
	GetNotifyFlow(context.Context, model.InternalID) (*modelnetzach.NotifyFlow, error)
	GetNotifyFlowIDsWithFeed(context.Context, model.InternalID) ([]model.InternalID, error)
}

type Netzach struct {
	repo                  NetzachRepo
	searcher              *client.Searcher
	feedToNotifyFlowCache *libcache.Map[model.InternalID, modelangela.FeedToNotifyFlowValue]
	notifyFlowCache       *libcache.Map[model.InternalID, modelnetzach.NotifyFlow]
	notifyTargetCache     *libcache.Map[model.InternalID, modelnetzach.NotifyTarget]
}

func NewNetzach(
	repo NetzachRepo,
	sClient *client.Searcher,
	feedToNotifyFlowCache *libcache.Map[model.InternalID, modelangela.FeedToNotifyFlowValue],
	notifyFlowCache *libcache.Map[model.InternalID, modelnetzach.NotifyFlow],
	notifyTargetCache *libcache.Map[model.InternalID, modelnetzach.NotifyTarget],
) *Netzach {
	y := &Netzach{
		repo,
		sClient,
		feedToNotifyFlowCache,
		notifyFlowCache,
		notifyTargetCache,
	}
	return y
}

func (n *Netzach) CreateNotifyTarget(ctx context.Context, target *modelnetzach.NotifyTarget) (
	model.InternalID, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	if target == nil {
		return 0, pb.ErrorErrorReasonBadRequest("notify target required")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return 0, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	id, err := n.searcher.NewID(ctx)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	target.ID = id
	err = n.repo.CreateNotifyTarget(ctx, claims.InternalID, target)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return target.ID, nil
}

func (n *Netzach) UpdateNotifyTarget(ctx context.Context, target *modelnetzach.NotifyTarget) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return pb.ErrorErrorReasonUnauthorized("empty token")
	}
	err := n.repo.UpdateNotifyTarget(ctx, claims.InternalID, target)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	err = n.notifyTargetCache.Delete(ctx, target.ID)
	if err != nil {
		logger.Errorf("failed to delete cache %s", err.Error())
	}
	return nil
}

func (n *Netzach) ListNotifyTargets(
	ctx context.Context,
	paging model.Paging,
	ids []model.InternalID,
	types []modelnetzach.NotifyTargetType,
	statuses []modelnetzach.NotifyTargetStatus,
) ([]*modelnetzach.NotifyTarget, int64, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, 0, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	targets, total, err := n.repo.ListNotifyTargets(ctx, paging, claims.InternalID, ids, types, statuses)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return targets, total, nil
}

func (n *Netzach) CreateNotifyFlow(ctx context.Context, flow *modelnetzach.NotifyFlow) (
	model.InternalID, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	if flow == nil {
		return 0, pb.ErrorErrorReasonBadRequest("notify target required")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return 0, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	id, err := n.searcher.NewID(ctx)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	flow.ID = id
	err = n.repo.CreateNotifyFlow(ctx, claims.InternalID, flow)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return flow.ID, nil
}

func (n *Netzach) UpdateNotifyFlow(ctx context.Context, flow *modelnetzach.NotifyFlow) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return pb.ErrorErrorReasonUnauthorized("empty token")
	}
	err := n.repo.UpdateNotifyFlow(ctx, claims.InternalID, flow)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if flow.Source != nil && len(flow.Source.FeedIDFilter) > 0 {
		for _, id := range flow.Source.FeedIDFilter {
			err = n.feedToNotifyFlowCache.Delete(ctx, id)
			if err != nil {
				logger.Errorf("failed to delete cache %s", err.Error())
			}
		}
	}
	err = n.notifyFlowCache.Delete(ctx, flow.ID)
	if err != nil {
		logger.Errorf("failed to delete cache %s", err.Error())
	}
	return nil
}

func (n *Netzach) ListNotifyFlows(
	ctx context.Context,
	paging model.Paging,
	ids []model.InternalID,
) ([]*modelnetzach.NotifyFlow, int64, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, 0, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	res, total, err := n.repo.ListNotifyFlows(ctx, paging, claims.InternalID, ids)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, total, nil
}
