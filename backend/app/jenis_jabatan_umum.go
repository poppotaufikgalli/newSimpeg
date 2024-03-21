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

func GetJenisJabatanUmum(searchString model.SearchJenisJabatanUmum) (jenis_jabatan_umum []model.JenisJabatanUmum, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisJabatanUmum{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_jabatan_umum.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_jabatan_umum.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&jenis_jabatan_umum)

	return
}

func FindJenisJabatanUmum(c echo.Context) error {
	var searchString model.SearchJenisJabatanUmum

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jenis_jabatan_umum, result := GetJenisJabatanUmum(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_jabatan_umum,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisJabatanUmum(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_jabatan_umum model.JenisJabatanUmum
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_jabatan_umum)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_jabatan_umum); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_jabatan_umum.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_jabatan_umum)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_jabatan_umum,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisJabatanUmum(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_jabatan_umum model.JenisJabatanUmum
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_jabatan_umum)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_jabatan_umum); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_jabatan_umum.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisJabatanUmum{}).Where("id = ?", jenis_jabatan_umum.Id).Updates(&jenis_jabatan_umum)

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
		"data":       jenis_jabatan_umum,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisJabatanUmum(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_jabatan_umum model.DeleteJenisJabatanUmum
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_jabatan_umum)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_jabatan_umum); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisJabatanUmum{}).Where("id = ?", jenis_jabatan_umum.Id).Delete(&jenis_jabatan_umum)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_jabatan_umum,
		"statucCode": http.StatusOK,
	})
}
