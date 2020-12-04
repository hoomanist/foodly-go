package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hoomanist/foodly/database"
	"net/http"
)

// create a user in database
func RegisterUser(c *gin.Context) {
	var user database.User
	err := user.Create(db, map[string]string{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
		"city":     c.PostForm("city"),
		"email":    c.PostForm("email"),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "یه مشکلی پیش اومده",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"token": user.Token,
	})
}

// login to a user and return a token
func LoginUser(c *gin.Context) {
	var user database.User
	err := user.Login(db, map[string]string{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "یه مشکلی پیش اومده",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"token": user.Token,
	})
}

// create a new restaurant in database
func CreateRestaurant(c *gin.Context) {
	var restaurant database.Restaurant
	err := restaurant.Create(db, map[string]string{
		"Username": c.PostForm("username"),
		"Name":     c.PostForm("name"),
		"Kind":     c.PostForm("type"),
		"Desc":     c.PostForm("desc"),
		"Address":  c.PostForm("address"),
		"Password": c.PostForm("password"),
		"City":     c.PostForm("city"),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "یه مشکلی پیش اومده",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"token": restaurant.Token,
	})
}

// get the restaurant credentials and return it's token
func LoginRestaurant(c *gin.Context) {
	var rest database.Restaurant
	err := rest.Login(db, map[string]string{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "یه مشکلی پیش اومده",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"token": rest.Token,
	})
}
