package app

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	model "newSimpegAPI/model"
	"strconv"
	"strings"
)

func FindPegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var searchString model.SearchPegawai
	var pegawai []model.Pegawai

	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		/*return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"message":    err.Error(),
			"statucCode": http.StatusNotImplemented,
		})*/
		searchString.Kstatus = []int{1, 2}
	}

	var findJabatan []string
	findJabatan = append(findJabatan, "left join master_riwayat_jabatan on master_pegawai.nip = master_riwayat_jabatan.nip and master_riwayat_jabatan.akhir = 1")

	result := db.Model(&model.Pegawai{})

	//IdOpd
	if searchString.IdOpd != "" {
		findJabatan = append(findJabatan, fmt.Sprintf("and master_riwayat_jabatan.id_opd = '%s'", searchString.IdOpd))
	}

	//id_eselon
	if searchString.Keselon != 0 {
		findJabatan = append(findJabatan, fmt.Sprintf("and master_riwayat_jabatan.keselon = '%v'", searchString.Keselon))
	}

	//jnsjab
	if searchString.Jnsjab > 0 {
		//jnsjab := strings.Trim(strings.Replace(fmt.Sprint(searchString.Jnsjab), " ", ",", -1), "[]")
		//findJabatan = append(findJabatan, fmt.Sprintf("and master_riwayat_jabatan.jnsjab IN (%v)", jnsjab))
		findJabatan = append(findJabatan, fmt.Sprintf("and master_riwayat_jabatan.jnsjab = '%v'", searchString.Jnsjab))
	}

	appendJab := strings.Join(findJabatan, " ")

	result = result.Joins(appendJab).Preload("JabatanAkhir", "akhir = ? ", 1)

	//nama
	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		result = result.Where("master_pegawai.nama LIKE ?", strings.Join(str, ""))
	}

	//kstatus
	if len(searchString.Kstatus) > 0 {
		result = result.Where("master_pegawai.kstatus IN (?)", searchString.Kstatus)
	}

	//kjpeg
	if len(searchString.Kjpeg) > 0 {
		result = result.Where("master_pegawai.kjpeg IN (?)", searchString.Kjpeg)
	}

	//nip
	if searchString.Nip != "" {
		result = result.Where("master_pegawai.nip = ?", searchString.Nip)
	}

	//no_ref_bkn
	if searchString.NoRefBkn != "" {
		result = result.Where("master_pegawai.no_ref_bkn = ?", searchString.NoRefBkn)
	}

	result = result.Find(&pegawai)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pegawai,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
		"filter":     searchString,
	})
}

func GetPegawaiByNip(c echo.Context) error {
	db, _ := model.CreateCon()
	var pegawai model.Pegawai
	nip, _ := strconv.Atoi(c.Param("nip"))

	result := db.Model(&model.Pegawai{}).Where("nip = ?", nip).Scan(&pegawai)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pegawai,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func FindListPegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var searchString model.SearchPegawai
	var pegawai []model.PegawaiList

	err := json.NewDecoder(c.Request().Body).Decode(&searchString)

	if err != nil {
		searchString.Kstatus = []int{1, 2}
	}

	//var findJabatan []string
	age := "DATE_FORMAT(FROM_DAYS(DATEDIFF(CURRENT_DATE, tlahir)),'%y th %m bl')"
	sqlSelect := fmt.Sprintf("select a.nip, a.nama, a.gldepan, a.glblk, a.nik, a.npwp, a.ktlahir, a.tlahir, %s AS umur, a.aljalan, a.alrt, a.alrw, a.alkoprop, a.alkokab, a.alkokec, a.alkodes, a.kstatus, g.nama as nstatus, a.kjpeg, a.no_ref_bkn, b.id_opd, b.jnsjab, b.keselon, b.tmtjab, b.nskjabat, c.nama as neselon, b.njab, b.nunker, h.nama as nopd, d.kgolru, d.tmtpang, i.kgolru as kgolru_cpns, i.tmtcpns, j.kgolru as kgolru_pns, j.tmtpns, l.nama as ntpu, k.nama as njur, f.nsek, a.kjkel from master_pegawai a left join master_riwayat_jabatan b on (a.nip = b.nip and b.akhir =1) left join master_eselon c on (b.keselon=c.id) left join master_riwayat_pangkat d on (a.nip = d.nip and d.akhir = 1) left join master_riwayat_pendum f on (a.nip =f.nip and f.akhir = 1) left join master_status_pegawai g on (a.kstatus=g.kstatus) left join master_opd h on (substr(b.id_opd,1,3) = h.id) left join master_cpns i on (a.nip=i.nip) left join master_pns j on (a.nip=j.nip) left join master_pendidikan k on (f.kjur=k.id) left join master_tingkat_pendidikan l on (f.ktpu=l.id)", age)

	var searchGroup []string
	//nip
	if searchString.Nip != "" {
		//result = result.Where("master_pegawai.nip = ?", searchString.Nip)
		searchGroup = append(searchGroup, fmt.Sprintf("a.nip = %s", searchString.Nip))
	}

	//IdOpd
	if searchString.IdOpd != "" {
		IdOpd := strings.TrimSpace(searchString.IdOpd)
		str := []string{IdOpd, "%"}
		searchGroup = append(searchGroup, fmt.Sprintf("b.id_opd LIKE '%s'", strings.Join(str, "")))
		//findJabatan = append(findJabatan, fmt.Sprintf("and master_riwayat_jabatan.id_opd = '%s'", searchString.IdOpd))
	}

	//id_eselon
	if searchString.Keselon != 0 {
		searchGroup = append(searchGroup, fmt.Sprintf("b.keselon = '%v'", searchString.Keselon))
	}

	//jnsjab
	if searchString.Jnsjab > 0 {
		//jnsjab := strings.Trim(strings.Replace(fmt.Sprint(searchString.Jnsjab), " ", ",", -1), "[]")
		//searchGroup = append(searchGroup, fmt.Sprintf("b.jnsjab IN (%v)", jnsjab))
		searchGroup = append(searchGroup, fmt.Sprintf("b.jnsjab = '%v'", searchString.Jnsjab))
	}

	if searchString.Nama != "" {
		nama := strings.TrimSpace(searchString.Nama)
		str := []string{"%", nama, "%"}
		searchGroup = append(searchGroup, fmt.Sprintf("a.nama LIKE '%s'", strings.Join(str, "")))
		//result = result.Where("master_pegawai.nama LIKE ?", strings.Join(str, ""))
	}

	//kstatus
	if len(searchString.Kstatus) > 0 {
		kstatus := strings.Trim(strings.Replace(fmt.Sprint(searchString.Kstatus), " ", ",", -1), "[]")
		searchGroup = append(searchGroup, fmt.Sprintf("a.kstatus IN (%v)", kstatus))
		//result = result.Where("master_pegawai.kstatus IN (?)", searchString.Kstatus)
	}

	//kjpeg
	if len(searchString.Kjpeg) > 0 {
		kjpeg := strings.Trim(strings.Replace(fmt.Sprint(searchString.Kjpeg), " ", ",", -1), "[]")
		searchGroup = append(searchGroup, fmt.Sprintf("a.kjpeg IN (%v)", kjpeg))
		//result = result.Where("master_pegawai.kjpeg IN (?)", searchString.Kjpeg)
	}

	//no_ref_bkn
	if searchString.NoRefBkn != "" {
		searchGroup = append(searchGroup, fmt.Sprintf("a.no_ref_bkn = %v", searchString.NoRefBkn))
		//result = result.Where("master_pegawai.no_ref_bkn = ?", searchString.NoRefBkn)
	}

	search := strings.Join(searchGroup, " and ")
	sqlString := fmt.Sprintf("%s where %s", sqlSelect, search)

	result := db.Raw(sqlString).Scan(&pegawai)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pegawai,
		"sqlString":  sqlString,
		"search":     search,
		"statucCode": http.StatusOK,
		"filter":     searchString,
		"count":      result.RowsAffected,
	})
}

func CreatePegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var pegawai model.Pegawai
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pegawai)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	pegawai.CreatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Create(&pegawai)

	//fmt.Println(pegawai)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pegawai,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func UpdatePegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var pegawai model.Pegawai
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pegawai)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	pegawai.UpdatedBy = fmt.Sprint(c.Get("nip"))
	result := db.Model(&model.Pegawai{}).Where("nip = ?", pegawai.Nip).Updates(&pegawai)

	//fmt.Println(pegawai)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pegawai,
		"statucCode": http.StatusOK,
		"count":      result.RowsAffected,
	})
}

func DeletePegawai(c echo.Context) error {
	db, _ := model.CreateCon()

	var pegawai model.DeletePegawai
	validate := validator.New()

	err := json.NewDecoder(c.Request().Body).Decode(&pegawai)

	if err != nil {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"statucCode": http.StatusNotImplemented,
			"errors":     err.Error(),
		})
	}

	if err = validate.Struct(pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"statucCode": http.StatusBadRequest,
			"errors":     err.Error(),
		})
	}

	result := db.Model(&model.Pegawai{}).Where("nip = ?", pegawai.Nip).Delete(&pegawai)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statucCode": http.StatusNotFound,
			"errors":     "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       pegawai,
		"statucCode": http.StatusOK,
	})
}
