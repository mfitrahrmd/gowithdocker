package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
  r := gin.Default()

  r.GET("/", func(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
      "name": "Rama",
      "age": 22,
    })
  })

  http.ListenAndServe(":3000", r)
}
