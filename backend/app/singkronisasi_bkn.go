package app

import (
	"bytes"
	_ "database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"net/url"
	model "newSimpegAPI/model"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func UpdateTokenSingkronisasi(c echo.Context) error {
	//db, _ := model.CreateCon()

	var updateTokens model.UpdateTokenSingkronisasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&updateTokens)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, err.Error())
	}

	if err = validate.Struct(updateTokens); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if updateTokens.Type == "token_sso" {
		//retVal := getTokenSSO(updateTokens.Ckey, updateTokens.Csecret)
		retVal := getTokenSSO(fmt.Sprint(c.Get("nip")))
		fmt.Println("A")

		/*now := time.Now()
		ExpiresIn := now.Add(time.Duration(retVal.ExpiresIn) * time.Second)

		singkronisasi := model.Singkronisasi{
			TokenSso:        &retVal.AccessToken,
			TokenSsoExpired: &ExpiresIn,
			UpdatedBy:       fmt.Sprint(c.Get("nip")),
		}

		db.Model(&model.Singkronisasi{}).Where("id = '1'").Updates(&singkronisasi)*/

		//return c.JSON(http.StatusOK, singkronisasi)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":       retVal,
			"statusCode": http.StatusOK,
		})

	} else if updateTokens.Type == "token_apimanager" {
		//retVal := getTokenApiManager(updateTokens.ClientId, updateTokens.Username, updateTokens.Password)
		retVal := getTokenApiManager(fmt.Sprint(c.Get("nip")))

		fmt.Println("B")

		/*now := time.Now()
		ExpiresIn := now.Add(time.Duration(retVal.ExpiresIn) * time.Second)

		singkronisasi := model.Singkronisasi{
			TokenApimanager:        &retVal.AccessToken,
			TokenApimanagerExpired: &ExpiresIn,
			UpdatedBy:              fmt.Sprint(c.Get("nip")),
		}

		db.Model(&model.Singkronisasi{}).Where("id = '1'").Updates(&singkronisasi)*/

		//return c.JSON(http.StatusOK, singkronisasi)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":       retVal,
			"statusCode": http.StatusOK,
		})
	}

	return c.JSON(http.StatusOK, nil)
}

func getTokenSSO(nip string) (res model.ResultTokenSso) {
	db, _ := model.CreateCon()
	payload := strings.NewReader("grant_type=client_credentials")

	var singkronisasi model.Singkronisasi

	db.Model(model.Singkronisasi{}).Find(&singkronisasi, 1)

	now := time.Now()

	if (*singkronisasi.TokenSsoExpired).Unix() < time.Now().Unix() {

		auth := fmt.Sprintf("%s:%s", *singkronisasi.Ckey, *singkronisasi.Csecret)
		credentials := base64.StdEncoding.EncodeToString([]byte(auth))

		url := "https://apimws.bkn.go.id/oauth2/token"

		client := &http.Client{}
		req, err := http.NewRequest("POST", url, payload)
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Authorization", "Basic "+credentials)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&res)

		ExpiresIn := now.Add(time.Duration(res.ExpiresIn) * time.Second)

		singkronisasi.TokenSso = &res.AccessToken
		singkronisasi.TokenSsoExpired = &ExpiresIn
		singkronisasi.UpdatedBy = nip

		db.Model(&model.Singkronisasi{}).Where("id = '1'").Updates(&singkronisasi)
	} else {
		res.AccessToken = *singkronisasi.TokenSso
	}

	return
}

func getTokenApiManager(nip string) (res model.ResultTokenSso) {
	db, _ := model.CreateCon()

	var singkronisasi model.Singkronisasi

	db.Model(model.Singkronisasi{}).Find(&singkronisasi, 1)

	now := time.Now()

	if (*singkronisasi.TokenApimanagerExpired).Unix() < now.Unix() {

		payload := strings.NewReader(fmt.Sprintf("client_id=%s&grant_type=password&username=%s&password=%s", *singkronisasi.ClientId, *singkronisasi.Username, *singkronisasi.Password))
		fmt.Println(payload)

		url := "https://sso-siasn.bkn.go.id/auth/realms/public-siasn/protocol/openid-connect/token"

		client := &http.Client{}
		req, err := http.NewRequest("POST", url, payload)
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&res)

		ExpiresIn := now.Add(time.Duration(res.ExpiresIn) * time.Second)

		singkronisasi.TokenApimanager = &res.AccessToken
		singkronisasi.TokenApimanagerExpired = &ExpiresIn
		singkronisasi.UpdatedBy = nip

		db.Model(&model.Singkronisasi{}).Where("id = '1'").Updates(&singkronisasi)
	} else {
		res.AccessToken = *singkronisasi.TokenApimanager
	}

	return
}

func GetSingkronisasiRiwayat(c echo.Context) error {
	db, _ := model.CreateCon()

	var riwayat_update_bkn []model.RiwayatUpdateBKN

	nip := fmt.Sprint(c.Get("nip"))

	db.Model(&model.RiwayatUpdateBKN{}).Where("created_by", nip).Scan(&riwayat_update_bkn)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       riwayat_update_bkn,
		"statusCode": http.StatusOK,
	})
}

func GetSingkronisasiGetDataBkn(c echo.Context) error {
	CreatedBy := c.Get("nip").(string)

	nip := c.Param("nip")
	page := c.Param("page")

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/pns/%s/%s", page, nip)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statusCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}
	req.Header.Add("accept", "application/json")
	a := getTokenApiManager(CreatedBy)
	b := getTokenSSO(CreatedBy)
	d := fmt.Sprintf("bearer %v", a.AccessToken)
	e := fmt.Sprintf("Bearer %v", b.AccessToken)

	req.Header.Add("Auth", d)
	req.Header.Add("Authorization", e)

	res, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statusCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}
	defer res.Body.Close()

	var retData map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&retData)
	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statusCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       retData,
		"statusCode": http.StatusOK,
	})
}

func GetSingkronisasiGetDataBknById(c echo.Context) error {
	CreatedBy := c.Get("nip").(string)

	page := c.Param("page")
	idSync := c.Param("idSync")

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/%s/id/%s", page, idSync)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statusCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}
	req.Header.Add("accept", "application/json")
	a := getTokenApiManager(CreatedBy)
	b := getTokenSSO(CreatedBy)
	d := fmt.Sprintf("bearer %v", a.AccessToken)
	e := fmt.Sprintf("Bearer %v", b.AccessToken)

	req.Header.Add("Auth", d)
	req.Header.Add("Authorization", e)

	res, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statusCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}
	defer res.Body.Close()

	var retData map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&retData)
	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statusCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       retData,
		"statusCode": http.StatusOK,
	})
}

func PostSinkronisasiGetDataBKN(updatas map[string]interface{}, page, CreatedBy string) (retData map[string]interface{}, err error) {
	url := fmt.Sprintf(`https://apimws.bkn.go.id:8243/apisiasn/1.0/%s/save`, page)

	out, err := json.Marshal(updatas)

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

	err = json.NewDecoder(res.Body).Decode(&retData)
	if err != nil {
		return
	}

	return
}

func DeleteSingkronisasiDataBknById(page, idSync, CreatedBy string) (retData map[string]interface{}, err error) {
	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/%s/delete/%s", page, idSync)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return
	}

	req.Header.Add("accept", "application/json")
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

	err = json.NewDecoder(res.Body).Decode(&retData)
	if err != nil {
		return
	}

	return
}

//========================================================================================================================================================//

func UpdateSingkronisasiPutDataUtamaBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	validate := validator.New()

	var updatas model.UpdateBKNDataUtama
	var singkronisasi model.Singkronisasi

	page := c.Param("page")
	host := "bkn"

	err := json.NewDecoder(c.Request().Body).Decode(&updatas)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statusCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(updatas); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//start doupdate
	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/pns/%s-update", page)

	out, err := json.Marshal(updatas)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "A",
			"errors":     err.Error(),
		})
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewReader(out))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "B",
			"errors":     err.Error(),
		})
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	d := fmt.Sprintf("bearer %v", *singkronisasi.TokenApimanager)
	e := fmt.Sprintf("Bearer %v", *singkronisasi.TokenSso)

	//fmt.Println(d)
	//fmt.Println(e)

	req.Header.Add("Auth", d)
	req.Header.Add("Authorization", e)

	res, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "C",
			"errors":     err.Error(),
		})
	}
	defer res.Body.Close()

	var retData map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&retData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "D",
			"errors":     err.Error(),
		})
	}
	fmt.Println(retData)
	//log data
	code := fmt.Sprintf("%v", retData["code"])
	riwayat_update_bkn := &model.RiwayatUpdateBKN{
		Deskripsi: "Update Data Utama/Identitas",
		CreatedBy: fmt.Sprint(c.Get("nip")),
		Code:      code,
		Content:   string(out),
		Message:   retData["message"].(string),
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create

	//end doupdate

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       retData,
		"statusCode": http.StatusOK,
	})
}

func PostSinkronisasiGetDocFromBKN(c echo.Context) error {
	db, _ := model.CreateCon()

	var docData model.DocKPBKN
	var singkronisasi model.Singkronisasi
	host := "bkn"

	err := json.NewDecoder(c.Request().Body).Decode(&docData)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statusCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	//start doupdate
	// Create the file
	ext := filepath.Ext(filepath.Base(docData.Path))

	filename := fmt.Sprintf("%s_pangkat%v_%v%s", docData.RiwayatPangkat.Nip, docData.RiwayatPangkat.Kgolru, docData.RiwayatPangkat.Knpang, ext)
	fmt.Println(filename)

	newpath := filepath.Join("assets", "dokumen", docData.RiwayatPangkat.Nip, filename)
	out, err := os.Create(newpath)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "A",
			"errors":     err.Error(),
		})
	}
	defer out.Close()

	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)
	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/download-dok?filePath=%s", url.QueryEscape(docData.Path))

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "B",
			"errors":     err.Error(),
		})
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	d := fmt.Sprintf("bearer %v", *singkronisasi.TokenApimanager)
	e := fmt.Sprintf("Bearer %v", *singkronisasi.TokenSso)

	//fmt.Println(d)
	//fmt.Println(e)

	req.Header.Add("Auth", d)
	req.Header.Add("Authorization", e)

	res, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "C",
			"errors":     err.Error(),
		})
	}
	defer res.Body.Close()

	// Writer the body to file
	retData, err := io.Copy(out, res.Body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "D",
			"errors":     err.Error(),
		})
	}

	//end doupdate

	Tmtpang := docData.RiwayatPangkat.Tmtpang.Format("2006-01-02")
	result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and tmtpang = ? and kgolru = ? and knpang = ?", docData.RiwayatPangkat.Nip, Tmtpang, docData.RiwayatPangkat.Kgolru, docData.RiwayatPangkat.Knpang).Update("filename", filename)
	fmt.Println(result.RowsAffected)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"size":     retData,
			"filename": filename,
			"path":     newpath,
		},
		"statusCode": http.StatusOK,
	})
}

func UpdateSingkronisasiPutDataBknAngkaKredit(c echo.Context) error {
	db, _ := model.CreateCon()

	nip := c.Param("nip")

	var retData model.RiwayatAngkaKredit
	//var toSendData model.UpdateAngkaKreditBKN
	var pegawai model.Pegawai
	var singkronisasi model.Singkronisasi
	host := "bkn"

	err := json.NewDecoder(c.Request().Body).Decode(&retData)
	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statusCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	fmt.Println(err)

	db.First(&pegawai, "nip = ?", nip)

	updatas := &model.UpdateAngkaKreditBKN{
		PnsId:                  *pegawai.NoRefBkn,
		BulanMulaiPenilaian:    fmt.Sprintf("%v", int(retData.Tmulai.Month())),
		TahunMulaiPenailan:     fmt.Sprintf("%v", retData.Tmulai.Year()),
		BulanSelesiaiPenilaian: fmt.Sprintf("%v", int(retData.Tselesai.Month())),
		TahunSelesaiPenailan:   fmt.Sprintf("%v", retData.Tselesai.Year()),
		KreditBaruTotal:        fmt.Sprintf("%v", *retData.Total),
		KreditPenunjangBaru:    fmt.Sprintf("%v", *retData.Tambahan),
		KreditUtamaBaru:        fmt.Sprintf("%v", retData.Utama),
		NomorSk:                *retData.Nsk,
		TanggalSk:              fmt.Sprintf("%v", retData.Tsk.Format("02-01-2006")),
	}

	if retData.Jns == "Pertama" {
		updatas.IsAngkaKreditPertama = "1"
	} else if retData.Jns == "Integrasi" {
		updatas.IsIntegrasi = "1"
	} else if retData.Jns == "Konversi" {
		updatas.IsKonversi = "1"
	}

	//start update
	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/angkakredit/save")

	out, err := json.Marshal(updatas)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "A",
			"errors":     err.Error(),
		})
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewReader(out))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "B",
			"errors":     err.Error(),
		})
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	d := fmt.Sprintf("bearer %v", *singkronisasi.TokenApimanager)
	e := fmt.Sprintf("Bearer %v", *singkronisasi.TokenSso)

	//fmt.Println(d)
	//fmt.Println(e)

	req.Header.Add("Auth", d)
	req.Header.Add("Authorization", e)

	res, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "C",
			"errors":     err.Error(),
		})
	}

	type ReturnBKNStatus struct {
		Code        string `json:"code" default:"0"`
		Description string `json:"description"`
		Message     string `json:"message"`
		Success     bool   `json:"success"`
		MapData     struct {
			RwAngkaKreditId string `json:"rwAngkaKreditId"`
		} `json:"mapData"`
	}

	defer res.Body.Close()

	var respData ReturnBKNStatus

	err = json.NewDecoder(res.Body).Decode(&respData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "D",
			"errors":     err.Error(),
		})
	}

	fmt.Println(respData)

	//log data
	/*code := respData["code"].(string)*/
	if respData.Success {
		respData.Code = "1"
	}

	riwayat_update_bkn := &model.RiwayatUpdateBKN{
		Deskripsi: "Riwayat Angka Kredit",
		CreatedBy: fmt.Sprint(c.Get("nip")),
		Code:      respData.Code,
		Content:   string(out),
		Message:   respData.Message,
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       respData,
		"statusCode": http.StatusOK,
	})
}

func UpdateSingkronisasiPutDataBknCpnsPns(c echo.Context) error {
	db, _ := model.CreateCon()

	nip := c.Param("nip")
	host := "bkn"

	var pegawai model.Pegawai
	var singkronisasi model.Singkronisasi
	cpns := make(map[string]string)

	pegawai.Nip = nip
	knpang := []int{211, 212}

	db.Debug().Preload("RiwayatPangkat", "knpang IN (?)", knpang).Find(&pegawai)

	cpns["pns_orang_id"] = *pegawai.NoRefBkn
	riwayat_pangkat := pegawai.RiwayatPangkat

	for _, value := range riwayat_pangkat {
		if value.Knpang == 211 {
			//cpns["nama_jabatan_angkat_cpns"] = value.Tntbakn
			cpns["nomor_sk_cpns"] = *value.Nskpang
			cpns["tgl_sk_cpns"] = value.Tskpang.Format("02-01-2006")
			//cpns["nomor_spmt"] = DerefString(value.Nlgas)
			if value.Nlgas != nil {
				cpns["nomor_spmt"] = *value.Nlgas
			}
			//cpns["nomor_sttpl"] = DerefString(value.Nsttpp)
			if value.Nsttpp != nil {
				cpns["nomor_sttpl"] = *value.Nsttpp
			}
			//cpns["tgl_sttpl"] = DerefTime(value.Tsttpp)
			if value.Tsttpp != nil {
				cpns["tgl_sttpl"] = value.Tsttpp.Format("02-01-2006")
			}
			//cpns["pertek_cpns_pns_l2th_nomor"] = value.Nntbakn
			if value.Nntbakn != "" {
				cpns["pertek_cpns_pns_l2th_nomor"] = value.Nntbakn
			}
			//cpns["pertek_cpns_pns_l2th_tanggal"] = DerefTime(value.Tntbakn)
			if value.Tntbakn != nil {
				cpns["pertek_cpns_pns_l2th_tanggal"] = value.Tntbakn.Format("02-01-2006")
			}
			cpns["status_cpns_pns"] = "cpns"
		}

		if value.Knpang == 212 {
			//cpns["nomor_sk_pns"] = *value.Nskpang
			if value.Nskpang != nil {
				cpns["nomor_sk_pns"] = *value.Nskpang
			}
			//cpns["tgl_sk_pns"] = DerefTime(value.Tskpang)
			if value.Tskpang != nil {
				cpns["tgl_sk_pns"] = value.Tskpang.Format("02-01-2006")
			}
			//cpns["tmt_pns"] = DerefTime(value.Tmtpang)
			if value.Tmtpang != nil {
				cpns["tmt_pns"] = value.Tmtpang.Format("02-01-2006")
			}
			//cpns["tanggal_dokter_pns"] = value.Nsttpp
			cpns["status_cpns_pns"] = "pns"
		}
	}

	//start update
	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/cpns/save")

	out, err := json.Marshal(cpns)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "A",
			"errors":     err.Error(),
		})
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewReader(out))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "B",
			"errors":     err.Error(),
		})
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	d := fmt.Sprintf("bearer %v", *singkronisasi.TokenApimanager)
	e := fmt.Sprintf("Bearer %v", *singkronisasi.TokenSso)

	fmt.Println(d)
	fmt.Println(e)

	req.Header.Add("Auth", d)
	req.Header.Add("Authorization", e)

	res, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "C",
			"errors":     err.Error(),
		})
	}

	defer res.Body.Close()

	/*bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	respData := bodyString*/

	type ReturnBKNStatus struct {
		Code    int    `json:"code"`
		Message string `json:"data"`
	}

	var respData ReturnBKNStatus

	err = json.NewDecoder(res.Body).Decode(&respData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "D",
			"errors":     err.Error(),
		})
	}

	riwayat_update_bkn := &model.RiwayatUpdateBKN{
		Deskripsi: "Riwayat CPNS/PNS",
		CreatedBy: fmt.Sprint(c.Get("nip")),
		Code:      fmt.Sprint(respData.Code),
		Content:   string(out),
		Message:   respData.Message,
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       respData,
		"statusCode": http.StatusOK,
	})
}

func DeleteSingkronisasiDelDataBknAngkaKredit(c echo.Context) error {
	db, _ := model.CreateCon()

	id := c.Param("id")
	var singkronisasi model.Singkronisasi
	host := "bkn"

	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/angkakredit/delete/%s", id)

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "B",
			"errors":     err.Error(),
		})
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	d := fmt.Sprintf("bearer %v", *singkronisasi.TokenApimanager)
	e := fmt.Sprintf("Bearer %v", *singkronisasi.TokenSso)

	//fmt.Println(d)
	//fmt.Println(e)

	req.Header.Add("Auth", d)
	req.Header.Add("Authorization", e)

	res, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "C",
			"errors":     err.Error(),
		})
	}
	defer res.Body.Close()

	var respData map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&respData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "D",
			"errors":     err.Error(),
		})
	}

	//log data
	code := "0"
	if respData["success"].(bool) {
		code = "1"
	}
	riwayat_update_bkn := &model.RiwayatUpdateBKN{
		Deskripsi: "Delete Riwayat Angka Kredit",
		CreatedBy: fmt.Sprint(c.Get("nip")),
		Code:      code,
		Message:   respData["message"].(string),
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create
	//end doupdate

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       respData,
		"statusCode": http.StatusOK,
	})
}

func UpdateSingkronisasiPutUnorJabatan(c echo.Context) error {
	//var retData model.RiwayatJabatan

	jabatan := make(map[string]string)

	err := json.NewDecoder(c.Request().Body).Decode(&jabatan)
	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statusCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	//fmt.Println(jabatan)

	db, _ := model.CreateCon()

	nip := c.Param("nip")

	var pegawai model.Pegawai
	var singkronisasi model.Singkronisasi
	host := "bkn"

	db.First(&pegawai, "nip = ?", nip)
	jabatan["pnsId"] = *pegawai.NoRefBkn

	fmt.Println(jabatan)

	//start update
	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/jabatan/unorjabatan/save")

	out, err := json.Marshal(jabatan)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "A",
			"errors":     err.Error(),
		})
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewReader(out))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "B",
			"errors":     err.Error(),
		})
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	d := fmt.Sprintf("bearer %v", *singkronisasi.TokenApimanager)
	e := fmt.Sprintf("Bearer %v", *singkronisasi.TokenSso)

	//fmt.Println(d)
	//fmt.Println(e)

	req.Header.Add("Auth", d)
	req.Header.Add("Authorization", e)

	res, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "C",
			"errors":     err.Error(),
		})
	}

	type ReturnBKNStatus struct {
		Code        string `json:"code" default:"0"`
		Description string `json:"description"`
		Message     string `json:"message"`
		Success     bool   `json:"success"`
		MapData     struct {
			RwJabatanId string `json:"rwJabatanId"`
		} `json:"mapData"`
	}

	defer res.Body.Close()

	var respData ReturnBKNStatus

	err = json.NewDecoder(res.Body).Decode(&respData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "D",
			"errors":     err.Error(),
		})
	}

	if respData.Success {
		respData.Code = "1"
	}

	riwayat_update_bkn := &model.RiwayatUpdateBKN{
		Deskripsi: fmt.Sprintf("Riwayat Jabatan %s", jabatan["nomorSk"]),
		CreatedBy: fmt.Sprint(c.Get("nip")),
		Code:      respData.Code,
		Content:   string(out),
		Message:   respData.Message,
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create*/

	if respData.Code == "1" && jabatan["id"] == "" {
		//db.Model(&model.RiwayatJabatan{}).Where().Update("idSync", respData.MapData.RwJabatanId)
		db.Model(&model.RiwayatJabatan{}).Debug().Where("nip = ? and tmtjab = ? and tmt_mutasi = ?", jabatan["nip"], jabatan["tmtJabatan"], jabatan["tmtMutasi"]).Update("idSync", respData.MapData.RwJabatanId)
	}

	//respData := bodyString

	/*bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	respData := bodyString*/

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       respData,
		"statusCode": http.StatusOK,
	})
}

func DeleteSingkronisasiDelRiwJabatan(c echo.Context) error {
	db, _ := model.CreateCon()

	id := c.Param("id")
	var singkronisasi model.Singkronisasi
	host := "bkn"

	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/jabatan/delete/%s", id)

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "B",
			"errors":     err.Error(),
		})
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	d := fmt.Sprintf("bearer %v", *singkronisasi.TokenApimanager)
	e := fmt.Sprintf("Bearer %v", *singkronisasi.TokenSso)

	//fmt.Println(d)
	//fmt.Println(e)

	req.Header.Add("Auth", d)
	req.Header.Add("Authorization", e)

	res, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "C",
			"errors":     err.Error(),
		})
	}
	defer res.Body.Close()

	var respData map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&respData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "D",
			"errors":     err.Error(),
		})
	}

	//log data
	code := "0"
	if respData["success"].(bool) {
		code = "1"
	}
	riwayat_update_bkn := &model.RiwayatUpdateBKN{
		Deskripsi: "Delete Riwayat Jabatan",
		CreatedBy: fmt.Sprint(c.Get("nip")),
		Code:      code,
		Message:   respData["message"].(string),
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create
	//end doupdate

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       respData,
		"statusCode": http.StatusOK,
	})
}
