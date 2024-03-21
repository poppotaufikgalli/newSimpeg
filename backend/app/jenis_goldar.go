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

func GetJenisGoldar(searchString model.SearchJenisGoldar) (jenis_goldar []model.JenisGoldar, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisGoldar{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_goldar.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_goldar.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&jenis_goldar)

	return
}

func FindJenisGoldar(c echo.Context) error {
	var searchString model.SearchJenisGoldar

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jenis_goldar, result := GetJenisGoldar(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_goldar,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisGoldar(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_goldar model.JenisGoldar
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_goldar)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_goldar); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_goldar.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_goldar)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_goldar,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisGoldar(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_goldar model.JenisGoldar
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_goldar)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_goldar); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_goldar.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisGoldar{}).Where("id = ?", jenis_goldar.Id).Updates(&jenis_goldar)

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
		"data":       jenis_goldar,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisGoldar(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_goldar model.DeleteJenisGoldar
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_goldar)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_goldar); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisGoldar{}).Where("id = ?", jenis_goldar.Id).Delete(&jenis_goldar)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_goldar,
		"statucCode": http.StatusOK,
	})
}
