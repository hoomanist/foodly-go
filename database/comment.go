package database

import (
  "gorm.io/gorm"
)

func (model Comment) All(database *gorm.DB, food Food) ([]Comment, error) {
  rsp := []Comment{}
  result := database.Where(&Comment{
    Food:food,
  }).Find(&rsp)
  if result.Error != nil {return []Comment{}, result.Error}
  return rsp, nil
}

func (model Comment) Create(database *gorm.DB, food Food, msg string, Username User) error {
  model = Comment{
    User: Username,
    Food: food,
    Msg: msg,
  }
  result := database.Create(&model)
  if result.Error != nil {return result.Error}
  return nil
}
