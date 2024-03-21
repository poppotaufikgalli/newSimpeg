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

func GetPekerjaan(searchString model.SearchPekerjaan) (pekerjaan []model.Pekerjaan, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Pekerjaan{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_pekerjaan.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_pekerjaan.nama LIKE ?", strings.Join(str, ""))
	}

	//ref simpeg
	if len(searchString.Status) > 0 {
		result = result.Where("master_pekerjaan.status IN (?)", searchString.Status)
	}

	result = result.Find(&pekerjaan)

	return
}

func FindPekerjaan(c echo.Context) error {
	var searchString model.SearchPekerjaan

	//json.NewDecoder(c.Request().Body).Decode(&searchString)
	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		searchString.Status = []int{1}
	}

	pekerjaan, result := GetPekerjaan(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pekerjaan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreatePekerjaan(c echo.Context) error {
	db, _ := model.CreateCon()

	var pekerjaan model.Pekerjaan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pekerjaan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pekerjaan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	pekerjaan.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&pekerjaan)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pekerjaan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdatePekerjaan(c echo.Context) error {
	db, _ := model.CreateCon()

	var pekerjaan model.Pekerjaan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pekerjaan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pekerjaan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	pekerjaan.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Pekerjaan{}).Where("id = ?", pekerjaan.Id).Updates(&pekerjaan)

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
		"data":       pekerjaan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeletePekerjaan(c echo.Context) error {
	db, _ := model.CreateCon()

	var pekerjaan model.DeletePekerjaan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pekerjaan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pekerjaan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Pekerjaan{}).Where("id = ?", pekerjaan.Id).Delete(&pekerjaan)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pekerjaan,
		"statucCode": http.StatusOK,
	})
}
