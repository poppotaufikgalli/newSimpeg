<script setup>
const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const toast = useToast()
const {snip} = route.params
const showSpinner = ref(false)
import customParseFormat from "dayjs/plugin/customParseFormat";
dayjs.extend(customParseFormat)

onMounted(() => {
	//refreshNuxtData(["getData"])
})

const result = ref([{
    "eselon": null,
    "eselonId": "",
    "id": "c13059bb-d24d-4749-b1f2-397fe75758bd",
    "idPns": "A8ACA7CD3C7F3912E040640A040269BB",
    "instansiKerjaId": "A5EB03E23CFBF6A0E040640A040252AD",
    "instansiKerjaNama": "Pemerintah Kota Tanjungpinang",
    "jabatanFungsionalId": "A5EB03E23EB0F6A0E040640A040252AD",
    "jabatanFungsionalNama": "Pranata Komputer Ahli Muda",
    "jabatanFungsionalUmumId": "",
    "jabatanFungsionalUmumNama": "",
    "jenisJabatan": "2",
    "jenisMutasiId": "MJ",
    "jenisPenugasanId": "",
    "namaJabatan": "Pranata Komputer Ahli Muda",
    "namaUnor": "",
    "nipBaru": "198602162008031001",
    "nipLama": "P20006592",
    "nomorSk": "17 TAHUN 2024",
    "path":
    {
        "872":
        {
            "dok_id": "872",
            "dok_nama": "Dok SK Jabatan",
            "dok_uri": "peremajaan/usulan/872_c13059bb-d24d-4749-b1f2-397fe75758bd.pdf",
            "object": "peremajaan/usulan/872_c13059bb-d24d-4749-b1f2-397fe75758bd.pdf",
            "slug": "872"
        }
    },
    "satuanKerjaId": "A5EB03E24326F6A0E040640A040252AD",
    "satuanKerjaNama": "Pemerintah Kota Tanjung Pinang",
    "subJabatanId": "0",
    "tanggalSk": "03-04-2024",
    "tmtJabatan": "03-04-2024",
    "tmtMutasi": "01-01-0001",
    "tmtPelantikan": "01-01-0001",
    "unorId": "8ae4828859b32c060159b5817cc82c8c",
    "unorIndukId": "8ae48287599445da015996893743048f",
    "unorIndukNama": "DINAS KOMUNIKASI DAN INFORMATIKA",
    "unorNama": "BIDANG TEKNOLOGI INFORMASI DAN KOMUNIKASI (TIK)"
},
{
    "eselon": "NON",
    "eselonId": "99",
    "id": "7bd7f18b-5b75-11ee-a19d-0a580a830006",
    "idPns": "A8ACA7CD3C7F3912E040640A040269BB",
    "instansiKerjaId": "A5EB03E23CFBF6A0E040640A040252AD",
    "instansiKerjaNama": "Pemerintah Kota Tanjungpinang",
    "jabatanFungsionalId": "",
    "jabatanFungsionalNama": "",
    "jabatanFungsionalUmumId": "8ae48288503bd9ee01505f6b026e0f3b",
    "jabatanFungsionalUmumNama": "FUNGSIONAL UMUM",
    "jenisJabatan": "4",
    "jenisMutasiId": "",
    "jenisPenugasanId": "",
    "namaJabatan": "FUNGSIONAL UMUM",
    "namaUnor": "DINAS KOMUNIKASI DAN INFORMATIKA",
    "nipBaru": "198602162008031001",
    "nipLama": "P20006592",
    "nomorSk": "121212",
    "path":
    {
        "872":
        {
            "dok_id": "872",
            "dok_nama": "Dok SK Jabatan",
            "dok_uri": "peremajaan/usulan/A8ACA7CD79C73912E040640A040269BB_20230925_073314_198602162008031001_0000_0000_06_09_2023.pdf",
            "object": "peremajaan/usulan/A8ACA7CD79C73912E040640A040269BB_20230925_073314_198602162008031001_0000_0000_06_09_2023.pdf",
            "slug": "872"
        },
        "":
        {
            "dok_id": "",
            "dok_nama": "",
            "dok_uri": "",
            "object": "",
            "slug": ""
        }
    },
    "satuanKerjaId": "A5EB03E24326F6A0E040640A040252AD",
    "satuanKerjaNama": "Pemerintah Kota Tanjung Pinang",
    "subJabatanId": "",
    "tanggalSk": "07-09-2023",
    "tmtJabatan": "06-09-2023",
    "tmtMutasi": "",
    "tmtPelantikan": "01-01-0001",
    "unorId": "8ae48287599445da015996893743048f",
    "unorIndukId": "A8ACA7436B153912E040640A040269BB",
    "unorIndukNama": "PEMERINTAH KOTA TANJUNGPINANG",
    "unorNama": "DINAS KOMUNIKASI DAN INFORMATIKA"
},
{
    "eselon": null,
    "eselonId": "",
    "id": "8ae483a67e63e2c3017e6b8114c823c4",
    "idPns": "A8ACA7CD3C7F3912E040640A040269BB",
    "instansiKerjaId": "A5EB03E23CFBF6A0E040640A040252AD",
    "instansiKerjaNama": "Pemerintah Kota Tanjungpinang",
    "jabatanFungsionalId": "A5EB03E23EB0F6A0E040640A040252AD",
    "jabatanFungsionalNama": "Pranata Komputer Ahli Muda",
    "jabatanFungsionalUmumId": "",
    "jabatanFungsionalUmumNama": "",
    "jenisJabatan": "2",
    "jenisMutasiId": "",
    "jenisPenugasanId": "",
    "namaJabatan": "",
    "namaUnor": "BIDANG TEKNOLOGI INFORMASI KOMUNIKASI (TIK)",
    "nipBaru": "198602162008031001",
    "nipLama": "P20006592",
    "nomorSk": "780 TAHUN 2021",
    "path": null,
    "satuanKerjaId": "A5EB03E24326F6A0E040640A040252AD",
    "satuanKerjaNama": "Pemerintah Kota Tanjung Pinang",
    "subJabatanId": "",
    "tanggalSk": "31-12-2021",
    "tmtJabatan": "31-12-2021",
    "tmtMutasi": "",
    "tmtPelantikan": "01-01-0001",
    "unorId": "8ae4828859b32c060159b5817cc82c8c",
    "unorIndukId": "8ae48287599445da015996893743048f",
    "unorIndukNama": "DINAS KOMUNIKASI DAN INFORMATIKA",
    "unorNama": "BIDANG TEKNOLOGI INFORMASI DAN KOMUNIKASI (TIK)"
},
{
    "eselon": "IV.a",
    "eselonId": "41",
    "id": "8ae483a57882ef57017888e551ba5747",
    "idPns": "A8ACA7CD3C7F3912E040640A040269BB",
    "instansiKerjaId": "A5EB03E23CFBF6A0E040640A040252AD",
    "instansiKerjaNama": "Pemerintah Kota Tanjungpinang",
    "jabatanFungsionalId": "",
    "jabatanFungsionalNama": "",
    "jabatanFungsionalUmumId": null,
    "jabatanFungsionalUmumNama": null,
    "jenisJabatan": "1",
    "jenisMutasiId": "",
    "jenisPenugasanId": "",
    "namaJabatan": "KEPALA SEKSI e-GOVERNMENT DAN KEAMANAN SISTEM INFORMASI",
    "namaUnor": "SEKSI e-GOVERNMENT DAN KEAMANAN SISTEM INFORMASI",
    "nipBaru": "198602162008031001",
    "nipLama": "P20006592",
    "nomorSk": "35 TAHUN 2021",
    "path":
    {
        "872":
        {
            "dok_id": "872",
            "dok_nama": "Dok SK Jabatan",
            "dok_uri": "peremajaan/usulan/A8ACA7CD79C73912E040640A040269BB_20230925_051227_198602162008031001_0_0_19_01_2021.pdf",
            "object": "peremajaan/usulan/A8ACA7CD79C73912E040640A040269BB_20230925_051227_198602162008031001_0_0_19_01_2021.pdf",
            "slug": "872"
        }
    },
    "satuanKerjaId": "A5EB03E24326F6A0E040640A040252AD",
    "satuanKerjaNama": "Pemerintah Kota Tanjung Pinang",
    "subJabatanId": "",
    "tanggalSk": "19-01-2021",
    "tmtJabatan": "19-01-2021",
    "tmtMutasi": "",
    "tmtPelantikan": "19-01-2021",
    "unorId": "8ae4828859b32c060159b58ca1ed44fa",
    "unorIndukId": "8ae48287599445da015996893743048f",
    "unorIndukNama": "DINAS KOMUNIKASI DAN INFORMATIKA",
    "unorNama": "SEKSI e-GOVERNMENT DAN KEAMANAN SISTEM INFORMASI"
},
{
    "eselon": "IV.a",
    "eselonId": "41",
    "id": "8ae482855cc5065c015cc8f75cf624b8",
    "idPns": "A8ACA7CD3C7F3912E040640A040269BB",
    "instansiKerjaId": "A5EB03E23CFBF6A0E040640A040252AD",
    "instansiKerjaNama": "Pemerintah Kota Tanjungpinang",
    "jabatanFungsionalId": "",
    "jabatanFungsionalNama": "",
    "jabatanFungsionalUmumId": null,
    "jabatanFungsionalUmumNama": null,
    "jenisJabatan": "1",
    "jenisMutasiId": "",
    "jenisPenugasanId": "",
    "namaJabatan": "KEPALA SUB BIDANG PENGOLAHAN DATA DAN INFORMASI",
    "namaUnor": "SUB BIDANG PENGOLAHAN DATA DAN INFORMASI",
    "nipBaru": "198602162008031001",
    "nipLama": "P20006592",
    "nomorSk": "44 TAHUN 2017",
    "path": null,
    "satuanKerjaId": "A5EB03E24326F6A0E040640A040252AD",
    "satuanKerjaNama": "Pemerintah Kota Tanjung Pinang",
    "subJabatanId": "",
    "tanggalSk": "23-01-2017",
    "tmtJabatan": "24-01-2017",
    "tmtMutasi": "",
    "tmtPelantikan": "24-01-2017",
    "unorId": "8ae4828559aa45020159aad53d4e043c",
    "unorIndukId": "8ae48287599445da015996732cdd73fc",
    "unorIndukNama": "BADAN PENGELOLAAN PAJAK DAN RETRIBUSI DAERAH",
    "unorNama": "SUB BIDANG PENGOLAHAN DATA DAN INFORMASI"
},
{
    "eselon": "NON",
    "eselonId": "99",
    "id": "7a0d1aef-5b05-11ee-a19d-0a580a830006",
    "idPns": "A8ACA7CD3C7F3912E040640A040269BB",
    "instansiKerjaId": "A5EB03E23CFBF6A0E040640A040252AD",
    "instansiKerjaNama": "Pemerintah Kota Tanjungpinang",
    "jabatanFungsionalId": "",
    "jabatanFungsionalNama": "",
    "jabatanFungsionalUmumId": "8ae48288503bd9ee01505f6b026e0f3b",
    "jabatanFungsionalUmumNama": "FUNGSIONAL UMUM",
    "jenisJabatan": "4",
    "jenisMutasiId": "",
    "jenisPenugasanId": "",
    "namaJabatan": "FUNGSIONAL UMUM",
    "namaUnor": "BADAN KEPEGAWAIAN DAN PENGEMBANGAN SUMBER DAYA MANUSIA",
    "nipBaru": "198602162008031001",
    "nipLama": "P20006592",
    "nomorSk": "5 TAHUN 2017",
    "path":
    {
        "":
        {
            "dok_id": "",
            "dok_nama": "",
            "dok_uri": "",
            "object": "",
            "slug": ""
        }
    },
    "satuanKerjaId": "A5EB03E24326F6A0E040640A040252AD",
    "satuanKerjaNama": "Pemerintah Kota Tanjung Pinang",
    "subJabatanId": "",
    "tanggalSk": "13-01-2017",
    "tmtJabatan": "13-01-2017",
    "tmtMutasi": "",
    "tmtPelantikan": "01-01-0001",
    "unorId": "8ae48287599445da01599671ec307355",
    "unorIndukId": "A8ACA7436B153912E040640A040269BB",
    "unorIndukNama": "PEMERINTAH KOTA TANJUNGPINANG",
    "unorNama": "BADAN KEPEGAWAIAN DAN PENGEMBANGAN SUMBER DAYA MANUSIA"
},
{
    "eselon": null,
    "eselonId": null,
    "id": "8ae482874f81c7b9014f820bc34c1812",
    "idPns": "A8ACA7CD3C7F3912E040640A040269BB",
    "instansiKerjaId": "A5EB03E23CFBF6A0E040640A040252AD",
    "instansiKerjaNama": "Pemerintah Kota Tanjungpinang",
    "jabatanFungsionalId": "",
    "jabatanFungsionalNama": "",
    "jabatanFungsionalUmumId": "ff80808135a3a1670135a3c3f119059d",
    "jabatanFungsionalUmumNama": "FUNGSIONAL UMUM",
    "jenisJabatan": "4",
    "jenisMutasiId": "",
    "jenisPenugasanId": "",
    "namaJabatan": null,
    "namaUnor": "BADAN KEPEGAWAIAN, PENDIDIKAN DAN PELATIHAN DAERAH",
    "nipBaru": "198602162008031001",
    "nipLama": "P20006592",
    "nomorSk": "112 TAHUN 2016",
    "path": null,
    "satuanKerjaId": "A5EB03E24326F6A0E040640A040252AD",
    "satuanKerjaNama": "Pemerintah Kota Tanjung Pinang",
    "subJabatanId": "",
    "tanggalSk": "31-03-2016",
    "tmtJabatan": "01-04-2016",
    "tmtMutasi": "",
    "tmtPelantikan": "",
    "unorId": "ff80808131ae95a70131aee8b4520178",
    "unorIndukId": "8ae483a66cace127016cd10be24240c2",
    "unorIndukNama": "SOTK LAMA",
    "unorNama": "BADAN KEPEGAWAIAN, PENDIDIKAN DAN PELATIHAN DAERAH"
},
{
    "eselon": null,
    "eselonId": null,
    "id": "8ae482a44f9984ba014fa11596822071",
    "idPns": "A8ACA7CD3C7F3912E040640A040269BB",
    "instansiKerjaId": "A5EB03E23CFBF6A0E040640A040252AD",
    "instansiKerjaNama": "Pemerintah Kota Tanjungpinang",
    "jabatanFungsionalId": "",
    "jabatanFungsionalNama": "",
    "jabatanFungsionalUmumId": "ff80808135a3a1670135a3c3f119059d",
    "jabatanFungsionalUmumNama": "FUNGSIONAL UMUM",
    "jenisJabatan": "4",
    "jenisMutasiId": "",
    "jenisPenugasanId": "",
    "namaJabatan": null,
    "namaUnor": "BADAN KEPEGAWAIAN, PENDIDIKAN DAN PELATIHAN DAERAH",
    "nipBaru": "198602162008031001",
    "nipLama": "P20006592",
    "nomorSk": "56 TAHUN 2015",
    "path": null,
    "satuanKerjaId": "A5EB03E24326F6A0E040640A040252AD",
    "satuanKerjaNama": "Pemerintah Kota Tanjung Pinang",
    "subJabatanId": "",
    "tanggalSk": "28-09-2015",
    "tmtJabatan": "28-09-2015",
    "tmtMutasi": "",
    "tmtPelantikan": "",
    "unorId": "ff80808131ae95a70131aee8b4520178",
    "unorIndukId": "8ae483a66cace127016cd10be24240c2",
    "unorIndukNama": "SOTK LAMA",
    "unorNama": "BADAN KEPEGAWAIAN, PENDIDIKAN DAN PELATIHAN DAERAH"
},
{
    "eselon": "NON",
    "eselonId": "99",
    "id": "dab6169c-5b0b-11ee-a19d-0a580a830006",
    "idPns": "A8ACA7CD3C7F3912E040640A040269BB",
    "instansiKerjaId": "A5EB03E23CFBF6A0E040640A040252AD",
    "instansiKerjaNama": "Pemerintah Kota Tanjungpinang",
    "jabatanFungsionalId": "",
    "jabatanFungsionalNama": "",
    "jabatanFungsionalUmumId": "8ae48288503bd9ee01505f6b026e0f3b",
    "jabatanFungsionalUmumNama": "FUNGSIONAL UMUM",
    "jenisJabatan": "4",
    "jenisMutasiId": "",
    "jenisPenugasanId": "",
    "namaJabatan": "FUNGSIONAL UMUM",
    "namaUnor": "BADAN KEPEGAWAIAN DAN PENGEMBANGAN SUMBER DAYA MANUSIA",
    "nipBaru": "198602162008031001",
    "nipLama": "P20006592",
    "nomorSk": "184/SP/2015",
    "path":
    {
        "":
        {
            "dok_id": "",
            "dok_nama": "",
            "dok_uri": "",
            "object": "",
            "slug": ""
        }
    },
    "satuanKerjaId": "A5EB03E24326F6A0E040640A040252AD",
    "satuanKerjaNama": "Pemerintah Kota Tanjung Pinang",
    "subJabatanId": "",
    "tanggalSk": "24-08-2015",
    "tmtJabatan": "24-08-2015",
    "tmtMutasi": "",
    "tmtPelantikan": "01-01-0001",
    "unorId": "8ae48287599445da01599671ec307355",
    "unorIndukId": "A8ACA7436B153912E040640A040269BB",
    "unorIndukNama": "PEMERINTAH KOTA TANJUNGPINANG",
    "unorNama": "BADAN KEPEGAWAIAN DAN PENGEMBANGAN SUMBER DAYA MANUSIA"
},
{
    "eselon": null,
    "eselonId": null,
    "id": "8ae48206504a5bbe0150566b077f70e6",
    "idPns": "A8ACA7CD3C7F3912E040640A040269BB",
    "instansiKerjaId": "A5EB03E23CFBF6A0E040640A040252AD",
    "instansiKerjaNama": "Pemerintah Kota Tanjungpinang",
    "jabatanFungsionalId": "",
    "jabatanFungsionalNama": "",
    "jabatanFungsionalUmumId": "ff80808135a3a1670135a3c3f119059d",
    "jabatanFungsionalUmumNama": "FUNGSIONAL UMUM",
    "jenisJabatan": "4",
    "jenisMutasiId": "",
    "jenisPenugasanId": "",
    "namaJabatan": null,
    "namaUnor": "PEMERINTAH KOTA TANJUNGPINANG",
    "nipBaru": "198602162008031001",
    "nipLama": "P20006592",
    "nomorSk": "SK.824.3/110/VII/2015",
    "path": null,
    "satuanKerjaId": "A5EB03E24326F6A0E040640A040252AD",
    "satuanKerjaNama": "Pemerintah Kota Tanjung Pinang",
    "subJabatanId": "",
    "tanggalSk": "24-07-2015",
    "tmtJabatan": "01-08-2015",
    "tmtMutasi": "",
    "tmtPelantikan": "",
    "unorId": "A8ACA7436B153912E040640A040269BB",
    "unorIndukId": "",
    "unorIndukNama": "",
    "unorNama": "PEMERINTAH KOTA TANJUNGPINANG"
},
{
    "eselon": null,
    "eselonId": null,
    "id": "8ae48206504a5bbe0150566822166d87",
    "idPns": "A8ACA7CD3C7F3912E040640A040269BB",
    "instansiKerjaId": "A5EB03E23D51F6A0E040640A040252AD",
    "instansiKerjaNama": "Pemerintah Kab. Bintan",
    "jabatanFungsionalId": "",
    "jabatanFungsionalNama": "",
    "jabatanFungsionalUmumId": "8ae482893e8315e6013e9bc3760f0af7",
    "jabatanFungsionalUmumNama": "PELAKSANA",
    "jenisJabatan": "4",
    "jenisMutasiId": "",
    "jenisPenugasanId": "",
    "namaJabatan": null,
    "namaUnor": "BADAN KEPEGAWAIAN DAERAH KAB. BINTAN",
    "nipBaru": "198602162008031001",
    "nipLama": "P20006592",
    "nomorSk": "SK.813.2 - 16 TAHUN 2008",
    "path": null,
    "satuanKerjaId": "A5EB03E24320F6A0E040640A040252AD",
    "satuanKerjaNama": "Pemerintah Kab. Bintan",
    "subJabatanId": "",
    "tanggalSk": "26-03-2008",
    "tmtJabatan": "01-03-2008",
    "tmtMutasi": "",
    "tmtPelantikan": "",
    "unorId": "A8ACA7435B513912E040640A040269BB",
    "unorIndukId": "",
    "unorIndukNama": "",
    "unorNama": "BADAN KEPEGAWAIAN DAERAH "
}])

const result_simpeg = ref([])
const result_bkn = ref([])

async function getDataBKN() {
	console.log("CariData")
	let nip = $decodeBase64(snip)
	var body = JSON.stringify({
		nip: nip,
	});

	result_simpeg.value = await $fetch('/api/posts/riwayat_jabatan', {
		method: 'POST',
		body: body,
	});

	//var bkn = await $fetch('/api/gets/singkronisasi_bkn/getDataBKN/rw-jabatan/'+nip)
	//result_bkn.value = bkn.data

    result_bkn.value = result.value

    if(result_simpeg.value){
            for (var i = 0; i < result_simpeg.value.length; i++) {
                var data_bkn = _find(result_bkn.value, {
                    tmtJabatan: (dayjs(dayjs(result_simpeg.value[i].tmtjab, "YYYY-MM-DD")).format("DD-MM-YYYY")).toString(),
                })
                //console.log(data_bkn)
                result_simpeg.value[i]['data_bkn'] = JSON.stringify(data_bkn)
            }
        }

        if(result_bkn.value){
            for (var i = 0; i < result_bkn.value.length; i++) {
                var data_simpeg = _find(result_simpeg.value, {
                    tmtjab: dayjs(dayjs(result_bkn.value[i].tmtJabatan, 'DD-MM-YYYY')).format("YYYY-MM-DD[T]HH:mm:ss+07:00"),
                })
                //console.log(data_simpeg)
                result_bkn.value[i]['data_simpeg'] = JSON.stringify(data_simpeg)
            }
        }
}

/*const editData = (item) => {
	let stmtjab = item ? $encodeBase64(item.tmtjab) : $encodeBase64(null)
	navigateTo({ path: `/pegawai/${snip}/RiwayatPegawai/jabatan/${stmtjab}` })
}

const hapusData = async(item) => {
	var body = JSON.stringify({
		nip: item.nip,
		tmtjab: item.tmtjab,
	});

	var tmtjab = dayjs(item.tmtjab).format("DD-MM-YYYY").toString()

	var {data, error} = await useFetch('/api/delete/riwayat_jabatan', {
		method: 'DELETE',
		body: body,
	});

	if(error.value){
		toast.add({
	    	id: 'error_delete_riwayat_jabatan',
	    	description: error.value.data.data,
	    	timeout: 6000
	  	}) 	
	}else{
		toast.add({
	    	id: 'success_delete_riwayat_jabatan',
	    	description: `Data jabatan ${item.njab} / TMT ${tmtjab} berhasil di Hapus`,
	    	timeout: 6000
	  	}) 	
	}

	//refreshNuxtData(["getData"])
}*/

function chekJenisJab(item){
    if(item.jenisJabatan == "1") return item.namaJabatan + " PADA " + item.unorNama + " " + item.unorIndukNama
    else if(item.jenisJabatan == "2") return item.jabatanFungsionalNama + " PADA " + item.unorNama + " " + item.unorIndukNama
    else if(item.jenisJabatan == "4") return item.jabatanFungsionalUmumNama + " PADA " + item.unorNama + " " + item.unorIndukNama
    else return item.jabatanFungsionalUmumNama + " PADA " + item.unorNama + " " + item.unorIndukNama
}

const refDataBkn = ref({
    "eselonId": "keselon",
    "id": "idSync",
    "instansiKerjaId": "id_instansi",
    "instansiKerjaNama": "nama_instansi",
    //"jabatanFungsionalId": "",
    //"jabatanFungsionalNama": "",
    //"jabatanFungsionalUmumId": "8ae482893e8315e6013e9bc3760f0af7",
    //"jabatanFungsionalUmumNama": "PELAKSANA",
    "jenisJabatan": "jnsjab",
    //"jenisMutasiId": "",
    //"jenisPenugasanId": "",
    //"namaJabatan": null,
    //"namaUnor": "BADAN KEPEGAWAIAN DAERAH KAB. BINTAN",
    "nipBaru": "nip",
    //"nipLama": "P20006592",
    "nomorSk": "nskjabat",
    //"path": null,
    //"satuanKerjaId": "A5EB03E24320F6A0E040640A040252AD",
    //"satuanKerjaNama": "Pemerintah Kab. Bintan",
    "subJabatanId": "",
    "tanggalSk": "tskjabat",
    "tmtJabatan": "tmtjab",
    //"tmtMutasi": "",
    //"tmtPelantikan": "",
    "unorId": "id_opd",
    //"unorIndukId": "",
    //"unorIndukNama": "",
    //"unorNama": "BADAN KEPEGAWAIAN DAERAH "
})

async function simpanData(item){
    //showSpinner.value = true
    let nip = $decodeBase64(snip)
    var nMethod = item.data_simpeg ? 'Update' : 'Simpan'; 

    //= _pickBy(item);
    var aitem = _pickBy(item, function(val, key) {
        return _includes(_keys(refDataBkn.value), key);
    })

    var compacted = _mapKeys(aitem, function(val, key){
        return _find(refDataBkn.value, function(val1, key1){
            return key == key1
        })
    })

    console.log(compacted)

    /*compacted.tsk = compacted.tsk ? dayjs(dayjs(compacted.tsk, "DD-MM-YYYY")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
    compacted.tmulai = dayjs(item.tahunMulaiPenailan+"-"+item.bulanMulaiPenailan+"-1").format("YYYY-MM-DD[T]HH:mm:ss[Z]")
    compacted.tselesai = dayjs(item.tahunSelesaiPenailan+"-"+item.bulanSelesaiPenailan+"-1").format("YYYY-MM-DD[T]HH:mm:ss[Z]")

    compacted.utama = parseFloat(compacted.utama)
    compacted.tambahan = parseFloat(compacted.tambahan)
    compacted.total = parseFloat(compacted.total)
    compacted.nip = nip
    compacted.jns = chekJenisAK(item)
    
    var dr = {
        nip: nip,
        tmulai: compacted.tmulai,
    }

    if(nMethod == 'Update'){
        var body = {
            "refdata" : dr,
            "data" : compacted
        }

        var {data, error} = await useFetch('/api/puts/riwayat_angka_kredit', {
            method: 'PUT',
            body: JSON.stringify(body),
        })

        if(error.value){
            toast.add({
                id: 'error_put_ak',
                description: error.value.data.data,
                timeout: 6000
            })  
        }else{
            toast.add({
                id: 'success_put_ak',
                description: "Data berhasil di Update",
                timeout: 6000
            })  
        }
    }else{
        compacted.nip = $decodeBase64(snip)
        //compacted.akhir = 1
        var body = JSON.stringify(compacted)

        var {data, error} = await useFetch('/api/posts/riwayat_angka_kredit/new', {
            method: 'POST',
            body: body,
        })

        if(error.value){
            toast.add({
                id: 'error_post_jabatan',
                description: error.value.data.data,
                timeout: 6000
            })  
        }else{
            toast.add({
                id: 'success_post_jabatan',
                description: "Data berhasil di Ditambahkan",
                timeout: 6000
            })  
        }
    }
    showSpinner.value = false
    getDataBKN()*/
}
</script>
<template>
	<AppLoadingSpinner :show="showSpinner" />
	<LayoutSingkronisasi>
		<TokenInfo />
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="flex justify-end mb-2">
					<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="getDataBKN()">Get Data BKN</button>
				</div>
				<div class="mb-8">
					<h2 class="font-bold text-teal-600">Data Simpeg</h2>
					<div class="overflow-x-auto border">
						<div class="min-w-full inline-block align-middle">
							<!-- Table -->
							<table class="min-w-full divide-y divide-gray-200">
								<thead class="bg-gray-50 h-12">
									<tr>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												No
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Nama
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												TMT Jabatan
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Nomor SK
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Tanggal
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Eselon
											</span>
										</th>
										<th scope="col" class="px-3 text-center"></th>
									</tr>
								</thead>
								<tbody class="divide-y divide-gray-200">
									<template v-if="result_simpeg" v-for="(item, idx) in result_simpeg" :key="idx">
										<tr class="odd:bg-white even:bg-gray-100">
											<td class="size-px px-1" style="vertical-align: top;">
												<button type="button" class="block" data-hs-overlay="#hs-ai-invoice-modal">
													<span class="flex justify-center text-xs px-3 py-0">{{idx+1}}</span>
												</button>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 200px;">
												<span class="flex justify-start text-xs py-0">{{item.njab}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 80px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.tmtjab ? $dayjs(item.tmtjab).format("DD-MM-YYYY").toString() : '-'}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 80px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.nskjabat}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 80px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.tskjabat ? $dayjs(item.tskjabat).format("DD-MM-YYYY").toString() : '-'}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 80px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.eselon?.nama}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top;">
												<!--<div class="sm:flex gap-x-1">
													<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-green-600 text-white hover:bg-green-700 disabled:opacity-50 disabled:pointer-events-none">
													  	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="flex-shrink-0 size-3">
															<path fill-rule="evenodd" d="M4.5 13a3.5 3.5 0 0 1-1.41-6.705A3.5 3.5 0 0 1 9.72 4.124a2.5 2.5 0 0 1 3.197 3.018A3.001 3.001 0 0 1 12 13H4.5Zm.72-5.03a.75.75 0 0 0 1.06 1.06l.97-.97v2.69a.75.75 0 0 0 1.5 0V8.06l.97.97a.75.75 0 1 0 1.06-1.06L8.53 5.72a.75.75 0 0 0-1.06 0L5.22 7.97Z" clip-rule="evenodd" />
														</svg>
													</button>
													<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="editData(item)">
													  	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="flex-shrink-0 size-3">
														  	<path d="M13.488 2.513a1.75 1.75 0 0 0-2.475 0L6.75 6.774a2.75 2.75 0 0 0-.596.892l-.848 2.047a.75.75 0 0 0 .98.98l2.047-.848a2.75 2.75 0 0 0 .892-.596l4.261-4.262a1.75 1.75 0 0 0 0-2.474Z" />
														  	<path d="M4.75 3.5c-.69 0-1.25.56-1.25 1.25v6.5c0 .69.56 1.25 1.25 1.25h6.5c.69 0 1.25-.56 1.25-1.25V9A.75.75 0 0 1 14 9v2.25A2.75 2.75 0 0 1 11.25 14h-6.5A2.75 2.75 0 0 1 2 11.25v-6.5A2.75 2.75 0 0 1 4.75 2H7a.75.75 0 0 1 0 1.5H4.75Z" />
														</svg>

													</button>
													<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-red-600 text-white hover:bg-red-700 disabled:opacity-50 disabled:pointer-events-none" @click="hapusData(item)">
													  	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="flex-shrink-0 size-3">
														  	<path fill-rule="evenodd" d="M5 3.25V4H2.75a.75.75 0 0 0 0 1.5h.3l.815 8.15A1.5 1.5 0 0 0 5.357 15h5.285a1.5 1.5 0 0 0 1.493-1.35l.815-8.15h.3a.75.75 0 0 0 0-1.5H11v-.75A2.25 2.25 0 0 0 8.75 1h-1.5A2.25 2.25 0 0 0 5 3.25Zm2.25-.75a.75.75 0 0 0-.75.75V4h3v-.75a.75.75 0 0 0-.75-.75h-1.5ZM6.05 6a.75.75 0 0 1 .787.713l.275 5.5a.75.75 0 0 1-1.498.075l-.275-5.5A.75.75 0 0 1 6.05 6Zm3.9 0a.75.75 0 0 1 .712.787l-.275 5.5a.75.75 0 0 1-1.498-.075l.275-5.5a.75.75 0 0 1 .786-.711Z" clip-rule="evenodd" />
														</svg>
													</button>
												</div>-->
											</td>
										</tr>
									</template>
								</tbody>
							</table>
							<!-- End Table -->	
						</div>
					</div>
				</div>
				<div class="mb-8">
					<h2 class="font-bold text-teal-600">Data BKN</h2>
					<div class="overflow-x-auto border">
						<div class="min-w-full inline-block align-middle">
							<!-- Table -->
							<table class="min-w-full divide-y divide-gray-200">
								<thead class="bg-gray-50 h-12">
									<tr>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												No
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Nama
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												TMT Jabatan
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Nomor SK
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Tanggal
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Eselon
											</span>
										</th>
										<th scope="col" class="px-3 text-center"></th>
									</tr>
								</thead>
								<tbody class="divide-y divide-gray-200">
									<template v-if="result_bkn" v-for="(item, idx) in result_bkn" :key="idx">
										<tr class="odd:bg-white even:bg-gray-100">
											<td class="size-px px-1" style="vertical-align: top;">
												<button type="button" class="block" data-hs-overlay="#hs-ai-invoice-modal">
													<span class="flex justify-center text-xs px-3 py-0">{{idx+1}}</span>
												</button>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 200px;">
												<span class="flex justify-start text-xs py-0">{{chekJenisJab(item)}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 80px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.tmtJabatan}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 80px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.nomorSk}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 80px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.tanggalSk}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 80px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.eselon}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top;">
												<div class="sm:flex gap-x-1">
                                                    <template v-if="item.data_simpeg">
                                                        <button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-green-600 text-white hover:bg-green-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="simpanData(item)">
                                                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4">
                                                                <path fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14Zm3.844-8.791a.75.75 0 0 0-1.188-.918l-3.7 4.79-1.649-1.833a.75.75 0 1 0-1.114 1.004l2.25 2.5a.75.75 0 0 0 1.15-.043l4.25-5.5Z" clip-rule="evenodd" />
                                                            </svg>
                                                        </button>
                                                    </template>
                                                    <template v-else>
                                                        <button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-yellow-600 text-white hover:bg-yellow-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="simpanData(item)">
                                                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4">
                                                                <path fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14Zm.75-10.25a.75.75 0 0 0-1.5 0v4.69L6.03 8.22a.75.75 0 0 0-1.06 1.06l2.5 2.5a.75.75 0 0 0 1.06 0l2.5-2.5a.75.75 0 1 0-1.06-1.06L8.75 9.44V4.75Z" clip-rule="evenodd" />
                                                            </svg>

                                                        </button>
                                                    </template>
                                                    <template v-if="item.path">
                                                        <button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="doSyncDoc(item)">
                                                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4">
                                                                <path fill-rule="evenodd" d="M2 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4Zm10.5 5.707a.5.5 0 0 0-.146-.353l-1-1a.5.5 0 0 0-.708 0L9.354 9.646a.5.5 0 0 1-.708 0L6.354 7.354a.5.5 0 0 0-.708 0l-2 2a.5.5 0 0 0-.146.353V12a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V9.707ZM12 5a1 1 0 1 1-2 0 1 1 0 0 1 2 0Z" clip-rule="evenodd" />
                                                            </svg>
                                                        </button>
                                                    </template>
                                                </div>
											</td>
										</tr>
									</template>
								</tbody>
							</table>
							<!-- End Table -->	
						</div>
					</div>
				</div>
			</div>
		</div>
	</LayoutSingkronisasi>
</template>