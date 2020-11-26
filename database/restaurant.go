package database

import (
	"gorm.io/gorm"
  "github.com/hoomanist/foodly/tools"
  "fmt"
)

func (model Restaurant) All(database *gorm.DB, city string) ([]Restaurant, error) {
  rsp := []Restaurant{}
  result := database.Where(&Restaurant{City: city}).Find(&rsp)
  if result.Error != nil {
    return  []Restaurant{}, result.Error
  }
  return rsp, nil
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
    Password: fmt.Sprintln(tools.Hash(data["password"])),
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
  model = Restaurant{
    Username: data["username"],
    Password: fmt.Sprintln(tools.Hash(data["password"])),
    Token: tools.GenerateToken(data["password"]),
    Address: data["address"],
    Kind: data["category"],
    Name: data["name"],
    Desc: data["desc"],
  }
  result = database.Create(&model)
  if result.Error != nil {
    return result.Error
  }
  return nil
}
