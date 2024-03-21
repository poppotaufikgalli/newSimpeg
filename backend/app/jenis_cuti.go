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

func GetJenisCuti(searchString model.SearchJenisCuti) (jenis_cuti []model.JenisCuti, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisCuti{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_cuti.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_cuti.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&jenis_cuti)

	return
}

func FindJenisCuti(c echo.Context) error {
	var searchString model.SearchJenisCuti
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jenis_cuti, result := GetJenisCuti(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_cuti,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisCuti(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_cuti model.JenisCuti
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_cuti)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_cuti); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_cuti.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_cuti)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_cuti,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisCuti(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_cuti model.JenisCuti
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_cuti)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_cuti); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_cuti.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisCuti{}).Where("id = ?", jenis_cuti.Id).Updates(&jenis_cuti)

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
		"data":       jenis_cuti,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisCuti(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_cuti model.DeleteJenisCuti
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_cuti)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_cuti); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisCuti{}).Where("id = ?", jenis_cuti.Id).Delete(&jenis_cuti)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_cuti,
		"statucCode": http.StatusOK,
	})
}
