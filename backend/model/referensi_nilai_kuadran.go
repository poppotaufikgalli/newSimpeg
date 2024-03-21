package model

import ()

func (ReferensiNilaiKuadran) TableName() string {
	return "master_referensi_nilai_kuadran"
}

type SearchReferensiNilaiKuadran struct {
	Id   string `json:"id,omitempty"`
	Jns  int    `json:"jns,omitempty"`
	Nama string `json:"nama,omitempty"`
}

type DeleteReferensiNilaiKuadran struct {
	Id  string `json:"id" validate:"required"`
	Jns int    `json:"jns" validate:"required"`
}

type ReferensiNilaiKuadran struct {
	Id   string `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Jns  int    `json:"jns" validate:"required"`
	Nama string `json:"nama" validate:"required"`
}
