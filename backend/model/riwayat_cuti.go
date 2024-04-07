package model

import (
	"time"
)

func (RiwayatCuti) TableName() string {
	return "master_riwayat_cuti"
}

type SearchRiwayatCuti struct {
	Nip    string    `json:"nip,omitempty"`
	Kcuti  int       `json:"kcuti,omitempty"`
	Nsk    string    `json:"nsk,omitempty"`
	Tmulai time.Time `json:"tmulai,omitempty"`
	Thn    string    `json:"thn,omitempty"`
}

type DeleteRiwayatCuti struct {
	Nip    string    `json:"nip" validate:"required"`
	Nsk    string    `json:"nsk" validate:"required"`
	Tmulai time.Time `json:"tmulai" validate:"required"`
}

type RiwayatCuti struct {
	Nip       string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Kcuti     int        `json:"kcuti"`
	JenisCuti JenisCuti  `gorm:"foreignKey:Kcuti" json:"JenisCuti"`
	Jcuti     string     `json:"jcuti"`
	Nsk       string     `json:"nsk" validate:"required"`
	Tsk       *time.Time `json:"tsk"`
	Tmulai    time.Time  `json:"tmulai" validate:"required"`
	Takhir    *time.Time `json:"takhir"`
	Thn       *string    `json:"thn"`
	Ptetap    *int       `json:"ptetap"`
	Jmlhari   *int       `json:"jmlhari"`
	Filename  *string    `json:"filename"`
	IdSync    *string    `gorm:"column:idSync" json:"idSync"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
