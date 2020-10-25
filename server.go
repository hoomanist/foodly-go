package main

import (
	"context"
	"fmt"
	"foodly/ent/account"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
// create a new account
func register(c *gin.Context) {
	ctx := context.Context(context.Background())
	// check weather username is repetidious or not
	_, err := cursor.Account.Query().Where(account.UsernameEQ(c.Query("username"))).First(ctx)
	if err == nil {
		c.JSON(400, gin.H{
			"error": "username was previously used. Please use diffrent username",
		})
		return
	}
	// generate a token and write the entry to db
	token := GenerateToken(c.Query("password"))
	_, err = cursor.Account.Create().
		SetUsername(c.Query("username")).
		SetPassword(Hash(c.Query("password"))).
		SetRole(c.Query("role")).
		SetEmail(c.Query("email")).
		SetCity(c.Query("city")).
		SetToken(token).
		Save(ctx)
	if err != nil {
		fmt.Printf("failed creating user: %v", err)
		return
	}
	// send back the token
	c.JSON(200, gin.H{
		"token": token,
	})
}
// login to account
func login(c *gin.Context) {
	ctx := context.Context(context.Background())
	u, err := cursor.Account.Query().Where(account.UsernameEqualFold(c.Query("username")),
		account.PasswordEqualFold(Hash(c.Query("password")))).First(ctx)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "no such user in db. maybe your entering a wrong password.",
		})
		return
	}
	c.JSON(200, gin.H{
		"token": u.Token,
	})
}
// submit food
func SubmitFood(c *gin.Context) {
	ctx := context.Context(context.Background())
	u, err := cursor.Account.Query().Where(account.TokenEqualFold(c.Query("token"))).First(ctx)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "no such token in db. maybe your entering a wrong token.",
		})
		return
	}
	if u.Username == c.Query("username") {
		c.JSON(400, gin.H{
			"error": "token is not yours.",
		})
		return
	}
	if u.Role != "restaurant" {
		c.JSON(400, gin.H{
			"error": "you are not a restaurant.",
		})
		return
	}
	_, err = cursor.Food.Create().
		SetRestaurant(u.Username).
		SetPrice(c.Query("price")).
		SetName(c.Query("name")).
		SetImageName(c.Query("image_filename")).
		SetDesc(c.Query("desc")).
		Save(ctx)
	if err != nil {
		fmt.Printf("failed creating food: %v", err)
		return
	}
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
// upload images in ./uploads folder
func UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{
			"err": fmt.Errorf("%s", err),
		})
	}
	err = c.SaveUploadedFile(file, "uploads/"+file.Filename)

	if err != nil {
		c.JSON(400, gin.H{
			"err": fmt.Errorf("%s", err),
		})
	}
	c.JSON(200, gin.H{
		"filename": file.Filename,
	})
}
// server an image based on their name
func ServeImage(c *gin.Context){
	name := c.Param("name")
	c.File("uploads/" + name)
}
// query restaurant based on their city
func QueryRestaurants(c *gin.Context){
	ctx := context.Context(context.Background())
	accounts, err := cursor.Account.Query().Where(account.CityEQ(c.Query("city")), account.RoleEQ("restaurant")).All(ctx)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "no restaurant in your city is registered. it is time for leaving this town.",
		})
	}
	for i := 0 ; i <= len(accounts); i++ {
		accounts[i].Token = ""
		accounts[i].Password = ""
	}
	fmt.Println(accounts)
	c.JSON(200, accounts)
}