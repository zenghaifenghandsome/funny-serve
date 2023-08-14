package api

import (
	"fmt"
	"funny-serve/model"
	errormessages "funny-serve/utils/errorMessages"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddBlogAuthor(c *gin.Context) {
	var newBlogAuthor model.BlogAuthor
	_ = c.ShouldBindJSON(&newBlogAuthor)
	code := model.AddBlogAuthor(&newBlogAuthor)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errormessages.GetErrMsg(code),
	})
}

func FindBlog(c *gin.Context) {
	result, findBlogErr := model.FindBlog()
	if findBlogErr != errormessages.SUCCESS_GETALLBLOG {
		log.Fatal(findBlogErr)
		c.JSON(http.StatusOK, gin.H{
			"code": findBlogErr,
			"msg":  errormessages.GetErrMsg(findBlogErr),
			"data": result,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": findBlogErr,
			"msg":  errormessages.GetErrMsg(findBlogErr),
			"data": result,
		})
	}

}

func GetAllOkBlogs(c *gin.Context) {
	var blogs []model.Blog
	blogs, code = model.GetAllOkBlogs()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errormessages.GetErrMsg(code),
		"data": blogs,
	})
}

func FindOneBlogAuthor(c *gin.Context) {
	fmt.Println("ok join")
	id := c.Query("authorid")
	result, findOneBlogAuthorErr := model.FindOneBlogAuthor(id)
	if findOneBlogAuthorErr != errormessages.SUCCESS_GETBLOGAUTHOR {
		c.JSON(http.StatusOK, gin.H{
			"code": findOneBlogAuthorErr,
			"msg":  errormessages.GetErrMsg(findOneBlogAuthorErr),
			"data": result,
		})
		log.Fatal(findOneBlogAuthorErr)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": findOneBlogAuthorErr,
			"msg":  errormessages.GetErrMsg(findOneBlogAuthorErr),
			"data": result,
		})
	}
}

func ReleaseBlog(c *gin.Context) {
	var blog model.Blog
	_ = c.ShouldBindJSON(&blog)
	code := model.AddBlog(&blog)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errormessages.GetErrMsg(code),
	})
}
func FindOneBlogArticle(c *gin.Context) {
	blogID := c.Query("blogID")
	result, err := model.FinOneBlogArticle(blogID)
	if err == errormessages.ERROR_FINDTHISBLOGARTICL {
		c.JSON(http.StatusOK, gin.H{
			"code": err,
			"msg":  errormessages.GetErrMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": err,
			"msg":  errormessages.GetErrMsg(err),
			"data": result,
		})
	}
}

func DeletBlog(c *gin.Context) {
	blogID := c.Query("id")
	resultcode := model.DeletBlog(blogID)
	c.JSON(http.StatusOK, gin.H{
		"code": resultcode,
		"msg":  errormessages.GetErrMsg(resultcode),
	})
}

func SetBlogStatusOk(c *gin.Context) {
	blogID := c.Query("id")
	model.SetBlogStatusOk(blogID)
	c.JSON(http.StatusOK, gin.H{
		"msg": "审核成功",
	})

}

func SetBlogStatusNo(c *gin.Context) {
	blogID := c.Query("id")
	model.SetBlogStatusNo(blogID)
	c.JSON(http.StatusOK, gin.H{
		"msg": "反审成功",
	})
}
