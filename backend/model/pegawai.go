package model

import (
	"gorm.io/gorm"
	"time"

	age "github.com/theTardigrade/golang-age"
)

func (Pegawai) TableName() string {
	return "master_pegawai"
}

type SearchPegawai struct {
	Nip      string `json:"nip,omitempty"`
	Nama     string `json:"nama,omitempty"`
	IdOpd    string `json:"id_opd,omitempty"`
	Keselon  int    `json:"keselon,omitempty"`
	Jnsjab   []int  `json:"jnsjab,omitempty"`
	Kstatus  []int  `json:"kstatus,omitempty"`
	Kjpeg    []int  `json:"kjpeg,omitempty"`
	NoRefBkn string `json:"no_ref_bkn,omitempty"`
}

type DeletePegawai struct {
	Nip string `json:"nip" validate:"required"`
}

func (u *Pegawai) AfterFind(tx *gorm.DB) (err error) {
	if u.Tlahir != nil {
		dur := age.CalculateToNow(*u.Tlahir)
		u.Umur = dur // Set default role if not specified
	}
	return
}

type Pegawai struct {
	Nip            string           `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Nama           string           `json:"nama" validate:"required"`
	Gldepan        *string          `json:"gldepan"`
	Glblk          *string          `json:"glblk"`
	Ktlahir        *string          `json:"ktlahir"`
	Tlahir         *time.Time       `json:"tlahir" validate:"required"`
	Umur           int              `gorm:"-" json:"umur"`
	Kjkel          float64          `json:"kjkel"`
	Kagama         float64          `json:"kagama"`
	Kstatus        float64          `json:"kstatus"`
	Kjpeg          float64          `json:"kjpeg"`
	Kduduk         float64          `json:"kduduk"`
	Kskawin        float64          `json:"kskawin"`
	TinggiBadan    int              `json:"tinggi_badan"`
	BeratBadan     int              `json:"berat_badan"`
	Kgoldar        float64          `json:"kgoldar"`
	Aljalan        *string          `json:"aljalan"`
	Alrt           *string          `json:"alrt"`
	Alrw           *string          `json:"alrw"`
	Altelp         *string          `json:"altelp"`
	Altelpwa       *string          `json:"altelpwa"`
	Alkoprop       *string          `json:"alkoprop"`
	Alkokab        *string          `json:"alkokab"`
	Alkokec        *string          `json:"alkokec"`
	Alkodes        *string          `json:"alkodes"`
	Kpos           float64          `json:"kpos"`
	Kaparpol       float64          `json:"kaparpol"`
	NpapG          *string          `json:"npap_g"`
	Nkarpeg        *string          `json:"nkarpeg"`
	Naskes         *string          `json:"naskes"`
	Ntaspen        *string          `json:"ntaspen"`
	NkarisSu       *string          `json:"nkaris_su"`
	Npwp           *string          `json:"npwp"`
	Nik            *string          `json:"nik"`
	FileBmp        *string          `json:"file_bmp"`
	Aktif          float64          `json:"aktif"`
	Jjiwa          float64          `json:"jjiwa"`
	Marga          *string          `json:"marga"`
	Suku           *string          `json:"suku"`
	TglPeg         *time.Time       `json:"tgl_peg"`
	Niplama        *string          `json:"niplama"`
	TglReg         *time.Time       `json:"tgl_reg"`
	StatKpe        int              `json:"statkpe"`
	NoPinpeg       int              `json:"no_pinpeg"`
	Alkoproplahir  *string          `json:"alkoproplahir"`
	Alkokablahir   *string          `json:"alkokablahir"`
	TglKpe         *time.Time       `json:"tgl_kpe"`
	ThnPendataan   int              `json:"thn_pendataan"`
	NoRefBkn       *string          `json:"no_ref_bkn"`
	Email          *string          `json:"email"  validate:"email"`
	RiwayatJabatan []RiwayatJabatan `gorm:"foreignKey:Nip" json:"riwayat_jabatan"`
	JabatanAkhir   RiwayatJabatan   `gorm:"foreignKey:Nip" json:"jabatan_akhir"`
	CreatedBy      string           `gorm:"<-:create" json:"created_by"`
	CreatedAt      *time.Time       `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy      string           `gorm:"<-:update" json:"updated_by"`
	UpdatedAt      *time.Time       `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
