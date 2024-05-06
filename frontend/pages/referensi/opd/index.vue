<script setup>
import { useModalStore } from '~/store/modalStore';

const toast = useToast()
const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const modalStore = useModalStore();
const selId = ref(0)
const statuss = ref([
	{id:0, nama: 'Tidak Aktif'},
	{id:1, nama: 'Aktif'},
])
const s_status = ref([1,0])

const { data, pending, error, refresh } = await useFetch('/api/posts/opd', {
	method: "POST",
	body: JSON.stringify({
		status: s_status.value,
	})
})

const refOpdInduk = ref(null)
const s_opdInduk = ref("1")

const opdInduks = computed({
	get(){
		var result = _filter(data.value, {sfilter : 1})
		return result
	}
})

const opdSemua = computed({
	get(){
		var result = _filter(data.value, function(o){
			return _startsWith(o.id, s_opdInduk.value)
		})
		return result
	}
})

const doAdd = (id) => {
	const sid = $encodeBase64(null)
	navigateTo({ path: '/referensi/opd/'+sid, query: {parent: id} })
}

const doData = (id) => {
	const sid = $encodeBase64(id)
	navigateTo({ path: '/referensi/opd/'+sid })
}

const doDelete = (item) => {
	selId.value = item.id;
	modalStore.showModal("Hapus Data OPD", "Apakah anda yakin untuk menghapus data OPD an. "+item.nama+" ?")
}

onMounted(() => {
	modalStore.$onAction(callback, true)
})

const callback = async(e) => {
	const {name} = e
	
	if(name == 'doAction'){
		await useFetch('/api/delete/opd', {
			method: 'delete',
			body: JSON.stringify({
				id: selId.value.toString()
			})
		})

		refresh()

		toast.add({
	    	id: 'success_delete_opd',
	    	description: "Berhasil menghapus Data",
	    	timeout: 6000
	  	})
	}
}

</script>
<template>
	<LayoutReferensi>
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="flex justify-end mb-2">
					<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doData(0)">Tambah</button>
				</div>
				<div class="grid sm:grid-cols-12 gap-2 gap-2.5 mb-2">
					<div class="sm:col-span-2">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Pilih OPD Induk</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-10">
						<SearchSelect2 ref="refOpdInduk" id="opdInduks" :options="opdInduks" keyList="id" namaList="nama" v-model="s_opdInduk" />
					</div>
					<!-- End Col -->
				</div>
				
				<div class="mb-8">
					<div class="overflow-x-auto border">
						<div class="min-w-full inline-block align-middle">
							<!-- Table -->
							<table class="min-w-full divide-y divide-gray-200">
								<thead class="bg-gray-50 h-12">
									<tr>
										<th scope="col" class="px-3 text-center w-[5%]">
											<span class="text-xs font-semibold uppercase text-gray-800">
												No
											</span>
										</th>
										<th scope="col" class="px-3 text-center w-[25%]">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Nama
											</span>
										</th>
										<th scope="col" class="px-3 text-center w-[5%]">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Eselon
											</span>
										</th>
										<th scope="col" class="px-3 text-center w-[25%]">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Nama Jabatan
											</span>
										</th>
										<th scope="col" class="px-3 text-center w-[25%]">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Unor BKN
											</span>
										</th>
										<th scope="col" class="px-3 text-center w-[5%]"></th>
									</tr>
								</thead>
								<template v-if="pending">
									<td colspan="4" align="center">
										<i>Loading</i>
									</td>
								</template>
								<template v-else>
									<tbody class="divide-y divide-gray-200">
										<template v-if="!pending" v-for="(item, idx) in opdSemua" :key="idx">
											<tr class="odd:bg-white even:bg-gray-100" :class="item.status == 0  ? 'italic' : '' ">
												<td class="size-px px-1" style="vertical-align: top;">
													<button type="button" class="block" data-hs-overlay="#hs-ai-invoice-modal">
														<span class="flex justify-center text-xs px-3 py-0">{{idx+1}}</span>
													</button>
												</td>
												<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
													<span class="flex justify-start text-xs py-0">{{item.nama}}</span>
												</td>
												<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
													<span class="flex justify-center text-xs py-0">{{item.eselon?.nama}}</span>
												</td>
												<td class="size-px px-1 text-center" style="vertical-align: top; min-width: 100px;">
													<span class="flex justify-center text-xs py-0">{{item.nama_jabatan}}</span>
												</td>
												<td class="size-px px-1" style="vertical-align: top">
													<span class="flex justify-center text-xs py-0">{{item.id_unor_bkn}}</span>
												</td>
												<td class="size-px px-1" style="vertical-align: top;">
													<div class="sm:flex gap-x-1">
														<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doData(item.id)">
														  	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="flex-shrink-0 size-3">
															  	<path d="M13.488 2.513a1.75 1.75 0 0 0-2.475 0L6.75 6.774a2.75 2.75 0 0 0-.596.892l-.848 2.047a.75.75 0 0 0 .98.98l2.047-.848a2.75 2.75 0 0 0 .892-.596l4.261-4.262a1.75 1.75 0 0 0 0-2.474Z" />
															  	<path d="M4.75 3.5c-.69 0-1.25.56-1.25 1.25v6.5c0 .69.56 1.25 1.25 1.25h6.5c.69 0 1.25-.56 1.25-1.25V9A.75.75 0 0 1 14 9v2.25A2.75 2.75 0 0 1 11.25 14h-6.5A2.75 2.75 0 0 1 2 11.25v-6.5A2.75 2.75 0 0 1 4.75 2H7a.75.75 0 0 1 0 1.5H4.75Z" />
															</svg>
														</button>
														<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-red-600 text-white hover:bg-red-700 disabled:opacity-50 disabled:pointer-events-none" @click="doDelete(item)">
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
	</LayoutReferensi>
</template>