<script setup>
onMounted(() => {
	refreshNuxtData(["getData"])
})

const { pending, data, refresh} = await useLazyAsyncData('getData', () => CariData())

const CariData = async() => {
	console.log("Cari Data Singkronisasi")

	var body = JSON.stringify({
		host: 'bkn',
	});

	return await $fetch('/api/posts/singkronisasi', {
		method: 'POST',
		body: body,
	});
}

const tokenInfo = computed({
	get(){
		return data.value[0]
	}
})
</script>
<template>
	<LayoutSingkronisasi>
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<form>
					<!-- Grid -->
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Token SSO Expired</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 disabled:opacity-50 disabled:pointer-events-none bg-white" :value="tokenInfo.token_sso_expired" v-if="data" :class="$dayjs().isBefore($dayjs(tokenInfo.token_sso_expired)) ? 'border-gray-200 focus:border-blue-500 focus:ring-blue-500' : 'border-red-500 focus:border-red-500 focus:ring-red-500'" disabled>
								<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[36px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none">
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
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" :value="tokenInfo.token_apimanager_expired" v-if="data" :class="$dayjs().isBefore($dayjs(tokenInfo.token_apimanager_expired)) ? 'border-gray-200 focus:border-blue-500 focus:ring-blue-500' : 'border-red-500 focus:border-red-500 focus:ring-red-500'" disabled>
								<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[36px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none">
								  	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-5 h-5">
  											<path fill-rule="evenodd" d="M15.312 11.424a5.5 5.5 0 0 1-9.201 2.466l-.312-.311h2.433a.75.75 0 0 0 0-1.5H3.989a.75.75 0 0 0-.75.75v4.242a.75.75 0 0 0 1.5 0v-2.43l.31.31a7 7 0 0 0 11.712-3.138.75.75 0 0 0-1.449-.39Zm1.23-3.723a.75.75 0 0 0 .219-.53V2.929a.75.75 0 0 0-1.5 0V5.36l-.31-.31A7 7 0 0 0 3.239 8.188a.75.75 0 1 0 1.448.389A5.5 5.5 0 0 1 13.89 6.11l.311.31h-2.432a.75.75 0 0 0 0 1.5h4.243a.75.75 0 0 0 .53-.219Z" clip-rule="evenodd" />
									</svg>
								</button>
							</div>
						</div>
						<!-- End Col -->
					</div>
				</form>
			</div>
		</div>
	</LayoutSingkronisasi>
</template>