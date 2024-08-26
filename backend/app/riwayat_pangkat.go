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

func GetRiwayatPangkat(searchString model.SearchRiwayatPangkat) (riwayat_pangkat []model.RiwayatPangkat, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatPangkat{})

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_pangkat.nip = ?", searchString.Nip)
	}

	//tmtpang
	if searchString.Tmtpang.IsZero() != true {
		result = result.Where("master_riwayat_pangkat.tmtpang = ?", searchString.Tmtpang)
	}

	//Kgolru
	if len(searchString.Kgolru) > 0 {
		result = result.Where("master_riwayat_pangkat.kgolru IN (?)", searchString.Kgolru)
	}

	//Nskpangkat
	if searchString.Nskpang != "" {
		nskpang := strings.TrimSpace(searchString.Nskpang)
		str := []string{"%", nskpang, "%"}
		result = result.Where("master_riwayat_pangkat.nskpang LIKE ?", strings.Join(str, ""))
	}

	//Akhir
	if len(searchString.Akhir) > 0 {
		result = result.Where("master_riwayat_pangkat.akhir IN (?)", searchString.Akhir)
	}

	result = result.Preload("Pangkat").Preload("JenisKp").Order("kgolru desc").Find(&riwayat_pangkat)

	return

}

func FindRiwayatPangkat(c echo.Context) error {
	var searchString model.SearchRiwayatPangkat
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_pangkat, result := GetRiwayatPangkat(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetRiwayatPangkatByNip(c echo.Context) error {
	db, _ := model.CreateCon()
	var pangkat []model.RiwayatPangkat
	nip := c.Param("nip")

	result := db.Model(&model.RiwayatPangkat{}).Preload("Pangkat").Preload("JenisKp").Order("kgolru desc").Where("nip = ?", nip).Scan(&pangkat)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatPangkatByNipByKgolru(c echo.Context) error {
	db, _ := model.CreateCon()
	var pangkat model.RiwayatPangkat
	nip := c.Param("nip")
	kgolru := c.Param("kgolru")

	result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and kgolru = ?", nip, kgolru).Scan(&pangkat)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatPangkatByNipByKgolruTmtpangKnpang(c echo.Context) error {
	db, _ := model.CreateCon()
	var pangkat model.RiwayatPangkat
	nip := c.Param("nip")
	kgolru := c.Param("kgolru")
	tmtpang := c.Param("tmtpang")
	knpang := c.Param("knpang")

	result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and kgolru = ? and tmtpang = ? and knpang = ?", nip, kgolru, tmtpang, knpang).Scan(&pangkat)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatPangkatByNipCPNS(c echo.Context) error {
	db, _ := model.CreateCon()
	var pangkat model.RiwayatPangkat
	nip := c.Param("nip")

	result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and knpang = 211", nip).Scan(&pangkat)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatPangkatByNipPNS(c echo.Context) error {
	db, _ := model.CreateCon()
	var pangkat model.RiwayatPangkat
	nip := c.Param("nip")

	result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and knpang = 212", nip).Scan(&pangkat)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func GetRiwayatPangkatByNipAkhir(c echo.Context) error {
	db, _ := model.CreateCon()
	var pangkat model.RiwayatPangkat
	nip := c.Param("nip")

	result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and akhir = 1", nip).Scan(&pangkat)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateRiwayatPangkat(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_pangkat model.RiwayatPangkat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_pangkat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_pangkat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_pangkat.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_pangkat)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pangkat,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatPangkat(c echo.Context) error {
	db, _ := model.CreateCon()

	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatPangkat `json:"refdata"`
		Data model.RiwayatPangkat       `json:"data"`
	}

	var reqData ToUpdateData

	//var riwayat_pangkat model.RiwayatPangkat

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
	Tmtpang := (reqData.Ref.Tmtpang).Format("2006-01-02")
	result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and tmtpang = ? and kgolru = ? and knpang = ?", reqData.Ref.Nip, Tmtpang, reqData.Ref.Kgolru, reqData.Ref.Knpang).Updates(&reqData.Data)

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

	if c.Get("gid").(int) == 1 {
		status_bkn, err := sendRiwayatPangkatBKN(reqData.Data, c.Get("nip").(string))
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
				"description": fmt.Sprintf("%s %v %s", status_bkn["message"], status_bkn["data"], status_bkn["description"]),
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       reqData.Data,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteRiwayatPangkat(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_pangkat model.DeleteRiwayatPangkat
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_pangkat)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_pangkat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and tmtpang = ? and kgolru = ? and knpang = ?", riwayat_pangkat.Nip, riwayat_pangkat.Tmtpang, riwayat_pangkat.Kgolru, riwayat_pangkat.Knpang).Delete(&riwayat_pangkat)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_pangkat,
		"statucCode": http.StatusOK,
	})
}

func sendRiwayatPangkatBKN(riwayat_pangkat model.RiwayatPangkat, CreatedBy string) (status_bkn map[string]interface{}, err error) {
	db, _ := model.CreateCon()

	var pegawai model.Pegawai
	pangkat := make(map[string]string)
	db.Select("no_ref_bkn, kstatus, nkarpeg").First(&pegawai, "nip = ?", riwayat_pangkat.Nip)

	var urlBkn string
	var Kstatus string //Status Pegawai 1 : cpns 2: pns

	if riwayat_pangkat.Knpang == 211 {
		//cpns

		var dtPns model.RiwayatPangkat
		db.Select("nskpang, tskpang, tmtpang").First(&dtPns, "nip = ? and knpang = 212", riwayat_pangkat.Nip)

		if pegawai.Kstatus == 1 {
			Kstatus = "cpns"
		} else {
			Kstatus = "pns"
		}

		cpns := map[string]string{
			"kartu_pegawai":                *pegawai.Nkarpeg,
			"nomor_dokter_pns":             *riwayat_pangkat.Srtsehatno,
			"nomor_sk_cpns":                *riwayat_pangkat.Nskpang,
			"nomor_sk_pns":                 *dtPns.Nskpang,
			"nomor_spmt":                   *riwayat_pangkat.Nlgas,
			"nomor_sttpl":                  *riwayat_pangkat.Nsttpp,
			"pertek_cpns_pns_l2th_nomor":   riwayat_pangkat.Nntbakn,
			"pertek_cpns_pns_l2th_tanggal": "",
			"pns_orang_id":                 *pegawai.NoRefBkn,
			"status_cpns_pns":              Kstatus,
			"tanggal_dokter_pns":           "",
			"tgl_sk_cpns":                  riwayat_pangkat.Tskpang.Format("02-01-2006"),
			"tgl_sk_pns":                   dtPns.Tskpang.Format("02-01-2006"),
			"tgl_spmt":                     "",
			"tgl_sttpl":                    "",
			"tmt_pns":                      dtPns.Tmtpang.Format("02-01-2006"),
		}

		if riwayat_pangkat.Tntbakn != nil {
			cpns["pertek_cpns_pns_l2th_tanggal"] = (*riwayat_pangkat.Tntbakn).Format("02-01-2006")
		}

		if riwayat_pangkat.Tmtlgas != nil {
			cpns["tgl_spmt"] = (*riwayat_pangkat.Tmtlgas).Format("02-01-2006")
		}

		if riwayat_pangkat.Srtsehattgl != nil {
			cpns["tanggal_dokter_pns"] = (*riwayat_pangkat.Srtsehattgl).Format("02-01-2006")
		}

		if riwayat_pangkat.Tsttpp != nil {
			cpns["tgl_sttpl"] = (*riwayat_pangkat.Tsttpp).Format("02-01-2006")
		}

		pangkat = cpns
		urlBkn = "cpns"
	} else if riwayat_pangkat.Knpang == 212 {
		//pns
		var dtCpns model.RiwayatPangkat
		db.Select("nskpang, tskpang, nlgas, tmtlgas, nsttpp, tsttpp, srtsehatno, srtsehattgl").First(&dtCpns, "nip = ? and knpang = 211", riwayat_pangkat.Nip)

		//Status Pegawai 1 : cpns 2: pns
		var Kstatus string
		if pegawai.Kstatus == 1 {
			Kstatus = "cpns"
		} else {
			Kstatus = "pns"
		}

		pns := map[string]string{
			"kartu_pegawai":                *pegawai.Nkarpeg,
			"nomor_dokter_pns":             *dtCpns.Srtsehatno,
			"nomor_sk_cpns":                *dtCpns.Nskpang,
			"nomor_sk_pns":                 *riwayat_pangkat.Nskpang,
			"nomor_spmt":                   *dtCpns.Nlgas,
			"nomor_sttpl":                  *dtCpns.Nsttpp,
			"pertek_cpns_pns_l2th_nomor":   dtCpns.Nntbakn,
			"pertek_cpns_pns_l2th_tanggal": "",
			"pns_orang_id":                 *pegawai.NoRefBkn,
			"status_cpns_pns":              Kstatus,
			"tanggal_dokter_pns":           "",
			"tgl_sk_cpns":                  dtCpns.Tskpang.Format("02-01-2006"),
			"tgl_sk_pns":                   riwayat_pangkat.Tskpang.Format("02-01-2006"),
			"tgl_spmt":                     "",
			"tgl_sttpl":                    "",
			"tmt_pns":                      riwayat_pangkat.Tmtpang.Format("02-01-2006"),
		}

		if dtCpns.Tntbakn != nil {
			pns["pertek_cpns_pns_l2th_tanggal"] = (*dtCpns.Tntbakn).Format("02-01-2006")
		}

		if dtCpns.Tmtlgas != nil {
			pns["tgl_spmt"] = (*dtCpns.Tmtlgas).Format("02-01-2006")
		}

		if dtCpns.Srtsehattgl != nil {
			pns["tanggal_dokter_pns"] = (*dtCpns.Srtsehattgl).Format("02-01-2006")
		}

		if dtCpns.Tsttpp != nil {
			pns["tgl_sttpl"] = (*dtCpns.Tsttpp).Format("02-01-2006")
		}

		pangkat = pns
		urlBkn = "cpns"
	} else {
		//pangkat

	}

	fmt.Println(pangkat)

	//start update
	var singkronisasi model.Singkronisasi
	host := "bkn"
	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)

	url := fmt.Sprintf(`https://apimws.bkn.go.id:8243/apisiasn/1.0/%s/save`, urlBkn)

	out, err := json.Marshal(pangkat)

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
		Deskripsi: fmt.Sprintf("Riwayat %s %s", urlBkn, riwayat_pangkat.Nskpang),
		CreatedBy: CreatedBy,
		Code:      fmt.Sprint(status_bkn["code"]),
		Message:   fmt.Sprint(status_bkn["message"]),
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create*/

	/*if status_bkn["success"] == true {
		//db.Model(&model.RiwayatJabatan{}).Where().Update("idSync", respData.MapData.RwJabatanId)
		mapData := status_bkn["mapData"]
		mapmapData := mapData.(map[string]interface{})

		//fmt.Println()
		db.Model(&model.RiwayatJabatan{}).Debug().Where("nip = ? and id = ?", riwayat_jabatan.Nip, riwayat_jabatan.ID).Update("idSync", mapmapData["rwJabatanId"])
	}*/

	return
}

/*func sendDeleteRiwayatPangkatBKN(riwayat_jabatan model.DeleteRiwayatPangkat, CreatedBy string) (status_bkn map[string]interface{}, err error) {
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

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create

	return
}*/
