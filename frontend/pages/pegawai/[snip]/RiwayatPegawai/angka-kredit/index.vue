<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';
import { useModalStore } from '~/store/modalStore';

const modalUploadDoc = useModalUploadDoc();
const modalStore = useModalStore();

const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
import customParseFormat from "dayjs/plugin/customParseFormat";
dayjs.extend(customParseFormat)

const toast = useToast()
const {snip} = route.params
const showSpinner = ref(false)

onMounted(() => {
	refreshNuxtData(["getAngkaKredit"])
	modalUploadDoc.$onAction(callback, false)
	modalStore.$onAction(callbackHapusData, false)
})

const dataSimpeg = ref([])
//const result_bkn = ref([])

const { pending, data : result_simpeg, refresh} = await useLazyAsyncData('getAngkaKredit', async() => {
	console.log("CariData angka-kredit")

	let nip = $decodeBase64(snip)

  	var result = await $fetch(`/api/posts/riwayat_angka_kredit`, {
		method: 'POST',
		body: JSON.stringify({
			nip: nip,
		})
	});

	dataSimpeg.value = result

  	return result
});

const { pending : pending_bkn, data : result_bkn, error} = await useLazyAsyncData('getAngkaKreditBkn', async() => {
	console.log("CariData angka-kredit Bkn")
	let nip = $decodeBase64(snip)

	var result = await $fetch(`/api/gets/singkronisasi_bkn/getDataBKN/rw-angkakredit/${nip}`)
	return result.data
});

const editData = (item, sbkn="") => {
	let stmulai = item ? $encodeBase64(item.tmulai) : $encodeBase64(null)
	navigateTo({ path: '/pegawai/'+snip+'/RiwayatPegawai/angka-kredit/'+stmulai, query: {
		sbkn : sbkn,
	}})
}

const selData = ref([])

const hapusData = (item) => {
	selData.value = item
	modalStore.showModal("Hapus Data Angka Kredit", `Apakah anda yakin untuk menghapus data Angka Kredit dengan Nomor SK : ${item.nsk}`);
}

const callbackHapusData = async(e) => {
	showSpinner.value = true
	const {name} = e
	if(name == 'doAction' && selData.value != ""){
		var nip = $decodeBase64(snip)
		var body = JSON.stringify(selData.value);

		var {data, error} = await useFetch('/api/delete/riwayat_angka_kredit', {
			method: 'DELETE',
			body: body,
		});

		if(error.value){
			const {title, message, color} = error.value.data.data
			toast.add({
		    	id: 'error_delete_riwayat_angka_kredit',
		    	icon: "i-heroicons-x-circle",
		    	title: title,
		    	description: message,
		    	timeout: 6000,
		    	color: color,
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_delete_riwayat_angka_kredit',
		    	icon: "i-heroicons-check-circle",
		    	title: "Data Riwayat Angka Kredit",
		    	description: "Data Riwayat Angka Kredit telah di Hapus",
		    	timeout: 6000,
		  	}) 	
		}
		showSpinner.value = false
		refreshNuxtData(["getAngkaKredit","getAngkaKreditBkn"])
	}
}

function chekJenisAK(item){
	if(item.isAngkaKreditPertama == "1") return "Pertama"
	else if(item.isIntegrasi == "1") return "Integrasi"
	else if(item.isKonversi == "1") return "Konversi"
	else return "Lanjutan"
}

function checkPath(path, idx){
	if(path != null){
		if(path[idx] != undefined){
			return true
		}
	}
	return false
}

function checkIdSync(item){
	var idSync = item.id

	return _find(result_simpeg.value, function(o){
		if (o.idSync == item.id) {
			return true
		}
		return false
	})
}

async function getDataBKN(item){
	let tmulai = dayjs(item.tahunMulaiPenailan+"-"+item.bulanMulaiPenailan+"-1").format("YYYY-MM-DD")
	var dd = _find(dataSimpeg.value, function(o){
		let tmulaiSimpeg = dayjs(o.tmulai).format("YYYY-MM-DD").toString()
		return o.idSync == item.id && tmulaiSimpeg == tmulai
	})

	if(dd){
		editData(dd)	
	}else{
		let ID = $encodeBase64(item.id)
		editData(dd, ID)
	}
}

const selBody = ref([])

async function downloadDoc(item, idx){
	showSpinner.value = true

	let judul = idx == 880 ? "SK PAK" : "Dok PAK"
	let filename_idx = 880 ? "filename_sk_pak" : "filename_dok_pak"

	let filepath = item.path[idx]
	let nip = $decodeBase64(snip)

	var {data, error} = await useFetch('/api/posts/download/dokRwBkn', {
		method: 'POST',
		body: JSON.stringify({
			path : filepath["dok_uri"],
			nip: nip,
		}),
	})

	if(error.value){
		console.log(error.value)
		const {title, message, color} = error.value.data.data
		toast.add({
	    	id: 'error_download_riwayat_angka_kredit',
	    	icon: "i-heroicons-x-circle",
	    	title: title,
	    	description: message,
	    	timeout: 6000,
	    	color: color,
	  	}) 	
	}else{
		var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${data.value.filename}`)
		if(fileDisplay != null){
			const blobUrl = URL.createObjectURL(fileDisplay)
			modalUploadDoc.showModal(`Download Dokumen Angka Kredit ${judul} BKN`, "application/pdf", blobUrl, 'download')	
		}

		selBody.value = {
			"refdata": {
				nip: nip,
				tmulai: dayjs(item.tahunMulaiPenailan+"-"+item.bulanMulaiPenailan+"-1").format("YYYY-MM-DD[T]HH:mm:ss[Z]"),
			},
			"data": {
				nip: nip,
				tmulai: dayjs(item.tahunMulaiPenailan+"-"+item.bulanMulaiPenailan+"-1").format("YYYY-MM-DD[T]HH:mm:ss[Z]"),
				[filename_idx]: data.value.filename,
				do_bkn: 1,
			}
		}
	}

	showSpinner.value = false
}

const viewDoc = async(judul, filename) => {
	let nip = $decodeBase64(snip)
	//let idx = "filename_"+fileIdx;

	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Lihat ${judul}`, "application/pdf", blobUrl, true)	
	}
}

const callback = async(e) => {
	const {name} = e
	console.log("callbackSaveFilename")
	if(name == 'doAction'){
		showSpinner.value = true

		var {data, error} = await useFetch('/api/puts/riwayat_angka_kredit', {
			method: 'PUT',
			body: JSON.stringify(selBody.value),
		})

		if(error.value){
			const {title, description} = error.value.data.data
			toast.add({
				id: 'error_put_jabatan',
				icon: "i-heroicons-x-circle",
				title: title,
				description: description,
				timeout: 6000,
				color: 'red',
			}) 	
		}else{
			toast.add({
				id: 'success_put_jabatan',
				icon: "i-heroicons-check-circle",
				title: "Dokumen berhasil didownload",
				description: "Data telah di Update",
				timeout: 6000
			}) 	
		}

		showSpinner.value = false
		refreshNuxtData(["getDataRef", "getdataAngkaKreditAkhir"])
	}
}
</script>
<template>
	<AppLoadingSpinner :show="showSpinner" />
	<LayoutRiwayatPegawai>
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="flex justify-end mb-2">
					<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="editData()">Tambah</button>
				</div>
				<div class="mb-8">
					<div class="overflow-x-auto border">
						<div class="min-w-full inline-block align-middle">
							<!-- Table -->
							<table class="min-w-full divide-y divide-gray-200">
								<thead class="bg-gray-50 h-12">
									<tr>
										<th scope="col" class="px-3 text-center">
											<span class="flex justify-center text-xs font-semibold uppercase text-gray-800">
												No
											</span>
										</th>
										<th scope="col" class="px-3 text-center"></th>
										<th scope="col" class="px-3 text-center">
											<span class="flex justify-center text-xs font-semibold uppercase text-gray-800">
												Jabatan FT
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="flex justify-center text-xs font-semibold uppercase text-gray-800">
												Jenis AK
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="flex justify-center text-xs font-semibold uppercase text-gray-800">
												Priode
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="flex justify-center text-xs font-semibold uppercase text-gray-800">
												Angka Kredit Utama
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="flex justify-center text-xs font-semibold uppercase text-gray-800">
												Angka Kredit Tambahan
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="flex justify-center text-xs font-semibold uppercase text-gray-800">
												Total
											</span>
										</th>
										<th scope="col" class="px-3 text-center"></th>
									</tr>
								</thead>
								<tbody class="divide-y divide-gray-200">
									<template v-if="pending">
										<tr>
											<td colspan="7" class="size-px p-1" style="vertical-align: top; min-width: 100px;">
												<div class="flex justify-center animate-pulse">
												  	<span class="text-sm italic">Loading</span>
												</div>
											</td>
										</tr>
									</template>
									<template v-else-if="result_simpeg" v-for="(item, idx) in result_simpeg" :key="idx">
										<tr class="odd:bg-white even:bg-gray-100">
											<td class="size-px px-1" style="vertical-align: top; width: 5%">
												<span class="flex justify-center text-xs px-3 py-0">{{idx+1}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; width: 5%">
												<div class="sm:grid justify-center gap-1">
													<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="viewDoc('SK PAK', item.filename_sk_pak)" v-if="item.filename_sk_pak != null " title="Dok PAK">
														<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4">
															<path fill-rule="evenodd" d="M2 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4Zm10.5 5.707a.5.5 0 0 0-.146-.353l-1-1a.5.5 0 0 0-.708 0L9.354 9.646a.5.5 0 0 1-.708 0L6.354 7.354a.5.5 0 0 0-.708 0l-2 2a.5.5 0 0 0-.146.353V12a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V9.707ZM12 5a1 1 0 1 1-2 0 1 1 0 0 1 2 0Z" clip-rule="evenodd" />
														</svg>
													</button>
													<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="viewDoc('Dokumen PAK', item.filename_dok_pak)" v-if="item.filename_dok_pak != null " title="Dokumen PAK">
														<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4">
															<path fill-rule="evenodd" d="M2 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4Zm10.5 5.707a.5.5 0 0 0-.146-.353l-1-1a.5.5 0 0 0-.708 0L9.354 9.646a.5.5 0 0 1-.708 0L6.354 7.354a.5.5 0 0 0-.708 0l-2 2a.5.5 0 0 0-.146.353V12a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V9.707ZM12 5a1 1 0 1 1-2 0 1 1 0 0 1 2 0Z" clip-rule="evenodd" />
														</svg>
													</button>
												</div>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 200px;">
												<span class="flex justify-center text-xs py-0">{{item.RwJabatan?.njab}}</span>
												<span class="flex justify-center text-xs py-0 italic">{{item.idSync}}</span>
											</td>
											<td class="size-px px-1 text-center" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">AK {{item.jns}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.tmulai ? $dayjs(item.tmulai).format("DD-MM-YYYY").toString() : '-'}} <span>&nbsp;&nbsp;&nbsp;s/d&nbsp;&nbsp;&nbsp;</span> {{item.tselesai ? $dayjs(item.tselesai).format("DD-MM-YYYY").toString() : '-'}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0" v-if="(_includes(['Pertama', 'Lanjutan'], item.jns))">{{item.utama}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0" v-if="(_includes(['Pertama', 'Lanjutan'], item.jns))">{{item.tambahan}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">{{item.total}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top;">
												<div class="sm:flex gap-x-1">
													<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="editData(item)">
													  	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="flex-shrink-0 size-3">
														  	<path d="M13.488 2.513a1.75 1.75 0 0 0-2.475 0L6.75 6.774a2.75 2.75 0 0 0-.596.892l-.848 2.047a.75.75 0 0 0 .98.98l2.047-.848a2.75 2.75 0 0 0 .892-.596l4.261-4.262a1.75 1.75 0 0 0 0-2.474Z" />
														  	<path d="M4.75 3.5c-.69 0-1.25.56-1.25 1.25v6.5c0 .69.56 1.25 1.25 1.25h6.5c.69 0 1.25-.56 1.25-1.25V9A.75.75 0 0 1 14 9v2.25A2.75 2.75 0 0 1 11.25 14h-6.5A2.75 2.75 0 0 1 2 11.25v-6.5A2.75 2.75 0 0 1 4.75 2H7a.75.75 0 0 1 0 1.5H4.75Z" />
														</svg>
													</button>
													<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-red-600 text-white hover:bg-red-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="hapusData(item)">
													  	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="flex-shrink-0 size-3">
														  	<path fill-rule="evenodd" d="M5 3.25V4H2.75a.75.75 0 0 0 0 1.5h.3l.815 8.15A1.5 1.5 0 0 0 5.357 15h5.285a1.5 1.5 0 0 0 1.493-1.35l.815-8.15h.3a.75.75 0 0 0 0-1.5H11v-.75A2.25 2.25 0 0 0 8.75 1h-1.5A2.25 2.25 0 0 0 5 3.25Zm2.25-.75a.75.75 0 0 0-.75.75V4h3v-.75a.75.75 0 0 0-.75-.75h-1.5ZM6.05 6a.75.75 0 0 1 .787.713l.275 5.5a.75.75 0 0 1-1.498.075l-.275-5.5A.75.75 0 0 1 6.05 6Zm3.9 0a.75.75 0 0 1 .712.787l-.275 5.5a.75.75 0 0 1-1.498-.075l.275-5.5a.75.75 0 0 1 .786-.711Z" clip-rule="evenodd" />
														</svg>
													</button>
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
											<span class="flex justify-center text-xs font-semibold uppercase text-gray-800">
												No
											</span>
										</th>
										<th scope="col" class="px-3 text-center"></th>
										<th scope="col" class="px-3 text-center">
											<span class="flex justify-center text-xs font-semibold uppercase text-gray-800">
												Jabatan FT
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="flex justify-center text-xs font-semibold uppercase text-gray-800">
												Jenis AK
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="flex justify-center text-xs font-semibold uppercase text-gray-800">
												Priode
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="flex justify-center text-xs font-semibold uppercase text-gray-800">
												Angka Kredit Utama
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="flex justify-center text-xs font-semibold uppercase text-gray-800">
												Angka Kredit Penunjang
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="flex justify-center text-xs font-semibold uppercase text-gray-800">
												Total
											</span>
										</th>
										<th scope="col" class="px-3 text-center"></th>
									</tr>
								</thead>
								<tbody class="divide-y divide-gray-200">
									<template v-if="pending_bkn">
										<tr>
											<td colspan="8" align="center">
												<div class="flex justify-center animate-pulse">
												  	<span class="text-sm italic">Loading</span>
												</div>
											</td>
										</tr>
									</template>
									<template v-else-if="error">
										<tr>
											<td colspan="8" align="center">
												<span class="flex justify-center text-xs py-1 text-red-500">Get Data BKN Gagal :<b> {{error.data.data.message || error.data.data.errors}}</b></span>
											</td>
										</tr>
									</template>
									<template v-else-if="result_bkn" v-for="(item, idx) in result_bkn" :key="idx">
										<tr class="odd:bg-white even:bg-gray-100">
											<td class="size-px px-1" style="vertical-align: top; width: 5%">
												<span class="flex justify-center text-xs px-3 py-0">{{idx+1}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; width: 5%">
												<div class="sm:grid justify-center gap-1">
													<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="downloadDoc(item, 879)" v-if="checkPath(item.path, 879)" title="Dok PAK">
														<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4">
															<path fill-rule="evenodd" d="M2 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4Zm10.5 5.707a.5.5 0 0 0-.146-.353l-1-1a.5.5 0 0 0-.708 0L9.354 9.646a.5.5 0 0 1-.708 0L6.354 7.354a.5.5 0 0 0-.708 0l-2 2a.5.5 0 0 0-.146.353V12a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V9.707ZM12 5a1 1 0 1 1-2 0 1 1 0 0 1 2 0Z" clip-rule="evenodd" />
														</svg>
													</button>
													<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="downloadDoc(item, 880)" v-if="checkPath(item.path, 880)" title="SK PAK">
														<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4">
															<path fill-rule="evenodd" d="M2 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4Zm10.5 5.707a.5.5 0 0 0-.146-.353l-1-1a.5.5 0 0 0-.708 0L9.354 9.646a.5.5 0 0 1-.708 0L6.354 7.354a.5.5 0 0 0-.708 0l-2 2a.5.5 0 0 0-.146.353V12a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V9.707ZM12 5a1 1 0 1 1-2 0 1 1 0 0 1 2 0Z" clip-rule="evenodd" />
														</svg>
													</button>
												</div>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 200px;">
												<span class="flex justify-start text-xs py-0">{{item.namaJabatan}}</span>
												<span class="flex justify-start text-xs py-0 italic">{{item.id}}</span>
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
												<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-green-600 text-white hover:bg-green-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="getDataBKN(item)">
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
			</div>
		</div>
	</LayoutRiwayatPegawai>
</template>