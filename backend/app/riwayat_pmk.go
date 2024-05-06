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

func GetRiwayatPmk(searchString model.SearchRiwayatPmk) (riwayat_pmk []model.RiwayatPmk, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatPmk{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_pmk.nip = ?", searchString.Nip)
	}

	//dinilai
	if searchString.Dinilai > 0 {
		result = result.Where("master_riwayat_pmk.dinilai = ?", searchString.Dinilai)
	}

	//Tanggal Awal
	if searchString.TanggalAwal.IsZero() != true {
		result = result.Where("master_riwayat_pmk.tanggalAwal = ?", searchString.TanggalAwal)
	}

	result = result.Find(&riwayat_pmk)

	return

}

func FindRiwayatPmk(c echo.Context) error {
	var searchString model.SearchRiwayatPmk
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_pmk, result := GetRiwayatPmk(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pmk,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetRiwayatPmkByNiptanggalAwal(c echo.Context) error {
	db, _ := model.CreateCon()
	var riwayat_pmk model.RiwayatPmk
	nip := c.Param("nip")
	tanggalAwal := c.Param("tanggalAwal")

	result := db.Model(&model.RiwayatPmk{}).Where("nip = ? and tanggalAwal = ?", nip, tanggalAwal).Scan(&riwayat_pmk)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pmk,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateRiwayatPmk(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_pmk model.RiwayatPmk
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_pmk)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_pmk); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_pmk.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_pmk)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pmk,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatPmk(c echo.Context) error {
	db, _ := model.CreateCon()

	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatPmk `json:"refdata"`
		Data model.RiwayatPmk       `json:"data"`
	}

	var reqData ToUpdateData

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
	TanggalAwal := (reqData.Ref.TanggalAwal).Format("2006-01-02")
	result := db.Model(&model.RiwayatPmk{}).Where("nip = ? and tanggalAwal = ?", reqData.Ref.Nip, TanggalAwal).Updates(&reqData.Data)

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

func DeleteRiwayatPmk(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_pmk model.DeleteRiwayatPmk
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_pmk)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_pmk); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatPmk{}).Where("nip = ? and tanggalAwal = ?", riwayat_pmk.Nip, riwayat_pmk.TanggalAwal).Delete(&riwayat_pmk)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pmk,
		"statucCode": http.StatusOK,
	})
}
