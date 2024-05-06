package app

import (
	_ "database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	model "newSimpegAPI/model"
	"strings"
	"time"
)

func GetSingkronisasi(searchString model.SearchSingkronisasi) (singkronisasi []model.Singkronisasi, result *gorm.DB) {
	db, _ := model.CreateCon()

	result = db.Model(&model.Singkronisasi{})

	//id
	if searchString.Id != "" {
		result = result.Where("master_singkronisasi.id = ?", searchString.Id)
	}

	//host
	if searchString.Host != "" {
		result = result.Where("master_singkronisasi.host = ?", searchString.Host)
	}

	result = result.Find(&singkronisasi)

	return
}

func FindSingkronisasi(c echo.Context) error {
	var searchString model.SearchSingkronisasi

	json.NewDecoder(c.Request().Body).Decode(&searchString)

	singkronisasi, result := GetSingkronisasi(searchString)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       singkronisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func CreateSingkronisasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var singkronisasi model.Singkronisasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&singkronisasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(singkronisasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	singkronisasi.CreatedBy = fmt.Sprint(c.Get("nip"))

	result := db.Create(&singkronisasi)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       singkronisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdateSingkronisasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var singkronisasi model.Singkronisasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&singkronisasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(singkronisasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	singkronisasi.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Singkronisasi{}).Where("id = ?", singkronisasi.Id).Updates(&singkronisasi)

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
		"data":       singkronisasi,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeleteSingkronisasi(c echo.Context) error {
	db, _ := model.CreateCon()

	var singkronisasi model.DeleteSingkronisasi
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&singkronisasi)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(singkronisasi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Singkronisasi{}).Where("id = ?", singkronisasi.Id).Delete(&singkronisasi)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       singkronisasi,
		"statucCode": http.StatusOK,
	})
}

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

	fmt.Println("C")

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
