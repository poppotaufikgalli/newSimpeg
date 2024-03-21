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

func GetJenisJabatan(searchString model.SearchJenisJabatan) (jenis_jabatan []model.JenisJabatan, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisJabatan{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_jabatan.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_jabatan.nama LIKE ?", strings.Join(str, ""))
	}

	//status
	if len(searchString.Status) > 0 {
		result = result.Where("master_jenis_jabatan.status IN (?)", searchString.Status)
	}

	result = result.Find(&jenis_jabatan)

	return
}

func FindJenisJabatan(c echo.Context) error {
	var searchString model.SearchJenisJabatan
	//var jenis_jabatan []model.JenisJabatan

	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		searchString.Status = []int{1}
	}

	jenis_jabatan, result := GetJenisJabatan(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_jabatan model.JenisJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_jabatan.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_jabatan)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_jabatan model.JenisJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_jabatan.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisJabatan{}).Where("id = ?", jenis_jabatan.Id).Updates(&jenis_jabatan)

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
		"data":       jenis_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_jabatan model.DeleteJenisJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisJabatan{}).Where("id = ?", jenis_jabatan.Id).Delete(&jenis_jabatan)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_jabatan,
		"statucCode": http.StatusOK,
	})
}
