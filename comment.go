package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func SubmitComment(c *gin.Context){
  msg := c.Query("msg")
  FoodName := c.Query("food")
  RestaurantName := c.Query("rest")
  token := c.Query("token")
  var user User
  var rest Restaurant
  var food Food
  db.Where(&Restaurant{Name: RestaurantName}).First(&rest)
  db.Where(&Food{name: FoodName, restaurant: rest}).First(&food)
  db.Where(&User{token: token}).First(&user)
  if user.username == "" {
    c.JSON(http.StatusNotFound, gin.H{
      "error": "token not found. it is invalid perhaps your login was not properly.",
    })
  }
  db.Create(&Comment{
    msg: msg,
    username: user.token,
    food: food,
  })
  c.JSON(http.StatusOK, gin.H{
    "status": "ok",
  })
}

func QueryComments(c *gin.Context){
  FoodName := c.Query("food")
  RestaurantName := c.Query("rest")
  var rest Restaurant
  var food Food
  var comments []Comment
  db.Where(&Restaurant{Name: RestaurantName}).First(&rest)
  db.Where(&Food{name: FoodName, restaurant: rest}).First(&food)
  db.Where(&Comment{food: food}).Find(&comments)
  c.JSON(http.StatusOK, comments)
}