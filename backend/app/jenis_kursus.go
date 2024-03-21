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

func GetJenisKursus(searchString model.SearchJenisKursus) (jenis_kursus []model.JenisKursus, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisKursus{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_kursus.id = ?", searchString.Id)
	}

	//cepat kode
	if searchString.CepatKode != "" {
		result = result.Where("master_jenis_kursus.cepat_kode = ?", searchString.CepatKode)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_kursus.nama LIKE ?", strings.Join(str, ""))
	}

	//ref Simpeg
	if searchString.RefSimpeg != "" {
		result = result.Where("master_jenis_kursus.ref_simpeg = ?", searchString.RefSimpeg)
	}

	result = result.Find(&jenis_kursus)

	return
}

func FindJenisKursus(c echo.Context) error {
	var searchString model.SearchJenisKursus

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jenis_kursus, result := GetJenisKursus(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_kursus,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisKursus(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_kursus model.JenisKursus
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_kursus)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_kursus); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_kursus.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_kursus)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_kursus,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisKursus(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_kursus model.JenisKursus
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_kursus)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_kursus); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_kursus.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisKursus{}).Where("id = ?", jenis_kursus.Id).Updates(&jenis_kursus)

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
		"data":       jenis_kursus,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisKursus(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_kursus model.DeleteJenisKursus
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_kursus)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_kursus); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisKursus{}).Where("id = ?", jenis_kursus.Id).Delete(&jenis_kursus)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_kursus,
		"statucCode": http.StatusOK,
	})
}
