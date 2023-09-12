package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qingstor/qingstor-sdk-go/v4/config"
	"github.com/qingstor/qingstor-sdk-go/v4/service"
	uuid "github.com/satori/go.uuid"
)

var qy_access_key_id = "TUZFWDOWMDJANHJFQNAG"
var qy_secret_access_key = "dHgi1n2NVML70i3X92s6JcLG7p8anVNoDEP6nm3y"

func UpData(ctx *gin.Context) {
	configuration, _ := config.New(qy_access_key_id, qy_secret_access_key)
	qsService, _ := service.Init(configuration)

	bucket, _ := qsService.Bucket("zzhh-server", "pek3b")

	f, fheader, err := ctx.Request.FormFile("imgfile")
	uid := uuid.NewV4()
	name := fmt.Sprintf("%s%s", uid.String(), fheader.Filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	result, err2 := bucket.PutObject(name, &service.PutObjectInput{Body: f})
	fmt.Println(service.IntValue(result.StatusCode))
	if err2 != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "上传成功",
		"url":    "https://zzhh-server.pek3b.qingstor.com" + "/" + name,
	})

}
