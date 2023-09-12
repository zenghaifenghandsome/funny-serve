package router

import (
	"funny-serve/api"
	moddleware "funny-serve/middleware"
	errormessages "funny-serve/utils/errorMessages"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter() {
	gin.SetMode(viper.GetString("App_mode"))
	r := gin.Default()
	r.Use(moddleware.Cors())
	router := r.Group("")
	{
		router.POST("/update", api.UpData)
		router.GET("/main", api.GetMainPage)
	}
	//注册
	registerRouter := r.Group("/register")
	{
		registerRouter.POST("", api.AddUser)
		registerRouter.POST("/code", api.CheckCode)
	}

	//登录
	loginRouter := r.Group("/login")
	{
		loginRouter.POST("", api.Login)
		loginRouter.POST("/userinfo", api.GetUser)
	}
	//user
	userRouter := r.Group("/user")
	{
		userRouter.POST("/updatauserinfo", api.UpDataUserInfo)
		userRouter.GET("/getUserInfo", api.GetUserInfo)
		userRouter.GET("/updataUserInfoOneField", api.UpDataUserInfoOneField)
	}

	//bianchengdaohang
	bianchengRouter := r.Group("/biancheng")
	{
		bianchengRouter.GET("", api.GetAllBiancheng)
		bianchengRouter.POST("", api.AddBianCheng)
		//bianchengRouter.POST("/update", api.UpDate)
		bianchengRouter.GET("/getOnlybiancheng", api.GetBianchengInfo)
		bianchengRouter.POST("/like", api.LikeBiancheng)
		bianchengRouter.POST("/comment", api.AddComment)
		bianchengRouter.GET("/comment", api.GetComment)
		bianchengRouter.GET("/deletBiancheng", api.DeletBiancheng)
		bianchengRouter.GET("/okAllBiancheng", api.GetOkAllBiancheng)
		bianchengRouter.GET("/setOkBiancheng", api.SetOk)
		bianchengRouter.GET("/setNoBiancheng", api.SetNo)
		bianchengRouter.GET("/comment/replyComment", api.GetReplyComments)
		bianchengRouter.POST("/comment/replyComment", api.AddReplyComment)

	}
	evd := r.Group("evd")
	{
		evd.POST("", api.AddEvd)
		evd.GET("", api.GetAllEvd)
	}
	share := r.Group("share")
	{
		share.POST("/add", api.AddHare)
		share.GET("/getAll", api.GetAllShare)
	}

	hell := r.Group("/test")

	{
		//测试token
		hell.POST("", moddleware.JwtAuthMiddleware(), moddleware.CheckAuth("user"), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "token ok",
			})
		})

		hell.POST("/refresh", moddleware.JwtAuthMiddleware(), moddleware.CheckAuth("admin"), func(c *gin.Context) {
			token, code := moddleware.GenToken(c.GetString("username"), moddleware.Passtime, "user")
			if code != errormessages.SUCCESS {
				c.JSON(http.StatusOK, gin.H{
					"msg": errormessages.GetErrMsg(errormessages.ERROR),
				})
				c.Abort()
				return

			}
			c.JSON(http.StatusOK, gin.H{
				"msg":      "token 刷新成功",
				"newToken": token,
			})
		})

		hell.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"msg": "11111111111111111111"}) })
	}

	blogRouter := r.Group("/blog")
	{
		blogRouter.POST("", api.AddBlogAuthor)
		blogRouter.GET("", api.FindBlog)
		blogRouter.GET("/blogAuthor", api.FindOneBlogAuthor)
		blogRouter.POST("/release", api.ReleaseBlog)
		blogRouter.GET("/blogArticle", api.FindOneBlogArticle)
		blogRouter.GET("/deletBlog", api.DeletBlog)
		blogRouter.GET("/setOkBlog", api.SetBlogStatusOk)
		blogRouter.GET("/setNoBlog", api.SetBlogStatusNo)
		blogRouter.GET("/allOkBlog", api.GetAllOkBlogs)
	}
	//fmt.Println(config.HttpPort)
	r.Run(viper.GetString("server.Http_port"))
}
