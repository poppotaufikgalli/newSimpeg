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

func GetRiwayatSkp(searchString model.SearchRiwayatSkp) (riwayat_skp []model.RiwayatSkp, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatSkp{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_skp.nip = ?", searchString.Nip)
	}

	//tahun
	if searchString.Tahun > 0 {
		result = result.Where("master_riwayat_skp.tahun = ?", searchString.Tahun)
	}

	//pns dinilai orang
	if searchString.PnsDinilaiOrang != "" {
		nama := strings.TrimSpace(searchString.PnsDinilaiOrang)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_skp.pnsDinilaiOrang LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&riwayat_skp)

	return
}

func FindRiwayatSkp(c echo.Context) error {
	var searchString model.SearchRiwayatSkp
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_skp, result := GetRiwayatSkp(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_skp,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateRiwayatSkp(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_skp model.RiwayatSkp
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_skp)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_skp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_skp.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_skp)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_skp,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatSkp(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_skp model.RiwayatSkp
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_skp)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_skp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_skp.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RiwayatSkp{}).Where("nip = ? and tahun = ?", riwayat_skp.Nip, riwayat_skp.Tahun).Updates(&riwayat_skp)

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
		"data":       riwayat_skp,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatSkp(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_skp model.DeleteRiwayatSkp
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_skp)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_skp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatSkp{}).Where("nip = ? and tahun = ?", riwayat_skp.Nip, riwayat_skp.Tahun).Delete(&riwayat_skp)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_skp,
		"statucCode": http.StatusOK,
	})
}
