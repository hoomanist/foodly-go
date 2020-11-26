package database

import (
  "github.com/go-playground/assert/v2"
  "fmt"
  "gorm.io/gorm"
  "testing"
  "gorm.io/driver/mysql"
)

func TestAllRestaurants(t *testing.T) {
  db, err := gorm.Open(mysql.Open("hooman:hooman86@tcp(127.0.0.1:3306)/foodly?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
  assert.Equal(t, err, nil)
  var restaurant Restaurant
  var (
    rsp []Restaurant
    expected []Restaurant

  )
  rsp, err = restaurant.All(db, "qom")
  assert.Equal(t, err, nil)
  db.Where(&Restaurant{City: "qom"}).Find(&expected)
  fmt.Println(rsp)
  fmt.Println(expected)
  assert.Equal(t, expected, rsp)
}
