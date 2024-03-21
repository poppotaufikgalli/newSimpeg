package model

import (
	"time"
)

func (RiwayatPmk) TableName() string {
	return "master_riwayat_pmk"
}

type SearchRiwayatPmk struct {
	Nip         string    `json:"nip,omitempty"`
	Dinilai     float64   `json:"dinilai,omitempty"`
	TanggalAwal time.Time `gorm:"column:tanggalAwal" json:"tanggalAwal,omitempty"`
}

type DeleteRiwayatPmk struct {
	Nip         string    `json:"nip" validate:"required"`
	TanggalAwal time.Time `gorm:"column:tanggalAwal" json:"tanggalAwal" validate:"required"`
}

type RiwayatPmk struct {
	Nip            string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Dinilai        *float64   `json:"dinilai"`
	TanggalAwal    time.Time  `gorm:"column:tanggalAwal" json:"tanggalAwal" validate:"required"`
	TanggalAkhir   *time.Time `gorm:"column:tanggalAkhir" json:"tanggalAkhir"`
	MasaKerjaBulan *int       `gorm:"column:masaKerjaBulan" json:"masaKerjaBulan"`
	MasaKerjaTahun *int       `gorm:"column:masaKerjaTahun" json:"masaKerjaTahun"`
	Ptetap         *int       `gorm:"column:ptetap" json:"ptetap"`
	NomorSk        *string    `gorm:"column:nomorSk" json:"nomorSk"`
	TanggalSk      *time.Time `gorm:"column:tanggalSk" json:"tanggalSk"`
	Ninstansi      *string    `json:"ninstansi"`
	Pengalaman     *string    `json:"pengalaman"`
	NomorBkn       *string    `gorm:"column:nomorBkn" json:"nomorBkn"`
	TanggalBkn     *time.Time `gorm:"column:tanggalBkn" json:"tanggalBkn"`
	Filename       *string    `json:"filename"`
	IdSync         *string    `gorm:"column:idSync" json:"idSync"`
	CreatedBy      string     `gorm:"<-:create" json:"created_by"`
	CreatedAt      *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy      string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt      *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
