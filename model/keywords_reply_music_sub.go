package model

import (
	"golangWeixin/common"
	"time"
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

func (keyReplyMusic *KeywordsReplyMusicSub) Insert(person string) error {
	keyReplyMusic.CreatedAt = time.Now()
	keyReplyMusic.CreatedPerson = person
	keyReplyMusic.Status = StatusNormal
	return common.DB.Create(keyReplyMusic).Error
}

func (keyReplyMusic *KeywordsReplyMusicSub) Update(person string) error {
	keyReplyMusic.UpdatedPerson = person
	keyReplyMusic.UpdatedAt = time.Now()
	return common.DB.Save(keyReplyMusic).Error
}

func (keyReplyMusic *KeywordsReplyMusicSub) Delete(person string) error {
	delteDate := time.Now()
	keyReplyMusic.DeletedAt = &delteDate
	keyReplyMusic.DeletedPerson = person
	keyReplyMusic.Status = StatusDelete
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


func FindKeywordsReplyMusicSubByReplyId(validStatus bool,Id string)(*KeywordsReplyMusicSub, error)  {
	var keySub *KeywordsReplyMusicSub
	var err error
	if validStatus {
		err = common.DB.Where("status = ?", StatusNormal).Where("reply_id = ? ", Id).Find(&keySub).Error
	}else{
		err = common.DB.Where("id = ? ", Id).Find(&keySub).Error
	}
	return keySub,err
}