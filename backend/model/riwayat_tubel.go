package model

import (
	"time"
)

func (RiwayatTubel) TableName() string {
	return "master_riwayat_tubel"
}

type SearchRiwayatTubel struct {
	ID     int    `json:"id,omitempty"`
	Nip    string `json:"nip,omitempty"`
	Ktpu   string `json:"ktpu,omitempty"`
	Nsek   string `json:"nsek,omitempty"`
	Tempat string `json:"tempat,omitempty"`
}

type DeleteRiwayatTubel struct {
	ID   int    `json:"id" validate:"required"`
	Nip  string `json:"nip" validate:"required"`
	Ktpu string `json:"ktpu" validate:"required"`
}

type RiwayatTubel struct {
	ID                int               `gorm:"primaryKey" json:"id"`
	Nip               string            `json:"nip" validate:"required"`
	Ktpu              string            `json:"ktpu" validate:"required"`
	TingkatPendidikan TingkatPendidikan `gorm:"foreignKey:Ktpu" json:"TingkatPendidikan" validate:"-"`
	Njur              string            `json:"njur"`
	Nsek              *string           `json:"nsek"`
	Tempat            *string           `json:"tempat"`
	Aktif             *int              `json:"aktif"`
	Tahun             *int              `json:"tahun"`
	Filename          *string           `json:"filename"`
	Nsk               string            `json:"nsk"`
	Tsk               *time.Time        `json:"tsk"`
	IdSync            *string           `gorm:"column:idSync" json:"idSync"`
	CreatedBy         string            `gorm:"<-:create" json:"created_by"`
	CreatedAt         *time.Time        `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy         string            `gorm:"<-:update" json:"updated_by"`
	UpdatedAt         *time.Time        `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
