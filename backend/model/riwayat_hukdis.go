package model

import (
	"time"
)

func (RiwayatHukdis) TableName() string {
	return "master_riwayat_hukdis"
}

type SearchRiwayatHukdis struct {
	Nip       string    `json:"nip,omitempty"`
	Jhukum    string    `json:"jhukum,omitempty"`
	JhukumBkn string    `json:"jhukum_bkn,omitempty"`
	Tmt       time.Time `json:"tmt,omitempty"`
}

type DeleteRiwayatHukdis struct {
	Nip    string    `json:"nip" validate:"required"`
	Jhukum string    `json:"jhukum" validate:"required"`
	Tmt    time.Time `json:"tmt" validate:"required"`
}

type RiwayatHukdis struct {
	Nip         string      `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Jhukum      string      `json:"jhukum" validate:"required"`
	JenisHukdis JenisHukdis `gorm:"foreignKey:Jhukum" json:"JenisHukdis" validate:"-"`
	JhukumBkn   *string     `json:"jhukum_bkn"`
	Deshukum    *string     `json:"deshukum"`
	Nsk         *string     `json:"nsk"`
	Tsk         *time.Time  `json:"tsk"`
	Kpej        float64     `json:"kpej"`
	Tmt         *time.Time  `json:"tmt" validate:"required"`
	Selesai     *time.Time  `json:"selesai"`
	Filename    *string     `json:"filename"`
	IdSync      *string     `gorm:"column:idSync" json:"idSync"`
	CreatedBy   string      `gorm:"<-:create" json:"created_by"`
	CreatedAt   *time.Time  `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy   string      `gorm:"<-:update" json:"updated_by"`
	UpdatedAt   *time.Time  `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
