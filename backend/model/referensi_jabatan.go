package model

import (
	"time"
)

func (ReferensiJabatan) TableName() string {
	return "master_referensi_jabatan"
}

type SearchReferensiJabatan struct {
	Id             string   `json:"id,omitempty"`
	JenisJabatanId int      `json:"jenis_jabatan_id,omitempty"`
	Nama           string   `json:"nama,omitempty"`
	BupUsia        float64  `json:"bup_usia,omitempty"`
	KelJabatanId   []string `json:"kel_jabatan_id,omitempty"`
	Status         []int    `json:"status,omitempty"`
	CepatKode      string   `json:"cepat_kode,omitempty"`
	RefSimpeg      string   `json:"ref_simpeg,omitempty"`
}

type DeleteReferensiJabatan struct {
	Id             string `json:"id" validate:"required"`
	JenisJabatanId int    `json:"jenis_jabatan_id" validate:"required"`
}

type ReferensiJabatan struct {
	Id             string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	JenisJabatanId int        `json:"jenis_jabatan_id" validate:"required"`
	Nama           string     `json:"nama" validate:"required"`
	BupUsia        float64    `json:"bup_usia"`
	KelJabatanId   *string    `json:"kel_jabatan_id"`
	Jenjang        *string    `json:"jenjang"`
	Status         *int       `json:"status"`
	CepatKode      *string    `json:"cepat_kode"`
	RefSimpeg      *string    `json:"ref_simpeg"`
	CreatedBy      string     `gorm:"<-:create" json:"created_by"`
	CreatedAt      *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy      string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt      *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
