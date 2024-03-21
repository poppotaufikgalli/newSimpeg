package model

import (
	"time"
)

func (RiwayatTugasTambahan) TableName() string {
	return "master_riwayat_tugas_tambahan"
}

type SearchRiwayatTugasTambahan struct {
	Nip    string    `json:"nip,omitempty"`
	Tmtjab time.Time `json:"tmtjab,omitempty"`
	Njab   string    `json:"njab,omitempty"`
	Nunker string    `json:"nunker,omitempty"`
	IdOpd  string    `json:"id_opd,omitempty"`
	Status []int     `json:"status,omitempty"`
}

type DeleteRiwayatTugasTambahan struct {
	Nip    string    `json:"nip" validate:"required"`
	Tmtjab time.Time `json:"tmtjab" validate:"required"`
}

type RiwayatTugasTambahan struct {
	Nip       string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Tmtjab    time.Time  `json:"tmtjab" validate:"required"`
	IdJab     *int       `json:"id_jab"`
	Kjab      *string    `json:"kjab"`
	Njab      *string    `json:"njab"`
	Nunker    *string    `json:"nunker"`
	Kpej      *float64   `json:"kpej"`
	Nskjabat  *string    `json:"nskjabat"`
	Tskjabat  *time.Time `json:"tskjabat"`
	IdOpd     *string    `json:"id_opd"`
	Status    *int       `json:"status"`
	IdSubOpd  *string    `json:"id_sub_opd"`
	Filename  *string    `json:"filename"`
	Uuid      *string    `json:"uuid"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
