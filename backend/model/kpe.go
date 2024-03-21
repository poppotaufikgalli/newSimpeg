package model

import ()

func (Kpe) TableName() string {
	return "master_kpe"
}

type SearchKpe struct {
	Kkpe string `json:"kkpe,omitempty"`
	Nkpe string `json:"nkpe,omitempty"`
}

type DeleteKpe struct {
	Kkpe string `json:"kkpe" validate:"required"`
}

type Kpe struct {
	Kkpe string `gorm:"primaryKey;autoIncrement:false" json:"kkpe" validate:"required"`
	Nkpe string `json:"nkpe" validate:"required"`
}
