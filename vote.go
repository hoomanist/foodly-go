package main


import (
  "github.com/gin-gonic/gin"
  "net/http"
)


func Vote(c *gin.Context){
  var (
    rest Restaurant
    food Food
  )
  FoodName := c.PostForm("food")
  RestaurantName := c.PostForm("restaurant")
  db.Where(&Restaurant{Username: RestaurantName}).First(&rest)
  db.Where(&Food{Restaurant: rest, Name: FoodName}).First(&food)
  if c.PostForm("dir") == "up" {
    food.Vote += 1
  } else if c.PostForm("dir") == "down" {
    food.Vote -= 1
  }
  db.Save(&food)
  c.JSON(http.StatusOK, gin.H{
    "status": "ok",
  })
}

func GetVotes(c *gin.Context){
  var (
    rest Restaurant
    food Food
  )
  FoodName := c.Query("food")
  RestaurantName := c.Query("restaurant")
  db.Where(&Restaurant{Username: RestaurantName}).First(&rest)
  db.Where(&Food{Restaurant: rest, Name: FoodName}).First(&food)
  c.JSON(http.StatusOK, gin.H{
    "votes": food.Vote,
  })
}
