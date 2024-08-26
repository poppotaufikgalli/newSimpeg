<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)

const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const {snip, snbintang, snsk} = route.params
const {sbkn} = route.query

const dataRefPenghargaan = ref({
	nip: '',
	nbintang:'',
	nsk: '',
})

const dataPenghargaan = ref({
	nip: '',
	id_jenis_penghargaan: '',
	nbintang:'',
	aoleh: '',
	nsk: '',
	tsk: '',
	thn: 0,
	filename: '',
	idSync: '',
})

const jnss = ref([])
const pejabats = ref([])

const refJns = ref(null)
const refPejabat = ref(null)
const method = ref('Simpan')

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
	
  	const [jns, pejabat] = await Promise.all([
    	$fetch('/api/gets/jenis_penghargaan'),
    	$fetch('/api/gets/pejabat'),
  	])

  	jnss.value = jns
  	pejabats.value = pejabat
  	
  	if(refJns){
		refJns.value.reinit()
	}

	if(refPejabat){
		refPejabat.value.reinit()
	}
});

const { pending, data, refresh} = await useAsyncData('getdataPenghargaan', async() => {
	console.log("CariData : ", sbkn)

	if(snip){
		let nip = $decodeBase64(snip)

		if(sbkn != ""){
			let sid = $decodeBase64(sbkn)
			var result = await $fetch(`/api/gets/singkronisasi_bkn/getDataBKNById/penghargaan/${sid}`)
			
			let tsk = dayjs(dayjs(result.data.skDate, "DD-MM-YYYY")).format("YYYY-MM-DD[T]HH:mm:ss[Z]")
			dataPenghargaan.value = {
				nip: nip,
				id_jenis_penghargaan: result.data.hargaId,
				nbintang: result.data.hargaNama,
				tsk: tsk,
				nsk: result.data.skNomor,
				thn: result.data.tahun,
				idSync: sid,
			}
		}else{

			let nbintang = $decodeBase64(snbintang)
			let nsk = $decodeBase64(snsk)
			var result = await $fetch(`/api/gets/riwayat_penghargaan/${nip}/${nbintang}/${nsk}`);

			if(result.nip == nip){
				method.value = "Update"

				dataPenghargaan.value = _pickBy(result, function(val, key) {
					return _includes(_keys(dataPenghargaan.value), key);
				})

				dataRefPenghargaan.value = _pickBy(result, function(val, key) {
					return _includes(_keys(dataRefPenghargaan.value), key);
				})
			}
		}
	}
})

const tsk = computed({
	get(){
		return dayjs(dataPenghargaan.value.tsk).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataPenghargaan.value.tsk = val
	}
})

const aoleh = computed({
	get(){
		var result = _find(pejabats.value, {npej: dataPenghargaan.value.aoleh})
		if(result){
			return result.kpej	
		}
		return null
	},
	set(val){
		var result = _find(pejabats.value, {kpej: val})
		dataPenghargaan.value.aoleh = result.npej
	}
})

onMounted(() => {
	//refreshNuxtData(["getDataRef", "getdataPenghargaanAkhir"])
	refreshNuxtData(["getDataRef", "getdataPenghargaan"])
	modalUploadDoc.$onAction(callback, false)
})

const showSpinner = ref(false)
const toast = useToast()

async function simpanData() {
	showSpinner.value = true
	var canUpdateBkn = false

	var compacted = _pickBy(dataPenghargaan.value);
	compacted.tsk = compacted.tsk ? dayjs(dayjs(compacted.tsk, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.thn = compacted.thn *1

	var dr = dataRefPenghargaan.value
	//dr.kgolru = dr.kgolru *1

	if(method.value == 'Update'){
		var body = {
			"refdata" : dr,
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_penghargaan', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_penghargaan',
		    	icon: "i-heroicons-x-circle",
				title: "Update Data Penghargaan",
		    	description: error.value.data.data,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_penghargaan',
		    	icon: "i-heroicons-check-circle",
				title: "Update Data Penghargaan",
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		  //	canUpdateBkn = true
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		//compacted.akhir = 1
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_penghargaan/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_penghargaan',
		    	icon: "i-heroicons-check-circle",
				title: "Tambah Data Penghargaan",
		    	description: error.value.data.data,
		    	timeout: 6000,
		    	color: 'red'
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_penghargaan',
		    	icon: "i-heroicons-x-circle",
				title: "Tambah Data Penghargaan",
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
		  	//canUpdateBkn = true
		}
	}

	//update bkn

	if(canUpdateBkn == true && sbkn == ""){
		var {data, error} = await useFetch('/api/posts/riwayat_penghargaan/bkn', {
			method: 'POST',
			body: JSON.stringify(compacted),
		})

		if(error.value){
			const {title, description} = error.value.data.data
			toast.add({
				id: 'error_put_penghargaan_bkn',
				icon: "i-heroicons-x-circle",
				title: title,
				description: description,
				timeout: 6000,
				color: 'red',
			}) 	
		}else{
			toast.add({
				id: 'success_put_penghargaan_bkn',
				icon: "i-heroicons-check-circle",
				title: "Update Data Penghargaan BKN Berhasil",
				description: "Data telah di Update",
				timeout: 6000
			}) 	
		}
	}

	//end update bkn

  	showSpinner.value = false
  	refreshNuxtData(["getDataRef"])
}

const doUploadDoc = async() => {
	let nip = $decodeBase64(snip)
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataPenghargaan.value.filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen Penghargaan`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen Penghargaan`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		formData.append("nbintang", dataPenghargaan.value.nbintang);
		formData.append("nsk", dataPenghargaan.value.nsk);
		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_penghargaan_"+dataPenghargaan.value.nbintang+"_"+dataPenghargaan.value.nsk);
		formData.append("path", 'dokumen/'+nip);

		var {data, error} = await useFetch(`/api/uploads/upload/penghargaan/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(error.value){
			toast.add({
		    	id: 'error_post_upload_penghargaan',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_penghargaan',
		    	description: `Dokumen Penghargaan berhasil di Upload`,
		    	timeout: 6000
		  	}) 	
		}

		refreshNuxtData(["getDataRef", "getdataPenghargaan"])

		showSpinner.value = false
	}
}

function onlyNumber ($event) {
   	//console.log($event.keyCode); //keyCodes value
   	let keyCode = ($event.keyCode ? $event.keyCode : $event.which);
   	if ((keyCode < 48 || keyCode > 57) && keyCode !== 46) { // 46 is dot
    	$event.preventDefault();
   	}
}
</script>
<template>
	<LayoutRiwayatPegawai>
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="flex justify-end mb-4">
		          	<button type="button" class="flex justify-center items-center size-7 text-sm font-semibold border border-transparent text-gray-200 hover:text-gray-800 bg-red-500 hover:bg-red-200 disabled:opacity-50 disabled:pointer-events-none" @click="$router.go(-1)">
		            	<span class="sr-only">Close</span>
		           		<svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
		          	</button>
		        </div>
				<div class="mb-8">
					<div class="inline-flex justify-between items-center w-full gap-x-2">
						<h2 class="text-xl font-bold text-blue-600">Penghargaan/Tanda Jasa</h2>
						<div class="flex gap-x-2 w-[60%]">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP" :value="dataPenghargaan.nip">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="Nama">
						</div>
					</div>
				</div>
				
				<form>
					<!-- Grid -->
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jenis Penghargaan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refJns" id="jnss" :options="jnss" keyList="id" namaList="nama" v-model="dataPenghargaan.id_jenis_penghargaan"/>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nama Penghargaan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPenghargaan.nbintang">
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Asal Penghargaan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refPejabat" id="pejabats" :options="pejabats" keyList="kpej" namaList="npej" v-model="aoleh"/>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">SK Penghargaan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPenghargaan.nsk">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tsk">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Tahun Perolehan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPenghargaan.thn" @keypress="onlyNumber">
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
		<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2 mt-6" v-if="method == 'Update'">
			<div class="mb-8">
				<h2 class="text-xl font-bold text-blue-600">Dokumen Penghargaan</h2>
			</div>
			<div class="grid sm:grid-cols-4 gap-2 gap-2.5">
				<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc()">
					Dokumen Penghargaan
				</button>
			</div>
		</div>
		<!-- End Card Section -->
	</LayoutRiwayatPegawai>
</template>