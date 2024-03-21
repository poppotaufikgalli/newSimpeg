package app

import (
	_ "database/sql"
	"encoding/json"
	//"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	model "newSimpegAPI/model"
	"strings"
)

func GetReferensiNilaiKuadran(searchString model.SearchReferensiNilaiKuadran) (referensi_nilai_kuadran []model.ReferensiNilaiKuadran, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.ReferensiNilaiKuadran{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_referensi_nilai_kuadran.id = ?", searchString.Id)
	}

	//Jns
	if searchString.Jns > 0 {
		result = result.Where("master_referensi_nilai_kuadran.jns = ?", searchString.Jns)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_referensi_nilai_kuadran.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&referensi_nilai_kuadran)

	return
}

func FindReferensiNilaiKuadran(c echo.Context) error {
	var searchString model.SearchReferensiNilaiKuadran

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	referensi_nilai_kuadran, result := GetReferensiNilaiKuadran(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       referensi_nilai_kuadran,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateReferensiNilaiKuadran(c echo.Context) error {
	db, _ := model.CreateCon()

	var referensi_nilai_kuadran model.ReferensiNilaiKuadran
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&referensi_nilai_kuadran)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(referensi_nilai_kuadran); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//referensi_nilai_kuadran.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&referensi_nilai_kuadran)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       referensi_nilai_kuadran,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateReferensiNilaiKuadran(c echo.Context) error {
	db, _ := model.CreateCon()

	var referensi_nilai_kuadran model.ReferensiNilaiKuadran
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&referensi_nilai_kuadran)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(referensi_nilai_kuadran); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//referensi_nilai_kuadran.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.ReferensiNilaiKuadran{}).Where("id = ? and jns = ?", referensi_nilai_kuadran.Id, referensi_nilai_kuadran.Jns).Updates(&referensi_nilai_kuadran)

	//fmt.Println(result.RowsAffected)

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
		"data":       referensi_nilai_kuadran,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteReferensiNilaiKuadran(c echo.Context) error {
	db, _ := model.CreateCon()

	var referensi_nilai_kuadran model.DeleteReferensiNilaiKuadran
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&referensi_nilai_kuadran)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(referensi_nilai_kuadran); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.ReferensiNilaiKuadran{}).Where("id = ? and jns = ?", referensi_nilai_kuadran.Id, referensi_nilai_kuadran.Jns).Delete(&referensi_nilai_kuadran)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       referensi_nilai_kuadran,
		"statucCode": http.StatusOK,
	})
}
