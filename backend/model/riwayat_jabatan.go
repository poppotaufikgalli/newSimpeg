package model

import (
	"gorm.io/gorm"
	"time"
)

func (RiwayatJabatan) TableName() string {
	return "master_riwayat_jabatan"
}

type SearchRiwayatJabatan struct {
	ID            string    `json:"id,omitempty"`
	Nip           string    `json:"nip,omitempty"`
	Tmtjab        time.Time `json:"tmtjab,omitempty"`
	Kjab          string    `json:"kjab,omitempty"`
	KjabBkn       string    `json:"kjab_bkn,omitempty"`
	Jnsjab        string    `json:"jnsjab"`
	Keselon       []float64 `json:"keselon,omitempty"`
	Nskjabat      string    `json:"nskjabat,omitempty"`
	IdOpd         string    `json:"id_opd,omitempty"`
	Akhir         []int     `json:"akhir,omitempty"`
	TugasTambahan string    `json:"tugas_tambahan,omitempty"`
	IdOpdTambahan string    `json:"id_opd_tambahan,omitempty"`
	IdSync        string    `json:"idSync,omitempty"`
}

type DeleteRiwayatJabatan struct {
	ID       int    `json:"id" validate:"required"`
	Nip      string `json:"nip" validate:"required"`
	Nskjabat string `json:"nskjabat,omitempty"`
	IdSync   string `gorm:"column:idSync" json:"idSync,omitempty"`
	//Tmtjab    time.Time `json:"tmtjab" validate:"required"`
	//TmtMutasi time.Time `json:"tmt_mutasi" validate:"required"`
}

type RiwayatJabatan struct {
	ID                  uint             `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Nip                 string           `json:"nip" validate:"required"`
	Tmtjab              *time.Time       `json:"tmtjab" validate:"required"`
	Kjab                string           `json:"kjab"`
	KjabBkn             *string          `json:"kjab_bkn"`
	Jnsjab              string           `json:"jnsjab"`
	JnsTugasMutasi      string           `json:"jns_tugas_mutasi"`
	JenisTugasMutasi    JenisTugasMutasi `gorm:"foreignKey:JnsTugasMutasi" validate:"-"`
	TmtMutasi           *time.Time       `json:"tmt_mutasi"`
	Keselon             float64          `json:"keselon"`
	Eselon              Eselon           `gorm:"foreignKey:Keselon" json:"eselon" validate:"-"`
	Njab                *string          `json:"njab"`
	Njab0               *string          `json:"njab0"`
	Sjab                *float64         `json:"sjab"`
	SubJab              *string          `json:"sub_jab"`
	Nunker              string           `json:"nunker"`
	Kpej                float64          `json:"kpej"`
	Nskjabat            string           `json:"nskjabat" validate:"required"`
	Tskjabat            time.Time        `json:"tskjabat" validate:"required"`
	Nlantik             *string          `json:"nlantik"`
	Tlantik             *time.Time       `json:"tlantik"`
	IdInstansi          *string          `json:"id_instansi"`
	NamaInstansi        *string          `json:"nama_instansi"`
	IdOpd               *string          `json:"id_opd" validate:"required"`
	IdSubOpd            *string          `json:"id_sub_opd"`
	IdUnorBkn           *string          `json:"id_unor_bkn"`
	Akhir               *int             `json:"akhir"`
	TugasTambahan       *string          `json:"tugas_tambahan"`
	NjabTambahan        *string          `json:"njab_tambahan"`
	KjabTambahan        *string          `json:"kjab_tambahan"`
	IdOpdTambahan       *string          `json:"id_opd_tambahan"`
	NunkerTambahan      *string          `json:"nunker_tambahan"`
	FilenameLantik      *string          `json:"filename_lantik"`
	FilenameTugasMutasi *string          `json:"filename_tugas_mutasi"`
	FilenameJab         *string          `json:"filename_jab"`
	Filename            *string          `json:"filename"`
	IdSync              *string          `gorm:"column:idSync" json:"idSync"`
	DoBkn               *int             `gorm:"-" json:"do_bkn"`
	CreatedBy           string           `gorm:"<-:create" json:"created_by"`
	CreatedAt           *time.Time       `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy           string           `gorm:"<-:update" json:"updated_by"`
	UpdatedAt           *time.Time       `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}

func (u *RiwayatJabatan) BeforeCreate(tx *gorm.DB) (err error) {
	//u.Akhir = uuid.New()

	if u.Akhir != nil && *u.Akhir == 1 {
		//err = errors.New("can't save invalid data")
		tx.Model(&RiwayatJabatan{}).Where("nip = ?", u.Nip).Update("akhir", 0)
	}
	return
}

func (u *RiwayatJabatan) AfterUpdate(tx *gorm.DB) (err error) {
	if u.Akhir != nil && *u.Akhir == 1 {
		Tmtjab := (u.Tmtjab).Format("2006-01-02")
		tx.Model(&RiwayatJabatan{}).Where("nip = ? and tmtjab <> ?", u.Nip, Tmtjab).Update("akhir", 0)
	}
	return
}

func CheckTmtMutasi(tmt_mutasi string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if tmt_mutasi == "" || tmt_mutasi == "null" {
			return db.Where("tmt_mutasi is null")
		} else {
			return db.Where("tmt_mutasi = ?", tmt_mutasi)
		}
	}
}
