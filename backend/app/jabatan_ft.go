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

func GetJabatanFt(searchString model.SearchJabatanFt) (jabatan_ft []model.JabatanFt, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JabatanFt{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jabatan_ft.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jabatan_ft.nama LIKE ?", strings.Join(str, ""))
	}

	//kel jabatan id
	if searchString.KelJabatanId != "" {
		result = result.Where("master_jabatan_ft.kel_jabatan_id = ?", searchString.KelJabatanId)
	}

	//jenjang
	if len(searchString.Jenjang) > 0 {
		result = result.Where("master_jabatan_ft.jenjang IN (?)", searchString.Jenjang)
	}

	//status
	if len(searchString.Status) > 0 {
		result = result.Where("master_jabatan_ft.status IN (?)", searchString.Status)
	}

	//cepat kode
	if searchString.CepatKode != "" {
		result = result.Where("master_jabatan_ft.cepat_kode = ?", searchString.CepatKode)
	}

	//Ref BKN
	if searchString.RefBkn != "" {
		result = result.Where("master_jabatan_ft.ref_bkn = ?", searchString.RefBkn)
	}

	//Ref Simpeg
	if searchString.RefSimpeg != "" {
		result = result.Where("master_jabatan_ft.ref_simpeg = ?", searchString.RefSimpeg)
	}

	result = result.Find(&jabatan_ft)

	return
}

func FindJabatanFt(c echo.Context) error {
	var searchString model.SearchJabatanFt
	//var jabatan_ft []model.JabatanFt

	//err := json.NewDecoder(c.Request().Body).Decode(&searchString)
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	//if err != nil {
	//	searchString.Status = []int{1}
	//}

	jabatan_ft, result := GetJabatanFt(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_ft,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJabatanFt(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_ft model.JabatanFt
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_ft)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_ft); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jabatan_ft.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jabatan_ft)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_ft,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJabatanFt(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_ft model.JabatanFt
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_ft)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_ft); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jabatan_ft.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JabatanFt{}).Where("id = ?", jabatan_ft.Id).Updates(&jabatan_ft)

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
		"data":       jabatan_ft,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJabatanFt(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_ft model.DeleteJabatanFt
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_ft)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_ft); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JabatanFt{}).Where("id = ?", jabatan_ft.Id).Delete(&jabatan_ft)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_ft,
		"statucCode": http.StatusOK,
	})
}
