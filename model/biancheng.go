package model

import (
	errormessages "funny-serve/utils/errorMessages"

	"gorm.io/gorm"
)

type BianCheng struct {
	gorm.Model
	IconAddr    string `gorm:"type:varchar(200)" json:"iconaddr"`
	Title       string `gorm:"type:varchar(20)" json:"title"`
	Link        string `gorm:"type:varchar(200)" json:"link"`
	Describ     string `gorm:"type:varchar(50)" json:"describ"`
	Detail      string `gorm:"type:varchar(200)" json:"detail"`
	Tags        string `gorm:"type:varchar(200)" json:"tags"`
	Userid      string `gorm:"type:varchar(50)" json:"userid"`
	Look        int    `gorm:"type:int ;default: 0" json:"look"`
	Like        int    `gorm:"type:int ;default: 0" json:"like"`
	Username    string `gorm:"type:varchar(50)" json:"username"`
	Avater      string `gorm:"type:varchar(200);default:''" json:"avater"`
	Commentnumb int    `gorm:"type:int ;default: 0" json:"commentnumb"`
	Status      int    `gorm:"type:int ;default: 0" json:"status"`
}

// type Comment struct {
// 	gorm.Model
// 	Speak       string    `gorm:"type:varchar(200)" json:"speak"`
// 	CommentUser string    `gorm:"type:varchar(50)" json:"commentuser"`
// 	Comments    []Comment `gorm:"type:varchar(200)" json:"comments"`
// 	Like        int       `gorm:"type:int default 0" json:"like"`
// }

//查询所有编程导航列表

func GetAllBiancheng() ([]BianCheng, int) {
	var biancheng []BianCheng
	result := db.Find(&biancheng)
	if result.Error != nil {
		return nil, errormessages.ERROR
	}
	return biancheng, errormessages.SUCCESS
}
func GetOkAllBiancheng() ([]BianCheng, int) {
	var biancheng []BianCheng
	result := db.Where("status = ?", 1).Find(&biancheng)
	if result.Error != nil {
		return nil, errormessages.ERROR
	}
	return biancheng, errormessages.SUCCESS
}

// 查询Bcdaohang是否存在
func CheckBC(title string) int {
	var bcs BianCheng
	db.Select("id").Where("title = ?", title).First(&bcs)
	if bcs.ID > 0 {
		return errormessages.ERROR_USERNAME_USED
	}
	return errormessages.SUCCESS
}

//新增推荐编程导航

func AddBiancheng(data *BianCheng) int {
	res := db.Create(&data)
	err := res.Error
	if err != nil {
		return errormessages.ERROR
	}
	return errormessages.SUCCESS
}

// 查询编程导航信息
func GetBianchengInfo(id string) (BianCheng, int) {
	var bc BianCheng
	result := db.Where("id = ?", id).First(&bc)
	if result.Error != nil {
		return bc, errormessages.ERROR
	}
	return bc, errormessages.SUCCESS
}

func DeletBc(id string) int {
	var BC BianCheng
	result := db.Where("ID = ?", id).Delete(&BC)
	if result.Error != nil {
		return errormessages.ERROR_DELETBIANCHENG_ERROR
	} else {
		return errormessages.SUCCESS_DELETBIANCHENG
	}

}

func SetOk(id string) int {
	var bc BianCheng
	db.Model(&bc).Where("id=?", id).Update("status", 1)
	return errormessages.SUCCESS
}
func SetNo(id string) int {
	var bc BianCheng
	db.Model(&bc).Where("id=?", id).Update("status", 0)
	return errormessages.SUCCESS
}
