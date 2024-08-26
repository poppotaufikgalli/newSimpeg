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

func GetJabatanSubBkn(searchString model.SearchJabatanSubBkn) (jabatan_sub_bkn []model.JabatanSubBkn, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.JabatanSubBkn{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_jabatan_sub_bkn.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_jabatan_sub_bkn.nama LIKE ?", strings.Join(str, ""))
	}

	//kel jabatan id
	if searchString.KelJabatanId != "" {
		result = result.Where("master_jabatan_sub_bkn.kel_jabatan_id = ?", searchString.KelJabatanId)
	}

	//status
	if len(searchString.Status) > 0 {
		result = result.Where("master_jabatan_sub_bkn.status IN (?)", searchString.Status)
	}

	result = result.Find(&jabatan_sub_bkn)

	return
}

func FindJabatanSubBkn(c echo.Context) error {
	var searchString model.SearchJabatanSubBkn

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	jabatan_sub_bkn, result := GetJabatanSubBkn(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_sub_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func FindJabatanSubBknFilterJft(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_sub_bkn []model.JabatanSubBkn

	id_jft := c.Param("id_jft")

	req := model.SearchInput{
		Limit: 10,
	}

	if err := c.Bind(&req); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result := db.Model(&model.JabatanSubBkn{}).Select("master_jabatan_sub_bkn.id,master_jabatan_sub_bkn.nama,master_jabatan_sub_bkn.kel_jabatan_id").Joins("join master_jabatan_ft_bkn on master_jabatan_sub_bkn.kel_jabatan_id = master_jabatan_ft_bkn.kel_jabatan_id").Where("master_jabatan_ft_bkn.id = ?", id_jft)

	if req.SearchNama != "" {
		nama := strings.TrimSpace(req.SearchNama)
		str := []string{nama, "%"}
		result = result.Where("master_jabatan_sub_bkn.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Order("nama asc").Limit(req.Limit).Find(&jabatan_sub_bkn)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_sub_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateJabatanSubBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_sub_bkn model.JabatanSubBkn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_sub_bkn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_sub_bkn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jabatan_sub_bkn.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&jabatan_sub_bkn)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_sub_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateJabatanSubBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_sub_bkn model.JabatanSubBkn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_sub_bkn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_sub_bkn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	jabatan_sub_bkn.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.JabatanSubBkn{}).Where("id = ?", jabatan_sub_bkn.Id).Updates(&jabatan_sub_bkn)

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
		"data":       jabatan_sub_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteJabatanSubBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var jabatan_sub_bkn model.DeleteJabatanSubBkn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan_sub_bkn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(jabatan_sub_bkn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.JabatanSubBkn{}).Where("id = ?", jabatan_sub_bkn.Id).Delete(&jabatan_sub_bkn)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       jabatan_sub_bkn,
		"statucCode": http.StatusOK,
	})
}
