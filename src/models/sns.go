package model

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type SNSId struct {
	gorm.Model
	snsType string `json:"type"`
	userId string `json:"userId"`
}