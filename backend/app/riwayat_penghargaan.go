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
	"strings"
)

func GetRiwayatPenghargaan(searchString model.SearchRiwayatPenghargaan) (riwayat_penghargaan []model.RiwayatPenghargaan, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatPenghargaan{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_penghargaan.nip = ?", searchString.Nip)
	}

	//id jenis penghargaan
	if searchString.IdJenisPenghargaan != "" {
		result = result.Where("master_riwayat_penghargaan.id_jenis_penghargaan = ?", searchString.IdJenisPenghargaan)
	}

	//nsk
	if searchString.Nsk != "" {
		nama := strings.TrimSpace(searchString.Nsk)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_penghargaan.nsk LIKE ?", strings.Join(str, ""))
	}

	//Thn
	if searchString.Thn > 0 {
		result = result.Where("master_riwayat_penghargaan.thn = ?", searchString.Thn)
	}

	result = result.Preload("JenisPenghargaan").Order("thn desc").Find(&riwayat_penghargaan)

	return

}

func FindRiwayatPenghargaan(c echo.Context) error {
	var searchString model.SearchRiwayatPenghargaan
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_penghargaan, result := GetRiwayatPenghargaan(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_penghargaan,
		"statusCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetRiwayatPenghargaanByNipIdNsk(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_penghargaan model.RiwayatPenghargaan

	nip := c.Param("nip")
	nbintang := c.Param("nbintang")
	nsk := c.Param("nsk")

	result := db.Model(&model.RiwayatPenghargaan{}).Where("nip =? and nbintang = ? and nsk = ?", nip, nbintang, nsk).Scan(&riwayat_penghargaan)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_penghargaan,
		"statusCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateRiwayatPenghargaan(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_penghargaan model.RiwayatPenghargaan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_penghargaan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statusCode": http.StatusNotImplemented,
			"message":    "gagal Decode JSON",
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_penghargaan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"message":    "gagal Validasi JSON",
			"errors":     err.Error(),
		})
	}

	riwayat_penghargaan.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_penghargaan)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"message":    "gagal menambahkan ke DB",
			"errors":     result.Error.Error(),
		})
	} else {
		if c.Get("gid").(int) == 1 {
			status_bkn, err := sendRiwayatPenghargaanBKN(riwayat_penghargaan, c.Get("nip").(string))
			//fmt.Println(status_bkn)

			if err != nil {
				return c.JSON(http.StatusNotImplemented, map[string]interface{}{
					"statusCode": http.StatusNotImplemented,
					"errors":     err.Error(),
				})
			}
			if status_bkn["success"] != true {
				return c.JSON(http.StatusNotImplemented, map[string]string{
					"title":       "Update Data Penghargaan BKN Gagal",
					"description": fmt.Sprintf("%s", status_bkn["message"]),
				})
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_penghargaan,
		"statusCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatPenghargaan(c echo.Context) error {
	db, _ := model.CreateCon()

	//var riwayat_penghargaan model.RiwayatPenghargaan
	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatPenghargaan `json:"refdata"`
		Data model.RiwayatPenghargaan       `json:"data"`
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
	result := db.Model(&model.RiwayatPenghargaan{}).Where("nip = ? and nbintang = ? and nsk = ?", reqData.Ref.Nip, reqData.Ref.Nbintang, reqData.Ref.Nsk).Updates(&reqData.Data)

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
	} else {
		if c.Get("gid").(int) == 1 {
			status_bkn, err := sendRiwayatPenghargaanBKN(reqData.Data, c.Get("nip").(string))
			//fmt.Println(status_bkn)

			if err != nil {
				return c.JSON(http.StatusNotImplemented, map[string]interface{}{
					"statusCode": http.StatusNotImplemented,
					"errors":     err.Error(),
				})
			}
			if status_bkn["success"] != true {
				return c.JSON(http.StatusNotImplemented, map[string]string{
					"title":       "Update Data Penghargaan BKN Gagal",
					"description": fmt.Sprintf("%s", status_bkn["message"]),
				})
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       reqData.Data,
		"statusCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatPenghargaan(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_penghargaan model.DeleteRiwayatPenghargaan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_penghargaan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statusCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_penghargaan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatPenghargaan{}).Where("nip = ? and nbintang = ? and nsk = ?", riwayat_penghargaan.Nip, riwayat_penghargaan.Nbintang, riwayat_penghargaan.Nsk).Delete(&riwayat_penghargaan)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statusCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	} else {
		if c.Get("gid").(int) == 1 {
			resultBkn, err := DeleteSingkronisasiDataBknById("penghargaan", riwayat_penghargaan.IdSync, c.Get("nip").(string))

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
		"data":       riwayat_penghargaan,
		"statusCode": http.StatusOK,
	})
}

func sendRiwayatPenghargaanBKN(riwayat_penghargaan model.RiwayatPenghargaan, CreatedBy string) (status_bkn map[string]interface{}, err error) {
	db, _ := model.CreateCon()

	var pegawai model.Pegawai
	db.Select("no_ref_bkn").First(&pegawai, "nip = ?", riwayat_penghargaan.Nip)

	Tsk := (riwayat_penghargaan.Tsk).Format("02-01-2006")

	updatas := map[string]interface{}{
		"hargaId":    *riwayat_penghargaan.IdJenisPenghargaan,
		"skNomor":    *riwayat_penghargaan.Nsk,
		"pnsOrangId": *pegawai.NoRefBkn,
		"tahun":      *riwayat_penghargaan.Thn,
		"skDate":     Tsk,
	}

	if riwayat_penghargaan.IdSync != nil {
		updatas["id"] = *riwayat_penghargaan.IdSync
	}

	status_bkn, err = PostSinkronisasiGetDataBKN(updatas, "penghargaan", CreatedBy)

	if err != nil {
		return
	}

	riwayat_update_bkn := &model.RiwayatUpdateBKN{
		Deskripsi: fmt.Sprintf("Riwayat Penghargaan %s, No SK %s", riwayat_penghargaan.Nip, *riwayat_penghargaan.Nsk),
		CreatedBy: CreatedBy,
		Code:      fmt.Sprint(status_bkn["code"]),
		Message:   fmt.Sprint(status_bkn["message"]),
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create*/

	if status_bkn["success"] == true || status_bkn["code"] == 1 {
		mapData := status_bkn["mapData"]
		//fmt.Println(mapData)
		mapmapData := mapData.(map[string]interface{})
		db.Model(&model.RiwayatPenghargaan{}).Debug().Where("nip = ? and nbintang = ? and nsk = ?", riwayat_penghargaan.Nip, riwayat_penghargaan.Nbintang, riwayat_penghargaan.Nsk).Update("idSync", mapmapData["rwPenghargaanId"])
	}

	return
}
