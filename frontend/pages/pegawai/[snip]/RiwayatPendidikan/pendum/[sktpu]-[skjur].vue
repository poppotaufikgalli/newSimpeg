<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)

const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const {snip, sktpu, skjur} = route.params

const dataRefPendum = ref({
	nip: '',
	ktpu: '',
	kjur: '',
})

const dataPendum = ref({
	nip: '',
	ktpu: '',
	kjur:'',
	nsek: '',
	tempat: '',
	nkepsek: '',
	nsttb: '',
	tsttb: '',
	akhir: 0,
	npdum: '',
	negara: '',
	ket:'',
	filename: '',
})

const jurusans = ref([])
const jurusanss = ref([])
const ktpus = ref([])
const refKjur = ref(null)
const refKtpu = ref(null)
const method = ref('Simpan')

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
	
  	const [ktpu, jurusan] = await Promise.all([
    	$fetch('/api/gets/tingkat_pendidikan'),
    	$fetch('/api/gets/pendidikan'),
  	])

  	var aktpu = _sortBy(ktpu, [function(o){
  		return o.id *1
  	}])

  	ktpus.value = aktpu
  	jurusans.value = jurusan
  	jurusanss.value = jurusan

	if(refKtpu){
		refKtpu.value.reinit()
	}

	if(refKjur){
		refKjur.value.reinit()
	}
});

const { pending, data, refresh} = await useAsyncData('getdataPendum', async() => {
	console.log("CariData")
	if(snip){
		let nip = $decodeBase64(snip)
		let ktpu = $decodeBase64(sktpu)
		let kjur = $decodeBase64(skjur)
		var result = await $fetch(`/api/gets/riwayat_pendum/${nip}/${ktpu}/${kjur}`);

		if(result.nip == nip){
			method.value = "Update"

			dataPendum.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataPendum.value), key);
			})

			dataRefPendum.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataRefPendum.value), key);
			})
		}
	}
})

const ktpu = computed({
	get(){
		return dataPendum.value.ktpu
	},
	set(val){
		dataPendum.value.ktpu = val

		cariJurusan(val)
	}
})

async function cariJurusan(val){
	var item = _find(ktpus.value, {id: val})
	var result = await $fetch(`/api/posts/pendidikan/`, {
		method: "POST",
		body: JSON.stringify({
			'tk_pendidikan_id' : item.ref_simpeg.split(", ")
		})
	});

	jurusans.value = result

	if(refKjur){
		refKjur.value.reinit()
	}
}

const tsttb = computed({
	get(){
		return dayjs(dataPendum.value.tsttb).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataPendum.value.tsttb = val
	}
})

const aakhir = computed({
	get(){
		return dataPendum.value.akhir == 0 ? 'Tidak' : 'Ya'
	},
})

const akhir = computed({
	get(){
		return dataPendum.value.akhir == 0 ? false : true
	},
	set(val){
		dataPendum.value.akhir = val
	}
})

onMounted(() => {
	//refreshNuxtData(["getDataRef", "getdataPendumAkhir"])
	refreshNuxtData(["getDataRef", "getdataPendum"])
	modalUploadDoc.$onAction(callback, false)
})

const showSpinner = ref(false)
const toast = useToast()

async function simpanData() {
	showSpinner.value = true

	var compacted = _pickBy(dataPendum.value);
	compacted.tsttb = compacted.tsttb ? dayjs(dayjs(compacted.tsttb, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.akhir = compacted.akhir == true ? 1 : 0
	
	var dr = dataRefPendum.value

	if(method.value == 'Update'){
		var body = {
			"refdata" : dr,
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_pendum', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_pendum',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_pendum',
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_pendum/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_pendum',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_pendum',
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
		}
	}
  	showSpinner.value = false
  	refreshNuxtData(["getDataRef"])
}

const doUploadDoc = async() => {
	let nip = $decodeBase64(snip)
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataPendum.value.filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen Pendidikan Umum`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen Pendidikan Umum`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		formData.append("ktpu", dataPendum.value.ktpu);
		formData.append("kjur", dataPendum.value.kjur);
		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_pendum_"+dataPendum.value.ktpu+"_"+dataPendum.value.kjur);
		formData.append("path", 'dokumen/'+nip);

		var {data, error} = await useFetch(`/api/uploads/upload/pendum/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(error.value){
			toast.add({
		    	id: 'error_post_upload_pendum',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_pendum',
		    	description: `Dokumen Pendidikan Umum berhasil di Upload`,
		    	timeout: 6000
		  	}) 	
		}

		refreshNuxtData(["getDataRef", "getdataPendum"])

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
						<h2 class="text-xl font-bold text-blue-600">Pendidikan Umum</h2>
						<div class="flex gap-x-2 w-[60%]">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP" :value="dataPendum.nip">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="Nama">
						</div>
					</div>
				</div>
				
				<form>
					<!-- Grid -->
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Tingkat Pendidikan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refKtpu" id="ktpus" :options="ktpus" keyList="id" namaList="nama" v-model="ktpu"/>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jurusan Pendidikan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refKjur" id="jurusans" :options="jurusans" keyList="id" namaList="nama" v-model="dataPendum.kjur"/>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nama Sekolah</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPendum.nsek">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Tempat</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPendum.tempat">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Negara</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPendum.negara">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nama Kepsek / Direktur / Rektor</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPendum.nkepsek">
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">STTB</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPendum.nsttb">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tsttb">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Pendidikan Terakhir</label>
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
	</LayoutRiwayatPegawai>
</template>