package model

import (
	"time"
)

func (Group) TableName() string {
	return "group"
}

type SearchGroup struct {
	Id   string `json:"id,omitempty"`
	Nama string `json:"nama,omitempty"`
}

type DeleteGroup struct {
	Id string `json:"id" validate:"required"`
}

type Group struct {
	Id        int        `gorm:"primaryKey" json:"id"`
	Nama      string     `json:"nama" validate:"required"`
	Lsakses   string     `json:"lsakses"`
	Users     []Users    `gorm:"foreignKey:Gid" json:"users"`
	CreatedBy string     `json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
