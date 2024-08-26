package model

import (
	"fmt"
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
	Jnsjab   int    `json:"jnsjab,omitempty"`
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

	if u.Gldepan == nil || len(*u.Gldepan) == 0 {
		if u.Glblk == nil || len(*u.Glblk) == 0 {
			u.Namapeg = u.Nama
			fmt.Println("1")
		} else {
			fmt.Println("2")
			u.Namapeg = fmt.Sprintf("%s, %s", u.Nama, *u.Glblk)
		}
	} else {
		if u.Glblk == nil || len(*u.Glblk) == 0 {
			fmt.Println("3")
			u.Namapeg = fmt.Sprintf("%s. %s", *u.Gldepan, u.Nama)
		} else {
			fmt.Println("4")
			u.Namapeg = fmt.Sprintf("%s. %s, %s", *u.Gldepan, u.Nama, *u.Glblk)
		}
	}

	return
}

type Pegawai struct {
	Nip               string           `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Nama              string           `json:"nama" validate:"required"`
	Namapeg           string           `gorm:"-" json:"namapeg"`
	Gldepan           *string          `json:"gldepan"`
	Glblk             *string          `json:"glblk"`
	Ktlahir           *string          `json:"ktlahir"`
	Tlahir            *time.Time       `json:"tlahir" validate:"required"`
	Umur              int              `gorm:"-" json:"umur"`
	Kjkel             float64          `json:"kjkel"`
	Kagama            float64          `json:"kagama"`
	Kstatus           float64          `json:"kstatus"`
	Kjpeg             float64          `json:"kjpeg"`
	Kduduk            float64          `json:"kduduk"`
	Kskawin           float64          `json:"kskawin"`
	TinggiBadan       int              `json:"tinggi_badan"`
	BeratBadan        int              `json:"berat_badan"`
	Kgoldar           float64          `json:"kgoldar"`
	Aljalan           *string          `json:"aljalan"`
	Alrt              *string          `json:"alrt"`
	Alrw              *string          `json:"alrw"`
	Altelp            *string          `json:"altelp"`
	Altelpwa          *string          `json:"altelpwa"`
	Alkoprop          *string          `json:"alkoprop"`
	Alkokab           *string          `json:"alkokab"`
	Alkokec           *string          `json:"alkokec"`
	Alkodes           *string          `json:"alkodes"`
	Kpos              *string          `json:"kpos"`
	Kaparpol          float64          `json:"kaparpol"`
	NpapG             *string          `json:"npap_g"`
	Nkarpeg           *string          `json:"nkarpeg"`
	Naskes            *string          `json:"naskes"`
	Bpjs              *string          `json:"bpjs"`
	KartuAsn          *string          `json:"kartu_asn"`
	Ntaspen           *string          `json:"ntaspen"`
	NkarisSu          *string          `json:"nkaris_su"`
	Npwp              *string          `json:"npwp"`
	Nik               *string          `json:"nik"`
	Tapera            *string          `json:"tapera"`
	FileBmp           *string          `json:"file_bmp"`
	Aktif             float64          `json:"aktif"`
	Jjiwa             float64          `json:"jjiwa"`
	Marga             *string          `json:"marga"`
	Suku              *string          `json:"suku"`
	TglPeg            *time.Time       `json:"tgl_peg"`
	Niplama           *string          `json:"niplama"`
	TglReg            *time.Time       `json:"tgl_reg"`
	StatKpe           int              `json:"stat_kpe"`
	NoPinpeg          int              `json:"no_pinpeg"`
	Alkoproplahir     *string          `json:"alkoproplahir"`
	Alkokablahir      *string          `json:"alkokablahir"`
	TglKpe            *time.Time       `json:"tgl_kpe"`
	ThnPendataan      int              `json:"thn_pendataan"`
	NoRefBkn          *string          `json:"no_ref_bkn"`
	Email             *string          `json:"email,omitempty" validate:"omitempty,email"`
	EmailPemerintahan *string          `json:"email_pemerintahan,omitempty" validate:"omitempty,email"`
	RiwayatJabatan    []RiwayatJabatan `gorm:"foreignKey:Nip" json:"riwayat_jabatan" validate:"-"`
	JabatanAkhir      RiwayatJabatan   `gorm:"foreignKey:Nip" json:"jabatan_akhir" validate:"-"`
	RiwayatPangkat    []RiwayatPangkat `gorm:"foreignKey:Nip" json:"riwayat_pangkat" validate:"-"`
	PangkatAkhir      RiwayatPangkat   `gorm:"foreignKey:Nip" json:"pangkat_akhir" validate:"-"`
	CreatedBy         string           `gorm:"<-:create" json:"created_by"`
	CreatedAt         *time.Time       `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy         string           `gorm:"<-:update" json:"updated_by"`
	UpdatedAt         *time.Time       `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}

type PegawaiList struct {
	Nip        string     `gorm:"primaryKey;autoIncrement:false" json:"nip" validate:"required"`
	Nama       string     `json:"nama"`
	Gldepan    *string    `json:"gldepan"`
	Glblk      *string    `json:"glblk"`
	Npwp       *string    `json:"npwp"`
	Nik        *string    `json:"nik"`
	Ktlahir    *string    `json:"ktlahir"`
	Tlahir     *time.Time `json:"tlahir" validate:"required"`
	Umur       *string    `json:"umur"`
	Kjkel      float64    `json:"kjkel"`
	Aljalan    *string    `json:"aljalan"`
	Alrt       *string    `json:"alrt"`
	Alrw       *string    `json:"alrw"`
	Alkoprop   *string    `json:"alkoprop"`
	Alkokab    *string    `json:"alkokab"`
	Alkokec    *string    `json:"alkokec"`
	Alkodes    *string    `json:"alkodes"`
	Kstatus    int        `json:"kstatus"`
	Nstatus    *string    `json:"nstatus"`
	Kjpeg      int        `json:"kjpeg"`
	NoRefBkn   string     `json:"no_ref_bkn"`
	IdOpd      *string    `json:"id_opd"`
	Jnsjab     int        `json:"jnsjab"`
	Keselon    *int       `json:"keselon"`
	Neselon    *string    `json:"neselon"`
	Njab       *string    `json:"njab"`
	Tmtjab     *time.Time `json:"tmtjab"`
	Nskjabat   *string    `json:"nskjabat"`
	Nunker     *string    `json:"nunker"`
	Nopd       *string    `json:"nopd"`
	Kgolru     *string    `json:"kgolru"`
	Tmtpang    *time.Time `json:"tmtpang"`
	KgolruCpns *string    `json:"kgolru_cpns"`
	Tmtcpns    *time.Time `json:"tmtcpns"`
	KgolruPns  *string    `json:"kgolru_pns"`
	Tmtpns     *time.Time `json:"tmtpns"`
	Ntpu       *string    `json:"ntpu"`
	Njur       *string    `json:"njur"`
	Nsek       *string    `json:"nsek"`
	Dikstr     *string    `json:"dik_str"`
	DikFung    *string    `json:"dik_fung"`
}
