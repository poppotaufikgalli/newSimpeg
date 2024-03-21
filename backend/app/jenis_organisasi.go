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

func GetJenisOrganisasi(searchString model.SearchJenisOrganisasi) (jenis_organisasi []model.JenisOrganisasi, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisOrganisasi{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_organisasi.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_organisasi.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&jenis_organisasi)

	return
}

func FindJenisOrganisasi(c echo.Context) error {
	var searchString model.SearchJenisOrganisasi

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jenis_organisasi, result := GetJenisOrganisasi(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_organisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisOrganisasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_organisasi model.JenisOrganisasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_organisasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_organisasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_organisasi.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_organisasi)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_organisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisOrganisasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_organisasi model.JenisOrganisasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_organisasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_organisasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_organisasi.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisOrganisasi{}).Where("id = ?", jenis_organisasi.Id).Updates(&jenis_organisasi)

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
		"data":       jenis_organisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisOrganisasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_organisasi model.DeleteJenisOrganisasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_organisasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_organisasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisOrganisasi{}).Where("id = ?", jenis_organisasi.Id).Delete(&jenis_organisasi)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_organisasi,
		"statucCode": http.StatusOK,
	})
}
