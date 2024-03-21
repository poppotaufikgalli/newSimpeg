package model

import (
	"time"
)

func (Pejabat) TableName() string {
	return "master_pejabat"
}

type SearchPejabat struct {
	Kpej string `json:"kpej,omitempty"`
	Npej string `json:"npej,omitempty"`
}

type DeletePejabat struct {
	Kpej string `json:"kpej" validate:"required"`
}

type Pejabat struct {
	Kpej      string     `gorm:"primaryKey;autoIncrement:false" json:"kpej" validate:"required"`
	Npej      string     `json:"npej"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
