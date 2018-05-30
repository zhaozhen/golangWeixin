package model

import (
	"golangWeixin/common"
)

type KeywordsReplyVideoSub struct {
	Model
	Title        string
	Description  string
	MediaId string
	ReplyId string `gorm:"column:reply_id"`
}



func _listPageReplyVedio(status bool, replyId string) ([]*KeywordsReply, error) {
	var pages []*KeywordsReply
	var err error
	if status {
		err = common.DB.Where("status = ?", StatusNormal).Where("key = ? ", replyId).Find(&pages).Error
	} else {
		err = common.DB.Where("key = ? ", replyId).Find(&pages).Error
	}
	return pages, err
}

func FindAllKeysReplyVedioPage(replyId string) ([]*KeywordsReply, error) {
	return _listPageReplyVedio(true, replyId)
}


func (keywordsReplyVideoSub *KeywordsReplyVideoSub) Insert() error {
	return common.DB.Create(keywordsReplyVideoSub).Error
}

func (keywordsReplyVideoSub *KeywordsReplyVideoSub) Update() error {
	return common.DB.Save(keywordsReplyVideoSub).Error
}