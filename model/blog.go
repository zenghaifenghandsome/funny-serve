package model

import (
	errormessages "funny-serve/utils/errorMessages"
	"log"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	UserId          int    `gorm:"type:int;not null" json:"userid"`
	Username        string `gorm:"type:varchar(30);not null" json:"username"`
	Useravator      string `gorm:"type:varchar(100);not null" json:"useravator"`
	AuthorId        int    `gorm:"type:int;not null" json:"authorid"`
	AuthorName      string `gorm:"type:varchar(30);not null" json:"authorname"`
	AuthorAvator    string `gorm:"type:varchar(100);not null" json:"authoravator"`
	BlogTags        string `gorm:"type:varchar(200)" json:"blogtags"`
	BlogArticlePic  string `gorm:"type:varchar(100);not null" json:"blogarticlepic"`
	BlogTitle       string `gorm:"type:varchar(50);not null" json:"blogtitle" `
	BlogCategory    string `gorm:"type:varchar(30);not null" json:"blogcategory"`
	BlogIntroduce   string `gorm:"type:varchar(200);not null" json:"blogintroduce"`
	BlogDetail      string `gorm:"type:text;not null" json:"blogdetail"`
	BlogReadNumb    int    `gorm:"type:int;default: 0" json:"blogreadnumb"`
	BlogLikeNumb    int    `gorm:"type:int;default: 0" json:"bloglikenumb"`
	BlogCommentNumb int    `gorm:"type:int;default: 0" json:"blogcommentnumb"`
	Status          int    `gorm:"type:int; default: 0" json:"status"`
}

type BlogAuthor struct {
	gorm.Model
	UserId       string `gorm:"type:int;not null" json:"userid"`
	AuthorName   string `gorm:"type:varchar(30);not null" json:"authorname"`
	AuthorAvator string `gorm:"type:varchar(100);not null" json:"authoravator"`
	Title        string `gorm:"type:varchar(50);not null" json:"title"`
	Detail       string `gorm:"type:varchar(500);not null" json:"detail"`
	WechatDetail string `gorm:"type:varchar(300);not null" json:"wechatdetail"`
	WechatCode   string `gorm:"type:varchar(100);not null" json:"wechatcode"`
	Qgroup       string `gorm:"type:varchar(15);not null" json:"qgroup"`
	Wechat       string `gorm:"type:varchar(30);not null" json:"wechat"`
	Gitee        string `gorm:"type:varchar(100);not null" json:"gitee"`
}

func AddBlogAuthor(b *BlogAuthor) int {
	res := db.Create(&b)
	if res.Error != nil {
		log.Fatal(err)
		return errormessages.ERROR_ADDBLOGAUTHOR
	}
	return errormessages.SUCCESS_ADDBLOGAUTHOR
}

func FindBlog() ([]Blog, int) {

	var blogs []Blog
	res := db.Find(&blogs)

	if res.Error != nil {
		log.Fatal(err)
		return nil, errormessages.ERROR_GETALLBLOG
	}
	return blogs, errormessages.SUCCESS_GETALLBLOG
}

func GetAllOkBlogs() ([]Blog, int) {
	var blogs []Blog
	result := db.Where("status = ?", 1).Find(&blogs)
	if result.Error != nil {
		return nil, errormessages.ERROR
	}
	return blogs, errormessages.SUCCESS
}

func FindOneBlogAuthor(id string) (BlogAuthor, int) {
	var blogAuthor BlogAuthor
	result := db.Where("id = ?", id).First(&blogAuthor)
	if result.Error != nil {
		log.Fatal(err)
		return blogAuthor, errormessages.ERROR_GETBLOGAUTHOR
	}
	//fmt.Printf("%v", result)
	return blogAuthor, errormessages.SUCCESS_GETBLOGAUTHOR
}

func AddBlog(blog *Blog) int {
	res := db.Create(&blog)
	if res.Error != nil {
		log.Fatal(err)
		return errormessages.ERROR_RELEASEBLOG
	}
	return errormessages.SUCCESS_RELEASEBLOG
}

func FinOneBlogArticle(blogID string) (Blog, int) {
	var blog Blog
	result := db.Where("id = ?", blogID).First(&blog)
	if result.Error != nil {
		//log.Fatal(result.Error)
		return blog, errormessages.ERROR_FINDTHISBLOGARTICL
	}
	return blog, errormessages.SUCCESS
}

func DeletBlog(blogID string) int {
	var blog Blog
	result := db.Where("ID=?", blogID).Delete(&blog)
	if result.Error != nil {
		return errormessages.ERROR_DELETBLOG
	} else {
		return errormessages.SUCCESS_DELETBLOG
	}
}
func SetBlogStatusOk(id string) int {
	var blog Blog
	db.Model(&blog).Where("id=?", id).Update("status", 1)
	return errormessages.SUCCESS
}

func SetBlogStatusNo(id string) int {
	var blog Blog
	db.Model(&blog).Where("id=?", id).Update("status", 0)
	return errormessages.SUCCESS
}
