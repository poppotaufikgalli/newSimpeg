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

func GetJenisKeluarga(searchString model.SearchJenisKeluarga) (jenis_keluarga []model.JenisKeluarga, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisKeluarga{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_keluarga.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_keluarga.nama LIKE ?", strings.Join(str, ""))
	}

	//ref simpeg
	if len(searchString.Status) > 0 {
		result = result.Where("master_jenis_keluarga.status IN (?)", searchString.Status)
	}

	result = result.Find(&jenis_keluarga)

	return
}

func FindJenisKeluarga(c echo.Context) error {
	var searchString model.SearchJenisKeluarga

	//json.NewDecoder(c.Request().Body).Decode(&searchString)
	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		searchString.Status = []int{1}
	}

	jenis_keluarga, result := GetJenisKeluarga(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_keluarga,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetJenisKeluargaByJkeluarga(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_keluarga model.JenisKeluarga
	id := c.Param("id")

	result := db.Model(&model.JenisKeluarga{}).Where("id=?", id).Scan(&jenis_keluarga)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_keluarga,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateJenisKeluarga(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_keluarga model.JenisKeluarga
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_keluarga)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_keluarga); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_keluarga.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_keluarga)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_keluarga,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisKeluarga(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_keluarga model.JenisKeluarga
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_keluarga)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_keluarga); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_keluarga.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisKeluarga{}).Where("id = ?", jenis_keluarga.Id).Updates(&jenis_keluarga)

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
		"data":       jenis_keluarga,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisKeluarga(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_keluarga model.DeleteJenisKeluarga
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_keluarga)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_keluarga); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisKeluarga{}).Where("id = ?", jenis_keluarga.Id).Delete(&jenis_keluarga)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_keluarga,
		"statucCode": http.StatusOK,
	})
}
