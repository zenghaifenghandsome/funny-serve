package model

import (
	errormessages "funny-serve/utils/errorMessages"

	"gorm.io/gorm"
)

// father-comment
type Comment struct {
	gorm.Model
	BcID       uint   `gorm:"type:int;not null" json:"BcID"`
	Content    string `gorm:"type:varchar(500);not null" json:"Content"`
	UserId     uint   `gorm:"type:int;not null" json:"UserId"`
	UserAvator string `gorm:"type:varchar(500);not null" json:"UserAvator"`
	UserName   string `gorm:"type:varchar(100);not null" json:"UserName"`
	LikeNumb   int    `gorm:"type:int; default: 0" json:"LikeNumb"`
	LowNumb    int    `gomr:"type:int; default :0" json:"LowNumb"`
}

// child-comment
type CommentReply struct {
	gorm.Model
	BcID       uint   `gorm:"type:int;not null" json:"BcID"`
	CommentID  uint   `gorm:"type:int;not null" json:"CommentID"`
	Content    string `gorm:"type:varchar(500);not null" json:"Content"`
	UserId     uint   `gorm:"type:int;not null" json:"UserId"`
	UserAvator string `gorm:"type:varchar(500);not null" json:"UserAvator"`
	UserName   string `gorm:"type:varchar(100);not null" json:"UserName"`
	LikeNumb   int    `gorm:"type:int; default: 0" json:"LikeNumb"`
	LowNumb    int    `gomr:"type:int; default: 0" json:"LowNumb"`
}

// get-comment
func GetComment(bcID string) ([]Comment, int) {
	var comments []Comment
	result := db.Where("bc_id = ?", bcID).Find(&comments)
	if result.Error != nil {
		return nil, errormessages.ERROR_COMMENT_GET_ERROR
	}
	return comments, errormessages.SUCCESS
}

// get-commentReply
func GetCommentReply(bcCommentID string) ([]CommentReply, int) {
	var commentReplys []CommentReply
	result := db.Where("comment_id = ?", bcCommentID).Find(&commentReplys)
	if result.Error != nil {
		return nil, errormessages.ERROR_REPLYCOMMENT_GET_ERROR
	}
	return commentReplys, errormessages.ERROR_REPLYCOMMENT_GET_SUCCESS
}

// add to comment
func AddComment(com *Comment) int {
	result := db.Create(&com)
	if result.Error != nil {
		return errormessages.ERROR_ADD_COMMENT_ERROR
	}
	return errormessages.SUCCESS
}

// add to commentReply
func AddCommentReply(reply *CommentReply) int {
	result := db.Create(&reply)
	if result.Error != nil {
		return errormessages.ERROR_ADD_REPLYCOMMENT_ERROR
	}
	return errormessages.ERROR_ADD_REPLYCOMMENT_SUCCESS
}
