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

func GetPendidikanBkn(searchString model.SearchPendidikanBkn) (pendidikan_bkn []model.PendidikanBkn, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.PendidikanBkn{}).Debug()

	//id
	if searchString.Id != "" {
		result = result.Where("master_pendidikan_bkn.id = ?", searchString.Id)
	}

	//tk PendidikanBkn id
	if len(searchString.TkPendidikanId) > 0 {
		result = result.Where("master_pendidikan_bkn.tk_PendidikanBkn_id IN (?)", searchString.TkPendidikanId)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_pendidikan_bkn.nama LIKE ?", strings.Join(str, ""))
	}

	//pembina id
	if len(searchString.Status) > 0 {
		result = result.Where("master_pendidikan_bkn.status IN (?)", searchString.Status)
	}

	result = result.Find(&pendidikan_bkn)

	return
}

func FindPendidikanBkn(c echo.Context) error {
	var searchString model.SearchPendidikanBkn

	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		searchString.Status = []int{1}
	}

	pendidikan_bkn, result := GetPendidikanBkn(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pendidikan_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func SearchPendidikanBknFilter(c echo.Context) error {
	db, _ := model.CreateCon()
	var pendidikan_bkn []model.PendidikanBkn

	req := model.SearchInput{
		Limit: 10,
	}

	fmt.Println(req)

	if err := c.Bind(&req); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result := db.Model(&model.PendidikanBkn{})

	if req.SearchNama != "" {
		nama := strings.TrimSpace(req.SearchNama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_pendidikan_bkn.nama LIKE ? or master_pendidikan_bkn.id LIKE ?", strings.Join(str, ""), strings.Join(str, ""))
	}

	result = result.Order("nama asc").Debug().Limit(req.Limit).Find(&pendidikan_bkn)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pendidikan_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreatePendidikanBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var pendidikan_bkn model.PendidikanBkn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pendidikan_bkn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pendidikan_bkn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	pendidikan_bkn.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&pendidikan_bkn)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pendidikan_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdatePendidikanBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var pendidikan_bkn model.PendidikanBkn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pendidikan_bkn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pendidikan_bkn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	pendidikan_bkn.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.PendidikanBkn{}).Where("id = ?", pendidikan_bkn.Id).Updates(&pendidikan_bkn)

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
		"data":       pendidikan_bkn,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeletePendidikanBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var pendidikan_bkn model.DeletePendidikanBkn
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pendidikan_bkn)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pendidikan_bkn); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.PendidikanBkn{}).Where("id = ?", pendidikan_bkn.Id).Delete(&pendidikan_bkn)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pendidikan_bkn,
		"statucCode": http.StatusOK,
	})
}
