package app

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	model "newSimpegAPI/model"
	"strings"
)

func FindPegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var searchString model.SearchPegawai
	var pegawai []model.Pegawai

	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		/*return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"message":    err.Error(),
			"statucCode": http.StatusNotImplemented,
		})*/
		searchString.Kstatus = []int{1, 2}
	}

	var findJabatan []string
	findJabatan = append(findJabatan, "left join master_riwayat_jabatan on master_pegawai.nip = master_riwayat_jabatan.nip and master_riwayat_jabatan.akhir = 1")

	result := db.Model(&model.Pegawai{})

	//IdOpd
	if searchString.IdOpd != "" {
		findJabatan = append(findJabatan, fmt.Sprintf("and master_riwayat_jabatan.id_opd = '%s'", searchString.IdOpd))
	}

	//id_eselon
	if searchString.Keselon != 0 {
		findJabatan = append(findJabatan, fmt.Sprintf("and master_riwayat_jabatan.keselon = '%v'", searchString.Keselon))
	}

	//jnsjab
	if len(searchString.Jnsjab) > 0 {
		jnsjab := strings.Trim(strings.Replace(fmt.Sprint(searchString.Jnsjab), " ", ",", -1), "[]")
		findJabatan = append(findJabatan, fmt.Sprintf("and master_riwayat_jabatan.jnsjab IN (%v)", jnsjab))

	}

	appendJab := strings.Join(findJabatan, " ")

	result = result.Joins(appendJab).Preload("JabatanAkhir", "akhir = ? ", 1)

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_pegawai.nama LIKE ?", strings.Join(str, ""))
	}

	//kstatus
	if len(searchString.Kstatus) > 0 {
		result = result.Where("master_pegawai.kstatus IN (?)", searchString.Kstatus)
	}

	//kjpeg
	if len(searchString.Kjpeg) > 0 {
		result = result.Where("master_pegawai.kjpeg IN (?)", searchString.Kjpeg)
	}

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_pegawai.nip = ?", searchString.Nip)
	}

	//no_ref_bkn
	if searchString.NoRefBkn != "" {
		result = result.Where("master_pegawai.no_ref_bkn = ?", searchString.NoRefBkn)
	}

	result = result.Find(&pegawai)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pegawai,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreatePegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var pegawai model.Pegawai
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pegawai)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	pegawai.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&pegawai)

	//fmt.Println(pegawai)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pegawai,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdatePegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var pegawai model.Pegawai
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pegawai)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	pegawai.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Pegawai{}).Where("nip = ?", pegawai.Nip).Updates(&pegawai)

	//fmt.Println(pegawai)

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
		"data":       pegawai,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeletePegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var pegawai model.DeletePegawai
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pegawai)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Pegawai{}).Where("nip = ?", pegawai.Nip).Delete(&pegawai)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pegawai,
		"statucCode": http.StatusOK,
	})
}
