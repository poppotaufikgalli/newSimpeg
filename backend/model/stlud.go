package model

import (
	"time"
)

func (Stlud) TableName() string {
	return "master_stlud"
}

type SearchStlud struct {
	Kstlud int    `json:"kstlud,omitempty"`
	Nstlud string `json:"nstlud,omitempty"`
}

type DeleteStlud struct {
	Kstlud int `json:"kstlud" validate:"required"`
}

type Stlud struct {
	Kstlud    int        `gorm:"primaryKey;autoIncrement:false" json:"kstlud" validate:"required"`
	Nstlud    *string    `json:"nstlud"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
