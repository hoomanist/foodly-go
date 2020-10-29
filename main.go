package main

import (
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

// database cursor
var db *gorm.DB

func main(){
  // connect to sqlite database
  db, err := gorm.Open(sqlite.Open("foodly.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  // migrate models
  db.AutoMigrate(&Restaurant{})
  db.AutoMigrate(&User{})
  // create router and assign routes
  router := gin.Default()
  router.GET("/ping", ping)
  router.POST("/register/user", RegisterUser)
  // start the web server
  router.Run()
}
