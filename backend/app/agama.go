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

func GetAgama(searchString model.SearchAgama) (agama []model.Agama, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Agama{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_agama.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_agama.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&agama)

	return
}

func FindAgama(c echo.Context) error {
	var searchString model.SearchAgama

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	agama, result := GetAgama(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       agama,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateAgama(c echo.Context) error {
	db, _ := model.CreateCon()

	var agama model.Agama
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&agama)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(agama); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	agama.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&agama)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       agama,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateAgama(c echo.Context) error {
	db, _ := model.CreateCon()

	var agama model.Agama
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&agama)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(agama); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	agama.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Agama{}).Where("id = ?", agama.Id).Updates(&agama)

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
		"data":       agama,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteAgama(c echo.Context) error {
	db, _ := model.CreateCon()

	var agama model.DeleteAgama
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&agama)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(agama); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Agama{}).Where("id = ?", agama.Id).Delete(&agama)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       agama,
		"statucCode": http.StatusOK,
	})
}
