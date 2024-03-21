package model

import (
	"time"
)

func (RiwayatSingkronisasi) TableName() string {
	return "master_riwayat_singkronisasi"
}

type SearchRiwayatSingkronisasi struct {
	Id     int    `json:"id,omitempty"`
	Table  string `json:"table,omitempty"`
	Method string `json:"method,omitempty"`
}

type DeleteRiwayatSingkronisasi struct {
	Id int `json:"id" validate:"required"`
}

type RiwayatSingkronisasi struct {
	Id        int        `gorm:"primaryKey;autoIncrement:true" json:"id" validate:"required"`
	Key       *float64   `json:"key"`
	Table     *string    `json:"table"`
	Jml       *string    `json:"jml"`
	Method    *string    `json:"method"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
