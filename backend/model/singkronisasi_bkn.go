package model

import ()

type UpdateBKNDataUtama struct {
	NoRefBkn          string `json:"pns_orang_id" validate:"required"`
	Kagama            string `json:"agama_id,omitempty"`
	Alamat            string `json:"alamat,omitempty"`
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
