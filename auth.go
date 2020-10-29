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
}

func LoginUser(c *gin.Context){
}
