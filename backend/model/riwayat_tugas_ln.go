package model

import (
	"time"
)

func (RiwayatTugasLn) TableName() string {
	return "master_riwayat_tugas_ln"
}

type SearchRiwayatTugasLn struct {
	Nip    string    `json:"nip,omitempty"`
	Nneg   string    `json:"nneg,omitempty"`
	Tujuan string    `json:"tujuan,omitempty"`
	Tmulai time.Time `json:"tmulai,omitempty"`
}

type DeleteRiwayatTugasLn struct {
	Nip    string    `json:"nip" validate:"required"`
	Tmulai time.Time `json:"tmulai" validate:"required"`
}

type RiwayatTugasLn struct {
	Nip       string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Nneg      *string    `json:"nneg"`
	Tujuan    *string    `json:"tujuan"`
	Nsk       *string    `json:"nsk" `
	Tsk       *time.Time `json:"tsk"`
	Ptetap    *int       `json:"ptetap"`
	Tmulai    time.Time  `json:"tmulai" validate:"required"`
	Tselesai  *time.Time `json:"tselesai"`
	Filename  *string    `json:"filename"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
