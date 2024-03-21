package model

import (
	"time"
)

func (JenisKompetensi) TableName() string {
	return "master_jenis_kompetensi"
}

type SearchJenisKompetensi struct {
	Id     string `json:"id,omitempty"`
	Nama   string `json:"nama,omitempty"`
	NamaId string `json:"nama_id,omitempty"`
}

type DeleteJenisKompetensi struct {
	Id string `json:"id" validate:"required"`
}

type JenisKompetensi struct {
	Id        string     `gorm:"primaryKey;autoIncrement:false" json:"id"`
	Nama      string     `json:"nama"`
	NamaId    *string    `json:"nama_id"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
