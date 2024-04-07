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

func GetRiwayatPindahInstansi(searchString model.SearchRiwayatPindahInstansi) (riwayat_pindah_instansi []model.RiwayatPindahInstansi, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatPindahInstansi{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_pindah_instansi.nip = ?", searchString.Nip)
	}

	//Instansi Induk Baru
	if searchString.InstansiIndukBaru != "" {
		nama := strings.TrimSpace(searchString.InstansiIndukBaru)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_pindah_instansi.instansiIndukBaru LIKE ?", strings.Join(str, ""))
	}

	//Instansi Induk Lama
	if searchString.InstansiIndukLama != "" {
		nama := strings.TrimSpace(searchString.InstansiIndukLama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_pindah_instansi.instansiIndukLama LIKE ?", strings.Join(str, ""))
	}

	//Jenis Pi
	if searchString.JenisPi != "" {
		result = result.Where("master_riwayat_pindah_instansi.jenisPi = ?", searchString.JenisPi)
	}

	//Tmt Pi
	if searchString.TmtPi.IsZero() != true {
		result = result.Where("master_riwayat_pindah_instansi.tmtPi = ?", searchString.TmtPi)
	}

	result = result.Preload("InstansiBaru").Preload("InstansiLama").Order("tmtPi desc").Find(&riwayat_pindah_instansi)

	return

}

func FindRiwayatPindahInstansi(c echo.Context) error {
	var searchString model.SearchRiwayatPindahInstansi
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_pindah_instansi, result := GetRiwayatPindahInstansi(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pindah_instansi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateRiwayatPindahInstansi(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_pindah_instansi model.RiwayatPindahInstansi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_pindah_instansi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_pindah_instansi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_pindah_instansi.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_pindah_instansi)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pindah_instansi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatPindahInstansi(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_pindah_instansi model.RiwayatPindahInstansi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_pindah_instansi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_pindah_instansi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_pindah_instansi.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RiwayatPindahInstansi{}).Where("nip = ? and tmtPi = ?", riwayat_pindah_instansi.Nip, riwayat_pindah_instansi.TmtPi).Updates(&riwayat_pindah_instansi)

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
		"data":       riwayat_pindah_instansi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatPindahInstansi(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_pindah_instansi model.DeleteRiwayatPindahInstansi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_pindah_instansi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_pindah_instansi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatPindahInstansi{}).Where("nip = ? and tmtPi = ?", riwayat_pindah_instansi.Nip, riwayat_pindah_instansi.TmtPi).Delete(&riwayat_pindah_instansi)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pindah_instansi,
		"statucCode": http.StatusOK,
	})
}
