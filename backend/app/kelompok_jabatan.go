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

func GetKelompokJabatan(searchString model.SearchKelompokJabatan) (kelompok_jabatan []model.KelompokJabatan, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.KelompokJabatan{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_kelompok_jabatan.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_kelompok_jabatan.nama LIKE ?", strings.Join(str, ""))
	}

	//jenis jabatan id
	if searchString.JenisJabatanId != "" {
		result = result.Where("master_kelompok_jabatan.jenis_jabatan_id = ?", searchString.JenisJabatanId)
	}

	//rumpun jabatan id
	if searchString.RumpunJabatanId != "" {
		result = result.Where("master_kelompok_jabatan.rumpun_jabatan_id = ?", searchString.RumpunJabatanId)
	}

	//pembina id
	if searchString.PembinaId != "" {
		result = result.Where("master_kelompok_jabatan.pembina_id = ?", searchString.PembinaId)
	}

	result = result.Find(&kelompok_jabatan)

	return
}

func FindKelompokJabatan(c echo.Context) error {
	var searchString model.SearchKelompokJabatan

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	kelompok_jabatan, result := GetKelompokJabatan(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       kelompok_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateKelompokJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var kelompok_jabatan model.KelompokJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&kelompok_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(kelompok_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	kelompok_jabatan.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&kelompok_jabatan)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       kelompok_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateKelompokJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var kelompok_jabatan model.KelompokJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&kelompok_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(kelompok_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	kelompok_jabatan.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.KelompokJabatan{}).Where("id = ?", kelompok_jabatan.Id).Updates(&kelompok_jabatan)

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
		"data":       kelompok_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteKelompokJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var kelompok_jabatan model.DeleteKelompokJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&kelompok_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(kelompok_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.KelompokJabatan{}).Where("id = ?", kelompok_jabatan.Id).Delete(&kelompok_jabatan)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       kelompok_jabatan,
		"statucCode": http.StatusOK,
	})
}
