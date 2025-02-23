package bizbinah

import (
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelbinah"
	"github.com/tuihub/librarian/internal/lib/libauth"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
)

type BinahRepo interface {
}

type Binah struct {
	callback *modelbinah.ControlBlock
	auth     *libauth.Auth
	mapper   mapper.LibrarianMapperServiceClient
	porter   porter.LibrarianPorterServiceClient
	searcher searcher.LibrarianSearcherServiceClient
}

func NewBinah(
	callback *modelbinah.ControlBlock,
	auth *libauth.Auth,
	mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient searcher.LibrarianSearcherServiceClient,
) *Binah {
	return &Binah{
		callback: callback,
		auth:     auth,
		mapper:   mClient,
		porter:   pClient,
		searcher: sClient,
	}
}

func NewControlBlock(a *libauth.Auth) *modelbinah.ControlBlock {
	return modelbinah.NewControlBlock(a)
}
