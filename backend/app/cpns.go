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

func GetCpns(searchString model.SearchCpns) (cpns []model.Cpns, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Cpns{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_cpns.nip = ?", searchString.Nip)
	}

	//sk cpns
	if searchString.Skcpns != "" {
		skcpns := strings.TrimSpace(searchString.Skcpns)
		str := []string{"%", skcpns, "%"}
		result = result.Where("master_cpns.skcpns LIKE ?", strings.Join(str, ""))
	}

	//TMT CPNS
	if searchString.Tmtcpns.IsZero() != true {
		result = result.Where("master_cpns.tmtcpns = ?", searchString.Tmtcpns)
	}

	result = result.Find(&cpns)

	return
}

func FindCpns(c echo.Context) error {
	var searchString model.SearchCpns

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	cpns, result := GetCpns(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       cpns,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateCpns(c echo.Context) error {
	db, _ := model.CreateCon()

	var cpns model.Cpns
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&cpns)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(cpns); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	cpns.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&cpns)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       cpns,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateCpns(c echo.Context) error {
	db, _ := model.CreateCon()

	var cpns model.Cpns
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&cpns)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(cpns); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	cpns.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Cpns{}).Where("nip = ?", cpns.Nip).Updates(&cpns)

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
		"data":       cpns,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteCpns(c echo.Context) error {
	db, _ := model.CreateCon()

	var cpns model.DeleteCpns
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&cpns)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(cpns); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Cpns{}).Where("nip = ?", cpns.Nip).Delete(&cpns)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       cpns,
		"statucCode": http.StatusOK,
	})
}
