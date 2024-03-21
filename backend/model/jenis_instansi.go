package model

import (
	"time"
)

func (JenisInstansi) TableName() string {
	return "master_jenis_instansi"
}

type SearchJenisInstansi struct {
	Id    string   `json:"id,omitempty"`
	Nama  string   `json:"nama,omitempty"`
	Jenis []string `json:"jenis,omitempty"`
}

type DeleteJenisInstansi struct {
	Id string `json:"id" validate:"required"`
}

type JenisInstansi struct {
	Id        string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama      string     `json:"nama"`
	Jenis     string     `json:"jenis"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
