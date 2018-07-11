package model

import (
	"golangWeixin/common"
	"time"
	"github.com/jinzhu/gorm"
)

type KeywordsReply struct {
	Model
	//ID      int `gorm:"primary_key;column:id"`
	//Status  int    `gorm:"column:status"`
	Key       string `gorm:"column:key_word"`
	MsgType   int    `gorm:"column:msg_type"` // text,image,voice, --video,music,news
	Value     string `gorm:"column:value"`
	AccountId string `gorm:"column:account_id"`
	RawId     string `gorm:"column:raw_id"`
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

func FindAllKeysReplyPage(key string, page int, limit int) (*[]KeywordsReply, error) {
	return _listPage(true, key, page, limit)
}

func _listPage(status bool, key string, page int, limit int) (*[]KeywordsReply, error) {
	var pages []KeywordsReply
	var err error
	if status {
		err = common.DB.Where("status = ?", StatusNormal).Where(" key_word like ? ", "%"+key+"%").Offset(page * limit).Limit(limit).Find(&pages).Error
	} else {
		err = common.DB.Where(" key_word like ? ", "%"+key+"%").Offset(page * limit).Limit(limit).Find(&pages).Error
	}
	return &pages, err
}

func FindKeyWordReplyByOne(validStatus bool, Id int) (*KeywordsReply, error) {
	var key KeywordsReply
	var err error
	if validStatus {
		err = common.DB.Where("status = ?", StatusNormal).Where("id = ? ", Id).Find(&key).Error
	} else {
		err = common.DB.Where("id = ? ", Id).Find(&key).Error
	}
	return &key, err
}

func FindKeyWordReplyByKey(keyWord string) (*KeywordsReply, error) {
	var key KeywordsReply
	var err error
	if err := common.DB.Where("status = ?", StatusNormal).Where("key_word = ? ", keyWord).Find(&key).Error; err != nil {
		return nil, err
	}
	return &key, err
}

// 按道理这个样子只是单个创建，如果多个操作要涉及事物的操作，用的再说
func (keyReply *KeywordsReply) Insert(tx *gorm.DB, person string) error {
	keyReply.CreatedAt = time.Now()
	keyReply.CreatedPerson = person
	keyReply.Status = StatusNormal
	//冗余account_id
	keyReply.AccountId = "default"
	keyReply.RawId = "default"
	if err := tx.Create(keyReply).Error; err != nil {
		tx.Rollback()
		return err
	} else {
		return nil
	}

}

// update
func (keyReply *KeywordsReply) Update(tx *gorm.DB, person string) error {
	updateDate := time.Now()
	keyReply.UpdatedPerson = person
	keyReply.UpdatedAt = &updateDate
	//骚操作，gorm字段等于0居然不更新。。。我。。。
	if keyReply.MsgType == 0 {
		tx.Table("keywords_reply").Model(&keyReply).Update("msg_type", 0)
	}
	if err := tx.Table("keywords_reply").Where("id = ? ", keyReply.ID).Update(keyReply).Error; err != nil {
		tx.Rollback()
		return err
	} else {
		return nil
	}
}

func (keyReply *KeywordsReply) Delete(tx *gorm.DB, person string) error {
	delteDate := time.Now()
	keyReply.DeletedAt = &delteDate
	keyReply.DeletedPerson = person
	keyReply.Status = StatusDelete
	if err := tx.Table("keywords_reply").Where("id = ? ", keyReply.ID).Update(keyReply).Error; err != nil {
		tx.Rollback()
		return err
	} else {
		return nil
	}
}
