package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SNSData struct {
	gorm.Model
	snsType string `json:"type"`
	accountId string `json:"accountId"`
}