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

func GetOPD(searchString model.SearchOpd) (opd []model.Opd, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Opd{}).Preload("SubOpd").Preload("ListFormasiJabatan").Preload("Eselon")

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_opd.nama LIKE ?", strings.Join(str, ""))
	}

	//status
	if len(searchString.Status) > 0 {
		result = result.Where("master_opd.status IN (?)", searchString.Status)
	}

	//id_eselon
	if len(searchString.IdEselon) > 0 {
		result = result.Where("master_opd.id_eselon IN (?)", searchString.IdEselon)
	}

	//sfilter
	if len(searchString.Sfilter) > 0 {
		result = result.Where("master_opd.sfilter IN (?)", searchString.Sfilter)
	}

	//id
	if searchString.Id != "" {
		result = result.Where("master_opd.id = ?", searchString.Id)
	}

	//parent_opd
	if searchString.ParentOpd != "" {
		result = result.Where("master_opd.parent_opd = ?", searchString.ParentOpd)
	}

	//kunker
	if searchString.Kunker != "" {
		result = result.Where("master_opd.kunker = ?", searchString.Kunker)
	}

	//disingkat
	if searchString.Disingkat != "" {
		disingkat := strings.TrimSpace(searchString.Disingkat)
		str := []string{"%", disingkat, "%"}
		result = result.Where("master_opd.disingkat LIKE ?", strings.Join(str, ""))
	}

	//id_unor_bkn
	if searchString.IdUnorBkn != "" {
		result = result.Where("master_opd.id_unor_bkn = ?", searchString.IdUnorBkn)
	}

	result = result.Find(&opd)

	return
}

func FindOPD(c echo.Context) error {
	var searchString model.SearchOpd
	//var opd []model.Opd

	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		searchString.Sfilter = []int{1}
	}

	opd, result := GetOPD(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       opd,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetOPDbyId(c echo.Context) error {
	db, _ := model.CreateCon()
	var opd model.Opd
	id := c.Param("id")

	result := db.Model(&model.Opd{}).Where("id = ?", id).Scan(&opd)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       opd,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateOPD(c echo.Context) error {
	db, _ := model.CreateCon()

	var opd model.Opd
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&opd)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(opd); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	opd.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&opd)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       opd,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateOPD(c echo.Context) error {
	db, _ := model.CreateCon()

	var opd model.Opd
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&opd)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(opd); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	opd.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Opd{}).Where("id = ?", opd.Id).Updates(&opd)

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
		"data":       opd,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteOPD(c echo.Context) error {
	db, _ := model.CreateCon()

	var opd model.DeleteOpd
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&opd)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(opd); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Opd{}).Where("id = ?", opd.Id).Delete(&opd)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       opd,
		"statucCode": http.StatusOK,
	})
}
