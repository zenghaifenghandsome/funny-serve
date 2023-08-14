package model

import (
	"encoding/base64"
	"fmt"
	"log"

	errormessages "funny-serve/utils/errorMessages"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

// 用户基本信息表-注册表
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Phone    string `gorm:"type:varchar(20);not null" json:"phone"`
	Email    string `gorm:"type:varchar(30)" json:"email"`
	Role     string `gorm:"type:varchar(10);not null" json:"role"`
}

// 用户详细信息表
type UserInfo struct {
	UserID   int    `gorm:"type:int;unique;not null;primaryKey" json:"userid"`
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Phone    string `gorm:"type:varchar(20);not null" json:"phone"`
	Email    string `gorm:"type:varchar(30)" json:"email"`
	Avater   string `gorm:"type:varchar(200);default:''" json:"avater"`
	NickName string `gorm:"type:varchar(20)" json:"nickname"`
	Decrib   string `gorm:"type:varchar(200)" json:"decrib"`
	Addr     string `gorm:"type:varchar(50)" json:"addr"`
	QQ       string `gorm:"type:varchar(12)" json:"QQ"`
	Wechat   string `gorm:"type:varchar(40)" json:"wechat"`
	Role     string `gorm:"type:varchar(10);not null" json:"role"`
}

// 用户登录状态表
type UserState struct {
	gorm.Model
	UserId     int    `gorm:"type:int" jason:"userid"`
	LoginState int    `gorm:"type:int" json:"loginstate"`
	Token      string `gorm:"type:varchar(50)" json:"token"`
}

// 查询用户是否存在
func CheckUser(name string) int {
	var user User
	fmt.Println("serch User start............")
	db.Select("id").Where("username = ?", name).First(&user)
	fmt.Println("serch user over..............")
	if user.ID > 0 {
		return errormessages.ERROR_USERNAME_USED
	}
	return errormessages.SUCCESS
}

// 注册用户
func CreateUser(data *User) int {
	data.Password = ScriptPW(data.Password)
	res := db.Create(&data)
	err := res.Error
	if err != nil {
		return errormessages.ERROR
	}
	return errormessages.SUCCESS
}

func CreateUserInfo(data *UserInfo) int {
	res := db.Create(&data)
	err := res.Error
	if err != nil {
		return errormessages.ERROR
	}
	return errormessages.SUCCESS
}

// 密码加密
func ScriptPW(password string) string {
	const KeyLen = 10
	var salt = make([]byte, 8)
	salt = []byte{12, 10, 2, 11, 22, 33, 44, 88}
	HashPW, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPW)
	return fpw
}

// 判断是否登录
func IsLogin(token string) int {
	if len(token) > 0 {
		return errormessages.SUCCESS
	}
	return errormessages.ERROR
}

// 登录
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		return errormessages.ERROR_USER_NOT_EXIST
	}
	if ScriptPW(password) != user.Password {
		return errormessages.ERROR_PASSWORD_WRONG
	} else {
		AddUserToUserLoginState(user)
		return errormessages.SUCCESS
	}

}

// 添加用户信息到用户登录状态表
func AddUserToUserLoginState(user User) int {
	var userstate UserState
	userstate.UserId = int(user.ID)
	userstate.LoginState = 1
	userstate.Token = ""

	//判断是否已登录
	db.Select("id").Where("userid = ?", userstate.UserId).First(&userstate)
	if userstate.ID > 0 {
		//已登录，直接返回
		return errormessages.ERROR
	}
	//未登录 添加登录状态
	res := db.Create(&userstate)
	err := res.Error
	if err != nil {
		return errormessages.ERROR
	}
	return errormessages.SUCCESS

}

// 查询用户信息
func GetUser(userid int) (int, UserInfo) {
	var user UserInfo
	db.Where("user_id = ?", userid).First(&user)
	return errormessages.SUCCESS, user
}

func GetUserID(username string) (int, int) {
	var user User
	d := db.Where("username = ?", username).First(&user)
	fmt.Println(int(user.ID))
	if d.Error != nil {
		return errormessages.ERROR_USER_ID_ERROR, 0
	}
	return errormessages.SUCCESS, int(user.ID)
}

// updata userInfo func
func UpDataUserInfo(userinfo UserInfo) int {

	result := db.Model(&userinfo).Updates(map[string]interface{}{
		"username":  userinfo.Username,
		"phone":     userinfo.Phone,
		"email":     userinfo.Email,
		"avater":    userinfo.Avater,
		"nick_name": userinfo.NickName,
		"decrib":    userinfo.Decrib,
		"addr":      userinfo.Addr,
		"qq":        userinfo.QQ,
		"wechat":    userinfo.Wechat,
	})
	if result.Error != nil {
		return errormessages.ERROR_UPDATA_USERINFO_FAIL
	}
	return errormessages.SUCCESS
}

// 获取userInfo
func GetUserInfo(userid string) (UserInfo, int) {
	var userinfo UserInfo
	result := db.First(&userinfo, userid)
	if result.Error != nil {
		return userinfo, 400
	}
	return userinfo, errormessages.SUCCESS

}

func UpDataUserInfoOneField(userid string, field string, newDate string) int {
	var userinfo UserInfo
	result := db.Model(&userinfo).Where("user_id = ?", userid).Update(field, newDate)
	if result.Error != nil {
		return 400
	}
	return errormessages.SUCCESS
}
