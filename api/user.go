package api

import (
	moddleware "funny-serve/middleware"
	"funny-serve/model"
	"net/http"

	errormessages "funny-serve/utils/errorMessages"

	"github.com/gin-gonic/gin"
)

var code int

func AddUser(c *gin.Context) {
	var data model.User

	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errormessages.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errormessages.ERROR_USERNAME_USED {
		code = errormessages.ERROR_USERNAME_USED
	}
	i, uid := model.GetUserID(data.Username)
	if i != errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"message": errormessages.GetErrMsg(i),
		})
	}
	var userinfo = model.UserInfo{
		UserID:   uid,
		Username: data.Username,
		Phone:    data.Phone,
		Email:    data.Email,
		Avater:   "",
		NickName: "",
		Decrib:   "",
		Addr:     "",
		QQ:       "",
		Wechat:   ""}
	err := model.CreateUserInfo(&userinfo)
	if err != errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"message": "创建info表错误",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormessages.GetErrMsg(code),
	})
}

func Login(c *gin.Context) {
	var data model.User
	var token string
	c.ShouldBindJSON(&data)
	code = model.CheckLogin(data.Username, data.Password)
	if code != errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":              code,
			"message":             errormessages.GetErrMsg(code),
			"token":               "",
			"refreshToken":        "",
			"refreshToken-status": "",
			"userid":              "",
			"username":            "",
			"avator":              "",
			"role":                "",
		})
		c.Abort()
		return
	}
	token, code = moddleware.GenToken(data.Username, moddleware.Passtime, "user")
	refreshToken, i := moddleware.GenToken(data.Username, moddleware.Refreshtime, "admin")
	//c.SetCookie("username", data.Username, 1000, "/", "localhost", false, true)
	ok, userID := model.GetUserID(data.Username)
	if ok != errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":              ok,
			"message":             errormessages.GetErrMsg(ok),
			"token":               "",
			"refreshToken":        "",
			"refreshToken-status": "",
			"userid":              "",
			"username":            "",
			"avator":              "",
			"role":                "",
		})
	}
	i2, user := model.GetUser(userID)
	if i2 != errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":              i2,
			"message":             errormessages.GetErrMsg(i2),
			"token":               "",
			"refreshToken":        "",
			"refreshToken-status": "",
			"userid":              "",
			"username":            "",
			"avator":              "",
			"role":                "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":              "200",
			"message":             errormessages.GetErrMsg(code),
			"token":               token,
			"refreshToken":        refreshToken,
			"refreshToken-status": i,
			"userid":              user.UserID,
			"username":            user.Username,
			"avator":              user.Avater,
			"role":                user.Role,
		})
	}

}

// 查询用户信息
func GetUser(c *gin.Context) {
	var userid int
	c.ShouldBindJSON(&userid)
	code, userinfo := model.GetUser(userid)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    userinfo,
		"message": errormessages.GetErrMsg(code),
	})

}

// updata userinfo func
func UpDataUserInfo(c *gin.Context) {
	var userinfo model.UserInfo
	c.ShouldBindJSON(&userinfo)
	err := model.UpDataUserInfo(userinfo)
	if err != errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  err,
			"message": errormessages.GetErrMsg(err),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "用户信息更新成功",
	})
}

func GetUserInfo(c *gin.Context) {
	userid := c.Query("userid")
	userinfo, code := model.GetUserInfo(userid)
	if code != errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errormessages.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errormessages.GetErrMsg(code),
			"data":   userinfo,
		})
	}
}

// 修改资料
func UpDataUserInfoOneField(c *gin.Context) {
	userid := c.Query("userid")
	field := c.Query("field")
	newDate := c.Query("newDate")

	i := model.UpDataUserInfoOneField(userid, field, newDate)

	if i != errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"msg":    "更新失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "更新成功",
		})
	}

}
