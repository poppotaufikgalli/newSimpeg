package model

import ()

func (Stlud) TableName() string {
	return "master_stlud"
}

type SearchStlud struct {
	Kstlud float64 `json:"kstlud,omitempty"`
	Nstlud string  `json:"nstlud,omitempty"`
}

type DeleteStlud struct {
	Kstlud float64 `json:"kstlud" validate:"required"`
}

type Stlud struct {
	Kstlud float64 `gorm:"primaryKey;autoIncrement:false" json:"kstlud"`
	Nstlud *string `json:"nstlud"`
}
