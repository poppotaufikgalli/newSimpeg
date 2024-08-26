<script setup>
const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
import customParseFormat from "dayjs/plugin/customParseFormat";
dayjs.extend(customParseFormat)

const toast = useToast()
const {snip} = route.params
const dataSimpeg = ref([])

onMounted(() => {
	refreshNuxtData(["getDataDp3"])
})

const { pending, data, refresh} = await useLazyAsyncData('getDataSkp', async() => {
	console.log("CariData")

	let nip = $decodeBase64(snip)
	var body = JSON.stringify({
		nip: nip,
	});

	var result = await $fetch('/api/posts/riwayat_skp', {
		method: 'POST',
		body: body,
	});

	dataSimpeg.value = result
	return result
});

const { pending: pending_bkn, data: data_bkn, error} = await useLazyAsyncData('getDataSkpBkn', async() => {
	console.log("CariData Skp Bkn")
	let nip = $decodeBase64(snip)

	var result = await $fetch(`/api/gets/singkronisasi_bkn/getDataBKN/rw-skp/${nip}`)
	
	var bkn = _orderBy(result.data, 'tahun', 'desc')
	return bkn
});

const editData = (item) => {
	let sid = item ? $encodeBase64(item.id) : $encodeBase64(null)
	let stahun = item ? $encodeBase64(item.tahun) : $encodeBase64(null)
	navigateTo({ path: `/pegawai/${snip}/RiwayatPegawai/skp/${sid}-${stahun}` })
}

const hapusData = async(item) => {
	var body = JSON.stringify({
		nip: item.nip,
		tahun: item.tahun,
		id: item.id,
	});

	var {data, error} = await useFetch('/api/delete/riwayat_skp', {
		method: 'DELETE',
		body: body,
	});

	if(error.value){
		toast.add({
	    	id: 'error_delete_riwayat_skp',
	    	description: error.value.data.data,
	    	timeout: 6000
	  	}) 	
	}else{
		toast.add({
	    	id: 'success_delete_riwayat_skp',
	    	description: "Data SKP Kerja berhasil di Hapus",
	    	timeout: 6000
	  	}) 	
	}

	refreshNuxtData(["getDataSkp", "getDataSkpBkn"])
}

async function getDataBkn(item) {
	var dd = _find(dataSimpeg.value, function(o){
		return o.idSync == item.id && o.tahun == item.tahun
	})

	console.log(dd)

	let nip = $decodeBase64(snip)

	var compacted = _pickBy(item);
	compacted.idSync = compacted.id
	compacted.id = 0
	compacted.nip = nip
	compacted.nilaiSkp = compacted.nilaiSkp *1
	compacted.orientasiPelayanan = compacted.orientasiPelayanan *1
	compacted.komitmen = compacted.komitmen *1
	compacted.disiplin = compacted.disiplin *1
	compacted.kerjasama = compacted.kerjasama *1
	compacted.nilaiSkp = compacted.nilaiSkp *1
	compacted.integritas = compacted.integritas *1
	compacted.konversiNilai = compacted.konversiNilai *1
	compacted.kepemimpinan = compacted.kepemimpinan *1
	compacted.inisiatifKerja = compacted.inisiatifKerja *1
	compacted.kepemimpinan = compacted.kepemimpinan *1
	compacted.jumlah = compacted.jumlah *1
	compacted.nilaiIntegrasi = compacted.nilaiIntegrasi *1
	compacted.nilaiPerilakuKerja = compacted.nilaiPerilakuKerja *1
	compacted.nilaiPrestasiKerja = compacted.nilaiPrestasiKerja *1
	compacted.nilairatarata = compacted.nilairatarata *1
	compacted.tahun = compacted.tahun *1

	//console.log(dayjs(compacted.penilaiTmtGolongan, 'DD-MM-YYYY', true).isValid())
	
	compacted.atasanPenilaiTmtGolongan = dayjs(compacted.atasanPenilaiTmtGolongan, 'DD-MM-YYYY', true).isValid() ? dayjs(compacted.atasanPenilaiTmtGolongan, "DD-MM-YYYY").format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.penilaiTmtGolongan = dayjs(compacted.penilaiTmtGolongan, 'DD-MM-YYYY', true).isValid() ? dayjs(compacted.penilaiTmtGolongan, "DD-MM-YYYY").format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null

	if(dd){
		console.log("ketemu")
		var body = {
			"refdata" : {
				id: dd.id,
				nip: dd.nip,
				tahun: dd.tahun,
			},
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_skp', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_kinerja',
		    	icon: "i-heroicons-x-circle",
		    	title: "Tambah dari SKP Data BKN",
		    	description: error.value.data.data,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_kinerja',
		    	icon: "i-heroicons-check-circle",
		    	title: "Tambah dari SKP Data BKN",
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
	}else{
		console.log("tidak ketemu")
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_skp/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_skp',
		    	icon: "i-heroicons-x-circle",
		    	title: "Tambah dari SKP Data BKN",
		    	description: error.value.data.data,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_skp',
		    	icon: "i-heroicons-check-circle",
		    	title: "Tambah dari SKP Data BKN",
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
		}
	}

	refreshNuxtData(["getDataSkp", "getDataSkpBkn"])
}

function checkPath(path){
	if(path != null){
		if(path["872"] != undefined){
			return true
		}
		//console.log(path["872"])
		
	}
	return false
}
</script>
<template>
	<LayoutRiwayatPegawai>
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="flex justify-end mb-2 gap-2">
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
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Jenis Jabatan
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Nilai SKP
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Orientasi Pelayanan
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Komitmen
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Integritas
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Disiplin
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Kerjasama
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Aturan
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
											<td class="size-px px-1" style="vertical-align: top; width: 5%;">
												<span class="flex justify-center text-xs px-3 py-0">{{idx+1}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; width: 5%;">
												<div class="sm:flex justify-center">
													<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="editData(item)">
													  	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="flex-shrink-0 size-3">
														  	<path d="M13.488 2.513a1.75 1.75 0 0 0-2.475 0L6.75 6.774a2.75 2.75 0 0 0-.596.892l-.848 2.047a.75.75 0 0 0 .98.98l2.047-.848a2.75 2.75 0 0 0 .892-.596l4.261-4.262a1.75 1.75 0 0 0 0-2.474Z" />
														  	<path d="M4.75 3.5c-.69 0-1.25.56-1.25 1.25v6.5c0 .69.56 1.25 1.25 1.25h6.5c.69 0 1.25-.56 1.25-1.25V9A.75.75 0 0 1 14 9v2.25A2.75 2.75 0 0 1 11.25 14h-6.5A2.75 2.75 0 0 1 2 11.25v-6.5A2.75 2.75 0 0 1 4.75 2H7a.75.75 0 0 1 0 1.5H4.75Z" />
														</svg>
													</button>
												</div>
											</td>
											<td class="size-px px-1" style="vertical-align: top;">
												<button type="button" class="block" data-hs-overlay="#hs-ai-invoice-modal">
													<span class="flex justify-center text-xs px-3 py-0">{{item.tahun}}</span>
												</button>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.jenisJabatan}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.nilaiSkp}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.orientasiPelayanan}}
												</span>
											</td>

											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.komitmen}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.integritas}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.disiplin}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.kerjasama}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.jenisPeraturanKinerjaKd}}
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
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Jenis Jabatan
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Nilai SKP
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Orientasi Pelayanan
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Komitmen
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Integritas
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Disiplin
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Kerjasama
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Aturan
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
												<span class="flex justify-center text-xs py-1 text-red-500">
													Get Data BKN Gagal :<b> {{error.data.data.message || error.data.data.errors}}</b>
												</span>
											</td>
										</tr>
									</template>
									<template v-else-if="data_bkn" v-for="(item, idx) in data_bkn" :key="idx">
										<tr class="odd:bg-white even:bg-gray-100">
											<td class="size-px px-1" style="vertical-align: top; width: 5%">
												<span class="flex justify-center text-xs px-3 py-0">{{idx+1}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; width: 5%">
												<div class="sm:flex justify-center">
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
												<span class="flex justify-center text-xs px-3 py-0">{{item.tahun}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.jenisJabatan}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.nilaiSkp}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.orientasiPelayanan}}
												</span>
											</td>

											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.komitmen}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.integritas}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.disiplin}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.kerjasama}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{item.jenisPeraturanKinerjaKd}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top;">
												<div class="sm:flex gap-x-1">
													<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-green-600 text-white hover:bg-green-700 disabled:opacity-50 disabled:pointer-events-none" @click="getDataBkn(item)">
													  	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="flex-shrink-0 size-3">
															<path fill-rule="evenodd" d="M4.5 13a3.5 3.5 0 0 1-1.41-6.705A3.5 3.5 0 0 1 9.72 4.124a2.5 2.5 0 0 1 3.197 3.018A3.001 3.001 0 0 1 12 13H4.5Zm.72-5.03a.75.75 0 0 0 1.06 1.06l.97-.97v2.69a.75.75 0 0 0 1.5 0V8.06l.97.97a.75.75 0 1 0 1.06-1.06L8.53 5.72a.75.75 0 0 0-1.06 0L5.22 7.97Z" clip-rule="evenodd" />
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
			</div>
		</div>
	</LayoutRiwayatPegawai>
</template>