package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title   string
	Url     string
	Img     string
	User_id string
}
