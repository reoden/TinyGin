package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.Default()
  r.GET("/", func(ctx *gee.Context) {
    ctx.String(http.StatusOK, "Hello World\n")
  })

  r.GET("/panic", func(ctx *gee.Context) {
    names := []string{"reoden"}
    ctx.String(http.StatusOK, names[100])
  })

  r.RUN(":9999")
}
