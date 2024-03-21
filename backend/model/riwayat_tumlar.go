package model

import (
	"time"
)

func (RiwayatTumlar) TableName() string {
	return "master_riwayat_tumlar"
}

type SearchRiwayatTumlar struct {
	Nip  string    `json:"nip,omitempty"`
	Tsk  time.Time `json:"tsk,omitempty"`
	Ktpu string    `json:"ktpu,omitempty"`
	Kjur string    `json:"kjur,omitempty"`
}

type DeleteRiwayatTumlar struct {
	Nip string    `json:"nip" validate:"required"`
	Tsk time.Time `json:"tsk" validate:"required"`
}

type RiwayatTumlar struct {
	Nip       string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Ptetap    *float64   `json:"ptetap"`
	Nsk       *string    `json:"nsk"`
	Tsk       time.Time  `json:"tsk" validate:"required"`
	Ktpu      string     `json:"ktpu"`
	Kjur      string     `json:"kjur"`
	Gldepan   *string    `json:"gldepan"`
	Glblk     *string    `json:"glblk"`
	Filename  *string    `json:"filename"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
