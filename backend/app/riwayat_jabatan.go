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

func GetRiwayatJabatan(searchString model.SearchRiwayatJabatan) (riwayat_jabatan []model.RiwayatJabatan, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatJabatan{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_jabatan.nip = ?", searchString.Nip)
	}

	//tmtjab
	if searchString.Tmtjab.IsZero() != true {
		result = result.Where("master_riwayat_jabatan.tmtjab = ?", searchString.Tmtjab)
	}

	//kjab
	if searchString.Kjab != "" {
		result = result.Where("master_riwayat_jabatan.kjab = ?", searchString.Kjab)
	}

	//kjab bkn
	if searchString.KjabBkn != "" {
		result = result.Where("master_riwayat_jabatan.kjab_bkn = ?", searchString.KjabBkn)
	}

	//Keselon
	if len(searchString.Keselon) > 0 {
		result = result.Where("master_riwayat_jabatan.keselon IN (?)", searchString.Keselon)
	}

	//Nskjabat
	if searchString.Nskjabat != "" {
		nskjabat := strings.TrimSpace(searchString.Nskjabat)
		str := []string{"%", nskjabat, "%"}
		result = result.Where("master_riwayat_jabatan.nskjabat LIKE ?", strings.Join(str, ""))
	}

	//id Opd
	if searchString.IdOpd != "" {
		result = result.Where("master_riwayat_jabatan.id_opd = ?", searchString.IdOpd)
	}

	//Akhir
	if len(searchString.Akhir) > 0 {
		result = result.Where("master_riwayat_jabatan.akhir IN (?)", searchString.Akhir)
	}

	//Tugas Tambahan
	if searchString.TugasTambahan != "" {
		result = result.Where("master_riwayat_jabatan.tugas_tambahan = ?", searchString.TugasTambahan)
	}

	//Id Opd Tambahan
	if searchString.IdOpdTambahan != "" {
		result = result.Where("master_riwayat_jabatan.id_opd_tambahan = ?", searchString.IdOpdTambahan)
	}

	result = result.Find(&riwayat_jabatan)

	return

}

func FindRiwayatJabatan(c echo.Context) error {
	var searchString model.SearchRiwayatJabatan
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_jabatan, result := GetRiwayatJabatan(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateRiwayatJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_jabatan model.RiwayatJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_jabatan.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_jabatan)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_jabatan model.RiwayatJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_jabatan.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RiwayatJabatan{}).Where("nip = ? and tmtjab = ?", riwayat_jabatan.Nip, riwayat_jabatan.Tmtjab).Updates(&riwayat_jabatan)

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
		"data":       riwayat_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_jabatan model.DeleteRiwayatJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatJabatan{}).Where("nip = ? and tmtjab = ?", riwayat_jabatan.Nip, riwayat_jabatan.Tmtjab).Delete(&riwayat_jabatan)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_jabatan,
		"statucCode": http.StatusOK,
	})
}
