package main


import (
  "github.com/gin-gonic/gin"
  "net/http"
)


func Vote(c *gin.Context){
  var rest Restaurant
  var food Food
  FoodName := c.Query("food")
  RestaurantName := c.Query("restaurant")
  db.Where(&Restaurant{username: RestaurantName}).First(&rest)
  db.Where(&Food{restaurant: rest, name: FoodName}).First(&food)
  if c.Query("dir") == "up" {
    food.vote += 1
  } else if c.Query("dir") == "down" {
    food.vote -= 1
  }
  db.Save(&food)
  c.JSON(http.StatusOK, gin.H{
    "status": "ok",
  })
}

func GetVotes(c *gin.Context){
  var rest Restaurant
  var food Food
  FoodName := c.Query("food")
  RestaurantName := c.Query("restaurant")
  db.Where(&Restaurant{username: RestaurantName}).First(&rest)
  db.Where(&Food{restaurant: rest, name: FoodName}).First(&food)
  c.JSON(http.StatusOK, gin.H{
    "votes": food.vote,
  })
}
