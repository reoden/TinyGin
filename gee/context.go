package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Method     string
	Path       string
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
  return &Context{
    Writer: w,
    Req: req,
    Method: req.Method,
    Path: req.URL.Path,
  }
}

func (ctx *Context) PostForm(key string) (value string) {
  value = ctx.Req.FormValue(key)
  return
}

func (ctx *Context) Query(key string) string {
  return ctx.Req.URL.Query().Get(key)
}

func (ctx *Context) Status(code int) {
  ctx.StatusCode = code
  ctx.Writer.WriteHeader(code)
}

func (ctx *Context) SetHeader(key, value string) {
  ctx.Writer.Header().Set(key, value)
}

func (ctx *Context) String(code int, format string, values ...interface{}) {
  ctx.SetHeader("Content-Type", "text/plain")
  ctx.Status(code)
  ctx.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (ctx *Context) JSON(code int, obj interface{}) {
  ctx.SetHeader("Content-Type", "application/json")
  ctx.Status(code)
  encoder := json.NewEncoder(ctx.Writer)
  if err := encoder.Encode(obj); err != nil {
    http.Error(ctx.Writer, err.Error(), 500)
  }
}

func (ctx *Context) Data(code int, data []byte) {
  ctx.Status(code)
  ctx.Writer.Write(data)
}

func (ctx *Context) HTML(code int, html string) {
  ctx.SetHeader("Content-Type", "text/html")
  ctx.Status(code)
  ctx.Writer.Write([]byte(html))
}
