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

func GetJenisInstansi(searchString model.SearchJenisInstansi) (jenis_instansi []model.JenisInstansi, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisInstansi{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_instansi.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_instansi.nama LIKE ?", strings.Join(str, ""))
	}

	//jenis
	if len(searchString.Jenis) > 0 {
		result = result.Where("master_jenis_instansi.jenis IN (?)", searchString.Jenis)
	}

	result = result.Find(&jenis_instansi)

	return
}

func FindJenisInstansi(c echo.Context) error {
	var searchString model.SearchJenisInstansi

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jenis_instansi, result := GetJenisInstansi(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_instansi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisInstansi(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_instansi model.JenisInstansi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_instansi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_instansi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_instansi.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_instansi)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_instansi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisInstansi(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_instansi model.JenisInstansi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_instansi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_instansi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_instansi.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisInstansi{}).Where("id = ?", jenis_instansi.Id).Updates(&jenis_instansi)

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
		"data":       jenis_instansi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisInstansi(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_instansi model.DeleteJenisInstansi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_instansi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_instansi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisInstansi{}).Where("id = ?", jenis_instansi.Id).Delete(&jenis_instansi)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_instansi,
		"statucCode": http.StatusOK,
	})
}
