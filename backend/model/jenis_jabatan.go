package model

import (
	"time"
)

func (JenisJabatan) TableName() string {
	return "master_jenis_jabatan"
}

type SearchJenisJabatan struct {
	Id     string `json:"id,omitempty"`
	Nama   string `json:"nama,omitempty"`
	Status []int  `json:"status,omitempty"`
}

type DeleteJenisJabatan struct {
	Id string `json:"id" validate:"required"`
}

type JenisJabatan struct {
	Id        string     `gorm:"primaryKey;autoIncrement:false" json:"id"`
	Nama      string     `json:"nama"`
	RefSimpeg string     `json:"ref_simpeg"`
	RefSiap   string     `json:"ref_siap"`
	IsAk      int        `json:"is_ak"`
	Status    *int       `json:"status"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
