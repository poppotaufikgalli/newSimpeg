<script setup>
const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const toast = useToast()

const id = ref(route.params.id)
const dataGroup = ref({
	nama: '',
})
const method = ref("Simpan")

const { pending, data, refresh } = await useAsyncData('getDataGroup', async() => {
	let sid = $decodeBase64(id.value)
	if(sid){
		dataGroup.value = await $fetch('/api/gets/group/'+sid);

		if(dataGroup.value.id > 0){
			method.value = "Update"
		}
	}
})

onMounted(() => {
	refreshNuxtData(["getDataGroup"])
})

async function simpanData() {
	//showSpinner.value = true

	var compacted = _pickBy(dataGroup.value);

	if(method.value == 'Update'){
		var body = JSON.stringify(compacted)
		//url = '/api/pegawai'
		var {data, error} = await useFetch('/api/puts/group', {
			method: 'PUT',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_group',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_group',
		    	description: "Data Group berhasil di Update",
		    	timeout: 6000
		  	}) 	
		  	navigateTo({ path: '/konfigurasi/sistem/group' })
		}
		//console.log(error.value.data.data)
		
	}else{
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/group/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_group',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_group',
		    	description: "Data Group berhasil di Ditambahkan",
		    	timeout: 6000
		  	})
		  	navigateTo({ path: '/konfigurasi/sistem/group' })
		}
	}
  	//showSpinner.value = false
}

</script>
<template>
	<LayoutKonfigurasi>
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="mb-8">
					<h2 class="text-xl font-bold text-blue-600">Data Group</h2>
				</div>
				<form>
					<!-- Grid -->
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nama Group</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataGroup.nama">
							</div>
						</div>
						<!-- End Col -->
					</div>
					<div class="mt-5 grid sm:grid-cols-12 gap-x-2">
						<div class="sm:col-span-6 sm:col-start-4">
							<div class="sm:flex gap-2">
								<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none" @click="refresh">
					  				Batal
								</button>
								<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="simpanData">
					  				{{method}} Data
								</button>	
							</div>
						</div>
						<!-- End Col -->
		  			</div>
				</form>
	  		</div>
	  		<!-- End Card -->
		</div>
	</LayoutKonfigurasi>
</template>