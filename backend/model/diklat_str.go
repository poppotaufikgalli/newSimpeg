package model

import (
	"time"
)

func (DiklatStr) TableName() string {
	return "master_diklat_str"
}

type SearchDiklatStr struct {
	Id   string `json:"id,omitempty"`
	Nama string `json:"nama,omitempty"`
}

type DeleteDiklatStr struct {
	Id string `json:"id" validate:"required"`
}

type DiklatStr struct {
	Id        string     `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Nama      string     `json:"nama"`
	RefSimpeg string     `json:"ref_simpeg"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
