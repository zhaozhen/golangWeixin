package model

import (
	"golangWeixin/common"
	"time"
)

type KeywordsReplyNewsSub struct {
	Model
	Title       string
	Description string
	PicUrl      string
	Url         string
	ReplyId     string `gorm:"column:reply_id"`
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

func (keyReplySub *KeywordsReplyNewsSub) Insert(person string) error {
	keyReplySub.CreatedAt = time.Now()
	keyReplySub.CreatedPerson = person
	keyReplySub.Status = StatusNormal
	return common.DB.Create(keyReplySub).Error
}

func (keyReplySub *KeywordsReplyNewsSub) Update(person string) error {
	keyReplySub.UpdatedPerson = person
	keyReplySub.UpdatedAt = time.Now()
	return common.DB.Save(keyReplySub).Error
}

func (keyReplySub *KeywordsReplyNewsSub) Delete(person string) error {
	delteDate := time.Now()
	keyReplySub.DeletedAt = &delteDate
	keyReplySub.DeletedPerson = person
	keyReplySub.Status = StatusDelete
	return common.DB.Save(keyReplySub).Error
}

func FindKeywordsReplyNewsSubByReplyId(validStatus bool, Id string) (*[]KeywordsReplyNewsSub, error) {
	var keySub *[]KeywordsReplyNewsSub
	var err error
	if validStatus {
		err = common.DB.Where("status = ?", StatusNormal).Where("reply_id = ? ", Id).Find(&keySub).Error
	} else {
		err = common.DB.Where("id = ? ", Id).Find(&keySub).Error
	}
	return keySub, err
}
