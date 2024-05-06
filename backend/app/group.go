package app

import (
	_ "database/sql"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	model "newSimpegAPI/model"
)

func GetGroup(searchString model.SearchGroup) (group []model.Group, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Group{})

	//id
	if searchString.Id != "" {
		result = result.Where("group.id = ?", searchString.Id)
	}

	//nama
	if searchString.Nama != "" {
		result = result.Where("group.nama = ?", searchString.Nama)
	}

	result = result.Preload("Users").Find(&group)

	return
}

func FindGroup(c echo.Context) error {
	var searchString model.SearchGroup

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	group, result := GetGroup(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       group,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetGroupById(c echo.Context) error {
	db, _ := model.CreateCon()
	var user model.Group
	id := c.Param("id")

	result := db.Model(&model.Group{}).Where("id = ?", id).Scan(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       user,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateGroup(c echo.Context) error {
	db, _ := model.CreateCon()

	var group model.Group
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&group)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(group); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//group.CreatedBy = fmt.Sprint(c.Get("nip"))

	result := db.Create(&group)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       group,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateGroup(c echo.Context) error {
	db, _ := model.CreateCon()

	var group model.Group
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&group)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(group); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//group.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Group{}).Where("id = ?", group.Id).Updates(&group)

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
		"data":       group,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteGroup(c echo.Context) error {
	db, _ := model.CreateCon()

	var group model.DeleteGroup
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&group)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(group); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Group{}).Where("id = ?", group.Id).Delete(&group)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       group,
		"statucCode": http.StatusOK,
	})
}
