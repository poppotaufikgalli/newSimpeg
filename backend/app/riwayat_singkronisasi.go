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

func GetRiwayatSingkronisasi(searchString model.SearchRiwayatSingkronisasi) (riwayat_singkronisasi []model.RiwayatSingkronisasi, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatSingkronisasi{})

	//id
	if searchString.Id > 0 {
		result = result.Where("master_riwayat_singkronisasi.id = ?", searchString.Id)
	}

	//table
	if searchString.Table != "" {
		nama := strings.TrimSpace(searchString.Table)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_singkronisasi.table LIKE ?", strings.Join(str, ""))
	}

	//Method
	if searchString.Method != "" {
		result = result.Where("master_riwayat_singkronisasi.method = ?", searchString.Method)
	}

	result = result.Find(&riwayat_singkronisasi)

	return
}

func FindRiwayatSingkronisasi(c echo.Context) error {
	var searchString model.SearchRiwayatSingkronisasi
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_singkronisasi, result := GetRiwayatSingkronisasi(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_singkronisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateRiwayatSingkronisasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_singkronisasi model.RiwayatSingkronisasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_singkronisasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_singkronisasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_singkronisasi.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_singkronisasi)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_singkronisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatSingkronisasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_singkronisasi model.RiwayatSingkronisasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_singkronisasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_singkronisasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_singkronisasi.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RiwayatSingkronisasi{}).Where("id = ?", riwayat_singkronisasi.Id).Updates(&riwayat_singkronisasi)

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
		"data":       riwayat_singkronisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatSingkronisasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_singkronisasi model.DeleteRiwayatSingkronisasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_singkronisasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_singkronisasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatSingkronisasi{}).Where("id = ?", riwayat_singkronisasi.Id).Delete(&riwayat_singkronisasi)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_singkronisasi,
		"statucCode": http.StatusOK,
	})
}
