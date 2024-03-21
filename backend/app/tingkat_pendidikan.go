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

func GetTingkatPendidikan(searchString model.SearchTingkatPendidikan) (tingkat_pendidikan []model.TingkatPendidikan, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.TingkatPendidikan{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_tingkat_pendidikan.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_tingkat_pendidikan.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&tingkat_pendidikan)

	return
}

func FindTingkatPendidikan(c echo.Context) error {
	var searchString model.SearchTingkatPendidikan

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	tingkat_pendidikan, result := GetTingkatPendidikan(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       tingkat_pendidikan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateTingkatPendidikan(c echo.Context) error {
	db, _ := model.CreateCon()

	var tingkat_pendidikan model.TingkatPendidikan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&tingkat_pendidikan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(tingkat_pendidikan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	tingkat_pendidikan.CreatedBy = fmt.Sprint(c.Get("nip"))

	result := db.Create(&tingkat_pendidikan)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       tingkat_pendidikan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateTingkatPendidikan(c echo.Context) error {
	db, _ := model.CreateCon()

	var tingkat_pendidikan model.TingkatPendidikan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&tingkat_pendidikan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(tingkat_pendidikan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	tingkat_pendidikan.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.TingkatPendidikan{}).Where("id = ?", tingkat_pendidikan.Id).Updates(&tingkat_pendidikan)

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
		"data":       tingkat_pendidikan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteTingkatPendidikan(c echo.Context) error {
	db, _ := model.CreateCon()

	var tingkat_pendidikan model.DeleteTingkatPendidikan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&tingkat_pendidikan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(tingkat_pendidikan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.TingkatPendidikan{}).Where("id = ?", tingkat_pendidikan.Id).Delete(&tingkat_pendidikan)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       tingkat_pendidikan,
		"statucCode": http.StatusOK,
	})
}
