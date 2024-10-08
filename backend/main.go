package main

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lpernett/godotenv"
	"log"
	"os"
	"strings"

	app "newSimpegAPI/app"
)

var (
	APP_KEY string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	APP_KEY = os.Getenv("APP_KEY")
}

func main() {
	e := echo.New()
	e.Static("/static", "assets")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(app.JwtCustomClaims)
		},
		SigningKey: []byte(APP_KEY),
		Skipper: func(c echo.Context) bool {
			if strings.HasSuffix(c.Path(), "/login") {
				return true
			}

			if strings.HasSuffix(c.Path(), "/fakelogin") {
				return true
			}
			return false
		},
	}

	e.Use(echojwt.WithConfig(config))
	e.Use(app.AuthorizeRequest)

	e.POST("/login", app.Login)
	e.POST("/fakelogin", app.FakeLogin)
	e.GET("/logout", app.Logout)
	e.GET("/islogin", app.IsLogin)

	//upload
	e.POST("/upload/pegawai/:nip", app.UploadPegawai)
	e.POST("/upload/pangkat/:nip", app.UploadPangkat)
	e.POST("/upload/jabatan/:nip", app.UploadJabatan)
	e.POST("/upload/tugas_tambahan/:nip", app.UploadTugasTambahan)
	e.POST("/upload/kgb/:nip", app.UploadKGB)
	e.POST("/upload/ak/:nip", app.UploadAK)
	e.POST("/upload/stlud/:nip", app.UploadSTLUD)
	e.POST("/upload/pmk/:nip", app.UploadPMK)
	e.POST("/upload/pindah_instansi/:nip", app.UploadPI)
	e.POST("/upload/skp/:nip", app.UploadSKP)
	e.POST("/upload/kinerja/:nip", app.UploadKinerja)
	e.POST("/upload/penghargaan/:nip", app.UploadPenghargaan)
	e.POST("/upload/hukdis/:nip", app.UploadHukdis)
	e.POST("/upload/cuti/:nip", app.UploadCuti)
	e.POST("/upload/organisasi/:nip", app.UploadOrganisasi)
	e.POST("/upload/tugas_ln/:nip", app.UploadTugasLn)
	e.POST("/upload/bahasa/:nip", app.UploadBahasa)
	e.POST("/upload/pendum/:nip", app.UploadPendum)
	e.POST("/upload/tubel/:nip", app.UploadTubel)
	e.POST("/upload/diklat/:nip", app.UploadDiklat)
	e.POST("/upload/keluarga/:nip", app.UploadKeluarga)

	e.POST("/upload/dokRwBkn/:nip", app.UploadDokRwBKN)

	e.POST("/download/dokRwBkn", app.DownloadDokRwBKN)

	//pegawai
	routePegawai := e.Group("/pegawai")
	routePegawai.GET("", app.FindPegawai)
	routePegawai.POST("", app.FindPegawai)
	routePegawai.GET("/:nip", app.GetPegawaiByNip)
	routePegawai.GET("/list", app.FindListPegawai)
	routePegawai.POST("/list", app.FindListPegawai)
	routePegawai.POST("/new", app.CreatePegawai)
	routePegawai.PUT("", app.UpdatePegawai)
	routePegawai.DELETE("", app.DeletePegawai)

	//OPD
	routeOPD := e.Group("/opd")
	routeOPD.GET("", app.FindOPD)
	routeOPD.POST("", app.FindOPD)
	routeOPD.GET("/:id", app.GetOPDbyId)
	routeOPD.POST("/new", app.CreateOPD)
	routeOPD.GET("/filter", app.SearchOPDFilter)
	routeOPD.PUT("", app.UpdateOPD)
	routeOPD.DELETE("", app.DeleteOPD)

	//Eselon
	routeEselon := e.Group("/eselon")
	routeEselon.GET("", app.FindEselon)
	routeEselon.POST("/new", app.CreateEselon)
	routeEselon.PUT("", app.UpdateEselon)
	routeEselon.DELETE("", app.DeleteEselon)

	//Status Pegawai
	routeStatusPegawai := e.Group("/status_pegawai")
	routeStatusPegawai.GET("", app.FindStatusPegawai)
	routeStatusPegawai.POST("/new", app.CreateStatusPegawai)
	routeStatusPegawai.PUT("", app.UpdateStatusPegawai)
	routeStatusPegawai.DELETE("", app.DeleteStatusPegawai)

	//Jenis Jabatan
	routeJenisJabatan := e.Group("/jenis_jabatan")
	routeJenisJabatan.GET("", app.FindJenisJabatan)
	routeJenisJabatan.POST("/new", app.CreateJenisJabatan)
	routeJenisJabatan.PUT("", app.UpdateJenisJabatan)
	routeJenisJabatan.DELETE("", app.DeleteJenisJabatan)

	//Jenis Pegawai
	routeJenisPegawai := e.Group("/jenis_pegawai")
	routeJenisPegawai.GET("", app.FindJenisPegawai)
	routeJenisPegawai.POST("/new", app.CreateJenisPegawai)
	routeJenisPegawai.PUT("", app.UpdateJenisPegawai)
	routeJenisPegawai.DELETE("", app.DeleteJenisPegawai)

	//Jenis Tugas Mutasi
	routeJenisTugasMutasi := e.Group("/jenis_tugas_mutasi")
	routeJenisTugasMutasi.GET("", app.FindJenisTugasMutasi)
	routeJenisTugasMutasi.GET("/:id", app.SearchJenisTugasMutasiById)
	routeJenisTugasMutasi.POST("/new", app.CreateJenisTugasMutasi)
	routeJenisTugasMutasi.PUT("", app.UpdateJenisTugasMutasi)
	routeJenisTugasMutasi.DELETE("", app.DeleteJenisTugasMutasi)

	//riwayat jabatan
	routeRiwayatJabatan := e.Group("/riwayat_jabatan")
	//routeRiwayatJabatan.GET("", app.FindRiwayatJabatan)
	routeRiwayatJabatan.GET("/:nip", app.GetRiwayatJabatanByNip)
	routeRiwayatJabatan.GET("/:nip/akhir", app.GetRiwayatJabatanByNipAkhir)
	//routeRiwayatJabatan.GET("/:nip/:tmtjab", app.GetRiwayatJabatanByNipTmtjab)
	routeRiwayatJabatan.GET("/:nip/:id", app.GetRiwayatJabatanByNipId)
	//routeRiwayatJabatan.GET("/:nip/:tmtjab/:tmt_mutasi", app.GetRiwayatJabatanByNipTmtjabTmtMutasi)
	routeRiwayatJabatan.POST("", app.FindRiwayatJabatan)
	routeRiwayatJabatan.POST("/new", app.CreateRiwayatJabatan)
	routeRiwayatJabatan.POST("/bkn", app.UpdateRiwayatJabatanBkn)
	routeRiwayatJabatan.PUT("", app.UpdateRiwayatJabatan)
	routeRiwayatJabatan.DELETE("", app.DeleteRiwayatJabatan)

	//riwayat Pegawai
	routeRiwayatPangkat := e.Group("/riwayat_pangkat")
	routeRiwayatPangkat.GET("", app.FindRiwayatPangkat)
	routeRiwayatPangkat.GET("/:nip", app.GetRiwayatPangkatByNip)
	routeRiwayatPangkat.GET("/:nip/cpns", app.GetRiwayatPangkatByNipCPNS)
	routeRiwayatPangkat.GET("/:nip/pns", app.GetRiwayatPangkatByNipPNS)
	routeRiwayatPangkat.GET("/:nip/akhir", app.GetRiwayatPangkatByNipAkhir)
	routeRiwayatPangkat.GET("/:nip/:kgolru", app.GetRiwayatPangkatByNipByKgolru)
	routeRiwayatPangkat.GET("/:nip/:kgolru/:tmtpang/:knpang", app.GetRiwayatPangkatByNipByKgolruTmtpangKnpang)
	routeRiwayatPangkat.POST("", app.FindRiwayatPangkat)
	routeRiwayatPangkat.POST("/new", app.CreateRiwayatPangkat)
	routeRiwayatPangkat.PUT("", app.UpdateRiwayatPangkat)
	routeRiwayatPangkat.DELETE("", app.DeleteRiwayatPangkat)

	//Agama
	routeAgama := e.Group("/agama")
	routeAgama.GET("", app.FindAgama)
	routeAgama.POST("/new", app.CreateAgama)
	routeAgama.PUT("", app.UpdateAgama)
	routeAgama.DELETE("", app.DeleteAgama)

	//CPNS
	routeCpns := e.Group("/cpns")
	routeCpns.GET("", app.FindCpns)
	routeCpns.GET("/:nip", app.GetCpnsByNip)
	routeCpns.POST("/new", app.CreateCpns)
	routeCpns.PUT("", app.UpdateCpns)
	routeCpns.DELETE("", app.DeleteCpns)

	//Diklat
	routeDiklat := e.Group("/diklat")
	routeDiklat.GET("", app.FindDiklat)
	routeDiklat.POST("", app.FindDiklat)
	routeDiklat.POST("/new", app.CreateDiklat)
	routeDiklat.PUT("", app.UpdateDiklat)
	routeDiklat.DELETE("", app.DeleteDiklat)

	//Diklat Str
	routeDiklatStr := e.Group("/diklat_str")
	routeDiklatStr.GET("", app.FindDiklatStr)
	routeDiklatStr.POST("/new", app.CreateDiklatStr)
	routeDiklatStr.PUT("", app.UpdateDiklatStr)
	routeDiklatStr.DELETE("", app.DeleteDiklatStr)

	//riwayat jabatan
	routeFormasiJabatan := e.Group("/formasi_jabatan")
	routeFormasiJabatan.GET("", app.FindFormasiJabatan)
	routeFormasiJabatan.POST("", app.FindFormasiJabatan)
	routeFormasiJabatan.POST("/new", app.CreateFormasiJabatan)
	routeFormasiJabatan.PUT("", app.UpdateFormasiJabatan)
	routeFormasiJabatan.DELETE("", app.DeleteFormasiJabatan)

	//Group Arsip
	routeGroupArsip := e.Group("/group_arsip")
	routeGroupArsip.GET("", app.FindGroupArsip)
	routeGroupArsip.POST("/new", app.CreateGroupArsip)
	routeGroupArsip.PUT("", app.UpdateGroupArsip)
	routeGroupArsip.DELETE("", app.DeleteGroupArsip)

	//Group Arsip
	routeInstansi := e.Group("/instansi")
	routeInstansi.GET("", app.FindInstansi)
	routeInstansi.POST("", app.FindInstansi)
	routeInstansi.POST("/new", app.CreateInstansi)
	routeInstansi.PUT("", app.UpdateInstansi)
	routeInstansi.DELETE("", app.DeleteInstansi)

	//Jabatan FT
	routeJabatanFt := e.Group("/jabatan_ft")
	routeJabatanFt.GET("", app.FindJabatanFt)
	routeJabatanFt.POST("", app.FindJabatanFt)
	routeJabatanFt.POST("/new", app.CreateJabatanFt)
	routeJabatanFt.PUT("", app.UpdateJabatanFt)
	routeJabatanFt.DELETE("", app.DeleteJabatanFt)

	//Jabatan FT BKN
	routeJabatanFtBkn := e.Group("/jabatan_ft_bkn")
	routeJabatanFtBkn.GET("", app.FindJabatanFtBkn)
	routeJabatanFtBkn.GET("/filter", app.SearchJabatanFtBkn)
	routeJabatanFtBkn.POST("/new", app.CreateJabatanFtBkn)
	routeJabatanFtBkn.PUT("", app.UpdateJabatanFtBkn)
	routeJabatanFtBkn.DELETE("", app.DeleteJabatanFtBkn)

	//Jabatan SUB BKN
	routeJabatanSubBkn := e.Group("/jabatan_sub_bkn")
	routeJabatanSubBkn.GET("", app.FindJabatanSubBkn)
	routeJabatanSubBkn.POST("", app.FindJabatanSubBkn)
	routeJabatanSubBkn.GET("/filter_jft/:id_jft", app.FindJabatanSubBknFilterJft)
	routeJabatanSubBkn.POST("/new", app.CreateJabatanSubBkn)
	routeJabatanSubBkn.PUT("", app.UpdateJabatanSubBkn)
	routeJabatanSubBkn.DELETE("", app.DeleteJabatanSubBkn)

	//Jabatan FU
	routeJabatanFu := e.Group("/jabatan_fu")
	routeJabatanFu.GET("", app.FindJabatanFu)
	routeJabatanFu.POST("", app.FindJabatanFu)
	routeJabatanFu.POST("/new", app.CreateJabatanFu)
	routeJabatanFu.PUT("", app.UpdateJabatanFu)
	routeJabatanFu.DELETE("", app.DeleteJabatanFu)

	//Jabatan FU BKN
	routeJabatanFuBkn := e.Group("/jabatan_fu_bkn")
	routeJabatanFuBkn.GET("", app.FindJabatanFuBkn)
	routeJabatanFuBkn.POST("", app.FindJabatanFuBkn)
	routeJabatanFuBkn.GET("/filter", app.SearchJabatanFuBkn)
	routeJabatanFuBkn.POST("/new", app.CreateJabatanFuBkn)
	routeJabatanFuBkn.PUT("", app.UpdateJabatanFuBkn)
	routeJabatanFuBkn.DELETE("", app.DeleteJabatanFuBkn)

	//Jabatan Negara
	routeJabatanNegara := e.Group("/jabatan_negara")
	routeJabatanNegara.GET("", app.FindJabatanNegara)
	routeJabatanNegara.POST("", app.FindJabatanNegara)
	routeJabatanNegara.POST("/new", app.CreateJabatanNegara)
	routeJabatanNegara.PUT("", app.UpdateJabatanNegara)
	routeJabatanNegara.DELETE("", app.DeleteJabatanNegara)

	//Jabatan Str
	routeJabatanStr := e.Group("/jabatan_str")
	routeJabatanStr.GET("", app.FindJabatanStr)
	routeJabatanStr.POST("", app.FindJabatanStr)
	routeJabatanStr.POST("/new", app.CreateJabatanStr)
	routeJabatanStr.PUT("", app.UpdateJabatanStr)
	routeJabatanStr.DELETE("", app.DeleteJabatanStr)

	//Jenis Cuti
	routeJenisCuti := e.Group("/jenis_cuti")
	routeJenisCuti.GET("", app.FindJenisCuti)
	routeJenisCuti.POST("/new", app.CreateJenisCuti)
	routeJenisCuti.PUT("", app.UpdateJenisCuti)
	routeJenisCuti.DELETE("", app.DeleteJenisCuti)

	//Jenis Diklat
	routeJenisDiklat := e.Group("/jenis_diklat")
	routeJenisDiklat.GET("", app.FindJenisDiklat)
	routeJenisDiklat.POST("", app.FindJenisDiklat)
	routeJenisDiklat.GET("/:id", app.GetJenisDiklatByJdiklat)
	routeJenisDiklat.POST("/new", app.CreateJenisDiklat)
	routeJenisDiklat.PUT("", app.UpdateJenisDiklat)
	routeJenisDiklat.DELETE("", app.DeleteJenisDiklat)

	//Jenis Goldar
	routeJenisGoldar := e.Group("/jenis_goldar")
	routeJenisGoldar.GET("", app.FindJenisGoldar)
	routeJenisGoldar.POST("/new", app.CreateJenisGoldar)
	routeJenisGoldar.PUT("", app.UpdateJenisGoldar)
	routeJenisGoldar.DELETE("", app.DeleteJenisGoldar)

	//Jenis Hukdis
	routeJenisHukdis := e.Group("/jenis_hukdis")
	routeJenisHukdis.GET("", app.FindJenisHukdis)
	routeJenisHukdis.POST("/new", app.CreateJenisHukdis)
	routeJenisHukdis.PUT("", app.UpdateJenisHukdis)
	routeJenisHukdis.DELETE("", app.DeleteJenisHukdis)

	//Jenis Instansi
	routeJenisInstansi := e.Group("/jenis_instansi")
	routeJenisInstansi.GET("", app.FindJenisInstansi)
	routeJenisInstansi.POST("/new", app.CreateJenisInstansi)
	routeJenisInstansi.PUT("", app.UpdateJenisInstansi)
	routeJenisInstansi.DELETE("", app.DeleteJenisInstansi)

	//Jenis Jabatan Umum
	routeJenisJabatanUmum := e.Group("/jenis_jabatan_umum")
	routeJenisJabatanUmum.GET("", app.FindJenisJabatanUmum)
	routeJenisJabatanUmum.POST("/new", app.CreateJenisJabatanUmum)
	routeJenisJabatanUmum.PUT("", app.UpdateJenisJabatanUmum)
	routeJenisJabatanUmum.DELETE("", app.DeleteJenisJabatanUmum)

	//Jenis Kawin
	routeJenisKawin := e.Group("/jenis_kawin")
	routeJenisKawin.GET("", app.FindJenisKawin)
	routeJenisKawin.POST("/new", app.CreateJenisKawin)
	routeJenisKawin.PUT("", app.UpdateJenisKawin)
	routeJenisKawin.DELETE("", app.DeleteJenisKawin)

	//Jenis Keluarga
	routeJenisKeluarga := e.Group("/jenis_keluarga")
	routeJenisKeluarga.GET("", app.FindJenisKeluarga)
	routeJenisKeluarga.GET("/:id", app.GetJenisKeluargaByJkeluarga)
	routeJenisKeluarga.POST("/new", app.CreateJenisKeluarga)
	routeJenisKeluarga.PUT("", app.UpdateJenisKeluarga)
	routeJenisKeluarga.DELETE("", app.DeleteJenisKeluarga)

	//Jenis Kompetensi
	routeJenisKompetensi := e.Group("/jenis_kompetensi")
	routeJenisKompetensi.GET("", app.FindJenisKompetensi)
	routeJenisKompetensi.POST("/new", app.CreateJenisKompetensi)
	routeJenisKompetensi.PUT("", app.UpdateJenisKompetensi)
	routeJenisKompetensi.DELETE("", app.DeleteJenisKompetensi)

	//Jenis Kp
	routeJenisKp := e.Group("/jenis_kp")
	routeJenisKp.GET("", app.FindJenisKp)
	routeJenisKp.POST("/new", app.CreateJenisKp)
	routeJenisKp.PUT("", app.UpdateJenisKp)
	routeJenisKp.DELETE("", app.DeleteJenisKp)

	//Jenis Kursus
	routeJenisKursus := e.Group("/jenis_kursus")
	routeJenisKursus.GET("", app.FindJenisKursus)
	routeJenisKursus.POST("/new", app.CreateJenisKursus)
	routeJenisKursus.PUT("", app.UpdateJenisKursus)
	routeJenisKursus.DELETE("", app.DeleteJenisKursus)

	//Jenis Organisasi
	routeJenisOrganisasi := e.Group("/jenis_organisasi")
	routeJenisOrganisasi.GET("", app.FindJenisOrganisasi)
	routeJenisOrganisasi.POST("/new", app.CreateJenisOrganisasi)
	routeJenisOrganisasi.PUT("", app.UpdateJenisOrganisasi)
	routeJenisOrganisasi.DELETE("", app.DeleteJenisOrganisasi)

	//Jenis Pemberhentian
	routeJenisPemberhentian := e.Group("/jenis_pemberhentian")
	routeJenisPemberhentian.GET("", app.FindJenisPemberhentian)
	routeJenisPemberhentian.POST("/new", app.CreateJenisPemberhentian)
	routeJenisPemberhentian.PUT("", app.UpdateJenisPemberhentian)
	routeJenisPemberhentian.DELETE("", app.DeleteJenisPemberhentian)

	//Jenis Pengadaaan
	routeJenisPengadaan := e.Group("/jenis_pengadaan")
	routeJenisPengadaan.GET("", app.FindJenisPengadaan)
	routeJenisPengadaan.POST("/new", app.CreateJenisPengadaan)
	routeJenisPengadaan.PUT("", app.UpdateJenisPengadaan)
	routeJenisPengadaan.DELETE("", app.DeleteJenisPengadaan)

	//Jenis Penghargaan
	routeJenisPenghargaan := e.Group("/jenis_penghargaan")
	routeJenisPenghargaan.GET("", app.FindJenisPenghargaan)
	routeJenisPenghargaan.POST("/new", app.CreateJenisPenghargaan)
	routeJenisPenghargaan.PUT("", app.UpdateJenisPenghargaan)
	routeJenisPenghargaan.DELETE("", app.DeleteJenisPenghargaan)

	//Jenis Pensiun
	routeJenisPensiun := e.Group("/jenis_pensiun")
	routeJenisPensiun.GET("", app.FindJenisPensiun)
	routeJenisPensiun.POST("/new", app.CreateJenisPensiun)
	routeJenisPensiun.PUT("", app.UpdateJenisPensiun)
	routeJenisPensiun.DELETE("", app.DeleteJenisPensiun)

	//jenjang Jabfung
	routeJenjangJabfung := e.Group("/jenjang_jabfung")
	routeJenjangJabfung.GET("", app.FindJenjangJabfung)
	routeJenjangJabfung.POST("/new", app.CreateJenjangJabfung)
	routeJenjangJabfung.PUT("", app.UpdateJenjangJabfung)
	routeJenjangJabfung.DELETE("", app.DeleteJenjangJabfung)

	//kedudukan pegawai
	routeKedudukanPegawai := e.Group("/kedudukan_pegawai")
	routeKedudukanPegawai.GET("", app.FindKedudukanPegawai)
	routeKedudukanPegawai.POST("/new", app.CreateKedudukanPegawai)
	routeKedudukanPegawai.PUT("", app.UpdateKedudukanPegawai)
	routeKedudukanPegawai.DELETE("", app.DeleteKedudukanPegawai)

	//kelompok jabatan
	routeKelompokJabatan := e.Group("/kelompok_jabatan")
	routeKelompokJabatan.GET("", app.FindKelompokJabatan)
	routeKelompokJabatan.POST("/new", app.CreateKelompokJabatan)
	routeKelompokJabatan.PUT("", app.UpdateKelompokJabatan)
	routeKelompokJabatan.DELETE("", app.DeleteKelompokJabatan)

	//kpe
	routeKpe := e.Group("/kpe")
	routeKpe.GET("", app.FindKpe)
	routeKpe.POST("/new", app.CreateKpe)
	routeKpe.PUT("", app.UpdateKpe)
	routeKpe.DELETE("", app.DeleteKpe)

	//kppn
	routeKppn := e.Group("/kppn")
	routeKppn.GET("", app.FindKppn)
	routeKppn.POST("/new", app.CreateKppn)
	routeKppn.PUT("", app.UpdateKppn)
	routeKppn.DELETE("", app.DeleteKppn)

	//pangkat
	routePangkat := e.Group("/pangkat")
	routePangkat.GET("", app.FindPangkat)
	routePangkat.POST("/new", app.CreatePangkat)
	routePangkat.PUT("", app.UpdatePangkat)
	routePangkat.DELETE("", app.DeletePangkat)

	//pejabat
	routePejabat := e.Group("/pejabat")
	routePejabat.GET("", app.FindPejabat)
	routePejabat.POST("/new", app.CreatePejabat)
	routePejabat.PUT("", app.UpdatePejabat)
	routePejabat.DELETE("", app.DeletePejabat)

	//pekerjaan
	routePekerjaan := e.Group("/pekerjaan")
	routePekerjaan.GET("", app.FindPekerjaan)
	routePekerjaan.POST("/new", app.CreatePekerjaan)
	routePekerjaan.PUT("", app.UpdatePekerjaan)
	routePekerjaan.DELETE("", app.DeletePekerjaan)

	//pendidikan
	routePendidikan := e.Group("/pendidikan")
	routePendidikan.GET("", app.FindPendidikan)
	routePendidikan.GET("/filter", app.SearchPendidikanFilter)
	routePendidikan.GET("/filter/:ktpu", app.SearchPendidikanFilterKtpu)
	routePendidikan.POST("", app.FindPendidikan)
	routePendidikan.POST("/new", app.CreatePendidikan)
	routePendidikan.PUT("", app.UpdatePendidikan)
	routePendidikan.DELETE("", app.DeletePendidikan)

	//pendidikan_bkn
	routePendidikanBkn := e.Group("/pendidikan_bkn")
	routePendidikanBkn.GET("", app.FindPendidikanBkn)
	routePendidikanBkn.GET("/filter", app.SearchPendidikanBknFilter)
	routePendidikanBkn.POST("", app.FindPendidikanBkn)
	routePendidikanBkn.POST("/new", app.CreatePendidikanBkn)
	routePendidikanBkn.PUT("", app.UpdatePendidikanBkn)
	routePendidikanBkn.DELETE("", app.DeletePendidikanBkn)

	//profesi
	routeProfesi := e.Group("/profesi")
	routeProfesi.GET("", app.FindProfesi)
	routeProfesi.POST("/new", app.CreateProfesi)
	routeProfesi.PUT("", app.UpdateProfesi)
	routeProfesi.DELETE("", app.DeleteProfesi)

	//referensi jabatan
	routeReferensiJabatan := e.Group("/referensi_jabatan")
	routeReferensiJabatan.GET("", app.FindReferensiJabatan)
	routeReferensiJabatan.POST("/new", app.CreateReferensiJabatan)
	routeReferensiJabatan.PUT("", app.UpdateReferensiJabatan)
	routeReferensiJabatan.DELETE("", app.DeleteReferensiJabatan)

	//referensi nilai kuadran
	routeReferensiNilaiKuadran := e.Group("/referensi_nilai_kuadran")
	routeReferensiNilaiKuadran.GET("", app.FindReferensiNilaiKuadran)
	routeReferensiNilaiKuadran.POST("/new", app.CreateReferensiNilaiKuadran)
	routeReferensiNilaiKuadran.PUT("", app.UpdateReferensiNilaiKuadran)
	routeReferensiNilaiKuadran.DELETE("", app.DeleteReferensiNilaiKuadran)

	//riwayat angka kredit
	routeRiwayatAngkaKredit := e.Group("/riwayat_angka_kredit")
	routeRiwayatAngkaKredit.GET("", app.FindRiwayatAngkaKredit)
	routeRiwayatAngkaKredit.POST("", app.FindRiwayatAngkaKredit)
	routeRiwayatAngkaKredit.GET("/:nip/:tmulai", app.GetRiwayatAngkaKreditByNipTmulai)
	routeRiwayatAngkaKredit.POST("/new", app.CreateRiwayatAngkaKredit)
	routeRiwayatAngkaKredit.POST("/bkn", app.UpdateRiwayatAngkaKreditBkn)
	routeRiwayatAngkaKredit.PUT("", app.UpdateRiwayatAngkaKredit)
	routeRiwayatAngkaKredit.DELETE("", app.DeleteRiwayatAngkaKredit)

	//riwayat bahasa
	routeRiwayatBahasa := e.Group("/riwayat_bahasa")
	routeRiwayatBahasa.GET("", app.FindRiwayatBahasa)
	routeRiwayatBahasa.POST("", app.FindRiwayatBahasa)
	routeRiwayatBahasa.GET("/:nip/:nbahasa", app.GetRiwayatBahasaByNipNbahasa)
	routeRiwayatBahasa.POST("/new", app.CreateRiwayatBahasa)
	routeRiwayatBahasa.PUT("", app.UpdateRiwayatBahasa)
	routeRiwayatBahasa.DELETE("", app.DeleteRiwayatBahasa)

	//riwayat cuti
	routeRiwayatCuti := e.Group("/riwayat_cuti")
	routeRiwayatCuti.GET("", app.FindRiwayatCuti)
	routeRiwayatCuti.POST("", app.FindRiwayatCuti)
	routeRiwayatCuti.GET("/:nip/:nsk/:tmulai", app.GetRiwayatCutiByNipNskTmulai)
	routeRiwayatCuti.POST("/new", app.CreateRiwayatCuti)
	routeRiwayatCuti.PUT("", app.UpdateRiwayatCuti)
	routeRiwayatCuti.DELETE("", app.DeleteRiwayatCuti)

	//riwayat diklat
	routeRiwayatDiklat := e.Group("/riwayat_diklat")
	routeRiwayatDiklat.GET("", app.FindRiwayatDiklat)
	routeRiwayatDiklat.POST("", app.FindRiwayatDiklat)
	routeRiwayatDiklat.GET("/:nip/:jdiklat/:tmulai", app.GetRiwayatDiklatByNipJdiklatTmulai)
	routeRiwayatDiklat.POST("/new", app.CreateRiwayatDiklat)
	routeRiwayatDiklat.PUT("", app.UpdateRiwayatDiklat)
	routeRiwayatDiklat.DELETE("", app.DeleteRiwayatDiklat)

	//riwayat gajiberkala
	routeRiwayatGajiberkala := e.Group("/riwayat_gajiberkala")
	routeRiwayatGajiberkala.GET("", app.FindRiwayatGajiberkala)
	routeRiwayatGajiberkala.GET("/:nip", app.GetRiwayatGajiBerkalaByNip)
	routeRiwayatGajiberkala.GET("/:nip/akhir", app.GetRiwayatGajiBerkalaByNipAkhir)
	routeRiwayatGajiberkala.GET("/:nip/:tmtngaj", app.GetRiwayatGajiBerkalaByNipTmtngaj)
	routeRiwayatGajiberkala.POST("", app.FindRiwayatGajiberkala)
	routeRiwayatGajiberkala.POST("/new", app.CreateRiwayatGajiberkala)
	routeRiwayatGajiberkala.PUT("", app.UpdateRiwayatGajiberkala)
	routeRiwayatGajiberkala.DELETE("", app.DeleteRiwayatGajiberkala)

	//riwayat hukdis
	routeRiwayatHukdis := e.Group("/riwayat_hukdis")
	routeRiwayatHukdis.GET("", app.FindRiwayatHukdis)
	routeRiwayatHukdis.POST("", app.FindRiwayatHukdis)
	routeRiwayatHukdis.GET("/:nip/:jhukum/:tmt", app.GetRiwayatHukdisByNipJhukumTmt)
	routeRiwayatHukdis.POST("/new", app.CreateRiwayatHukdis)
	routeRiwayatHukdis.PUT("", app.UpdateRiwayatHukdis)
	routeRiwayatHukdis.DELETE("", app.DeleteRiwayatHukdis)

	//riwayat Keluarga
	routeRiwayatKeluarga := e.Group("/riwayat_keluarga")
	routeRiwayatKeluarga.GET("", app.FindRiwayatKeluarga)
	routeRiwayatKeluarga.POST("", app.FindRiwayatKeluarga)
	routeRiwayatKeluarga.GET("/:nip/:jkeluarga/:nama_kel", app.GetRiwayatKeluargaByNipJkeluargaNamakel)
	routeRiwayatKeluarga.POST("/new", app.CreateRiwayatKeluarga)
	routeRiwayatKeluarga.PUT("", app.UpdateRiwayatKeluarga)
	routeRiwayatKeluarga.DELETE("", app.DeleteRiwayatKeluarga)

	//riwayat organisasi
	routeRiwayatOrganisasi := e.Group("/riwayat_organisasi")
	routeRiwayatOrganisasi.GET("", app.FindRiwayatOrganisasi)
	routeRiwayatOrganisasi.POST("", app.FindRiwayatOrganisasi)
	routeRiwayatOrganisasi.GET("/:nip/:norg/:jborg", app.GetRiwayatOrganisasiByNipNorgJborg)
	routeRiwayatOrganisasi.POST("/new", app.CreateRiwayatOrganisasi)
	routeRiwayatOrganisasi.PUT("", app.UpdateRiwayatOrganisasi)
	routeRiwayatOrganisasi.DELETE("", app.DeleteRiwayatOrganisasi)

	//riwayat SKP
	routeRiwayatSKP := e.Group("/riwayat_skp")
	routeRiwayatSKP.GET("", app.FindRiwayatSkp)
	routeRiwayatSKP.GET("/:nip/:id/:tahun", app.GetRiwayatSkpByNipIdTahun)
	routeRiwayatSKP.POST("", app.FindRiwayatSkp)
	routeRiwayatSKP.POST("/new", app.CreateRiwayatSkp)
	routeRiwayatSKP.POST("/bkn", app.UpdateRiwayatSkpBkn)
	routeRiwayatSKP.PUT("", app.UpdateRiwayatSkp)
	routeRiwayatSKP.DELETE("", app.DeleteRiwayatSkp)

	//riwayat Kinerja
	routeRiwayatKinerja := e.Group("/riwayat_kinerja")
	routeRiwayatKinerja.GET("", app.FindRiwayatKinerja)
	routeRiwayatKinerja.GET("/:nip/:id/:thnilai", app.GetRiwayatKinerjaByNipIdThn)
	routeRiwayatKinerja.POST("", app.FindRiwayatKinerja)
	routeRiwayatKinerja.POST("/new", app.CreateRiwayatKinerja)
	routeRiwayatKinerja.POST("/bkn", app.UpdateRiwayatKinerjaBkn)
	routeRiwayatKinerja.PUT("", app.UpdateRiwayatKinerja)
	routeRiwayatKinerja.DELETE("", app.DeleteRiwayatKinerja)

	//riwayat pendum
	routeRiwayatPendum := e.Group("/riwayat_pendum")
	routeRiwayatPendum.GET("", app.FindRiwayatPendum)
	routeRiwayatPendum.POST("", app.FindRiwayatPendum)
	routeRiwayatPendum.GET("/:nip/:ktpu/:kjur", app.GetRiwayatPendumByNipKtpuKjur)
	routeRiwayatPendum.POST("/new", app.CreateRiwayatPendum)
	routeRiwayatPendum.PUT("", app.UpdateRiwayatPendum)
	routeRiwayatPendum.DELETE("", app.DeleteRiwayatPendum)

	//riwayat pendum
	routeRiwayatTubel := e.Group("/riwayat_tubel")
	routeRiwayatTubel.GET("", app.FindRiwayatTubel)
	routeRiwayatTubel.POST("", app.FindRiwayatTubel)
	routeRiwayatTubel.GET("/:nip/:id", app.GetRiwayatTubelByNipId)
	routeRiwayatTubel.POST("/new", app.CreateRiwayatTubel)
	routeRiwayatTubel.PUT("", app.UpdateRiwayatTubel)
	routeRiwayatTubel.DELETE("", app.DeleteRiwayatTubel)

	//riwayat penghargaan
	routeRiwayatPenghargaan := e.Group("/riwayat_penghargaan")
	routeRiwayatPenghargaan.GET("", app.FindRiwayatPenghargaan)
	routeRiwayatPenghargaan.POST("", app.FindRiwayatPenghargaan)
	routeRiwayatPenghargaan.GET("/:nip/:nbintang/:nsk", app.GetRiwayatPenghargaanByNipIdNsk)
	routeRiwayatPenghargaan.POST("/new", app.CreateRiwayatPenghargaan)
	routeRiwayatPenghargaan.PUT("", app.UpdateRiwayatPenghargaan)
	routeRiwayatPenghargaan.DELETE("", app.DeleteRiwayatPenghargaan)

	//riwayat pindah instansi
	routeRiwayatPI := e.Group("/riwayat_pindah_instansi")
	routeRiwayatPI.GET("", app.FindRiwayatPindahInstansi)
	routeRiwayatPI.GET("/:nip/:tmtPi", app.GetRiwayatPindahInstansiByNipTmtPi)
	routeRiwayatPI.POST("", app.FindRiwayatPindahInstansi)
	routeRiwayatPI.POST("/new", app.CreateRiwayatPindahInstansi)
	routeRiwayatPI.PUT("", app.UpdateRiwayatPindahInstansi)
	routeRiwayatPI.DELETE("", app.DeleteRiwayatPindahInstansi)

	//riwayat PMK
	routeRiwayatPMK := e.Group("/riwayat_pmk")
	routeRiwayatPMK.GET("", app.FindRiwayatPmk)
	routeRiwayatPMK.POST("", app.FindRiwayatPmk)
	routeRiwayatPMK.GET("/:nip/:tanggalAwal", app.GetRiwayatPmkByNiptanggalAwal)
	routeRiwayatPMK.POST("/new", app.CreateRiwayatPmk)
	routeRiwayatPMK.PUT("", app.UpdateRiwayatPmk)
	routeRiwayatPMK.DELETE("", app.DeleteRiwayatPmk)

	//riwayat singkronisasi
	routeRiwayatSingkronisasi := e.Group("/riwayat_singkronisasi")
	routeRiwayatSingkronisasi.GET("", app.FindRiwayatSingkronisasi)
	routeRiwayatSingkronisasi.POST("/new", app.CreateRiwayatSingkronisasi)
	routeRiwayatSingkronisasi.PUT("", app.UpdateRiwayatSingkronisasi)
	routeRiwayatSingkronisasi.DELETE("", app.DeleteRiwayatSingkronisasi)

	//riwayat Stlud
	routeRiwayatStlud := e.Group("/riwayat_stlud")
	routeRiwayatStlud.GET("", app.FindRiwayatStlud)
	routeRiwayatStlud.GET("/:nip/:kstlud/:tstlud", app.GetRiwayatStludByNipKstludTstlud)
	routeRiwayatStlud.POST("", app.FindRiwayatStlud)
	routeRiwayatStlud.POST("/new", app.CreateRiwayatStlud)
	routeRiwayatStlud.PUT("", app.UpdateRiwayatStlud)
	routeRiwayatStlud.DELETE("", app.DeleteRiwayatStlud)

	//riwayat Tugas Ln
	routeRiwayatTugasLn := e.Group("/riwayat_tugas_ln")
	routeRiwayatTugasLn.GET("", app.FindRiwayatTugasLn)
	routeRiwayatTugasLn.POST("", app.FindRiwayatTugasLn)
	routeRiwayatTugasLn.GET("/:nip/:tmulai", app.GetRiwayatTugasLnByNipTmulai)
	routeRiwayatTugasLn.POST("/new", app.CreateRiwayatTugasLn)
	routeRiwayatTugasLn.PUT("", app.UpdateRiwayatTugasLn)
	routeRiwayatTugasLn.DELETE("", app.DeleteRiwayatTugasLn)

	//riwayat Tugas tambahan
	routeRiwayatTugasTambahan := e.Group("/riwayat_tugas_tambahan")
	routeRiwayatTugasTambahan.GET("", app.FindRiwayatTugasTambahan)
	routeRiwayatTugasTambahan.GET("/:nip", app.GetRiwayatTugasTambahanByNip)
	routeRiwayatTugasTambahan.GET("/:nip/akhir", app.GetRiwayatTugasTambahanByNipAkhir)
	routeRiwayatTugasTambahan.GET("/:nip/:tmtjab", app.GetRiwayatTugasTambahanByNipTmtjab)
	routeRiwayatTugasTambahan.POST("", app.FindRiwayatTugasTambahan)
	routeRiwayatTugasTambahan.POST("/new", app.CreateRiwayatTugasTambahan)
	routeRiwayatTugasTambahan.PUT("", app.UpdateRiwayatTugasTambahan)
	routeRiwayatTugasTambahan.DELETE("", app.DeleteRiwayatTugasTambahan)

	//riwayat Tumlar
	routeRiwayatTumlar := e.Group("/riwayat_tumlar")
	routeRiwayatTumlar.GET("", app.FindRiwayatTumlar)
	routeRiwayatTumlar.POST("/new", app.CreateRiwayatTumlar)
	routeRiwayatTumlar.PUT("", app.UpdateRiwayatTumlar)
	routeRiwayatTumlar.DELETE("", app.DeleteRiwayatTumlar)

	//rumpun jabfung
	routeRumpunJabfung := e.Group("/rumpun_jabfung")
	routeRumpunJabfung.GET("", app.FindRumpunJabfung)
	routeRumpunJabfung.POST("/new", app.CreateRumpunJabfung)
	routeRumpunJabfung.PUT("", app.UpdateRumpunJabfung)
	routeRumpunJabfung.DELETE("", app.DeleteRumpunJabfung)

	//rumpun jabatan
	routeRumpunJabatan := e.Group("/rumpun_jabatan")
	routeRumpunJabatan.GET("", app.FindRumpunJabatan)
	routeRumpunJabatan.POST("/new", app.CreateRumpunJabatan)
	routeRumpunJabatan.PUT("", app.UpdateRumpunJabatan)
	routeRumpunJabatan.DELETE("", app.DeleteRumpunJabatan)

	//Singkronisasi
	routeSingkronisasi := e.Group("/singkronisasi")
	routeSingkronisasi.GET("", app.FindSingkronisasi)
	routeSingkronisasi.GET("/:host", app.GetSingkronisasiByHost)
	routeSingkronisasi.GET("/:host/expiredTime", app.GetSingkronisasiExpiredTimeByHost)
	routeSingkronisasi.POST("", app.FindSingkronisasi)
	routeSingkronisasi.POST("/new", app.CreateSingkronisasi)
	routeSingkronisasi.PUT("", app.UpdateSingkronisasi)
	routeSingkronisasi.DELETE("", app.DeleteSingkronisasi)

	routeSingkronisasiBKN := e.Group("/singkronisasi_bkn")
	routeSingkronisasiBKN.POST("/riwayat", app.GetSingkronisasiRiwayat)
	routeSingkronisasiBKN.POST("/updateToken", app.UpdateTokenSingkronisasi)
	routeSingkronisasiBKN.GET("/getDataBKN/:page/:nip", app.GetSingkronisasiGetDataBkn)
	routeSingkronisasiBKN.GET("/getDataBKNById/:page/:idSync", app.GetSingkronisasiGetDataBknById)

	//routeSingkronisasiBKN.DELETE("/delDataBKNById/:page/:idSync", app.DeleteSingkronisasiDataBknById)

	routeSingkronisasiBKN.PUT("/putsDataBKN/:page", app.UpdateSingkronisasiPutDataUtamaBkn)
	routeSingkronisasiBKN.POST("/getDocFromBkn", app.PostSinkronisasiGetDocFromBKN)
	routeSingkronisasiBKN.PUT("/putsDataBKNAK/:nip", app.UpdateSingkronisasiPutDataBknAngkaKredit)
	routeSingkronisasiBKN.PUT("/putsDataBKNCpnsPns/:nip", app.UpdateSingkronisasiPutDataBknCpnsPns)
	routeSingkronisasiBKN.PUT("/putsUnorJabatan/:nip", app.UpdateSingkronisasiPutUnorJabatan)
	routeSingkronisasiBKN.DELETE("/delDataBKNAK/:id", app.DeleteSingkronisasiDelDataBknAngkaKredit)
	routeSingkronisasiBKN.DELETE("/delRiwJabatan/:id", app.DeleteSingkronisasiDelRiwJabatan)

	//Stlud
	routeStlud := e.Group("/stlud")
	routeStlud.GET("", app.FindStlud)
	routeStlud.POST("/new", app.CreateStlud)
	routeStlud.PUT("", app.UpdateStlud)
	routeStlud.DELETE("", app.DeleteStlud)

	//Tingkat Pendidikan
	routeTingkatPendidikan := e.Group("/tingkat_pendidikan")
	routeTingkatPendidikan.GET("", app.FindTingkatPendidikan)
	routeTingkatPendidikan.POST("/new", app.CreateTingkatPendidikan)
	routeTingkatPendidikan.PUT("", app.UpdateTingkatPendidikan)
	routeTingkatPendidikan.DELETE("", app.DeleteTingkatPendidikan)

	//Tugas Tambahan
	routeTugasTambahan := e.Group("/tugas_tambahan")
	routeTugasTambahan.GET("", app.FindTugasTambahan)
	routeTugasTambahan.POST("/new", app.CreateTugasTambahan)
	routeTugasTambahan.PUT("", app.UpdateTugasTambahan)
	routeTugasTambahan.DELETE("", app.DeleteTugasTambahan)

	//Unor BKN
	routeUnorBkn := e.Group("/unor_bkn")
	routeUnorBkn.GET("", app.FindUnorBkn)
	routeUnorBkn.POST("/new", app.CreateUnorBkn)
	routeUnorBkn.PUT("", app.UpdateUnorBkn)
	routeUnorBkn.DELETE("", app.DeleteUnorBkn)

	//Wilayah
	routeWilayah := e.Group("/wilayah")
	routeWilayah.GET("", app.FindWilayah)
	routeWilayah.POST("", app.FindWilayah)
	routeWilayah.POST("/new", app.CreateWilayah)
	routeWilayah.PUT("", app.UpdateWilayah)
	routeWilayah.DELETE("", app.DeleteWilayah)

	//ref jenpeg
	routeRefJenpeg := e.Group("/ref_jenpeg")
	routeRefJenpeg.GET("", app.FindRefJenpeg)
	routeRefJenpeg.POST("/new", app.CreateRefJenpeg)
	routeRefJenpeg.PUT("", app.UpdateRefJenpeg)
	routeRefJenpeg.DELETE("", app.DeleteRefJenpeg)

	//users
	routeUsers := e.Group("/users")
	routeUsers.GET("", app.FindUsers)
	routeUsers.GET("/:id", app.GetUsersById)
	routeUsers.POST("/new", app.CreateUsers)
	routeUsers.PUT("", app.UpdateUsers)
	routeUsers.DELETE("", app.DeleteUsers)

	//group
	routeGroup := e.Group("/group")
	routeGroup.GET("", app.FindGroup)
	routeGroup.GET("/:id", app.GetGroupById)
	routeGroup.POST("/new", app.CreateGroup)
	routeGroup.PUT("", app.UpdateGroup)
	routeGroup.DELETE("", app.DeleteGroup)

	e.Logger.Fatal(e.Start(":1323"))
}
