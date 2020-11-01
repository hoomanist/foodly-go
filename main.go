package main

import (
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

// database cursor
var db *gorm.DB

func main(){
  //// connect to sqlite database
  db, err := gorm.Open(sqlite.Open("foodly.db"), &gorm.Config{})
  if err != nil  {
    panic("failed to connect database")
  }
  //// migrate models
  db.AutoMigrate(&Restaurant{})
  db.AutoMigrate(&User{})
  db.AutoMigrate(&Food{})
  //// create router and assign routes
  router := gin.Default()
  // misc
  router.GET("/ping", ping)
  // auth
  router.POST("/register/user", RegisterUser)
  router.POST("/login/user", LoginUser)
  router.POST("/register/restaurant", CreateRestaurant)
  router.POST("/login/restaurant", LoginRestaurant)
  // food
  router.GET("/query/food", GetFoodData)
  router.GET("/query/foods", GetFoodByRestaurantName)
  router.POST("/submit/food", CreateFood)
  // vote
  router.GET("/query/vote", GetVotes)
  router.POST("/submit/vote", Vote)
  // comment
  router.GET("/query/comments", QueryComments)
  router.POST("/submit/comments", SubmitComment)
  // images
  router.POST("/upload/images", UploadImages)
  router.GET("/images/:name", GetImage)
  //// start the web server
  router.Run(":5000")
}
