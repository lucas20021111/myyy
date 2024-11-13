package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]Handlerfunc
}

func newRoute() *router {
	return &router{handlers: make(map[string]Handlerfunc)}
}
func (r *router) addRoute(method string, pattern string, handler Handlerfunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND", c.Path)
	}

}
