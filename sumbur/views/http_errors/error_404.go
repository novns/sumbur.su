package http_errors

import (
	"sumbur/data"
	"sumbur/views"

	"github.com/savsgio/atreugo/v11"
)

type NotFound struct {
	*views.BasePage

	path []byte
}

func NotFoundView(ctx *atreugo.RequestCtx) error {
	data.Init(ctx)

	ctx.SetStatusCode(404)
	views.WritePage(ctx, data.Get(ctx), &NotFound{path: ctx.Path()})

	return nil
}
