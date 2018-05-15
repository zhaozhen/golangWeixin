package model

import "time"

type KeywordsReplyNewsSub struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	PicUrl       string    `json:"pic_url"`
	Url          string    `json:"url"`
	Status       int       `json:"status"`
	CreatDate    time.Time `json:"creat_date" time_format:"sql_datetime" time_utc:"false"`
	CreatePerson string    `json:"create_person"`
	UpdateDate   time.Time `json:"update_date"`
	DeleteAt     time.Time `json:"delete_at"`
	UpdatePerson string    `json:"update_person"`
}
