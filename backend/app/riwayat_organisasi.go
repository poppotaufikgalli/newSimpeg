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

func GetRiwayatOrganisasi(searchString model.SearchRiwayatOrganisasi) (riwayat_organisasi []model.RiwayatOrganisasi, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatOrganisasi{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_organisasi.nip = ?", searchString.Nip)
	}

	//norg
	if searchString.Norg != "" {
		nama := strings.TrimSpace(searchString.Norg)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_organisasi.norg LIKE ?", strings.Join(str, ""))
	}

	//jborg
	if searchString.Jborg != "" {
		nama := strings.TrimSpace(searchString.Jborg)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_organisasi.jborg LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&riwayat_organisasi)

	return

}

func FindRiwayatOrganisasi(c echo.Context) error {
	var searchString model.SearchRiwayatOrganisasi
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_organisasi, result := GetRiwayatOrganisasi(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_organisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetRiwayatOrganisasiByNipNorgJborg(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_organisasi model.RiwayatOrganisasi
	nip := c.Param("nip")
	norg := c.Param("norg")
	jborg := c.Param("jborg")

	result := db.Model(&model.RiwayatOrganisasi{}).Where("nip =? and norg =? and jborg =? ", nip, norg, jborg).Scan(&riwayat_organisasi)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_organisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateRiwayatOrganisasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_organisasi model.RiwayatOrganisasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_organisasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_organisasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_organisasi.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_organisasi)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_organisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatOrganisasi(c echo.Context) error {
	db, _ := model.CreateCon()

	//var riwayat_organisasi model.RiwayatOrganisasi
	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatOrganisasi `json:"refdata"`
		Data model.RiwayatOrganisasi       `json:"data"`
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
	result := db.Model(&model.RiwayatOrganisasi{}).Where("nip = ? and norg = ? and jborg = ?", reqData.Ref.Nip, reqData.Ref.Norg, reqData.Ref.Jborg).Updates(&reqData.Data)

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

func DeleteRiwayatOrganisasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_organisasi model.DeleteRiwayatOrganisasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_organisasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_organisasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatOrganisasi{}).Where("nip = ? and norg = ? and jborg = ?", riwayat_organisasi.Nip, riwayat_organisasi.Norg, riwayat_organisasi.Jborg).Delete(&riwayat_organisasi)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_organisasi,
		"statucCode": http.StatusOK,
	})
}
