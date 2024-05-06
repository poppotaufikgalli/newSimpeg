package model

import (
	"gorm.io/gorm"
	"time"
)

func (RiwayatPangkat) TableName() string {
	return "master_riwayat_pangkat"
}

type SearchRiwayatPangkat struct {
	Nip     string    `json:"nip,omitempty"`
	Nskpang string    `json:"nskpang,omitempty"`
	Tmtpang time.Time `json:"tmtpang,omitempty"`
	Kgolru  []float64 `json:"kgolru,omitempty"`
	Akhir   []int     `json:"akhir,omitempty"`
}

type DeleteRiwayatPangkat struct {
	Nip     string    `json:"nip" validate:"required"`
	Tmtpang time.Time `json:"tmtpang" validate:"required"`
	Kgolru  float64   `json:"kgolru" validate:"required"`
	Knpang  float64   `json:"knpang" validate:"required"`
}

type RiwayatPangkat struct {
	Nip         string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Kstlud      float64    `json:"kstlud"`
	Nstlud      string     `json:"nstlud"`
	Tstlud      *time.Time `json:"tstlud"`
	Nntbakn     string     `json:"nntbakn"`
	Tntbakn     *time.Time `json:"tntbakn"`
	Akredit     float64    `json:"akredit"`
	Akredit2    float64    `gorm:"column:akredit_2" json:"akredit_2"`
	Ptetap      float64    `json:"ptetap"`
	Nskpang     *string    `json:"nskpang"`
	Tskpang     *time.Time `json:"tskpang"`
	Tmtpang     *time.Time `json:"tmtpang" validate:"required"`
	Kgolru      float64    `json:"kgolru" validate:"required"`
	Pangkat     Pangkat    `gorm:"foreignKey:Kgolru" json:"pangkat" validate:"-"`
	Nsttpp      *string    `json:"nsttpp"`
	Tsttpp      *time.Time `json:"tsttpp"`
	Tmtlgas     *time.Time `json:"tmtlgas"`
	Kjpns       *int       `json:"kjpns"`
	Knpang      float64    `json:"knpang" validate:"required"`
	JenisKp     JenisKp    `gorm:"foreignKey:Knpang" json:"jenis_kp" validate:"-"`
	Mskerja     float64    `json:"mskerja"`
	Blnkerja    float64    `json:"blnkerja"`
	Srtsehatno  *string    `json:"srtsehatno"`
	Srtsehattgl *time.Time `json:"srtsehattgl"`
	Gpok        *float64   `json:"gpok"`
	Akhir       *int       `json:"akhir"`
	Filename    *string    `json:"filename"`
	IdSync      *string    `gorm:"column:idSync" json:"idSync"`
	CreatedBy   string     `gorm:"<-:create" json:"created_by"`
	CreatedAt   *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy   string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt   *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}

func (u *RiwayatPangkat) BeforeCreate(tx *gorm.DB) (err error) {
	//u.Akhir = uuid.New()

	if *u.Akhir == 1 {
		//err = errors.New("can't save invalid data")
		tx.Model(&RiwayatPangkat{}).Where("nip = ?", u.Nip).Update("akhir", 0)
	}
	return
}

func (u *RiwayatPangkat) AfterUpdate(tx *gorm.DB) (err error) {
	if *u.Akhir == 1 {
		Tmtpang := (u.Tmtpang).Format("2006-01-02")
		tx.Model(&RiwayatPangkat{}).Debug().Where("nip = ? and kgolru <> ? and tmtpang <> ? and knpang <> ?", u.Nip, u.Kgolru, Tmtpang, u.Knpang).Update("akhir", 0)
	}
	return
}
