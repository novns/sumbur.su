package http_errors

import (
	"runtime/debug"
	"sumbur/data"
	"sumbur/views"

	"github.com/savsgio/atreugo/v11"
)

type Panic struct {
	*views.BasePage

	err   any
	stack []byte
}

func PanicView(ctx *atreugo.RequestCtx, err any) {
	data.Init(ctx)

	ctx.SetStatusCode(500)
	ctx.ResetBody()

	views.WritePage(ctx, data.Get(ctx), &Panic{err: err, stack: debug.Stack()})
}
