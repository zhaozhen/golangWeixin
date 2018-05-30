package model

import (
	"github.com/gin-gonic/gin/json"
	"github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model
	ID     string `gorm:"primary_key;column:id"`
	Status int  `gorm:"column:status"`
	//CreatedAt     time.Time  `gorm:"column:created_at"`
	CreatedPerson string `gorm:"column:created_person"`
	//UpdatedAt     time.Time  `gorm:"column:updated_at"`
	UpdatedPerson string `gorm:"column:updated_person"`
	//DeletedAt     *time.Time `gorm:"column:deleted_at"`
	DeletedPerson string `gorm:"column:deleted_person"`
}

const (
	//正常的状态，
	StatusNormal = 0

	// 被删除的状态
	StatusDelete = 9
)

const (
	SystemUser = "SYSTEM"
)

func (this Model) MarshalJSON() ([]byte, error) {
	// 定义一个该结构体的别名
	type BaseModel Model
	// 定义一个新的结构体
	tmpStudent := struct {
		BaseModel
		CreatedAt string
		UpdatedAt string
		DeletedAt string
	}{
		BaseModel: (BaseModel)(this),
		CreatedAt: this.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: this.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt: this.DeletedAt.Format("2006-01-02 15:04:05"),
	}
	return json.Marshal(tmpStudent)
}
