package model

import (
	errormessages "funny-serve/utils/errorMessages"

	"gorm.io/gorm"
)

type Share struct {
	gorm.Model
	UserId  uint   `gorm:"type:int;not null" json:"userid"`
	Content string `gorm:"type:text;not null" json:"content"`
	PicList string `gorm:"type:text;default null" json:"piclist"`
	Video   string `gorm:"type:text;default null" json:"video"`
}

func AddHare(share *Share)int {
	res := db.Create(&share)
	err := res.Error
	if err != nil {
		return errormessages.ERROR
	}
	return errormessages.SUCCESS
}

func GetAllShare() ([]Share,int){
	var share []Share
	result := db.Find(&share)
	if result.Error != nil {
		return nil, errormessages.ERROR
	}
	return share, errormessages.SUCCESS
}

