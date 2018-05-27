package model

import (
	"golangWeixin/common"
)

type KeywordsReply struct {
	Model
	Key     string `gorm:"column:key"`
	MsgType int    `gorm:"column:msg_type"` // text,image,voice, --video,music,news
	Value   string `gorm:"column:value"`
}

//处理表名复数
func (KeywordsReply) TableName() string {
	return "keywords_reply"
}

const (
	KeywordsReplyMsgText  = 0
	KeywordsReplyMsgImage = 1
	KeywordsReplyMsgVoice = 2
	KeywordsReplyMsgVideo = 3
	KeywordsReplyMsgMusic = 4
	KeywordsReplyMsgNews  = 5
)

func FindAllKeysReplyPage(key string,page int,limit int) ([]*KeywordsReply, error) {
	return _listPage(true, key,page,limit)
}

func _listPage(status bool, key string,page int,limit int) ([]*KeywordsReply, error) {
	var pages []*KeywordsReply
	var err error
	if status {
		err = common.DB.Where("status = ?", StatusNormal).Where("key like ? ", "%"+key+"%").Offset(page * limit).Limit(limit).Find(&pages).Error
	} else {
		err = common.DB.Where("key like ? ", "%"+key+"%").Offset(page * limit).Limit(limit).Find(&pages).Error
	}
	return pages, err
}

// 按道理这个样子只是单个创建，如果多个操作要涉及事物的操作，用的再说
func (keyReply *KeywordsReply) Insert() error {
	return common.DB.Create(keyReply).Error
}

// update
func (keyreply *KeywordsReply) Update() error {
	return common.DB.Save(keyreply).Error
}
