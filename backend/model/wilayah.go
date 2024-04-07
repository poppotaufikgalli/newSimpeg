package model

import (
	"fmt"
	"gorm.io/gorm"
)

func (Wilayah) TableName() string {
	return "master_wilayah"
}

type SearchWilayah struct {
	Kwil   string `json:"kwil,omitempty"`
	Nwil   string `json:"nwil,omitempty"`
	Twil   string `json:"twil,omitempty"`
	Kprov  string `json:"kprov,omitempty"`
	Kkab   string `json:"kkab,omitempty"`
	Kkec   string `json:"kkec,omitempty"`
	Tkdesa string `json:"tkdesa,omitempty"`
	Kdesa  string `json:"kdesa,omitempty"`
}

type DeleteWilayah struct {
	Kwil string `json:"kwil" validate:"required"`
}

func (u *Wilayah) AfterFind(tx *gorm.DB) (err error) {
	u.TkKdesa = fmt.Sprintf("%s%s", u.Tkdesa, u.Kdesa)
	return
}

type Wilayah struct {
	Kwil    string `gorm:"primaryKey;autoIncrement:false" json:"kwil" validate:"required"`
	Nwil    string `json:"nwil"`
	Twil    string `json:"twil"`
	Kprov   string `json:"kprov"`
	Kkab    string `json:"kkab"`
	Kkec    string `json:"kkec"`
	Tkdesa  string `json:"tkdesa"`
	Kdesa   string `json:"kdesa"`
	TkKdesa string `json:"tk_kdesa"`
}
