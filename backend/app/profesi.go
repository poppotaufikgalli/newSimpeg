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

func GetProfesi(searchString model.SearchProfesi) (profesi []model.Profesi, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Profesi{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_profesi.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_profesi.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&profesi)

	return
}

func FindProfesi(c echo.Context) error {
	var searchString model.SearchProfesi

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	profesi, result := GetProfesi(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       profesi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateProfesi(c echo.Context) error {
	db, _ := model.CreateCon()

	var profesi model.Profesi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&profesi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(profesi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	profesi.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&profesi)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       profesi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateProfesi(c echo.Context) error {
	db, _ := model.CreateCon()

	var profesi model.Profesi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&profesi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(profesi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	profesi.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Profesi{}).Where("id = ?", profesi.Id).Updates(&profesi)

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
		"data":       profesi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteProfesi(c echo.Context) error {
	db, _ := model.CreateCon()

	var profesi model.DeleteProfesi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&profesi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(profesi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Profesi{}).Where("id = ?", profesi.Id).Delete(&profesi)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       profesi,
		"statucCode": http.StatusOK,
	})
}
