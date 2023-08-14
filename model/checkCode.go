package model

import (
	"bytes"
	"crypto/rand"
	"fmt"
	errormessages "funny-serve/utils/errorMessages"
	"math/big"
	"strconv"

	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

type CheckCode struct {
	gorm.Model
	Username     string `gorm:"type:varchar(20);not null" json:"username"`
	Password     string `gorm:"type:varchar(20);not null" json:"password"`
	ChekPassword string `gorm:"type:varchar(20);not null" json:"chekpassword"`
	Phone        string `gorm:"type:varchar(20);not null" json:"phone"`
	Email        string `gorm:"type:varchar(30)" json:"email"`
	Role         string `gorm:"type:varchar(10);not null" json:"role"`
	Code         string `gorm:"type:varchar(10);not null" json:"code"`
}

//check register info

func ChekeRegisterInfo(info CheckCode) int {
	fmt.Print(info)
	if info.Username == "" {
		return errormessages.USERNAME_NULL
	} else if info.Password == "" {
		return errormessages.PASSWORD_NULL
	} else if info.Password != info.ChekPassword {
		return errormessages.CHECKPASSWORD_ERROR
	} else if info.Phone == "" {
		return errormessages.PHONE_NULL
	} else if info.Email == "" {
		return errormessages.EMAIL_NULL
	}
	return errormessages.SUCCESS
}
func SendCode(email string) (int, string) {
	g := gomail.NewMessage()
	g.SetHeader("From", "zenghaifenghandsom@163.com")
	g.SetHeader("To", email)
	g.SetHeader("Subject", "ceshi")
	codeString := RadomString(6)
	g.SetBody("text/html", "<h1>zzhh 验证码："+codeString+"</h1>")

	d := gomail.NewDialer("smtp.163.com", 25, "zenghaifenghandsom@163.com", "QRFSVWNRFQNLMJBF")

	if err := d.DialAndSend(g); err != nil {
		fmt.Printf("DialAndSend err %v", err)
		panic(err)
	}
	//fmt.Printf("send mail success\n")
	return errormessages.SEND_CODE_SUCCESS, codeString
}

func RadomString(lengh int) string {
	var VerificationCode string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	len := b.Len()
	bigInt := big.NewInt(int64(len))
	for i := 0; i < lengh; i++ {
		radomInt, _ := rand.Int(rand.Reader, bigInt)
		VerificationCode += string(str[radomInt.Int64()])
	}

	return VerificationCode
}

func SaveCode(code *CheckCode) (string, int) {
	res := db.Create(&code)
	if res.Error != nil {
		return errormessages.GetErrMsg(errormessages.ERROR_SAVECODE), errormessages.ERROR_SAVECODE
	} else {
		db.Where("code=? AND username=? AND password=? AND phone=? AND email=?", code.Code, code.Username, code.Password, code.Phone, code.Email).First(&code)
		return string(strconv.Itoa(int(code.ID))), errormessages.SUCCESS_SAVECODE
	}
}
