package model

import (
	"gorm.io/gorm"
	"time"
)

func (RiwayatGajiberkala) TableName() string {
	return "master_riwayat_gajiberkala"
}

type SearchRiwayatGajiberkala struct {
	Nip     string    `json:"nip,omitempty"`
	Nstahu  string    `json:"nstahu,omitempty"`
	Tmtngaj time.Time `json:"tmtngaj,omitempty"`
	IdOpd   string    `json:"id_opd,omitempty"`
}

type DeleteRiwayatGajiberkala struct {
	Nip     string    `json:"nip" validate:"required"`
	Tmtngaj time.Time `json:"tmtngaj" validate:"required"`
	IdOpd   string    `json:"id_opd"  validate:"required"`
}

type RiwayatGajiberkala struct {
	Nip       string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Nstahu    *string    `json:"nstahu"`
	Tstahu    *time.Time `json:"tstahu"`
	Tmtngaj   *time.Time `json:"tmtngaj" validate:"required"`
	Gpokkhir  float64    `json:"gpokkhir"`
	Kkantor   *string    `json:"kkantor"`
	Mskerja   *int       `json:"mskerja"`
	Mskerjabl *int       `json:"mskerjabl"`
	Kpej      *int       `json:"kpej"`
	IdOpd     *string    `json:"id_opd"  validate:"required"`
	Nunker    *string    `json:"nunker"`
	Akhir     *int       `json:"akhir"`
	Filename  *string    `json:"filename"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}

func (u *RiwayatGajiberkala) BeforeCreate(tx *gorm.DB) (err error) {
	//u.Akhir = uuid.New()

	if *u.Akhir == 1 {
		//err = errors.New("can't save invalid data")
		tx.Model(&RiwayatGajiberkala{}).Where("nip = ?", u.Nip).Update("akhir", 0)
	}
	return
}

func (u *RiwayatGajiberkala) AfterUpdate(tx *gorm.DB) (err error) {
	if *u.Akhir == 1 {
		Tmtngaj := (u.Tmtngaj).Format("2006-01-02")
		tx.Model(&RiwayatGajiberkala{}).Where("nip = ? and tmtngaj <> ?", u.Nip, Tmtngaj).Update("akhir", 0)
	}
	return
}
