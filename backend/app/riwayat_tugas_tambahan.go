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

func GetRiwayatTugasTambahan(searchString model.SearchRiwayatTugasTambahan) (riwayat_tugas_tambahan []model.RiwayatTugasTambahan, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatTugasTambahan{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_tugas_tambahan.nip = ?", searchString.Nip)
	}

	//tmtjab
	if searchString.Tmtjab.IsZero() != true {
		result = result.Where("master_riwayat_tugas_tambahan.tmtjab = ?", searchString.Tmtjab)
	}

	//Njab
	if searchString.Njab != "" {
		nama := strings.TrimSpace(searchString.Njab)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_tugas_tambahan.njab LIKE ?", strings.Join(str, ""))
	}

	//Nunker
	if searchString.Nunker != "" {
		nama := strings.TrimSpace(searchString.Nunker)
		str := []string{"%", nama, "%"}
		result = result.Where("master_riwayat_tugas_tambahan.nunker LIKE ?", strings.Join(str, ""))
	}

	//id OPD
	if searchString.IdOpd != "" {
		result = result.Where("master_riwayat_tugas_tambahan.id_opd = ?", searchString.IdOpd)
	}

	//status
	if len(searchString.Status) > 0 {
		result = result.Where("master_riwayat_tugas_tambahan.status IN (?)", searchString.Status)
	}

	result = result.Order("tmtjab desc").Find(&riwayat_tugas_tambahan)

	return

}

func FindRiwayatTugasTambahan(c echo.Context) error {
	var searchString model.SearchRiwayatTugasTambahan
	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		searchString.Status = []int{1}
	}

	riwayat_tugas_tambahan, result := GetRiwayatTugasTambahan(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_tugas_tambahan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetRiwayatTugasTambahanByNip(c echo.Context) error {
	db, _ := model.CreateCon()
	var riwayat_tugas_tambahan model.RiwayatTugasTambahan
	nip := c.Param("nip")

	result := db.Model(&model.RiwayatTugasTambahan{}).Where("nip = ?", nip).Scan(&riwayat_tugas_tambahan)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_tugas_tambahan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatTugasTambahanByNipAkhir(c echo.Context) error {
	db, _ := model.CreateCon()
	var riwayat_tugas_tambahan model.RiwayatTugasTambahan
	nip := c.Param("nip")

	result := db.Model(&model.RiwayatTugasTambahan{}).Where("nip = ? and status = 1", nip).Scan(&riwayat_tugas_tambahan)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_tugas_tambahan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatTugasTambahanByNipTmtjab(c echo.Context) error {
	db, _ := model.CreateCon()
	var riwayat_tugas_tambahan model.RiwayatTugasTambahan
	nip := c.Param("nip")
	tmtjab := c.Param("tmtjab")

	result := db.Model(&model.RiwayatTugasTambahan{}).Where("nip = ? and tmtjab = ?", nip, tmtjab).Scan(&riwayat_tugas_tambahan)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_tugas_tambahan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateRiwayatTugasTambahan(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_tugas_tambahan model.RiwayatTugasTambahan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_tugas_tambahan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_tugas_tambahan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_tugas_tambahan.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_tugas_tambahan)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_tugas_tambahan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatTugasTambahan(c echo.Context) error {
	db, _ := model.CreateCon()

	//var riwayat_tugas_tambahan model.RiwayatTugasTambahan
	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatTugasTambahan `json:"refdata"`
		Data model.RiwayatTugasTambahan       `json:"data"`
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
	Tmtjab := (reqData.Ref.Tmtjab).Format("2006-01-02")
	result := db.Model(&model.RiwayatTugasTambahan{}).Where("nip = ? and tmtjab = ?", reqData.Ref.Nip, Tmtjab).Updates(&reqData.Data)

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

func DeleteRiwayatTugasTambahan(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_tugas_tambahan model.DeleteRiwayatTugasTambahan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_tugas_tambahan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_tugas_tambahan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatTugasTambahan{}).Where("nip = ? and tmtjab = ?", riwayat_tugas_tambahan.Nip, riwayat_tugas_tambahan.Tmtjab).Delete(&riwayat_tugas_tambahan)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_tugas_tambahan,
		"statucCode": http.StatusOK,
	})
}
