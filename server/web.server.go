package server

import (
	"io/ioutil"
	"log"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)
type engine struct {
	router map[string]HandlerFunc
}

func (t *engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	t.router[key] = handler
}
func (t *engine) GET(pattern string, handler HandlerFunc) {
	t.addRoute("GET", pattern, handler)
}

func (t *engine) POST(pattern string, handler HandlerFunc) {
	t.addRoute("POST", pattern, handler)
}

func (t *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handle, ok := t.router[key]; ok {
		handle(w, r)

	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
	}
}

func serve() {
	engine := new(engine)
	engine.router = make(map[string]HandlerFunc)

	engine.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("test"))
	})

	engine.POST("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		log.Println(r.Header)
		s, _ := ioutil.ReadAll(r.Body)

		w.Write(s)
	})

	http.ListenAndServe(":9999", engine)
}
