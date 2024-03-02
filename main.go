package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee.HandlerFunc {
	return func(ctx *gee.Context) {
		t := time.Now()
		ctx.Fail(500, "Internet Server Error")
		log.Printf("[%d] %s in %v for group v2\n", ctx.StatusCode, ctx.Req.RequestURI, t)
	}
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	v2 := r.Group("/v2")
  v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(ctx *gee.Context) {
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
		})
	}

	r.RUN(":9999")
}
