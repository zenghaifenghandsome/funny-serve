package api

import (
	"fmt"
	"funny-serve/model"
	errormessages "funny-serve/utils/errorMessages"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBiancheng(c *gin.Context) {
	var bcs []model.BianCheng
	bcs, code = model.GetAllBiancheng()
	if code == 200 {
		c.JSON(http.StatusOK, gin.H{
			"message": errormessages.GetErrMsg(code),
			"data":    bcs,
			"status":  code,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": errormessages.GetErrMsg(code),
			"data":    bcs,
			"status":  code,
		})
	}
}

func GetOkAllBiancheng(c *gin.Context) {
	var bcs []model.BianCheng
	bcs, code = model.GetOkAllBiancheng()
	if code == 200 {
		c.JSON(http.StatusOK, gin.H{
			"message": errormessages.GetErrMsg(code),
			"data":    bcs,
			"status":  code,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": errormessages.GetErrMsg(code),
			"data":    bcs,
			"status":  code,
		})
	}
}
func SetOk(c *gin.Context) {
	bcid := c.Query("id")
	model.SetOk(bcid)
	c.JSON(http.StatusOK, gin.H{
		"msg": "审核成功",
	})
}
func SetNo(c *gin.Context) {
	bcid := c.Query("id")
	model.SetNo(bcid)
	c.JSON(http.StatusOK, gin.H{
		"msg": "反审成功",
	})
}

// 新增编程导航
func AddBianCheng(c *gin.Context) {
	var biancheng model.BianCheng
	_ = c.ShouldBindJSON(&biancheng)
	fmt.Print(biancheng)
	code = model.CheckBC(biancheng.Title)
	if code != errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    "biancheng 已存在",
		})
	} else {
		err := model.AddBiancheng(&biancheng)
		if err != errormessages.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    "biancheng 写入数据库失败",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"data":   biancheng,
				"msg":    errormessages.GetErrMsg(code),
			})
		}
	}

}

func GetBianchengInfo(c *gin.Context) {
	bcInfoId := c.Query("id")
	fmt.Println(bcInfoId)
	bc, err := model.GetBianchengInfo(bcInfoId)
	if err != errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"msg": "获取失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": err,
		"msg":    errormessages.GetErrMsg(err),
		"bcInfo": bc,
	})

}

func LikeBiancheng(c *gin.Context) {
	var bclike model.BianchengLikes
	_ = c.ShouldBindJSON(&bclike)
	i := model.AddLike(bclike.BianchengId, bclike.UserId)
	if i == errormessages.ERROR_BIANCHENG_LIKE_SEARCH_VOID {
		i2 := model.CreateLikes(&bclike)
		if i2 != errormessages.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code": i2,
				"msg":  errormessages.GetErrMsg(i2),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "点赞创建成功",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": i,
		"msg":    "点赞成功",
	})
}

// delet bianchengdaohang
func DeletBiancheng(c *gin.Context) {
	fmt.Println("进来了")
	bcid := c.Query("id")
	fmt.Println(bcid)
	resultcode := model.DeletBc(bcid)
	if resultcode == errormessages.ERROR_DELETBIANCHENG_ERROR {
		c.JSON(http.StatusOK, gin.H{
			"code": resultcode,
			"msg":  errormessages.GetErrMsg(resultcode),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": resultcode,
			"msg":  errormessages.GetErrMsg(resultcode),
		})
	}
}
