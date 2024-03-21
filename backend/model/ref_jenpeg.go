package model

import ()

func (RefJenpeg) TableName() string {
	return "ref_jenpeg"
}

type SearchRefJenpeg struct {
	Kjpeg float64 `json:"kjpeg,omitempty"`
	Njpeg string  `json:"njpeg,omitempty"`
}

type DeleteRefJenpeg struct {
	Kjpeg float64 `json:"kjpeg" validate:"required"`
}

type RefJenpeg struct {
	Kjpeg float64 `gorm:"primaryKey;autoIncrement:false" json:"kjpeg" validate:"required"`
	Njpeg *string `json:"njpeg"`
}
