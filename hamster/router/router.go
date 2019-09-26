package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/kataras/iris"
	"github.com/winjeg/hamster/conf"
	"github.com/winjeg/hamster/log"
	"github.com/winjeg/hamster/store"
)

var (
	application *iris.Application
	config      = conf.GetConf()
	lite        = store.NewLiteStore()
)

func SetApp(app *iris.Application) {
	application = app
	app.RegisterView(iris.HTML("./static", ".html"))
	svc := app.Party(config.Web.Path)
	{
		svc.Get("/", serveIndex)
	}
	RegisterStaticPath("/static", "./static")
	sites := lite.GetAll()
	for _, site := range sites {
		RegisterStaticPath(*site.Mapping, *site.Dir)
	}
}

func RegisterStaticPath(route, path string) {
	if route[0] == '/' {
		route = route[1:]
	}
	if path[0] == '/' {
		path = "." + path
	}
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	rt := config.Web.Path + route
	application.StaticWeb(rt, path)
	application.Get(rt, func(c iris.Context) {
		d, _ := ioutil.ReadFile(path + "/index.html")
		_, err := c.HTML(string(d))
		log.LogErr(err)
	})
	log.LogErr(application.RefreshRouter())
	printRoutes()
}

func printRoutes() {
	r := application.GetRoutes()
	d, _ := json.Marshal(r)
	fmt.Println(string(d))
}

func serveIndex(c iris.Context) {
	sites := lite.GetAll()
	c.ViewData("sites", sites)
	log.LogErr(c.View("index.html"))
}

func RefreshRouter() {
	log.LogErr(application.RefreshRouter())
}
