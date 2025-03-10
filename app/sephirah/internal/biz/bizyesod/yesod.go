package bizyesod

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
)

type YesodRepo interface {
	CreateFeedConfig(context.Context, model.InternalID, *modelyesod.FeedConfig) error
	UpdateFeedConfig(context.Context, model.InternalID, *modelyesod.FeedConfig) error
	ListFeedConfigCategories(context.Context, model.InternalID) ([]string, error)
	ListFeedConfigNeedPull(context.Context, []modelyesod.FeedConfigSource, []modelyesod.FeedConfigStatus,
		modelyesod.ListFeedOrder, time.Time, int) ([]*modelyesod.FeedConfig, error)
	UpdateFeedConfigAsInQueue(context.Context, model.InternalID) error
	ListFeedConfigs(context.Context, model.InternalID, model.Paging, []model.InternalID, []model.InternalID,
		[]modelyesod.FeedConfigSource, []modelyesod.FeedConfigStatus, []string) ([]*modelyesod.FeedWithConfig, int, error)
	ListFeedItems(context.Context, model.InternalID, model.Paging, []model.InternalID,
		[]model.InternalID, []string, *model.TimeRange, []string) ([]*modelyesod.FeedItemDigest, int, error)
	GroupFeedItems(context.Context, model.InternalID, []model.TimeRange, []model.InternalID,
		[]model.InternalID, []string, int, []string) (
		map[model.TimeRange][]*modelyesod.FeedItemDigest, error)
	GetFeedItems(context.Context, model.InternalID, []model.InternalID) ([]*modelfeed.Item, error)
}

type Yesod struct {
	repo     YesodRepo
	mapper   mapper.LibrarianMapperServiceClient
	searcher *client.Searcher
	pullFeed *libmq.Topic[modelyesod.PullFeed]
}

func NewYesod(
	repo YesodRepo,
	cron *libcron.Cron,
	mClient mapper.LibrarianMapperServiceClient,
	sClient *client.Searcher,
	pullFeed *libmq.Topic[modelyesod.PullFeed],
) (*Yesod, error) {
	y := &Yesod{
		repo:     repo,
		mapper:   mClient,
		searcher: sClient,
		pullFeed: pullFeed,
	}
	err := cron.BySeconds(60, y.PullFeeds, context.Background()) //nolint:gomnd // hard code min interval
	if err != nil {
		return nil, err
	}
	return y, nil
}

func (y *Yesod) PullFeeds(ctx context.Context) {
	configs, err := y.repo.ListFeedConfigNeedPull(ctx, nil,
		[]modelyesod.FeedConfigStatus{modelyesod.FeedConfigStatusActive},
		modelyesod.ListFeedOrderNextPull, time.Now(), 32) //nolint:gomnd // TODO
	if err != nil {
		logger.Errorf("%s", err.Error())
		return
	}
	for _, c := range configs {
		err = y.pullFeed.Publish(ctx, modelyesod.PullFeed{
			InternalID: c.ID,
			URL:        c.FeedURL,
			Source:     c.Source,
		})
		if err != nil {
			logger.Errorf("%s", err.Error())
			continue
		}
		err = y.repo.UpdateFeedConfigAsInQueue(ctx, c.ID)
		if err != nil {
			logger.Errorf("%s", err.Error())
		}
	}
}
