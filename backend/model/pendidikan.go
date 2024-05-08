package model

import (
	"time"
)

func (Pendidikan) TableName() string {
	return "master_pendidikan"
}

type SearchPendidikan struct {
	Id             string   `json:"id,omitempty"`
	RefBkn         string   `json:"ref_bkn,omitempty"`
	TkPendidikanId []string `json:"tk_pendidikan_id,omitempty"`
	Nama           string   `json:"nama,omitempty"`
	Status         []int    `json:"status,omitempty"`
}

type DeletePendidikan struct {
	Id string `json:"id" validate:"required"`
}

type Pendidikan struct {
	Id             string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	RefBkn         *string    `json:"ref_bkn"`
	TkPendidikanId string     `json:"tk_pendidikan_id"`
	Nama           string     `json:"nama"`
	Status         *int       `json:"status"`
	CreatedBy      string     `gorm:"<-:create" json:"created_by"`
	CreatedAt      *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy      string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt      *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
