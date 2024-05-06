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

func GetUsers(searchString model.SearchUsers) (users []model.Users, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Users{})

	//id
	if searchString.Id != "" {
		result = result.Where("users.id = ?", searchString.Id)
	}

	//nip
	if searchString.Nip != "" {
		result = result.Where("users.nip = ?", searchString.Nip)
	}

	result = result.Find(&users)

	return
}

func FindUsers(c echo.Context) error {
	var searchString model.SearchUsers

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	users, result := GetUsers(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       users,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetUsersById(c echo.Context) error {
	db, _ := model.CreateCon()
	var user model.Users
	id := c.Param("id")

	result := db.Model(&model.Users{}).Where("id = ?", id).Scan(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       user,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateUsers(c echo.Context) error {
	db, _ := model.CreateCon()

	var users model.Users
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&users)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(users); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//users.CreatedBy = fmt.Sprint(c.Get("nip"))

	result := db.Create(&users)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       users,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateUsers(c echo.Context) error {
	db, _ := model.CreateCon()

	var users model.Users
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&users)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(users); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//users.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Users{}).Where("id = ?", users.Id).Updates(&users)

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
		"data":       users,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteUsers(c echo.Context) error {
	db, _ := model.CreateCon()

	var users model.DeleteUsers
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&users)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(users); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Users{}).Where("id = ?", users.Id).Delete(&users)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       users,
		"statucCode": http.StatusOK,
	})
}
