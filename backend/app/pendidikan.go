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

func GetPendidikan(searchString model.SearchPendidikan) (pendidikan []model.Pendidikan, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Pendidikan{}).Debug()

	//id
	if searchString.Id != "" {
		result = result.Where("master_pendidikan.id = ?", searchString.Id)
	}

	//ref bkn
	if searchString.RefBkn != "" {
		result = result.Where("master_pendidikan.ref_bkn = ?", searchString.RefBkn)
	}

	//tk pendidikan id
	if len(searchString.TkPendidikanId) > 0 {
		result = result.Where("master_pendidikan.tk_pendidikan_id IN (?)", searchString.TkPendidikanId)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_pendidikan.nama LIKE ?", strings.Join(str, ""))
	}

	//pembina id
	if len(searchString.Status) > 0 {
		result = result.Where("master_pendidikan.status IN (?)", searchString.Status)
	}

	result = result.Find(&pendidikan)

	return
}

func FindPendidikan(c echo.Context) error {
	var searchString model.SearchPendidikan

	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		searchString.Status = []int{1}
	}

	pendidikan, result := GetPendidikan(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pendidikan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreatePendidikan(c echo.Context) error {
	db, _ := model.CreateCon()

	var pendidikan model.Pendidikan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pendidikan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pendidikan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	pendidikan.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&pendidikan)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pendidikan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdatePendidikan(c echo.Context) error {
	db, _ := model.CreateCon()

	var pendidikan model.Pendidikan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pendidikan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pendidikan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	pendidikan.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Pendidikan{}).Where("id = ?", pendidikan.Id).Updates(&pendidikan)

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
		"data":       pendidikan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeletePendidikan(c echo.Context) error {
	db, _ := model.CreateCon()

	var pendidikan model.DeletePendidikan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pendidikan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pendidikan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Pendidikan{}).Where("id = ?", pendidikan.Id).Delete(&pendidikan)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pendidikan,
		"statucCode": http.StatusOK,
	})
}
