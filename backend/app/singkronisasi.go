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
)

func GetSingkronisasi(searchString model.SearchSingkronisasi) (singkronisasi []model.Singkronisasi, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Singkronisasi{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_singkronisasi.id = ?", searchString.Id)
	}

	//host
	if searchString.Host != "" {
		result = result.Where("master_singkronisasi.host = ?", searchString.Host)
	}

	result = result.Find(&singkronisasi)

	return
}

func FindSingkronisasi(c echo.Context) error {
	var searchString model.SearchSingkronisasi

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	singkronisasi, result := GetSingkronisasi(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       singkronisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetSingkronisasiByHost(c echo.Context) error {
	db, _ := model.CreateCon()

	var singkronisasi model.Singkronisasi
	host := c.Param("host")

	result := db.Model(&model.Singkronisasi{}).Where("host =?", host).Scan(&singkronisasi)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       singkronisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetSingkronisasiExpiredTimeByHost(c echo.Context) error {
	db, _ := model.CreateCon()

	var singkronisasi model.Singkronisasi
	host := c.Param("host")

	result := db.Model(&model.Singkronisasi{}).Select("token_sso_expired", "token_apimanager_expired").Where("host =?", host).Scan(&singkronisasi)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       singkronisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateSingkronisasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var singkronisasi model.Singkronisasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&singkronisasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(singkronisasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	singkronisasi.CreatedBy = fmt.Sprint(c.Get("nip"))

	result := db.Create(&singkronisasi)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       singkronisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateSingkronisasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var singkronisasi model.Singkronisasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&singkronisasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(singkronisasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	singkronisasi.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Singkronisasi{}).Where("id = ?", singkronisasi.Id).Updates(&singkronisasi)

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
		"data":       singkronisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteSingkronisasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var singkronisasi model.DeleteSingkronisasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&singkronisasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(singkronisasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Singkronisasi{}).Where("id = ?", singkronisasi.Id).Delete(&singkronisasi)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       singkronisasi,
		"statucCode": http.StatusOK,
	})
}
