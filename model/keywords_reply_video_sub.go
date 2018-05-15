package model

import "time"

type KeywordsReplyVideoSub struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Status       int       `json:"status"`
	CreatDate    time.Time `json:"creat_date" time_format:"sql_datetime" time_utc:"false"`
	CreatePerson string    `json:"create_person"`
	UpdateDate   time.Time `json:"update_date"`
	UpdatePerson string    `json:"update_person"`
}
