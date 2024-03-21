package model

import (
	"time"
)

func (Instansi) TableName() string {
	return "master_instansi"
}

type SearchInstansi struct {
	Id              string   `json:"id,omitempty"`
	Nama            string   `json:"nama,omitempty"`
	Jenis           []string `json:"jenis,omitempty"`
	CepatKode       string   `json:"cepat_kode,omitempty"`
	JenisInstansiId []string `json:"jenis_instansi_id,omitempty"`
}

type DeleteInstansi struct {
	Id string `json:"id" validate:"required"`
}

type Instansi struct {
	Id              string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama            string     `json:"nama"`
	Jenis           string     `json:"jenis"`
	CepatKode       *string    `json:"cepat_kode"`
	JenisInstansiId *string    `json:"jenis_instansi_id"`
	RefSimpeg       *string    `json:"ref_simpeg"`
	CreatedBy       string     `gorm:"<-:create" json:"created_by"`
	CreatedAt       *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy       string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt       *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
