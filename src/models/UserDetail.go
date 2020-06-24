package model

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserDetail struct {
	gorm.Model
	userId string `json:"userId"`
	fullName string `json:"name"`
	birth string `json:"email"`
	description string `json:"emailVerified"`
}

