package model

import (
	"time"
)

func (JenisKursus) TableName() string {
	return "master_jenis_kursus"
}

type SearchJenisKursus struct {
	Id        string `json:"id,omitempty"`
	CepatKode string `json:"cepat_kode,omitempty"`
	Nama      string `json:"nama,omitempty"`
	RefSimpeg string `json:"ref_simpeg,omitempty"`
}

type DeleteJenisKursus struct {
	Id string `json:"id" validate:"required"`
}

type JenisKursus struct {
	Id        string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	CepatKode *string    `json:"cepat_kode"`
	Nama      string     `json:"nama"`
	RefSimpeg *string    `json:"ref_simpeg"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
