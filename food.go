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
  db.Where(&Restaurant{Username: RestaurantUserName}).First(&MakerRestaurant)
  db.Where(&Food{Restaurant: MakerRestaurant}).Find(&FoodList)
  c.JSON(http.StatusOK, FoodList)
}

// Get a Food Data in Json format
func GetFoodData(c *gin.Context){
  var food Food
  var MakerRestaurant Restaurant
  RestaurantUserName := c.Query("username")
  db.Where(&Restaurant{Username: RestaurantUserName}).First(&MakerRestaurant)
  db.Where(&Food{Name: c.Query("name"), Restaurant: MakerRestaurant}).First(&food)
  c.JSON(http.StatusOK, food)
}

// Create a Food in DateBase
func CreateFood(c *gin.Context){
  var rest Restaurant
  price, _ := strconv.Atoi(c.Query("desc"))
  RestaurantUserName := c.Query("username")
  db.Where(&Restaurant{Username: RestaurantUserName}).First(&rest)
  db.Create(&Food{
    Restaurant: rest,
    Desc: c.Query("desc"),
    Name: c.Query("name"),
    Price: price,
    Vote: 0,
  })
}
