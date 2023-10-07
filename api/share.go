package api

import (
	"fmt"
	"funny-serve/model"
	errormessages "funny-serve/utils/errorMessages"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddHare(c *gin.Context) {
	var share model.Share
	_ = c.ShouldBindJSON(&share)
	code := model.AddHare(&share)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormessages.GetErrMsg(code),
	})
}

func GetAllShare(c *gin.Context) {
	var share []model.Share
	share, code = model.GetAllShare()
	if code == errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errormessages.GetErrMsg(code),
			"data":    share,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errormessages.GetErrMsg(code),
			"data":    share,
		})
	}
}

// topic
func AddTopic(c *gin.Context) {
	var topic model.Topic
	_ = c.ShouldBindJSON(&topic)
	fmt.Println(topic)
	code := model.AddTopic(&topic)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormessages.GetErrMsg(code),
	})
}

func GetTopicByType(c *gin.Context) {
	var topics []model.Topic
	topics, code = model.GetTopicByType(c.Query("topictype"))
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormessages.GetErrMsg(code),
		"data":   topics,
	})
}
