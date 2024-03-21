package model

import (
	"time"
)

func (RiwayatJabatan) TableName() string {
	return "master_riwayat_jabatan"
}

type SearchRiwayatJabatan struct {
	Nip           string    `json:"nip,omitempty"`
	Tmtjab        time.Time `json:"tmtjab,omitempty"`
	Kjab          string    `json:"kjab,omitempty"`
	KjabBkn       string    `json:"kjab_bkn,omitempty"`
	Keselon       []float64 `json:"keselon,omitempty"`
	Nskjabat      string    `json:"nskjabat,omitempty"`
	IdOpd         string    `json:"id_opd,omitempty"`
	Akhir         []int     `json:"akhir,omitempty"`
	TugasTambahan string    `json:"tugas_tambahan,omitempty"`
	IdOpdTambahan string    `json:"id_opd_tambahan,omitempty"`
}

type DeleteRiwayatJabatan struct {
	Nip    string    `json:"nip" validate:"required"`
	Tmtjab time.Time `json:"tmtjab" validate:"required"`
}

type RiwayatJabatan struct {
	Nip            string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Tmtjab         *time.Time `json:"tmtjab" validate:"required"`
	Kjab           string     `json:"kjab"`
	KjabBkn        *string    `json:"kjab_bkn"`
	Jnsjab         string     `json:"jnsjab"`
	Keselon        float64    `json:"keselon"`
	Njab           *string    `json:"njab"`
	Sjab           *float64   `json:"sjab"`
	Nunker         string     `json:"nunker"`
	Kpej           float64    `json:"kpej"`
	Nskjabat       string     `json:"nskjabat" validate:"required"`
	Tskjabat       time.Time  `json:"tskjabat" validate:"required"`
	Nlantik        *string    `json:"nlantik"`
	Tlantik        *time.Time `json:"tlantik"`
	IdInstansi     *string    `json:"id_instansi"`
	NamaInstansi   *string    `json:"nama_instansi"`
	IdOpd          *string    `json:"id_opd" validate:"required"`
	IdSubOpd       *string    `json:"id_sub_opd"`
	Akhir          int        `json:"akhir"`
	TugasTambahan  *string    `json:"tugas_tambahan"`
	NjabTambahan   *string    `json:"njab_tambahan"`
	KjabTambahan   *string    `json:"kjab_tambahan"`
	IdOpdTambahan  *string    `json:"id_opd_tambahan"`
	NunkerTambahan *string    `json:"nunker_tambahan"`
	FilenamePp     *string    `json:"filename_pp"`
	FilenameBa     *string    `json:"filename_ba"`
	FilenameLamp   *string    `json:"filename_lamp"`
	Filename       *string    `json:"filename"`
	IdSync         *string    `gorm:"column:idSync" json:"idSync"`
	CreatedBy      string     `gorm:"<-:create" json:"created_by"`
	CreatedAt      *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy      string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt      *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
