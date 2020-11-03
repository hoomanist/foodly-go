package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "fmt"
  "crypto/md5"
	"crypto/sha1"
	"encoding/hex"
  "golang.org/x/crypto/bcrypt"
)

// generate a token based on password with a random salt
func GenerateToken(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}

// create a hash
func Hash(s string) []byte {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return bs
}

// create a user in database
func RegisterUser(c *gin.Context){
  var user User
  fmt.Println(c.Query("username"))
  result := db.Where(&User{Username: c.Query("username")}).Take(&user)
  fmt.Println(result)
  if result.Error == nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "repitidious username",
    })
    return
  }
  token := GenerateToken(c.Query("password"))
  result = db.Create(&User{
    Username: c.Query("username"),
    Password: fmt.Sprintln(Hash(c.Query("password"))),
    Email: c.Query("email"),
    City: c.Query("city"),
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
  db.Where(&User{Username: c.Query("username"), Password: fmt.Sprintln(Hash(c.Query("password")))}).Find(&user)
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
  result := db.Where(&Restaurant{Username: c.Query("username")}).First(&restaurant)
  fmt.Println(restaurant)
  if result.Error == nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "repitidious username",
    })
    return
  }
  token := GenerateToken(c.Query("password"))
  db.Create(&Restaurant{
    Username: c.Query("username"),
    Name: c.Query("name"),
    Kind: c.Query("type"),
    Desc: c.Query("desc"),
    Address: c.Query("address"),
    Password: fmt.Sprintln(Hash(c.Query("password"))),
    City: c.Query("city"),
    Token: fmt.Sprintln(token),
  })
  c.JSON(http.StatusOK, gin.H{
    "token": fmt.Sprintln(token),
  })
}

// get the restaurant credentials and return it's token
func LoginRestaurant(c *gin.Context){
  var rest []Restaurant
  db.Where(&Restaurant{Username: c.Query("username"), Password: fmt.Sprintln(Hash(c.Query("password")))}).Find(&rest)
  if len(rest) == 0{
    c.JSON(http.StatusNotFound, gin.H{
      "error": "user not found",
    })
  }
  c.JSON(http.StatusOK, gin.H{
    "token": rest[0].Token,
  })
}
