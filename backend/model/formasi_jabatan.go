package model

import (
	"time"
)

func (FormasiJabatan) TableName() string {
	return "master_formasi_jabatan"
}

type SearchFormasiJabatan struct {
	Id           string `json:"id,omitempty"`
	IdOpd        string `json:"id_opd,omitempty"`
	ParentId     string `json:"parent_id,omitempty"`
	Nama         string `json:"nama,omitempty"`
	KelasJabatan []int  `json:"kelas_jabatan,omitempty"`
	Status       []int  `json:"status,omitempty"`
}

type DeleteFormasiJabatan struct {
	Id string `json:"id" validate:"required"`
}

type FormasiJabatan struct {
	Id             string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	ParentId       string     `json:"parent_id" validate:"required"`
	IdOpd          string     `json:"id_opd" validate:"required"`
	Nama           string     `json:"nama" validate:"required"`
	IdEselon       int        `json:"id_eselon"`
	IdJenisJabatan int        `json:"id_jenis_jabatan"`
	RefJabatanId   *string    `json:"ref_jabatan_id"`
	KelasJabatan   int        `json:"kelas_jabatan"`
	NilaiJabatan   int        `json:"nilai_jabatan"`
	IndeksJabatan  int        `json:"indeks_jabatan"`
	Status         *int       `json:"status"`
	CreatedBy      string     `gorm:"<-:create" json:"created_by"`
	CreatedAt      *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy      string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt      *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
