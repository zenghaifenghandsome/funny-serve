package api

import (
	"funny-serve/model"
	errormessages "funny-serve/utils/errorMessages"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	code := model.AddComment(&data)
	if code != errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errormessages.GetErrMsg(code),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormessages.GetErrMsg(code),
	})
}

func GetComment(c *gin.Context) {
	comId := c.Query("id")
	var coms []model.Comment
	coms, code = model.GetComment(comId)
	if code != errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errormessages.GetErrMsg(code),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormessages.GetErrMsg(code),
		"data":   coms,
	})
}

// get replyComments
func GetReplyComments(c *gin.Context) {
	replyCommentID := c.Query("bcCommentID")
	var replyComments []model.CommentReply
	replyComments, code = model.GetCommentReply(replyCommentID)
	if code != errormessages.ERROR_REPLYCOMMENT_GET_SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errormessages.GetErrMsg(code),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormessages.GetErrMsg(code),
		"data":   replyComments,
	})

}

// add replyComment
func AddReplyComment(c *gin.Context) {
	var replyComment model.CommentReply
	_ = c.ShouldBindJSON(&replyComment)
	code := model.AddCommentReply(&replyComment)
	if code != errormessages.ERROR_ADD_REPLYCOMMENT_SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errormessages.GetErrMsg(code),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormessages.GetErrMsg(code),
	})
}
