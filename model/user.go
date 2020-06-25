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
	snsId interface.SNSId `json:snsId`
}

// ここでDB情報を拡張する必要がある
type UserDetail struct {
	gorm.Model
	userId string `json:"userId"`
	fullName string `json:"name"`
	birth string `json:"birth"`
	description string `json:"description"`
}
