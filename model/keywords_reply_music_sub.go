package model

import (
	"time"
	"golangWeixin/common"
)

type KeywordsReplyMusicSub struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	MusicUrl     string    `json:"music_url"`
	HqMusicUrl   string    `json:"hq_music_url"`
	ThumbMediaId string    `json:"thumb_media_id"`
	Status       int       `json:"status"`
	CreatDate    time.Time `json:"creat_date" time_format:"sql_datetime" time_utc:"false"`
	CreatePerson string    `json:"create_person"`
	UpdateDate   time.Time `json:"update_date"`
	DeleteAt     time.Time `json:"delete_at"`
	UpdatePerson string    `json:"update_person"`
}


func (keywordsReplyMusicSub KeywordsReplyMusicSub) findAll()([]KeywordsReply){
	reply := make([]KeywordsReply, 0)
	common.DB.Where("status = ?" ,StatusNormal).Find(&reply);
	return reply;
}
