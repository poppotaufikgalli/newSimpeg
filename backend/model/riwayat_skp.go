package model

import (
	"time"
)

func (RiwayatSkp) TableName() string {
	return "master_riwayat_skp"
}

type SearchRiwayatSkp struct {
	Nip             string `json:"nip,omitempty"`
	Tahun           int    `json:"tahun,omitempty"`
	PnsDinilaiOrang string `gorm:"column:pnsDinilaiOrang" json:"pnsDinilaiOrang,omitempty"`
}

type DeleteRiwayatSkp struct {
	Nip   string `json:"nip" validate:"required"`
	Tahun int    `json:"tahun" validate:"required"`
}

type RiwayatSkp struct {
	Nip                      string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Tahun                    int        `json:"tahun" validate:"required"`
	JenisJabatan             string     `gorm:"column:jenisJabatan" json:"jenisJabatan" validate:"required"`
	NilaiSkp                 int        `gorm:"column:nilaiSkp" json:"nilaiSkp" validate:"required"`
	Kerjasama                int        `json:"kerjasama" validate:"required"`
	Integritas               int        `json:"integritas" validate:"required"`
	Komitmen                 int        `json:"komitmen" validate:"required"`
	Disiplin                 int        `json:"disiplin" validate:"required"`
	OrientasiPelayanan       int        `gorm:"column:orientasiPelayanan" json:"orientasiPelayanan" validate:"required"`
	Kepemimpinan             int        `json:"kepemimpinan" validate:"required"`
	Jumlah                   int        `json:"jumlah" validate:"required"`
	Nilairatarata            int        `gorm:"column:nilairatarata" json:"nilairatarata" validate:"required"`
	NilaiPerilakuKerja       int        `gorm:"column:nilaiPerilakuKerja" json:"nilaiPerilakuKerja" validate:"required"`
	NilaiPrestasiKerja       int        `gorm:"column:nilaiPrestasiKerja" json:"nilaiPrestasiKerja" validate:"required"`
	AtasanPejabatPenilai     *string    `gorm:"column:atasanPejabatPenilai" json:"atasanPejabatPenilai"`
	AtasanPenilaiGolongan    *string    `gorm:"column:atasanPenilaiGolongan" json:"atasanPenilaiGolongan"`
	AtasanPenilaiJabatan     *string    `gorm:"column:atasanPenilaiJabatan" json:"atasanPenilaiJabatan"`
	AtasanPenilaiNama        *string    `gorm:"column:atasanPenilaiNama" json:"atasanPenilaiNama"`
	AtasanPenilaiNipNrp      *string    `gorm:"column:atasanPenilaiNipNrp" json:"atasanPenilaiNipNrp"`
	AtasanPenilaiTmtGolongan *time.Time `gorm:"column:atasanPenilaiTmtGolongan" json:"atasanPenilaiTmtGolongan"`
	AtasanPenilaiUnorNama    *string    `gorm:"column:atasanPenilaiUnorNama" json:"atasanPenilaiUnorNama"`
	IdSync                   *string    `gorm:"column:idSync" json:"idSync"`
	PejabatPenilai           *string    `gorm:"column:pejabatPenilai" json:"pejabatPenilai"`
	PenilaiGolongan          *string    `gorm:"column:penilaiGolongan" json:"penilaiGolongan"`
	PenilaiJabatan           *string    `gorm:"column:penilaiJabatan" json:"penilaiJabatan"`
	PenilaiNama              *string    `gorm:"column:penilaiNama" json:"penilaiNama"`
	PenilaiNipNrp            *string    `gorm:"column:penilaiNipNrp" json:"penilaiNipNrp"`
	PenilaiTmtGolongan       *time.Time `gorm:"column:penilaiTmtGolongan" json:"penilaiTmtGolongan"`
	PenilaiUnorNama          *string    `gorm:"column:penilaiUnorNama" json:"penilaiUnorNama"`
	PnsDinilaiOrang          string     `gorm:"column:pnsDinilaiOrang" json:"pnsDinilaiOrang" validate:"required"`
	PnsUserId                *string    `gorm:"column:pnsUserId" json:"pnsUserId"`
	StatusAtasanPenilai      *string    `gorm:"column:statusAtasanPenilai" json:"statusAtasanPenilai"`
	StatusPenilai            *string    `gorm:"column:statusPenilai" json:"statusPenilai"`
	CreatedBy                string     `gorm:"<-:create" json:"created_by"`
	CreatedAt                *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy                string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt                *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
