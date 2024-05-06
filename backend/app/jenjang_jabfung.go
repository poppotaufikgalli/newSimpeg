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

func GetJenjangJabfung(searchString model.SearchJenjangJabfung) (jenjang_jabfung []model.JenjangJabfung, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenjangJabfung{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenjang_jabfung.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenjang_jabfung.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Order("urut").Find(&jenjang_jabfung)

	return
}

func FindJenjangJabfung(c echo.Context) error {
	var searchString model.SearchJenjangJabfung

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jenjang_jabfung, result := GetJenjangJabfung(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenjang_jabfung,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenjangJabfung(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenjang_jabfung model.JenjangJabfung
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenjang_jabfung)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenjang_jabfung); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenjang_jabfung.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenjang_jabfung)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenjang_jabfung,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenjangJabfung(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenjang_jabfung model.JenjangJabfung
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenjang_jabfung)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenjang_jabfung); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenjang_jabfung.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenjangJabfung{}).Where("id = ?", jenjang_jabfung.Id).Updates(&jenjang_jabfung)

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
		"data":       jenjang_jabfung,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenjangJabfung(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenjang_jabfung model.DeleteJenjangJabfung
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenjang_jabfung)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenjang_jabfung); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenjangJabfung{}).Where("id = ?", jenjang_jabfung.Id).Delete(&jenjang_jabfung)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenjang_jabfung,
		"statucCode": http.StatusOK,
	})
}
