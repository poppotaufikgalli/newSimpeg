package model

import (
	"time"
)

func (JenisKp) TableName() string {
	return "master_jenis_kp"
}

type SearchJenisKp struct {
	Id        string `json:"id,omitempty"`
	Nama      string `json:"nama,omitempty"`
	RefSimpeg string `json:"ref_simpeg,omitempty"`
	Status    []int  `json:"status,omitempty"`
}

type DeleteJenisKp struct {
	Id string `json:"id" validate:"required"`
}

type JenisKp struct {
	Id        string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama      string     `json:"nama"`
	RefSimpeg *string    `json:"ref_simpeg"`
	Status    *int       `json:"status"`
	Noformat  *string    `json:"noformat"`
	Alasan    *string    `json:"alasan"`
	Noformat2 *string    `json:"noformat2"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
