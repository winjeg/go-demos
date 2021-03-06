package depositor

import (
	"github.com/kataras/iris"
	"github.com/winjeg/hamster/conf"
	"github.com/winjeg/hamster/log"
	"io"
	"os"
	"strings"
)

var config = conf.GetConf()

func RegisterDepositor(app *iris.Application) {
	party := app.Party(config.Web.Path)
	{
		party.Post("/api/upload", iris.LimitRequestBodySize(maxSize*(1<<20)), upload)
	}
}

const maxSize = 30 // 30MB

func upload(ctx iris.Context) {
	file, info, err := ctx.FormFile("uploadFile")
	author := ctx.FormValue("author")
	path := ctx.FormValue("path")
	if strings.EqualFold("/", path) || strings.EqualFold("/static", path) {
		ctx.Text("系统路径不被允许")
		return
	}
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, err := ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		log.LogErr(err)
		return
	}
	defer file.Close()
	fileName := config.Web.UploadDir + "/" + info.Filename
	out, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, err := ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		log.LogErr(err)
		return
	}
	defer out.Close()
	_, wErr := io.Copy(out, file)
	log.LogErr(wErr)
	uploadAndRegister(fileName, author, path, path)
	ctx.Redirect("/")
}
