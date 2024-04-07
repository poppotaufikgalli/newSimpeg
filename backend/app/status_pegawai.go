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

func GetStatusPegawai(searchString model.StatusPegawai) (status_pegawai []model.StatusPegawai, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.StatusPegawai{})

	//id
	if searchString.IdStatusPegawai != 0 {
		result = result.Where("master_status_pegawai.id = ?", searchString.IdStatusPegawai)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_status_pegawai.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Order("cast(kstatus as unsigned) asc").Find(&status_pegawai)

	return

}

func FindStatusPegawai(c echo.Context) error {
	var searchString model.StatusPegawai
	//var status_pegawai []model.StatusPegawai

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	status_pegawai, result := GetStatusPegawai(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       status_pegawai,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateStatusPegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var status_pegawai model.StatusPegawai
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&status_pegawai)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(status_pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//status_pegawai.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&status_pegawai)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       status_pegawai,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateStatusPegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var status_pegawai model.StatusPegawai
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&status_pegawai)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(status_pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//status_pegawai.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.StatusPegawai{}).Where("id_status_pegawai = ?", status_pegawai.IdStatusPegawai).Updates(&status_pegawai)

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
		"data":       status_pegawai,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteStatusPegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var status_pegawai model.DeleteStatusPegawai
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&status_pegawai)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(status_pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.StatusPegawai{}).Where("id_status_pegawai = ?", status_pegawai.IdStatusPegawai).Delete(&status_pegawai)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       status_pegawai,
		"statucCode": http.StatusOK,
	})
}
