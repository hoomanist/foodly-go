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
