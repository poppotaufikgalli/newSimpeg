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

func GetJabatanStr(searchString model.SearchJabatanStr) (jabatan_str []model.JabatanStr, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JabatanStr{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jabatan_str.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jabatan_str.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&jabatan_str)

	return
}

func FindJabatanStr(c echo.Context) error {
	var searchString model.SearchJabatanStr

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jabatan_str, result := GetJabatanStr(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_str,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJabatanStr(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_str model.JabatanStr
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_str)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_str); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jabatan_str.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jabatan_str)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_str,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJabatanStr(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_str model.JabatanStr
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_str)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_str); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jabatan_str.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JabatanStr{}).Where("id = ?", jabatan_str.Id).Updates(&jabatan_str)

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
		"data":       jabatan_str,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJabatanStr(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_str model.DeleteJabatanStr
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_str)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_str); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JabatanStr{}).Where("id = ?", jabatan_str.Id).Delete(&jabatan_str)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_str,
		"statucCode": http.StatusOK,
	})
}
