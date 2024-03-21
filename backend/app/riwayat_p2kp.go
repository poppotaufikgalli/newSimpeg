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

func GetRiwayatP2kp(searchString model.SearchRiwayatP2kp) (riwayat_p2kp []model.RiwayatP2kp, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatP2kp{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_p2kp.nip = ?", searchString.Nip)
	}

	//thnilai
	if searchString.Thnilai > 0 {
		result = result.Where("master_riwayat_p2kp.thnilai = ?", searchString.Thnilai)
	}

	//tmulai
	if searchString.Tmulai.IsZero() != true {
		result = result.Where("master_riwayat_p2kp.tmulai = ?", searchString.Tmulai)
	}

	//tselesai
	if searchString.Tselesai.IsZero() != true {
		result = result.Where("master_riwayat_p2kp.tselesai = ?", searchString.Tselesai)
	}

	result = result.Find(&riwayat_p2kp)

	return

}

func FindRiwayatP2kp(c echo.Context) error {
	var searchString model.SearchRiwayatP2kp
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_p2kp, result := GetRiwayatP2kp(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_p2kp,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateRiwayatP2kp(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_p2kp model.RiwayatP2kp
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_p2kp)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_p2kp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_p2kp.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_p2kp)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_p2kp,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatP2kp(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_p2kp model.RiwayatP2kp
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_p2kp)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_p2kp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_p2kp.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RiwayatP2kp{}).Where("nip = ? and thnilai = ? and tmulai = ? and tselesai = ?", riwayat_p2kp.Nip, riwayat_p2kp.Thnilai, riwayat_p2kp.Tmulai, riwayat_p2kp.Tselesai).Updates(&riwayat_p2kp)

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
		"data":       riwayat_p2kp,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatP2kp(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_p2kp model.DeleteRiwayatP2kp
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_p2kp)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_p2kp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatP2kp{}).Where("nip = ? and thnilai = ? and tmulai = ? and tselesai = ?", riwayat_p2kp.Nip, riwayat_p2kp.Thnilai, riwayat_p2kp.Tmulai, riwayat_p2kp.Tselesai).Delete(&riwayat_p2kp)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_p2kp,
		"statucCode": http.StatusOK,
	})
}
