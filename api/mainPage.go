package api

import (
	"encoding/json"
	"fmt"
	"funny-serve/model"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetMainPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "首页请求成功",
		"data": getMain_joke(),
	})
}

func getMain_joke() model.JokeInfo {
	jokeInfo := &model.JokeInfo{}
	var baseUrl = "http://v.juhe.cn/joke/content/list.php?key=404bed3010a1af3497b7726ac9e9c3c2&page=1&pagesize=20&sort=desc&time=" + strconv.FormatInt(time.Now().Unix(), 10)
	resp, err := http.Get(baseUrl)
	if err != nil {
		fmt.Println("请求失败")
	}
	defer resp.Body.Close()
	result, errs := ioutil.ReadAll(resp.Body)
	if errs != nil {
		return *jokeInfo
	}
	code := json.Unmarshal(result, jokeInfo)
	if code != nil {
		return *jokeInfo
	}
	fmt.Print(jokeInfo)
	return *jokeInfo
}
