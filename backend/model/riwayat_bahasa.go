package model

import (
	"time"
)

func (RiwayatBahasa) TableName() string {
	return "master_riwayat_bahasa"
}

type SearchRiwayatBahasa struct {
	Nip     string   `json:"nip,omitempty"`
	Nbahasa string   `json:"nbahasa,omitempty"`
	Kbahasa []string `json:"kbahasa,omitempty"`
	Jbahasa []string `json:"jbahasa,omitempty"`
}

type DeleteRiwayatBahasa struct {
	Nip     string `json:"nip" validate:"required"`
	Nbahasa string `json:"nbahasa" validate:"required"`
}

type RiwayatBahasa struct {
	Nip       string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Nbahasa   string     `json:"nbahasa" validate:"required"`
	Kbahasa   *string    `json:"kbahasa"`
	Jbahasa   *string    `json:"jbahasa"`
	Nilai     *float64   `json:"nilai"`
	Ttoelf    *string    `json:"ttoelf"`
	Ftoelf    *float64   `json:"ftoelf"`
	Filename  *string    `json:"filename"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
