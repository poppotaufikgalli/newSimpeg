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

func GetJenisKp(searchString model.SearchJenisKp) (jenis_kp []model.JenisKp, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisKp{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_kp.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_kp.nama LIKE ?", strings.Join(str, ""))
	}

	//ref Simpeg
	if searchString.RefSimpeg != "" {
		result = result.Where("master_jenis_kp.ref_simpeg = ?", searchString.RefSimpeg)
	}

	//status
	if len(searchString.Status) > 0 {
		result = result.Where("master_jenis_kp.status IN (?)", searchString.Status)
	}

	result = result.Find(&jenis_kp)

	return
}

func FindJenisKp(c echo.Context) error {
	var searchString model.SearchJenisKp

	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		searchString.Status = []int{1}
	}

	jenis_kp, result := GetJenisKp(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_kp,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisKp(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_kp model.JenisKp
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_kp)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_kp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_kp.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_kp)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_kp,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisKp(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_kp model.JenisKp
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_kp)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_kp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_kp.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisKp{}).Where("id = ?", jenis_kp.Id).Updates(&jenis_kp)

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
		"data":       jenis_kp,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisKp(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_kp model.DeleteJenisKp
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_kp)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_kp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisKp{}).Where("id = ?", jenis_kp.Id).Delete(&jenis_kp)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_kp,
		"statucCode": http.StatusOK,
	})
}
