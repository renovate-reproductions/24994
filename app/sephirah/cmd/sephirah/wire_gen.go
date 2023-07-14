// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizangela"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizchesed"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biznetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	client2 "github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/data"
	"github.com/tuihub/librarian/app/sephirah/internal/service"
	"github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/server"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(sephirah_Server *conf.Sephirah_Server, sephirah_Data *conf.Sephirah_Data, auth *conf.Auth, mq *conf.MQ, cache *conf.Cache, settings *libapp.Settings) (*kratos.App, func(), error) {
	libauthAuth, err := libauth.NewAuth(auth)
	if err != nil {
		return nil, nil, err
	}
	libmqMQ, cleanup, err := libmq.NewMQ(mq, settings)
	if err != nil {
		return nil, nil, err
	}
	entClient, cleanup2, err := data.NewSQLClient(sephirah_Data)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	dataData := data.NewData(entClient)
	angelaRepo := data.NewAngelaRepo(dataData)
	geburaRepo := data.NewGeburaRepo(dataData)
	librarianMapperServiceClient, err := client.NewMapperClient()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	librarianPorterServiceClient, err := client.NewPorterClient()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	librarianSearcherServiceClient, err := client.NewSearcherClient()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	searcher := client2.NewSearcher(librarianSearcherServiceClient)
	angelaBase, err := bizangela.NewAngelaBase(angelaRepo, geburaRepo, librarianMapperServiceClient, librarianPorterServiceClient, searcher)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	topic := bizangela.NewPullSteamAppTopic(angelaBase)
	libmqTopic := bizangela.NewPullSteamAccountAppRelationTopic(angelaBase, topic)
	topic2 := bizangela.NewPullAccountTopic(angelaBase, libmqTopic)
	netzachRepo := data.NewNetzachRepo(dataData)
	store, err := libcache.NewStore(cache)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	libcacheMap := bizangela.NewNotifyFlowCache(netzachRepo, store)
	map2 := bizangela.NewFeedToNotifyFlowMap(netzachRepo, store)
	map3 := bizangela.NewNotifyTargetCache(netzachRepo, store)
	topic3 := bizangela.NewNotifyPushTopic(angelaBase, map3)
	topic4 := bizangela.NewNotifyRouterTopic(angelaBase, libcacheMap, map2, topic3)
	topic5 := bizangela.NewParseFeedItemDigestTopic(angelaBase)
	topic6 := bizangela.NewPullFeedTopic(angelaBase, topic4, topic5)
	angela, err := bizangela.NewAngela(libmqMQ, topic2, libmqTopic, topic, topic6, topic4, topic3, topic5)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tipherethRepo := data.NewTipherethRepo(dataData)
	tiphereth, err := biztiphereth.NewTiphereth(tipherethRepo, libauthAuth, librarianMapperServiceClient, searcher, topic2)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	gebura := bizgebura.NewGebura(geburaRepo, libauthAuth, librarianMapperServiceClient, searcher)
	controlBlock := bizbinah.NewControlBlock(libauthAuth)
	binah := bizbinah.NewBinah(controlBlock, libauthAuth, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	yesodRepo := data.NewYesodRepo(dataData)
	cron := libcron.NewCron()
	yesod, err := bizyesod.NewYesod(yesodRepo, cron, librarianMapperServiceClient, searcher, topic6)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	netzach := biznetzach.NewNetzach(netzachRepo, searcher, map2, libcacheMap, map3)
	chesedRepo := data.NewChesedRepo(dataData)
	librarianMinerServiceClient, err := client.NewMinerClient()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	map4 := bizchesed.NewImageCache(store)
	chesed, err := bizchesed.NewChesed(chesedRepo, cron, librarianPorterServiceClient, searcher, librarianMinerServiceClient, controlBlock, map4)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	v := server.NewAuthMiddleware(libauthAuth)
	librarianSephirahServiceServer := service.NewLibrarianSephirahServiceService(angela, tiphereth, gebura, binah, yesod, netzach, chesed, settings, libauthAuth, v)
	grpcServer := server.NewGRPCServer(sephirah_Server, libauthAuth, librarianSephirahServiceServer, settings)
	httpServer := server.NewGrpcWebServer(grpcServer, sephirah_Server, libauthAuth, settings)
	app := newApp(grpcServer, httpServer, libmqMQ, cron)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
