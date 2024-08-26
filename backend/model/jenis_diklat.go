package model

import (
	"time"
)

func (JenisDiklat) TableName() string {
	return "master_jenis_diklat"
}

type SearchJenisDiklat struct {
	Id     int    `json:"id,omitempty"`
	Nama   string `json:"nama,omitempty"`
	Status []int  `json:"status,omitempty"`
	RefBkn string `json:"ref_bkn,omitempty"`
}

type DeleteJenisDiklat struct {
	Id string `json:"id" validate:"required"`
}

type JenisDiklat struct {
	Id        int        `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama      string     `json:"nama"`
	RefBkn    *string    `json:"ref_bkn"`
	Status    *int       `json:"status"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
