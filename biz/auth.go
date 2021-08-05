package protos

import (
	"github.com/kataras/iris/v12"
	"github.com/winjeg/go-commons/cryptos"

	"bytes"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

const (
	signParamName = "sign"
)

// EnableAuth 此方法应该尽早调用， 凡是需要鉴权的方法，都应该放在这个方法调用之后
// 不想启用鉴权，请勿使用此方法
func EnableAuth(app *iris.Application, key string, validTime time.Duration) {
	manager := &authManager{
		Key:   key,
		Valid: validTime / time.Second,
	}
	app.Use(manager.RequireChaosAuth())
}

type authManager struct {
	Key    string        `json:"key"`   // 双方约定的key
	Valid  time.Duration `json:"valid"` // 有效期为多久
	Enable bool          `json:"enable"`
}

// RequireChaosAuth generate a middle ware for those who may need a auth method.
func (a *authManager) RequireChaosAuth() func(iris.Context) {
	return func(ctx iris.Context) {
		sign := ctx.URLParam(signParamName)
		t := time.Now().Unix() / int64(a.Valid)
		m := ctx.URLParams()
		paramStr := joinUrlParams(m)
		// 下一阶段的也可以通过，原因：有可能请求发出去，这边处理的时候已经到下一个时间段了
		signCurrent := cryptos.Sha1([]byte(paramStr + fmt.Sprintf("time=%d&key=%s", t, a.Key)))
		signNext := cryptos.Sha1([]byte(paramStr + fmt.Sprintf("time=%d&key=%s", t+1, a.Key)))
		if strings.EqualFold(signCurrent, sign) || strings.EqualFold(signNext, sign) {
			ctx.Next()
			return
		} else {
			_, err := ctx.JSON(map[string]interface{}{
				"code": "401",
				"msg":  "sign doesn't match, maybe an attack",
			})
			if err != nil {
				log.Println(err.Error())
			}
			ctx.StopExecution()
			return
		}
	}
}

// sort and join url params except the sign
func joinUrlParams(m map[string]string) string {
	// sort the params
	keys := make([]string, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	b := bytes.Buffer{}
	for _, v := range keys {
		if !strings.EqualFold(signParamName, v) {
			b.WriteString(fmt.Sprintf("%s=%s&", v, m[v]))
		}
	}
	return b.String()
}
