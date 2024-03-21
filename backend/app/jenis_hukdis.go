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

func GetJenisHukdis(searchString model.SearchJenisHukdis) (jenis_hukdis []model.JenisHukdis, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisHukdis{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_hukdis.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_hukdis.nama LIKE ?", strings.Join(str, ""))
	}

	//jenis tingkat hukuman id
	if len(searchString.JenisTingkatHukumanId) > 0 {
		result = result.Where("master_jenis_hukdis.jenis_tingkat_hukuman_id IN (?)", searchString.JenisTingkatHukumanId)
	}

	//id ref bkn
	if len(searchString.IdRefBkn) > 0 {
		result = result.Where("master_jenis_hukdis.id_ref_bkn IN (?)", searchString.IdRefBkn)
	}

	result = result.Find(&jenis_hukdis)

	return
}

func FindJenisHukdis(c echo.Context) error {
	var searchString model.SearchJenisHukdis

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jenis_hukdis, result := GetJenisHukdis(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_hukdis,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateJenisHukdis(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_hukdis model.JenisHukdis
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_hukdis)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_hukdis); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_hukdis.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_hukdis)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_hukdis,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisHukdis(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_hukdis model.JenisHukdis
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_hukdis)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_hukdis); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_hukdis.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisHukdis{}).Where("id = ?", jenis_hukdis.Id).Updates(&jenis_hukdis)

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
		"data":       jenis_hukdis,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisHukdis(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_hukdis model.DeleteJenisHukdis
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_hukdis)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_hukdis); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisHukdis{}).Where("id = ?", jenis_hukdis.Id).Delete(&jenis_hukdis)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_hukdis,
		"statucCode": http.StatusOK,
	})
}
