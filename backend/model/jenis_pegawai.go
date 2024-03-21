package model

import (
	"time"
)

func (JenisPegawai) TableName() string {
	return "master_jenis_pegawai"
}

type SearchJenisPegawai struct {
	Id     string `json:"id,omitempty"`
	Nama   string `json:"nama,omitempty"`
	Status []int  `json:"status,omitempty"`
}

type DeleteJenisPegawai struct {
	Id string `json:"id" validate:"required"`
}

type JenisPegawai struct {
	Id        string     `gorm:"primaryKey;autoIncrement:false" json:"id"`
	Nama      string     `json:"nama"`
	Status    *int       `json:"status"`
	RefSimpeg string     `json:"ref_simpeg"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
