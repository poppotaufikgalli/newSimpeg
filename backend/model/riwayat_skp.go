package model

import (
	"time"
)

func (RiwayatSkp) TableName() string {
	return "master_riwayat_skp"
}

type SearchRiwayatSkp struct {
	ID              int    `json:"id,omitempty"`
	Nip             string `json:"nip,omitempty"`
	Tahun           int    `json:"tahun,omitempty"`
	PnsDinilaiOrang string `gorm:"column:pnsDinilaiOrang" json:"pnsDinilaiOrang,omitempty"`
}

type DeleteRiwayatSkp struct {
	ID    int    `json:"id" validate:"required"`
	Nip   string `json:"nip" validate:"required"`
	Tahun int    `json:"tahun" validate:"required"`
}

type RiwayatSkp struct {
	ID                       int        `gorm:"primaryKey" json:"id" validate:"-"`
	Nip                      string     `json:"nip" validate:"required"`
	Tahun                    int        `json:"tahun" validate:"required"`
	JenisJabatan             string     `gorm:"column:jenisJabatan" json:"jenisJabatan" validate:"required"`
	NilaiSkp                 float64    `gorm:"column:nilaiSkp" json:"nilaiSkp" validate:"required"`
	Kerjasama                float64    `json:"kerjasama"`
	Integritas               float64    `json:"integritas"`
	Komitmen                 float64    `json:"komitmen"`
	Disiplin                 float64    `json:"disiplin"`
	OrientasiPelayanan       float64    `gorm:"column:orientasiPelayanan" json:"orientasiPelayanan"`
	Kepemimpinan             float64    `json:"kepemimpinan"`
	Jumlah                   float64    `json:"jumlah" validate:"required"`
	Nilairatarata            float64    `gorm:"column:nilairatarata" json:"nilairatarata"`
	NilaiPerilakuKerja       float64    `gorm:"column:nilaiPerilakuKerja" json:"nilaiPerilakuKerja"`
	NilaiPrestasiKerja       float64    `gorm:"column:nilaiPrestasiKerja" json:"nilaiPrestasiKerja"`
	InisiatifKerja           float64    `gorm:"column:inisiatifKerja" json:"inisiatifKerja"`
	KonversiNilai            float64    `gorm:"column:konversiNilai" json:"konversiNilai"`
	NilaiIntegrasi           float64    `gorm:"column:nilaiIntegrasi" json:"nilaiIntegrasi"`
	AtasanPejabatPenilai     *string    `gorm:"column:atasanPejabatPenilai" json:"atasanPejabatPenilai"`
	AtasanPenilaiGolongan    *string    `gorm:"column:atasanPenilaiGolongan" json:"atasanPenilaiGolongan"`
	AtasanPenilaiJabatan     *string    `gorm:"column:atasanPenilaiJabatan" json:"atasanPenilaiJabatan"`
	AtasanPenilaiNama        *string    `gorm:"column:atasanPenilaiNama" json:"atasanPenilaiNama"`
	AtasanPenilaiNipNrp      *string    `gorm:"column:atasanPenilaiNipNrp" json:"atasanPenilaiNipNrp"`
	AtasanPenilaiTmtGolongan *time.Time `gorm:"column:atasanPenilaiTmtGolongan" json:"atasanPenilaiTmtGolongan"`
	AtasanPenilaiUnorNama    *string    `gorm:"column:atasanPenilaiUnorNama" json:"atasanPenilaiUnorNama"`
	IdSync                   *string    `gorm:"column:idSync" json:"idSync"`
	Filename                 *string    `json:"filename"`
	PejabatPenilai           *string    `gorm:"column:pejabatPenilai" json:"pejabatPenilai"`
	PenilaiGolongan          *string    `gorm:"column:penilaiGolongan" json:"penilaiGolongan"`
	PenilaiJabatan           *string    `gorm:"column:penilaiJabatan" json:"penilaiJabatan"`
	PenilaiNama              *string    `gorm:"column:penilaiNama" json:"penilaiNama"`
	PenilaiNipNrp            *string    `gorm:"column:penilaiNipNrp" json:"penilaiNipNrp"`
	PenilaiTmtGolongan       *time.Time `gorm:"column:penilaiTmtGolongan" json:"penilaiTmtGolongan"`
	PenilaiUnorNama          *string    `gorm:"column:penilaiUnorNama" json:"penilaiUnorNama"`
	PnsDinilaiOrang          string     `gorm:"column:pnsDinilaiOrang" json:"pnsDinilaiOrang"`
	PnsUserId                *string    `gorm:"column:pnsUserId" json:"pnsUserId"`
	StatusAtasanPenilai      *string    `gorm:"column:statusAtasanPenilai" json:"statusAtasanPenilai"`
	StatusPenilai            *string    `gorm:"column:statusPenilai" json:"statusPenilai"`
	JenisPeraturanKinerjaKd  *string    `gorm:"column:jenisPeraturanKinerjaKd" json:"jenisPeraturanKinerjaKd"`
	CreatedBy                string     `gorm:"<-:create" json:"created_by"`
	CreatedAt                *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy                string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt                *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
