package model

import (
	"time"
)

func (RiwayatP2kp) TableName() string {
	return "master_riwayat_p2kp"
}

type SearchRiwayatP2kp struct {
	Nip      string    `json:"nip,omitempty"`
	Thnilai  float64   `json:"thnilai,omitempty"`
	Tmulai   time.Time `json:"tmulai,omitempty"`
	Tselesai time.Time `json:"tselesai,omitempty"`
}

type DeleteRiwayatP2kp struct {
	Nip      string    `json:"nip" validate:"required"`
	Thnilai  float64   `json:"thnilai" validate:"required"`
	Tmulai   time.Time `json:"tmulai" validate:"required"`
	Tselesai time.Time `json:"tselesai" validate:"required"`
}

type RiwayatP2kp struct {
	Nip            string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Thnilai        float64    `json:"thnilai" validate:"required"`
	Tmulai         time.Time  `json:"tmulai" validate:"required"`
	Tselesai       time.Time  `json:"tselesai" validate:"required"`
	HasilKinerja   *int       `json:"hasil_kinerja"`
	PerilakuKerja  *int       `json:"perilaku_kerja"`
	KuadranKinerja *int       `json:"kuadran_kinerja"`
	PejPns         *int       `json:"pej_pns"`
	PejNikNip      *string    `json:"pej_nik_nip"`
	PejNama        *string    `json:"pej_nama"`
	PejKgolru      *int       `json:"pej_kgolru"`
	PejNunker      *string    `json:"pej_nunker"`
	PejNjab        *string    `json:"pej_njab"`
	Filename       *string    `json:"filename"`
	CreatedBy      string     `gorm:"<-:create" json:"created_by"`
	CreatedAt      *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy      string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt      *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
