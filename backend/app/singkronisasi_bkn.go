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
	db, _ := model.CreateCon()

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
		retVal := getTokenSSO(updateTokens.Ckey, updateTokens.Csecret)

		fmt.Println("A")

		now := time.Now()
		ExpiresIn := now.Add(time.Duration(retVal.ExpiresIn) * time.Second)

		singkronisasi := model.Singkronisasi{
			TokenSso:        &retVal.AccessToken,
			TokenSsoExpired: &ExpiresIn,
			UpdatedBy:       fmt.Sprint(c.Get("nip")),
		}

		db.Model(&model.Singkronisasi{}).Where("id = '1'").Updates(&singkronisasi)

		//return c.JSON(http.StatusOK, singkronisasi)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":       singkronisasi,
			"statucCode": http.StatusOK,
		})

	} else if updateTokens.Type == "token_apimanager" {
		retVal := getTokenApiManager(updateTokens.ClientId, updateTokens.Username, updateTokens.Password)

		fmt.Println("B")

		now := time.Now()
		ExpiresIn := now.Add(time.Duration(retVal.ExpiresIn) * time.Second)

		singkronisasi := model.Singkronisasi{
			TokenApimanager:        &retVal.AccessToken,
			TokenApimanagerExpired: &ExpiresIn,
			UpdatedBy:              fmt.Sprint(c.Get("nip")),
		}

		db.Model(&model.Singkronisasi{}).Where("id = '1'").Updates(&singkronisasi)

		//return c.JSON(http.StatusOK, singkronisasi)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":       singkronisasi,
			"statucCode": http.StatusOK,
		})
	}

	return c.JSON(http.StatusOK, nil)
}

func getTokenSSO(Ckey, Csecret string) (res model.ResultTokenSso) {
	payload := strings.NewReader("grant_type=client_credentials")

	auth := fmt.Sprintf("%s:%s", Ckey, Csecret)
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

	return
}

func getTokenApiManager(ClientId, Username, Password string) (res model.ResultTokenSso) {
	payload := strings.NewReader(fmt.Sprintf("client_id=%s&grant_type=password&username=%s&password=%s", ClientId, Username, Password))
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

	return
}

func GetSingkronisasiGetDataBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	var singkronisasi model.Singkronisasi

	nip := c.Param("nip")
	page := c.Param("page")
	host := "bkn"

	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/pns/%s/%s", page, nip)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("accept", "application/json")
	d := fmt.Sprintf("bearer %v", *singkronisasi.TokenApimanager)
	e := fmt.Sprintf("Bearer %v", *singkronisasi.TokenSso)

	fmt.Println(d)
	fmt.Println(e)

	req.Header.Add("Auth", d)
	req.Header.Add("Authorization", e)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	var retData map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&retData)
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       retData,
		"statucCode": http.StatusOK,
	})
}

func UpdateSingkronisasiPutDataBkn(c echo.Context) error {
	db, _ := model.CreateCon()

	validate := validator.New()

	var updatas model.UpdateBKNDataUtama
	var singkronisasi model.Singkronisasi

	page := c.Param("page")
	host := "bkn"

	err := json.NewDecoder(c.Request().Body).Decode(&updatas)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(updatas); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	//start doupdate
	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/pns/%s-update", page)

	out, err := json.Marshal(updatas)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"pos":        "A",
			"errors":     err.Error(),
		})
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewReader(out))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
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
			"statucCode": http.StatusBadRequest,
			"pos":        "C",
			"errors":     err.Error(),
		})
	}
	defer res.Body.Close()

	var retData map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&retData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"pos":        "D",
			"errors":     err.Error(),
		})
	}

	//end doupdate

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       retData,
		"statucCode": http.StatusOK,
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
			"statucCode": http.StatusNotImplemented,
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
			"statucCode": http.StatusBadRequest,
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
			"statucCode": http.StatusBadRequest,
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
			"statucCode": http.StatusBadRequest,
			"pos":        "C",
			"errors":     err.Error(),
		})
	}
	defer res.Body.Close()

	// Writer the body to file
	retData, err := io.Copy(out, res.Body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
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
		"statucCode": http.StatusOK,
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
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

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
			"statucCode": http.StatusBadRequest,
			"pos":        "A",
			"errors":     err.Error(),
		})
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewReader(out))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
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
			"statucCode": http.StatusBadRequest,
			"pos":        "C",
			"errors":     err.Error(),
		})
	}
	defer res.Body.Close()

	var respData map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&respData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"pos":        "D",
			"errors":     err.Error(),
		})
	}

	//end doupdate

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       respData,
		"statucCode": http.StatusOK,
	})
}

func DeleteSingkronisasiDelDataBknAngkaKredit(c echo.Context) error {
	db, _ := model.CreateCon()

	id := c.Param("id")
	var singkronisasi model.Singkronisasi
	host := "bkn"

	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)

	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/angkakredit/delete/%s", id)

	/*out, err := json.Marshal(updatas)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"pos":        "A",
			"errors":     err.Error(),
		})
	}*/

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
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
			"statucCode": http.StatusBadRequest,
			"pos":        "C",
			"errors":     err.Error(),
		})
	}
	defer res.Body.Close()

	var respData map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&respData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"pos":        "D",
			"errors":     err.Error(),
		})
	}

	//end doupdate

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       respData,
		"statucCode": http.StatusOK,
	})

}
