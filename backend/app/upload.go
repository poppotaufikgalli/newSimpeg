package app

import (
	"bytes"
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	model "newSimpegAPI/model"
	"os"
	"path/filepath"
)

func UploadPegawai(c echo.Context) error {
	nip := c.Param("nip")
	filename := c.FormValue("filename")
	path := c.FormValue("path")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	filename = fmt.Sprintf("%s%s", filename, filepath.Ext(file.Filename))
	newpath := filepath.Join("assets", path, filename)
	if err = os.MkdirAll(filepath.Dir(newpath), os.ModeDir|0775); err != nil {
		return err
	}

	dst, err := os.Create(newpath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	if nip != "" && path == "photo" {
		db, _ := model.CreateCon()
		result := db.Model(&model.Pegawai{}).Where("nip = ?", nip).Update("file_bmp", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadPangkat(c echo.Context) error {
	nip := c.Param("nip")
	tmtpang := c.FormValue("tmtpang")
	kgolru := c.FormValue("kgolru")
	knpang := c.FormValue("knpang")
	updateField := c.FormValue("updateField")
	filename := c.FormValue("filename")
	path := c.FormValue("path")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	filename = fmt.Sprintf("%s%s", filename, filepath.Ext(file.Filename))
	newpath := filepath.Join("assets", path, filename)
	if err = os.MkdirAll(filepath.Dir(newpath), os.ModeDir|0775); err != nil {
		return err
	}

	dst, err := os.Create(newpath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	if nip != "" {
		db, _ := model.CreateCon()

		result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and tmtpang = ? and kgolru = ? and knpang = ?", nip, tmtpang, kgolru, knpang).Update(updateField, filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadJabatan(c echo.Context) error {
	nip := c.Param("nip")
	id := c.FormValue("id")
	fieldName := c.FormValue("field_name")
	filename := c.FormValue("filename")
	path := c.FormValue("path")

	id_ref_dokumen := c.FormValue("id_ref_dokumen")
	idSync := c.FormValue("idSync")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	filename = fmt.Sprintf("%s%s", filename, filepath.Ext(file.Filename))
	newpath := filepath.Join("assets", path, filename)
	if err = os.MkdirAll(filepath.Dir(newpath), os.ModeDir|0775); err != nil {
		return err
	}

	dst, err := os.Create(newpath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	if nip != "" {
		db, _ := model.CreateCon()

		result := db.Model(&model.RiwayatJabatan{}).Where("nip = ? and id = ?", nip, id).Update(fieldName, filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}

		if c.Get("gid").(int) == 1 {
			status_bkn, err := sendDokRwBkn(filename, newpath, id_ref_dokumen, idSync, c.Get("nip").(string))

			if err != nil {
				return c.JSON(http.StatusNotImplemented, map[string]interface{}{
					"statusCode": http.StatusNotImplemented,
					"errors":     err.Error(),
					"data":       err.Error(),
					"title":      fmt.Sprintf("Gagal Upload Dokumen %s ke BKN", fieldName),
				})
			}

			if fmt.Sprint(status_bkn["code"]) != "1" {
				return c.JSON(http.StatusNotImplemented, map[string]string{
					"title":       "Update BKN Gagal",
					"description": fmt.Sprintf("%s-%s", status_bkn["code"], status_bkn["message"]),
				})
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func sendDokRwBkn(filename, path, id_ref_dokumen, idSync, CreatedBy string) (status_bkn map[string]interface{}, err error) {
	db, _ := model.CreateCon()

	fmt.Println(id_ref_dokumen, idSync)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	src, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer src.Close()

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = io.Copy(part, src)
	if err != nil {
		return
	}

	url := "https://apimws.bkn.go.id:8243/apisiasn/1.0/upload-dok-rw"

	_ = writer.WriteField("id_riwayat", idSync)
	_ = writer.WriteField("id_ref_dokumen", id_ref_dokumen)

	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Set("Content-Type", writer.FormDataContentType())
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

	riwayat_update_bkn := &model.RiwayatUpdateBKN{
		Deskripsi: fmt.Sprintf("Upload Dokuem Riwayat Jabatan %s", filename),
		CreatedBy: CreatedBy,
		Code:      fmt.Sprint(status_bkn["code"]),
		Message:   fmt.Sprint(status_bkn["message"]),
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn) // pass pointer of data to Create*/

	return
}

func DownloadDokRwBKN(c echo.Context) error {
	//db, _ := model.CreateCon()

	docData := make(map[string]string)

	err := json.NewDecoder(c.Request().Body).Decode(&docData)
	//	fmt.Println(docData)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"title":       "Gagal Download Dokumen BKN",
			"status":      http.StatusNotImplemented,
			"errors":      err.Error(),
			"description": "Gagal decode data input",
		})
	}

	//start doupdate
	// Create the file
	filename := filepath.Base(fmt.Sprintf("/%s", docData["path"]))

	newpath := filepath.Join("assets", "dokumen", docData["nip"], filename)
	out, err := os.Create(newpath)

	if err = os.MkdirAll(filepath.Dir(newpath), os.ModeDir|0775); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"title":       "Gagal Download Dokumen BKN",
			"status":      http.StatusNotImplemented,
			"errors":      err.Error(),
			"description": "Gagal membuat direktori",
		})
	}

	defer out.Close()

	//db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)
	url := fmt.Sprintf("https://apimws.bkn.go.id:8243/apisiasn/1.0/download-dok?filePath=%s", url.QueryEscape(docData["path"]))

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"title":       "Gagal Download Dokumen BKN",
			"status":      http.StatusNotImplemented,
			"errors":      err.Error(),
			"description": "Gagal get dari url BKN",
		})
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	a := getTokenApiManager(fmt.Sprint(c.Get("nip")))
	b := getTokenSSO(fmt.Sprint(c.Get("nip")))
	d := fmt.Sprintf("bearer %v", a.AccessToken)
	e := fmt.Sprintf("Bearer %v", b.AccessToken)

	//fmt.Println(d)
	//fmt.Println(e)

	req.Header.Add("Auth", d)
	req.Header.Add("Authorization", e)

	res, err := client.Do(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"title":       "Gagal Download Dokumen BKN",
			"status":      http.StatusNotImplemented,
			"errors":      err.Error(),
			"description": "Gagal get dari url BKN 2",
		})
	}

	defer res.Body.Close()

	size, err := io.Copy(out, res.Body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"title":       "Gagal Download Dokumen BKN",
			"status":      http.StatusNotImplemented,
			"errors":      err.Error(),
			"description": "Gagal copy body ke out",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"size":     size,
			"filename": filename,
			"path":     newpath,
		},
		"statusCode": http.StatusOK,
	})
}

func UploadTugasTambahan(c echo.Context) error {
	nip := c.Param("nip")
	tmtjab := c.FormValue("tmtjab")
	filename := c.FormValue("filename")
	path := c.FormValue("path")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	filename = fmt.Sprintf("%s%s", filename, filepath.Ext(file.Filename))
	newpath := filepath.Join("assets", path, filename)
	if err = os.MkdirAll(filepath.Dir(newpath), os.ModeDir|0775); err != nil {
		return err
	}

	dst, err := os.Create(newpath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	if nip != "" {
		db, _ := model.CreateCon()
		//result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and tmtpang = ? and kgolru = ? and knpang = ?", nip, tmtpang, kgolru, knpang).Update("filename", filename)
		result := db.Model(&model.RiwayatTugasTambahan{}).Where("nip = ? and tmtjab = ?", nip, tmtjab).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadKGB(c echo.Context) error {
	nip := c.Param("nip")
	tmtngaj := c.FormValue("tmtngaj")
	filename := c.FormValue("filename")
	path := c.FormValue("path")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	filename = fmt.Sprintf("%s%s", filename, filepath.Ext(file.Filename))
	newpath := filepath.Join("assets", path, filename)
	if err = os.MkdirAll(filepath.Dir(newpath), os.ModeDir|0775); err != nil {
		return err
	}

	dst, err := os.Create(newpath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	if nip != "" {
		db, _ := model.CreateCon()
		//result := db.Model(&model.RiwayatTugasTambahan{}).Where("nip = ? and tmtjab = ?", nip, tmtjab).Update("filename", filename)
		result := db.Model(&model.RiwayatGajiberkala{}).Where("nip = ? and tmtngaj = ?", nip, tmtngaj).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadAK(c echo.Context) error {
	nip := c.Param("nip")
	tmulai := c.FormValue("tmulai")
	filename := c.FormValue("filename")
	path := c.FormValue("path")
	idSync := c.FormValue("idSync")
	id_ref_dokumen := c.FormValue("id_ref_dokumen")
	fieldName := c.FormValue("field_name")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	filename = fmt.Sprintf("%s%s", filename, filepath.Ext(file.Filename))
	newpath := filepath.Join("assets", path, filename)
	if err = os.MkdirAll(filepath.Dir(newpath), os.ModeDir|0775); err != nil {
		return err
	}

	dst, err := os.Create(newpath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	if nip != "" {
		db, _ := model.CreateCon()
		result := db.Model(&model.RiwayatAngkaKredit{}).Where("nip = ? and tmulai = ?", nip, tmulai).Update(fieldName, filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}

		if c.Get("gid").(int) == 1 {
			status_bkn, err := sendDokRwBkn(filename, newpath, id_ref_dokumen, idSync, c.Get("nip").(string))

			fmt.Println(fmt.Sprint(status_bkn["code"]))

			if err != nil {
				return c.JSON(http.StatusNotImplemented, map[string]interface{}{
					"statusCode": http.StatusNotImplemented,
					"errors":     err.Error(),
					"data":       err.Error(),
					"title":      fmt.Sprintf("Gagal Upload Dokumen %s ke BKN", fieldName),
				})
			}

			if fmt.Sprint(status_bkn["code"]) != "1" {
				return c.JSON(http.StatusNotImplemented, map[string]string{
					"title":       "Update BKN Gagal",
					"description": fmt.Sprintf("%s-%s", status_bkn["code"], status_bkn["message"]),
				})
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadSTLUD(c echo.Context) error {
	nip := c.Param("nip")
	kstlud := c.FormValue("kstlud")
	tstlud := c.FormValue("tstlud")

	filename := c.FormValue("filename")
	path := c.FormValue("path")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	filename = fmt.Sprintf("%s%s", filename, filepath.Ext(file.Filename))
	newpath := filepath.Join("assets", path, filename)
	if err = os.MkdirAll(filepath.Dir(newpath), os.ModeDir|0775); err != nil {
		return err
	}

	dst, err := os.Create(newpath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	if nip != "" {
		db, _ := model.CreateCon()
		//result := db.Model(&model.RiwayatAngkaKredit{}).Where("nip = ? and tmulai = ?", nip, tmulai).Update("filename", filename)
		result := db.Model(&model.RiwayatStlud{}).Where("nip = ? and kstlud = ? and tstlud =?", nip, kstlud, tstlud).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadPMK(c echo.Context) error {
	nip := c.Param("nip")
	tanggalAwal := c.FormValue("tanggalAwal")

	filename := c.FormValue("filename")
	path := c.FormValue("path")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	filename = fmt.Sprintf("%s%s", filename, filepath.Ext(file.Filename))
	newpath := filepath.Join("assets", path, filename)
	if err = os.MkdirAll(filepath.Dir(newpath), os.ModeDir|0775); err != nil {
		return err
	}

	dst, err := os.Create(newpath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	if nip != "" {
		db, _ := model.CreateCon()
		result := db.Model(&model.RiwayatPmk{}).Where("nip = ? and tanggalAwal = ?", nip, tanggalAwal).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadPI(c echo.Context) error {
	nip := c.Param("nip")
	tmtPi := c.FormValue("tmtPi")

	filename := c.FormValue("filename")
	path := c.FormValue("path")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	filename = fmt.Sprintf("%s%s", filename, filepath.Ext(file.Filename))
	newpath := filepath.Join("assets", path, filename)
	if err = os.MkdirAll(filepath.Dir(newpath), os.ModeDir|0775); err != nil {
		return err
	}

	dst, err := os.Create(newpath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	if nip != "" {
		db, _ := model.CreateCon()
		result := db.Model(&model.RiwayatPindahInstansi{}).Where("nip = ? and tmtPi = ?", nip, tmtPi).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadSKP(c echo.Context) error {
	nip := c.Param("nip")
	tahun := c.FormValue("tahun")
	id := c.FormValue("id")

	id_ref_dokumen := c.FormValue("id_ref_dokumen")
	idSync := c.FormValue("idSync")

	filename, newpath, _ := doUpload(c)

	if nip != "" {
		db, _ := model.CreateCon()
		result := db.Model(&model.RiwayatSkp{}).Where("nip = ? and tahun = ? and id=?", nip, tahun, id).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}

		if c.Get("gid").(int) == 1 {
			status_bkn, err := sendDokRwBkn(filename, newpath, id_ref_dokumen, idSync, c.Get("nip").(string))

			fmt.Println(fmt.Sprint(status_bkn["code"]))

			if err != nil {
				return c.JSON(http.StatusNotImplemented, map[string]interface{}{
					"statusCode": http.StatusNotImplemented,
					"errors":     err.Error(),
					"data":       err.Error(),
					"title":      fmt.Sprintf("Gagal Upload Dokumen ke BKN"),
				})
			}

			if fmt.Sprint(status_bkn["code"]) != "1" {
				return c.JSON(http.StatusNotImplemented, map[string]string{
					"title":       "Update BKN Gagal",
					"description": fmt.Sprintf("%v-%s", status_bkn["code"], status_bkn["message"]),
				})
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadKinerja(c echo.Context) error {
	nip := c.Param("nip")
	thnilai := c.FormValue("thnilai")
	id := c.FormValue("id")
	id_ref_dokumen := c.FormValue("id_ref_dokumen")
	idSync := c.FormValue("idSync")

	filename, newpath, _ := doUpload(c)

	if nip != "" {
		db, _ := model.CreateCon()
		result := db.Model(&model.RiwayatKinerja{}).Where("nip = ? and thnilai = ? and id =?", nip, thnilai, id).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}

		if c.Get("gid").(int) == 1 {
			status_bkn, err := sendDokRwBkn(filename, newpath, id_ref_dokumen, idSync, c.Get("nip").(string))

			fmt.Println(fmt.Sprint(status_bkn["code"]))

			if err != nil {
				return c.JSON(http.StatusNotImplemented, map[string]interface{}{
					"statusCode": http.StatusNotImplemented,
					"errors":     err.Error(),
					"data":       err.Error(),
					"title":      fmt.Sprintf("Gagal Upload Dokumen ke BKN"),
				})
			}

			if fmt.Sprint(status_bkn["code"]) != "1" {
				return c.JSON(http.StatusNotImplemented, map[string]string{
					"title":       "Update BKN Gagal",
					"description": fmt.Sprintf("%v-%s", status_bkn["code"], status_bkn["message"]),
				})
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadPenghargaan(c echo.Context) error {
	nip := c.Param("nip")
	nbintang := c.FormValue("nbintang")
	nsk := c.FormValue("nsk")

	filename, newpath, _ := doUpload(c)

	if nip != "" {
		db, _ := model.CreateCon()
		result := db.Model(&model.RiwayatPenghargaan{}).Where("nip = ? and nbintang = ? and nsk =?", nip, nbintang, nsk).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadHukdis(c echo.Context) error {
	nip := c.Param("nip")
	jhukum := c.FormValue("jhukum")
	tmt := c.FormValue("tmt")

	filename, newpath, _ := doUpload(c)

	if nip != "" {
		db, _ := model.CreateCon()
		result := db.Model(&model.RiwayatHukdis{}).Where("nip = ? and jhukum = ? and tmt =?", nip, jhukum, tmt).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadCuti(c echo.Context) error {
	nip := c.Param("nip")
	nsk := c.FormValue("nsk")
	tmulai := c.FormValue("tmulai")

	filename, newpath, _ := doUpload(c)

	if nip != "" {
		db, _ := model.CreateCon()
		result := db.Model(&model.RiwayatCuti{}).Where("nip = ? and nsk = ? and tmulai =?", nip, nsk, tmulai).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadOrganisasi(c echo.Context) error {
	nip := c.Param("nip")
	norg := c.FormValue("norg")
	jborg := c.FormValue("jborg")

	filename, newpath, _ := doUpload(c)

	if nip != "" {
		db, _ := model.CreateCon()
		result := db.Model(&model.RiwayatOrganisasi{}).Where("nip = ? and norg = ? and jborg =?", nip, norg, jborg).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadBahasa(c echo.Context) error {
	nip := c.Param("nip")
	nbahasa := c.FormValue("nbahasa")

	filename, newpath, _ := doUpload(c)

	if nip != "" {
		db, _ := model.CreateCon()
		result := db.Model(&model.RiwayatBahasa{}).Where("nip = ? and nbahasa =?", nip, nbahasa).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadTugasLn(c echo.Context) error {
	nip := c.Param("nip")
	tmulai := c.FormValue("tmulai")

	filename, newpath, _ := doUpload(c)

	if nip != "" {
		db, _ := model.CreateCon()
		result := db.Model(&model.RiwayatTugasLn{}).Where("nip = ? and tmulai = ?", nip, tmulai).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadPendum(c echo.Context) error {
	nip := c.Param("nip")
	ktpu := c.FormValue("ktpu")
	kjur := c.FormValue("kjur")

	filename, newpath, _ := doUpload(c)

	if nip != "" {
		db, _ := model.CreateCon()
		result := db.Model(&model.RiwayatPendum{}).Where("nip = ? and ktpu = ? and kjur =?", nip, ktpu, kjur).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadTubel(c echo.Context) error {
	nip := c.Param("nip")
	id := c.FormValue("id")
	//kjur := c.FormValue("kjur")

	filename, newpath, _ := doUpload(c)

	if nip != "" {
		db, _ := model.CreateCon()
		result := db.Model(&model.RiwayatTubel{}).Where("id = ? and nip = ?", id, nip).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadDiklat(c echo.Context) error {
	nip := c.Param("nip")
	jdiklat := c.FormValue("jdiklat")
	tmulai := c.FormValue("tmulai")

	filename, newpath, _ := doUpload(c)

	if nip != "" {
		db, _ := model.CreateCon()
		result := db.Model(&model.RiwayatDiklat{}).Where("nip = ? and jdiklat = ? and tmulai =?", nip, jdiklat, tmulai).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func UploadKeluarga(c echo.Context) error {
	nip := c.Param("nip")
	jkeluarga := c.FormValue("jkeluarga")
	nama_kel := c.FormValue("nama_kel")

	filename, newpath, _ := doUpload(c)

	if nip != "" {
		db, _ := model.CreateCon()
		result := db.Model(&model.RiwayatKeluarga{}).Where("nip = ? and jkeluarga = ? and nama_kel =?", nip, jkeluarga, nama_kel).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statusCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statusCode": http.StatusOK,
	})
}

func doUpload(c echo.Context) (filename, newpath string, err error) {
	filename = c.FormValue("filename")
	path := c.FormValue("path")

	file, err := c.FormFile("file")
	if err != nil {
		return
	}
	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()

	// Destination
	filename = fmt.Sprintf("%s%s", filename, filepath.Ext(file.Filename))
	newpath = filepath.Join("assets", path, filename)
	if err = os.MkdirAll(filepath.Dir(newpath), os.ModeDir|0775); err != nil {
		return
	}

	dst, err := os.Create(newpath)
	if err != nil {
		return
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return
	}

	return
}

func UploadDokRwBKN(c echo.Context) error {
	//nip := c.Param("nip")
	db, _ := model.CreateCon()

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "file not found",
			"errors":     err.Error(),
		})
	}

	//var singkronisasi model.Singkronisasi
	//host := "bkn"

	//db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)

	url := "https://apimws.bkn.go.id:8243/apisiasn/1.0/upload-dok-rw"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer src.Close()

	part, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		fmt.Println(err)
	}

	_, err = io.Copy(part, src)
	if err != nil {
		fmt.Println(err)
		//return
	}

	//part.Write(fileContents)

	_ = writer.WriteField("id_riwayat", c.FormValue("idSync"))
	_ = writer.WriteField("id_ref_dokumen", c.FormValue("id_ref_dokumen"))

	writer.Close()

	/*out, err := json.Marshal(map[string]interface{}{
		"id_riwayat":     c.FormValue("idSync"),
		"id_ref_dokumen": c.FormValue("id_ref_dokumen"),
		"file":           c.FormFile("file"),
	})*/

	//fmt.Println(body)

	/*if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "A",
			"errors":     err.Error(),
		})
	}*/

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"pos":        "B",
			"errors":     err.Error(),
		})
	}

	req.Header.Add("accept", "application/json")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	a := getTokenApiManager(fmt.Sprint(c.Get("nip")))
	b := getTokenSSO(fmt.Sprint(c.Get("nip")))
	d := fmt.Sprintf("bearer %v", a.AccessToken)
	e := fmt.Sprintf("Bearer %v", b.AccessToken)

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

	//fmt.Println(req)

	/*bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)*/

	type ReturnBKNStatus struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
		Message     string `json:"message"`
		Success     bool   `json:"success"`
	}

	var retData ReturnBKNStatus

	err = json.NewDecoder(res.Body).Decode(&retData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statusCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	riwayat_update_bkn := &model.RiwayatUpdateBKN{
		Deskripsi: "Upload Dokumen PAK",
		CreatedBy: fmt.Sprint(c.Get("nip")),
		Code:      fmt.Sprintf("%v", retData.Code),
		Message:   retData.Message,
	}

	db.Model(&model.RiwayatUpdateBKN{}).Create(&riwayat_update_bkn)

	return c.JSON(http.StatusOK, retData)
}
