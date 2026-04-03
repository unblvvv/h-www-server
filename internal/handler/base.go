package handler

import (
	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/fx"
)

type registerRoute interface {
	GetMeta() huma.Operation
	Register(api huma.API)
}

func AsFxRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(registerRoute)),
		fx.ResultTags(`group:"api-routes"`),
	)
}

type RegisterRoutesOpts struct {
	fx.In

	Api    huma.API
	Routes []registerRoute `group:"api-routes"`
}

func RegisterRoutes(opts RegisterRoutesOpts) {
	for _, r := range opts.Routes {
		r.Register(opts.Api)
	}
}
