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
	//refreshNuxtData(["getAngkaKredit"])
})

const result_simpeg = ref([])
const result_bkn = ref([])
const result = ref([{
    "Sumber": null,
    "bulanMulaiPenailan": "1",
    "bulanSelesaiPenailan": "12",
    "created_at": "0001-01-01T00:00:00Z",
    "id": "8ae483a763849741016390bd144f43d5",
    "isAngkaKreditPertama": "",
    "isIntegrasi": "",
    "isKonversi": "",
    "kreditBaruTotal": "13.250",
    "kreditPenunjangBaru": "0.750",
    "kreditUtamaBaru": "12.500",
    "namaJabatan": "Guru Pertama",
    "nomorSk": "800/3255/5.3.05/2016",
    "path": null,
    "pns": "A8ACA7CDDD6C3912E040640A040269BB",
    "rwJabatan": "8ae482874aed68ed014b106e370a17c9",
    "tahunMulaiPenailan": "2016",
    "tahunSelesaiPenailan": "2016",
    "tanggalSk": "30-12-2016",
    "updated_at": "0001-01-01T00:00:00Z"
},
{
    "Sumber": null,
    "bulanMulaiPenailan": "7",
    "bulanSelesaiPenailan": "12",
    "created_at": "0001-01-01T00:00:00Z",
    "id": "8ae4828735d2a0600135e5ae523604c9",
    "isAngkaKreditPertama": "1",
    "isIntegrasi": "",
    "isKonversi": "",
    "kreditBaruTotal": "19.608",
    "kreditPenunjangBaru": "2.000",
    "kreditUtamaBaru": "17.608",
    "namaJabatan": "Guru Pratama",
    "nomorSk": "420/830/1/4929/148",
    "path": null,
    "pns": "A8ACA7CDDD6C3912E040640A040269BB",
    "rwJabatan": "8ae482873542f4c6013556414089003c",
    "tahunMulaiPenailan": "2009",
    "tahunSelesaiPenailan": "2011",
    "tanggalSk": "30-12-2011",
    "updated_at": "0001-01-01T00:00:00Z"
},
{
    "Sumber": null,
    "bulanMulaiPenailan": "7",
    "bulanSelesaiPenailan": "6",
    "created_at": "0001-01-01T00:00:00Z",
    "id": "8ae4828735d2a0600135e5adbe6a043a",
    "isAngkaKreditPertama": "1",
    "isIntegrasi": "",
    "isKonversi": "",
    "kreditBaruTotal": "35.860",
    "kreditPenunjangBaru": "1.150",
    "kreditUtamaBaru": "34.710",
    "namaJabatan": "Kredit Mulai Sebagai Fung.Tertentu",
    "nomorSk": "823/SET/1549/2009/007",
    "path": null,
    "pns": "A8ACA7CDDD6C3912E040640A040269BB",
    "rwJabatan": "",
    "tahunMulaiPenailan": "2008",
    "tahunSelesaiPenailan": "2009",
    "tanggalSk": "30-06-2009",
    "updated_at": "0001-01-01T00:00:00Z"
},
{
    "Sumber": null,
    "bulanMulaiPenailan": "1",
    "bulanSelesaiPenailan": "12",
    "created_at": "0001-01-01T00:00:00Z",
    "id": "8ae4829d4e1fc8d8014e28b7c4e87687",
    "isAngkaKreditPertama": "",
    "isIntegrasi": "",
    "isKonversi": "",
    "kreditBaruTotal": "2.925",
    "kreditPenunjangBaru": "0.000",
    "kreditUtamaBaru": "2.925",
    "namaJabatan": null,
    "nomorSk": "800/3011/5.5.01/2014/69",
    "path": null,
    "pns": "A8ACA7CDDD6C3912E040640A040269BB",
    "rwJabatan": "",
    "tahunMulaiPenailan": "2014",
    "tahunSelesaiPenailan": "2014",
    "tanggalSk": "31-12-2014",
    "updated_at": "0001-01-01T00:00:00Z"
},
{
    "Sumber": null,
    "bulanMulaiPenailan": "1",
    "bulanSelesaiPenailan": "12",
    "created_at": "0001-01-01T00:00:00Z",
    "id": "8ae4829d4e1fc8d8014e28b64946748d",
    "isAngkaKreditPertama": "",
    "isIntegrasi": "",
    "isKonversi": "",
    "kreditBaruTotal": "77.925",
    "kreditPenunjangBaru": "0.000",
    "kreditUtamaBaru": "77.925",
    "namaJabatan": null,
    "nomorSk": "800/3820/5.5.01/2013/139",
    "path": null,
    "pns": "A8ACA7CDDD6C3912E040640A040269BB",
    "rwJabatan": "",
    "tahunMulaiPenailan": "2013",
    "tahunSelesaiPenailan": "2013",
    "tanggalSk": "31-12-2013",
    "updated_at": "0001-01-01T00:00:00Z"
},
{
    "Sumber": null,
    "bulanMulaiPenailan": "1",
    "bulanSelesaiPenailan": "12",
    "created_at": "0001-01-01T00:00:00Z",
    "id": "8ae4829d4e1fc8d8014e28b5bdf37426",
    "isAngkaKreditPertama": "",
    "isIntegrasi": "",
    "isKonversi": "",
    "kreditBaruTotal": "4.568",
    "kreditPenunjangBaru": "0.000",
    "kreditUtamaBaru": "4.568",
    "namaJabatan": null,
    "nomorSk": "420/830/1/1783/139",
    "path": null,
    "pns": "A8ACA7CDDD6C3912E040640A040269BB",
    "rwJabatan": "",
    "tahunMulaiPenailan": "2012",
    "tahunSelesaiPenailan": "2012",
    "tanggalSk": "31-12-2012",
    "updated_at": "0001-01-01T00:00:00Z"
},
{
    "Sumber": null,
    "bulanMulaiPenailan": "1",
    "bulanSelesaiPenailan": "12",
    "created_at": "0001-01-01T00:00:00Z",
    "id": "8ae483a763849741016390bbfa054310",
    "isAngkaKreditPertama": "",
    "isIntegrasi": "",
    "isKonversi": "",
    "kreditBaruTotal": "11.250",
    "kreditPenunjangBaru": "0.750",
    "kreditUtamaBaru": "10.500",
    "namaJabatan": "Guru Pertama",
    "nomorSk": "800/4302/5.3.05/2015",
    "path": null,
    "pns": "A8ACA7CDDD6C3912E040640A040269BB",
    "rwJabatan": "8ae482874aed68ed014b106e370a17c9",
    "tahunMulaiPenailan": "2015",
    "tahunSelesaiPenailan": "2015",
    "tanggalSk": "31-12-2015",
    "updated_at": "0001-01-01T00:00:00Z"
},
{
    "Sumber": null,
    "bulanMulaiPenailan": "1",
    "bulanSelesaiPenailan": "6",
    "created_at": "0001-01-01T00:00:00Z",
    "id": "8ae483a763849741016390c218d64781",
    "isAngkaKreditPertama": "",
    "isIntegrasi": "",
    "isKonversi": "",
    "kreditBaruTotal": "5.250",
    "kreditPenunjangBaru": "0.000",
    "kreditUtamaBaru": "5.250",
    "namaJabatan": "Guru Pertama",
    "nomorSk": "800/1085.2/5.3.05/2018/",
    "path": null,
    "pns": "A8ACA7CDDD6C3912E040640A040269BB",
    "rwJabatan": "8ae482874aed68ed014b106e370a17c9",
    "tahunMulaiPenailan": "2018",
    "tahunSelesaiPenailan": "2018",
    "tanggalSk": "30-06-2018",
    "updated_at": "0001-01-01T00:00:00Z"
},
{
    "Sumber": null,
    "bulanMulaiPenailan": "1",
    "bulanSelesaiPenailan": "12",
    "created_at": "0001-01-01T00:00:00Z",
    "id": "8ae483a763849741016390c0b312468b",
    "isAngkaKreditPertama": "",
    "isIntegrasi": "",
    "isKonversi": "",
    "kreditBaruTotal": "12.250",
    "kreditPenunjangBaru": "0.750",
    "kreditUtamaBaru": "11.500",
    "namaJabatan": "Guru Pertama",
    "nomorSk": "800/1069.2/5.3.05/2017",
    "path": null,
    "pns": "A8ACA7CDDD6C3912E040640A040269BB",
    "rwJabatan": "8ae482874aed68ed014b106e370a17c9",
    "tahunMulaiPenailan": "2017",
    "tahunSelesaiPenailan": "2017",
    "tanggalSk": "31-12-2017",
    "updated_at": "0001-01-01T00:00:00Z"
},
{
    "Sumber": null,
    "bulanMulaiPenailan": "1",
    "bulanSelesaiPenailan": "12",
    "created_at": "0001-01-01T00:00:00Z",
    "id": "8ae483a577259067017737e365955f91",
    "isAngkaKreditPertama": "",
    "isIntegrasi": "",
    "isKonversi": "",
    "kreditBaruTotal": "14.250",
    "kreditPenunjangBaru": "0.750",
    "kreditUtamaBaru": "13.500",
    "namaJabatan": null,
    "nomorSk": "800/1387.1/5.3.05/2019",
    "path": null,
    "pns": "A8ACA7CDDD6C3912E040640A040269BB",
    "rwJabatan": "",
    "tahunMulaiPenailan": "2019",
    "tahunSelesaiPenailan": "2019",
    "tanggalSk": "31-12-2019",
    "updated_at": "0001-01-01T00:00:00Z"
},
{
    "Sumber": null,
    "bulanMulaiPenailan": "1",
    "bulanSelesaiPenailan": "12",
    "created_at": "0001-01-01T00:00:00Z",
    "id": "8ae483a577259067017737e405995fb0",
    "isAngkaKreditPertama": "",
    "isIntegrasi": "",
    "isKonversi": "",
    "kreditBaruTotal": "12.250",
    "kreditPenunjangBaru": "0.750",
    "kreditUtamaBaru": "11.500",
    "namaJabatan": null,
    "nomorSk": "800/1629/5.3.05/2020",
    "path": null,
    "pns": "A8ACA7CDDD6C3912E040640A040269BB",
    "rwJabatan": "",
    "tahunMulaiPenailan": "2020",
    "tahunSelesaiPenailan": "2020",
    "tanggalSk": "31-12-2020",
    "updated_at": "0001-01-01T00:00:00Z"
},
{
    "Sumber": null,
    "bulanMulaiPenailan": "7",
    "bulanSelesaiPenailan": "12",
    "created_at": "0001-01-01T00:00:00Z",
    "id": "8ae483a577259067017737e2db6a5f5a",
    "isAngkaKreditPertama": "",
    "isIntegrasi": "",
    "isKonversi": "",
    "kreditBaruTotal": "5.750",
    "kreditPenunjangBaru": "0.000",
    "kreditUtamaBaru": "5.750",
    "namaJabatan": null,
    "nomorSk": "800/1085.2/5.3.05/2018",
    "path": null,
    "pns": "A8ACA7CDDD6C3912E040640A040269BB",
    "rwJabatan": "",
    "tahunMulaiPenailan": "2018",
    "tahunSelesaiPenailan": "2018",
    "tanggalSk": "31-12-2018",
    "updated_at": "0001-01-01T00:00:00Z"
}])

async function getDataBKN(){
	if(snip){
		showSpinner.value = true
		console.log("CariData")

		let nip = $decodeBase64(snip)
		var body = JSON.stringify({
			nip: nip,
		});

		result_simpeg.value = await $fetch('/api/posts/riwayat_angka_kredit', {
			method: 'POST',
			body: body,
		});

		var bkn = await $fetch('/api/gets/singkronisasi_bkn/getDataBKN/rw-angkakredit/'+nip)
		result_bkn.value = bkn.data
		result_bkn.value = _orderBy(bkn.data, ['tahunSelesaiPenailan', 'bulanSelesaiPenailan'], ['desc', 'desc'])
		//result_bkn.value = _orderBy(result.value, ['tahunSelesaiPenailan', 'bulanSelesaiPenailan'], ['desc', 'desc'])


		if(result_simpeg.value){
			for (var i = 0; i < result_simpeg.value.length; i++) {
				var data_bkn = _find(result_bkn.value, {
					bulanMulaiPenailan: (dayjs(dayjs(result_simpeg.value[i].tmulai, "YYYY-MM-DD")).format("MM") *1).toString(),
					tahunMulaiPenailan: dayjs(dayjs(result_simpeg.value[i].tmulai, "YYYY-MM-DD")).format("YYYY"),
				})
				result_simpeg.value[i]['data_bkn'] = JSON.stringify(data_bkn)
			}
		}

		if(result_bkn.value){
			for (var i = 0; i < result_bkn.value.length; i++) {
				var tmulai = dayjs(result_bkn.value[i].tahunMulaiPenailan+"-"+result_bkn.value[i].bulanMulaiPenailan+"-1").format("YYYY-MM-DD[T]HH:mm:ss+07:00")
				var data_simpeg = _find(result_simpeg.value, {
					tmulai: tmulai,
				})
				//console.log(data_simpeg)
				result_bkn.value[i]['data_simpeg'] = JSON.stringify(data_simpeg)
			}
		}

		showSpinner.value = false
	}
}

/*const editAK = (item) => {
	let stmulai = item ? $encodeBase64(item.tmulai) : $encodeBase64(null)
	navigateTo({ path: '/pegawai/'+snip+'/RiwayatPegawai/angka-kredit/'+stmulai })
}

const hapusAK = async(item) => {
	var body = JSON.stringify({
		nip: item.nip,
		tmulai: item.tmulai,
	});

	var {data, error} = await useFetch('/api/delete/riwayat_angka_kredit', {
		method: 'DELETE',
		body: body,
	});

	if(error.value){
		toast.add({
	    	id: 'error_delete_riwayat_angka_kredit',
	    	description: error.value.data.data,
	    	timeout: 6000
	  	}) 	
	}else{
		toast.add({
	    	id: 'success_delete_riwayat_angka_kredit',
	    	description: "Data berhasil di Hapus",
	    	timeout: 6000
	  	}) 	
	}

	//refreshNuxtData(["getAngkaKredit"])
}*/
const refDataBkn = ref({
	//"Sumber": null,
    //"bulanMulaiPenailan": "7",
    //"bulanSelesaiPenailan": "12",
    //"created_at": "0001-01-01T00:00:00Z",
    "id": "idSync",
    //"isAngkaKreditPertama": "",
    //"isIntegrasi": "",
    //"isKonversi": "",
    "kreditBaruTotal": "total",
    "kreditPenunjangBaru": "tambahan",
    "kreditUtamaBaru": "utama",
    //"namaJabatan": null,
    "nomorSk": "nsk",
    //"path": null,
    //"pns": "A8ACA7CDDD6C3912E040640A040269BB",
    //"rwJabatan": "",
    //"tahunMulaiPenailan": "2018",
    //"tahunSelesaiPenailan": "2018",
    "tanggalSk": "tsk",
    //"updated_at": "0001-01-01T00:00:00Z"
})

async function simpanData(item) {
	showSpinner.value = true
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

	compacted.tsk = compacted.tsk ? dayjs(dayjs(compacted.tsk, "DD-MM-YYYY")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
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
  	getDataBKN()
}

function chekJenisAK(item){
	if(item.isAngkaKreditPertama == "1") return "Pertama"
	else if(item.isIntegrasi == "1") return "Integrasi"
	else if(item.isKonversi == "1") return "Konversi"
	else return "Lanjutan"
}

const refUpdateBkn = ref({
    
})

async function addToBKN(item){
    showSpinner.value = true

    let body = JSON.stringify(item)
    console.log(body)

    let nip = $decodeBase64(snip)
    var {data, error} = await useFetch('/api/puts/singkronisasi_bkn/postDataBKNAK/'+nip, {
        method: 'POST',
        body: body,
    })

    console.log(data, error)

    if(error.value){
        toast.add({
            id: 'error_put_singkronisasi_ak',
            description: error.value.data.data,
            timeout: 6000
        })  
    }else{
        toast.add({
            id: 'success_post_singkronisasi_ak',
            description: "Data berhasil di Ditambahkan",
            timeout: 6000
        })  
    }

    showSpinner.value = false
    getDataBKN()
}

async function doDeleteBKN(item) {
    showSpinner.value = true

    var {data, error} = await useFetch('/api/delete/singkronisasi_bkn/delDataBKNAK/'+item.id, {
        method: 'DELETE'
    })

    console.log(data, error)

    if(error.value){
        toast.add({
            id: 'error_delete_singkronisasi_ak',
            description: error.value.data.data,
            timeout: 6000
        })  
    }else{
        toast.add({
            id: 'success_delete_singkronisasi_ak',
            description: "Data AK BKN berhasil di Dihapus",
            timeout: 6000
        })  
    }

    showSpinner.value = false
    getDataBKN()
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
					<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="getDataBKN()">Get Data BKN</button>
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
												Jabatan FT
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Jenis AK
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Priode
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Angka Kredit Utama
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Angka Kredit Tambahan
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Total
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
											<td class="size-px px-1 text-center" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">{{item.jabatan_ft?.nama}}</span>
											</td>
											<td class="size-px px-1 text-center" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">AK {{item.jns}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.tmulai ? $dayjs(item.tmulai).format("DD-MM-YYYY").toString() : '-'}} <span>&nbsp;&nbsp;&nbsp;s/d&nbsp;&nbsp;&nbsp;</span> {{item.tselesai ? $dayjs(item.tselesai).format("DD-MM-YYYY").toString() : '-'}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.utama}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.tambahan}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">{{item.total}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top;">
												<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-green-600 text-white hover:bg-green-700 disabled:opacity-50 disabled:pointer-events-none" @click="addToBKN(item)">
                                                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="flex-shrink-0 size-3">
                                                    <path fill-rule="evenodd" d="M4.5 13a3.5 3.5 0 0 1-1.41-6.705A3.5 3.5 0 0 1 9.72 4.124a2.5 2.5 0 0 1 3.197 3.018A3.001 3.001 0 0 1 12 13H4.5Zm.72-5.03a.75.75 0 0 0 1.06 1.06l.97-.97v2.69a.75.75 0 0 0 1.5 0V8.06l.97.97a.75.75 0 1 0 1.06-1.06L8.53 5.72a.75.75 0 0 0-1.06 0L5.22 7.97Z" clip-rule="evenodd" />
                                                </svg>
                                            </button>
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
												Jabatan FT
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Jenis AK
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Priode
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Angka Kredit Utama
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Angka Kredit Penunjang
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Total
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
											<td class="size-px px-1 text-center" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">{{item.namaJabatan}}</span>
											</td>
											<td class="size-px px-1 text-center" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">AK {{chekJenisAK(item)}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0">1-{{item.bulanMulaiPenailan}}-{{item.tahunMulaiPenailan}} <span>&nbsp;&nbsp;&nbsp;s/d&nbsp;&nbsp;&nbsp;</span>1-{{item.bulanSelesaiPenailan}}-{{item.tahunSelesaiPenailan}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.kreditUtamaBaru}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.kreditPenunjangBaru}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">{{item.kreditBaruTotal}}</span>
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
                                                    <button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-red-600 text-white hover:bg-red-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="doDeleteBKN(item)">
                                                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4">
                                                            <path fill-rule="evenodd" d="M5 3.25V4H2.75a.75.75 0 0 0 0 1.5h.3l.815 8.15A1.5 1.5 0 0 0 5.357 15h5.285a1.5 1.5 0 0 0 1.493-1.35l.815-8.15h.3a.75.75 0 0 0 0-1.5H11v-.75A2.25 2.25 0 0 0 8.75 1h-1.5A2.25 2.25 0 0 0 5 3.25Zm2.25-.75a.75.75 0 0 0-.75.75V4h3v-.75a.75.75 0 0 0-.75-.75h-1.5ZM6.05 6a.75.75 0 0 1 .787.713l.275 5.5a.75.75 0 0 1-1.498.075l-.275-5.5A.75.75 0 0 1 6.05 6Zm3.9 0a.75.75 0 0 1 .712.787l-.275 5.5a.75.75 0 0 1-1.498-.075l.275-5.5a.75.75 0 0 1 .786-.711Z" clip-rule="evenodd" />
                                                        </svg>
                                                    </button>
													<template v-if="item.path == {}">
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