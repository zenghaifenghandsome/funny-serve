package api

import (
	"net/http"
	"funny-serve/model"
	errormessages "funny-serve/utils/errorMessages"

	"github.com/gin-gonic/gin"
)

func AddEvd(ctx *gin.Context) {
	var newEvd model.Evd
	_ = ctx.ShouldBindJSON(&newEvd)

	code := model.AddEvd(&newEvd)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormessages.GetErrMsg(code),
	})
}

func GetAllEvd(ctx *gin.Context) {
	result, code := model.GetAllEvd()
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormessages.GetErrMsg(code),
		"data":   result,
	})
}
