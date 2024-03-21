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

func GetReferensiJabatan(searchString model.SearchReferensiJabatan) (referensi_jabatan []model.ReferensiJabatan, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.ReferensiJabatan{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_referensi_jabatan.id = ?", searchString.Id)
	}

	//jenis jabatan id
	if searchString.JenisJabatanId > 0 {
		result = result.Where("master_referensi_jabatan.jenis_jabatan_id = ?", searchString.JenisJabatanId)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_referensi_jabatan.nama LIKE ?", strings.Join(str, ""))
	}

	//bup usia
	if searchString.BupUsia > 0 {
		result = result.Where("master_referensi_jabatan.bup_usia = ?", searchString.BupUsia)
	}

	//kel jabatan id
	if len(searchString.KelJabatanId) > 0 {
		result = result.Where("master_referensi_jabatan.kel_jabatan_id IN (?)", searchString.KelJabatanId)
	}

	//status
	if len(searchString.Status) > 0 {
		result = result.Where("master_referensi_jabatan.status IN (?)", searchString.Status)
	}

	//cepat kode
	if searchString.CepatKode != "" {
		result = result.Where("master_referensi_jabatan.cepat_kode = ?", searchString.CepatKode)
	}

	//ref Simpeg
	if searchString.RefSimpeg != "" {
		result = result.Where("master_referensi_jabatan.ref_simpeg = ?", searchString.RefSimpeg)
	}

	result = result.Find(&referensi_jabatan)

	return
}

func FindReferensiJabatan(c echo.Context) error {
	var searchString model.SearchReferensiJabatan

	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		searchString.Status = []int{1}
	}

	referensi_jabatan, result := GetReferensiJabatan(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       referensi_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateReferensiJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var referensi_jabatan model.ReferensiJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&referensi_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(referensi_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	referensi_jabatan.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&referensi_jabatan)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       referensi_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateReferensiJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var referensi_jabatan model.ReferensiJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&referensi_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(referensi_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	referensi_jabatan.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.ReferensiJabatan{}).Where("id = ? and jenis_jabatan_id = ?", referensi_jabatan.Id, referensi_jabatan.JenisJabatanId).Updates(&referensi_jabatan)

	//fmt.Println(result.RowsAffected)

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
		"data":       referensi_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteReferensiJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var referensi_jabatan model.DeleteReferensiJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&referensi_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(referensi_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.ReferensiJabatan{}).Where("id = ? and jenis_jabatan_id = ?", referensi_jabatan.Id, referensi_jabatan.JenisJabatanId).Delete(&referensi_jabatan)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       referensi_jabatan,
		"statucCode": http.StatusOK,
	})
}
