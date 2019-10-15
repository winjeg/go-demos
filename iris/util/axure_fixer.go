package util

import (
	"github.com/winjeg/go-commons/str"
	"regexp"
	"strings"
)

const (
	staticFileExpr = `[a-zA-Z\/\-_0-9.]+[.](jpg|gif|bmg|png|svg|css|js|html)('|")`
)

func FixUrls(data []byte, path string) []byte {
	// 字符串以 .css, .js, .html， .png, .svg, .jpg 结尾, 切前面没有域名
	exp, _ := regexp.Compile(staticFileExpr)
	return exp.ReplaceAllFunc(data, func(bytes []byte) []byte {
		if len(bytes) < 1 {
			return bytes
		}
		seg := str.FromBytes(bytes)
		if strings.Index(seg, "../") == 0 {
			idx := strings.LastIndex(seg, "../")

			result := seg[:idx+3] + path[1:] + "/" + seg[idx+3:]

			return str.ToBytes(result)
		} else {
			if seg[0] == '/' {
				seg = seg[1:]
			}
			return str.ToBytes(path + "/" + seg)
		}
	})
}

func isRemoteAddr(str string) bool {
	if len(str) < 3 {
		return false
	}
	if len(str) < 15 {
		return strings.Contains(str, "://")
	}
	return strings.Contains(str[:15], "://")
}
