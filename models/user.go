package models

import (
	"example.com/jakkrit/ginbackendapi/utils"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Fullname string `json:"fullname" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`
	IsAdmin  bool   `json:"is_admin" gorm:"type:bool;default:false"` // is_active column:is_active (map column)
	ImageName string `json:"image_url"`
	Blogs 	 []Blog `json:"blogs"`
	Base
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.Password = utils.HashPassword(user.Password)
	return nil
}
