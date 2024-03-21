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

func GetInstansi(searchString model.SearchInstansi) (instansi []model.Instansi, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Instansi{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_instansi.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_instansi.nama LIKE ?", strings.Join(str, ""))
	}

	//jenis
	if len(searchString.Jenis) > 0 {
		result = result.Where("master_instansi.jenis = ?", searchString.Jenis)
	}

	//kode cepat
	if searchString.CepatKode != "" {
		result = result.Where("master_instansi.cepat_kode = ?", searchString.CepatKode)
	}

	//jenis
	if len(searchString.JenisInstansiId) > 0 {
		result = result.Where("master_instansi.jenis_instansi_id = ?", searchString.JenisInstansiId)
	}

	result = result.Find(&instansi)

	return
}

func FindInstansi(c echo.Context) error {
	var searchString model.SearchInstansi

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	instansi, result := GetInstansi(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       instansi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateInstansi(c echo.Context) error {
	db, _ := model.CreateCon()

	var instansi model.Instansi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&instansi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(instansi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	instansi.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&instansi)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       instansi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateInstansi(c echo.Context) error {
	db, _ := model.CreateCon()

	var instansi model.Instansi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&instansi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(instansi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	instansi.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Instansi{}).Where("id = ?", instansi.Id).Updates(&instansi)

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
		"data":       instansi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteInstansi(c echo.Context) error {
	db, _ := model.CreateCon()

	var instansi model.DeleteInstansi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&instansi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(instansi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Instansi{}).Where("id = ?", instansi.Id).Delete(&instansi)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       instansi,
		"statucCode": http.StatusOK,
	})
}
