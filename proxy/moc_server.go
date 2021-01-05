package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"time"
)

var server *http.Server

const defaultPort = ":8001"

func resolvePort() string {
	args := os.Args
	var port = defaultPort
	if len(args) > 1 {
		portStr := args[1]
		p, err := strconv.Atoi(portStr)
		if err == nil && p > 0 && p < 65535 {
			port = fmt.Sprintf(":%d", p)
		}
	}
	return port
}

func main() {
	port := resolvePort()
	// 一个通知退出的chan
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})

	server = &http.Server{
		Addr:         port,
		WriteTimeout: time.Minute * 60,
		Handler:      mux,
	}

	go func() {
		// 接收退出信号
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Close server:", err)
		}
	}()

	log.Printf("Starting v3 httpserver on %s\n", port)
	err := server.ListenAndServe()
	if err != nil {
		// 正常退出
		if err == http.ErrServerClosed {
			log.Fatal("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected", err)
		}
	}
	log.Fatal("Server exited")
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.RequestURI()
	bodyData, _ := ioutil.ReadAll(r.Body)
	r.Body.Read(bodyData)
	var bodyMap = make(map[string]interface{}, 16)
	err := json.Unmarshal(bodyData, &bodyMap)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error parsing body data: %s", err.Error())))
		return
	}
	var request string
	if v, ok := bodyMap["request"]; ok {
		d, _ := json.Marshal(v)
		request = string(d)
	}
	log.Println(fmt.Sprintf("uri:%s\tthread:%d\t request:%s", uri, GetGID(), request))
	if v, ok := bodyMap["timeout"]; ok {
		if timeout, ok := v.(float64); ok {
			time.Sleep(time.Millisecond * time.Duration(timeout))
		}
	}
	if v, ok := bodyMap["response"]; ok {
		r, e := json.Marshal(v)
		if e != nil {
			log.Println(e.Error())
		}
		w.Write(r)
		return
	}
	w.Write([]byte("ok"))
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
