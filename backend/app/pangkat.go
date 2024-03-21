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

func GetPangkat(searchString model.SearchPangkat) (pangkat []model.Pangkat, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Pangkat{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_pangkat.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_pangkat.nama LIKE ?", strings.Join(str, ""))
	}

	//nama pangkat
	if searchString.NamaPangkat != "" {
		nama := strings.TrimSpace(searchString.NamaPangkat)
		str := []string{"%", nama, "%"}
		result = result.Where("master_pangkat.nama_pangkat LIKE ?", strings.Join(str, ""))
	}

	//Ref Simpeg
	if searchString.RefSimpeg != "" {
		result = result.Where("master_pangkat.ref_simpeg = ?", searchString.RefSimpeg)
	}

	//Pajak
	if len(searchString.Pajak) > 0 {
		result = result.Where("master_pangkat.pajak IN (?)", searchString.Pajak)
	}

	result = result.Find(&pangkat)

	return
}

func FindPangkat(c echo.Context) error {
	var searchString model.SearchPangkat

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	pangkat, result := GetPangkat(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreatePangkat(c echo.Context) error {
	db, _ := model.CreateCon()

	var pangkat model.Pangkat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pangkat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pangkat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	pangkat.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&pangkat)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdatePangkat(c echo.Context) error {
	db, _ := model.CreateCon()

	var pangkat model.Pangkat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pangkat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pangkat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	pangkat.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Pangkat{}).Where("id = ?", pangkat.Id).Updates(&pangkat)

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
		"data":       pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeletePangkat(c echo.Context) error {
	db, _ := model.CreateCon()

	var pangkat model.DeletePangkat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pangkat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pangkat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Pangkat{}).Where("id = ?", pangkat.Id).Delete(&pangkat)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pangkat,
		"statucCode": http.StatusOK,
	})
}
