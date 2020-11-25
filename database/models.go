package database

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Username string
	Msg      string
	Food     Food `gorm:"foreignKey:name"`
}

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	City     string
	Token    string
}

type users interface {
  Create(database *gorm.DB, data map[string]string) error
  Login(database *gorm.DB, data map[string]string) error
}

type Food struct {
	gorm.Model
	Restaurant Restaurant `gorm:"foreignKey:username"`
	Name       string
	Desc       string
	Price      int
	Vote       int
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
type restaurant interface {
  All(database *gorm.DB, city string) error
  One(database *gorm.DB, username string) error
  Create(database *gorm.DB, data map[string]string) error
  Login(database *gorm.DB, data map[string]string) error
}

