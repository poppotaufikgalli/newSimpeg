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

func GetJenisKawin(searchString model.SearchJenisKawin) (jenis_kawin []model.JenisKawin, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisKawin{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_kawin.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_kawin.nama LIKE ?", strings.Join(str, ""))
	}

	//ref simpeg
	if len(searchString.RefSimpeg) > 0 {
		result = result.Where("master_jenis_kawin.ref_simpeg IN (?)", searchString.RefSimpeg)
	}

	result = result.Find(&jenis_kawin)

	return
}

func FindJenisKawin(c echo.Context) error {
	var searchString model.SearchJenisKawin

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jenis_kawin, result := GetJenisKawin(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_kawin,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisKawin(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_kawin model.JenisKawin
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_kawin)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_kawin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_kawin.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_kawin)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_kawin,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisKawin(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_kawin model.JenisKawin
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_kawin)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_kawin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_kawin.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisKawin{}).Where("id = ?", jenis_kawin.Id).Updates(&jenis_kawin)

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
		"data":       jenis_kawin,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisKawin(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_kawin model.DeleteJenisKawin
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_kawin)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_kawin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisKawin{}).Where("id = ?", jenis_kawin.Id).Delete(&jenis_kawin)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_kawin,
		"statucCode": http.StatusOK,
	})
}
