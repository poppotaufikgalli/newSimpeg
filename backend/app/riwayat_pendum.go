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

func GetRiwayatPendum(searchString model.SearchRiwayatPendum) (riwayat_pendum []model.RiwayatPendum, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatPendum{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_pendum.nip = ?", searchString.Nip)
	}

	//ktpu
	if searchString.Ktpu != "" {
		result = result.Where("master_riwayat_pendum.ktpu = ?", searchString.Ktpu)
	}

	//kjur
	if searchString.Kjur != "" {
		result = result.Where("master_riwayat_pendum.kjur = ?", searchString.Kjur)
	}

	//nsek
	if searchString.Nsek != "" {
		nama := strings.TrimSpace(searchString.Nsek)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_organisasi.nsek LIKE ?", strings.Join(str, ""))
	}

	//tempat
	if searchString.Tempat != "" {
		nama := strings.TrimSpace(searchString.Tempat)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_organisasi.tempat LIKE ?", strings.Join(str, ""))
	}

	result = result.Preload("TingkatPendidikan").Preload("Pendidikan").Order("cast(ktpu as unsigned)").Find(&riwayat_pendum)

	return

}

func FindRiwayatPendum(c echo.Context) error {
	var searchString model.SearchRiwayatPendum
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_pendum, result := GetRiwayatPendum(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pendum,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetRiwayatPendumByNipKtpuKjur(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_pendum model.RiwayatPendum
	nip := c.Param("nip")
	ktpu := c.Param("ktpu")
	kjur := c.Param("kjur")

	result := db.Model(&model.RiwayatPendum{}).Where("nip =? and ktpu =? and kjur =?", nip, ktpu, kjur).Scan(&riwayat_pendum)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pendum,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateRiwayatPendum(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_pendum model.RiwayatPendum
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_pendum)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_pendum); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_pendum.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_pendum)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pendum,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatPendum(c echo.Context) error {
	db, _ := model.CreateCon()

	//var riwayat_pendum model.RiwayatPendum
	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatPendum `json:"refdata"`
		Data model.RiwayatPendum       `json:"data"`
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
	result := db.Model(&model.RiwayatPendum{}).Where("nip = ? and ktpu = ? and kjur = ?", reqData.Ref.Nip, reqData.Ref.Ktpu, reqData.Ref.Kjur).Updates(&reqData.Data)

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

func DeleteRiwayatPendum(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_pendum model.DeleteRiwayatPendum
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_pendum)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_pendum); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatPendum{}).Where("nip = ? and ktpu = ? and kjur = ?", riwayat_pendum.Nip, riwayat_pendum.Ktpu, riwayat_pendum.Kjur).Delete(&riwayat_pendum)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pendum,
		"statucCode": http.StatusOK,
	})
}
