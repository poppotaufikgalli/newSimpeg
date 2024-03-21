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

func GetKedudukanPegawai(searchString model.SearchKedudukanPegawai) (kedudukan_pegawai []model.KedudukanPegawai, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.KedudukanPegawai{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_kedudukan_pegawai.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_kedudukan_pegawai.nama LIKE ?", strings.Join(str, ""))
	}

	//id
	if searchString.RefSimpeg != "" {
		result = result.Where("master_kedudukan_pegawai.ref_simpeg = ?", searchString.RefSimpeg)
	}

	result = result.Find(&kedudukan_pegawai)

	return
}

func FindKedudukanPegawai(c echo.Context) error {
	var searchString model.SearchKedudukanPegawai

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	kedudukan_pegawai, result := GetKedudukanPegawai(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       kedudukan_pegawai,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateKedudukanPegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var kedudukan_pegawai model.KedudukanPegawai
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&kedudukan_pegawai)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(kedudukan_pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	kedudukan_pegawai.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&kedudukan_pegawai)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       kedudukan_pegawai,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateKedudukanPegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var kedudukan_pegawai model.KedudukanPegawai
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&kedudukan_pegawai)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(kedudukan_pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	kedudukan_pegawai.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.KedudukanPegawai{}).Where("id = ?", kedudukan_pegawai.Id).Updates(&kedudukan_pegawai)

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
		"data":       kedudukan_pegawai,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteKedudukanPegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var kedudukan_pegawai model.DeleteKedudukanPegawai
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&kedudukan_pegawai)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(kedudukan_pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.KedudukanPegawai{}).Where("id = ?", kedudukan_pegawai.Id).Delete(&kedudukan_pegawai)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       kedudukan_pegawai,
		"statucCode": http.StatusOK,
	})
}
