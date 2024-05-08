package model

import (
	"time"
)

func (RiwayatPendum) TableName() string {
	return "master_riwayat_pendum"
}

type SearchRiwayatPendum struct {
	Nip    string `json:"nip,omitempty"`
	Ktpu   string `json:"ktpu,omitempty"`
	Kjur   string `json:"kjur,omitempty"`
	Nsek   string `json:"nsek,omitempty"`
	Tempat string `json:"tempat,omitempty"`
}

type DeleteRiwayatPendum struct {
	Nip  string `json:"nip" validate:"required"`
	Ktpu string `json:"ktpu" validate:"required"`
	Kjur string `json:"kjur" validate:"required"`
}

type RiwayatPendum struct {
	Nip               string            `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Ktpu              string            `json:"ktpu" validate:"required"`
	TingkatPendidikan TingkatPendidikan `gorm:"foreignKey:Ktpu" json:"TingkatPendidikan" validate:"-"`
	Kjur              string            `json:"kjur" validate:"required"`
	Pendidikan        Pendidikan        `gorm:"foreignKey:Kjur" json:"Pendidikan" validate:"-"`
	Nsek              *string           `json:"nsek"`
	Tempat            *string           `json:"tempat"`
	Nkepsek           *string           `json:"nkepsek"`
	Nsttb             *string           `json:"nsttb"`
	Tsttb             *time.Time        `json:"tsttb"`
	Akhir             *int              `json:"akhir"`
	Npdum             *string           `json:"npdum"`
	Negara            *string           `json:"negara"`
	Ket               *string           `json:"ket"`
	Filename          *string           `json:"filename"`
	CreatedBy         string            `gorm:"<-:create" json:"created_by"`
	CreatedAt         *time.Time        `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy         string            `gorm:"<-:update" json:"updated_by"`
	UpdatedAt         *time.Time        `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
