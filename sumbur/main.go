package main

import (
	"os"
	"sumbur/views"
	"sumbur/views/http_errors"
	"sumbur/views/static"

	"github.com/savsgio/atreugo/v11"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Debug bool

	Server atreugo.Config
}

func main() {
	// Configuration

	config := Config{
		Server: atreugo.Config{
			NoDefaultContentType: true,

			MethodNotAllowedView: http_errors.NotFoundView,
			NotFoundView:         http_errors.NotFoundView,
			PanicView:            http_errors.PanicView,
		},
	}

	file, err := os.ReadFile(os.Getenv("SUMBUR_CONFIG"))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	// Server

	server := atreugo.New(config.Server)

	server.GET("/", func(ctx *atreugo.RequestCtx) error {
		views.WritePage(ctx, &views.BasePage{})
		return nil
	})

	// Static files

	static.Debug = config.Debug

	server.Static("/static", "static")
	server.Static("/static-stamp/{stamp}", "static")

	// Test routes

	server.GET("/panic", func(ctx *atreugo.RequestCtx) error {
		panic("Тестовая ошибка")
	})

	// Run

	err = server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
