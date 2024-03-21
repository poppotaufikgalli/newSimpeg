package model

import (
	"time"
)

func (Cpns) TableName() string {
	return "master_cpns"
}

type SearchCpns struct {
	Nip     string    `json:"nip,omitempty"`
	Skcpns  string    `json:"skcpns,omitempty"`
	Tmtcpns time.Time `json:"tmtcpns,omitempty"`
}

type DeleteCpns struct {
	Nip string `json:"nip" validate:"required"`
}

type Cpns struct {
	Nip         string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Ntbakn      *string    `json:"ntbakn"`
	Tntbakn     *time.Time `json:"tntbakn"`
	Kpej        *float64   `json:"kpej"`
	Skcpns      string     `json:"skcpns"`
	Tskcpns     time.Time  `json:"tskcpns"`
	Tmtcpns     time.Time  `json:"tmtcpns"`
	Kgolru      float64    `json:"kgolru"`
	Nsttpp      *string    `json:"nsttpp"`
	Tsttpp      *time.Time `json:"tsttpp"`
	Tmtlgas     *time.Time `json:"tmtlgas"`
	Mskerjabl   *int       `json:"mskerjabl"`
	Mskerjath   *int       `json:"mskerjath"`
	Srtsehatno  *string    `json:"srtsehatno"`
	Srtsehattgl *time.Time `json:"srtsehattgl"`
	Filename    *string    `json:"filename"`
	CreatedBy   string     `gorm:"<-:create" json:"created_by"`
	CreatedAt   *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy   string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt   *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
