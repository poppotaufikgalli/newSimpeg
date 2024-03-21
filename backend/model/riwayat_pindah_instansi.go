package model

import (
	"time"
)

func (RiwayatPindahInstansi) TableName() string {
	return "master_riwayat_pindah_instansi"
}

type SearchRiwayatPindahInstansi struct {
	Nip               string    `json:"nip,omitempty"`
	InstansiIndukBaru string    `gorm:"column:instansiIndukBaru" json:"instansiIndukBaru,omitempty"`
	InstansiIndukLama string    `gorm:"column:instansiIndukLama" json:"instansiIndukLama,omitempty"`
	JenisPi           string    `gorm:"column:jenisPi" json:"jenisPi,omitempty"`
	TmtPi             time.Time `gorm:"column:tmtPi" json:"tmtPi,omitempty"`
}

type DeleteRiwayatPindahInstansi struct {
	Nip   string    `json:"nip" validate:"required"`
	TmtPi time.Time `gorm:"column:tmtPi" json:"tmtPi" validate:"required"`
}

type RiwayatPindahInstansi struct {
	Nip                       string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Eselon                    *string    `json:"eselon"`
	InstansiIndukBaru         *string    `gorm:"column:instansiIndukBaru" json:"instansiIndukBaru"`
	InstansiIndukLama         *string    `gorm:"column:instansiIndukLama" json:"instansiIndukLama"`
	InstansiKerjaBaru         *string    `gorm:"column:instansiKerjaBaru" json:"instansiKerjaBaru"`
	InstansiKerjaLama         *string    `gorm:"column:instansiKerjaLama" json:"instansiKerjaLama"`
	JabatanFungsionalBaru     *string    `gorm:"column:jabatanFungsionalBaru" json:"jabatanFungsionalBaru"`
	JabatanFungsionalLama     *string    `gorm:"column:jabatanFungsionalLama" json:"jabatanFungsionalLama"`
	JabatanFungsionalUmumBaru *string    `gorm:"column:jabatanFungsionalUmumBaru" json:"jabatanFungsionalUmumBaru"`
	JenisJabatanBaru          *string    `gorm:"column:jenisJabatanBaru" json:"jenisJabatanBaru"`
	JenisJabatanLama          *string    `gorm:"column:jenisJabatanLama" json:"jenisJabatanLama"`
	JenisPegawai              *string    `gorm:"column:jenisPegawai" json:"jenisPegawai"`
	JenisPi                   *string    `gorm:"column:jenisPi" json:"jenisPi"`
	KpknBaru                  *string    `gorm:"column:kpknBaru" json:"kpknBaru"`
	LokasiKerjaBaru           *string    `gorm:"column:lokasiKerjaBaru" json:"lokasiKerjaBaru"`
	LokasiKerjaLama           *string    `gorm:"column:lokasiKerjaLama" json:"lokasiKerjaLama"`
	OrangUsulPeremajaanId     *string    `gorm:"column:orangUsulPeremajaanId" json:"orangUsulPeremajaanId"`
	PnsOrang                  *string    `gorm:"column:pnsOrang" json:"pnsOrang"`
	SatuanKerjaBaru           *string    `gorm:"column:satuanKerjaBaru" json:"satuanKerjaBaru"`
	SatuanKerjaIndukBaru      *string    `gorm:"column:satuanKerjaIndukBaru" json:"satuanKerjaIndukBaru"`
	SatuanKerjaIndukLama      *string    `gorm:"column:satuanKerjaIndukLama" json:"satuanKerjaIndukLama"`
	SatuanKerjaLama           *string    `gorm:"column:satuanKerjaLama" json:"satuanKerjaLama"`
	SkAsalNomor               *string    `gorm:"column:skAsalNomor" json:"skAsalNomor"`
	SkAsalProvNomor           *string    `gorm:"column:skAsalProvNomor" json:"skAsalProvNomor"`
	SkAsalProvTanggal         *string    `gorm:"column:skAsalProvTanggal" json:"skAsalProvTanggal"`
	SkAsalTanggal             *string    `gorm:"column:skAsalTanggal" json:"skAsalTanggal"`
	SkBknNomor                *string    `gorm:"column:skBknNomor" json:"skBknNomor"`
	SkBknTanggal              *string    `gorm:"column:skBknTanggal" json:"skBknTanggal"`
	SkTujuanNomor             *string    `gorm:"column:skTujuanNomor" json:"skTujuanNomor"`
	SkTujuanTanggal           *string    `gorm:"column:skTujuanTanggal" json:"skTujuanTanggal"`
	SkUsulNomor               *string    `gorm:"column:skUsulNomor" json:"skUsulNomor"`
	SkUsulTanggal             *string    `gorm:"column:skUsulTanggal" json:"skUsulTanggal"`
	TmtPi                     time.Time  `gorm:"column:tmtPi" json:"tmtPi" validate:"required"`
	UnorBaru                  *string    `gorm:"column:unorBaru"json:"unorBaru"`
	UnorLama                  *string    `gorm:"column:unorLama"json:"unorLama"`
	Filename                  *string    `json:"filename"`
	IdSync                    *string    `gorm:"column:idSync" json:"idSync"`
	CreatedBy                 string     `gorm:"<-:create" json:"created_by"`
	CreatedAt                 *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy                 string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt                 *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
