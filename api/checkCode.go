package api

import (
	"funny-serve/model"
	errormessages "funny-serve/utils/errorMessages"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckCode(c *gin.Context) {
	var registerInfo model.CheckCode

	_ = c.ShouldBindJSON(&registerInfo)
	code := model.ChekeRegisterInfo(registerInfo)
	if code != errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errormessages.GetErrMsg(code),
		})
	} else {
		err, codeString := model.SendCode(registerInfo.Email)
		if err != errormessages.SEND_CODE_SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"status":     err,
				"msg":        errormessages.GetErrMsg(err),
				"codestring": codeString,
			})
		} else {
			registerInfo.Code = codeString
			saveCode, saveCodeError := model.SaveCode(&registerInfo)
			c.JSON(http.StatusOK, gin.H{
				"status":     saveCodeError,
				"msg":        errormessages.GetErrMsg(saveCodeError),
				"codestring": saveCode,
			})
		}

	}

}
