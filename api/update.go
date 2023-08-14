package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qingstor/qingstor-sdk-go/v4/config"
	"github.com/qingstor/qingstor-sdk-go/v4/service"
)

func UpDate(ctx *gin.Context) {
	fmt.Println("update...")
	//qy_access_key_id := "TUZFWDOWMDJANHJFQNAG"
	//qy_secret_access_key := "dHgi1n2NVML70i3X92s6JcLG7p8anVNoDEP6nm3y"
	configuration, _ := config.New("TUZFWDOWMDJANHJFQNAG", "dHgi1n2NVML70i3X92s6JcLG7p8anVNoDEP6nm3y")
	qsService, _ := service.Init(configuration)
	baseUrl := "https://zzhh-server.pek3b.qingstor.com/"
	bucketService, _ := qsService.Bucket("zzhh-server", "pek3b")

	//name := "test/objectput.jpg"

	//f, err := os.Open("C://Users//zengh//Desktop//5c3c751112b7a119a4d696419677cbf.jpg")
	f, fheader, err := ctx.Request.FormFile("imgfile")
	name := fmt.Sprintf("%d%s", time.Now().Unix(), fheader.Filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//	_, err2 := c.Object.Put(context.Background(), name, f, nil)
	output, err2 := bucketService.PutObject(name, &service.PutObjectInput{Body: f})
	if err2 != nil {
		panic(err)
	} else {
		fmt.Println(output.StatusCode)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "上传成功",
		"url":    baseUrl + name,
	})
}
