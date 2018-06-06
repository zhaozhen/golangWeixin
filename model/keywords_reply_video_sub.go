package model

import (
	"golangWeixin/common"
	"time"
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


func (keywordsReplyVideoSub *KeywordsReplyVideoSub) Insert(person string) error {
	keywordsReplyVideoSub.CreatedAt = time.Now()
	keywordsReplyVideoSub.CreatedPerson = person
	keywordsReplyVideoSub.Status = StatusNormal
	return common.DB.Create(keywordsReplyVideoSub).Error
}

func (keywordsReplyVideoSub *KeywordsReplyVideoSub) Update(person string) error {
	keywordsReplyVideoSub.UpdatedPerson = person
	keywordsReplyVideoSub.UpdatedAt = time.Now()
	return common.DB.Save(keywordsReplyVideoSub).Error
}

func (keywordsReplyVideoSub *KeywordsReplyVideoSub) Delete(person string) error {
	delteDate := time.Now()
	keywordsReplyVideoSub.DeletedAt = &delteDate
	keywordsReplyVideoSub.DeletedPerson = person
	keywordsReplyVideoSub.Status = StatusDelete
	return common.DB.Save(keywordsReplyVideoSub).Error
}

func FindKeywordsReplyVideoSubByReplyId(validStatus bool,Id string)(*KeywordsReplyVideoSub, error)  {
	var keySub *KeywordsReplyVideoSub
	var err error
	if validStatus {
		err = common.DB.Where("status = ?", StatusNormal).Where("reply_id = ? ", Id).Find(&keySub).Error
	}else{
		err = common.DB.Where("id = ? ", Id).Find(&keySub).Error
	}
	return keySub,err
}