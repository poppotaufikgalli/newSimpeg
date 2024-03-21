package model

import (
	"time"
)

func (TingkatPendidikan) TableName() string {
	return "master_tingkat_pendidikan"
}

type SearchTingkatPendidikan struct {
	Id   string `json:"id,omitempty"`
	Nama string `json:"nama,omitempty"`
}

type DeleteTingkatPendidikan struct {
	Id string `json:"id" validate:"required"`
}

type TingkatPendidikan struct {
	Id            string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama          *string    `json:"nama"`
	GroupTkPendNm *string    `json:"group_tk_pend_nm"`
	Maxkgolru     *string    `json:"maxkgolru"`
	RefSimpeg     *string    `json:"ref_simpeg"`
	CreatedBy     string     `gorm:"<-:create" json:"created_by"`
	CreatedAt     *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy     string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt     *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
