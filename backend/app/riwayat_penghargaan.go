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

func GetRiwayatPenghargaan(searchString model.SearchRiwayatPenghargaan) (riwayat_penghargaan []model.RiwayatPenghargaan, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatPenghargaan{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_penghargaan.nip = ?", searchString.Nip)
	}

	//id jenis penghargaan
	if searchString.IdJenisPenghargaan != "" {
		result = result.Where("master_riwayat_penghargaan.id_jenis_penghargaan = ?", searchString.IdJenisPenghargaan)
	}

	//nsk
	if searchString.Nsk != "" {
		nama := strings.TrimSpace(searchString.Nsk)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_penghargaan.nsk LIKE ?", strings.Join(str, ""))
	}

	//Thn
	if searchString.Thn > 0 {
		result = result.Where("master_riwayat_penghargaan.thn = ?", searchString.Thn)
	}

	result = result.Preload("JenisPenghargaan").Find(&riwayat_penghargaan)

	return

}

func FindRiwayatPenghargaan(c echo.Context) error {
	var searchString model.SearchRiwayatPenghargaan
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_penghargaan, result := GetRiwayatPenghargaan(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_penghargaan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateRiwayatPenghargaan(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_penghargaan model.RiwayatPenghargaan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_penghargaan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_penghargaan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_penghargaan.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_penghargaan)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_penghargaan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatPenghargaan(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_penghargaan model.RiwayatPenghargaan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_penghargaan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_penghargaan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_penghargaan.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RiwayatPenghargaan{}).Where("nip = ? and nbintang = ? and nsk = ?", riwayat_penghargaan.Nip, riwayat_penghargaan.Nbintang, riwayat_penghargaan.Nsk).Updates(&riwayat_penghargaan)

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
		"data":       riwayat_penghargaan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatPenghargaan(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_penghargaan model.DeleteRiwayatPenghargaan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_penghargaan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_penghargaan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatPenghargaan{}).Where("nip = ? and nbintang = ? and nsk = ?", riwayat_penghargaan.Nip, riwayat_penghargaan.Nbintang, riwayat_penghargaan.Nsk).Delete(&riwayat_penghargaan)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_penghargaan,
		"statucCode": http.StatusOK,
	})
}
