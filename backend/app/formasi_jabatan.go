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

func GetFormasiJabatan(searchString model.SearchFormasiJabatan) (formasi_jabatan []model.FormasiJabatan, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.FormasiJabatan{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_formasi_jabatan.id = ?", searchString.Id)
	}

	//id_opd
	if searchString.IdOpd != "" {
		result = result.Where("master_formasi_jabatan.id_opd = ?", searchString.IdOpd)
	}

	//parent_id
	if searchString.ParentId != "" {
		result = result.Where("master_formasi_jabatan.parent_id = ?", searchString.ParentId)
	}

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_formasi_jabatan.nama LIKE ?", strings.Join(str, ""))
	}

	//id_eselon
	if searchString.IdEselon > 0 {
		result = result.Where("master_formasi_jabatan.id_eselon = ?", searchString.IdEselon)
	}

	//id jenis jabatan
	if searchString.IdJenisJabatan > 0 {
		result = result.Where("master_formasi_jabatan.id_jenis_jabatan = ?", searchString.IdJenisJabatan)
	}

	//Kelas Jabatan
	if len(searchString.KelasJabatan) > 0 {
		result = result.Where("master_formasi_jabatan.kelas_jabatan IN (?)", searchString.KelasJabatan)
	}

	//status
	if len(searchString.Status) > 0 {
		result = result.Where("master_formasi_jabatan.status IN (?)", searchString.Status)
	}

	result = result.Find(&formasi_jabatan)

	return

}

func FindFormasiJabatan(c echo.Context) error {
	var searchString model.SearchFormasiJabatan

	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		searchString.Status = []int{1}
	}

	formasi_jabatan, result := GetFormasiJabatan(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       formasi_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateFormasiJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var formasi_jabatan model.FormasiJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&formasi_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(formasi_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	formasi_jabatan.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&formasi_jabatan)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       formasi_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateFormasiJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var formasi_jabatan model.FormasiJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&formasi_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(formasi_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	formasi_jabatan.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.FormasiJabatan{}).Where("id = ?", formasi_jabatan.Id).Updates(&formasi_jabatan)

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
		"data":       formasi_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteFormasiJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var formasi_jabatan model.DeleteFormasiJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&formasi_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(formasi_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.FormasiJabatan{}).Where("id = ?", formasi_jabatan.Id).Delete(&formasi_jabatan)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       formasi_jabatan,
		"statucCode": http.StatusOK,
	})
}
