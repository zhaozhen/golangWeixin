package model

import (
	"golangWeixin/common"
	"time"
	"github.com/jinzhu/gorm"
)

type KeywordsReplyMusicSub struct {
	Model
	Title        string `gorm:"column:title"`
	Description  string `gorm:"column:description"`
	MusicUrl     string `gorm:"column:music_url"`
	HqMusicUrl   string `gorm:"column:hq_music_url"`
	ThumbMediaId string `gorm:"column:thumb_media_id"`
	ReplyId      int `gorm:"column:reply_id"`
}


func (KeywordsReplyMusicSub) TableName() string {
	return "keywords_reply_music_sub"
}

func (keyReplyMusic *KeywordsReplyMusicSub) Insert(tx *gorm.DB, person string) error {
	keyReplyMusic.CreatedAt = time.Now()
	keyReplyMusic.CreatedPerson = person
	keyReplyMusic.Status = StatusNormal
	if err := tx.Create(keyReplyMusic).Error; err != nil {
		tx.Rollback()
		return err
	} else {
		return nil
	}
}

func (keyReplyMusic *KeywordsReplyMusicSub) Update(tx *gorm.DB,person string) error {
	updateDate := time.Now()
	keyReplyMusic.UpdatedPerson = person
	keyReplyMusic.UpdatedAt = &updateDate
	if err:=tx.Table("keywords_reply_music_sub").Where("reply_id = ? ",keyReplyMusic.ReplyId).Update(keyReplyMusic).Error;err!=nil{
		tx.Rollback()
		return err
	}else {
		return nil
	}

}

func (keyReplyMusic *KeywordsReplyMusicSub) Delete(tx *gorm.DB, person string) error {
	delteDate := time.Now()
	keyReplyMusic.DeletedAt = &delteDate
	keyReplyMusic.DeletedPerson = person
	keyReplyMusic.Status = StatusDelete
	if err := tx.Table("keywords_reply_music_sub").Where("reply_id = ? ", keyReplyMusic.ReplyId).Update(keyReplyMusic).Error; err != nil {
		tx.Rollback()
		return err
	} else {
		return nil
	}
}

func _listPageReplyMusic(status bool, replyId string) (*[]KeywordsReply, error) {
	var pages []KeywordsReply
	var err error
	if status {
		err = common.DB.Where("status = ?", StatusNormal).Where("key = ? ", replyId).Find(&pages).Error
	} else {
		err = common.DB.Where("key = ? ", replyId).Find(&pages).Error
	}
	return &pages, err
}

func FindAllKeysReplyMusicPage(replyId string) (*[]KeywordsReply, error) {
	return _listPageReplyMusic(true, replyId)
}


func FindKeywordsReplyMusicSubByReplyId(validStatus bool,Id int)(*KeywordsReplyMusicSub, error)  {
	var keySub KeywordsReplyMusicSub
	var err error
	if validStatus {
		err = common.DB.Where("status = ?", StatusNormal).Where("reply_id = ? ", Id).Find(&keySub).Error
	}else{
		err = common.DB.Where("id = ? ", Id).Find(&keySub).Error
	}
	return &keySub,err
}