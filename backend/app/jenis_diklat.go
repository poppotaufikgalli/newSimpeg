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

func GetJenisDiklat(searchString model.SearchJenisDiklat) (jenis_diklat []model.JenisDiklat, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisDiklat{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_diklat.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_diklat.nama LIKE ?", strings.Join(str, ""))
	}

	//status
	if len(searchString.Status) > 0 {
		result = result.Where("master_jenis_diklat.status IN (?)", searchString.Status)
	}

	result = result.Find(&jenis_diklat)

	return
}

func FindJenisDiklat(c echo.Context) error {
	var searchString model.SearchJenisDiklat

	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		searchString.Status = []int{1}
	}

	jenis_diklat, result := GetJenisDiklat(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_diklat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisDiklat(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_diklat model.JenisDiklat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_diklat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_diklat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_diklat.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_diklat)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_diklat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisDiklat(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_diklat model.JenisDiklat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_diklat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_diklat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_diklat.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisDiklat{}).Where("id = ?", jenis_diklat.Id).Updates(&jenis_diklat)

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
		"data":       jenis_diklat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisDiklat(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_diklat model.DeleteJenisDiklat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_diklat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_diklat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisDiklat{}).Where("id = ?", jenis_diklat.Id).Delete(&jenis_diklat)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_diklat,
		"statucCode": http.StatusOK,
	})
}
