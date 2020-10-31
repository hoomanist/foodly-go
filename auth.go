package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "crypto/md5"
	"crypto/sha1"
	"encoding/hex"
  "golang.org/x/crypto/bcrypt"
)

func GenerateToken(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}

func Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return string(bs)
}

func RegisterUser(c *gin.Context){
  var user User
  db.Where(&User{username: c.Query("username")}).Find(&user)
  if user.username != "" {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "repitidious username",
    })
  }
  token := GenerateToken(c.Query("password"))
  db.Create(&User{
    username: c.Query("username"),
    password: Hash(c.Query("password")),
    email: c.Query("email"),
    city: c.Query("city"),
    token: token,
  })
  c.JSON(http.StatusOK, gin.H{
    "token": token,
  })
}

func LoginUser(c *gin.Context){
  var user User
  db.Where(&User{username: c.Query("username"), password: Hash(c.Query("password"))}).Find(&user)
  if user.username == "" {
    c.JSON(http.StatusNotFound, gin.H{
      "error": "user not found",
    })
  }
  c.JSON(http.StatusOK, gin.H{
    "token": user.token,
  })
}

func CreateRestaurant(c *gin.Context){
  var restaurant Restaurant
  db.Where(&Restaurant{username: c.Query("username")}).Find(&restaurant)
  if restaurant.username != "" {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "repitidious username",
    })
  }
  token := GenerateToken(c.Query("password"))
  db.Create(&Restaurant{
    username: c.Query("username"),
    Name: c.Query("name"),
    kind: c.Query("type"),
    desc: c.Query("desc"),
    address: c.Query("address"),
    password: Hash(c.Query("password")),
    city: c.Query("city"),
    token: token,
  })
  c.JSON(http.StatusOK, gin.H{
    "token": token,
  })
}
func LoginRestaurant(c *gin.Context){
  var rest Restaurant
  db.Where(&Restaurant{username: c.Query("username"), password: Hash(c.Query("password"))}).Find(&rest)
  if rest.username == "" {
    c.JSON(http.StatusNotFound, gin.H{
      "error": "user not found",
    })
  }
  c.JSON(http.StatusOK, gin.H{
    "token": rest.token,
  })
}
