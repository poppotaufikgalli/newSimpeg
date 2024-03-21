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

func GetDiklatStr(searchString model.SearchDiklatStr) (diklat_str []model.DiklatStr, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.DiklatStr{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_diklat_str.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_diklat_str.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&diklat_str)

	return
}

func FindDiklatStr(c echo.Context) error {
	var searchString model.SearchDiklatStr

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	diklat_str, result := GetDiklatStr(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       diklat_str,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateDiklatStr(c echo.Context) error {
	db, _ := model.CreateCon()

	var diklat_str model.DiklatStr
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&diklat_str)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(diklat_str); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	diklat_str.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&diklat_str)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       diklat_str,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateDiklatStr(c echo.Context) error {
	db, _ := model.CreateCon()

	var diklat_str model.DiklatStr
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&diklat_str)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(diklat_str); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	diklat_str.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.DiklatStr{}).Where("id = ?", diklat_str.Id).Updates(&diklat_str)

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
		"data":       diklat_str,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteDiklatStr(c echo.Context) error {
	db, _ := model.CreateCon()

	var diklat_str model.DeleteDiklatStr
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&diklat_str)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(diklat_str); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.DiklatStr{}).Where("id = ?", diklat_str.Id).Delete(&diklat_str)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       diklat_str,
		"statucCode": http.StatusOK,
	})
}
