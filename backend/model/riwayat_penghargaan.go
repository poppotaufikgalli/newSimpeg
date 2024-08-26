package model

import (
	"time"
)

func (RiwayatPenghargaan) TableName() string {
	return "master_riwayat_penghargaan"
}

type SearchRiwayatPenghargaan struct {
	Nip                string  `json:"nip,omitempty"`
	IdJenisPenghargaan string  `json:"id_jenis_penghargaan,omitempty"`
	Nsk                string  `json:"nsk,omitempty"`
	Thn                float64 `json:"thn,omitempty"`
}

type DeleteRiwayatPenghargaan struct {
	Nip      string `json:"nip" validate:"required"`
	Nbintang string `json:"nbintang" validate:"required"`
	Nsk      string `json:"nsk" validate:"required"`
	IdSync   string `json:"idSync"`
}

type RiwayatPenghargaan struct {
	Nip                string           `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	IdJenisPenghargaan *string          `json:"id_jenis_penghargaan"`
	JenisPenghargaan   JenisPenghargaan `gorm:"foreignKey:IdJenisPenghargaan" json:"JenisPenghargaan" validate:"-"`
	Nbintang           string           `json:"nbintang" validate:"required"`
	Aoleh              *string          `json:"aoleh"`
	Nsk                *string          `json:"nsk" validate:"required"`
	Tsk                *time.Time       `json:"tsk"`
	Thn                *float64         `json:"thn"`
	Filename           *string          `json:"filename"`
	IdSync             *string          `gorm:"column:idSync" json:"idSync"`
	CreatedBy          string           `gorm:"<-:create" json:"created_by"`
	CreatedAt          *time.Time       `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy          string           `gorm:"<-:update" json:"updated_by"`
	UpdatedAt          *time.Time       `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
