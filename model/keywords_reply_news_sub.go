package model

import (
	"github.com/jinzhu/gorm"
	"golangWeixin/common"
)

type KeywordsReplyNewsSub struct {
	gorm.Model
	Title       string
	Description string
	PicUrl      string
	Url         string
	ReplyId string `gorm:"column:reply_id"`
}



func _listPageReplyNews(status bool, replyId string) ([]*KeywordsReply, error) {
	var pages []*KeywordsReply
	var err error
	if status {
		err = common.DB.Where("status = ?", StatusNormal).Where("key = ? ", replyId).Find(&pages).Error
	} else {
		err = common.DB.Where("key = ? ", replyId).Find(&pages).Error
	}
	return pages, err
}

func FindAllKeysReplyNewsPage(replyId string) ([]*KeywordsReply, error) {
	return _listPageReplyNews(true, replyId);
}