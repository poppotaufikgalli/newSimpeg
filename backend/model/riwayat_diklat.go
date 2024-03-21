package model

import (
	"time"
)

func (RiwayatDiklat) TableName() string {
	return "master_riwayat_diklat"
}

type SearchRiwayatDiklat struct {
	Nip      string    `json:"nip,omitempty"`
	Jdiklat  int       `json:"jdiklat,omitempty"`
	Kdiklat  float64   `json:"kdiklat,omitempty"`
	Tempat   string    `json:"tempat,omitempty"`
	Angkatan string    `json:"angkatan,omitempty"`
	Tmulai   time.Time `json:"tmulai,omitempty"`
}

type DeleteRiwayatDiklat struct {
	Nip     string    `json:"nip" validate:"required"`
	Jdiklat int       `json:"jdiklat" validate:"required"`
	Tmulai  time.Time `json:"tmulai" validate:"required"`
}

type RiwayatDiklat struct {
	Nip       string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Jdiklat   int        `json:"jdiklat" validate:"required"`
	Kdiklat   float64    `json:"kdiklat"`
	Ndiklat   *string    `json:"ndiklat"`
	Tempat    *string    `json:"tempat"`
	Pan       *string    `json:"pan"`
	Sebagai   *string    `json:"sebagai"`
	Angkatan  *string    `json:"angkatan"`
	Tmulai    *time.Time `json:"tmulai" validate:"required"`
	Takhir    *time.Time `json:"takhir"`
	Jam       float64    `json:"jam"`
	Nsttpp    *string    `json:"nsttpp"`
	Tsttpp    *time.Time `json:"tsttpp"`
	Akhir     *int       `json:"akhir"`
	Filename  *string    `json:"filename"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
