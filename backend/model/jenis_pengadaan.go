package model

import (
	"time"
)

func (JenisPengadaan) TableName() string {
	return "master_jenis_pengadaan"
}

type SearchJenisPengadaan struct {
	Id   string `json:"id,omitempty"`
	Nama string `json:"nama,omitempty"`
}

type DeleteJenisPengadaan struct {
	Id string `json:"id" validate:"required"`
}

type JenisPengadaan struct {
	Id        string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama      string     `json:"nama"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
