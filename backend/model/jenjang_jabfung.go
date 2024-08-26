package model

import (
	"time"
)

func (JenjangJabfung) TableName() string {
	return "master_jenjang_jabfung"
}

type SearchJenjangJabfung struct {
	Id   string `json:"id,omitempty"`
	Nama string `json:"nama,omitempty"`
}

type DeleteJenjangJabfung struct {
	Id string `json:"id" validate:"required"`
}

type JenjangJabfung struct {
	Id             string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama           string     `json:"nama"`
	Urut           int        `json:"urut"`
	KoefisienId    int        `json:"koefisien_id"`
	NilaiKoefisien float64    `json:"nilai_koefisien"`
	CreatedBy      string     `gorm:"<-:create" json:"created_by"`
	CreatedAt      *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy      string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt      *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
