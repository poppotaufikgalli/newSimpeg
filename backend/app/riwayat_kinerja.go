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
	"bytes"
)

func GetRiwayatKinerja(searchString model.SearchRiwayatKinerja) (riwayat_kinerja []model.RiwayatKinerja, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.RiwayatKinerja{})

	//ID
	if searchString.ID > 0 {
		result = result.Where("master_riwayat_kinerja.id = ?", searchString.ID)
	}

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_riwayat_kinerja.nip = ?", searchString.Nip)
	}

	//thnilai
	if searchString.Thnilai > 0 {
		result = result.Where("master_riwayat_kinerja.thnilai = ?", searchString.Thnilai)
	}

	//jenis
	if searchString.Jns != "" {
		result = result.Where("master_riwayat_kinerja.jns = ?", searchString.Jns)
	}

	//tmulai
	if searchString.Tmulai.IsZero() != true {
		result = result.Where("master_riwayat_kinerja.tmulai = ?", searchString.Tmulai)
	}

	//tselesai
	if searchString.Tselesai.IsZero() != true {
		result = result.Where("master_riwayat_kinerja.tselesai = ?", searchString.Tselesai)
	}

	result = result.Debug().Order("thnilai desc").Find(&riwayat_kinerja)

	return

}

func FindRiwayatKinerja(c echo.Context) error {
	var searchString model.SearchRiwayatKinerja
	json.NewDecoder(c.Request().Body).Decode(&searchString)

	riwayat_kinerja, result := GetRiwayatKinerja(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_kinerja,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetRiwayatKinerjaByNipIdThn(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_kinerja model.RiwayatKinerja
	nip := c.Param("nip")
	thnilai := c.Param("thnilai")
	id := c.Param("id")

	result := db.Model(&model.RiwayatKinerja{}).Where("nip = ? and thnilai = ? and id = ?", nip, thnilai, id).Scan(&riwayat_kinerja)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_kinerja,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func CreateRiwayatKinerja(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_kinerja model.RiwayatKinerja
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_kinerja)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_kinerja); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_kinerja.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&riwayat_kinerja)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_kinerja,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateRiwayatKinerja(c echo.Context) error {
	db, _ := model.CreateCon()

	//var riwayat_kinerja model.RiwayatKinerja
	validate := validator.New()

	type ToUpdateData struct {
		Ref  model.DeleteRiwayatKinerja `json:"refdata"`
		Data model.RiwayatKinerja       `json:"data"`
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
	//Tmulai := (reqData.Ref.Tmulai).Format("2006-01-02")
	//Tselesai := (reqData.Ref.Tselesai).Format("2006-01-02")
	result := db.Model(&model.RiwayatKinerja{}).Where("nip = ? and thnilai = ? and id = ?", reqData.Ref.Nip, reqData.Ref.Thnilai, reqData.Ref.ID).Updates(&reqData.Data)

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

func DeleteRiwayatKinerja(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_kinerja model.DeleteRiwayatKinerja
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_kinerja)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_kinerja); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	if c.Get("gid").(int) == 1 && riwayat_kinerja.IdSync != "" {
		status_bkn, err := sendDeleteRiwayatKinerjaBKN(riwayat_kinerja, c.Get("nip").(string))
		fmt.Println(status_bkn)

		if err != nil {
			return c.JSON(http.StatusNotImplemented, map[string]interface{}{
				"title":   "Gagal Menghapus Data Kinerja BKN",
				"status":  http.StatusNotImplemented,
				"message": err.Error(),
				"color":   "red",
			})
		}

		if status_bkn["success"] == false {
			return c.JSON(http.StatusNotImplemented, map[string]interface{}{
				"title":   "Gagal Menghapus Data Kinerja BKN",
				"status":  http.StatusNotImplemented,
				"message": status_bkn["message"],
				"color":   "red",
			})
		}
	}

	result := db.Model(&model.RiwayatKinerja{}).Where("nip = ? and thnilai = ? and id = ?", riwayat_kinerja.Nip, riwayat_kinerja.Thnilai, riwayat_kinerja.ID).Delete(&riwayat_kinerja)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_kinerja,
		"statucCode": http.StatusOK,
	})
}

func UpdateRiwayatKinerjaBkn(c echo.Context) error {
	var riwayat_kinerja model.RiwayatKinerja
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&riwayat_kinerja)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(riwayat_kinerja); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	if c.Get("gid").(int) == 1 {
		status_bkn, err := sendRiwayatKinerjaBKN(riwayat_kinerja, c.Get("nip").(string))
		fmt.Println(status_bkn)

		if err != nil {
			return c.JSON(http.StatusNotImplemented, map[string]interface{}{
				"statucCode": http.StatusNotImplemented,
				"errors":     err.Error(),
			})
		}
		if status_bkn["success"] != true {
			return c.JSON(http.StatusNotImplemented, map[string]string{
				"title":   "Update BKN Gagal",
				"message": fmt.Sprintf("%v-%s", status_bkn["code"], status_bkn["message"]),
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_kinerja,
		"statucCode": http.StatusOK,
	})
}

func sendRiwayatKinerjaBKN(riwayat_kinerja model.RiwayatKinerja, CreatedBy string) (status_bkn map[string]interface{}, err error) {
	db, _ := model.CreateCon()

	var pegawai model.Pegawai
	db.Select("no_ref_bkn").First(&pegawai, "nip = ?", riwayat_kinerja.Nip)

	kinerja := map[string]interface{}{
		"hasilKinerjaNilai":   riwayat_kinerja.HasilKinerja,
		"kuadranKinerjaNilai": riwayat_kinerja.KuadranKinerja,
		"penilaiGolongan":     fmt.Sprint(*riwayat_kinerja.PejKgolru),
		"penilaiJabatan":      checkPointer(riwayat_kinerja.PejNjab),
		"penilaiNama":         checkPointer(riwayat_kinerja.PejNama),
		"penilaiNipNrp":       checkPointer(riwayat_kinerja.PejNikNip),
		"penilaiUnorNama":     checkPointer(riwayat_kinerja.PejNunker),
		"perilakuKerjaNilai":  riwayat_kinerja.PerilakuKerja,
		"pnsDinilaiOrang":     *pegawai.NoRefBkn,
		"statusPenilai":       checkPointer(riwayat_kinerja.PejPns),
		"tahun":               riwayat_kinerja.Thnilai,
	}

	if checkPointer(riwayat_kinerja.IdSync) != "" {
		kinerja["id"] = checkPointer(riwayat_kinerja.IdSync)
	}

	//start update
	/*var singkronisasi model.Singkronisasi
	host := "bkn"
	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)*/

	urlBkn := "/kinerjaperiodik/save"
	if *riwayat_kinerja.Jns == "T" {
		urlBkn = "/skp22/save"
	} else {
		//0-7 errors occurred:
		//* PnsDinilaiId is required
		kinerja["PnsDinilaiId"] = *pegawai.NoRefBkn
		//* KoefisienId is required
		kinerja["KoefisienId"] = *riwayat_kinerja.KoefisienId
		//* BulanMulaiPenilaian is required
		mTmulai := (*riwayat_kinerja.Tmulai).Month()
		kinerja["BulanMulaiPenilaian"] = fmt.Sprint(int(mTmulai))
		//* BulanSelesaiPenilaian is required
		mTselesai := (*riwayat_kinerja.Tselesai).Month()
		kinerja["BulanSelesaiPenilaian"] = fmt.Sprint(int(mTselesai))
		//* TahunMulaiPenilaian is required
		kinerja["TahunMulaiPenilaian"] = (*riwayat_kinerja.Tmulai).Year()
		//* TahunSelesaiPenilaian is required
		kinerja["TahunSelesaiPenilaian"] = (*riwayat_kinerja.Tselesai).Year()
		//* PeriodikId is required
		kinerja["PeriodikId"] = fmt.Sprint(*riwayat_kinerja.PeriodikId)
	}

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0" + urlBkn)

	out, err := json.Marshal(kinerja)

	if err != nil {
		return
	}

	fmt.Println(string(out))

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewReader(out))

	if err != nil {
		return
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	//d := fmt.Sprintf("bearer %v", *singkronisasi.TokenApimanager)
	//e := fmt.Sprintf("Bearer %v", *singkronisasi.TokenSso)
	a := getTokenApiManager(CreatedBy)
	b := getTokenSSO(CreatedBy)
	d := fmt.Sprintf("bearer %v", a.AccessToken)
	e := fmt.Sprintf("Bearer %v", b.AccessToken)

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

	if status_bkn["success"] == true {
		status_bkn["code"] = 1
	}

	riwayat_update_bkn := &model.RiwayatUpdateBKN{
		Deskripsi: fmt.Sprintf("Riwayat Kinerja, NIP %s - Jenis %s Tahun. %v", riwayat_kinerja.Nip, *riwayat_kinerja.Jns, riwayat_kinerja.Thnilai),
		CreatedBy: CreatedBy,
		Code:      fmt.Sprint(status_bkn["code"]),
		Message:   fmt.Sprint(status_bkn["message"]),
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create

	if status_bkn["success"] == true {
		//db.Model(&model.RiwayatJabatan{}).Where().Update("idSync", respData.MapData.RwJabatanId)
		mapData := status_bkn["mapData"]
		//mapmapData := mapData.(map[string]string)

		//fmt.Println(mapData)
		db.Model(&model.RiwayatKinerja{}).Debug().Where("nip = ? and id = ?", riwayat_kinerja.Nip, riwayat_kinerja.ID).Update("idSync", mapData)
	}

	return
}

func sendDeleteRiwayatKinerjaBKN(riwayat_kinerja model.DeleteRiwayatKinerja, CreatedBy string) (status_bkn map[string]interface{}, err error) {
	db, _ := model.CreateCon()

	/*var singkronisasi model.Singkronisasi
	host := "bkn"
	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)*/

	urlBkn := "/kinerjaperiodik/delete"
	/*if riwayat_kinerja.Jns == "T" {
		urlBkn = "/skp22/delete"
	}*/

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/%s/%s", urlBkn, riwayat_kinerja.IdSync)

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	a := getTokenApiManager(CreatedBy)
	b := getTokenSSO(CreatedBy)
	d := fmt.Sprintf("bearer %v", a.AccessToken)
	e := fmt.Sprintf("Bearer %v", b.AccessToken)

	req.Header.Add("Auth", d)
	req.Header.Add("Authorization", e)

	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&status_bkn)

	if status_bkn["success"] == true {
		status_bkn["code"] = 1
	}

	riwayat_update_bkn := &model.RiwayatUpdateBKN{
		Deskripsi: fmt.Sprintf("Delete Riwayat Kinerja NIP %s - Jenis %s Tahun. %v", riwayat_kinerja.Nip, riwayat_kinerja.Jns, riwayat_kinerja.Thnilai),
		CreatedBy: CreatedBy,
		Code:      fmt.Sprint(status_bkn["code"]),
		Message:   fmt.Sprint(status_bkn["message"]),
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create*/

	return
}
