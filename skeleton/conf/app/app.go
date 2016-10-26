package app

import (
	mcache "github.com/go-macaron/cache"
	"github.com/go-macaron/gzip"
	"github.com/go-macaron/i18n"
	"github.com/go-macaron/jade"
	"github.com/go-macaron/session"
	"{{.AppPath}}/conf"
	"{{.AppPath}}/lib/cache"
	"{{.AppPath}}/lib/context"
	"{{.AppPath}}/lib/template"
	"gopkg.in/macaron.v1"
)

func SetupMiddlewares(app *macaron.Macaron) {
	app.Use(macaron.Logger())
	app.Use(macaron.Recovery())
	app.Use(gzip.Gziper())
	app.Use(macaron.Static("public"))
	app.Use(i18n.I18n(i18n.Options{
		Directory: "locale",
		Langs:     []string{"pt-BR", "en-US"},
		Names:     []string{"Português do Brasil", "American English"},
	}))
	app.Use(jade.Renderer(jade.Options{
		Directory: "public/templates",
		Funcs:     template.FuncMaps(),
	}))
	app.Use(macaron.Renderer())
	app.Use(mcache.Cacher(
		cache.Option(conf.Cfg.Section("").Key("cache_adapter").Value()),
	))
	app.Use(session.Sessioner())
	app.Use(context.Contexter())
}

func SetupRoutes(app *macaron.Macaron) {
	app.Get("", func() string {
		return "Mercurius Works!"
	})
}