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

func GetRefJenpeg(searchString model.SearchRefJenpeg) (ref_jenpeg []model.RefJenpeg, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RefJenpeg{})

	//id
	if searchString.Kjpeg != 0 {
		result = result.Where("master_ref_jenpeg.kjpeg = ?", searchString.Kjpeg)
	}

	//nama
	if searchString.Njpeg != "" {
		nama := strings.TrimSpace(searchString.Njpeg)
		str := []string{"%", nama, "%"}
		result = result.Where("master_ref_jenpeg.njpeg LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&ref_jenpeg)

	return

}

func FindRefJenpeg(c echo.Context) error {
	var searchString model.SearchRefJenpeg
	//var ref_jenpeg []model.RefJenpeg

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	ref_jenpeg, result := GetRefJenpeg(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       ref_jenpeg,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateRefJenpeg(c echo.Context) error {
	db, _ := model.CreateCon()

	var ref_jenpeg model.RefJenpeg
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&ref_jenpeg)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(ref_jenpeg); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//ref_jenpeg.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&ref_jenpeg)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       ref_jenpeg,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRefJenpeg(c echo.Context) error {
	db, _ := model.CreateCon()

	var ref_jenpeg model.RefJenpeg
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&ref_jenpeg)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(ref_jenpeg); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//ref_jenpeg.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RefJenpeg{}).Where("kjpeg = ?", ref_jenpeg.Kjpeg).Updates(&ref_jenpeg)

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
		"data":       ref_jenpeg,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRefJenpeg(c echo.Context) error {
	db, _ := model.CreateCon()

	var ref_jenpeg model.DeleteRefJenpeg
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&ref_jenpeg)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(ref_jenpeg); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RefJenpeg{}).Where("kjpeg = ?", ref_jenpeg.Kjpeg).Delete(&ref_jenpeg)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       ref_jenpeg,
		"statucCode": http.StatusOK,
	})
}
