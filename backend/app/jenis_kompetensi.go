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

func GetJenisKompetensi(searchString model.SearchJenisKompetensi) (jenis_kompetensi []model.JenisKompetensi, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisKompetensi{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_kompetensi.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_kompetensi.nama LIKE ?", strings.Join(str, ""))
	}

	//nama id
	if searchString.NamaId != "" {
		nama := strings.TrimSpace(searchString.NamaId)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_kompetensi.nama_id LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&jenis_kompetensi)

	return
}

func FindJenisKompetensi(c echo.Context) error {
	var searchString model.SearchJenisKompetensi

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jenis_kompetensi, result := GetJenisKompetensi(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_kompetensi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisKompetensi(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_kompetensi model.JenisKompetensi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_kompetensi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_kompetensi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_kompetensi.CreatedBy = fmt.Sprint(c.Get("nip"))

	result := db.Create(&jenis_kompetensi)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_kompetensi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisKompetensi(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_kompetensi model.JenisKompetensi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_kompetensi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_kompetensi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_kompetensi.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisKompetensi{}).Where("id = ?", jenis_kompetensi.Id).Updates(&jenis_kompetensi)

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
		"data":       jenis_kompetensi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisKompetensi(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_kompetensi model.DeleteJenisKompetensi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_kompetensi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_kompetensi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisKompetensi{}).Where("id = ?", jenis_kompetensi.Id).Delete(&jenis_kompetensi)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_kompetensi,
		"statucCode": http.StatusOK,
	})
}
