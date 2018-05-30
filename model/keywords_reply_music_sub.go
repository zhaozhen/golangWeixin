package model

import (
	"golangWeixin/common"
)

type KeywordsReplyMusicSub struct {
	Model
	Title        string `gorm:"column:title"`
	Description  string `gorm:"column:description"`
	MusicUrl     string `gorm:"column:music_url"`
	HqMusicUrl   string `gorm:"column:hq_music_url"`
	ThumbMediaId string `gorm:"column:thumb_media_id"`
	ReplyId      string `gorm:"column:reply_id"`
}

func (keyReplyMusic *KeywordsReplyMusicSub) Insert() error {
	return common.DB.Create(keyReplyMusic).Error
}

func (keyReplyMusic *KeywordsReplyMusicSub) Update() error {
	return common.DB.Save(keyReplyMusic).Error
}

func _listPageReplyMusic(status bool, replyId string) ([]*KeywordsReply, error) {
	var pages []*KeywordsReply
	var err error
	if status {
		err = common.DB.Where("status = ?", StatusNormal).Where("key = ? ", replyId).Find(&pages).Error
	} else {
		err = common.DB.Where("key = ? ", replyId).Find(&pages).Error
	}
	return pages, err
}

func FindAllKeysReplyMusicPage(replyId string) ([]*KeywordsReply, error) {
	return _listPageReplyMusic(true, replyId)
}