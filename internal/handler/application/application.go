package application

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/unblvvv/h-www-server/internal/handler"
	"go.uber.org/fx"
)

type Opts struct {
	fx.In

	Api huma.API
}

var FxModule = fx.Provide(
	handler.AsFxRoute(NewUpdateHandler),
	handler.AsFxRoute(NewListHandler),
	handler.AsFxRoute(NewCreateHandler),
)
