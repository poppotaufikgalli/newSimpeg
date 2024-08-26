package app

import (
	"bytes"
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

func GetRiwayatJabatan(searchString model.SearchRiwayatJabatan) (riwayat_jabatan []model.RiwayatJabatan, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatJabatan{}).Preload("JenisTugasMutasi")

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_jabatan.nip = ?", searchString.Nip)
	}

	//tmtjab
	if searchString.Tmtjab.IsZero() != true {
		result = result.Where("master_riwayat_jabatan.tmtjab = ?", searchString.Tmtjab)
	}

	//kjab
	if searchString.Kjab != "" {
		result = result.Where("master_riwayat_jabatan.kjab = ?", searchString.Kjab)
	}

	//kjab bkn
	if searchString.KjabBkn != "" {
		result = result.Where("master_riwayat_jabatan.kjab_bkn = ?", searchString.KjabBkn)
	}

	//kjab bkn
	if searchString.Jnsjab != "" {
		result = result.Where("master_riwayat_jabatan.jnsjab = ?", searchString.Jnsjab)
	}

	//Keselon
	if len(searchString.Keselon) > 0 {
		result = result.Where("master_riwayat_jabatan.keselon IN (?)", searchString.Keselon)
	}

	//Nskjabat
	if searchString.Nskjabat != "" {
		nskjabat := strings.TrimSpace(searchString.Nskjabat)
		str := []string{"%", nskjabat, "%"}
		result = result.Where("master_riwayat_jabatan.nskjabat LIKE ?", strings.Join(str, ""))
	}

	//id Opd
	if searchString.IdOpd != "" {
		result = result.Where("master_riwayat_jabatan.id_opd = ?", searchString.IdOpd)
	}

	//Akhir
	if len(searchString.Akhir) > 0 {
		result = result.Where("master_riwayat_jabatan.akhir IN (?)", searchString.Akhir)
	}

	//Tugas Tambahan
	if searchString.TugasTambahan != "" {
		result = result.Where("master_riwayat_jabatan.tugas_tambahan = ?", searchString.TugasTambahan)
	}

	//Id Opd Tambahan
	if searchString.IdOpdTambahan != "" {
		result = result.Where("master_riwayat_jabatan.id_opd_tambahan = ?", searchString.IdOpdTambahan)
	}

	//idSync
	if searchString.IdSync != "" {
		result = result.Where("master_riwayat_jabatan.idSync = ?", searchString.IdSync)
	}

	result = result.Debug().Preload("Eselon").Order("tmtjab desc").Find(&riwayat_jabatan)

	return

}

func FindRiwayatJabatan(c echo.Context) error {
	var searchString model.SearchRiwayatJabatan
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_jabatan, result := GetRiwayatJabatan(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetRiwayatJabatanByNip(c echo.Context) error {
	db, _ := model.CreateCon()
	var riwayat_jabatan model.RiwayatJabatan
	nip := c.Param("nip")

	result := db.Model(&model.RiwayatJabatan{}).Where("nip = ?", nip).Scan(&riwayat_jabatan)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatJabatanByNipAkhir(c echo.Context) error {
	db, _ := model.CreateCon()
	var riwayat_jabatan model.RiwayatJabatan
	nip := c.Param("nip")

	result := db.Model(&model.RiwayatJabatan{}).Debug().Where("nip = ? and akhir = 1", nip).Scan(&riwayat_jabatan)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatJabatanByNipTmtjab(c echo.Context) error {
	db, _ := model.CreateCon()
	var riwayat_jabatan model.RiwayatJabatan
	nip := c.Param("nip")
	tmtjab := c.Param("tmtjab")

	result := db.Model(&model.RiwayatJabatan{}).Debug().Where("nip = ? and tmtjab = ?", nip, tmtjab).Scan(&riwayat_jabatan)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatJabatanByNipId(c echo.Context) error {
	db, _ := model.CreateCon()
	var riwayat_jabatan model.RiwayatJabatan
	nip := c.Param("nip")
	id := c.Param("id")

	result := db.Model(&model.RiwayatJabatan{}).Debug().Where("nip = ? and id = ?", nip, id).Scan(&riwayat_jabatan)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatJabatanByNipTmtjabTmtMutasi(c echo.Context) error {
	db, _ := model.CreateCon()
	var riwayat_jabatan model.RiwayatJabatan
	nip := c.Param("nip")
	tmtjab := c.Param("tmtjab")
	tmt_mutasi := c.Param("tmt_mutasi")

	result := db.Model(&model.RiwayatJabatan{}).Debug().Scopes(model.CheckTmtMutasi(tmt_mutasi)).Where("nip = ? and tmtjab = ?", nip, tmtjab).Scan(&riwayat_jabatan)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateRiwayatJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_jabatan model.RiwayatJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_jabatan.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_jabatan)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_jabatan,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	//var riwayat_jabatan model.RiwayatJabatan
	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatJabatan `json:"refdata"`
		Data model.RiwayatJabatan       `json:"data"`
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

	result := db.Model(&model.RiwayatJabatan{}).Where("nip= ? and id = ?", reqData.Ref.Nip, reqData.Ref.ID).Updates(&reqData.Data)

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

func DeleteRiwayatJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_jabatan model.DeleteRiwayatJabatan
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"title":      "Gagal Hapus Data",
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"title":      "Gagal Hapus Data",
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//result := db.Model(&model.RiwayatJabatan{}).Where("nip = ? and tmtjab = ? and tmt_mutasi = ?", riwayat_jabatan.Nip, riwayat_jabatan.Tmtjab, riwayat_jabatan.TmtMutasi).Delete(&riwayat_jabatan)

	/*if c.Get("gid").(int) == 1 && riwayat_jabatan.IdSync != "" {
		status_bkn, err := sendDeleteRiwayatJabatanBKN(riwayat_jabatan, c.Get("nip").(string))
		fmt.Println(status_bkn)

		if err != nil {
			return c.JSON(http.StatusNotImplemented, map[string]interface{}{
				"statucCode": http.StatusNotImplemented,
				"errors":     err.Error(),
			})
		}

		if status_bkn["success"] != true {
			return c.JSON(http.StatusNotImplemented, map[string]string{
				"title":       "Hapus Data Jabatan BKN Gagal",
				"description": fmt.Sprintf("%s-%s", status_bkn["code"], status_bkn["message"]),
			})
		}
	}*/

	result := db.Model(&model.RiwayatJabatan{}).Where("nip=? and id = ?", riwayat_jabatan.Nip, riwayat_jabatan.ID).Delete(&riwayat_jabatan)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"title":      "Gagal Hapus Data",
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_jabatan,
		"statucCode": http.StatusOK,
	})
}

func UpdateRiwayatJabatanBkn(c echo.Context) error {
	var riwayat_jabatan model.RiwayatJabatan

	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_jabatan)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_jabatan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	if c.Get("gid").(int) == 1 {
		status_bkn, err := sendRiwayatJabatanBKN(riwayat_jabatan, c.Get("nip").(string))
		fmt.Println(status_bkn)

		if err != nil {
			return c.JSON(http.StatusNotImplemented, map[string]interface{}{
				"statucCode": http.StatusNotImplemented,
				"errors":     err.Error(),
			})
		}
		if status_bkn["success"] != true {
			return c.JSON(http.StatusNotImplemented, map[string]string{
				"title":       "Update BKN Gagal",
				"description": fmt.Sprintf("%s-%s", status_bkn["code"], status_bkn["message"]),
			})
		}
	} else {
		return c.JSON(http.StatusNotImplemented, map[string]string{
			"title":       "Update BKN Gagal",
			"description": "Anda tidak memiliki hak akses",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_jabatan,
		"statucCode": http.StatusOK,
	})
}

func sendRiwayatJabatanBKN(riwayat_jabatan model.RiwayatJabatan, CreatedBy string) (status_bkn map[string]interface{}, err error) {
	db, _ := model.CreateCon()

	var pegawai model.Pegawai
	db.Select("no_ref_bkn").First(&pegawai, "nip = ?", riwayat_jabatan.Nip)

	Tskjabat := (riwayat_jabatan.Tskjabat).Format("02-01-2006")
	Tmtjab := (riwayat_jabatan.Tmtjab).Format("02-01-2006")
	Tlantik := (riwayat_jabatan.Tlantik).Format("02-01-2006")

	jabatan := map[string]string{
		"eselonId":                fmt.Sprint(riwayat_jabatan.Keselon),
		"instansiId":              *riwayat_jabatan.IdInstansi,
		"jabatanFungsionalId":     "",
		"jabatanFungsionalUmumId": "",
		"jenisJabatan":            fmt.Sprint(riwayat_jabatan.Jnsjab),
		"jenisMutasiId":           "",
		"jenisPenugasanId":        "",
		"nomorSk":                 riwayat_jabatan.Nskjabat,
		"pnsId":                   *pegawai.NoRefBkn,
		"subJabatanId":            fmt.Sprint(riwayat_jabatan.SubJab),
		"tanggalSk":               Tskjabat,
		"tmtJabatan":              Tmtjab,
		"tmtMutasi":               "",
		"tmtPelantikan":           Tlantik,
		"unorId":                  *riwayat_jabatan.IdUnorBkn,
		"satuanKerjaId":           "A5EB03E24326F6A0E040640A040252AD",
	}

	if riwayat_jabatan.Jnsjab == "1" {
		jabatan["jenisPenugasanId"] = fmt.Sprint(riwayat_jabatan.JnsTugasMutasi)
	} else {
		jabatan["jenisMutasiId"] = fmt.Sprint(riwayat_jabatan.JnsTugasMutasi)

		if riwayat_jabatan.Jnsjab == "2" {
			jabatan["jabatanFungsionalId"] = fmt.Sprint(riwayat_jabatan.Kjab)
		} else if riwayat_jabatan.Jnsjab == "4" {
			jabatan["jabatanFungsionalUmumId"] = fmt.Sprint(riwayat_jabatan.Kjab)
		}
	}

	if riwayat_jabatan.JnsTugasMutasi == "MU" {
		TmtMutasi := (riwayat_jabatan.TmtMutasi).Format("02-01-2006")
		jabatan["tmtMutasi"] = fmt.Sprint(TmtMutasi)
	}
	//fmt.Println(jabatan)

	//start update
	var singkronisasi model.Singkronisasi
	host := "bkn"
	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/jabatan/unorjabatan/save")

	out, err := json.Marshal(jabatan)

	if err != nil {
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewReader(out))

	if err != nil {
		return
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	d := fmt.Sprintf("bearer %v", *singkronisasi.TokenApimanager)
	e := fmt.Sprintf("Bearer %v", *singkronisasi.TokenSso)

	req.Header.Add("Auth", d)
	req.Header.Add("Authorization", e)

	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&status_bkn)

	if err != nil {
		return
	}

	riwayat_update_bkn := &model.RiwayatUpdateBKN{
		Deskripsi: fmt.Sprintf("Riwayat Jabatan, NIP %s - SK NO. %s", riwayat_jabatan.Nip, riwayat_jabatan.Nskjabat),
		CreatedBy: CreatedBy,
		Code:      fmt.Sprint(status_bkn["code"]),
		Message:   fmt.Sprint(status_bkn["message"]),
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create*/

	if status_bkn["success"] == true {
		//db.Model(&model.RiwayatJabatan{}).Where().Update("idSync", respData.MapData.RwJabatanId)
		mapData := status_bkn["mapData"]
		mapmapData := mapData.(map[string]interface{})

		//fmt.Println()
		db.Model(&model.RiwayatJabatan{}).Debug().Where("nip = ? and id = ?", riwayat_jabatan.Nip, riwayat_jabatan.ID).Update("idSync", mapmapData["rwJabatanId"])
	}

	return
}

func sendDeleteRiwayatJabatanBKN(riwayat_jabatan model.DeleteRiwayatJabatan, CreatedBy string) (status_bkn map[string]interface{}, err error) {
	db, _ := model.CreateCon()

	fmt.Println(riwayat_jabatan)

	var singkronisasi model.Singkronisasi
	host := "bkn"
	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/jabatan/delete/%s", riwayat_jabatan.IdSync)

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	d := fmt.Sprintf("bearer %v", *singkronisasi.TokenApimanager)
	e := fmt.Sprintf("Bearer %v", *singkronisasi.TokenSso)

	req.Header.Add("Auth", d)
	req.Header.Add("Authorization", e)

	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&status_bkn)

	riwayat_update_bkn := &model.RiwayatUpdateBKN{
		Deskripsi: fmt.Sprintf("Delete Riwayat Jabatan %s", riwayat_jabatan.Nskjabat),
		CreatedBy: CreatedBy,
		Code:      fmt.Sprint(status_bkn["code"]),
		Message:   fmt.Sprint(status_bkn["message"]),
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create*/

	return
}
