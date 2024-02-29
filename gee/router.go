package gee

import "net/http"

type router struct {
  handlers map[string]HandlerFunc
}

func newRouter() *router {
  return &router{
    make(map[string]HandlerFunc),
  }
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
  key := method + "-" + pattern
  r.handlers[key] = handler
}

func (r *router) handle(ctx *Context) {
  key := ctx.Method + "-" + ctx.Path
  if handler, ok := r.handlers[key]; ok{
    handler(ctx) 
  } else {
    ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
  }
}
