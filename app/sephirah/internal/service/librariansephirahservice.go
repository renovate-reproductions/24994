package service

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizangela"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizchesed"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biznetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type LibrarianSephirahServiceService struct {
	pb.UnimplementedLibrarianSephirahServiceServer

	t        *biztiphereth.Tiphereth
	g        *bizgebura.Gebura
	b        *bizbinah.Binah
	y        *bizyesod.Yesod
	n        *biznetzach.Netzach
	c        *bizchesed.Chesed
	app      *libapp.Settings
	auth     *libauth.Auth
	authFunc func(context.Context) (context.Context, error)
}

func NewLibrarianSephirahServiceService(
	_ *bizangela.Angela,
	t *biztiphereth.Tiphereth,
	g *bizgebura.Gebura,
	b *bizbinah.Binah,
	y *bizyesod.Yesod,
	n *biznetzach.Netzach,
	c *bizchesed.Chesed,
	app *libapp.Settings,
	auth *libauth.Auth,
	authFunc func(context.Context) (context.Context, error),
) pb.LibrarianSephirahServiceServer {
	if enable, err := app.GetEnvBool(libapp.EnvCreateAdmin); err == nil && enable {
		t.CreateDefaultAdmin(context.Background(), &modeltiphereth.User{
			ID:       0,
			UserName: "admin",
			PassWord: "admin",
			Type:     0,
			Status:   0,
		})
	}
	return &LibrarianSephirahServiceService{
		UnimplementedLibrarianSephirahServiceServer: pb.UnimplementedLibrarianSephirahServiceServer{},
		t:        t,
		g:        g,
		b:        b,
		y:        y,
		n:        n,
		c:        c,
		app:      app,
		auth:     auth,
		authFunc: authFunc,
	}
}

func (s *LibrarianSephirahServiceService) GetServerInformation(_ context.Context,
	_ *pb.GetServerInformationRequest) (*pb.GetServerInformationResponse, error) {
	return &pb.GetServerInformationResponse{
		ServerBinarySummary: &pb.ServerBinarySummary{
			SourceCodeAddress: s.app.SourceCodeAddress,
			BuildVersion:      s.app.Version,
			BuildDate:         s.app.BuildDate,
		},
		ProtocolSummary: &pb.ServerProtocolSummary{
			Version: s.app.ProtoVersion,
		},
		CurrentTime: timestamppb.New(time.Now()),
	}, nil
}
