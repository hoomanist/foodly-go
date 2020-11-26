package database

import (
  "gorm.io/gorm"
  "strconv"
)


func (model Food) All(database *gorm.DB, restaurant Restaurant) ([]Food,error) {
  rsp := []Food{}
  result := database.Where(&Food{
    Restaurant: restaurant,
  }).Find(&rsp)
  if result.Error != nil{
    return []Food{}, result.Error
  }
  return rsp, nil
}

func (model Food) One(database *gorm.DB, ID uint) error {
  result := database.Take(&model, ID)
  if result.Error != nil{return result.Error}
  return nil
}

func (model Food) Create(database *gorm.DB, restaurant Restaurant, data map[string]string) error {
  price, _:=strconv.Atoi(data["price"])
  model = Food{
    Restaurant: restaurant,
    Name: data["name"],
    Desc: data["desc"],
    Price: price,
    Vote: 0,
  }
  result := database.Create(model)
  if result.Error != nil {return result.Error}
  return nil
}
func (model Food) increaseVote(database *gorm.DB){
  model.Vote += 1
  database.Save(&model)
}
func (model Food) decreaseVote(database *gorm.DB){
  model.Vote -= 1
  database.Save(&model)
}
