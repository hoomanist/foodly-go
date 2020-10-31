package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

// connection check
func ping(c *gin.Context){
  c.JSON(http.StatusOK, gin.H{
    "message": "pong",
  })
}
