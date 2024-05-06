<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)

const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const {snip, stmtngaj} = route.params

const dataRefKgb = ref({
	nip: '',
	tmtngaj: '',
})

const dataKgb = ref({
	nip: '',
	kpej: '',
	nstahu:'',
	tstahu: '',
	tmtngaj: '',
	id_opd: '',
	kkantor: '',
	mskerja: 0,
	mskerjabl: 0,
	nunker: '',
	gpokkhir: 0,
	akhir: 0,
	filename: '',
})

const pejabats = ref([])
const opds = ref([])
const kppns = ref([])

const refPejabat = ref()
const refOpd = ref()
const refKppn = ref()
const method = ref("Simpan")

const { pending, data, refresh} = await useAsyncData('getDataGajiBerkala', async() => {
	if(snip){
		let nip = $decodeBase64(snip)

		if(route.query.akhir && (route.query.akhir == "true" || route.query.akhir == true)){
			var result = await $fetch('/api/gets/riwayat_gajiberkala/'+nip+'/akhir');

			if(result.nip == nip){
				method.value = "Update"

				dataKgb.value = _pickBy(result, function(val, key) {
					return _includes(_keys(dataKgb.value), key);
				})

				dataRefKgb.value = _pickBy(result, function(val, key) {
					return _includes(_keys(dataRefKgb.value), key);
				})
			}
			akhir.value = route.query.akhir ? 1 : 0
		}else{
			let tmtngaj = $decodeBase64(stmtngaj)
			var result = await $fetch('/api/gets/riwayat_gajiberkala/'+nip+'/'+tmtngaj);

			if(result.nip == nip){
				method.value = "Update"

				dataKgb.value = _pickBy(result, function(val, key) {
					return _includes(_keys(dataKgb.value), key);
				})

				dataRefKgb.value = _pickBy(result, function(val, key) {
					return _includes(_keys(dataRefKgb.value), key);
				})
			}
		}
	}
	
})

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
	//loading.value = true
  	const [opd, pejabat, kppn] = await Promise.all([
    	$fetch('/api/gets/opd'),
    	$fetch('/api/gets/pejabat'),
    	$fetch('/api/gets/kppn'),
  	])

  	opds.value = opd
  	pejabats.value = pejabat
  	kppns.value = kppn

  	if(refPejabat){
		refPejabat.value.reinit()
	}

	if(refOpd){
		refOpd.value.reinit()
	}

	if(refKppn){
		refKppn.value.reinit()
	}
});

const tstahu = computed({
	get(){
		return dayjs(dataKgb.value.tstahu).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataKgb.value.tstahu = val
	}
})

const tmtngaj = computed({
	get(){
		return dayjs(dataKgb.value.tmtngaj).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataKgb.value.tmtngaj = val
	}
})

const gpokkhir = computed({
	get(){
		var a = dataKgb.value.gpokkhir.toString().split('.').join("") 
		console.log(a)
		return a.replace(/(\d)(?=(\d\d\d)+(?!\d))/g, "$1.")
	},
	set(val){
		dataKgb.value.gpokkhir = val
	}
})

onMounted(() => {
	refreshNuxtData(["getDataRef", "getDataGajiBerkala"])
	modalUploadDoc.$onAction(callback, false)
})

function onlyNumber ($event) {
   	//console.log($event.keyCode); //keyCodes value
   	let keyCode = ($event.keyCode ? $event.keyCode : $event.which);
   	if ((keyCode < 48 || keyCode > 57) && keyCode !== 46) { // 46 is dot
    	$event.preventDefault();
   	}
}

const showSpinner = ref(false)
const toast = useToast()

async function simpanData(nMethod) {
	showSpinner.value = true

	//console.log(dataJabatan.value)
	var compacted = _pickBy(dataKgb.value);
	compacted.tstahu = compacted.tstahu ? dayjs(dayjs(compacted.tstahu, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tmtngaj = compacted.tmtngaj ? dayjs(dayjs(compacted.tmtngaj, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.kpej = compacted.kpej *1
	compacted.mskerja = compacted.mskerja *1
	compacted.mskerjabl = compacted.mskerjabl *1
	compacted.gpokkhir = compacted.gpokkhir.toString().split(".").join("") *1
	compacted.akhir = compacted.akhir == true ? 1 : 0

	var dr = dataRefKgb.value
	//dr.kgolru = dr.kgolru *1

	if(nMethod == 'Update'){
		var body = {
			"refdata" : dr,
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_gajiberkala', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_kgb',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_kgb',
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		//compacted.akhir = 1
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_gajiberkala/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_kgb',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_kgb',
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
		}
	}
  	showSpinner.value = false
  	refreshNuxtData(["getDataRef", "getDataGajiBerkala"])
}

const akhir = computed({
	get(){
		return dataKgb.value.akhir == 0 ? false : true
	},
	set(val){
		dataKgb.value.akhir = val
	}
})

const aakhir = computed({
	get(){
		return dataKgb.value.akhir == 0 ? 'Tidak' : 'Ya'
	},
})

const doUploadDoc = async() => {
	let nip = $decodeBase64(snip)
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataKgb.value.filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen Kenaikan Gaji Berkala`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen Kenaikan Gaji Berkala`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		formData.append("tmtngaj", tmtngaj.value);
		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_kgb"+tmtngaj.value);
		formData.append("path", 'dokumen/'+nip);

		var {data, error} = await useFetch(`/api/uploads/upload/kgb/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(error.value){
			toast.add({
		    	id: 'error_post_upload_kgb',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_kgb',
		    	description: `Dokumen Kenaikan Gaji Berkala berhasil di Upload`,
		    	timeout: 6000
		  	}) 	
		}

		refreshNuxtData(["getDataRef", "getDataGajiBerkala"])

		showSpinner.value = false
	}
}
</script>
<template>
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
					<h2 class="text-xl font-bold text-blue-600">Gaji Berkala {{route.query.akhir ? 'Terakhir' : ''}}</h2>
					<div class="flex gap-x-2 w-[60%]">
						<input type="text" name="hs-input-with-add-on-url" id="hs-input-with-add-on-url" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP" :value="dataKgb.nip">
						<input type="text" name="hs-input-with-add-on-url" id="hs-input-with-add-on-url" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="Nama">
					</div>
				</div>
			</div>

			<form>
				<!-- Grid -->
				<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Pejabat Penetap</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<SearchSelect2 ref="refPejabat" id="pejabats" :options="pejabats" keyList="kpej" namaList="npej" v-model="dataKgb.kpej"/>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">SK Gaji Berkala</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKgb.nstahu">
							<input id="af-account-full-name" type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tstahu">
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">TMT Berkala</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input id="af-account-full-name" type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tmtngaj">
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Masa Kerja</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-[20%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKgb.mskerja" placeholder="thn" @keypress="onlyNumber">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-[20%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKgb.mskerjabl" placeholder="bln" @keypress="onlyNumber">
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Gaji Pokok</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="gpokkhir" @keypress="onlyNumber">
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Unit Kerja</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<SearchSelect2 ref="refOpd" id="opds" :options="opds" keyList="id" namaList="nama" v-model="dataKgb.id_opd"/>
					</div>

					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Kantor Pembayaran</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<SearchSelect2 ref="refKppn" id="kppns" :options="kppns" keyList="id" namaList="nama" v-model="dataKgb.kkantor"/>
					</div>

					<template v-if="!route.query.akhir">
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Gaji Berkala Terakhir</label>
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
					<div class="sm:col-span-6 sm:col-start-4">
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

					<div class="sm:col-span-3 flex justify-end" v-if="method == 'Update'">
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