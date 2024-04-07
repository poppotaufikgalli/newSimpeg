package model

import (
	"time"
)

func (RiwayatStlud) TableName() string {
	return "master_riwayat_stlud"
}

type SearchRiwayatStlud struct {
	Nip           string    `json:"nip,omitempty"`
	Kstlud        int       `json:"kstlud,omitempty"`
	Tstlud        time.Time `json:"tstlud,omitempty"`
	Penyelenggara string    `json:"penyelenggara,omitempty"`
}

type DeleteRiwayatStlud struct {
	Nip    string    `json:"nip" validate:"required"`
	Kstlud int       `json:"kstlud" validate:"required"`
	Tstlud time.Time `json:"tstlud" validate:"required"`
}

type RiwayatStlud struct {
	Nip           string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Kstlud        int        `json:"kstlud" validate:"required"`
	Stlud         Stlud      `gorm:"foreignKey:Kstlud;references:kstlud" json:"stlud"`
	Nostlud       *string    `json:"nostlud"`
	Tstlud        time.Time  `json:"tstlud" validate:"required"`
	Nilai         *float64   `json:"nilai"`
	Penyelenggara *string    `json:"penyelenggara"`
	Tempat        *string    `json:"tempat"`
	Filename      *string    `json:"filename"`
	CreatedBy     string     `gorm:"<-:create" json:"created_by"`
	CreatedAt     *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy     string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt     *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
