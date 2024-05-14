<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)
const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const toast = useToast()
import customParseFormat from "dayjs/plugin/customParseFormat";
dayjs.extend(customParseFormat)

const showSpinner = ref(false)
const {snip} = route.params
//const dataPangkat = ref([])
const result_bknx = ref([
{
    "golongan": "III/d",
    "golonganId": "34",
    "id": "8ae483c68026f7960180314c64470589",
    "idPns": "A8ACA7CD3C7F3912E040640A040269BB",
    "jenisKPId": "202",
    "jenisKPNama": "Pilihan (Jabatan Fungsional Tertentu)",
    "jumlahKreditTambahan": "0",
    "jumlahKreditUtama": "0",
    "masaKerjaGolonganBulan": "1",
    "masaKerjaGolonganTahun": "12",
    "nipBaru": "198602162008031001",
    "nipLama": "",
    "noPertekBkn": "MG-22172000204",
    "pangkat": "Penata Tingkat I",
    "path": null,
    "skNomor": "224 TAHUN 2022",
    "skTanggal": "31-03-2022",
    "tglPertekBkn": "31-03-2022",
    "tmtGolongan": "2022-04-01T00:00:00Z"
},
{
    "golongan": "III/c",
    "golonganId": "33",
    "id": "8ae483a76188d3200161925279da6d05",
    "idPns": "A8ACA7CD3C7F3912E040640A040269BB",
    "jenisKPId": "201",
    "jenisKPNama": "Pilihan (Jabatan Struktural)",
    "jumlahKreditTambahan": "0",
    "jumlahKreditUtama": "0",
    "masaKerjaGolonganBulan": "1",
    "masaKerjaGolonganTahun": "8",
    "nipBaru": "198602162008031001",
    "nipLama": "",
    "noPertekBkn": "MG-22172000076",
    "pangkat": "Penata",
    "path":
    {
        "858":
        {
            "dok_id": "858",
            "dok_nama": "Dok SK KP",
            "dok_uri": "peremajaan/usulan/858_16da5b25-ed94-4252-9ab8-93b3ba51253a.pdf",
            "object": "peremajaan/usulan/858_16da5b25-ed94-4252-9ab8-93b3ba51253a.pdf",
            "slug": "858"
        }
    },
    "skNomor": "147 TAHUN 2018",
    "skTanggal": "29-03-2018",
    "tglPertekBkn": "14-02-2018",
    "tmtGolongan": "2018-04-01T00:00:00Z"
},
{
    "golongan": "III/b",
    "golonganId": "32",
    "id": "ff80808152cac93b0152da44457c4938",
    "idPns": "A8ACA7CD3C7F3912E040640A040269BB",
    "jenisKPId": "101",
    "jenisKPNama": "Reguler",
    "jumlahKreditTambahan": "0",
    "jumlahKreditUtama": "0",
    "masaKerjaGolonganBulan": "1",
    "masaKerjaGolonganTahun": "6",
    "nipBaru": "198602162008031001",
    "nipLama": "",
    "noPertekBkn": "MG-22172000134",
    "pangkat": "Penata Muda Tingkat I",
    "path": null,
    "skNomor": "112 TAHUN 2016",
    "skTanggal": "31-03-2016",
    "tglPertekBkn": "13-02-2016",
    "tmtGolongan": "2016-04-01T00:00:00Z"
},
{
    "golongan": "III/a",
    "golonganId": "31",
    "id": "ff80808136c894560136c8c5e54706d9",
    "idPns": "A8ACA7CD3C7F3912E040640A040269BB",
    "jenisKPId": "203",
    "jenisKPNama": "Pilihan (Penyesuaian Ijazah)",
    "jumlahKreditTambahan": "0",
    "jumlahKreditUtama": "0",
    "masaKerjaGolonganBulan": "1",
    "masaKerjaGolonganTahun": "2",
    "nipBaru": "198602162008031001",
    "nipLama": "",
    "noPertekBkn": "MG-22102000160",
    "pangkat": "Penata Muda",
    "path": null,
    "skNomor": "SK.823.5 - 29.b TAHUN 2012",
    "skTanggal": null,
    "tglPertekBkn": "30-03-2012",
    "tmtGolongan": "2012-04-01T00:00:00Z"
},
{
    "golongan": "II/c",
    "golonganId": "23",
    "id": "A8ACA8A0B4653912E040640A040269BB",
    "idPns": "A8ACA7CD3C7F3912E040640A040269BB",
    "jenisKPId": "211",
    "jenisKPNama": "Gol. dari Pengadaan CPNS/PNS",
    "jumlahKreditTambahan": "0",
    "jumlahKreditUtama": "0",
    "masaKerjaGolonganBulan": "0",
    "masaKerjaGolonganTahun": "3",
    "nipBaru": "198602162008031001",
    "nipLama": "",
    "noPertekBkn": "02/KEP/1998",
    "pangkat": "Pengatur",
    "path": null,
    "skNomor": "SK.813.2 - 16 Tahun 2008",
    "skTanggal": null,
    "tglPertekBkn": null,
    "tmtGolongan": "2008-03-01T00:00:00Z"
}])

const refDataBkn = ref({
	"golonganId": "kgolru",
	"id": "idSync",
	"jenisKPId": "knpang",
	"nipBaru": "nip",
	//"jumlahKreditTambahan": "string",
	//"jumlahKreditUtama": "string",
	"masaKerjaGolonganBulan": "mskerja",
	"masaKerjaGolonganTahun": "blnkerja",
	"noPertekBkn": "nntbakn",
	"skNomor": "nskpang",
	"skTanggal": "tskpang",
	"tglPertekBkn": "tntbakn",
	"tmtGolongan": "tmtpang"
})

const result_simpeg = ref([])
const result_bkn = ref([])
onMounted(() => {
	//modalUploadDoc.$onAction(callback, false)
})

const getDataBKN = async() => {
	if(snip){
		let nip = $decodeBase64(snip)
		var body = JSON.stringify({
			nip: nip,
		});

		result_simpeg.value = await $fetch('/api/posts/riwayat_pangkat', {
			method: 'POST',
			body: body,
		});

		
		var bkn = await $fetch('/api/gets/singkronisasi_bkn/getDataBKN/rw-golongan/'+nip)

		result_bkn.value = bkn.data

		if(result_simpeg.value){
			for (var i = 0; i < result_simpeg.value.length; i++) {
				var tmtpang = dayjs(dayjs(result_simpeg.value[i].tmtpang, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]").toString();
				var data_bkn = _find(result_bkn.value, {
					golonganId: result_simpeg.value[i].kgolru.toString(),
					tmtGolongan: tmtpang,
					jenisKPId: result_simpeg.value[i].knpang.toString(),
				})

				result_simpeg.value[i]['data_bkn'] = data_bkn
			}
		}

		if(result_bkn.value){
			for (var i = 0; i < result_bkn.value.length; i++) {
				var tmtGolongan = dayjs(dayjs(result_bkn.value[i].tmtGolongan, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss").toString()+"+07:00";
				var data_simpeg = _find(result_simpeg.value, {
					kgolru: result_bkn.value[i].golonganId *1,
					tmtpang: tmtGolongan,
					knpang: result_bkn.value[i].jenisKPId *1,
				})

				result_bkn.value[i]['data_simpeg'] = data_simpeg
			}
		}
	}
}

function checkData (item, akey, bkey){
	let key = item.kgolru +'-'+item.tmtpang
	let aval = item[akey]
	if(akey == 'tskpang' && aval != null){
		aval = dayjs(item[akey]).format("DD-MM-YYYY").toString()
	}
	if(akey == 'tmtpang' && aval != null){
		aval = dayjs(dayjs(item[akey], "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]").toString();
	}
	if(item.data_bkn && aval == item.data_bkn[bkey]){
		return false
	}
	return true
}

async function simpanData(item) {
	showSpinner.value = true
	let nip = $decodeBase64(snip)
	var nMethod = item.data_simpeg ? 'Update' : 'Simpan'; 
	
	item = _pickBy(item, function(val, key) {
		return _includes(_keys(refDataBkn.value), key);
	})

	var compacted = _mapKeys(item, function(val, key){
		return _find(refDataBkn.value, function(val1, key1){
			return key == key1
		})
	})

	compacted.tntbakn = compacted.tntbakn ? dayjs(dayjs(compacted.tntbakn, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tskpang = compacted.tskpang ? dayjs(dayjs(compacted.tskpang, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tmtpang = compacted.tmtpang ? dayjs(dayjs(compacted.tmtpang, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.kgolru = compacted.kgolru *1
	compacted.mskerja = compacted.mskerja *1
	compacted.blnkerja = compacted.blnkerja *1	
	compacted.knpang = compacted.knpang *1
	//compacted.akhir = compacted.akhir == true ? 1 : 0

	var dr = {
		nip: nip,
		kgolru: compacted.kgolru,
		knpang: compacted.knpang,
		tmtpang: compacted.tmtpang,
	}

	if(nMethod == 'Update'){
		var body = {
			"refdata" : dr,
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_pangkat', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_pangkat',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_pangkat',
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		//compacted.akhir = 1
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_pangkat/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_pangkat',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_pangkat',
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
		}
	}
  	showSpinner.value = false
  	getDataBKN()
}

async function doSyncDoc(item){
	//let nip = $decodeBase64(snip)
	showSpinner.value = true
	var filepath = item.path['858'].dok_uri
	var {data, error} = await useFetch('/api/posts/singkronisasi_bkn/getDocFromBkn/', {
		method: 'POST',
		body: JSON.stringify({
			path : filepath,
			RiwayatPangkat: {
				nip: item.nipBaru,
				kgolru: item.golonganId*1,
				knpang: item.jenisKPId*1,
				tmtpang: item.tmtGolongan,
			},
		})
	})

	if(error.value){
		toast.add({
	    	id: 'error_post_upload_pangkat',
	    	description: error.value.data.data,
	    	timeout: 6000
	  	}) 	
	}else{
		toast.add({
	    	id: 'success_post_upload_pangkat',
	    	description: `Dokumen Pangkat berhasil di Download`,
	    	timeout: 6000
	  	}) 	
	}
	showSpinner.value = false
	getDataBKN()
	//console.log(result)
	doUploadDoc(result)
}

const doUploadDoc = async(item) => {
	let nip = $decodeBase64(snip)
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${item.filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen Pangkat`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen Pangkat`, "application/pdf", null)	
	}
}
</script>
<template>
	<AppLoadingSpinner :show="showSpinner" />
	<LayoutSingkronisasi>
		<TokenInfo />
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="flex justify-between mb-6">
					<h2 class="text-xl font-bold text-blue-600">Riwayat Pangkat</h2>
					<button class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="getDataBKN()">Get Data BKN</button>
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
												Pangkat/Gol
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												TMT
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Pertek BKN
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Nomor SK
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Tanggal SK
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Masa Kerja
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Jenis KP
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
											<td class="size-px px-1 text-center" style="vertical-align: top; min-width: 100px;" :class="checkData(item, 'kgolru', 'golonganId') ? 'bg-red-200': '' ">
												<span class="flex justify-center text-xs py-0">
													{{item.pangkat?.pangkat_gol}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;" :class="checkData(item, 'tmtpang', 'tmtGolongan') ? 'bg-red-200': '' ">
												<span class="flex justify-center text-center text-xs py-0">{{item.tmtpang ? $dayjs(item.tmtpang).format("DD-MM-YYYY").toString() : '-'}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;" :class="checkData(item, 'nntbakn', 'noPertekBkn') ? 'bg-red-200': '' ">
												<span class="flex justify-center text-center text-xs py-0">{{item.nntbakn}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;" :class="checkData(item, 'nskpang', 'skNomor') ? 'bg-red-200': '' ">
												<span class="flex justify-center text-center text-xs py-0">{{item.nskpang}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;" :class="checkData(item, 'tskpang', 'skTanggal') ? 'bg-red-200': '' ">
												<span class="flex justify-center text-center text-xs py-0">{{item.tskpang ? $dayjs(item.tskpang).format("DD-MM-YYYY").toString() : '-'}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;" :class="checkData(item, 'mskerja', 'masaKerjaGolonganTahun') && checkData(item, 'blnkerja', 'masaKerjaGolonganBulan') ? 'bg-red-200': '' ">
												<span class="flex justify-center text-center text-xs py-0">
													{{item.mskerja ?? 0}} th {{item.blnkerja ?? 0}} bln
												</span>
												
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;" :class="checkData(item, 'knpang', 'jenisKPId') ? 'bg-red-200': '' ">
												<span class="flex justify-start text-xs py-0">{{item.jenis_kp?.nama}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top;">
												<div class="sm:flex gap-x-1">
													<template v-if="item.filename">
														<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="doUploadDoc(item)">
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
												Pangkat/Gol
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												TMT
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Pertek BKN
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Nomor SK
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Tanggal SK
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Masa Kerja
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Jenis KP
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
												<span class="flex justify-center text-xs py-0">
													{{item.pangkat}} ({{item.golongan}})
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.tmtGolongan ? $dayjs(item.tmtGolongan).format("DD-MM-YYYY").toString() : '-'}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.noPertekBkn}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.skNomor}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.skTanggal}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0">
													{{item.masaKerjaGolonganTahun ?? 0}} th {{item.masaKerjaGolonganBulan ?? 0}} bln
												</span>
												
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-start text-xs py-0">{{item.jenisKPNama}}</span>
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