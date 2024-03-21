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

func GetRiwayatBahasa(searchString model.SearchRiwayatBahasa) (riwayat_bahasa []model.RiwayatBahasa, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatBahasa{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_bahasa.nip = ?", searchString.Nip)
	}

	//Nbahasa
	if searchString.Nbahasa != "" {
		result = result.Where("master_riwayat_bahasa.nbahasa = ?", searchString.Nbahasa)
	}

	//kbahasa
	if len(searchString.Kbahasa) > 0 {
		result = result.Where("master_riwayat_bahasa.kbahasa IN (?)", searchString.Kbahasa)
	}

	//jbahasa
	if len(searchString.Jbahasa) > 0 {
		result = result.Where("master_riwayat_bahasa.jbahasa IN (?)", searchString.Jbahasa)
	}

	result = result.Find(&riwayat_bahasa)

	return

}

func FindRiwayatBahasa(c echo.Context) error {
	var searchString model.SearchRiwayatBahasa
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_bahasa, result := GetRiwayatBahasa(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_bahasa,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateRiwayatBahasa(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_bahasa model.RiwayatBahasa
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_bahasa)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_bahasa); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_bahasa.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_bahasa)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_bahasa,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatBahasa(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_bahasa model.RiwayatBahasa
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_bahasa)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_bahasa); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_bahasa.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RiwayatBahasa{}).Where("nip = ? and nbahasa = ?", riwayat_bahasa.Nip, riwayat_bahasa.Nbahasa).Updates(&riwayat_bahasa)

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
		"data":       riwayat_bahasa,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatBahasa(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_bahasa model.DeleteRiwayatBahasa
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_bahasa)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_bahasa); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatBahasa{}).Where("nip = ? and nbahasa = ?", riwayat_bahasa.Nip, riwayat_bahasa.Nbahasa).Delete(&riwayat_bahasa)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_bahasa,
		"statucCode": http.StatusOK,
	})
}
