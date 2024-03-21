package model

import (
	"time"
)

func (RiwayatAngkaKredit) TableName() string {
	return "master_riwayat_angka_kredit"
}

type SearchRiwayatAngkaKredit struct {
	Nip    string    `json:"nip,omitempty"`
	Jns    string    `json:"jns,omitempty"`
	Kjab   string    `json:"kjab,omitempty"`
	Tmulai time.Time `json:"tmulai,omitempty"`
	Thn    string    `json:"thn,omitempty"`
}

type DeleteRiwayatAngkaKredit struct {
	Nip    string    `json:"nip" validate:"required"`
	Tmulai time.Time `json:"tmulai" validate:"required"`
}

type RiwayatAngkaKredit struct {
	Nip       string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Jns       string     `json:"jns"`
	Kjab      string     `json:"kjab"`
	Tmulai    *time.Time `json:"tmulai" validate:"required"`
	Tselesai  *time.Time `json:"tselesai"`
	Utama     float64    `json:"utama"`
	Tambahan  *float64   `json:"tambahan"`
	Total     *float64   `json:"total"`
	Nsk       *string    `json:"nsk"`
	Tsk       *time.Time `json:"tsk"`
	Thn       *string    `json:"thn"`
	Filename  *string    `json:"filename"`
	IdSync    *string    `gorm:"column:idSync" json:"idSync"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
