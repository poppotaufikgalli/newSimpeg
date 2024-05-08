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
	//"strings"
)

func GetRiwayatDiklat(searchString model.SearchRiwayatDiklat) (riwayat_diklat []model.RiwayatDiklat, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatDiklat{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_diklat.nip = ?", searchString.Nip)
	}

	//jdiklat
	if searchString.Jdiklat > 0 {
		result = result.Where("master_riwayat_diklat.jdiklat = ?", searchString.Jdiklat)
	}

	//kdiklat
	if searchString.Kdiklat > 0 {
		result = result.Where("master_riwayat_diklat.kdiklat = ?", searchString.Kdiklat)
	}

	//tempat
	if searchString.Tempat != "" {
		result = result.Where("master_riwayat_diklat.tempat = ?", searchString.Tempat)
	}

	//angkatan
	if searchString.Angkatan != "" {
		result = result.Where("master_riwayat_diklat.angkatan = ?", searchString.Angkatan)
	}

	//tmulai
	if searchString.Tmulai.IsZero() != true {
		result = result.Where("master_riwayat_diklat.tmulai = ?", searchString.Tmulai)
	}

	result = result.Find(&riwayat_diklat)

	return

}

func FindRiwayatDiklat(c echo.Context) error {
	var searchString model.SearchRiwayatDiklat
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_diklat, result := GetRiwayatDiklat(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_diklat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetRiwayatDiklatByNipJdiklatTmulai(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_diklat model.RiwayatDiklat
	nip := c.Param("nip")
	jdiklat := c.Param("jdiklat")
	tmulai := c.Param("tmulai")

	result := db.Model(&model.RiwayatDiklat{}).Where("nip =? and jdiklat = ? and tmulai = ?", nip, jdiklat, tmulai).Scan(&riwayat_diklat)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_diklat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateRiwayatDiklat(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_diklat model.RiwayatDiklat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_diklat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_diklat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_diklat.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_diklat)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_diklat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatDiklat(c echo.Context) error {
	db, _ := model.CreateCon()

	//var riwayat_diklat model.RiwayatDiklat
	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatDiklat `json:"refdata"`
		Data model.RiwayatDiklat       `json:"data"`
	}

	var reqData ToUpdateData

	err := json.NewDecoder(c.Request().Body).Decode(&reqData)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(reqData.Data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	reqData.Data.UpdatedBy = fmt.Sprint(c.Get("nip"))
	Tmulai := (reqData.Ref.Tmulai).Format("2006-01-02")
	result := db.Model(&model.RiwayatDiklat{}).Where("nip = ? and jdiklat = ? and tmulai = ?", reqData.Ref.Nip, reqData.Ref.Jdiklat, Tmulai).Updates(&reqData.Data)

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
		"data":       reqData.Data,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatDiklat(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_diklat model.DeleteRiwayatDiklat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_diklat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_diklat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatDiklat{}).Where("nip = ? and jdiklat = ? and tmulai = ?", riwayat_diklat.Nip, riwayat_diklat.Jdiklat, riwayat_diklat.Tmulai).Delete(&riwayat_diklat)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_diklat,
		"statucCode": http.StatusOK,
	})
}
