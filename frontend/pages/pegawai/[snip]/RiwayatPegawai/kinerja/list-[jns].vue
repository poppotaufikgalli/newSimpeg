<script setup>
const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const toast = useToast()
const {snip, jns} = route.params

const showSpinner = ref(false)

onMounted(() => {
	refreshNuxtData(["getDataKinerja"])
})

const ratings = ref([
	{id: 1, nama: 'Diatas Ekspektasi'},
	{id: 2, nama: 'Sesuai Ekspektasi'},
	{id: 3, nama: 'Dibawah Ekspektasi'},
])
const kuadrankinerjas = ref([
	{id: 1, nama: 'Sangat Baik'},
	{id: 2, nama: 'Baik'},
	{id: 3, nama: 'Butuh Perbaikan'},
	{id: 4, nama: 'Kurang'},
	{id: 5, nama: 'Sangat Kurang'},
])

const dataSimpeg = ref([])
const { pending, data, refresh} = await useLazyAsyncData('getDataKinerja', async() => {
	console.log("CariData Kinerja")

	let nip = $decodeBase64(snip)
	var body = JSON.stringify({
		nip: nip,
		jns:jns,
	});

	var result = await $fetch('/api/posts/riwayat_kinerja', {
		method: 'POST',
		body: body,
	});
	dataSimpeg.value = result
	console.log(result)
	return result
});

const { pending: pending_bkn, data: data_bkn, error} = await useLazyAsyncData('getDataKinerjaBkn', async() => {
	console.log("CariData Kinerja Bkn")
	let nip = $decodeBase64(snip)

	var url = "rw-skp22"

	if(jns == 'P'){
		url = "rw-kinerjaperiodik"
	}

	var result = await $fetch(`/api/gets/singkronisasi_bkn/getDataBKN/${url}/${nip}`)
	
	var bkn = _orderBy(result.data, 'tahun', 'desc')
	return bkn
});

const editData = (item) => {
	let sthnilai = item ? $encodeBase64(item.thnilai) : $encodeBase64(null)
	let sid = item ? $encodeBase64(item.id) : $encodeBase64(null)
	navigateTo({ path: `/pegawai/${snip}/RiwayatPegawai/kinerja/${sid}-${sthnilai}-${jns}` })
}

const hapusData = async(item) => {
	var body = JSON.stringify({
		nip: item.nip,
		thnilai: item.thnilai,
		id: item.id,
		idSync: item.idSync,
		jns: item.jns
	});

	var {data, error} = await useFetch('/api/delete/riwayat_kinerja', {
		method: 'DELETE',
		body: body,
	});

	if(error.value){
		toast.add({
	    	id: 'error_delete_riwayat_skp',
	    	icon: "i-heroicons-x-circle",
		    title: "Hapus Data Kinerja",
	    	description: error.value.data.data,
	    	timeout: 6000,
	    	color: 'red',
	  	}) 	
	}else{
		toast.add({
	    	id: 'success_delete_riwayat_skp',
	    	icon: "i-heroicons-check-circle",
		    title: "Hapus Data Kinerja",
	    	description: "Data kinerja berhasil di Hapus",
	    	timeout: 6000
	  	}) 	
	}

	refreshNuxtData(["getDataKinerja", "getDataKinerjaBkn"])
}

async function getDataBkn(item) {
	var dd = _find(dataSimpeg.value, function(o){
		return o.idSync == item.id && o.thnilai == item.tahun
	})

	let nip = $decodeBase64(snip)

	var compacted = _pickBy(item);
	compacted.idSync = compacted.id
	compacted.id = 0
	compacted.nip = nip
	compacted.thnilai = compacted.tahun *1
	compacted.kuadran_kinerja = compacted.KuadranKinerjaNilai *1
	compacted.perilaku_kerja = compacted.PerilakuKerjaNilai *1
	compacted.hasilKinerja = compacted.hasilKinerja
	compacted.hasil_kinerja = compacted.hasilKinerjaNilai *1
	compacted.kuadranKinerja = compacted.kuadranKinerja
	compacted.pej_nama = compacted.namaPenilai
	compacted.pej_nik_nip = compacted.nipNrpPenilai
	compacted.pej_kgolru = compacted.penilaiGolonganId *1
	compacted.pej_njab = compacted.penilaiJabatanNm
	compacted.pej_nunker = compacted.penilaiUnorNm
	compacted.perilakuKerja = compacted.perilakuKerja
	compacted.pej_pns = compacted.statusPenilai == "ASN" ? 1 : 0
	compacted.jns = jns
	
	if(dd){
		console.log("ketemu")
		var body = {
			"refdata" : {
				id: dd.id,
				nip: dd.nip,
				thnilai: dd.thnilai,
			},
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_kinerja', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_kinerja',
		    	icon: "i-heroicons-x-circle",
		    	title: "Tambah dari Kinerja Data BKN",
		    	description: error.value.data.data,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_kinerja',
		    	icon: "i-heroicons-check-circle",
		    	title: "Tambah dari Kinerj Data BKN",
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
	}else{
		console.log("tidak ketemu")
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_kinerja/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_skp',
		    	icon: "i-heroicons-x-circle",
		    	title: "Tambah dari Kinerja Data BKN",
		    	description: error.value.data.data,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_skp',
		    	icon: "i-heroicons-check-circle",
		    	title: "Tambah dari Kinerja Data BKN",
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
		}
	}

	refreshNuxtData(["getDataKinerja", "getDataKinerjaBkn"])
}

function getRatings(idx){
	var result = _find(ratings.value, {id: idx})
	return result.nama
}

function getKuadrankinerjas(idx){
	var result = _find(kuadrankinerjas.value, {id: idx})
	return result.nama
}

function checkPath(path){
	if(path != null){
		if(path["891"] != undefined){
			return true
		}
		//console.log(path["872"])
		
	}
	return false
}

async function downloadDoc(item){
	showSpinner.value = true

	let filepath = _values(item.path)[0]
	let nip = $decodeBase64(snip)

	var {data, error} = await useFetch('/api/posts/download/dokRwBkn', {
		method: 'POST',
		body: JSON.stringify({
			path : filepath["dok_uri"],
			nip: nip,
		}),
	})

	if(error.value){
		toast.add({
	    	id: 'error_post_download_kinerja',
	    	icon: "i-heroicons-x-circle",
	    	title: "Gagal Mendownload",
	    	description: "Gagal Mendownload data BKN",
	    	timeout: 6000,
	    	color: 'red',
	  	}) 	
	}else{
		console.log(data.value)
		const {filename} = data.value
		/*toast.add({
	    	id: 'success_post_download_kinerja',
	    	icon: "i-heroicons-check-circle",
	    	title: "Berhasil download dokumen BKN",
	    	description: "Silahkan upload kembali dokumen tersebut",
	    	timeout: 6000,
	  	})*/ 	

	  	//update simpeg

		var dd = _find(dataSimpeg.value, function(o){
			return o.idSync == item.id && o.thnilai == item.tahun
		})

		if(dd){
			var body = {
				"refdata" : {
					id: dd.id,
					nip: dd.nip,
					thnilai: dd.thnilai,
				},
				"data" : {
					nip: dd.nip,
					thnilai: dd.thnilai,
					filename: filename,
				}
			}

			var {data, error} = await useFetch('/api/puts/riwayat_kinerja', {
				method: 'PUT',
				body: JSON.stringify(body),
			})

			if(error.value){
				toast.add({
			    	id: 'error_put_kinerja',
			    	icon: "i-heroicons-x-circle",
			    	title: "Download dokumen BKN",
			    	description: error.value.data.data,
			    	timeout: 6000,
			    	color: 'red',
			  	}) 	
			}else{
				toast.add({
			    	id: 'success_put_kinerja',
			    	icon: "i-heroicons-check-circle",
			    	title: "Download dokumen BKN",
			    	description: "Data Simpeg berhasil di Update",
			    	timeout: 6000
			  	}) 	
			}	
		}

	  	//end update simpeg

	  	/*var fwd = await $fetch(`/api/fileserver/static/dokumen/temp/${data.value.filename}`)

	  	console.log(data.value.filename, fwd)
	  	console.log(`/api/fileserver/static/temp/${data.value.filename}`)

	  	const blob = new Blob([fwd], {type: 'application/pdf'})

		const blobUrl = URL.createObjectURL(blob)

		const config = useRuntimeConfig()
	  	const link = document.createElement('a')
	    link.href = blobUrl

	    link.download = data.value.filename
	    link.target = '_blank'
	    link.click()*/
	}
	
	showSpinner.value = false
}
</script>
<template>
	<AppLoadingSpinner :show="showSpinner" />
	<LayoutRiwayatPegawai>
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="flex justify-end mb-2">
					<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="editData()">Tambah</button>
				</div>
				<div class="mb-8">
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
										<th scope="col" class="px-3 text-center"></th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Tahun
											</span>
										</th>
										<template v-if="jns=='P'">
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Priode Penilaian
												</span>
											</th>		
										</template>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Hasil Kinerja
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Perilaku Kerja
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Kuadran Kinerja
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
									<template v-else-if="data" v-for="(item, idx) in data" :key="idx">
										<tr class="odd:bg-white even:bg-gray-100">
											<td class="size-px px-1" style="vertical-align: top; width: 5%">
												<span class="flex justify-center text-xs px-3 py-0">{{idx+1}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; width: 5%">
												<div class="flex justify-center">
													<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="editData(item)">
													  	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="flex-shrink-0 size-3">
														  	<path d="M13.488 2.513a1.75 1.75 0 0 0-2.475 0L6.75 6.774a2.75 2.75 0 0 0-.596.892l-.848 2.047a.75.75 0 0 0 .98.98l2.047-.848a2.75 2.75 0 0 0 .892-.596l4.261-4.262a1.75 1.75 0 0 0 0-2.474Z" />
														  	<path d="M4.75 3.5c-.69 0-1.25.56-1.25 1.25v6.5c0 .69.56 1.25 1.25 1.25h6.5c.69 0 1.25-.56 1.25-1.25V9A.75.75 0 0 1 14 9v2.25A2.75 2.75 0 0 1 11.25 14h-6.5A2.75 2.75 0 0 1 2 11.25v-6.5A2.75 2.75 0 0 1 4.75 2H7a.75.75 0 0 1 0 1.5H4.75Z" />
														</svg>
													</button>
												</div>
											</td>
											<td class="size-px px-1" style="vertical-align: top;">
												<span class="flex justify-center text-center text-xs py-0">{{item.thnilai}}</span>
											</td>
											<template v-if="jns=='P'">
												<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
													<span class="flex justify-center text-center text-xs py-0">{{item.tmulai ? $dayjs(item.tmulai).format("DD-MM-YYYY").toString() : '-'}} <span>&nbsp;&nbsp;&nbsp;s/d&nbsp;&nbsp;&nbsp;</span> {{item.tselesai ? $dayjs(item.tselesai).format("DD-MM-YYYY").toString() : '-'}}</span>
												</td>
											</template>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.hasil_kinerja}}-{{getRatings(item.hasil_kinerja)}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.perilaku_kerja}}-{{getRatings(item.perilaku_kerja)}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.kuadran_kinerja}}-{{getKuadrankinerjas(item.perilaku_kerja)}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top;">
												<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-red-600 text-white hover:bg-red-700 disabled:opacity-50 disabled:pointer-events-none" @click="hapusData(item)">
												  	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="flex-shrink-0 size-3">
													  	<path fill-rule="evenodd" d="M5 3.25V4H2.75a.75.75 0 0 0 0 1.5h.3l.815 8.15A1.5 1.5 0 0 0 5.357 15h5.285a1.5 1.5 0 0 0 1.493-1.35l.815-8.15h.3a.75.75 0 0 0 0-1.5H11v-.75A2.25 2.25 0 0 0 8.75 1h-1.5A2.25 2.25 0 0 0 5 3.25Zm2.25-.75a.75.75 0 0 0-.75.75V4h3v-.75a.75.75 0 0 0-.75-.75h-1.5ZM6.05 6a.75.75 0 0 1 .787.713l.275 5.5a.75.75 0 0 1-1.498.075l-.275-5.5A.75.75 0 0 1 6.05 6Zm3.9 0a.75.75 0 0 1 .712.787l-.275 5.5a.75.75 0 0 1-1.498-.075l.275-5.5a.75.75 0 0 1 .786-.711Z" clip-rule="evenodd" />
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
										<th scope="col" class="px-3 text-center"></th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Tahun
											</span>
										</th>
										<template v-if="jns=='P'">
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Priode Penilaian
												</span>
											</th>		
										</template>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Hasil Kinerja
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Perilaku Kerja
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Kuadran Kinerja
											</span>
										</th>
										<th scope="col" class="px-3 text-center"></th>
									</tr>
								</thead>
								<tbody class="divide-y divide-gray-200">
									<template v-if="pending_bkn">
										<tr>
											<td colspan="7" class="size-px p-1" style="vertical-align: top; min-width: 100px;">
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
									<template v-else-if="data_bkn" v-for="(item, idx) in data_bkn" :key="idx">
										<tr class="odd:bg-white even:bg-gray-100">
											<td class="size-px px-1" style="vertical-align: top; width: 5%">
												<span class="flex justify-center text-xs px-3 py-0">{{idx+1}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; width: 5%">
												<div class="flex justify-center">
													<template v-if="checkPath(item.path)">
														<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="downloadDoc(item)">
															<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4">
																<path fill-rule="evenodd" d="M2 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4Zm10.5 5.707a.5.5 0 0 0-.146-.353l-1-1a.5.5 0 0 0-.708 0L9.354 9.646a.5.5 0 0 1-.708 0L6.354 7.354a.5.5 0 0 0-.708 0l-2 2a.5.5 0 0 0-.146.353V12a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V9.707ZM12 5a1 1 0 1 1-2 0 1 1 0 0 1 2 0Z" clip-rule="evenodd" />
															</svg>
														</button>
													</template>
												</div>
											</td>
											<td class="size-px px-1" style="vertical-align: top;">
												<span class="flex justify-center text-center text-xs py-0">{{item.tahun}}</span>
											</td>
											<template v-if="jns=='P'">
												<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
													<span class="flex justify-center text-center text-xs py-0">1-{{item.bulanMulaiPenilaian}}-{{item.tahunMulaiPenilaian}} <span>&nbsp;&nbsp;&nbsp;s/d&nbsp;&nbsp;&nbsp;</span>1-{{item.bulanSelesaiPenilaian}}-{{item.tahunSelesaiPenilaian}}</span>
												</td>
											</template>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{parseInt(item.hasilKinerjaNilai).toFixed(0)}}-{{item.hasilKinerja}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{parseInt(item.perilakuKerjaNilai).toFixed(0)}}-{{item.perilakuKerja}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{parseInt(item.kuadranKinerjaNilai).toFixed(0)}}-{{item.kuadranNilai}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top;">
												<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-green-600 text-white hover:bg-green-700 disabled:opacity-50 disabled:pointer-events-none" @click="getDataBkn(item)">
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