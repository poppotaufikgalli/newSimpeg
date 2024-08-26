package app

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	model "newSimpegAPI/model"
)

func GetRiwayatAngkaKredit(searchString model.SearchRiwayatAngkaKredit) (riwayat_angka_kredit []model.RiwayatAngkaKredit, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatAngkaKredit{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_angka_kredit.nip = ?", searchString.Nip)
	}

	//Jns
	if searchString.Jns != "" {
		result = result.Where("master_riwayat_angka_kredit.jns = ?", searchString.Jns)
	}

	//kjab
	if searchString.Kjab != "" {
		result = result.Where("master_riwayat_angka_kredit.kjab = ?", searchString.Kjab)
	}

	//tmtjab
	if searchString.Tmulai.IsZero() != true {
		result = result.Where("master_riwayat_angka_kredit.tmulai = ?", searchString.Tmulai)
	}

	//kjab bkn
	if searchString.Thn != "" {
		result = result.Where("master_riwayat_angka_kredit.thn = ?", searchString.Thn)
	}

	result = result.Preload("JabatanFt").Preload("JenjangJabfung").Preload("RwJabatan").Find(&riwayat_angka_kredit)

	return

}

func FindRiwayatAngkaKredit(c echo.Context) error {
	var searchString model.SearchRiwayatAngkaKredit
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_angka_kredit, result := GetRiwayatAngkaKredit(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_angka_kredit,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetRiwayatAngkaKreditByNipTmulai(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_angka_kredit model.RiwayatAngkaKredit
	nip := c.Param("nip")
	tmulai := c.Param("tmulai")

	//tmulai = (tmulai).Format("2006-01-02")
	result := db.Model(&model.RiwayatAngkaKredit{}).Where("nip = ? and tmulai = ? ", nip, tmulai).Scan(&riwayat_angka_kredit)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_angka_kredit,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateRiwayatAngkaKredit(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_angka_kredit model.RiwayatAngkaKredit
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_angka_kredit)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"status": http.StatusNotImplemented,
			"error":  err.Error(),
		})
	}

	if err = validate.Struct(riwayat_angka_kredit); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	riwayat_angka_kredit.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_angka_kredit)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"error":  result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_angka_kredit,
		"statusCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatAngkaKredit(c echo.Context) error {
	db, _ := model.CreateCon()

	//var riwayat_angka_kredit model.RiwayatAngkaKredit
	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatAngkaKredit `json:"refdata"`
		Data model.RiwayatAngkaKredit       `json:"data"`
	}

	var reqData ToUpdateData

	err := json.NewDecoder(c.Request().Body).Decode(&reqData)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(reqData.Data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	reqData.Data.UpdatedBy = fmt.Sprint(c.Get("nip"))
	Tmulai := (reqData.Ref.Tmulai).Format("2006-01-02")
	result := db.Model(&model.RiwayatAngkaKredit{}).Where("nip = ? and tmulai = ?", reqData.Ref.Nip, Tmulai).Updates(&reqData.Data)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       reqData.Data,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatAngkaKredit(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_angka_kredit model.DeleteRiwayatAngkaKredit
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_angka_kredit)

	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"title":   "Gagal Menghapus Data Angka Kredit Simpeg",
			"status":  http.StatusNotImplemented,
			"message": err.Error(),
			"color":   "red",
		})
	}

	if err = validate.Struct(riwayat_angka_kredit); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"title":   "Gagal Menghapus Data Angka Kredit Simpeg",
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"color":   "red",
		})
	}

	result := db.Model(&model.RiwayatAngkaKredit{}).Where("nip = ? and tmulai = ?", riwayat_angka_kredit.Nip, riwayat_angka_kredit.Tmulai).Delete(&riwayat_angka_kredit)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"title":   "Gagal Menghapus Data Angka Kredit Simpeg",
			"status":  http.StatusNotFound,
			"message": "Data Angka Kredit Tidak ditemukan",
			"color":   "red",
		})
	} else {
		if c.Get("gid").(int) == 1 {
			resultBkn, err := DeleteSingkronisasiDataBknById("angkakredit", *riwayat_angka_kredit.IdSync, c.Get("nip").(string))
			fmt.Println(resultBkn)

			if err != nil {
				return c.JSON(http.StatusNotImplemented, map[string]interface{}{
					"statusCode": http.StatusNotImplemented,
					"errors":     err.Error(),
				})
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"title":      "Hapus Data Angka Kredit Simpeg Berhasil",
		"data":       riwayat_angka_kredit,
		"message":    "Data Angka Kredit Berhasil dihapus",
		"color":      "green",
		"statucCode": http.StatusOK,
	})
}

func sendRiwayatAngkaKreditBKN(riwayat_angka_kredit model.RiwayatAngkaKredit, CreatedBy string) (status_bkn map[string]interface{}, err error) {
	db, _ := model.CreateCon()

	var pegawai model.Pegawai
	db.Select("no_ref_bkn").First(&pegawai, "nip = ?", riwayat_angka_kredit.Nip)

	Tsk := (riwayat_angka_kredit.Tsk).Format("02-01-2006")

	updatas := map[string]interface{}{
		"bulanMulaiPenailan":   fmt.Sprintf("%v", int(riwayat_angka_kredit.Tmulai.Month())),
		"bulanSelesaiPenailan": fmt.Sprintf("%v", int(riwayat_angka_kredit.Tselesai.Month())),
		"isAngkaKreditPertama": "",
		"isIntegrasi":          "",
		"isKonversi":           "",
		"kreditBaruTotal":      fmt.Sprintf("%v", *riwayat_angka_kredit.Total),
		"kreditPenunjangBaru":  fmt.Sprintf("%v", *riwayat_angka_kredit.Tambahan),
		"kreditUtamaBaru":      fmt.Sprintf("%v", riwayat_angka_kredit.Utama),
		"nomorSk":              *riwayat_angka_kredit.Nsk,
		"pnsId":                *pegawai.NoRefBkn,
		"rwJabatanId":          *riwayat_angka_kredit.RwJabatanIdBkn,
		"tahunMulaiPenailan":   fmt.Sprintf("%v", riwayat_angka_kredit.Tmulai.Year()),
		"tahunSelesaiPenailan": fmt.Sprintf("%v", riwayat_angka_kredit.Tselesai.Year()),
		"tanggalSk":            Tsk,
	}

	if riwayat_angka_kredit.Jns == "Pertama" {
		updatas["isAngkaKreditPertama"] = "1"
	} else if riwayat_angka_kredit.Jns == "Integrasi" {
		updatas["isIntegrasi"] = "1"
	} else if riwayat_angka_kredit.Jns == "Konversi" {
		updatas["isKonversi"] = "1"
	}

	if riwayat_angka_kredit.IdSync != nil {
		updatas["id"] = *riwayat_angka_kredit.IdSync
	}

	status_bkn, err = PostSinkronisasiGetDataBKN(updatas, "angkakredit", CreatedBy)

	if err != nil {
		return
	}

	riwayat_update_bkn := &model.RiwayatUpdateBKN{
		Deskripsi: fmt.Sprintf("Riwayat Angka Kredit %s, No SK %s", riwayat_angka_kredit.Nip, *riwayat_angka_kredit.Nsk),
		CreatedBy: CreatedBy,
		Code:      fmt.Sprint(status_bkn["code"]),
		Message:   fmt.Sprint(status_bkn["message"]),
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create*/

	if status_bkn["success"] == true || status_bkn["code"] == 1 {
		mapData := status_bkn["mapData"]
		//fmt.Println(mapData)
		mapmapData := mapData.(map[string]interface{})

		Tmulai := (riwayat_angka_kredit.Tmulai).Format("2006-01-02")
		db.Model(&model.RiwayatAngkaKredit{}).Debug().Where("nip = ? and tmulai = ?", riwayat_angka_kredit.Nip, Tmulai).Update("idSync", mapmapData["rwAngkaKreditId"])
	}

	return
}

func UpdateRiwayatAngkaKreditBkn(c echo.Context) error {
	var riwayat_angka_kredit model.RiwayatAngkaKredit

	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_angka_kredit)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"status": http.StatusNotImplemented,
			"error":  err.Error(),
		})
	}

	if err = validate.Struct(riwayat_angka_kredit); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	if c.Get("gid").(int) == 1 {
		status_bkn, err := sendRiwayatAngkaKreditBKN(riwayat_angka_kredit, c.Get("nip").(string))
		//fmt.Println(status_bkn)

		if err != nil {
			return c.JSON(http.StatusNotImplemented, map[string]interface{}{
				"status": http.StatusNotImplemented,
				"error":  err.Error(),
			})
		}
		if status_bkn["success"] == false {
			return c.JSON(http.StatusNotImplemented, map[string]string{
				"title":       "Tambah Angka Kredit BKN Gagal",
				"description": fmt.Sprintf("%s", status_bkn["message"]),
				"error":       fmt.Sprintf("%s", status_bkn["message"]),
			})
		}
	} else {
		return c.JSON(http.StatusNotImplemented, map[string]string{
			"title":       "Update BKN Gagal",
			"description": "Anda tidak memiliki hak akses",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_angka_kredit,
		"statusCode": http.StatusOK,
	})
}

func sendDeleteRiwayatAngkaKreditBKN(riwayat_angka_kredit model.DeleteRiwayatAngkaKredit, CreatedBy string) (status_bkn map[string]interface{}, err error) {
	db, _ := model.CreateCon()

	var singkronisasi model.Singkronisasi
	host := "bkn"
	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/angkakredit/delete/%s", *riwayat_angka_kredit.IdSync)

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	d := fmt.Sprintf("bearer %v", *singkronisasi.TokenApimanager)
	e := fmt.Sprintf("Bearer %v", *singkronisasi.TokenSso)

	req.Header.Add("Auth", d)
	req.Header.Add("Authorization", e)

	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&status_bkn)

	riwayat_update_bkn := &model.RiwayatUpdateBKN{
		Deskripsi: fmt.Sprintf("Delete Riwayat Angka Kredit %s %s", riwayat_angka_kredit.Tmulai, riwayat_angka_kredit.Nip),
		CreatedBy: CreatedBy,
		Code:      fmt.Sprint(status_bkn["code"]),
		Message:   fmt.Sprint(status_bkn["message"]),
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create*/

	return
}
