package main

import (
  // "os"

  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/logger"
  "github.com/rs/zerolog"
  "github.com/rs/zerolog/log"
)

import "net/http"

func main() {

  zerolog.SetGlobalLevel(zerolog.InfoLevel)
  r := gin.New()

  r.GET("/", logger.SetLogger() , func(c *gin.Context) {

    log.Warn().
      Str("foo", "bar").
      Msg("/")

    c.JSON(http.StatusOK, gin.H{
      "message": "hello world",
    })
  })

  r.GET("/api/health", func(c *gin.Context) {

    log.Warn().
      Str("foo", "bar").
      Msg("/api/health")

    c.JSON(http.StatusOK, gin.H{
      "message": "ok",
    })
  })

  if gin.Mode() == "release" {
    r.Run(":80")
  } else {
    r.Run(":3000")
  }
}
