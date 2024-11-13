package gee

import "net/http"

type Handlerfunc func(*Context)
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRoute()}
}
func (engine *Engine) AddRoute(method string, pattern string, handle Handlerfunc) {
	engine.router.addRoute(method, pattern, handle)
}
func (engine *Engine) GET(pattern string, handle Handlerfunc) {
	engine.AddRoute("GET", pattern, handle)
}
func (engine *Engine) POST(pattern string, handle Handlerfunc) {
	engine.AddRoute("POST", pattern, handle)
}
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := Newcontext(w, req)
	engine.router.handle(c)
}
