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

func GetUnorBkn(searchString model.SearchUnorBkn) (unor_bkn []model.UnorBkn, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.UnorBkn{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_unor_bkn.id = ?", searchString.Id)
	}

	//nama unor
	if searchString.NamaUnor != "" {
		nama := strings.TrimSpace(searchString.NamaUnor)
		str := []string{"%", nama, "%"}
		result = result.Where("master_unor_bkn.NamaUnor LIKE ?", strings.Join(str, ""))
	}

	//nama jabatan
	if searchString.NamaJabatan != "" {
		nama := strings.TrimSpace(searchString.NamaJabatan)
		str := []string{"%", nama, "%"}
		result = result.Where("master_unor_bkn.NamaJabatan LIKE ?", strings.Join(str, ""))
	}

	//cepat kode
	if searchString.CepatKode != "" {
		result = result.Where("master_unor_bkn.CepatKode = ?", searchString.CepatKode)
	}

	result = result.Find(&unor_bkn)

	return
}

func FindUnorBkn(c echo.Context) error {
	var searchString model.SearchUnorBkn

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	unor_bkn, result := GetUnorBkn(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       unor_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateUnorBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var unor_bkn model.UnorBkn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&unor_bkn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(unor_bkn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	unor_bkn.CreatedBy = fmt.Sprint(c.Get("nip"))

	result := db.Create(&unor_bkn)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       unor_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateUnorBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var unor_bkn model.UnorBkn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&unor_bkn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(unor_bkn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	unor_bkn.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.UnorBkn{}).Where("id = ?", unor_bkn.Id).Updates(&unor_bkn)

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
		"data":       unor_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteUnorBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var unor_bkn model.DeleteUnorBkn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&unor_bkn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(unor_bkn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.UnorBkn{}).Where("id = ?", unor_bkn.Id).Delete(&unor_bkn)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       unor_bkn,
		"statucCode": http.StatusOK,
	})
}
