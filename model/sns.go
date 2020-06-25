package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/okakyo/user-api/interface"
)

type SNSId struct {
	gorm.Model
	snsType string `json:"type"`
	accountId string `json:"accountId"`
}