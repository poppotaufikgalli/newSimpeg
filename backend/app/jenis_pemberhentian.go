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

func GetJenisPemberhentian(searchString model.SearchJenisPemberhentian) (jenis_pemberhentian []model.JenisPemberhentian, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisPemberhentian{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_pemberhentian.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_pemberhentian.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&jenis_pemberhentian)

	return
}

func FindJenisPemberhentian(c echo.Context) error {
	var searchString model.SearchJenisPemberhentian

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jenis_pemberhentian, result := GetJenisPemberhentian(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_pemberhentian,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisPemberhentian(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_pemberhentian model.JenisPemberhentian
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_pemberhentian)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_pemberhentian); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_pemberhentian.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_pemberhentian)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_pemberhentian,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisPemberhentian(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_pemberhentian model.JenisPemberhentian
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_pemberhentian)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_pemberhentian); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_pemberhentian.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisPemberhentian{}).Where("id = ?", jenis_pemberhentian.Id).Updates(&jenis_pemberhentian)

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
		"data":       jenis_pemberhentian,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisPemberhentian(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_pemberhentian model.DeleteJenisPemberhentian
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_pemberhentian)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_pemberhentian); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisPemberhentian{}).Where("id = ?", jenis_pemberhentian.Id).Delete(&jenis_pemberhentian)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_pemberhentian,
		"statucCode": http.StatusOK,
	})
}
