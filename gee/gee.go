package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type RouteGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouteGroup
	engine      *Engine
}

type Engine struct {
	*RouteGroup
	router *router
	groups  []*RouteGroup
}

func New() *Engine {
  engine := &Engine{router: newRouter()}
  engine.RouteGroup = &RouteGroup{engine: engine}
  engine.groups = []*RouteGroup{engine.RouteGroup}
  return engine
}

func (group *RouteGroup) Group(prefix string) *RouteGroup {
  engine := group.engine  
  newGroup := &RouteGroup{
    prefix: group.prefix + prefix,
    parent: group,
    engine: engine,
  }
  
  engine.groups = append(engine.groups, newGroup)
  return newGroup
}

func (group *RouteGroup) addRoute(method string, comp string, handler HandlerFunc) {
  pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouteGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouteGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine *Engine) RUN(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := newContext(w, req)
	engine.router.handle(ctx)
}
