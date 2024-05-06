<script setup>
const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const toast = useToast()

const id = ref(route.params.id)
const sid = $decodeBase64(id.value)
const dataOpd = ref({
	id: sid,
	parent_opd: '',
	nama: '',
	id_eselon: '',
	id_jabatan: '',
	nama_jabatan: '',
	id_unor_bkn: '',
	status: 1,
})

const atasan = ref();
const method = ref("Simpan")

const { pending, data, refresh } = await useAsyncData('getDataOpd', async() => {
	//let sid = $decodeBase64(id.value)
	if(sid){
		var result = await $fetch('/api/gets/opd/'+sid);

		if(result.id == sid){
			method.value = "Update"

			dataOpd.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataOpd.value), key);
			})
		}
	}
})

const opdInduks = ref([])
const eselons = ref([])
const unor_bkns = ref([])

const refOpdInduk = ref(null)
const refEselon = ref(null)
const refUnorBkn = ref(null)

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useAsyncData('getDataRef', async () => {
  	const [opdInduk, eselon, unor_bkn] = await Promise.all([
    	$fetch('/api/posts/opd', {
    		method: 'POST',
    		body: JSON.stringify({
    			sfilter: [0,1]
    		})
    	}),
    	$fetch('/api/gets/eselon'),
    	$fetch('/api/gets/unor_bkn')
  	])

  	opdInduks.value = opdInduk
  	eselons.value = eselon
  	unor_bkns.value = unor_bkn

  	if(refOpdInduk){
  		refOpdInduk.value.reinit()
  	}

  	if(refEselon){
  		refEselon.value.reinit()
  	}

  	if(refUnorBkn){
  		refUnorBkn.value.reinit()
  	}
});

onMounted(() => {
	refreshNuxtData(["getDataOpd", "getDataRef"])
})

async function simpanData() {
	//showSpinner.value = true

	var compacted = _pickBy(dataOpd.value);
	compacted.id_eselon = compacted.id_eselon *1
	compacted.status = compacted.status *1

	if(method.value == 'Update'){
		var body = JSON.stringify(compacted)
		//url = '/api/pegawai'
		var {data, error} = await useFetch('/api/puts/opd', {
			method: 'PUT',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_opd',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_opd',
		    	description: "Data OPD berhasil di Update",
		    	timeout: 6000
		  	}) 	
		  	navigateTo({ path: '/referensi/opd' })
		}
		//console.log(error.value.data.data)
		
	}else{
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/opd/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_opd',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_opd',
		    	description: "Data OPD berhasil di Ditambahkan",
		    	timeout: 6000
		  	})
		  	navigateTo({ path: '/referensi/opd' })
		}
	}
  	//showSpinner.value = false
}

watch(atasan, (newVal) => {
	dataOpd.value.parent_opd = newVal

	if(method.value != "Update"){
		findNewIDOPD(newVal)	
	}
})

const findNewIDOPD = (parent) => {
	var result = _filter(opdInduks.value, {parent_opd: parent})
	if(result.length > 0){
		var maxID = _maxBy(result, 'id')

		dataOpd.value.id = parseInt(maxID.id) +1
	}else{
		dataOpd.value.id = parent + '01'
	}
	
}

</script>
<template>
	<LayoutReferensi>
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="mb-8">
					<h2 class="text-xl font-bold text-blue-600">Data OPD</h2>
				</div>
				
				<form>
					<!-- Grid -->
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Atasan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refOpdInduk" id="opdInduks" :options="opdInduks" keyList="id" namaList="nama" v-model="atasan" />
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">ID</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex  gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataOpd.id" :disabled="method == 'Update'" required>
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nama</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataOpd.nama">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Eselon</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<SearchSelect2 ref="refEselon" id="eselons" :options="eselons" keyList="id" namaList="nama" v-model="dataOpd.id_eselon" />
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nama Jabatan Kepala</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataOpd.nama_jabatan">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Unor BKN</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refUnorBkn" id="unor_bkns" :options="unor_bkns" keyList="Id" namaList="NamaUnor" v-model="dataOpd.id_unor_bkn" />
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Status</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex">
  								<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataOpd.status">
								  	<option selected>Pilih Status</option>
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
	</LayoutReferensi>
</template>