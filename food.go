package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// All foods of a Restaurant
func GetFoodByRestaurantName(c *gin.Context) {
	var MakerRestaurant Restaurant
	var FoodList []Food
	RestaurantUserName := c.Query("username")
	db.Where(&Restaurant{Username: RestaurantUserName}).First(&MakerRestaurant)
	db.Where(&Food{Restaurant: MakerRestaurant}).Find(&FoodList)
	c.JSON(http.StatusOK, FoodList)
}

// Get a Food Data in Json format
func GetFoodData(c *gin.Context) {
	var food Food
	var MakerRestaurant Restaurant
	RestaurantUserName := c.Query("username")
	db.Where(&Restaurant{Username: RestaurantUserName}).First(&MakerRestaurant)
	db.Where(&Food{Name: c.Query("name"), Restaurant: MakerRestaurant}).First(&food)
	c.JSON(http.StatusOK, food)
}

// Create a Food in DateBase
func CreateFood(c *gin.Context) {
	var rest Restaurant
	price, _ := strconv.Atoi(c.PostForm("desc"))
	RestaurantUserName := c.PostForm("username")
	db.Where(&Restaurant{Username: RestaurantUserName}).First(&rest)
	db.Create(&Food{
		Restaurant: rest,
		Desc:       c.PostForm("desc"),
		Name:       c.PostForm("name"),
		Price:      price,
		Vote:       0,
	})
}

// Get Restaurant by their votes
func FoodByVote(c *gin.Context) {
	restaurants := []Restaurant{}
	result := db.Where(&Restaurant{
		City: c.PostForm("city"),
	}).Order("total").Limit(20).Find(&restaurants)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"error": "رستورانی پیدا نشد",
		})
	}
	c.JSON(200, restaurants)
}
