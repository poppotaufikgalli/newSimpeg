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

func GetJabatanNegara(searchString model.SearchJabatanNegara) (jabatan_negara []model.JabatanNegara, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JabatanNegara{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jabatan_negara.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jabatan_negara.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&jabatan_negara)

	return
}

func FindJabatanNegara(c echo.Context) error {
	var searchString model.SearchJabatanNegara

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jabatan_negara, result := GetJabatanNegara(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_negara,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJabatanNegara(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_negara model.JabatanNegara
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_negara)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_negara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jabatan_negara.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jabatan_negara)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_negara,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJabatanNegara(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_negara model.JabatanNegara
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_negara)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_negara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jabatan_negara.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JabatanNegara{}).Where("id = ?", jabatan_negara.Id).Updates(&jabatan_negara)

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
		"data":       jabatan_negara,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJabatanNegara(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_negara model.DeleteJabatanNegara
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_negara)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_negara); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JabatanNegara{}).Where("id = ?", jabatan_negara.Id).Delete(&jabatan_negara)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_negara,
		"statucCode": http.StatusOK,
	})
}
