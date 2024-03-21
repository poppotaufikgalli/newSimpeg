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

func GetPejabat(searchString model.SearchPejabat) (pejabat []model.Pejabat, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Pejabat{})

	//kpej
	if searchString.Kpej != "" {
		result = result.Where("master_pejabat.kpej = ?", searchString.Kpej)
	}

	//npej
	if searchString.Npej != "" {
		nama := strings.TrimSpace(searchString.Npej)
		str := []string{"%", nama, "%"}
		result = result.Where("master_pejabat.npej LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&pejabat)

	return
}

func FindPejabat(c echo.Context) error {
	var searchString model.SearchPejabat

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	pejabat, result := GetPejabat(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pejabat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreatePejabat(c echo.Context) error {
	db, _ := model.CreateCon()

	var pejabat model.Pejabat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pejabat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pejabat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	pejabat.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&pejabat)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pejabat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdatePejabat(c echo.Context) error {
	db, _ := model.CreateCon()

	var pejabat model.Pejabat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pejabat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pejabat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	pejabat.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Pejabat{}).Where("kpej = ?", pejabat.Kpej).Updates(&pejabat)

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
		"data":       pejabat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeletePejabat(c echo.Context) error {
	db, _ := model.CreateCon()

	var pejabat model.DeletePejabat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pejabat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pejabat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Pejabat{}).Where("kpej = ?", pejabat.Kpej).Delete(&pejabat)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pejabat,
		"statucCode": http.StatusOK,
	})
}
