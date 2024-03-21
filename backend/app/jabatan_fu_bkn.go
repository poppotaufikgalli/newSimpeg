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

func GetJabatanFuBkn(searchString model.SearchJabatanFuBkn) (jabatan_fu_bkn []model.JabatanFuBkn, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JabatanFuBkn{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jabatan_fu_bkn.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jabatan_fu_bkn.nama LIKE ?", strings.Join(str, ""))
	}

	//status
	if len(searchString.Status) > 0 {
		result = result.Where("master_jabatan_fu_bkn.status IN (?)", searchString.Status)
	}

	//cepat kode
	if searchString.CepatKode != "" {
		result = result.Where("master_jabatan_fu_bkn.cepat_kode = ?", searchString.CepatKode)
	}

	//ref Simpeg
	if searchString.RefSimpeg != "" {
		result = result.Where("master_jabatan_fu_bkn.ref_simpeg = ?", searchString.RefSimpeg)
	}

	result = result.Find(&jabatan_fu_bkn)

	return
}

func FindJabatanFuBkn(c echo.Context) error {
	var searchString model.SearchJabatanFuBkn
	//var jabatan_fu_bkn []model.JabatanFuBkn

	err := json.NewDecoder(c.Request().Body).Decode(&searchString)
	//json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		searchString.Status = []int{1}
	}

	jabatan_fu_bkn, result := GetJabatanFuBkn(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_fu_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJabatanFuBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_fu_bkn model.JabatanFuBkn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_fu_bkn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_fu_bkn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jabatan_fu_bkn.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jabatan_fu_bkn)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_fu_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJabatanFuBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_fu_bkn model.JabatanFuBkn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_fu_bkn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_fu_bkn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jabatan_fu_bkn.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JabatanFuBkn{}).Where("id = ?", jabatan_fu_bkn.Id).Updates(&jabatan_fu_bkn)

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
		"data":       jabatan_fu_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJabatanFuBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_fu_bkn model.DeleteJabatanFuBkn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_fu_bkn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_fu_bkn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JabatanFuBkn{}).Where("id = ?", jabatan_fu_bkn.Id).Delete(&jabatan_fu_bkn)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_fu_bkn,
		"statucCode": http.StatusOK,
	})
}
