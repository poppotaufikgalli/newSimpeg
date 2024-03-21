package model

import (
	"time"
)

func (JabatanFtBkn) TableName() string {
	return "master_jabatan_ft_bkn"
}

type SearchJabatanFtBkn struct {
	Id           string   `json:"id,omitempty"`
	Nama         string   `json:"nama,omitempty"`
	KelJabatanId string   `json:"kel_jabatan_id,omitempty"`
	Jenjang      []string `json:"jenjang,omitempty"`
	Status       []string `json:"status,omitempty"`
	CepatKode    string   `json:"cepat_kode,omitempty"`
	RefSimpeg    string   `json:"ref_simpeg,omitempty"`
}

type DeleteJabatanFtBkn struct {
	Id string `json:"id" validate:"required"`
}

type JabatanFtBkn struct {
	Id           string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama         string     `json:"nama"`
	BupUsia      *float64   `json:"bup_usia" validate:"required"`
	KelJabatanId *string    `json:"kel_jabatan_id" validate:"required"`
	Jenjang      *string    `json:"jenjang"`
	Status       string     `json:"status"`
	CepatKode    *string    `json:"cepat_kode"`
	RefSimpeg    *string    `json:"ref_simpeg"`
	CreatedBy    string     `gorm:"<-:create" json:"created_by"`
	CreatedAt    *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy    string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt    *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
