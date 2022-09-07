// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/tuihub/librarian/app/searcher/internal/biz"
	"github.com/tuihub/librarian/app/searcher/internal/data"
	"github.com/tuihub/librarian/app/searcher/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/protos/pkg/librarian/searcher/v1"
)

// Injectors from wire.go:

func NewSearcherService(searcher_Data *conf.Searcher_Data) (v1.LibrarianSearcherServiceServer, func(), error) {
	dataData, cleanup, err := data.NewData(searcher_Data)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData)
	greeterUseCase := biz.NewGreeterUseCase(greeterRepo)
	librarianSearcherServiceServer := service.NewLibrarianSearcherServiceService(greeterUseCase)
	return librarianSearcherServiceServer, func() {
		cleanup()
	}, nil
}
