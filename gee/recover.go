package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)


func Recover() HandlerFunc {
  return func(ctx *Context) {
    defer func ()  {
      if err := recover(); err != nil {
        message := fmt.Sprintf("%s", err)
        log.Printf("%s\n\n", trace(message))
        ctx.Fail(http.StatusInternalServerError, "Internal Server Error")
      }
    }()

    ctx.Next()
  }
}

func trace(message string) string {
  var pcs [32]uintptr
  n := runtime.Callers(3, pcs[:])

  var str strings.Builder
  str.WriteString(message + "\nTraceback:")
  for _, pc := range pcs[:n] {
    fn := runtime.FuncForPC(pc)
    file, line := fn.FileLine(pc)
    str.WriteString(fmt.Sprintf("\n\t%s:\t%d", file, line))
  }

  return str.String()
}
