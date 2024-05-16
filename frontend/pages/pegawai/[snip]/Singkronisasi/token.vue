<script setup>
const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const dayjs = useDayjs()
const toast = useToast()
const route = useRoute()

const {snip} = route.params
const showSpinner = ref(false)
const dataSyncs = ref([])
const dataRiwayats = ref([])


const { pending, data, refresh } = await useAsyncData('getDataTokenInfo', async () => {
	console.log("getDataTokenInfo")
	
	let nip = $decodeBase64(snip)
  	const [dataSync, dataRiwayat] = await Promise.all([
    	$fetch('/api/gets/singkronisasi/bkn'),
    	$fetch('/api/posts/singkronisasi_bkn/riwayat', {
    		method: 'POST',
    		body: JSON.stringify({
    			nip: nip
    		})
    	}),
  	])

  	dataSyncs.value = dataSync
  	dataRiwayats.value = _orderBy(dataRiwayat, 'created_at', 'desc');
});

async function updateToken(type){
	showSpinner.value = true
	var result = await $fetch('/api/posts/singkronisasi_bkn/updateToken', {
		method: 'POST',
		body: JSON.stringify({
			type: type,
			ckey: dataSyncs.value.ckey,
			csecret: dataSyncs.value.csecret,
			client_id: dataSyncs.value.client_id,
			username: dataSyncs.value.username,
			password: dataSyncs.value.password,
		})
	})
	
	//const {token_sso_expired, token_apimanager_expired} = result
	if(type == 'token_sso'){
		dataSyncs.value.token_sso_expired = dayjs(result.token_sso_expired).format('YYYY-MM-DDTHH:mm:ssZ')
	}

	if(type == 'token_apimanager'){
		dataSyncs.value.token_apimanager_expired = dayjs(result.token_apimanager_expired).format('YYYY-MM-DDTHH:mm:ssZ')
	}
	showSpinner.value = false
}

onMounted(() => {
	refreshNuxtData(["getDataTokenInfo"])
})


</script>
<template>
	<AppLoadingSpinner :show="showSpinner" />
	<LayoutSingkronisasi>
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2 mb-8">
				<!-- Grid -->
				<div class="grid sm:grid-cols-12 gap-2 gap-2.5" v-if="!pending">
					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Token SSO Expired</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 disabled:opacity-50 disabled:pointer-events-none bg-white" :value="dataSyncs.token_sso_expired" v-if="dataSyncs" :class="$dayjs().isBefore($dayjs(dataSyncs.token_sso_expired)) ? 'border-gray-200 focus:border-blue-500 focus:ring-blue-500' : 'border-red-500 focus:border-red-500 focus:ring-red-500'" disabled>
							<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[36px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" :disabled="$dayjs().isBefore($dayjs(dataSyncs.token_sso_expired))"  @click="updateToken('token_sso')">
							  	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-5 h-5">
									<path fill-rule="evenodd" d="M15.312 11.424a5.5 5.5 0 0 1-9.201 2.466l-.312-.311h2.433a.75.75 0 0 0 0-1.5H3.989a.75.75 0 0 0-.75.75v4.242a.75.75 0 0 0 1.5 0v-2.43l.31.31a7 7 0 0 0 11.712-3.138.75.75 0 0 0-1.449-.39Zm1.23-3.723a.75.75 0 0 0 .219-.53V2.929a.75.75 0 0 0-1.5 0V5.36l-.31-.31A7 7 0 0 0 3.239 8.188a.75.75 0 1 0 1.448.389A5.5 5.5 0 0 1 13.89 6.11l.311.31h-2.432a.75.75 0 0 0 0 1.5h4.243a.75.75 0 0 0 .53-.219Z" clip-rule="evenodd" />
								</svg>
							</button>
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">API Manager Expired</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" :value="dataSyncs.token_apimanager_expired" v-if="dataSyncs" :class="$dayjs().isBefore($dayjs(dataSyncs.token_apimanager_expired)) ? 'border-gray-200 focus:border-blue-500 focus:ring-blue-500' : 'border-red-500 focus:border-red-500 focus:ring-red-500'" disabled>
							<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[36px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" :disabled="$dayjs().isBefore($dayjs(dataSyncs.token_apimanager_expired))" @click="updateToken('token_apimanager')">
							  	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-5 h-5">
									<path fill-rule="evenodd" d="M15.312 11.424a5.5 5.5 0 0 1-9.201 2.466l-.312-.311h2.433a.75.75 0 0 0 0-1.5H3.989a.75.75 0 0 0-.75.75v4.242a.75.75 0 0 0 1.5 0v-2.43l.31.31a7 7 0 0 0 11.712-3.138.75.75 0 0 0-1.449-.39Zm1.23-3.723a.75.75 0 0 0 .219-.53V2.929a.75.75 0 0 0-1.5 0V5.36l-.31-.31A7 7 0 0 0 3.239 8.188a.75.75 0 1 0 1.448.389A5.5 5.5 0 0 1 13.89 6.11l.311.31h-2.432a.75.75 0 0 0 0 1.5h4.243a.75.75 0 0 0 .53-.219Z" clip-rule="evenodd" />
								</svg>
							</button>
						</div>
					</div>
					<!-- End Col -->
				</div>
			</div>
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="flex justify-between mb-6">
					<h2 class="text-xl font-bold text-blue-600">Riwayat Singkronisasi</h2>
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
												Tanggal
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Deskripsi
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Status
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Pesan Status
											</span>
										</th>
										<th scope="col" class="px-3 text-center"></th>
									</tr>
								</thead>
								<tbody class="divide-y divide-gray-200">
									<template v-if="dataRiwayats" v-for="(item, idx) in dataRiwayats" :key="idx">
										<tr class="odd:bg-white even:bg-gray-100">
											<td class="size-px px-1" style="vertical-align: top;">
												<button type="button" class="block" data-hs-overlay="#hs-ai-invoice-modal">
													<span class="flex justify-center text-xs px-3 py-0">{{idx+1}}</span>
												</button>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-start text-xs py-0">{{item.created_at ? $dayjs(item.created_at).format("DD-MM-YYYY HH:mm:ss").toString() : '-'}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-start text-xs py-0">{{item.deskripsi}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">{{item.code == 1 ? 'Berhasil' : 'Gagal'}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-start text-xs py-0">{{item.message}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top;">
												<div class="sm:flex gap-x-1">
													<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-green-600 text-white hover:bg-green-700 disabled:opacity-50 disabled:pointer-events-none">
													  	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="flex-shrink-0 size-3">
															<path fill-rule="evenodd" d="M4.5 13a3.5 3.5 0 0 1-1.41-6.705A3.5 3.5 0 0 1 9.72 4.124a2.5 2.5 0 0 1 3.197 3.018A3.001 3.001 0 0 1 12 13H4.5Zm.72-5.03a.75.75 0 0 0 1.06 1.06l.97-.97v2.69a.75.75 0 0 0 1.5 0V8.06l.97.97a.75.75 0 1 0 1.06-1.06L8.53 5.72a.75.75 0 0 0-1.06 0L5.22 7.97Z" clip-rule="evenodd" />
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
							</table>
							<!-- End Table -->	
						</div>
					</div>
				</div>
			</div>
		</div>
	</LayoutSingkronisasi>
</template>