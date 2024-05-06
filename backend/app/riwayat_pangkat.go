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

	result = result.Preload("Pangkat").Preload("JenisKp").Order("kgolru desc").Find(&riwayat_pangkat)

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

func GetRiwayatPangkatByNip(c echo.Context) error {
	db, _ := model.CreateCon()
	var pangkat []model.RiwayatPangkat
	nip := c.Param("nip")

	result := db.Model(&model.RiwayatPangkat{}).Preload("Pangkat").Preload("JenisKp").Order("kgolru desc").Where("nip = ?", nip).Scan(&pangkat)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatPangkatByNipByKgolru(c echo.Context) error {
	db, _ := model.CreateCon()
	var pangkat model.RiwayatPangkat
	nip := c.Param("nip")
	kgolru := c.Param("kgolru")

	result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and kgolru = ?", nip, kgolru).Scan(&pangkat)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatPangkatByNipByKgolruTmtpangKnpang(c echo.Context) error {
	db, _ := model.CreateCon()
	var pangkat model.RiwayatPangkat
	nip := c.Param("nip")
	kgolru := c.Param("kgolru")
	tmtpang := c.Param("tmtpang")
	knpang := c.Param("knpang")

	result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and kgolru = ? and tmtpang = ? and knpang = ?", nip, kgolru, tmtpang, knpang).Scan(&pangkat)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatPangkatByNipCPNS(c echo.Context) error {
	db, _ := model.CreateCon()
	var pangkat model.RiwayatPangkat
	nip := c.Param("nip")

	result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and knpang = 211", nip).Scan(&pangkat)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatPangkatByNipPNS(c echo.Context) error {
	db, _ := model.CreateCon()
	var pangkat model.RiwayatPangkat
	nip := c.Param("nip")

	result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and knpang = 212", nip).Scan(&pangkat)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatPangkatByNipAkhir(c echo.Context) error {
	db, _ := model.CreateCon()
	var pangkat model.RiwayatPangkat
	nip := c.Param("nip")

	result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and akhir = 1", nip).Scan(&pangkat)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
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

	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatPangkat `json:"refdata"`
		Data model.RiwayatPangkat       `json:"data"`
	}

	var reqData ToUpdateData

	//var riwayat_pangkat model.RiwayatPangkat

	err := json.NewDecoder(c.Request().Body).Decode(&reqData)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(reqData.Data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	reqData.Data.UpdatedBy = fmt.Sprint(c.Get("nip"))
	Tmtpang := (reqData.Ref.Tmtpang).Format("2006-01-02")
	result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and tmtpang = ? and kgolru = ? and knpang = ?", reqData.Ref.Nip, Tmtpang, reqData.Ref.Kgolru, reqData.Ref.Knpang).Updates(&reqData.Data)

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
		"data":       reqData.Data,
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
