package database

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
  User     User
	Msg      string
	Food     Food `gorm:"foreignKey:name"`
}

type Food struct {
	gorm.Model
	Restaurant Restaurant `gorm:"foreignKey:username"`
	Name       string
	Desc       string
	Price      int
	Vote       int
}

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	City     string
	Token    string
}

type Restaurant struct {
	gorm.Model
	Kind     string
	Address  string
	Name     string
	City     string
	Username string
	Password string
	Desc     string
	Token    string
}

