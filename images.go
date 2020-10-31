package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func UploadImages(c *gin.Context){
  file, _ := c.FormFile("file")
  dst := "uploads/" + file.Filename
	c.SaveUploadedFile(file, dst)
  c.JSON(http.StatusOK,gin.H{
    "name": file.Filename,
  })
}

func GetImage(c *gin.Context){
  name := c.Param("name")
  c.File("uploads/"+name)
}
