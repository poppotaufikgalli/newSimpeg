package model

import (
	"time"
)

func (RiwayatUpdateBKN) TableName() string {
	return "master_riwayat_update_bkn"
}

type RiwayatUpdateBKN struct {
	ID        uint       `gorm:"primaryKey;autoIncrement:false" json:"id" validate:"required"`
	Deskripsi string     `json:"deskripsi"`
	Content   string     `json:"content"`
	Code      string     `json:"code"`
	Message   string     `json:"message"`
	CreatedBy string     `gorm:"<-:create" json:"created_by"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}

type UpdateBKNDataUtama struct {
	NoRefBkn          string `json:"pns_orang_id" validate:"required"`
	Kagama            string `json:"agama_id,omitempty"`
	Aljalan           string `json:"alamat,omitempty"`
	Email             string `json:"email,omitempty"`
	EmailPemerintahan string `json:"email_gov,omitempty"`
	NkarisSu          string `json:"karis_karsu,omitempty"`
	Bpjs              string `json:"nomor_bpjs,omitempty"`
	Altelp            string `json:"nomor_telpon,omitempty"`
	AltelpWa          string `json:"nomor_hp,omitempty"`
	Npwp              string `json:"npwp_nomor,omitempty"`
	Tapera            string `json:"tapera_nomor,omitempty"`
	Ntaspen           string `json:"taspen_nomor,omitempty"`
	//kabupaten_id    *string `json:"kabupaten_id,omitempty"`
	//kelas_jabatan   *string `json:"kelas_jabatan,omitempty"`
	//kpkn_id         *string `json:"kpkn_id,omitempty"`
	//lokasi_kerja_id *string `json:"lokasi_kerja_id,omitempty"`
	//nomor_telpon    *string `json:"nomor_telpon,omitempty"`
	//npwp_tanggal    *string `json:"npwp_tanggal,omitempty"`
	//tanggal_taspen  *string `json:"tanggal_taspen,omitempty"`
}

type DocKPBKN struct {
	Path           string               `json:"path"`
	RiwayatPangkat DeleteRiwayatPangkat `json:"riwayat_pangkat"`
}

type UpdateAngkaKreditBKN struct {
	Id                     string `json:"id,omitempty"`
	PnsId                  string `json:"pnsId"`
	BulanMulaiPenilaian    string `json:"bulanMulaiPenailan"`
	BulanSelesiaiPenilaian string `json:"bulanSelesaiPenailan"`
	IsAngkaKreditPertama   string `json:"isAngkaKreditPertama,omitempty"`
	IsIntegrasi            string `json:"isIntegrasi,omitempty"`
	IsKonversi             string `json:"isKonversi,omitempty"`
	KreditBaruTotal        string `json:"kreditBaruTotal"`
	KreditPenunjangBaru    string `json:"kreditPenunjangBaru":`
	KreditUtamaBaru        string `json:"kreditUtamaBaru"`
	NomorSk                string `json:"nomorSk"`
	RwJabatanId            string `json:"rwJabatanId,omitempty"`
	TahunMulaiPenailan     string `json:"tahunMulaiPenailan"`
	TahunSelesaiPenailan   string `json:"tahunSelesaiPenailan"`
	TanggalSk              string `json:"tanggalSk"`
}
