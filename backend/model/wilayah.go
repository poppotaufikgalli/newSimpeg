package model

import ()

func (Wilayah) TableName() string {
	return "master_wilayah"
}

type SearchWilayah struct {
	Kwil string `json:"kwil,omitempty"`
	Nwil string `json:"nwil,omitempty"`
}

type DeleteWilayah struct {
	Kwil string `json:"kwil" validate:"required"`
}

type Wilayah struct {
	Kwil   string  `gorm:"primaryKey;autoIncrement:false" json:"kwil" validate:"required"`
	Nwil   *string `json:"nwil"`
	Twil   *string `json:"Twil"`
	Kprov  *string `json:"kprov"`
	Kkab   *string `json:"kkab"`
	Kkec   *string `json:"kkec"`
	Tkdesa *string `json:"tkdesa"`
	Kdesa  *string `json:"kdesa"`
}
