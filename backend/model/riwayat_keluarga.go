package model

import (
	"time"
)

func (RiwayatKeluarga) TableName() string {
	return "master_riwayat_keluarga"
}

type SearchRiwayatKeluarga struct {
	Nip       string `json:"nip,omitempty"`
	Jkeluarga int    `json:"jkeluarga,omitempty"`
	NamaKel   string `json:"nama_kel,omitempty"`
}

type DeleteRiwayatKeluarga struct {
	Nip       string `json:"nip" validate:"required"`
	Jkeluarga int    `json:"jkeluarga" validate:"required"`
	NamaKel   string `json:"nama_kel" validate:"required"`
}

type RiwayatKeluarga struct {
	Nip       string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Jkeluarga int        `json:"jkeluarga" validate:"required"`
	NamaKel   string     `json:"nama_kel" validate:"required"`
	Ktlahir   *string    `json:"ktlahir"`
	Tlahir    *time.Time `json:"tlahir"`
	Tijazah   *string    `json:"tijazah"`
	Tkawin    *time.Time `json:"tkawin"`
	Stunj     string     `json:"stunj"`
	Kjkel     int        `json:"kjkel"`
	Kkerja    float64    `json:"kkerja"`
	Pekerjaan Pekerjaan  `gorm:"foreignKey:Kkerja" json:"Pekerjaan"`
	Instansi  *string    `json:"instansi"`
	NipKel    *string    `json:"nip_kel"`
	Hubkel    string     `json:"hubkel"`
	Akhir     int        `json:"akhir"`
	Aljalan   *string    `json:"aljalan"`
	Alrt      *string    `json:"alrt"`
	Alrw      *string    `json:"alrw"`
	Kodepos   *string    `json:"kodepos"`
	Notelp    *string    `json:"notelp"`
	Wil       *string    `json:"wil"`
	Filename  *string    `json:"filename"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
