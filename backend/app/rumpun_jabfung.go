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

func GetRumpunJabfung(searchString model.SearchRumpunJabfung) (rumpun_jabfung []model.RumpunJabfung, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RumpunJabfung{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_rumpun_jabfung.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_rumpun_jabfung.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&rumpun_jabfung)

	return
}

func FindRumpunJabfung(c echo.Context) error {
	var searchString model.SearchRumpunJabfung

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	rumpun_jabfung, result := GetRumpunJabfung(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       rumpun_jabfung,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateRumpunJabfung(c echo.Context) error {
	db, _ := model.CreateCon()

	var rumpun_jabfung model.RumpunJabfung
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&rumpun_jabfung)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(rumpun_jabfung); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	rumpun_jabfung.CreatedBy = fmt.Sprint(c.Get("nip"))

	result := db.Create(&rumpun_jabfung)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       rumpun_jabfung,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRumpunJabfung(c echo.Context) error {
	db, _ := model.CreateCon()

	var rumpun_jabfung model.RumpunJabfung
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&rumpun_jabfung)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(rumpun_jabfung); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	rumpun_jabfung.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RumpunJabfung{}).Where("id = ?", rumpun_jabfung.Id).Updates(&rumpun_jabfung)

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
		"data":       rumpun_jabfung,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRumpunJabfung(c echo.Context) error {
	db, _ := model.CreateCon()

	var rumpun_jabfung model.DeleteRumpunJabfung
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&rumpun_jabfung)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(rumpun_jabfung); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RumpunJabfung{}).Where("id = ?", rumpun_jabfung.Id).Delete(&rumpun_jabfung)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       rumpun_jabfung,
		"statucCode": http.StatusOK,
	})
}
