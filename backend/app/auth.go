package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lpernett/godotenv"
	"gorm.io/gorm/clause"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	model "newSimpegAPI/model"
)

var (
	SIAP_AUTH_API       string
	SIAP_REST_API_TOKEN string
	APP_KEY             string
)

type JwtCustomClaims struct {
	Nip     string `json:"nip"`
	TokenId string `json:"token_id"`
	Gid     int    `json:"gid"`
	jwt.RegisteredClaims
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	SIAP_AUTH_API = os.Getenv("SIAP_AUTH_API")
	//SIAP_AUTH_API = os.Getenv("SIAP_AUTH_API_DEV")
	SIAP_REST_API_TOKEN = os.Getenv("SIAP_REST_API_TOKEN")
	APP_KEY = os.Getenv("APP_KEY")
}

func DerefString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

func DerefTime(s *time.Time) string {
	if s != nil {
		return s.Format("02-01-2006")
	}

	return ""
}

func UserInfo(nip, token_id string) (result model.UsersToken, RowsAffected int64) {
	db, _ := model.CreateCon()
	r := db.Model(&model.UsersToken{}).First(&result, "nip = ? and token_id = ?", nip, token_id)

	RowsAffected = r.RowsAffected

	return
}

func Login(c echo.Context) error {
	ret, err := siapLogin(c)

	token_id := uuid.New().String()
	expires_at := time.Now().Add(time.Hour * 24)

	if err != nil {
		return err
	}

	// Throws unauthorized error
	if !ret.LoginStatus {
		return echo.ErrUnauthorized
	}

	ti := time.Now()
	ret.Datauser.LoginAt = &ti
	ret.Datauser.ExpiresAt = &expires_at
	ret.Datauser.TokenId = token_id

	//var isAdmin int64
	ret.Datauser = saveloginInfo(ret.Datauser)
	fmt.Println(ret.Datauser.Gid)

	// Set custom claims
	claims := &JwtCustomClaims{
		ret.Datauser.Nip,
		token_id,
		ret.Datauser.Gid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expires_at),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(APP_KEY))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token":      t,
		"datauser":   ret.Datauser,
		"statusCode": http.StatusOK,
	})
}

func Logout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	logout := time.Now()

	db, _ := model.CreateCon()
	db.Model(&model.UsersToken{}).Where("nip = ? and token_id = ?", claims.Nip, claims.TokenId).Updates(map[string]interface{}{"token_id": "", "logout_at": &logout})

	return c.JSON(http.StatusOK, echo.Map{
		"nip":        claims.Nip,
		"message":    "Logout Success",
		"statusCode": http.StatusOK,
	})
}

func IsLogin(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	//dtuser := claims.Datauser

	result, _ := UserInfo(claims.Nip, claims.TokenId)

	return c.JSON(http.StatusOK, echo.Map{
		"datauser":   result,
		"statusCode": http.StatusOK,
	})
}

func AuthorizeRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Validating User")

		if strings.HasSuffix(c.Path(), "/login") {
			return next(c)
		}

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtCustomClaims)
		//dtuser := claims.Datauser

		_, ra := UserInfo(claims.Nip, claims.TokenId)
		if ra == 0 {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message":    "Request Unauthorized",
				"statusCode": http.StatusUnauthorized,
			})
		}
		c.Set("nip", claims.Nip)
		c.Set("gid", claims.Gid)
		return next(c)
	}
}

func saveloginInfo(dtuser model.UsersToken) model.UsersToken {
	db, _ := model.CreateCon()

	db.Model(&model.UsersToken{}).Clauses(clause.Returning{}, clause.OnConflict{
		Columns: []clause.Column{{Name: "nip"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"id_opd":     dtuser.IdOpd,
			"parent_opd": dtuser.ParentOpd,
			"token_id":   dtuser.TokenId,
			"login_at":   dtuser.LoginAt,
			"expires_at": dtuser.ExpiresAt,
		}),
	}).Create(&dtuser)

	//var user model.Users
	db.Model(&model.UsersToken{}).Where("nip = ? and stts = ?", dtuser.Nip, 1).First(&dtuser)

	//dtuser.Gid = user.Gid

	//fmt.Println(user)
	fmt.Println(dtuser.Gid)

	return dtuser
}

func siapLogin(c echo.Context) (res model.SIAPResponse, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", c.FormValue("username"))
	_ = writer.WriteField("password", c.FormValue("password"))
	err = writer.Close()
	if err != nil {
		//fmt.Println("=>?", err)
		return
	}

	url := fmt.Sprintf(`%s/loginUser`, SIAP_AUTH_API)

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header["token"] = []string{SIAP_REST_API_TOKEN}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func FakeLogin(c echo.Context) error {
	/*ret := &model.SIAPResponse{
		Message:     "",
		Success:     true,
		LoginStatus: true,
		Datauser: model.UsersToken{
			Nip:       "198602162008031001",
			Nama:      "Muhammad Taufik Hidayat",
			IdOpd:     "125",
			ParentOpd: "125",
		},
	}*/

	token_id := uuid.New().String()
	expires_at := time.Now().Add(time.Hour * 24)

	fmt.Println(token_id, expires_at)

	// Set custom claims
	/*claims := &JwtCustomClaims{
		"198602162008031001",
		token_id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expires_at),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("simpeg2023"))
	if err != nil {
		return err
	}

	ti := time.Now()
	ret.Datauser.LoginAt = &ti
	ret.Datauser.ExpiresAt = &expires_at
	ret.Datauser.TokenId = token_id

	var isAdmin int64
	ret.Datauser, isAdmin = saveloginInfo(ret.Datauser)*/

	return c.JSON(http.StatusOK, echo.Map{

		"datauser": map[string]interface{}{
			"nip":        "198602162008031001",
			"nama":       "Muhammad Taufik Hidayat",
			"id_opd":     "125",
			"parent_opd": "125",
			"token_id":   token_id,
		},
		"is_admin":   1,
		"statusCode": http.StatusOK,
	})
}

func CheckLevel(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Get("gid"))
		return next(c)
	}
}
