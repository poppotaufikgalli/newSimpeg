package model

import ()

func (StatusPegawai) TableName() string {
	return "master_status_pegawai"
}

type DeleteStatusPegawai struct {
	IdStatusPegawai int `json:"id_status_pegawai" validate:"required"`
}

type StatusPegawai struct {
	IdStatusPegawai int     `gorm:"primaryKey;autoIncrement:false" json:"id_status_pegawai" validate:"required"`
	Kstatus         string  `json:"kstatus"`
	Nama            string  `json:"nama" validate:"required"`
	Persentase      float64 `json:"persentase"`
}
