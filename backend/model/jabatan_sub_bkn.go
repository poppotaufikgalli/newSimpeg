package model

import (
	"time"
)

func (JabatanSubBkn) TableName() string {
	return "master_jabatan_sub_bkn"
}

type SearchJabatanSubBkn struct {
	Id           string   `json:"id,omitempty"`
	KelJabatanId string   `json:"kel_jabatan_id,omitempty"`
	Nama         string   `json:"nama,omitempty"`
	Status       []string `json:"status,omitempty"`
}

type DeleteJabatanSubBkn struct {
	Id string `json:"id" validate:"required"`
}

type JabatanSubBkn struct {
	Id           string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	KelJabatanId *string    `json:"kel_jabatan_id" validate:"required"`
	Nama         string     `json:"nama"`
	Status       string     `json:"status"`
	CreatedBy    string     `gorm:"<-:create" json:"created_by"`
	CreatedAt    *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy    string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt    *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
