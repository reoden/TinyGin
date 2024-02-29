package main

import "github.com/gin-gonic/gin"

func ginExample() {
  r := gin.Default()

  r.GET("/", func(ctx *gin.Context) {
    ctx.String(200, "hello world")
  })

  r.Run(":9999")
}
