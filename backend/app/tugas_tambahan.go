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

func GetTugasTambahan(searchString model.SearchTugasTambahan) (tugas_tambahan []model.TugasTambahan, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.TugasTambahan{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_tugas_tambahan.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_tugas_tambahan.nama LIKE ?", strings.Join(str, ""))
	}

	//singkatan
	if searchString.Singkatan != "" {
		nama := strings.TrimSpace(searchString.Singkatan)
		str := []string{"%", nama, "%"}
		result = result.Where("master_tugas_tambahan.singkatan LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&tugas_tambahan)

	return
}

func FindTugasTambahan(c echo.Context) error {
	var searchString model.SearchTugasTambahan

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	tugas_tambahan, result := GetTugasTambahan(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       tugas_tambahan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateTugasTambahan(c echo.Context) error {
	db, _ := model.CreateCon()

	var tugas_tambahan model.TugasTambahan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&tugas_tambahan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(tugas_tambahan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	tugas_tambahan.CreatedBy = fmt.Sprint(c.Get("nip"))

	result := db.Create(&tugas_tambahan)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       tugas_tambahan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateTugasTambahan(c echo.Context) error {
	db, _ := model.CreateCon()

	var tugas_tambahan model.TugasTambahan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&tugas_tambahan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(tugas_tambahan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	tugas_tambahan.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.TugasTambahan{}).Where("id = ?", tugas_tambahan.Id).Updates(&tugas_tambahan)

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
		"data":       tugas_tambahan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteTugasTambahan(c echo.Context) error {
	db, _ := model.CreateCon()

	var tugas_tambahan model.DeleteTugasTambahan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&tugas_tambahan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(tugas_tambahan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.TugasTambahan{}).Where("id = ?", tugas_tambahan.Id).Delete(&tugas_tambahan)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       tugas_tambahan,
		"statucCode": http.StatusOK,
	})
}
