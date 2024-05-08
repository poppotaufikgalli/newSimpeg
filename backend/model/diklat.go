package model

import (
	"time"
)

func (Diklat) TableName() string {
	return "master_diklat"
}

type SearchDiklat struct {
	Id      string `json:"id,omitempty"`
	Nama    string `json:"nama,omitempty"`
	Jdiklat []int  `json:"jdiklat,omitempty"`
	Status  []int  `json:"status,omitempty"`
}

type DeleteDiklat struct {
	Id      string `json:"id" validate:"required"`
	Jdiklat int    `json:"jdiklat" validate:"required"`
}

type Diklat struct {
	Id        string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama      string     `json:"nama"`
	Jdiklat   int        `json:"jdiklat" validate:"required"`
	KetGroup  *string    `json:"ket_group"`
	Status    *int       `json:"status"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
