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

func GetJenisTugasMutasi(searchString model.SearchJenisTugasMutasi) (jenis_tugas_mutasi []model.JenisTugasMutasi, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JenisTugasMutasi{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jenis_tugas_mutasi.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jenis_tugas_mutasi.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Order("cast(id as unsigned) asc").Find(&jenis_tugas_mutasi)

	return
}

func FindJenisTugasMutasi(c echo.Context) error {
	var searchString model.SearchJenisTugasMutasi
	//var jenis_tugas_mutasi []model.JenisTugasMutasi

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jenis_tugas_mutasi, result := GetJenisTugasMutasi(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_tugas_mutasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func SearchJenisTugasMutasiById(c echo.Context) error {
	db, _ := model.CreateCon()
	var jenis_tugas_mutasi model.JenisTugasMutasi
	id := c.Param("id")

	result := db.Model(&model.JenisTugasMutasi{}).Where("id = ?", id).Scan(&jenis_tugas_mutasi)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_tugas_mutasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateJenisTugasMutasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_tugas_mutasi model.JenisTugasMutasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_tugas_mutasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_tugas_mutasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_tugas_mutasi.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jenis_tugas_mutasi)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_tugas_mutasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJenisTugasMutasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_tugas_mutasi model.JenisTugasMutasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_tugas_mutasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_tugas_mutasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jenis_tugas_mutasi.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JenisTugasMutasi{}).Where("id = ?", jenis_tugas_mutasi.Id).Updates(&jenis_tugas_mutasi)

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
		"data":       jenis_tugas_mutasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJenisTugasMutasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var jenis_tugas_mutasi model.DeleteJenisTugasMutasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jenis_tugas_mutasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jenis_tugas_mutasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JenisTugasMutasi{}).Where("id = ?", jenis_tugas_mutasi.Id).Delete(&jenis_tugas_mutasi)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jenis_tugas_mutasi,
		"statucCode": http.StatusOK,
	})
}
