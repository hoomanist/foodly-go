package database

import (
	"gorm.io/gorm"
  "github.com/hoomanist/foodly"
  "fmt"
)

func (model Restaurant) All(database *gorm.DB, city string) error {
  result := database.Where(&Restaurant{City: city}).Find(&model)
  if result.Error != nil {
    return  result.Error
  }
  return nil
}


func (model Restaurant) One(database *gorm.DB, username string) error {
  result := database.Where(&Restaurant{Username: username}).First(&model)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func (model Restaurant) Login(database *gorm.DB, data map[string]string)  error {
  rest := Restaurant{}
  result := database.Where(&Restaurant{
    Username: data["username"],
    Password: fmt.Sprintln(main.Hash(data["password"])),
  }).First(&rest)
  if result.Error != nil {
    return result.Error
  }
  return nil
}


func (model Restaurant) Create(database *gorm.DB, data map[string]string) error {
  result := database.Where(&Restaurant{
    Username: data["username"],
  })
  if result.Error == nil {
    return result.Error
  }
  result = database.Create(&Restaurant{
    Username: data["username"],
    Password: fmt.Sprintln(main.Hash(data["password"])),
    Token: main.GenerateToken(data["password"]),
    Address: data["address"],
    Kind: data["category"],
    Name: data["name"],
    Desc: data["desc"],
  }).First(&model)
  if result.Error != nil {
    return result.Error
  }
  return nil
}
