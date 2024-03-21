package model

import (
	"time"
)

func (JenisKeluarga) TableName() string {
	return "master_jenis_keluarga"
}

type SearchJenisKeluarga struct {
	Id     string `json:"id,omitempty"`
	Nama   string `json:"nama,omitempty"`
	Status []int  `json:"status,omitempty"`
}

type DeleteJenisKeluarga struct {
	Id string `json:"id" validate:"required"`
}

type JenisKeluarga struct {
	Id        string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama      string     `json:"nama"`
	Status    *int       `json:"status"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
