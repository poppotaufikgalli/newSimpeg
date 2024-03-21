package model

import (
	"time"
)

func (JenisHukdis) TableName() string {
	return "master_jenis_hukdis"
}

type SearchJenisHukdis struct {
	Id                    string   `json:"id,omitempty"`
	Nama                  string   `json:"nama,omitempty"`
	JenisTingkatHukumanId []string `json:"jenis_tingkat_hukuman_id,omitempty"`
	IdRefBkn              []string `json:"id_ref_bkn,omitempty"`
}

type DeleteJenisHukdis struct {
	Id string `json:"id" validate:"required"`
}

type JenisHukdis struct {
	Id                    string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama                  string     `json:"nama"`
	JenisTingkatHukumanId string     `json:"jenis_tingkat_hukuman_id"`
	IdRefBkn              string     `json:"id_ref_bkn"`
	CreatedBy             string     `gorm:"<-:create" json:"created_by"`
	CreatedAt             *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy             string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt             *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
