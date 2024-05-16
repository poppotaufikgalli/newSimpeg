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
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
	})
}

func UploadPangkat(c echo.Context) error {
	nip := c.Param("nip")
	tmtpang := c.FormValue("tmtpang")
	kgolru := c.FormValue("kgolru")
	knpang := c.FormValue("knpang")
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
		//result := db.Model(&model.Pegawai{}).Where("nip = ?", nip).Update("file_bmp", filename)
		//tmtpang = tmtpang.Format("2006-01-02")
		result := db.Model(&model.RiwayatPangkat{}).Where("nip = ? and tmtpang = ? and kgolru = ? and knpang = ?", nip, tmtpang, kgolru, knpang).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
	})
}

func UploadJabatan(c echo.Context) error {
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
		result := db.Model(&model.RiwayatJabatan{}).Where("nip = ? and tmtjab = ?", nip, tmtjab).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
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
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
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
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
	})
}

func UploadAK(c echo.Context) error {
	nip := c.Param("nip")
	tmulai := c.FormValue("tmulai")
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
		result := db.Model(&model.RiwayatAngkaKredit{}).Where("nip = ? and tmulai = ?", nip, tmulai).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
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
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
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
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
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
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
	})
}

func UploadP2kp(c echo.Context) error {
	nip := c.Param("nip")
	thnilai := c.FormValue("thnilai")
	tmulai := c.FormValue("tmulai")
	tselesai := c.FormValue("tselesai")

	filename, newpath, _ := doUpload(c)

	if nip != "" {
		db, _ := model.CreateCon()
		result := db.Model(&model.RiwayatP2kp{}).Where("nip = ? and thnilai = ? and tmulai =? and tselesai=?", nip, thnilai, tmulai, tselesai).Update("filename", filename)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
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
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
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
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
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
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
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
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
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
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
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
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
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
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
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
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
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
				"statucCode": http.StatusNotFound,
				"errors":     "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename":   filename,
		"path":       newpath,
		"statucCode": http.StatusOK,
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

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"pos":        "file not found",
			"errors":     err.Error(),
		})
	}

	db, _ := model.CreateCon()
	var singkronisasi model.Singkronisasi
	host := "bkn"

	db.Model(&model.Singkronisasi{Host: &host}).First(&singkronisasi)

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
			"statucCode": http.StatusBadRequest,
			"pos":        "A",
			"errors":     err.Error(),
		})
	}*/

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"pos":        "B",
			"errors":     err.Error(),
		})
	}

	req.Header.Add("accept", "application/json")
	req.Header.Set("Content-Type", writer.FormDataContentType())
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
			"statucCode": http.StatusBadRequest,
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
