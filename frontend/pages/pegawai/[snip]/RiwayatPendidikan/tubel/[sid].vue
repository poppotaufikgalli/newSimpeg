<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';
const dayjs = useDayjs()
import customParseFormat from "dayjs/plugin/customParseFormat";
dayjs.extend(customParseFormat)

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)

const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()

const {snip, sid} = route.params
const showSpinner = ref(false)
const toast = useToast()

const dataRefTubel = ref({
	id: '',
	nip: '',
	ktpu: '',
})

const dataTubel = ref({
	id: 0,
	nip: '',
	ktpu: '',
	kjur:'',
	njur: '',
	nsek: '',
	tempat: '',
	nsk: '',
	tsk: '',
	tahun: '',
	aktif: 0,
	filename: '',
})

const jurusans = ref([])
const jurusans_bkns = ref([])
const ktpus = ref([])
const refKjur = ref(null)
const refKtpu = ref(null)
const method = ref('Simpan')
const apiUrlPendidikan = ref("/api/gets/pendidikan/filter")

const { pending, data, refresh} = await useAsyncData('getdataTubel', async() => {
	console.log("CariData")
	if(snip){
		let nip = $decodeBase64(snip)
		let id = $decodeBase64(sid)
		var result = await $fetch(`/api/gets/riwayat_tubel/${nip}/${id}`);

		console.log(result)

		if(result.nip == nip){
			method.value = "Update"

			dataTubel.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataTubel.value), key);
			})

			dataRefTubel.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataRefTubel.value), key);
			})
		}
	}

  	const [ktpu] = await Promise.all([
    	$fetch('/api/gets/tingkat_pendidikan'),
    	$fetch('/api/gets/pendidikan'),
  	])

  	var aktpu = _sortBy(ktpu, [function(o){
  		return o.id *1
  	}])

  	ktpus.value = aktpu
  	//jurusans.value = jurusan

	if(refKtpu){
		refKtpu.value.reinit()
	}
});

const ktpu = computed({
	get(){
		return dataTubel.value.ktpu
	},
	set(val){
		console.log("ktpu value")
		dataTubel.value.ktpu = val
	}
})

const tsk = computed({
	get(){
		return dayjs(dataTubel.value.tsk).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataTubel.value.tsk = val
	}
})

const aaktif = computed({
	get(){
		return dataTubel.value.aktif == 0 ? 'Tidak' : 'Ya'
	},
})

const aktif = computed({
	get(){
		return dataTubel.value.aktif == 0 ? false : true
	},
	set(val){
		dataTubel.value.aktif = val
	}
})

onMounted(() => {
	refreshNuxtData(["getdataTubel"]) 
	modalUploadDoc.$onAction(callback, false)
})

async function simpanData() {
	showSpinner.value = true

	var compacted = _pickBy(dataTubel.value);
	compacted.tsk = compacted.tsk ? dayjs(dayjs(compacted.tsk, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.aktif = compacted.aktif == true ? 1 : 0
	compacted.tahun = compacted.tahun *1
	
	var dr = dataRefTubel.value

	if(method.value == 'Update'){
		var body = {
			"refdata" : dr,
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_tubel', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_tubel',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_tubel',
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_tubel/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_tubel',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_tubel',
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
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataTubel.value.filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen Tugas Belajar`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen Tugas Belajar`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		formData.append("id", dataTubel.value.id);
		formData.append("ktpu", dataTubel.value.ktpu);
		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_tubel_"+dataTubel.value.id+"_"+dataTubel.value.ktpu);
		formData.append("path", 'dokumen/'+nip);

		var {data, error} = await useFetch(`/api/uploads/upload/tubel/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(error.value){
			toast.add({
		    	id: 'error_post_upload_tubel',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_tubel',
		    	description: `Dokumen Tugas Belajar berhasil di Upload`,
		    	timeout: 6000
		  	}) 	
		}

		refreshNuxtData(["getdataTubel"])

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
	<AppLoadingSpinner :show="showSpinner" />
	<LayoutRiwayatPendidikan>
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
						<h2 class="text-xl font-bold text-blue-600">Tugas Belajar</h2>
						<div class="flex gap-x-2 w-[60%]">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP" :value="dataTubel.nip">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="Nama">
						</div>
					</div>
				</div>
				
				<form>
					<!-- Grid -->
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Surat Keputusan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataTubel.nsk">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tsk">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Tahun</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataTubel.tahun" @keypress="onlyNumber">
						</div>
						<!-- End Col -->

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
							<!--<SearchSelect3 ref="refKjur" id="jurusans_simpegs" :apiUrl="apiUrlPendidikan" keyList="id" namaList="nama" v-model="dataTubel.kjur"/>-->
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataTubel.njur">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nama Sekolah</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataTubel.nsek">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Tempat</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataTubel.tempat">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Sedang Tubel</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<div class="flex mt-2">
								  	<input type="checkbox" class="shrink-0 mt-0.5 border-gray-200 rounded text-blue-600 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none" id="hs-default-checkbox" v-model="aktif">
								  	<label for="hs-default-checkbox" class="text-sm text-gray-500 ms-3">{{aaktif}}</label>
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
		  			</div>
				</form>
	  		</div>
	  		<!-- End Card -->
		</div>
		<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2 mt-6" v-if="method == 'Update'">
			<div class="mb-8">
				<h2 class="text-xl font-bold text-blue-600">Dokumen Tugas Belajar</h2>
			</div>
			<div class="grid sm:grid-cols-4 gap-2 gap-2.5">
				<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc()">
					SK Tugas Belajar
				</button>
			</div>
		</div>
		<!-- End Card Section -->
	</LayoutRiwayatPendidikan>
</template>