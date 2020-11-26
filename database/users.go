package database

import (
	"gorm.io/gorm"
  "github.com/hoomanist/foodly/tools"
  "fmt"
)

func (model User) Create(database *gorm.DB, data map[string]string) error {
  result := database.Where(&User{
    Username: data["username"],
  })
  if result.Error != nil {
    return result.Error
  }
  result = database.Create(&User{
    Username: data["username"],
    Password: fmt.Sprintln(tools.Hash(data["password"])),
    City: data["username"],
    Email: data["email"],
    Token: tools.GenerateToken(data["password"]),
  }).First(&model)
  return nil
}

func (model User) Login(database *gorm.DB, data map[string]string) error {
  result := database.Where(&User{
    Username: data["username"],
    Password: fmt.Sprintln(tools.Hash(data["password"])),
  }).Take(&model)
  if result.Error != nil {
    return result.Error
  }
  return nil
}
