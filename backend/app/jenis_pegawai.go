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

func GetJenisPegawai(searchString model.SearchJenisPegawai) (jenis_pegawai []model.JenisPegawai, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisPegawai{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_pegawai.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_pegawai.nama LIKE ?", strings.Join(str, ""))
	}

	//status
	if len(searchString.Status) > 0 {
		result = result.Where("master_jenis_pegawai.status IN (?)", searchString.Status)
	}

	result = result.Find(&jenis_pegawai)

	return
}

func FindJenisPegawai(c echo.Context) error {
	var searchString model.SearchJenisPegawai
	//var jenis_pegawai []model.JenisPegawai

	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		searchString.Status = []int{1}
	}

	jenis_pegawai, result := GetJenisPegawai(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_pegawai,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisPegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_pegawai model.JenisPegawai
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_pegawai)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_pegawai.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_pegawai)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_pegawai,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisPegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_pegawai model.JenisPegawai
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_pegawai)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_pegawai.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisPegawai{}).Where("id = ?", jenis_pegawai.Id).Updates(&jenis_pegawai)

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
		"data":       jenis_pegawai,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisPegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_pegawai model.DeleteJenisPegawai
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_pegawai)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisPegawai{}).Where("id = ?", jenis_pegawai.Id).Delete(&jenis_pegawai)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_pegawai,
		"statucCode": http.StatusOK,
	})
}
