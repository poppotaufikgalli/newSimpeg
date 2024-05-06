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

func GetRiwayatStlud(searchString model.SearchRiwayatStlud) (riwayat_stlud []model.RiwayatStlud, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatStlud{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_stlud.nip = ?", searchString.Nip)
	}

	//kstlud
	if searchString.Kstlud > 0 {
		result = result.Where("master_riwayat_stlud.kstlud = ?", searchString.Kstlud)
	}

	//Penyelenggara
	if searchString.Penyelenggara != "" {
		nama := strings.TrimSpace(searchString.Penyelenggara)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_stlud.penyelenggara LIKE ?", strings.Join(str, ""))
	}

	//Tstlud
	if searchString.Tstlud.IsZero() != true {
		result = result.Where("master_riwayat_stlud.tstlud = ?", searchString.Tstlud)
	}

	result = result.Preload("Stlud").Find(&riwayat_stlud)

	return

}

func FindRiwayatStlud(c echo.Context) error {
	var searchString model.SearchRiwayatStlud
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_stlud, result := GetRiwayatStlud(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_stlud,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetRiwayatStludByNipKstludTstlud(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_stlud model.RiwayatStlud
	nip := c.Param("nip")
	kstlud := c.Param("kstlud")
	tstlud := c.Param("tstlud")

	//tmulai = (tmulai).Format("2006-01-02")
	result := db.Model(&model.RiwayatStlud{}).Where("nip = ? and kstlud = ? and tstlud =? ", nip, kstlud, tstlud).Scan(&riwayat_stlud)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_stlud,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateRiwayatStlud(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_stlud model.RiwayatStlud
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_stlud)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_stlud); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_stlud.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_stlud)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_stlud,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatStlud(c echo.Context) error {
	db, _ := model.CreateCon()

	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatStlud `json:"refdata"`
		Data model.RiwayatStlud       `json:"data"`
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
	Tstlud := (reqData.Ref.Tstlud).Format("2006-01-02")
	result := db.Model(&model.RiwayatStlud{}).Where("nip = ? and kstlud = ? and tstlud = ?", reqData.Ref.Nip, reqData.Ref.Kstlud, Tstlud).Updates(&reqData.Data)

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

func DeleteRiwayatStlud(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_stlud model.DeleteRiwayatStlud
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_stlud)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_stlud); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatStlud{}).Where("nip = ? and kstlud = ? and tstlud = ?", riwayat_stlud.Nip, riwayat_stlud.Kstlud, riwayat_stlud.Tstlud).Delete(&riwayat_stlud)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_stlud,
		"statucCode": http.StatusOK,
	})
}
