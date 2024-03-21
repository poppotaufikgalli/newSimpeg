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

func GetJabatanFu(searchString model.SearchJabatanFu) (jabatan_fu []model.JabatanFu, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JabatanFu{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jabatan_fu.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jabatan_fu.nama LIKE ?", strings.Join(str, ""))
	}

	//status
	if len(searchString.Status) > 0 {
		result = result.Where("master_jabatan_fu.status IN (?)", searchString.Status)
	}

	//cepat kode
	if searchString.CepatKode != "" {
		result = result.Where("master_jabatan_fu.cepat_kode = ?", searchString.CepatKode)
	}

	//ref Simpeg
	if searchString.RefSimpeg != "" {
		result = result.Where("master_jabatan_fu.ref_simpeg = ?", searchString.RefSimpeg)
	}

	result = result.Find(&jabatan_fu)

	return
}

func FindJabatanFu(c echo.Context) error {
	var searchString model.SearchJabatanFu

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jabatan_fu, result := GetJabatanFu(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_fu,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJabatanFu(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_fu model.JabatanFu
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_fu)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_fu); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jabatan_fu.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jabatan_fu)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_fu,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJabatanFu(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_fu model.JabatanFu
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_fu)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_fu); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jabatan_fu.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JabatanFu{}).Where("id = ?", jabatan_fu.Id).Updates(&jabatan_fu)

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
		"data":       jabatan_fu,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJabatanFu(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_fu model.DeleteJabatanFu
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_fu)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_fu); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JabatanFu{}).Where("id = ?", jabatan_fu.Id).Delete(&jabatan_fu)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_fu,
		"statucCode": http.StatusOK,
	})
}
