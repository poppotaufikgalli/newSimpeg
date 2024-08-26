package model

import (
	"time"
)

func (RiwayatKinerja) TableName() string {
	return "master_riwayat_kinerja"
}

type SearchRiwayatKinerja struct {
	ID       int       `json:"id,omitempty"`
	Nip      string    `json:"nip,omitempty"`
	Thnilai  float64   `json:"thnilai,omitempty"`
	Jns      string    `json:"jns,omitempty"`
	Tmulai   time.Time `json:"tmulai,omitempty"`
	Tselesai time.Time `json:"tselesai,omitempty"`
}

type DeleteRiwayatKinerja struct {
	ID      int     `json:"id" validate:"required"`
	Nip     string  `json:"nip" validate:"required"`
	Thnilai float64 `json:"thnilai" validate:"required"`
	IdSync  string  `json:"idSync,omitempty"`
	Jns     string  `json:"jns" validate:"required"`
}

type RiwayatKinerja struct {
	ID             int        `gorm:"primaryKey" json:"id" validate:"-"`
	Nip            string     `json:"nip" validate:"required"`
	Thnilai        float64    `json:"thnilai" json:"tahun" validate:"required"`
	Jns            *string    `json:"jns"`
	Tmulai         *time.Time `json:"tmulai"`
	Tselesai       *time.Time `json:"tselesai"`
	HasilKinerja   *int       `json:"hasil_kinerja" json:"hasilKinerjaNilai"`
	hasilKinerja   *string    `gorm:"column:hasilKinerja" json:"hasilKinerja"`
	PerilakuKerja  *int       `json:"perilaku_kerja" json:"PerilakuKerjaNilai"`
	perilakuKerja  *string    `gorm:"column:perilakuKerja" json:"perilakuKerja"`
	KuadranKinerja *int       `json:"kuadran_kinerja" json:"KuadranKinerjaNilai"`
	kuadranKinerja *string    `gorm:"column:kuadranKinerja" json:"kuadranKinerja"`
	PeriodikId     *int       `json:"periodik_id"`
	KoefisienId    *string    `json:"koefisien_id"`
	Jenjang        *string    `json:"jenjang"`
	NilaiKoefisien *float32   `json:"nilai_koefisien"`
	PejPns         *string    `json:"pej_pns" json:"statusPenilai"`
	PejNikNip      *string    `json:"pej_nik_nip" json:"nipNrpPenilai"`
	PejNama        *string    `json:"pej_nama" json:"namaPenilai"`
	PejKgolru      *int       `json:"pej_kgolru" json:"penilaiGolonganId"`
	PejNunker      *string    `json:"pej_nunker" json:"penilaiUnorNm"`
	PejNjab        *string    `json:"pej_njab" json:"penilaiJabatanNm"`
	IdSync         *string    `gorm:"column:idSync" json:"idSync"`
	Filename       *string    `json:"filename"`
	CreatedBy      string     `gorm:"<-:create" json:"created_by"`
	CreatedAt      *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy      string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt      *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
