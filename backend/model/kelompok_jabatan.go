package model

import (
	"time"
)

func (KelompokJabatan) TableName() string {
	return "master_kelompok_jabatan"
}

type SearchKelompokJabatan struct {
	Id              string `json:"id,omitempty"`
	Nama            string `json:"nama,omitempty"`
	JenisJabatanId  string `json:"jenis_jabatan_id,omitempty"`
	RumpunJabatanId string `json:"rumpun_jabatan_id,omitempty"`
	PembinaId       string `json:"pembina_id,omitempty"`
}

type DeleteKelompokJabatan struct {
	Id string `json:"id" validate:"required"`
}

type KelompokJabatan struct {
	Id              string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama            string     `json:"nama"`
	JenisJabatanId  *string    `json:"jenis_jabatan_id"`
	RumpunJabatanId *string    `json:"rumpun_jabatan_id"`
	PembinaId       *string    `json:"pembina_id"`
	CreatedBy       string     `gorm:"<-:create" json:"created_by"`
	CreatedAt       *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy       string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt       *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
