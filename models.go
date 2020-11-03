package main

import (
  "gorm.io/gorm"
)

type Comment struct {
  gorm.Model
  Username string
  Msg string
  Food Food `gorm:"foreignKey:name"`
}

type User struct {
  gorm.Model
  Username string
  Password string
  Email string
  City string
  Token string
}

type Food struct {
  gorm.Model
  Restaurant Restaurant `gorm:"foreignKey:username"`
  Name string
  Desc string
  Price int
  Vote int
}

type Restaurant struct {
  gorm.Model
  Kind string
  Address string
  Name string
  City string
  Username string
  Password string
  Desc string
  Token string
}
