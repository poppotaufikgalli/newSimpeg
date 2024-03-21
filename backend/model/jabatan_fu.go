package model

import (
	"time"
)

func (JabatanFu) TableName() string {
	return "master_jabatan_fu"
}

type SearchJabatanFu struct {
	Id        string `json:"id,omitempty"`
	Nama      string `json:"nama,omitempty"`
	CepatKode string `json:"cepat_kode,omitempty"`
	Status    []int  `json:"status,omitempty"`
	RefSimpeg string `json:"ref_simpeg,omitempty"`
}

type DeleteJabatanFu struct {
	Id string `json:"id" validate:"required"`
}

type JabatanFu struct {
	Id        string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama      string     `json:"nama"`
	CepatKode *string    `json:"cepat_kode"`
	Status    *int       `json:"status"`
	RefSimpeg *string    `json:"ref_simpeg"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
