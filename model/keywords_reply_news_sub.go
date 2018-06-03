package model

import (
	"golangWeixin/common"
)

type KeywordsReplyNewsSub struct {
	Model
	Title       string
	Description string
	PicUrl      string
	Url         string
	ReplyId string `gorm:"column:reply_id"`
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

func (keyReplyVideo *KeywordsReplyNewsSub) Insert() error {
	return common.DB.Create(keyReplyVideo).Error
}

func (keyReplyVideo *KeywordsReplyNewsSub) Update() error {
	return common.DB.Save(keyReplyVideo).Error
}


func FindKeywordsReplyNewsSubByReplyId(validStatus bool,Id string)(*[]KeywordsReplyNewsSub, error)  {
	var keySub *[]KeywordsReplyNewsSub
	var err error
	if validStatus {
		err = common.DB.Where("status = ?", StatusNormal).Where("reply_id = ? ", Id).Find(&keySub).Error
	}else{
		err = common.DB.Where("id = ? ", Id).Find(&keySub).Error
	}
	return keySub,err
}
