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

func GetJabatanFtBkn(searchString model.SearchJabatanFtBkn) (jabatan_ft_bkn []model.JabatanFtBkn, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JabatanFtBkn{}).Preload("SubJabatan")

	//id
	if searchString.Id != "" {
		result = result.Where("master_jabatan_ft_bkn.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jabatan_ft_bkn.nama LIKE ?", strings.Join(str, ""))
	}

	//kel jabatan id
	if searchString.KelJabatanId != "" {
		result = result.Where("master_jabatan_ft_bkn.kel_jabatan_id = ?", searchString.KelJabatanId)
	}

	//jenjang
	if len(searchString.Jenjang) > 0 {
		result = result.Where("master_jabatan_ft_bkn.jenjang IN (?)", searchString.Jenjang)
	}

	//status
	if len(searchString.Status) > 0 {
		result = result.Where("master_jabatan_ft_bkn.status IN (?)", searchString.Status)
	}

	//cepat kode
	if searchString.CepatKode != "" {
		result = result.Where("master_jabatan_ft_bkn.cepat_kode = ?", searchString.CepatKode)
	}

	//ref Simpeg
	if searchString.RefSimpeg != "" {
		result = result.Where("master_jabatan_ft_bkn.ref_simpeg = ?", searchString.RefSimpeg)
	}

	result = result.Find(&jabatan_ft_bkn)

	return
}

func FindJabatanFtBkn(c echo.Context) error {
	var searchString model.SearchJabatanFtBkn

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jabatan_ft_bkn, result := GetJabatanFtBkn(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_ft_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func SearchJabatanFtBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_ft_bkn []model.JabatanFtBkn

	req := model.SearchInput{
		Limit: 10,
	}

	if err := c.Bind(&req); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result := db.Model(&model.JabatanFtBkn{})

	if req.SearchNama != "" {
		nama := strings.TrimSpace(req.SearchNama)
		str := []string{nama, "%"}
		result = result.Where("master_jabatan_ft_bkn.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Order("nama asc").Limit(req.Limit).Find(&jabatan_ft_bkn)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_ft_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateJabatanFtBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_ft_bkn model.JabatanFtBkn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_ft_bkn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_ft_bkn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jabatan_ft_bkn.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jabatan_ft_bkn)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_ft_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJabatanFtBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_ft_bkn model.JabatanFtBkn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_ft_bkn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_ft_bkn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jabatan_ft_bkn.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JabatanFtBkn{}).Where("id = ?", jabatan_ft_bkn.Id).Updates(&jabatan_ft_bkn)

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
		"data":       jabatan_ft_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJabatanFtBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_ft_bkn model.DeleteJabatanFtBkn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_ft_bkn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_ft_bkn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JabatanFtBkn{}).Where("id = ?", jabatan_ft_bkn.Id).Delete(&jabatan_ft_bkn)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_ft_bkn,
		"statucCode": http.StatusOK,
	})
}
