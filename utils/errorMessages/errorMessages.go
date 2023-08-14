package errormessages

const (
	SUCCESS = 200
	ERROR   = 400

	//code = 1000.。。用户模块错误
	ERROR_USERNAME_USED        = 1001
	ERROR_PASSWORD_WRONG       = 1002
	ERROR_USER_NOT_EXIST       = 1003
	ERROR_USER_NO_RIGHT        = 1004
	ERROR_USER_ID_ERROR        = 1005
	ERROR_UPDATA_USERINFO_FAIL = 1006

	//token
	ERROR_TOKEN_GENTOKEN_ERROR   = 2001
	ERROR_TOKEN_PARSETOKEN_ERROR = 2002
	ERROR_TOKEN_INVALIDTOKEN     = 2003
	ERROR_TOKEN_ROLE_ERROR       = 2004
	ERROR_TOKEN_HEADER_ERROR     = 2005
	ERROR_TOKEN_TIMA_PASS        = 2006

	//
	ERROR_BIANCHENG_LIKE_SEARCH_VOID = 3001

	//comment
	ERROR_COMMENT_GET_ERROR = 4001
	ERROR_ADD_COMMENT_ERROR = 4002
	//replyComment
	ERROR_REPLYCOMMENT_GET_ERROR   = 4102
	ERROR_REPLYCOMMENT_GET_SUCCESS = 4101
	ERROR_ADD_REPLYCOMMENT_ERROR   = 4103
	ERROR_ADD_REPLYCOMMENT_SUCCESS = 4104
	//register
	USERNAME_NULL       = 5001
	PASSWORD_NULL       = 5002
	EMAIL_NULL          = 5003
	CODE_ERROR          = 5004
	PHONE_NULL          = 5005
	CHECKPASSWORD_ERROR = 5006
	SEND_CODE_SUCCESS   = 5007
	SEND_CODE_ERROE     = 5008

	//bianchengdaohang
	SUCCESS_DELETBIANCHENG     = 6001
	ERROR_DELETBIANCHENG_ERROR = 6002

	//blog
	SUCCESS_ADDBLOGAUTHOR = 7001
	ERROR_ADDBLOGAUTHOR   = 7002
	SUCCESS_GETALLBLOG    = 7003
	ERROR_GETALLBLOG      = 7004
	SUCCESS_GETBLOGAUTHOR = 7005
	ERROR_GETBLOGAUTHOR   = 7006

	SUCCESS_RELEASEBLOG = 7007
	ERROR_RELEASEBLOG   = 7008

	ERROR_FINDTHISBLOGARTICL = 7009

	ERROR_DELETBLOG   = 7010
	SUCCESS_DELETBLOG = 7011

	ERROR_SAVECODE   = 8001
	SUCCESS_SAVECODE = 8002
)

var codeMsg = map[int]string{
	SUCCESS:                          "OK",
	ERROR:                            "FAIL",
	ERROR_USERNAME_USED:              "用户名已存在",
	ERROR_PASSWORD_WRONG:             "密码错误",
	ERROR_USER_NOT_EXIST:             "用户名已存在",
	ERROR_USER_NO_RIGHT:              "该用户无权限",
	ERROR_TOKEN_GENTOKEN_ERROR:       "token生成错误",
	ERROR_TOKEN_PARSETOKEN_ERROR:     "token解析失败",
	ERROR_TOKEN_INVALIDTOKEN:         "token验证失败",
	ERROR_TOKEN_ROLE_ERROR:           "token权限错误",
	ERROR_TOKEN_HEADER_ERROR:         "token请求头错误",
	ERROR_TOKEN_TIMA_PASS:            "token过期",
	ERROR_USER_ID_ERROR:              "用户ID错误",
	ERROR_UPDATA_USERINFO_FAIL:       "更新用户信息失败",
	ERROR_BIANCHENG_LIKE_SEARCH_VOID: "编程点赞，数据库查询为空",
	ERROR_COMMENT_GET_ERROR:          "获取评论失败",
	ERROR_ADD_COMMENT_ERROR:          "添加评论失败",
	USERNAME_NULL:                    "用户名不能为空",
	PASSWORD_NULL:                    "密码不能为空",
	EMAIL_NULL:                       "邮箱不能为空",
	CODE_ERROR:                       "验证码错误",
	PHONE_NULL:                       "电话号码不能为空",
	CHECKPASSWORD_ERROR:              "两次密码不一致",
	SEND_CODE_SUCCESS:                "验证码发生成功",
	SEND_CODE_ERROE:                  "验证码发送失败",
	SUCCESS_DELETBIANCHENG:           "删除编程导航成功",
	ERROR_DELETBIANCHENG_ERROR:       "删除编程导航失败",
	SUCCESS_ADDBLOGAUTHOR:            "添加博客作者成功",
	ERROR_ADDBLOGAUTHOR:              "添加博客作者失败",
	SUCCESS_GETALLBLOG:               "获取所有博客成功",
	ERROR_GETALLBLOG:                 "获取所有博客失败",
	SUCCESS_GETBLOGAUTHOR:            "获取博客作者成功",
	ERROR_GETBLOGAUTHOR:              "获取博客作者失败",
	SUCCESS_RELEASEBLOG:              "文章发布成功",
	ERROR_RELEASEBLOG:                "文章发布失败",
	ERROR_FINDTHISBLOGARTICL:         "查询博客文章失败",
	ERROR_DELETBLOG:                  "博客删除失败",
	SUCCESS_DELETBLOG:                "博客删除成功",
	ERROR_SAVECODE:                   "保存验证码失败，联系站长修复后再验证",
	SUCCESS_SAVECODE:                 "保存验证码成功，可以验证",
	ERROR_REPLYCOMMENT_GET_ERROR:     "获取回复失败",
	ERROR_REPLYCOMMENT_GET_SUCCESS:   "获取回复成功",
	ERROR_ADD_REPLYCOMMENT_ERROR:     "回复失败",
	ERROR_ADD_REPLYCOMMENT_SUCCESS:   "回复成功",
}

// 获取错误信息
func GetErrMsg(code int) string {
	return codeMsg[code]
}
