package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	r.GET("/hello", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "hello %s, you are at %s\n", ctx.Query("name"), ctx.Path)
	})
  
  r.GET("/hello/:name", func(ctx *gee.Context) {
    ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
  })

  r.GET("/assets/*filePath", func(ctx *gee.Context) {
    ctx.JSON(http.StatusOK, gee.H{"filePath": ctx.Param("filePath")})
  })

	r.RUN(":9999")
}
