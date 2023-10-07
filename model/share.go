package model

import (
	errormessages "funny-serve/utils/errorMessages"

	"gorm.io/gorm"
)

type Share struct {
	gorm.Model
	UserId  uint   `gorm:"type:int;not null" json:"userid"`
	Content string `gorm:"type:text;not null" json:"content"`
	PicList string `gorm:"type:text;default null" json:"piclist"`
	Video   string `gorm:"type:text;default null" json:"video"`
}

type Topic struct {
	gorm.Model
	TopicUuid    string `gorm:"type:text;not null" json:"topicuuid"`
	TopicName    string `gorm:"type:text;not null" json:"topicname"`
	TopicType    string `gorm:"type:text;not null" json:"topictype"`
	TopicUrl     string `gorm:"type:text;not null" json:"topicurl"`
	TopicCreator string `gorm:"type:text;not null" json:"topiccreator"`
	TopicContent string `gorm:"type:text" json:"topiccontent"`
}

func AddHare(share *Share) int {
	res := db.Create(&share)
	err := res.Error
	if err != nil {
		return errormessages.ERROR
	}
	return errormessages.SUCCESS
}

func GetAllShare() ([]Share, int) {
	var share []Share
	result := db.Find(&share)
	if result.Error != nil {
		return nil, errormessages.ERROR
	}
	return share, errormessages.SUCCESS
}

// topic
func AddTopic(topic *Topic) int {
	res := db.Create(&topic)
	err := res.Error
	if err != nil {
		return errormessages.ERROR
	}
	return errormessages.SUCCESS
}

func GetTopicByType(topicType string) ([]Topic, int) {
	var topics []Topic
	result := db.Find(&topics).Where("topic_type = ?", topicType)
	if result.Error != nil {
		return nil, errormessages.ERROR
	}
	return topics, errormessages.SUCCESS
}
