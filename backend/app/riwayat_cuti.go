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

func GetRiwayatCuti(searchString model.SearchRiwayatCuti) (riwayat_cuti []model.RiwayatCuti, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatCuti{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_cuti.nip = ?", searchString.Nip)
	}

	//Kcuti
	if searchString.Kcuti > 0 {
		result = result.Where("master_riwayat_cuti.kcuti = ?", searchString.Kcuti)
	}

	//Nsk
	if searchString.Nsk != "" {
		nama := strings.TrimSpace(searchString.Nsk)
		str := []string{"%", nama, "%"}
		result = result.Where("master_pekerjaan.nsk LIKE ?", strings.Join(str, ""))
	}

	//tmulai
	if searchString.Tmulai.IsZero() != true {
		result = result.Where("master_riwayat_cuti.tmulai = ?", searchString.Tmulai)
	}

	//thn
	if searchString.Thn != "" {
		result = result.Where("master_riwayat_cuti.thn = ?", searchString.Thn)
	}

	result = result.Preload("JenisCuti").Find(&riwayat_cuti)

	return

}

func FindRiwayatCuti(c echo.Context) error {
	var searchString model.SearchRiwayatCuti
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_cuti, result := GetRiwayatCuti(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_cuti,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetRiwayatCutiByNipNskTmulai(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_cuti model.RiwayatCuti
	nip := c.Param("nip")
	nsk := c.Param("nsk")
	tmulai := c.Param("tmulai")

	result := db.Model(&model.RiwayatCuti{}).Where("nip =? and nsk =? and tmulai =?", nip, nsk, tmulai).Scan(&riwayat_cuti)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_cuti,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateRiwayatCuti(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_cuti model.RiwayatCuti
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_cuti)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_cuti); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_cuti.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_cuti)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_cuti,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatCuti(c echo.Context) error {
	db, _ := model.CreateCon()

	//var riwayat_cuti model.RiwayatCuti
	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatCuti `json:"refdata"`
		Data model.RiwayatCuti       `json:"data"`
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
	Tmulai := (reqData.Ref.Tmulai).Format("2006-01-02")
	result := db.Model(&model.RiwayatCuti{}).Debug().Where("nip = ? and tmulai = ? and nsk = ?", reqData.Ref.Nip, Tmulai, reqData.Ref.Nsk).Updates(&reqData.Data)

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

func DeleteRiwayatCuti(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_cuti model.DeleteRiwayatCuti
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_cuti)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_cuti); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatCuti{}).Where("nip = ? and tmulai = ? and nsk = ?", riwayat_cuti.Nip, riwayat_cuti.Tmulai, riwayat_cuti.Nsk).Delete(&riwayat_cuti)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_cuti,
		"statucCode": http.StatusOK,
	})
}
