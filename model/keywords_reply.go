package model

import (
	"time"
)

type KeywordsReply struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	Key          string    `json:"key"`
	MsgType      int       `json:"msg_type"` // text,image,voice, --video,music,news
	Value        string    `json:"value"`
	Status       int       `json:"status"`
	CreatDate    time.Time `json:"creat_date" time_format:"sql_datetime" time_utc:"false"`
	CreatePerson string    `json:"create_person"`
	UpdateDate   time.Time `json:"update_date"`
	DeleteAt     time.Time `json:"delete_at"`
	UpdatePerson string    `json:"update_person"`
}

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


////得到所有有效的记录
//func (keywordsReply KeywordsReply) FindAll()(keys []KeywordsReply, err error){
//	reply := make([]KeywordsReply, 0)
//	if err :=common.DB.Where("status = ?" ,StatusNormal).Find(&reply);err != nil{
//		return nil,err.Error;
//	}
//	return reply,nil;
//}


