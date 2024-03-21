package model

import (
	"time"
)

func (JenisKawin) TableName() string {
	return "master_jenis_kawin"
}

type SearchJenisKawin struct {
	Id        string   `json:"id,omitempty"`
	Nama      string   `json:"nama,omitempty"`
	RefSimpeg []string `json:"ref_simpeg,omitempty"`
}

type DeleteJenisKawin struct {
	Id string `json:"id" validate:"required"`
}

type JenisKawin struct {
	Id        string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama      string     `json:"nama"`
	RefSimpeg *string    `json:"ref_simpeg"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
