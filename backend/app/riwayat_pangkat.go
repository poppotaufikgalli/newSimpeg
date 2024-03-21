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

func GetRiwayatPangkat(searchString model.SearchRiwayatPangkat) (riwayat_pangkat []model.RiwayatPangkat, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatPangkat{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_pangkat.nip = ?", searchString.Nip)
	}

	//tmtpang
	if searchString.Tmtpang.IsZero() != true {
		result = result.Where("master_riwayat_pangkat.tmtpang = ?", searchString.Tmtpang)
	}

	//Kgolru
	if len(searchString.Kgolru) > 0 {
		result = result.Where("master_riwayat_pangkat.kgolru IN (?)", searchString.Kgolru)
	}

	//Nskpangkat
	if searchString.Nskpang != "" {
		nskpang := strings.TrimSpace(searchString.Nskpang)
		str := []string{"%", nskpang, "%"}
		result = result.Where("master_riwayat_pangkat.nskpang LIKE ?", strings.Join(str, ""))
	}

	//Akhir
	if len(searchString.Akhir) > 0 {
		result = result.Where("master_riwayat_pangkat.akhir IN (?)", searchString.Akhir)
	}

	result = result.Find(&riwayat_pangkat)

	return

}

func FindRiwayatPangkat(c echo.Context) error {
	var searchString model.SearchRiwayatPangkat
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_pangkat, result := GetRiwayatPangkat(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateRiwayatPangkat(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_pangkat model.RiwayatPangkat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_pangkat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_pangkat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_pangkat.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_pangkat)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatPangkat(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_pangkat model.RiwayatPangkat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_pangkat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_pangkat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_pangkat.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and tmtpang = ? and kgolru = ? and knpang = ?", riwayat_pangkat.Nip, riwayat_pangkat.Tmtpang, riwayat_pangkat.Kgolru, riwayat_pangkat.Knpang).Updates(&riwayat_pangkat)

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
		"data":       riwayat_pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatPangkat(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_pangkat model.DeleteRiwayatPangkat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_pangkat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_pangkat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and tmtpang = ? and kgolru = ? and knpang = ?", riwayat_pangkat.Nip, riwayat_pangkat.Tmtpang, riwayat_pangkat.Kgolru, riwayat_pangkat.Knpang).Delete(&riwayat_pangkat)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pangkat,
		"statucCode": http.StatusOK,
	})
}
