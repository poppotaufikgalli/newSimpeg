package app

import (
	"bytes"
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	model "newSimpegAPI/model"
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^\p{L}\p{N} ]+`)

func GetRiwayatSkp(searchString model.SearchRiwayatSkp) (riwayat_skp []model.RiwayatSkp, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatSkp{})

	//ID
	if searchString.ID > 0 {
		result = result.Where("master_riwayat_skp.id = ?", searchString.ID)
	}

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_skp.nip = ?", searchString.Nip)
	}

	//tahun
	if searchString.Tahun > 0 {
		result = result.Where("master_riwayat_skp.tahun = ?", searchString.Tahun)
	}

	//pns dinilai orang
	if searchString.PnsDinilaiOrang != "" {
		nama := strings.TrimSpace(searchString.PnsDinilaiOrang)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_skp.pnsDinilaiOrang LIKE ?", strings.Join(str, ""))
	}

	result = result.Order("tahun desc").Find(&riwayat_skp)

	return
}

func FindRiwayatSkp(c echo.Context) error {
	var searchString model.SearchRiwayatSkp
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_skp, result := GetRiwayatSkp(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_skp,
		"statusCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetRiwayatSkpByNipIdTahun(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_skp model.RiwayatSkp
	Nip := c.Param("nip")
	ID := c.Param("id")
	Tahun := c.Param("tahun")

	result := db.Model(model.RiwayatSkp{}).Debug().Where("id = ? and nip = ? and tahun = ?", ID, Nip, Tahun).Scan(&riwayat_skp)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statusCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_skp,
		"statusCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateRiwayatSkp(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_skp model.RiwayatSkp
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_skp)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statusCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_skp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_skp.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_skp)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_skp,
		"statusCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatSkp(c echo.Context) error {
	db, _ := model.CreateCon()

	//var riwayat_skp model.RiwayatSkp
	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatSkp `json:"refdata"`
		Data model.RiwayatSkp       `json:"data"`
	}

	var reqData ToUpdateData

	err := json.NewDecoder(c.Request().Body).Decode(&reqData)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statusCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(reqData.Data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	reqData.Data.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RiwayatSkp{}).Where("nip = ? and tahun = ? and id=?", reqData.Ref.Nip, reqData.Ref.Tahun, reqData.Ref.ID).Updates(&reqData.Data)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statusCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       reqData.Data,
		"statusCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatSkp(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_skp model.DeleteRiwayatSkp
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_skp)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statusCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_skp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatSkp{}).Debug().Where("id = ? and nip = ? and tahun = ?", riwayat_skp.ID, riwayat_skp.Nip, riwayat_skp.Tahun).Delete(&riwayat_skp)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statusCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_skp,
		"statusCode": http.StatusOK,
	})
}

func checkPointer(str *string) string {
	if str != nil {
		//return *str
		return strings.TrimSpace(*str)
	}
	return ""
}

func UpdateRiwayatSkpBkn(c echo.Context) error {
	var riwayat_skp model.RiwayatSkp
	validate := validator.New()
	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_skp)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statusCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_skp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	if c.Get("gid").(int) == 1 {
		status_bkn, err := sendRiwayatSkpBKN(riwayat_skp, c.Get("nip").(string))
		fmt.Println(status_bkn)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"statucCode": http.StatusBadRequest,
				"errors":     err.Error(),
			})
		}

		if status_bkn["success"] != true {
			return c.JSON(http.StatusNotImplemented, map[string]string{
				"title":       "Update SKP BKN Gagal",
				"description": fmt.Sprintf("%s", status_bkn["message"]),
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_skp,
		"statusCode": http.StatusOK,
	})
}

func sendRiwayatSkpBKN(riwayat_skp model.RiwayatSkp, CreatedBy string) (status_bkn map[string]interface{}, err error) {
	db, _ := model.CreateCon()

	var pegawai model.Pegawai
	db.Select("no_ref_bkn").First(&pegawai, "nip = ?", riwayat_skp.Nip)

	//AtasanPenilaiTmtGolongan := (riwayat_skp.AtasanPenilaiTmtGolongan).Format("02-01-2006")
	//PenilaiTmtGolongan := (riwayat_skp.PenilaiTmtGolongan).Format("02-01-2006")
	updatas := map[string]interface{}{
		"atasanPejabatPenilai":    checkPointer(riwayat_skp.AtasanPejabatPenilai),
		"atasanPenilaiGolongan":   checkPointer(riwayat_skp.AtasanPenilaiGolongan),
		"atasanPenilaiJabatan":    checkPointer(riwayat_skp.AtasanPenilaiJabatan),
		"atasanPenilaiNama":       checkPointer(riwayat_skp.AtasanPenilaiNama),
		"atasanPenilaiNipNrp":     checkPointer(riwayat_skp.AtasanPenilaiNipNrp),
		"atasanPenilaiUnorNama":   checkPointer(riwayat_skp.AtasanPenilaiUnorNama),
		"disiplin":                riwayat_skp.Disiplin,
		"id":                      checkPointer(riwayat_skp.IdSync),
		"integritas":              riwayat_skp.Integritas,
		"jenisJabatan":            riwayat_skp.JenisJabatan,
		"jumlah":                  riwayat_skp.Jumlah,
		"kepemimpinan":            riwayat_skp.Kepemimpinan,
		"kerjasama":               riwayat_skp.Kerjasama,
		"komitmen":                riwayat_skp.Komitmen,
		"nilaiPerilakuKerja":      riwayat_skp.NilaiPerilakuKerja,
		"nilaiPrestasiKerja":      riwayat_skp.NilaiPrestasiKerja,
		"nilaiSkp":                riwayat_skp.NilaiSkp,
		"nilairatarata":           riwayat_skp.Nilairatarata,
		"orientasiPelayanan":      riwayat_skp.OrientasiPelayanan,
		"pejabatPenilai":          checkPointer(riwayat_skp.PejabatPenilai),
		"penilaiGolongan":         checkPointer(riwayat_skp.PenilaiGolongan),
		"penilaiJabatan":          checkPointer(riwayat_skp.PenilaiJabatan),
		"penilaiNama":             checkPointer(riwayat_skp.PenilaiNama),
		"penilaiNipNrp":           checkPointer(riwayat_skp.PenilaiNipNrp),
		"penilaiUnorNama":         checkPointer(riwayat_skp.PenilaiUnorNama),
		"pnsDinilaiOrang":         *pegawai.NoRefBkn,
		"statusAtasanPenilai":     checkPointer(riwayat_skp.StatusAtasanPenilai),
		"statusPenilai":           checkPointer(riwayat_skp.StatusPenilai),
		"jenisPeraturanKinerjaKd": checkPointer(riwayat_skp.JenisPeraturanKinerjaKd),
		"tahun":                   riwayat_skp.Tahun,
	}

	if riwayat_skp.AtasanPenilaiTmtGolongan != nil {
		updatas["atasanPenilaiTmtGolongan"] = (riwayat_skp.AtasanPenilaiTmtGolongan).Format("02-01-2006")
	}

	if riwayat_skp.PenilaiTmtGolongan != nil {
		updatas["penilaiTmtGolongan"] = (riwayat_skp.PenilaiTmtGolongan).Format("02-01-2006")
	}

	//start update
	var singkronisasi model.Singkronisasi
	host := "bkn"
	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)

	pathBkn := "skp/save"

	if riwayat_skp.Tahun == 2021 {
		pathBkn = "skp/2021/save"

		updatas["inisiatifKerja"] = riwayat_skp.InisiatifKerja
		updatas["konversiNilai"] = riwayat_skp.KonversiNilai
		updatas["nilaiIntegrasi"] = riwayat_skp.NilaiIntegrasi
	}

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/" + pathBkn)

	//fmt.Println(updatas)

	//updatas = riwayat_skp

	out, err := json.Marshal(updatas)
	fmt.Println(string(out))

	if err != nil {
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewReader(out))

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

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	code := "0"
	if status_bkn["success"] == true {
		code = "1"
	}

	riwayat_update_bkn := &model.RiwayatUpdateBKN{
		Deskripsi: fmt.Sprintf("Riwayat SKP Tahun %s %s", riwayat_skp.Tahun, riwayat_skp.Nip),
		CreatedBy: CreatedBy,
		Code:      code,
		Message:   fmt.Sprint(status_bkn["message"]),
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create*/

	if status_bkn["success"] == true {
		mapData := status_bkn["mapData"]
		fmt.Println(mapData)
		//mapmapData := mapData.(map[string]string)

		db.Model(&model.RiwayatSkp{}).Debug().Where("nip = ? and tahun = ? and id =?", riwayat_skp.Nip, riwayat_skp.Tahun, riwayat_skp.ID).Update("idSync", mapData)
	}

	return
}
