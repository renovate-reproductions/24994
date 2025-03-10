//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	mapperService "github.com/tuihub/librarian/app/mapper/pkg/service"
	minerService "github.com/tuihub/librarian/app/miner/pkg/service"
	porterService "github.com/tuihub/librarian/app/porter/pkg/service"
	searcherService "github.com/tuihub/librarian/app/searcher/pkg/service"
	sephirahService "github.com/tuihub/librarian/app/sephirah/pkg/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/inprocgrpc"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/server"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(
	*conf.Librarian_EnableServiceDiscovery,
	*conf.Sephirah_Server,
	*conf.Sephirah_Data,
	*conf.Mapper_Data,
	*conf.Searcher_Data,
	*conf.Porter_Data,
	*conf.Miner_Data,
	*conf.Auth,
	*conf.MQ,
	*conf.Cache,
	*libapp.Settings,
) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			sephirahService.ProviderSet,
			mapperService.ProviderSet,
			searcherService.ProviderSet,
			porterService.ProviderSet,
			minerService.ProviderSet,
			server.ProviderSet,
			inprocgrpc.ProviderSet,
			libauth.ProviderSet,
			libmq.ProviderSet,
			libcron.ProviderSet,
			libcache.ProviderSet,
			ProviderSet,
		),
	)
}
