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

func GetWilayah(searchString model.SearchWilayah) (wilayah []model.Wilayah, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Wilayah{})

	//kwil
	if searchString.Kwil != "" {
		result = result.Where("master_wilayah.kwil = ?", searchString.Kwil)
	}

	//Nwil
	if searchString.Nwil != "" {
		nama := strings.TrimSpace(searchString.Nwil)
		str := []string{"%", nama, "%"}
		result = result.Where("master_wilayah.nwil LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&wilayah)

	return

}

func FindWilayah(c echo.Context) error {
	var searchString model.SearchWilayah
	//var wilayah []model.Wilayah

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	wilayah, result := GetWilayah(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       wilayah,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateWilayah(c echo.Context) error {
	db, _ := model.CreateCon()

	var wilayah model.Wilayah
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&wilayah)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(wilayah); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//wilayah.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&wilayah)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       wilayah,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateWilayah(c echo.Context) error {
	db, _ := model.CreateCon()

	var wilayah model.Wilayah
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&wilayah)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(wilayah); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//wilayah.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Wilayah{}).Where("kwil = ?", wilayah.Kwil).Updates(&wilayah)

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
		"data":       wilayah,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteWilayah(c echo.Context) error {
	db, _ := model.CreateCon()

	var wilayah model.DeleteWilayah
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&wilayah)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(wilayah); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Wilayah{}).Where("kwil = ?", wilayah.Kwil).Delete(&wilayah)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       wilayah,
		"statucCode": http.StatusOK,
	})
}
