package model

import (
	"time"
)

func (RiwayatOrganisasi) TableName() string {
	return "master_riwayat_organisasi"
}

type SearchRiwayatOrganisasi struct {
	Nip   string `json:"nip,omitempty"`
	Norg  string `json:"norg,omitempty"`
	Jborg string `json:"jborg,omitempty"`
}

type DeleteRiwayatOrganisasi struct {
	Nip   string `json:"nip" validate:"required"`
	Norg  string `json:"norg" validate:"required"`
	Jborg string `json:"jborg" validate:"required"`
}

type RiwayatOrganisasi struct {
	Nip       string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Jorg      *string    `json:"jorg"`
	Norg      string     `json:"norg" validate:"required"`
	Jborg     string     `json:"jborg" validate:"required"`
	Tmulai    *time.Time `json:"tmulai"`
	Takhir    *time.Time `json:"takhir"`
	Npimp     *string    `json:"npimp"`
	Tempat    *string    `json:"tempat"`
	Filename  *string    `json:"filename"`
	IdSync    *string    `gorm:"column:idSync" json:"idSync"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
