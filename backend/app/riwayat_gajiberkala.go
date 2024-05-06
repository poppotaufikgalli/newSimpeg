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
	//"strings"
)

func GetRiwayatGajiberkala(searchString model.SearchRiwayatGajiberkala) (riwayat_gajiberkala []model.RiwayatGajiberkala, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatGajiberkala{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_gajiberkala.nip = ?", searchString.Nip)
	}

	//Nstahu
	if searchString.Nstahu != "" {
		result = result.Where("master_riwayat_gajiberkala.nstahu = ?", searchString.Nstahu)
	}

	//Tmtngaj
	if searchString.Tmtngaj.IsZero() != true {
		result = result.Where("master_riwayat_gajiberkala.tmtngaj = ?", searchString.Tmtngaj)
	}

	//id opd
	if searchString.IdOpd != "" {
		result = result.Where("master_riwayat_gajiberkala.id_opd = ?", searchString.IdOpd)
	}

	result = result.Find(&riwayat_gajiberkala)

	return

}

func FindRiwayatGajiberkala(c echo.Context) error {
	var searchString model.SearchRiwayatGajiberkala
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_gajiberkala, result := GetRiwayatGajiberkala(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_gajiberkala,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetRiwayatGajiBerkalaByNip(c echo.Context) error {
	db, _ := model.CreateCon()
	var riwayat_gajiberkala model.RiwayatGajiberkala
	nip := c.Param("nip")

	result := db.Model(&model.RiwayatGajiberkala{}).Where("nip = ?", nip).Scan(&riwayat_gajiberkala)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_gajiberkala,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatGajiBerkalaByNipAkhir(c echo.Context) error {
	db, _ := model.CreateCon()
	var riwayat_gajiberkala model.RiwayatGajiberkala
	nip := c.Param("nip")

	result := db.Model(&model.RiwayatGajiberkala{}).Debug().Where("nip = ? and akhir = 1", nip).Scan(&riwayat_gajiberkala)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_gajiberkala,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatGajiBerkalaByNipTmtngaj(c echo.Context) error {
	db, _ := model.CreateCon()
	var riwayat_gajiberkala model.RiwayatGajiberkala
	nip := c.Param("nip")
	tmtngaj := c.Param("tmtngaj")

	result := db.Model(&model.RiwayatGajiberkala{}).Debug().Where("nip = ? and tmtngaj = ?", nip, tmtngaj).Scan(&riwayat_gajiberkala)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_gajiberkala,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateRiwayatGajiberkala(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_gajiberkala model.RiwayatGajiberkala
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_gajiberkala)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_gajiberkala); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_gajiberkala.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_gajiberkala)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_gajiberkala,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatGajiberkala(c echo.Context) error {
	db, _ := model.CreateCon()

	//var riwayat_gajiberkala model.RiwayatGajiberkala
	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatGajiberkala `json:"refdata"`
		Data model.RiwayatGajiberkala       `json:"data"`
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
	Tmtngaj := (reqData.Ref.Tmtngaj).Format("2006-01-02")
	result := db.Model(&model.RiwayatGajiberkala{}).Where("nip = ? and tmtngaj = ?", reqData.Ref.Nip, Tmtngaj).Updates(&reqData.Data)

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

func DeleteRiwayatGajiberkala(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_gajiberkala model.DeleteRiwayatGajiberkala
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_gajiberkala)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_gajiberkala); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatGajiberkala{}).Where("nip = ? and tmtngaj = ? and id_opd = ?", riwayat_gajiberkala.Nip, riwayat_gajiberkala.Tmtngaj, riwayat_gajiberkala.IdOpd).Delete(&riwayat_gajiberkala)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_gajiberkala,
		"statucCode": http.StatusOK,
	})
}
