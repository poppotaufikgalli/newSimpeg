package model

import (
	"time"
)

func (Eselon) TableName() string {
	return "master_eselon"
}

type SearchEselon struct {
	Id             string   `json:"id,omitempty"`
	Nama           string   `json:"nama,omitempty"`
	JabatanAsn     string   `json:"jabatan_asn,omitempty"`
	IdJenisJabatan []string `json:"id_jenis_jabatan,omitempty"`
	Status         []int    `json:"status,omitempty"`
}

type DeleteEselon struct {
	Id string `json:"id" validate:"required"`
}

type Eselon struct {
	Id             string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama           string     `json:"nama" validate:"required"`
	JabatanAsn     string     `json:"jabatan_asn"`
	IdJenisJabatan string     `json:"id_jenis_jabatan"`
	Status         int        `json:"status"`
	CreatedBy      string     `gorm:"<-:create" json:"created_by"`
	CreatedAt      *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy      string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt      *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
