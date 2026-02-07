package middlewares

import (
  "log"

  "github.com/gin-gonic/gin"
)

func DummyLogger() gin.HandlerFunc {
  return func(c *gin.Context) {
    log.Printf("JT Dummy logger middleware executed before request")
    c.Next()
    log.Printf("JT Dummy logger middleware executed after request")
  }
}


