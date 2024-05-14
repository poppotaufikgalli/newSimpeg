<script setup>
const dayjs = useDayjs()
const toast = useToast()

const { data, pending, error, refresh } = await useFetch('/api/posts/singkronisasi', {
  	method: 'POST',
  	body: JSON.stringify({
  		host: 'bkn'
  	})
})

const proxyData = computed({
	get(){
		return data.value[0]
	}
})

const method = computed({
	get(){
		if(data){
			return "Update"
		}else{
			return "Tambah"
		}
	}
})

async function simpanData(){
	if(method.value == 'Update'){
		var compacted = _pick(proxyData.value, ['id', 'uri', 'ckey', 'csecret', 'username', 'password', 'grant_type']);
		var body = JSON.stringify(compacted)
		
		var {data, error} = await useFetch('/api/puts/singkronisasi', {
			method: 'PUT',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_singkronisasi',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_singkronisasi',
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
	}else{
		var compacted = _pick(proxyData.value, ['uri', 'ckey', 'csecret', 'username', 'password', 'grant_type']);
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/singkronisasi/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_identitas',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_identitas',
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
		}
	}
}

async function updateToken(type){
	var result = await $fetch('/api/posts/singkronisasi/updateToken', {
		method: 'POST',
		body: JSON.stringify({
			type: type,
			ckey: proxyData.value.ckey,
			csecret: proxyData.value.csecret,
			client_id: proxyData.value.client_id,
			username: proxyData.value.username,
			password: proxyData.value.password,
		})
	})
	
	const {token_sso_expired, token_apimanager_expired} = result

	if(type == 'token_sso'){
		proxyData.value.token_sso_expired = dayjs(token_sso_expired).format('YYYY-MM-DDTHH:mm:ssZ')
	}

	if(type == 'token_apimanager'){
		proxyData.value.token_apimanager_expired = dayjs(token_apimanager_expired).format('YYYY-MM-DDTHH:mm:ssZ')
	}
}

</script>
<template>
	<LayoutKonfigurasi>
		<div class="mx-auto">
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2 mb-8">
				<div class="mb-8">
					<h2 class="text-xl font-bold text-blue-600">Aktivasi Token</h2>
				</div>
				<template v-if="pending">
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5 animate-pulse">
						<div class="sm:col-span-3">
							<p class="h-8 bg-gray-200 rounded-lg"></p>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<p class="h-10 bg-gray-200 rounded-lg"></p>
						</div>
						<!-- End Col -->
						<div class="sm:col-span-3">
							<p class="h-8 bg-gray-200 rounded-lg"></p>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<span class="h-10 w-full bg-gray-200 rounded-lg"></span>
								<span class="h-10 w-full bg-gray-200 rounded-lg"></span>
							</div>
						</div>
						<!-- End Col -->
						<div class="sm:col-span-3">
							<p class="h-8 bg-gray-200 rounded-lg"></p>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<p class="h-10 bg-gray-200 rounded-lg"></p>
						</div>
						<!-- End Col -->
					</div>
				</template>
				<template v-else>
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Token SSO Expired</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 disabled:opacity-50 disabled:pointer-events-none bg-white" :value="proxyData.token_sso_expired" :class="$dayjs().isBefore($dayjs(proxyData.token_sso_expired)) ? 'border-gray-200 focus:border-blue-500 focus:ring-blue-500' : 'border-red-500 focus:border-red-500 focus:ring-red-500'" disabled>
								<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[36px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" :disabled="$dayjs().isBefore($dayjs(proxyData.token_sso_expired))"  @click="updateToken('token_sso')">
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
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" :value="proxyData.token_apimanager_expired" :class="$dayjs().isBefore($dayjs(proxyData.token_apimanager_expired)) ? 'border-gray-200 focus:border-blue-500 focus:ring-blue-500' : 'border-red-500 focus:border-red-500 focus:ring-red-500'" disabled>
								<button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[36px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" :disabled="$dayjs().isBefore($dayjs(proxyData.token_apimanager_expired))" @click="updateToken('token_apimanager')">
								  	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-5 h-5">
										<path fill-rule="evenodd" d="M15.312 11.424a5.5 5.5 0 0 1-9.201 2.466l-.312-.311h2.433a.75.75 0 0 0 0-1.5H3.989a.75.75 0 0 0-.75.75v4.242a.75.75 0 0 0 1.5 0v-2.43l.31.31a7 7 0 0 0 11.712-3.138.75.75 0 0 0-1.449-.39Zm1.23-3.723a.75.75 0 0 0 .219-.53V2.929a.75.75 0 0 0-1.5 0V5.36l-.31-.31A7 7 0 0 0 3.239 8.188a.75.75 0 1 0 1.448.389A5.5 5.5 0 0 1 13.89 6.11l.311.31h-2.432a.75.75 0 0 0 0 1.5h4.243a.75.75 0 0 0 .53-.219Z" clip-rule="evenodd" />
									</svg>
								</button>
							</div>
						</div>
						<!-- End Col -->
					</div>
				</template>
			</div>
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="mb-8">
					<h2 class="text-xl font-bold text-blue-600">Informasi Koneksi</h2>
				</div>
				<template v-if="pending">
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5 animate-pulse">
						<div class="sm:col-span-3">
							<p class="h-8 bg-gray-200 rounded-lg"></p>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<p class="h-10 bg-gray-200 rounded-lg"></p>
						</div>
						<!-- End Col -->
						<div class="sm:col-span-3">
							<p class="h-8 bg-gray-200 rounded-lg"></p>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<span class="h-10 w-full bg-gray-200 rounded-lg"></span>
								<span class="h-10 w-full bg-gray-200 rounded-lg"></span>
							</div>
						</div>
						<!-- End Col -->
						<div class="sm:col-span-3">
							<p class="h-8 bg-gray-200 rounded-lg"></p>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<p class="h-10 bg-gray-200 rounded-lg"></p>
						</div>
						<!-- End Col -->
					</div>
				</template>
				<template v-else>
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Url</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="proxyData.uri">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Customer Key</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="proxyData.ckey">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Customer Secret</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="proxyData.csecret">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Username</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="proxyData.username">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Password</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="proxyData.password">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Grant Type</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="proxyData.grant_type">
							</div>
						</div>
						<!-- End Col -->
					</div>
					<div class="mt-5 flex justify-end gap-x-2">
						<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none" @click="refresh">
			  				Batal
						</button>
						<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="simpanData">
			  				{{method}} Data
						</button>
		  			</div>
				</template>
			</div>
		</div>
	</LayoutKonfigurasi>
</template>