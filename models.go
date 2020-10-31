package main

import (
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  username string
  password string
  email string
  city string
  token string
}

type Food struct {
  gorm.Model
  restaurant Restaurant `gorm:"foreignKey:username"`
  name string
  desc string
  price int
  vote int
}

type Restaurant struct {
  gorm.Model
  kind string
  address string
  Name string
  city string
  username string
  password string
  desc string
  token string
}
