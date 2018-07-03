package model

import (
	"golangWeixin/common"
	"time"
	"github.com/jinzhu/gorm"
)

type KeywordsReplyNewsSub struct {
	Model
	//ID          int `gorm:"primary_key;column:id"`
	//Status      int    `gorm:"column:status"`
	Title       string
	Description string
	PicUrl      string
	Url         string
	ReplyId     int `gorm:"column:reply_id"`
}

func (KeywordsReplyNewsSub) TableName() string {
	return "keywords_reply_news_sub"
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
	return _listPageReplyNews(true, replyId)
}

func (keyReplySub *KeywordsReplyNewsSub) Insert(tx *gorm.DB,person string) error {
	keyReplySub.CreatedAt = time.Now()
	keyReplySub.CreatedPerson = person
	keyReplySub.Status = StatusNormal
	if err:=tx.Create(keyReplySub).Error;err!=nil{
		tx.Rollback()
		return err
	}else {
		return nil
	}
}

func (keyReplySub *KeywordsReplyNewsSub) Update(tx *gorm.DB, person string) error {
	updateDate := time.Now()
	keyReplySub.UpdatedPerson = person
	keyReplySub.UpdatedAt = &updateDate
	if err := tx.Table("keywords_reply_news_sub").Where("reply_id = ? ", keyReplySub.ReplyId).Update(keyReplySub).Error; err != nil {
		tx.Rollback()
		return err
	} else {
		return nil
	}
}

func (keyReplySub *KeywordsReplyNewsSub) Delete(tx *gorm.DB,person string) error {
	delteDate := time.Now()
	keyReplySub.DeletedAt = &delteDate
	keyReplySub.DeletedPerson = person
	keyReplySub.Status = StatusDelete
	if err:=tx.Table("keywords_reply_news_sub").Where("reply_id = ? ",keyReplySub.ReplyId).Update(keyReplySub).Error;err!=nil{
		return err
	} else {
		return nil
	}
}

func FindKeywordsReplyNewsSubByReplyId(validStatus bool, Id int) (*[]KeywordsReplyNewsSub, error) {
	var keySub []KeywordsReplyNewsSub
	var err error
	if validStatus {
		err = common.DB.Where("status = ?", StatusNormal).Where("reply_id = ? ", Id).Find(&keySub).Error
	} else {
		err = common.DB.Where("id = ? ", Id).Find(&keySub).Error
	}
	return &keySub, err
}
