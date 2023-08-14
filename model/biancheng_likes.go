package model

import (
	errormessages "funny-serve/utils/errorMessages"
	"strconv"

	"gorm.io/gorm"
)

type BianchengLikes struct {
	gorm.Model
	Like           int    `gorm:"type:int" json:"like"`
	UserId         int    `gorm:"type:int;not null" json:"userid"`
	Username       string `gorm:"type:varchar(200);not null" json:"username"`
	BianchengId    int    `gorm:"type:int;not null" json:"bianchengid"`
	BianchengTitle string `gorm:"varchar(200);not null" json:"bianchengtitle"`
}

// 创建
func CreateLikes(bianchenglikes *BianchengLikes) int {
	result := db.Create(&bianchenglikes)
	if result.Error != nil {
		return errormessages.ERROR
	}
	return errormessages.SUCCESS
}

// 点赞
func AddLike(bcId int, usId int) int {
	bc, err := GetLike(bcId, usId)
	bcinfo, _ := GetBianchengInfo(strconv.Itoa(bcId))
	if err != errormessages.SUCCESS {

		return errormessages.ERROR_BIANCHENG_LIKE_SEARCH_VOID
	}
	if bc.Like == 0 {
		db.Model(&BianchengLikes{}).Where("id=? ", bc.ID).Update("like", 1)
		db.Model(&BianCheng{}).Where("id=?", bcinfo.ID).Update("like", bcinfo.Like+1)
		return errormessages.SUCCESS
	}
	db.Model(&BianchengLikes{}).Where("id=? ", bc.ID).Update("like", 0)
	db.Model(&BianCheng{}).Where("id=?", bcinfo.ID).Update("like", bcinfo.Like-1)
	return errormessages.SUCCESS
}

// 获取点赞信息
func GetLike(bcId int, usId int) (BianchengLikes, int) {
	var like BianchengLikes
	result := db.Where("biancheng_id = ? AND user_id=?", bcId, usId).First(&like)
	if result.Error != nil {
		return like, errormessages.ERROR
	}
	return like, errormessages.SUCCESS
}
