package model

import (
	"time"
)

func (Opd) TableName() string {
	return "master_opd"
}

type SearchOpd struct {
	Id        string    `json:"id,omitempty"`         //
	ParentOpd string    `json:"parent_opd,omitempty"` //
	Kunker    string    `json:"kunker,omitempty"`     //
	Nama      string    `json:"nama,omitempty"`       //
	IdEselon  []float64 `json:"id_eselon,omitempty"`  //
	Status    []int     `json:"status,omitempty"`     //
	Sfilter   []int     `json:"sfilter,omitempty"`    //
	Disingkat string    `json:"disingkat,omitempty"`
	IdUnorBkn string    `json:"id_unor_bkn,omitempty"`
}

type DeleteOpd struct {
	Id string `json:"id" validate:"required"`
}

type Opd struct {
	Id                 string           `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	ParentOpd          *string          `json:"parent_opd"`
	Kunker             *string          `json:"kunker" validate:"required"`
	Nama               string           `json:"nama" validate:"required"`
	IdJabatan          *string          `json:"id_jabatan"`
	NamaJabatan        *string          `json:"nama_jabatan"`
	IdEselon           float64          `json:"id_eselon"`
	Alamat             *string          `json:"alamat"`
	Status             *int             `json:"status"`
	Tkpem              string           `json:"tkpem"`
	Kunkom             string           `json:"kunkom"`
	Kununit            string           `json:"kununit"`
	Ksatker            string           `json:"ksatker"`
	Kssatker           string           `json:"kssatker"`
	Levelunker         int              `json:"levelunker"`
	Nip1               *string          `json:"nip1"`
	Nip2               *string          `json:"nip2"`
	Sreport            *int             `json:"sreport"`
	Sfilter            *int             `json:"sfilter"`
	Disingkat          *string          `json:"disingkat"`
	Peringkat          int              `json:"peringkat"`
	IdUnorBkn          *string          `json:"id_unor_bkn"`
	SubOpd             []Opd            `gorm:"foreignKey:ParentOpd" json:"sub_opd"`
	ListFormasiJabatan []FormasiJabatan `gorm:"foreignKey:IdOpd" json:"formasi_jabatan"`
	CreatedBy          string           `gorm:"<-:create" json:"created_by"`
	CreatedAt          *time.Time       `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy          string           `gorm:"<-:update" json:"updated_by"`
	UpdatedAt          *time.Time       `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
