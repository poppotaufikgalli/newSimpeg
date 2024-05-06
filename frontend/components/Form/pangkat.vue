<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)

const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const {snip, skgolru, stmtpang, sknpang} = route.params

const dataRefPangkat = ref({
	nip: '',
	knpang: '',
	tmtpang: '',
	kgolru: '',
})

const dataPangkat = ref({
	nip: '',
	nntbakn: '',
	tntbakn: '',
	ptetap: '',
	nskpang:'',
	tskpang: '',
	tmtpang: '',
	kgolru: '',
	mskerja: 0,
	blnkerja: 0,
	knpang: '',
	akredit: 0,
	akhir:0,
	filename:'',
})

const pejabats = ref([])
const pangkats = ref([])
const jenis_kps = ref([])

const refPejabat = ref(null)
const refPangkat = ref(null)
const refKjpns = ref(null)
const method = ref('Simpan')

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
  	const [pejabat, pangkat, jenis_kp] = await Promise.all([
    	$fetch('/api/gets/pejabat'),
    	$fetch('/api/gets/pangkat'),
    	$fetch('/api/gets/jenis_kp'),
  	])

  	pejabats.value = pejabat
  	pangkats.value = pangkat
  	jenis_kps.value = jenis_kp

  	if(refPejabat){
		refPejabat.value.reinit()
	}

	if(refPangkat){
		refPangkat.value.reinit()
	}

	if(refKjpns){
		refKjpns.value.reinit()
	}
});

const { pending, data, refresh} = await useAsyncData('getDataPangkat', async() => {
	if(snip){
		let nip = $decodeBase64(snip)

		if(route.query.akhir && (route.query.akhir == "true" || route.query.akhir == true)){
			var result = await $fetch('/api/gets/riwayat_pangkat/'+nip+'/akhir');

			if(result.nip == nip){
				method.value = "Update"

				dataPangkat.value = _pickBy(result, function(val, key) {
					return _includes(_keys(dataPangkat.value), key);
				})

				dataRefPangkat.value = _pickBy(result, function(val, key) {
					return _includes(_keys(dataRefPangkat.value), key);
				})
			}	
			akhir.value = route.query.akhir ? 1 : 0
		}else{
			let kgolru = $decodeBase64(skgolru)
			let tmtpang = $decodeBase64(stmtpang)
			let knpang = $decodeBase64(sknpang)

			var result = await $fetch('/api/gets/riwayat_pangkat/'+nip+'/'+kgolru+'/'+tmtpang+'/'+knpang);

			if(result.nip == nip){
				method.value = "Update"

				dataPangkat.value = _pickBy(result, function(val, key) {
					return _includes(_keys(dataPangkat.value), key);
				})

				dataRefPangkat.value = _pickBy(result, function(val, key) {
					return _includes(_keys(dataRefPangkat.value), key);
				})
			}	
		}
	}
})

onMounted(() => {
	refreshNuxtData(["getDataRef", "getDataPangkat"])
	modalUploadDoc.$onAction(callback, false)
})

const aakhir = computed({
	get(){
		return dataPangkat.value.akhir == 0 ? 'Tidak' : 'Ya'
	},
})

const tntbakn = computed({
	get(){
		return dayjs(dataPangkat.value.tntbakn).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataPangkat.value.tntbakn = val
	}
})

const tskpang = computed({
	get(){
		return dayjs(dataPangkat.value.tskpang).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataPangkat.value.tskpang = val
	}
})

const tmtpang = computed({
	get(){
		return dayjs(dataPangkat.value.tmtpang).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataPangkat.value.tmtpang = val
	}
})

const akhir = computed({
	get(){
		return dataPangkat.value.akhir == 0 ? false : true
	},
	set(val){
		dataPangkat.value.akhir = val
	}
})

const showSpinner = ref(false)
const toast = useToast()

async function simpanData(nMethod) {
	showSpinner.value = true

	var compacted = dataPangkat.value;
	compacted.tntbakn = compacted.tntbakn ? dayjs(dayjs(compacted.tntbakn, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tskpang = compacted.tskpang ? dayjs(dayjs(compacted.tskpang, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tmtpang = compacted.tmtpang ? dayjs(dayjs(compacted.tmtpang, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.kgolru = compacted.kgolru *1
	compacted.ptetap = compacted.ptetap *1
	compacted.mskerja = compacted.mskerja *1
	compacted.blnkerja = compacted.blnkerja *1	
	compacted.knpang = compacted.knpang *1
	compacted.akhir = compacted.akhir == true ? 1 : 0

	var dr = dataRefPangkat.value
	dr.kgolru = dr.kgolru *1

	if(nMethod == 'Update'){
		var body = {
			"refdata" : dr,
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_pangkat', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_pangkat',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_pangkat',
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		//compacted.akhir = 1
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_pangkat/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_pangkat',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_pangkat',
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
		}
	}
  	showSpinner.value = false
  	refreshNuxtData(["getDataRef", "getDataPangkat"])
}

function onlyNumber ($event) {
   	//console.log($event.keyCode); //keyCodes value
   	let keyCode = ($event.keyCode ? $event.keyCode : $event.which);
   	if ((keyCode < 48 || keyCode > 57) && keyCode !== 46) { // 46 is dot
    	$event.preventDefault();
   	}
}

const doUploadDoc = async() => {
	let nip = $decodeBase64(snip)
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataPangkat.value.filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen Pangkat`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen Pangkat`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		formData.append("tmtpang", tmtpang.value);
		formData.append("kgolru", dataPangkat.value.kgolru);
		formData.append("knpang", dataPangkat.value.knpang);

		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_pangkat"+dataPangkat.value.kgolru);
		formData.append("path", 'dokumen/'+nip);

		var {data, error} = await useFetch(`/api/uploads/upload/pangkat/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(error.value){
			toast.add({
		    	id: 'error_post_upload_pangkat',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_pangkat',
		    	description: `Dokumen Pangkat berhasil di Upload`,
		    	timeout: 6000
		  	}) 	
		}

		refreshNuxtData(["getDataRef", "getDataPangkat"])

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
					<h2 class="text-xl font-bold text-blue-600">Pangkat {{route.query.akhir ? 'Terakhir' : ''}}</h2>
					<div class="flex gap-x-2 w-[60%]">
						<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP" :value="dataPangkat.nip ?? snip">
						<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="Nama">
					</div>
				</div>
			</div>
			<form>
				<!-- Grid -->
				<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Nota BKN</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPangkat.nntbakn" placeholder="xx-xxxxxxxxxxxx">
							<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tntbakn">
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Pejabat Penetap</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<SearchSelect2 ref="refPejabat" id="pejabats" :options="pejabats" keyList="kpej" namaList="npej" v-model="dataPangkat.ptetap"/>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">SK Pangkat</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPangkat.nskpang">
							<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tskpang">
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">TMT Pangkat</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tmtpang">
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Pangkat / Golongan</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<SearchSelect2 ref="refPangkat" id="pangkats" :options="pangkats" keyList="id" namaList="pangkat_gol" v-model="dataPangkat.kgolru"/>
					</div>

					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Masa Kerja</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input type="text" class="py-2 px-3 block w-[20%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPangkat.mskerja" placeholder="thn" @keypress="onlyNumber">
							<input type="text" class="py-2 px-3 block w-[20%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPangkat.blnkerja" placeholder="bln" @keypress="onlyNumber">
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Jenis Kenaikan Pangkat</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<SearchSelect2 ref="refKjpns" id="jenis_kps" :options="jenis_kps" keyList="id" namaList="nama" v-model="dataPangkat.knpang" />
					</div>

					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Angka Kredit</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<div class="sm:flex gap-2">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPangkat.akredit" @keypress="onlyNumber">
						</div>
					</div>
					<!-- End Col -->
					<div class="sm:col-span-6"></div>

					<template v-if="!route.query.akhir">
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Pangkat Terakhir</label>
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
					  				Pangkat Akhir Baru
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
</template>