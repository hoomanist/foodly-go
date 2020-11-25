package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "fmt"
)

// create a user in database
func RegisterUser(c *gin.Context){
  var user User
  fmt.Println(c.PostForm("username"))
  result := db.Where(&User{Username: c.PostForm("username")}).Take(&user)
  fmt.Println(result)
  if result.Error == nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "repitidious username",
    })
    return
  }
  token := GenerateToken(c.PostForm("password"))
  result = db.Create(&User{
    Username: c.PostForm("username"),
    Password: fmt.Sprintln(Hash(c.PostForm("password"))),
    Email: c.PostForm("email"),
    City: c.PostForm("city"),
    Token: fmt.Sprintln(token),
  })
  fmt.Println(result.Error)
  c.JSON(http.StatusOK, gin.H{
    "token": fmt.Sprintln(token)[:len(fmt.Sprintln(token))-1],
  })
}

// login to a user and return a token
func LoginUser(c *gin.Context){
  var user []User
  db.Where(&User{Username: c.PostForm("username"), Password: fmt.Sprintln(Hash(c.PostForm("password")))}).Find(&user)
  if len(user) == 0 {
    c.JSON(http.StatusNotFound, gin.H{
      "error": "user not found",
    })
    return
  }
  c.JSON(http.StatusOK, gin.H{
    "token": user[0].Token[:len(user[0].Token)-1],
  })
}

// create a new restaurant in database
func CreateRestaurant(c *gin.Context){
  var restaurant Restaurant
  result := db.Where(&Restaurant{Username: c.PostForm("username")}).First(&restaurant)
  fmt.Println(restaurant)
  if result.Error == nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "repitidious username",
    })
    return
  }
  token := GenerateToken(c.PostForm("password"))
  db.Create(&Restaurant{
    Username: c.PostForm("username"),
    Name: c.PostForm("name"),
    Kind: c.PostForm("type"),
    Desc: c.PostForm("desc"),
    Address: c.PostForm("address"),
    Password: fmt.Sprintln(Hash(c.PostForm("password"))),
    City: c.PostForm("city"),
    Token: fmt.Sprintln(token),
  })
  c.JSON(http.StatusOK, gin.H{
    "token": fmt.Sprintln(token),
  })
}

// get the restaurant credentials and return it's token
func LoginRestaurant(c *gin.Context){
  var rest Restaurant
  result := db.Where(&Restaurant{Username: c.PostForm("username"), Password: fmt.Sprintln(Hash(c.PostForm("password")))}).Find(&rest)
  if result.Error != nil{
    c.JSON(http.StatusNotFound, gin.H{
      "error": "user not found",
    })
  }
  c.JSON(http.StatusOK, gin.H{
    "token": rest.Token[:len(rest.Token)-1],
  })
}
