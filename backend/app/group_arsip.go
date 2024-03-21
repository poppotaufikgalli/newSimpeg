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

func GetGroupArsip(searchString model.SearchGroupArsip) (group_arsip []model.GroupArsip, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.GroupArsip{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_group_arsip.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_group_arsip.nama LIKE ?", strings.Join(str, ""))
	}

	result = result.Find(&group_arsip)

	return
}

func FindGroupArsip(c echo.Context) error {
	var searchString model.SearchGroupArsip

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	group_arsip, result := GetGroupArsip(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       group_arsip,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateGroupArsip(c echo.Context) error {
	db, _ := model.CreateCon()

	var group_arsip model.GroupArsip
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&group_arsip)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(group_arsip); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	group_arsip.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&group_arsip)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       group_arsip,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateGroupArsip(c echo.Context) error {
	db, _ := model.CreateCon()

	var group_arsip model.GroupArsip
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&group_arsip)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(group_arsip); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	group_arsip.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.GroupArsip{}).Where("id = ?", group_arsip.Id).Updates(&group_arsip)

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
		"data":       group_arsip,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteGroupArsip(c echo.Context) error {
	db, _ := model.CreateCon()

	var group_arsip model.DeleteGroupArsip
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&group_arsip)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(group_arsip); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.GroupArsip{}).Where("id = ?", group_arsip.Id).Delete(&group_arsip)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       group_arsip,
		"statucCode": http.StatusOK,
	})
}
