package app

import (
	_ "database/sql"
	"encoding/json"
	//"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	model "newSimpegAPI/model"
	"strings"
)

func GetStlud(searchString model.SearchStlud) (stlud []model.Stlud, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Stlud{})

	//kstlud
	if searchString.Kstlud > 0 {
		result = result.Where("master_stlud.kstlud = ?", searchString.Kstlud)
	}

	//nstlud
	if searchString.Nstlud != "" {
		nama := strings.TrimSpace(searchString.Nstlud)
		str := []string{"%", nama, "%"}
		result = result.Where("master_stlud.nstlud LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&stlud)

	return
}

func FindStlud(c echo.Context) error {
	var searchString model.SearchStlud

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	stlud, result := GetStlud(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       stlud,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateStlud(c echo.Context) error {
	db, _ := model.CreateCon()

	var stlud model.Stlud
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&stlud)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(stlud); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//stlud.CreatedBy = fmt.Sprint(c.Get("nip"))

	result := db.Create(&stlud)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       stlud,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateStlud(c echo.Context) error {
	db, _ := model.CreateCon()

	var stlud model.Stlud
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&stlud)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(stlud); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//stlud.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Stlud{}).Where("kstlud = ?", stlud.Kstlud).Updates(&stlud)

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
		"data":       stlud,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteStlud(c echo.Context) error {
	db, _ := model.CreateCon()

	var stlud model.DeleteStlud
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&stlud)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(stlud); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Stlud{}).Where("kstlud = ?", stlud.Kstlud).Delete(&stlud)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       stlud,
		"statucCode": http.StatusOK,
	})
}
