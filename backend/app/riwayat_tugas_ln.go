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

func GetRiwayatTugasLn(searchString model.SearchRiwayatTugasLn) (riwayat_tugas_ln []model.RiwayatTugasLn, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatTugasLn{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_tugas_ln.nip = ?", searchString.Nip)
	}

	//Nneg
	if searchString.Nneg != "" {
		nama := strings.TrimSpace(searchString.Nneg)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_tugas_ln.nneg LIKE ?", strings.Join(str, ""))
	}

	//Tujuan
	if searchString.Tujuan != "" {
		nama := strings.TrimSpace(searchString.Tujuan)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_tugas_ln.tujuan LIKE ?", strings.Join(str, ""))
	}

	//Tmulai
	if searchString.Tmulai.IsZero() != true {
		result = result.Where("master_riwayat_tugas_ln.tmulai = ?", searchString.Tmulai)
	}

	result = result.Find(&riwayat_tugas_ln)

	return

}

func FindRiwayatTugasLn(c echo.Context) error {
	var searchString model.SearchRiwayatTugasLn
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_tugas_ln, result := GetRiwayatTugasLn(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_tugas_ln,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateRiwayatTugasLn(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_tugas_ln model.RiwayatTugasLn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_tugas_ln)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_tugas_ln); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_tugas_ln.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_tugas_ln)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_tugas_ln,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatTugasLn(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_tugas_ln model.RiwayatTugasLn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_tugas_ln)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_tugas_ln); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_tugas_ln.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RiwayatTugasLn{}).Where("nip = ? and tmulai = ?", riwayat_tugas_ln.Nip, riwayat_tugas_ln.Tmulai).Updates(&riwayat_tugas_ln)

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
		"data":       riwayat_tugas_ln,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatTugasLn(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_tugas_ln model.DeleteRiwayatTugasLn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_tugas_ln)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_tugas_ln); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatTugasLn{}).Where("nip = ? and tmulai = ?", riwayat_tugas_ln.Nip, riwayat_tugas_ln.Tmulai).Delete(&riwayat_tugas_ln)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_tugas_ln,
		"statucCode": http.StatusOK,
	})
}
