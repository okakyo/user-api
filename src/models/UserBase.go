package model

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserBase struct {
	gorm.Model
	userId string `json:"userId"`
	name string `json:"name"`
	email string `json:"email"`
	emailVerified string `json:"emailVerified"`
	snsId string `json:snsId`
}

