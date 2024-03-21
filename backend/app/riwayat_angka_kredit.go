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

func GetRiwayatAngkaKredit(searchString model.SearchRiwayatAngkaKredit) (riwayat_angka_kredit []model.RiwayatAngkaKredit, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatAngkaKredit{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_angka_kredit.nip = ?", searchString.Nip)
	}

	//Jns
	if searchString.Jns != "" {
		result = result.Where("master_riwayat_angka_kredit.jns = ?", searchString.Jns)
	}

	//kjab
	if searchString.Kjab != "" {
		result = result.Where("master_riwayat_angka_kredit.kjab = ?", searchString.Kjab)
	}

	//tmtjab
	if searchString.Tmulai.IsZero() != true {
		result = result.Where("master_riwayat_angka_kredit.tmulai = ?", searchString.Tmulai)
	}

	//kjab bkn
	if searchString.Thn != "" {
		result = result.Where("master_riwayat_angka_kredit.thn = ?", searchString.Thn)
	}

	result = result.Find(&riwayat_angka_kredit)

	return

}

func FindRiwayatAngkaKredit(c echo.Context) error {
	var searchString model.SearchRiwayatAngkaKredit
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_angka_kredit, result := GetRiwayatAngkaKredit(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_angka_kredit,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateRiwayatAngkaKredit(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_angka_kredit model.RiwayatAngkaKredit
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_angka_kredit)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_angka_kredit); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_angka_kredit.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_angka_kredit)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_angka_kredit,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatAngkaKredit(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_angka_kredit model.RiwayatAngkaKredit
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_angka_kredit)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_angka_kredit); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_angka_kredit.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RiwayatAngkaKredit{}).Where("nip = ? and tmulai = ?", riwayat_angka_kredit.Nip, riwayat_angka_kredit.Tmulai).Updates(&riwayat_angka_kredit)

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
		"data":       riwayat_angka_kredit,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatAngkaKredit(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_angka_kredit model.DeleteRiwayatAngkaKredit
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_angka_kredit)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_angka_kredit); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatAngkaKredit{}).Where("nip = ? and tmulai = ?", riwayat_angka_kredit.Nip, riwayat_angka_kredit.Tmulai).Delete(&riwayat_angka_kredit)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_angka_kredit,
		"statucCode": http.StatusOK,
	})
}
