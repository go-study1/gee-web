package gee

import (
	"net/http"
)

type Engine struct {
	router *route
}

func New() *Engine {
	return &Engine{router: NewRoute()}
}

func (engine *Engine) GET(pattern string, handle HandleFunc) {
	engine.router.addRoute("GET", pattern, handle)
}

func (engine *Engine) POST(pattern string, handle HandleFunc) {
	engine.router.addRoute("POST", pattern, handle)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

func (engine *Engine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}
