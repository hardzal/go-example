package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"pasword"`
}

func (u *User) CreateUser() *User {
	Db.NewRecord(u)
	Db.Create(&u)
	return u
}

func GetUsers() []User {
	var Users []User
	Db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := Db.Where("ID = ?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(Id int64) User {
	var user User
	Db.Where("ID = ?", Id).Delete(user)
	return user
}
