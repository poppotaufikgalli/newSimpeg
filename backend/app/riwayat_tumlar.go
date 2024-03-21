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

func GetRiwayatTumlar(searchString model.SearchRiwayatTumlar) (riwayat_tumlar []model.RiwayatTumlar, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatTumlar{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_tumlar.nip = ?", searchString.Nip)
	}

	//tsk
	if searchString.Tsk.IsZero() != true {
		result = result.Where("master_riwayat_tumlar.tsk = ?", searchString.Tsk)
	}

	//Ktpu
	if searchString.Ktpu != "" {
		result = result.Where("master_riwayat_tumlar.ktpu = ?", searchString.Ktpu)
	}

	//Kjur
	if searchString.Kjur != "" {
		result = result.Where("master_riwayat_tumlar.kjur = ?", searchString.Kjur)
	}

	result = result.Find(&riwayat_tumlar)

	return
}

func FindRiwayatTumlar(c echo.Context) error {
	var searchString model.SearchRiwayatTumlar
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_tumlar, result := GetRiwayatTumlar(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_tumlar,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateRiwayatTumlar(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_tumlar model.RiwayatTumlar
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_tumlar)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_tumlar); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_tumlar.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_tumlar)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_tumlar,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatTumlar(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_tumlar model.RiwayatTumlar
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_tumlar)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_tumlar); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_tumlar.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RiwayatTumlar{}).Where("nip = ? and tsk = ?", riwayat_tumlar.Nip, riwayat_tumlar.Tsk).Updates(&riwayat_tumlar)

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
		"data":       riwayat_tumlar,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatTumlar(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_tumlar model.DeleteRiwayatTumlar
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_tumlar)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_tumlar); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatTumlar{}).Where("nip = ? and tsk = ?", riwayat_tumlar.Nip, riwayat_tumlar.Tsk).Delete(&riwayat_tumlar)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_tumlar,
		"statucCode": http.StatusOK,
	})
}
