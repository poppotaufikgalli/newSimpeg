<script setup>
const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const toast = useToast()
const {snip, jdiklat} = route.params

onMounted(() => {
	refreshNuxtData(["getData"])
})

const { pending, data, refresh} = await useLazyAsyncData('getData', async() => {
	console.log("CariData")

	let nip = $decodeBase64(snip);
	var body = JSON.stringify({
		nip: nip,
		jdiklat: jdiklat *1,
	});

	return await $fetch('/api/posts/riwayat_diklat', {
		method: 'POST',
		body: body,
	});
});

const editData = (item) => {
	let stmulai = item ? $encodeBase64(item.tmulai) : $encodeBase64(null)
	navigateTo({ path: `/pegawai/${snip}/RiwayatPendidikan/diklat/${jdiklat}/${stmulai}` })
}

const hapusData = async(item) => {
	var body = JSON.stringify({
		nip: item.nip,
		jdiklat: jdiklat*1,
		tmulai: item.tmulai,
	});

	var {data, error} = await useFetch('/api/delete/riwayat_diklat', {
		method: 'DELETE',
		body: body,
	});

	if(error.value){
		toast.add({
	    	id: 'error_delete_riwayat_diklat',
	    	description: error.value.data.data,
	    	timeout: 6000
	  	}) 	
	}else{
		toast.add({
	    	id: 'success_delete_riwayat_diklat',
	    	description: "Data Diklat berhasil di Hapus",
	    	timeout: 6000
	  	}) 	
	}

	refreshNuxtData(["getData"])
}
</script>
<template>
	<LayoutRiwayatPendidikan>
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="flex justify-end mb-2">
					<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="editData()">Tambah {{id}}</button>
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
										<template v-if="jdiklat == 8">
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Nama Seminar
												</span>
											</th>
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Peran Sebagai
												</span>
											</th>
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Tempat
												</span>
											</th>
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Waktu Pelaksanaan
												</span>
											</th>
										</template>
										<template v-else-if="jdiklat == 9">
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Judul Buku/Karya Tulis/Makalah
												</span>
											</th>
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Tempat Pembuatan/Terbit
												</span>
											</th>
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Tgl Pembuatan
												</span>
											</th>
										</template>
										<template v-else-if="jdiklat == 10">
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Nama
												</span>
											</th>
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Tempat
												</span>
											</th>
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Penyelenggara
												</span>
											</th>
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Tanggal Ujian
												</span>
											</th>
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Berlaku Sampai
												</span>
											</th>
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Status
												</span>
											</th>
										</template>
										<template v-else>
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Nama
												</span>
											</th>
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Tempat
												</span>
											</th>
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Panitia
												</span>
											</th>
											<th scope="col" class="px-3 text-center">
												<span class="text-xs font-semibold uppercase text-gray-800">
													Waktu Pelaksanaan
												</span>
											</th>
											<template v-if="jdiklat != 7">
												<th scope="col" class="px-3 text-center">
													<span class="text-xs font-semibold uppercase text-gray-800">
														Akhir
													</span>
												</th>
											</template>
										</template>
										
										<th scope="col" class="px-3 text-center"></th>
									</tr>
								</thead>
								<template v-if="pending">
									<td colspan="7" align="center">
										<i>Loading</i>
									</td>
								</template>
								<template v-else>
									<tbody class="divide-y divide-gray-200">
										<template v-if="data" v-for="(item, idx) in data" :key="idx">
											<tr class="odd:bg-white even:bg-gray-100">
												<td class="size-px px-1" style="vertical-align: top;">
													<button type="button" class="block" data-hs-overlay="#hs-ai-invoice-modal">
														<span class="flex justify-center text-xs px-3 py-0">{{idx+1}}</span>
													</button>
												</td>
												<template v-if="jdiklat == 8">
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-start text-xs py-0">{{item.ndiklat}}</span>
													</td>
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-start text-xs py-0">{{item.pan}}</span>
													</td>
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-start text-xs py-0">{{item.tempat}}</span>
													</td>
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-center gap-x-1 text-xs py-0">{{item.tmulai ? $dayjs(item.tmulai).format("DD-MM-YYYY").toString() : '-'}}<span>s/d</span>{{item.takhir ? $dayjs(item.takhir).format("DD-MM-YYYY").toString() : '-'}}</span>
													</td>
												</template>
												<template v-else-if="jdiklat == 9">
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-start text-xs py-0">{{item.ndiklat}}</span>
													</td>
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-start text-xs py-0">{{item.tempat}}</span>
													</td>
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-center gap-x-1 text-xs py-0">{{item.tmulai ? $dayjs(item.tmulai).format("DD-MM-YYYY").toString() : '-'}}</span>
													</td>
												</template>
												<template v-else-if="jdiklat == 10">
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-start text-xs py-0">{{item.ndiklat}}</span>
													</td>
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-start text-xs py-0">{{item.tempat}}</span>
													</td>
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-start text-xs py-0">{{item.pan}}</span>
													</td>
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-center gap-x-1 text-xs py-0">{{item.tmulai ? $dayjs(item.tmulai).format("DD-MM-YYYY").toString() : '-'}}</span>
													</td>
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-center gap-x-1 text-xs py-0">{{item.takhir ? $dayjs(item.takhir).format("DD-MM-YYYY").toString() : '-'}}</span>
													</td>
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-center text-xs py-0" :class="$dayjs().isBefore(item.takhir, 'day') || $dayjs().isSame(item.takhir, 'day') ? 'bg-green-300' : 'bg-red-300'">
															{{$dayjs().isBefore(item.takhir, 'day') || $dayjs().isSame(item.takhir, 'day') ? 'Berlaku' : '-'}}
														</span>
													</td>
												</template>
												<template v-else>
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-start text-xs py-0">{{item.ndiklat}}</span>
													</td>
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-start text-xs py-0">{{item.tempat}}</span>
													</td>
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-start text-xs py-0">{{item.pan}}</span>
													</td>
													<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
														<span class="flex justify-center gap-x-1 text-xs py-0">{{item.tmulai ? $dayjs(item.tmulai).format("DD-MM-YYYY").toString() : '-'}}<span>s/d</span>{{item.takhir ? $dayjs(item.takhir).format("DD-MM-YYYY").toString() : '-'}}</span>
													</td>
													<template v-if="jdiklat != 7">
														<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
															<span class="flex justify-center text-xs py-0">{{item.akhir == 1 ? 'Ya' : ''}}</span>
														</td>
													</template>
												</template>
												<td class="size-px px-1" style="vertical-align: top;">
													<div class="sm:flex gap-x-1">
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
													</div>
												</td>
											</tr>
										</template>
									</tbody>
								</template>
								
							</table>
							<!-- End Table -->	
						</div>
					</div>
				</div>
			</div>
		</div>
	</LayoutRiwayatPendidikan>
</template>