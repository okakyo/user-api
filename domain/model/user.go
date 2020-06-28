package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
)

type mydqlDB struct {
	db *gorm.DB
}

type UserBase struct {
	gorm.Model
	name string `json:"name"`
	userId string `json:"userId"`
	email string `json:"email"`
	password string `json:"password"`
}

// ここでDB情報を拡張する必要がある
type UserDetail struct {
	gorm.Model
	user UserBase
	emailVerified string `json:"emailVerified"`
	thumbnail string "json:thumbnail"
	snsId SNSData `json:sns` // Has To の関係で実装をするuserId string `json:"userId"`
	fullName string `json:"name"`
	birth string `json:"birth"`
	description string `json:"description"`
}

func FindAllUserData() {
	var users []UserBase 
	if error := db.Find(&users).Error;error!=nil {
		panic(error)
	}
}

func FindUserByUserId(userId string){
	var userBase UserBase
	var userDetail UserDetail
	if error:= db.Where(&userBase).Find(&userDetail).Error;error!=nil{
		panic(error)
	}

}

