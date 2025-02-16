package data

import "github.com/savsgio/atreugo/v11"

type Data struct {
	ctx *atreugo.RequestCtx
}

func Init(ctx *atreugo.RequestCtx) error {
	ctx.SetUserValue("data", &Data{
		ctx: ctx,
	})

	return ctx.Next()
}

func Get(ctx *atreugo.RequestCtx) *Data {
	return ctx.UserValue("data").(*Data)
}
