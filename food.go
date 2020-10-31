package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "strconv"
)

// All foods of a Restaurant
func GetFoodByRestaurantName(c *gin.Context){
  var MakerRestaurant Restaurant
  var FoodList []Food
  RestaurantUserName := c.Query("username")
  db.Where(&Restaurant{username: RestaurantUserName}).First(&MakerRestaurant)
  db.Where(&Food{restaurant: MakerRestaurant}).Find(&FoodList)
  c.JSON(http.StatusOK, FoodList)
}

// Get a Food Data in Json format
func GetFoodData(c *gin.Context){
  var food Food
  var MakerRestaurant Restaurant
  RestaurantUserName := c.Query("username")
  db.Where(&Restaurant{username: RestaurantUserName}).First(&MakerRestaurant)
  db.Where(&Food{name: c.Query("name"), restaurant: MakerRestaurant}).First(&food)
  c.JSON(http.StatusOK, food)
}

// Create a Food in DateBase
func CreateFood(c *gin.Context){
  var rest Restaurant
  price, _ := strconv.Atoi(c.Query("desc"))
  RestaurantUserName := c.Query("username")
  db.Where(&Restaurant{username: RestaurantUserName}).First(&rest)
  db.Create(&Food{
    restaurant: rest,
    desc: c.Query("desc"),
    name: c.Query("name"),
    price: price,
    vote: 0,
  })
}
