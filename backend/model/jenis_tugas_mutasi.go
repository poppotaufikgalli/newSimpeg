package model

import (
	"time"
)

func (JenisTugasMutasi) TableName() string {
	return "master_jenis_tugas_mutasi"
}

type SearchJenisTugasMutasi struct {
	Id   string `json:"id,omitempty"`
	Nama string `json:"nama,omitempty"`
}

type DeleteJenisTugasMutasi struct {
	Id string `json:"id" validate:"required"`
}

type JenisTugasMutasi struct {
	Id        string     `gorm:"primaryKey;autoIncrement:false" json:"id"`
	Nama      string     `json:"nama"`
	Jns       *string    `json:"jns"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
