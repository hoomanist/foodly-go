package main

import (
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

// database cursor
var db *gorm.DB

func SetupRouter() *gin.Engine {

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
  return router

}

func main(){
  var err error
  //// connect to sqlite database
  db, err = gorm.Open(mysql.Open("hooman:hooman86@tcp(127.0.0.1:3306)/foodly?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
  if err != nil  {
    panic("failed to connect database")
  }
  //// migrate models
  db.AutoMigrate(&Restaurant{})
  db.AutoMigrate(&User{})
  db.AutoMigrate(&Food{})
  db.AutoMigrate(&Comment{})
  //// start the web server
  router := SetupRouter()
  router.Run(":5000")
}
