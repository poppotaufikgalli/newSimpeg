package model

import (
	"time"
)

func (RiwayatAngkaKredit) TableName() string {
	return "master_riwayat_angka_kredit"
}

type SearchRiwayatAngkaKredit struct {
	Nip    string    `json:"nip,omitempty"`
	Jns    string    `json:"jns,omitempty"`
	Kjab   string    `json:"kjab,omitempty"`
	Tmulai time.Time `json:"tmulai,omitempty"`
	Thn    string    `json:"thn,omitempty"`
}

type DeleteRiwayatAngkaKredit struct {
	Nip    string    `json:"nip" validate:"required"`
	Tmulai time.Time `json:"tmulai" validate:"required"`
	IdSync *string   `json:"idSync,omitempty"`
}

type RiwayatAngkaKredit struct {
	Nip            string         `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Jns            string         `json:"jns"`
	JenjangJabfung JenjangJabfung `gorm:"foreignKey:Jns" json:"jenjang_ft" validate:"-"`
	Kjab           string         `json:"kjab"`
	RwJabatanId    uint           `json:"rw_jabatan_id"`
	RwJabatan      RiwayatJabatan `gorm:"foreignKey:RwJabatanId" validate:"-"`
	RwJabatanIdBkn *string        `json:"rw_jabatan_id_bkn"`
	JabatanFt      JabatanFt      `gorm:"foreignKey:Kjab" json:"jabatan_ft" validate:"-"`
	Tmulai         *time.Time     `json:"tmulai" validate:"required"`
	Tselesai       *time.Time     `json:"tselesai"`
	Utama          float64        `json:"utama"`
	Tambahan       *float64       `json:"tambahan"`
	Total          *float64       `json:"total"`
	Nsk            *string        `json:"nsk"`
	Tsk            *time.Time     `json:"tsk"`
	Thn            *string        `json:"thn"`
	Filename       *string        `json:"filename"`
	FilenameSkPak  *string        `json:"filename_sk_pak"`
	FilenameDokPak *string        `json:"filename_dok_pak"`
	IdSync         *string        `gorm:"column:idSync" json:"idSync"`
	DoBkn          *int           `gorm:"-" json:"do_bkn"`
	CreatedBy      string         `gorm:"<-:create" json:"created_by"`
	CreatedAt      *time.Time     `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy      string         `gorm:"<-:update" json:"updated_by"`
	UpdatedAt      *time.Time     `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
