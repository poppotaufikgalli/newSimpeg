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

func GetDiklat(searchString model.SearchDiklat) (diklat []model.Diklat, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Diklat{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_diklat.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_diklat.nama LIKE ?", strings.Join(str, ""))
	}

	//id_jenis_jabatan
	if len(searchString.Jdiklat) > 0 {
		result = result.Where("master_diklat.jdiklat IN (?)", searchString.Jdiklat)
	}

	//status
	if len(searchString.Status) > 0 {
		result = result.Where("master_diklat.status IN (?)", searchString.Status)
	}

	result = result.Find(&diklat)

	return
}

func FindDiklat(c echo.Context) error {
	var searchString model.SearchDiklat

	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		searchString.Status = []int{1}
	}

	diklat, result := GetDiklat(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       diklat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateDiklat(c echo.Context) error {
	db, _ := model.CreateCon()

	var diklat model.Diklat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&diklat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(diklat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	diklat.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&diklat)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       diklat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateDiklat(c echo.Context) error {
	db, _ := model.CreateCon()

	var diklat model.Diklat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&diklat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(diklat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	diklat.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Diklat{}).Where("id = ? and jdiklat = ?", diklat.Id, diklat.Jdiklat).Updates(&diklat)

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
		"data":       diklat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteDiklat(c echo.Context) error {
	db, _ := model.CreateCon()

	var diklat model.DeleteDiklat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&diklat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(diklat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Diklat{}).Where("id = ? & jdiklat = ?", diklat.Id, diklat.Jdiklat).Delete(&diklat)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       diklat,
		"statucCode": http.StatusOK,
	})
}
