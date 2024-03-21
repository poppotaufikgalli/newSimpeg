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

func GetRiwayatKeluarga(searchString model.SearchRiwayatKeluarga) (riwayat_keluarga []model.RiwayatKeluarga, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatKeluarga{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_keluarga.nip = ?", searchString.Nip)
	}

	//Jkeluarga
	if searchString.Jkeluarga > 0 {
		result = result.Where("master_riwayat_keluarga.jkeluarga = ?", searchString.Jkeluarga)
	}

	//nama kel
	if searchString.NamaKel != "" {
		nama := strings.TrimSpace(searchString.NamaKel)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_keluarga.nama_kel LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&riwayat_keluarga)

	return

}

func FindRiwayatKeluarga(c echo.Context) error {
	var searchString model.SearchRiwayatKeluarga
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_keluarga, result := GetRiwayatKeluarga(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_keluarga,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateRiwayatKeluarga(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_keluarga model.RiwayatKeluarga
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_keluarga)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_keluarga); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_keluarga.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_keluarga)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_keluarga,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatKeluarga(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_keluarga model.RiwayatKeluarga
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_keluarga)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_keluarga); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_keluarga.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RiwayatKeluarga{}).Where("nip = ? and jkeluarga = ? and nama_kel = ?", riwayat_keluarga.Nip, riwayat_keluarga.Jkeluarga, riwayat_keluarga.NamaKel).Updates(&riwayat_keluarga)

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
		"data":       riwayat_keluarga,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatKeluarga(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_keluarga model.DeleteRiwayatKeluarga
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_keluarga)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_keluarga); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatKeluarga{}).Where("nip = ? and jkeluarga = ? and nama_kel = ?", riwayat_keluarga.Nip, riwayat_keluarga.Jkeluarga, riwayat_keluarga.NamaKel).Delete(&riwayat_keluarga)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_keluarga,
		"statucCode": http.StatusOK,
	})
}
