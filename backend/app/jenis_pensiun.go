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

func GetJenisPensiun(searchString model.SearchJenisPensiun) (jenis_pensiun []model.JenisPensiun, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisPensiun{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_pensiun.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_pensiun.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&jenis_pensiun)

	return
}

func FindJenisPensiun(c echo.Context) error {
	var searchString model.SearchJenisPensiun

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jenis_pensiun, result := GetJenisPensiun(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_pensiun,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisPensiun(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_pensiun model.JenisPensiun
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_pensiun)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_pensiun); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_pensiun.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_pensiun)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_pensiun,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisPensiun(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_pensiun model.JenisPensiun
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_pensiun)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_pensiun); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_pensiun.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisPensiun{}).Where("id = ?", jenis_pensiun.Id).Updates(&jenis_pensiun)

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
		"data":       jenis_pensiun,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisPensiun(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_pensiun model.DeleteJenisPensiun
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_pensiun)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_pensiun); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisPensiun{}).Where("id = ?", jenis_pensiun.Id).Delete(&jenis_pensiun)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_pensiun,
		"statucCode": http.StatusOK,
	})
}
