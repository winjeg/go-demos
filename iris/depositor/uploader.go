package depositor

import (
	"github.com/winjeg/hamster/router"
	"github.com/winjeg/hamster/store"
	"github.com/winjeg/hamster/zip"
)

var liteStore = store.NewLiteStore()

func uploadAndRegister(fileName, author, dir, route string) error {
	// 1. unzio to target dir
	err := zip.Unzip(fileName, dir, route)
	if err != nil {
		return err
	}
	// 3. store it to db
	err = liteStore.Store(author, dir, route)
	if err != nil {
		router.RefreshRouter()
		return err
	}
	// 2. register route
	router.RegisterStaticPath(route, dir)
	return nil
}
