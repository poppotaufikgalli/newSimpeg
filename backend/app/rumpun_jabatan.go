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

func GetRumpunJabatan(searchString model.SearchRumpunJabatan) (rumpun_jabatan []model.RumpunJabatan, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RumpunJabatan{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_rumpun_jabatan.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_rumpun_jabatan.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&rumpun_jabatan)

	return
}

func FindRumpunJabatan(c echo.Context) error {
	var searchString model.SearchRumpunJabatan

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	rumpun_jabatan, result := GetRumpunJabatan(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       rumpun_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateRumpunJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var rumpun_jabatan model.RumpunJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&rumpun_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(rumpun_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	rumpun_jabatan.CreatedBy = fmt.Sprint(c.Get("nip"))

	result := db.Create(&rumpun_jabatan)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       rumpun_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRumpunJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var rumpun_jabatan model.RumpunJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&rumpun_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(rumpun_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	rumpun_jabatan.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RumpunJabatan{}).Where("id = ?", rumpun_jabatan.Id).Updates(&rumpun_jabatan)

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
		"data":       rumpun_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRumpunJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var rumpun_jabatan model.DeleteRumpunJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&rumpun_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(rumpun_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RumpunJabatan{}).Where("id = ?", rumpun_jabatan.Id).Delete(&rumpun_jabatan)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       rumpun_jabatan,
		"statucCode": http.StatusOK,
	})
}
