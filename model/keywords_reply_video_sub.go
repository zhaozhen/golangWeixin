package model

import (
	"golangWeixin/common"
	"time"
	"github.com/jinzhu/gorm"
)

type KeywordsReplyVideoSub struct {
	Model
	Title       string
	Description string
	MediaId     string
	ReplyId     int `gorm:"column:reply_id"`
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

func (keywordsReplyVideoSub *KeywordsReplyVideoSub) Insert(tx *gorm.DB, person string) error {
	keywordsReplyVideoSub.CreatedAt = time.Now()
	keywordsReplyVideoSub.CreatedPerson = person
	keywordsReplyVideoSub.Status = StatusNormal
	if err := tx.Save(keywordsReplyVideoSub).Error; err != nil {
		tx.Rollback()
		return err
	} else {
		return nil
	}
}

func (keywordsReplyVideoSub *KeywordsReplyVideoSub) Update(tx *gorm.DB, person string) error {
	updateDate := time.Now()
	keywordsReplyVideoSub.UpdatedPerson = person
	keywordsReplyVideoSub.UpdatedAt = &updateDate
	if err := tx.Table("keywords_reply_video_sub").Where("reply_id = ? ", keywordsReplyVideoSub.ReplyId).Update(keywordsReplyVideoSub).Error; err != nil {
		tx.Rollback()
		return err
	} else {
		return nil
	}
}

func (keywordsReplyVideoSub *KeywordsReplyVideoSub) Delete(tx *gorm.DB, person string) error {
	delteDate := time.Now()
	keywordsReplyVideoSub.DeletedAt = &delteDate
	keywordsReplyVideoSub.DeletedPerson = person
	keywordsReplyVideoSub.Status = StatusDelete
	if err := tx.Table("keywords_reply_video_sub").Where("reply_id = ? ", keywordsReplyVideoSub.ReplyId).Update(keywordsReplyVideoSub).Error; err != nil {
		tx.Rollback()
		return err
	} else {
		return nil
	}
}

func FindKeywordsReplyVideoSubByReplyId(validStatus bool,Id int)(*KeywordsReplyVideoSub, error)  {
	var keySub KeywordsReplyVideoSub
	var err error
	if validStatus {
		err = common.DB.Where("status = ?", StatusNormal).Where("reply_id = ? ", Id).Find(&keySub).Error
	}else{
		err = common.DB.Where("id = ? ", Id).Find(&keySub).Error
	}
	return &keySub,err
}