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

func GetJenisPenghargaan(searchString model.SearchJenisPenghargaan) (jenis_penghargaan []model.JenisPenghargaan, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisPenghargaan{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_penghargaan.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_penghargaan.nama LIKE ?", strings.Join(str, ""))
	}

	//id
	if searchString.RefSimpeg != "" {
		result = result.Where("master_jenis_penghargaan.ref_simpeg = ?", searchString.RefSimpeg)
	}

	result = result.Find(&jenis_penghargaan)

	return
}

func FindJenisPenghargaan(c echo.Context) error {
	var searchString model.SearchJenisPenghargaan

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jenis_penghargaan, result := GetJenisPenghargaan(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_penghargaan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisPenghargaan(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_penghargaan model.JenisPenghargaan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_penghargaan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_penghargaan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_penghargaan.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_penghargaan)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_penghargaan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisPenghargaan(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_penghargaan model.JenisPenghargaan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_penghargaan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_penghargaan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_penghargaan.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisPenghargaan{}).Where("id = ?", jenis_penghargaan.Id).Updates(&jenis_penghargaan)

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
		"data":       jenis_penghargaan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisPenghargaan(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_penghargaan model.DeleteJenisPenghargaan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_penghargaan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_penghargaan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisPenghargaan{}).Where("id = ?", jenis_penghargaan.Id).Delete(&jenis_penghargaan)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_penghargaan,
		"statucCode": http.StatusOK,
	})
}
