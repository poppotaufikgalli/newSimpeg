<script setup>
const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const toast = useToast()

const id = ref(route.params.id)
const dataUser = ref({
	nip: '',
	name: '',
	email: '',
	gid: '',
	stts: '',
})
const method = ref("Simpan")

const {data: dtGroup} = await useFetch('/api/gets/group')

const { pending, data, refresh } = await useAsyncData('getDataUser', async() => {
	let sid = $decodeBase64(id.value)
	if(sid){
		dataUser.value = await $fetch('/api/gets/users/'+sid);

		if(dataUser.value.id > 0){
			method.value = "Update"
		}
	}
})

const cariNIP = async() => {
	const {nip} = dataUser.value
	if(nip.length != 18){
		toast.add({
	    	id: 'error_search_nip',
	    	description: "NIP Tidak Valid",
	    	timeout: 6000
	  	})
	}else{
		var result = await $fetch('/api/gets/pegawai/'+nip);
		console.log(result.nama)
		if(result){
			const {nama} = result
			dataUser.value.name = nama
		}
	}
}

onMounted(() => {
	refreshNuxtData(["getDataUser"])
})

async function simpanData() {
	//showSpinner.value = true

	var compacted = _pickBy(dataUser.value);
	compacted.gid = compacted.gid *1
	compacted.stts = compacted.stts *1

	if(method.value == 'Update'){
		var body = JSON.stringify(compacted)
		//url = '/api/pegawai'
		var {data, error} = await useFetch('/api/puts/users', {
			method: 'PUT',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_users',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_users',
		    	description: "Data Users berhasil di Update",
		    	timeout: 6000
		  	}) 	
		  	navigateTo({ path: '/konfigurasi/sistem/pengguna/' })
		}
		//console.log(error.value.data.data)
		
	}else{
		compacted.password = compacted.nip
		//console.log(compacted)
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/users/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_users',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_users',
		    	description: "Data Users berhasil di Ditambahkan",
		    	timeout: 6000
		  	})
		  	navigateTo({ path: '/konfigurasi/sistem/pengguna/' })
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
					<h2 class="text-xl font-bold text-blue-600">Data Pengguna</h2>
				</div>
				<form>
					<!-- Grid -->
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">NIP</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataUser.nip" placeholder="xxxxxxxxxxxxxxxxxxx" maxlength="18">
								<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none whitespace-nowrap" @click="cariNIP">
					  				Cari NIP
								</button>	
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nama</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" disabled v-model="dataUser.name">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Email</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="email" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataUser.email">
							</div>
						</div>
						<!-- End Col -->


						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Group Pengguna</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataUser.gid">
								  	<option selected>Pilih Status Pengguna</option>
								  	<template v-if="dtGroup" v-for="(item, idx) in dtGroup">
								  		<option :value=item.id>{{item.id}} - {{item.nama}}</option>	
								  	</template>
								</select>
							</div>
						</div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Status Pengguna</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex">
  								<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataUser.stts">
								  	<option selected>Pilih Status Pengguna</option>
								  	<option value=1>1 - Aktif</option>
								  	<option value=0>0 - Tidak Aktif</option>
								</select>
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