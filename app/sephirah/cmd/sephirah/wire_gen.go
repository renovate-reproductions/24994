// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizangela"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/data"
	"github.com/tuihub/librarian/app/sephirah/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/server"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(sephirah_Server *conf.Sephirah_Server, sephirah_Data *conf.Sephirah_Data, auth *conf.Auth, mq *conf.MQ, settings *libapp.Settings) (*kratos.App, func(), error) {
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
	tipherethRepo := data.NewTipherethRepo(dataData)
	geburaRepo := data.NewGeburaRepo(dataData)
	yesodRepo := data.NewYesodRepo(dataData)
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
	angelaBase, err := bizangela.NewAngelaBase(tipherethRepo, geburaRepo, yesodRepo, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	topicImpl := bizangela.NewPullSteamAppTopic(angelaBase)
	libmqTopicImpl := bizangela.NewPullSteamAccountAppRelationTopic(angelaBase, topicImpl)
	topicImpl2 := bizangela.NewPullAccountTopic(angelaBase, libmqTopicImpl)
	topicImpl3 := bizangela.NewPullFeedTopic(angelaBase)
	angela, err := bizangela.NewAngela(libmqMQ, topicImpl2, libmqTopicImpl, topicImpl, topicImpl3)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tiphereth, err := biztiphereth.NewTiphereth(tipherethRepo, libauthAuth, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient, topicImpl2)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	callbackControlBlock := bizbinah.NewCallbackControl()
	gebura := bizgebura.NewGebura(geburaRepo, libauthAuth, callbackControlBlock, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	binah := bizbinah.NewBinah(callbackControlBlock, libauthAuth, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	cron := libcron.NewCron()
	yesod, err := bizyesod.NewYesod(yesodRepo, cron, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient, topicImpl3)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	v := server.NewAuthMiddleware(libauthAuth)
	librarianSephirahServiceServer := service.NewLibrarianSephirahServiceService(angela, tiphereth, gebura, binah, yesod, settings, v)
	grpcServer := server.NewGRPCServer(sephirah_Server, libauthAuth, librarianSephirahServiceServer, settings)
	httpServer := server.NewGrpcWebServer(grpcServer, sephirah_Server, libauthAuth, settings)
	app := newApp(grpcServer, httpServer, libmqMQ, cron)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
