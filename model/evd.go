package model

import (
	errormessages "funny-serve/utils/errorMessages"

	"gorm.io/gorm"
)

type Evd struct {
	gorm.Model
	UserId  string `gorm:"type:varchar(200);not null" json:"userid"`
	ImgList string `gorm:"type:varchar(5000);default ''" json:"imgList"`
	Content string `gorm:"type:varchar(5000);default ''" json:"content"`
}

func AddEvd(evd *Evd) int {
	result := db.Create(&evd)
	if result.Error != nil {
		return errormessages.ERROR
	}
	return errormessages.SUCCESS
}

func GetAllEvd() ([]Evd, int) {
	var evds []Evd
	result := db.Find(&evds)
	if result.Error != nil {
		return nil, errormessages.ERROR
	}
	return evds, errormessages.SUCCESS
}
