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

func GetKppn(searchString model.SearchKppn) (kppn []model.Kppn, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Kppn{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_kppn.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_kppn.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&kppn)

	return
}

func FindKppn(c echo.Context) error {
	var searchString model.SearchKppn

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	kppn, result := GetKppn(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       kppn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateKppn(c echo.Context) error {
	db, _ := model.CreateCon()

	var kppn model.Kppn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&kppn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(kppn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	kppn.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&kppn)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       kppn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateKppn(c echo.Context) error {
	db, _ := model.CreateCon()

	var kppn model.Kppn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&kppn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(kppn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	kppn.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Kppn{}).Where("id = ?", kppn.Id).Updates(&kppn)

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
		"data":       kppn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteKppn(c echo.Context) error {
	db, _ := model.CreateCon()

	var kppn model.DeleteKppn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&kppn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(kppn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Kppn{}).Where("id = ?", kppn.Id).Delete(&kppn)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       kppn,
		"statucCode": http.StatusOK,
	})
}
