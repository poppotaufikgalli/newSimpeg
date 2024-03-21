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

func GetKpe(searchString model.SearchKpe) (kpe []model.Kpe, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Kpe{})

	//id
	if searchString.Kkpe != "" {
		result = result.Where("master_kpe.kkpe = ?", searchString.Kkpe)
	}

	//nama
	if searchString.Nkpe != "" {
		nama := strings.TrimSpace(searchString.Nkpe)
		str := []string{"%", nama, "%"}
		result = result.Where("master_kpe.nkpe LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&kpe)

	return
}

func FindKpe(c echo.Context) error {
	var searchString model.SearchKpe

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	kpe, result := GetKpe(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       kpe,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateKpe(c echo.Context) error {
	db, _ := model.CreateCon()

	var kpe model.Kpe
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&kpe)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(kpe); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//kpe.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&kpe)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       kpe,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateKpe(c echo.Context) error {
	db, _ := model.CreateCon()

	var kpe model.Kpe
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&kpe)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(kpe); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//kpe.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Kpe{}).Where("kkpe = ?", kpe.Kkpe).Updates(&kpe)

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
		"data":       kpe,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteKpe(c echo.Context) error {
	db, _ := model.CreateCon()

	var kpe model.DeleteKpe
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&kpe)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(kpe); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Kpe{}).Where("kkpe = ?", kpe.Kkpe).Delete(&kpe)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       kpe,
		"statucCode": http.StatusOK,
	})
}
