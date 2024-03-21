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

func GetEselon(searchString model.SearchEselon) (eselon []model.Eselon, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Eselon{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_eselon.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_eselon.nama LIKE ?", strings.Join(str, ""))
	}

	//jabatan Asn
	if searchString.JabatanAsn != "" {
		jabatan := strings.TrimSpace(searchString.JabatanAsn)
		str := []string{"%", jabatan, "%"}
		result = result.Where("master_eselon.jabatan_asn LIKE ?", strings.Join(str, ""))
	}

	//id_jenis_jabatan
	if len(searchString.IdJenisJabatan) > 0 {
		result = result.Where("master_eselon.id_jenis_jabatan IN (?)", searchString.IdJenisJabatan)
	}

	//status
	if len(searchString.Status) > 0 {
		result = result.Where("master_eselon.status IN (?)", searchString.Status)
	}

	result = result.Find(&eselon)

	return
}

func FindEselon(c echo.Context) error {
	var searchString model.SearchEselon
	//var eselon []model.Eselon

	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		searchString.Status = []int{1}
	}

	eselon, result := GetEselon(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       eselon,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateEselon(c echo.Context) error {
	db, _ := model.CreateCon()

	var eselon model.Eselon
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&eselon)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(eselon); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	eselon.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&eselon)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       eselon,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateEselon(c echo.Context) error {
	db, _ := model.CreateCon()

	var eselon model.Eselon
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&eselon)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(eselon); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	eselon.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Eselon{}).Where("id = ?", eselon.Id).Updates(&eselon)

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
		"data":       eselon,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteEselon(c echo.Context) error {
	db, _ := model.CreateCon()

	var eselon model.DeleteEselon
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&eselon)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(eselon); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Eselon{}).Where("id = ?", eselon.Id).Delete(&eselon)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       eselon,
		"statucCode": http.StatusOK,
	})
}
