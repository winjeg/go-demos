package zip

import (
	"archive/zip"
	"fmt"
	"github.com/winjeg/hamster/log"
	"github.com/winjeg/hamster/util"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func Unzip(fileName, dest, path string) error {
	// Open a zip archive for reading.
	target := util.GetCurrentPath(dest)
	r, err := zip.OpenReader(fileName)
	if err != nil {
		log.LogErr(err)
		return err
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			log.LogErr(err)
			return err
		}
		data, fErr := ioutil.ReadAll(rc)
		log.LogErr(fErr)
		name := target + "/" + f.Name
		writeErr := writeFile(name, data, path)
		log.LogErr(rc.Close())
		log.LogErr(writeErr)
	}
	return nil
}

func writeFile(filename string, data []byte, path string) error {
	if !util.CheckFileIsExist(filename) {
		idx := strings.LastIndex(filename, "/")
		path := filename[:idx]
		log.LogErr(os.MkdirAll(path, 0755))
		f, err := os.Create(filename)
		log.LogErr(err)
		log.LogErr(f.Close())
		if err != nil {
			return err
		}
	}
	var ioErr error
	if isStaticFile(filename) {
		fmt.Println(filename)

		ioErr = ioutil.WriteFile(filename, util.FixUrls(data, findPath(filename, path)), 0666)
	} else {
		ioErr = ioutil.WriteFile(filename, data, 0666)
	}
	log.LogErr(ioErr)
	return ioErr
}

const webFileExpr = `[.](js|css|html)$`

func findPath(name, path string) string {

	if contains(name, "init.js", "sitemap.js", "jquery") {
		return path
	}

	idx := strings.Index(name, path)
	idx2 := strings.LastIndex(name, "/")
	return name[idx:idx2]
}

func isStaticFile(name string) bool {
	reg, _ := regexp.Compile(webFileExpr)
	return reg.Match([]byte(name))
}

func contains(content string, arr ...string) bool {
	for _, v := range arr {
		if strings.Contains(content, v) {
			return true
		}
	}
	return false
}
