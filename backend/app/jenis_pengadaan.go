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

func GetJenisPengadaan(searchString model.SearchJenisPengadaan) (jenis_pengadaan []model.JenisPengadaan, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisPengadaan{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_pengadaan.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_pengadaan.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&jenis_pengadaan)

	return
}

func FindJenisPengadaan(c echo.Context) error {
	var searchString model.SearchJenisPengadaan

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jenis_pengadaan, result := GetJenisPengadaan(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_pengadaan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisPengadaan(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_pengadaan model.JenisPengadaan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_pengadaan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_pengadaan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_pengadaan.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_pengadaan)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_pengadaan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisPengadaan(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_pengadaan model.JenisPengadaan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_pengadaan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_pengadaan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_pengadaan.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisPengadaan{}).Where("id = ?", jenis_pengadaan.Id).Updates(&jenis_pengadaan)

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
		"data":       jenis_pengadaan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisPengadaan(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_pengadaan model.DeleteJenisPengadaan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_pengadaan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_pengadaan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisPengadaan{}).Where("id = ?", jenis_pengadaan.Id).Delete(&jenis_pengadaan)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_pengadaan,
		"statucCode": http.StatusOK,
	})
}
