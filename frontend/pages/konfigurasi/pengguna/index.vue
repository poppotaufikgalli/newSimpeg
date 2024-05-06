<script setup>
import { useModalStore } from '~/store/modalStore';

const toast = useToast()
const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const modalStore = useModalStore();
const selId = ref(0)

const { data, pending, error, refresh } = await useFetch('/api/gets/users')

const doData = (id) => {
	const sid = $encodeBase64(id)
	navigateTo({ path: '/konfigurasi/pengguna/'+sid })
}

const doDelete = (item) => {
	selId.value = item.id;
	modalStore.showModal("Hapus Data Pengguna", "Apakah anda yakin untuk menghapus data Pengguna an. "+item.name+" ?")
}

onMounted(() => {
	modalStore.$onAction(callback, true)
})

const callback = async(e) => {
	const {name} = e
	
	if(name == 'doAction'){
		await useFetch('/api/delete/users', {
			method: 'delete',
			body: JSON.stringify({
				id: selId.value.toString()
			})
		})

		refresh()

		toast.add({
	    	id: 'success_delete_nip',
	    	description: "Berhasil menghapus Data",
	    	timeout: 6000
	  	})
	}
}

</script>
<template>
	<LayoutKonfigurasi>
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="flex justify-end mb-2">
					<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doData(0)">Tambah</button>
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
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												NIP
											</span>
										</th>
										<th scope="col" class="px-3 text-center w-[30%]">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Nama
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Group ID
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Status
											</span>
										</th>
										<th scope="col" class="px-3 text-center"></th>
									</tr>
								</thead>
								<template v-if="pending">
									<td colspan="4" align="center">
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
												<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
													<span class="flex justify-center text-xs py-0">{{item.nip}}</span>
												</td>
												<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
													<span class="flex justify-center text-xs py-0">{{item.name}}</span>
												</td>
												<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
													<span class="flex justify-center text-xs py-0">{{item.gid == 1 ? 'Admin' : 'User Biasa'}}</span>
												</td>
												<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
													<span class="flex justify-center text-xs py-0">{{item.stts == 1 ? 'Aktif' : ''}}</span>
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
	</LayoutKonfigurasi>
</template>