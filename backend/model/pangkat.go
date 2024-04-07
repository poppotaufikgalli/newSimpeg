package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

func (Pangkat) TableName() string {
	return "master_pangkat"
}

type SearchPangkat struct {
	Id          string    `json:"id,omitempty"`
	Nama        string    `json:"nama,omitempty"`
	NamaPangkat string    `json:"nama_pangkat,omitempty"`
	RefSimpeg   string    `json:"ref_simpeg,omitempty"`
	Pajak       []float64 `json:"pajak,omitempty"`
}

type DeletePangkat struct {
	Id string `json:"id" validate:"required"`
}

func (u *Pangkat) AfterFind(tx *gorm.DB) (err error) {
	a := u.NamaPangkat
	u.PangkatGol = fmt.Sprintf("%s (%s)", *a, u.Nama)
	return
}

type Pangkat struct {
	Id          string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama        string     `json:"nama"`
	NamaPangkat *string    `json:"nama_pangkat"`
	PangkatGol  string     `gorm:"-" json:"pangkat_gol"`
	RefSimpeg   *string    `json:"ref_simpeg"`
	Pajak       *float64   `json:"pajak"`
	CreatedBy   string     `gorm:"<-:create" json:"created_by"`
	CreatedAt   *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy   string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt   *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
