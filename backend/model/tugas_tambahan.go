package model

import (
	"time"
)

func (TugasTambahan) TableName() string {
	return "master_tugas_tambahan"
}

type SearchTugasTambahan struct {
	Id        string `json:"id,omitempty"`
	Nama      string `json:"nama,omitempty"`
	Singkatan string `json:"singkatan"`
}

type DeleteTugasTambahan struct {
	Id string `json:"id" validate:"required"`
}

type TugasTambahan struct {
	Id        string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama      *string    `json:"nama"`
	Singkatan *string    `json:"singkatan"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
