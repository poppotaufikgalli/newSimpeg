<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)

const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const {snip, stmtjab} = route.params

const dataRefJabatan = ref({
	nip: '',
	tmtjab: '',
})

const dataJabatan = ref({
	nip: '',
	kpej: '',
	nskjabat:'',
	tskjabat: '',
	tmtjab: '',
	id_opd: '',
	jnsjab: '',
	keselon: 0,
	kjab: 0,
	njab: '',
	akhir: 0,
	filename: '',
})

const pejabats = ref([])
const jenis_jabatans = ref([])
const eselons = ref([])
const opds = ref([])
const unitKerja = ref([])

const seljnsjab = ref()
const selopd = ref()
const selkjab = ref()

const kpej = ref(null)
const kgolru = ref(null)
const jns_jab = ref(null)
const kjab = ref(null)
const fjab = ref(null)
const keselon = ref(null)
const kunker = ref(null)

const refPejabat = ref(null)
const refOpd = ref(null)
const refSubOpd = ref(null)
const refJnsJab = ref(null)
const refEselon = ref(null)
const refJabatan = ref(null)
const method = ref('Simpan')

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
	//loading.value = true
  	const [pejabat, opd, jenis_jabatan, eselon, unit_kerja] = await Promise.all([
  		$fetch('/api/gets/pejabat'),
    	$fetch('/api/gets/opd'),
    	$fetch('/api/gets/jenis_jabatan'),
    	$fetch('/api/gets/eselon'),
    	$fetch('/api/posts/opd', {
    		method: 'POST',
    		body: JSON.stringify({
    			sfilter: [0,1]
    		})
    	}),
  	])

  	pejabats.value = pejabat
  	opds.value = unit_kerja
  	jenis_jabatans.value = jenis_jabatan
  	eselons.value = eselon
  	unitKerja.value = unit_kerja
  	
  	if(refPejabat){
		refPejabat.value.reinit()
	}

	if(refOpd){
		refOpd.value.reinit()
	}

	if(refJnsJab){
		refJnsJab.value.reinit()
	}

	if(refEselon){
		refEselon.value.reinit()
	}
});

const { pending, data, refresh} = await useAsyncData('getDataJabatan', async() => {
	if(snip){
		let nip = $decodeBase64(snip)

		if(route.query.akhir && (route.query.akhir == "true" || route.query.akhir == true)){
			var result = await $fetch('/api/gets/riwayat_jabatan/'+nip+'/akhir');

			if(result.nip == nip){
				method.value = "Update"

				dataJabatan.value = _pickBy(result, function(val, key) {
					return _includes(_keys(dataJabatan.value), key);
				})

				dataRefJabatan.value = _pickBy(result, function(val, key) {
					return _includes(_keys(dataRefJabatan.value), key);
				})

				selopd.value = result.id_opd
				seljnsjab.value = result.jnsjab
				selkjab.value = result.kjab
				CariJabatanNonFormasi()
			}
			akhir.value = route.query.akhir ? 1 : 0
		}else{
			let tmtjab = $decodeBase64(stmtjab)
			var result = await $fetch('/api/gets/riwayat_jabatan/'+nip+'/'+tmtjab);

			if(result.nip == nip){
				method.value = "Update"

				dataJabatan.value = _pickBy(result, function(val, key) {
					return _includes(_keys(dataJabatan.value), key);
				})

				dataRefJabatan.value = _pickBy(result, function(val, key) {
					return _includes(_keys(dataRefJabatan.value), key);
				})

				selopd.value = result.id_opd
				seljnsjab.value = result.jnsjab
				selkjab.value = result.kjab
				CariJabatanNonFormasi()
			}
		}
	}
})

const tskjabat = computed({
	get(){
		return dayjs(dataJabatan.value.tskjabat).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataJabatan.value.tskjabat = val
	}
})

const tmtjab = computed({
	get(){
		return dayjs(dataJabatan.value.tmtjab).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataJabatan.value.tmtjab = val
	}
})

//const proxyEselons = ref([])

watch(selopd, (value) => {
	dataJabatan.value.id_opd = value

	//CariJabatanNonFormasi()
})

watch(seljnsjab, (value) => {
	dataJabatan.value.jnsjab = value
	
	CariJabatanNonFormasi()
})

watch(selkjab, (value) => {
	console.log("selkjab change")
	dataJabatan.value.kjab = value

	if(jabatans.value.length > 0){
		var jab = _find(jabatans.value, function(o){
			return o.id == dataJabatan.value.kjab
		})

		var njab = [jab.nama, "PADA"];	
		//var njab = [refJabatan.value.readOptions(), "PADA"];	

		var opd = _find(opds.value, function(o){
			return o.id == dataJabatan.value.id_opd
		})

		if(dataJabatan.value.jnsjab == 1){
			dataJabatan.value.keselon = opd.id_eselon
		}else{
			dataJabatan.value.keselon = 99
		}

		if(refEselon){
			refEselon.value.reinit()	
		}

		if(jab.nama !== opd.nama){
			njab.push(opd.nama)
		}

		var newJab = stackOPD(opd, njab)

		dataJabatan.value.njab =  newJab.join(" ")
	}else{
		//CariJabatanNonFormasi()
	}
})

function stackOPD (opd, njab) {
	var parent_opd_id = opd.parent_opd 

	while(parent_opd_id != null){
		var result = _find(opds.value, function(o){
			return o.id == parent_opd_id
		})	

		if(result.id != 1){
			njab.push(result.nama)	
		}else{
			njab.push("KOTA TANJUNGPINANG")	
		}

		parent_opd_id = result.parent_opd
	}

	return njab
}

const jabatans = ref([])
const { data: dataJabatans, refresh: refreshDataJabatan } = await useAsyncData('getListDataJabatan', () => CariJabatanNonFormasi())

const CariJabatanNonFormasi = async() => {
	console.log("CariJabatanNonFormasi", dataJabatan.value.jnsjab)
	var url = ""

	switch(dataJabatan.value.jnsjab){
	case "2":
		url = "jabatan_ft"
		break;
	case "4":
		url = "jabatan_fu"
		break;
	case "5":
		url = "jabatan_negara"
		break;
	}

	if(url != ""){
		var result = await $fetch('/api/gets/'+url, {
			method: 'GET',
		})

		if(result){
			jabatans.value = result
			//console.log(jabatans.value)
			if(refJabatan){
				refJabatan.value.reinit()
			}
		}
	}else{
		var result = await $fetch('/api/gets/opd/'+dataJabatan.value.id_opd);

		if(result){
			var newJab = [{
				"id" : result.id,
				"nama": result.nama_jabatan,
			}]

			jabatans.value = newJab

			//await nextTick()
			//console.log(jabatans.value)
			if(jabatans.value != null && refJabatan){
				try {
					refJabatan.value.reinit()
				}catch(err) {
				  	console.log(err)
				}
			}	
		}
		
	}
}

onMounted(() => {
	refreshNuxtData(["getDataRef", "getDataJabatan"])
	modalUploadDoc.$onAction(callback, false)
})

const showSpinner = ref(false)
const toast = useToast()

async function simpanData(nMethod) {
	showSpinner.value = true

	//console.log(dataJabatan.value)
	var compacted = _pickBy(dataJabatan.value);
	compacted.tmtjab = compacted.tmtjab ? dayjs(dayjs(compacted.tmtjab, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tskjabat = compacted.tskjabat ? dayjs(dayjs(compacted.tskjabat, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.id_eselon = compacted.id_eselon *1
	compacted.status = compacted.status *1
	compacted.kpej = compacted.kpej *1
	compacted.akhir = compacted.akhir == true ? 1 : 0

	var dr = dataRefJabatan.value
	//dr.kgolru = dr.kgolru *1

	if(nMethod == 'Update'){
		var body = {
			"refdata" : dr,
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_jabatan', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_jabatan',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_jabatan',
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		//compacted.akhir = 1
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_jabatan/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_jabatan',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_jabatan',
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
		}
	}
  	showSpinner.value = false
  	refreshNuxtData(["getDataRef", "getDataJabatan"])
}

const akhir = computed({
	get(){
		return dataJabatan.value.akhir == 0 ? false : true
	},
	set(val){
		dataJabatan.value.akhir = val
	}
})

const aakhir = computed({
	get(){
		return dataJabatan.value.akhir == 0 ? 'Tidak' : 'Ya'
	},
})

const doUploadDoc = async() => {
	let nip = $decodeBase64(snip)
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataJabatan.value.filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen Jabatan`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen Jabatan`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		formData.append("tmtjab", tmtjab.value);
		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_jabatan"+tmtjab.value);
		formData.append("path", 'dokumen/'+nip);

		var {data, error} = await useFetch(`/api/uploads/upload/jabatan/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(error.value){
			toast.add({
		    	id: 'error_post_upload_jabatan',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_jabatan',
		    	description: `Dokumen Jabatan berhasil di Upload`,
		    	timeout: 6000
		  	}) 	
		}

		refreshNuxtData(["getDataRef", "getDataJabatan"])

		showSpinner.value = false
	}
}
</script>
<template>
	<AppLoadingSpinner :show="showSpinner" />
	<div class="mx-auto">
		<!-- Card -->
		<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
			<div class="flex justify-end mb-4" v-if="!route.query.akhir">
	          	<button type="button" class="flex justify-center items-center size-7 text-sm font-semibold border border-transparent text-gray-200 hover:text-gray-800 bg-red-500 hover:bg-red-200 disabled:opacity-50 disabled:pointer-events-none" @click="$router.go(-1)">
	            	<span class="sr-only">Close</span>
	           		<svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
	          	</button>
	        </div>	        
			<div class="mb-8">
				<div class="inline-flex justify-between items-center w-full gap-x-2">
					<h2 class="text-xl font-bold text-blue-600">Jabatan {{route.query.akhir ? 'Terakhir' : ''}}</h2>
					<div class="flex gap-x-2 w-[60%]">
						<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP" :value="dataJabatan.nip">
						<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="Nama">
					</div>
				</div>
			</div>
			<form>
				<!-- Grid -->
				<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Pejabat Penetap</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<SearchSelect2 ref="refPejabat" id="pejabats" :options="pejabats" keyList="kpej" namaList="npej" v-model="dataJabatan.kpej"/>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">SK Jabatan</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataJabatan.nskjabat">
							<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tskjabat">
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">TMT Jabatan</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tmtjab">
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Unit Kerja</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<SearchSelect2 ref="refOpd" id="opds" :options="opds" keyList="id" namaList="nama" v-model="selopd"/>
					</div>

					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Jenis Jabatan</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-4">
						<!--<SearchSelect2 ref="refJnsJab" id="jenis_jabatans" :options="jenis_jabatans" keyList="id" namaList="nama" v-model="dataJabatan.jnsjab"/>-->
						<SearchSelect2 ref="refJnsJab" id="jenis_jabatans" :options="jenis_jabatans" keyList="id" namaList="nama" v-model="seljnsjab"/>
					</div>

					<div class="sm:col-span-2">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Eselon</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<SearchSelect2 ref="refEselon" id="eselons" :options="eselons" keyList="id" namaList="nama" v-model="dataJabatan.keselon"/>
					</div>

					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Jabatan</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<SearchSelect2 ref="refJabatan" id="jabatans" :options="jabatans" keyList="id" namaList="nama" v-model="selkjab"/>
					</div>

					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Nama Jabatan</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<textarea id="af-account-bio" class="py-2 px-3 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" rows="3" placeholder="Nama Jabatan" v-model="dataJabatan.njab"></textarea>
						</div>
					</div>
					<!-- End Col -->

					<template v-if="!route.query.akhir">
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jabatan Terakhir</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<div class="flex mt-2">
								  	<input type="checkbox" class="shrink-0 mt-0.5 border-gray-200 rounded text-blue-600 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none" id="hs-default-checkbox" v-model="akhir">
								  	<label for="hs-default-checkbox" class="text-sm text-gray-500 ms-3">{{aakhir}}</label>
								</div>
							</div>
						</div>
						<!-- End Col -->
					</template>
				</div>
				<div class="mt-5 grid sm:grid-cols-12 gap-x-2">
					<div class="sm:col-span-7 sm:col-start-4">
						<div class="sm:flex gap-2">
							<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none" @click="refresh">
				  				Batal
							</button>
							<template v-if="route.query.akhir">
								<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-yellow-600 text-white hover:bg-yellow-700 disabled:opacity-50 disabled:pointer-events-none" @click="simpanData('Baru')">
					  				Jabatan Akhir Baru
								</button>	
							</template>
							<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="simpanData(method)">
				  				{{method}} Data
							</button>	
						</div>					
					</div>
					<!-- End Col -->

					<div class="sm:col-span-2 flex justify-end" v-if="method == 'Update'">
						<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-green-600 text-white hover:bg-green-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="doUploadDoc()">
			  				Upload File
						</button>
					</div>
					<!-- End Col -->
	  			</div>
			</form>
  		</div>
  		<!-- End Card -->
	</div>
	<!-- End Card Section -->
</template>