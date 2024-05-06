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

func GetRiwayatHukdis(searchString model.SearchRiwayatHukdis) (riwayat_hukdis []model.RiwayatHukdis, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatHukdis{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_hukdis.nip = ?", searchString.Nip)
	}

	//Jhukum
	if searchString.Jhukum != "" {
		result = result.Where("master_riwayat_hukdis.jhukum = ?", searchString.Jhukum)
	}

	//Jhukum Bkn
	if searchString.JhukumBkn != "" {
		result = result.Where("master_riwayat_hukdis.jhukum_bkn = ?", searchString.JhukumBkn)
	}

	//Tmt
	if searchString.Tmt.IsZero() != true {
		result = result.Where("master_riwayat_hukdis.tmt = ?", searchString.Tmt)
	}

	result = result.Preload("JenisHukdis").Find(&riwayat_hukdis)

	return

}

func FindRiwayatHukdis(c echo.Context) error {
	var searchString model.SearchRiwayatHukdis
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_hukdis, result := GetRiwayatHukdis(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_hukdis,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetRiwayatHukdisByNipJhukumTmt(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_hukdis model.RiwayatHukdis
	nip := c.Param("nip")
	jhukum := c.Param("jhukum")
	tmt := c.Param("tmt")

	result := db.Model(&model.RiwayatHukdis{}).Where("nip =? and jhukum = ? and tmt =? ", nip, jhukum, tmt).Scan(&riwayat_hukdis)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_hukdis,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateRiwayatHukdis(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_hukdis model.RiwayatHukdis
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_hukdis)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_hukdis); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_hukdis.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_hukdis)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_hukdis,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatHukdis(c echo.Context) error {
	db, _ := model.CreateCon()

	//var riwayat_hukdis model.RiwayatHukdis
	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatHukdis `json:"refdata"`
		Data model.RiwayatHukdis       `json:"data"`
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
	Tmt := (reqData.Ref.Tmt).Format("2006-01-02")
	result := db.Model(&model.RiwayatHukdis{}).Where("nip = ? and jhukum = ? and tmt = ?", reqData.Ref.Nip, reqData.Ref.Jhukum, Tmt).Updates(&reqData.Data)

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

func DeleteRiwayatHukdis(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_hukdis model.DeleteRiwayatHukdis
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_hukdis)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_hukdis); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatHukdis{}).Where("nip = ? and jhukum = ? and tmt = ?", riwayat_hukdis.Nip, riwayat_hukdis.Jhukum, riwayat_hukdis.Tmt).Delete(&riwayat_hukdis)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_hukdis,
		"statucCode": http.StatusOK,
	})
}
