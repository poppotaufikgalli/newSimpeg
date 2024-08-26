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

func GetRiwayatTubel(searchString model.SearchRiwayatTubel) (riwayat_tubel []model.RiwayatTubel, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatTubel{})

	//id
	if searchString.ID > 0 {
		result = result.Where("master_riwayat_tubel.id = ?", searchString.ID)
	}

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_tubel.nip = ?", searchString.Nip)
	}

	//ktpu
	if searchString.Ktpu != "" {
		result = result.Where("master_riwayat_tubel.ktpu = ?", searchString.Ktpu)
	}

	//nsek
	if searchString.Nsek != "" {
		nama := strings.TrimSpace(searchString.Nsek)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_tubel.nsek LIKE ?", strings.Join(str, ""))
	}

	//tempat
	if searchString.Tempat != "" {
		nama := strings.TrimSpace(searchString.Tempat)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_tubel.tempat LIKE ?", strings.Join(str, ""))
	}

	result = result.Preload("TingkatPendidikan").Order("id asc").Find(&riwayat_tubel)

	return

}

func FindRiwayatTubel(c echo.Context) error {
	var searchString model.SearchRiwayatTubel
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_tubel, result := GetRiwayatTubel(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_tubel,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetRiwayatTubelByNipId(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_tubel model.RiwayatTubel
	nip := c.Param("nip")
	id := c.Param("id")

	result := db.Model(&model.RiwayatTubel{}).Where("nip =? and id =?", nip, id).Scan(&riwayat_tubel)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_tubel,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateRiwayatTubel(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_tubel model.RiwayatTubel
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_tubel)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_tubel); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_tubel.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_tubel)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_tubel,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatTubel(c echo.Context) error {
	db, _ := model.CreateCon()

	//var riwayat_tubel model.RiwayatTubel
	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatTubel `json:"refdata"`
		Data model.RiwayatTubel       `json:"data"`
	}

	var reqData ToUpdateData

	err := json.NewDecoder(c.Request().Body).Decode(&reqData)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(reqData.Data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	reqData.Data.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.RiwayatTubel{}).Where("id = ? and nip = ? and ktpu = ?", reqData.Ref.ID, reqData.Ref.Nip, reqData.Ref.Ktpu).Updates(&reqData.Data)

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
		"data":       reqData.Data,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatTubel(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_tubel model.DeleteRiwayatTubel
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_tubel)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_tubel); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatTubel{}).Where("id = ? and nip = ? and ktpu = ?", riwayat_tubel.ID, riwayat_tubel.Nip, riwayat_tubel.Ktpu).Delete(&riwayat_tubel)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_tubel,
		"statucCode": http.StatusOK,
	})
}
