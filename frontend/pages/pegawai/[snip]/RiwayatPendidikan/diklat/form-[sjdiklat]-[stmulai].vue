<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)

const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const {snip, sjdiklat, stmulai} = route.params

const dataRefDiklat = ref({
	nip: '',
	jdiklat:'',
	tmulai: '',
})

const dataDiklat = ref({
	nip: '',
	jdiklat: '',
	kdiklat:'',
	ndiklat: '',
	tempat: '',
	pan: '',
	sebagai: '',
	angkatan: '',
	tmulai:'',
	takhir: '',
	tahun:'',
	jam:'',
	nsttpp: '',
	tsttpp: '',
	akhir:'',
	filename: '',
	rumpun_id: '',
})

const njdiklat = ref('')
const jenis_diklats = ref([])
const diklats = ref([])
const rumpuns = ref([])
const refJnsDiklat = ref(null)
const refDiklat = ref(null)
const refRumpun = ref(null)
const method = ref('Simpan')

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
	let jdiklat = $decodeBase64(sjdiklat)
  	const [jenis_diklat, diklat, rumpun] = await Promise.all([
  		$fetch('/api/gets/jenis_diklat'),
    	$fetch('/api/posts/diklat', {
    		method: "POST",
    		body: JSON.stringify({
    			status: [1],
    			jdiklat: 1,
    		})
    	}),
    	$fetch('/api/gets/rumpun_jabatan'),
  	])

  	//jenis_diklats.value = jenis_diklat
  	let n = _split(jdiklat, ',')
  	//console.log(n.length)
  	if(n.length > 1){
  		jenis_diklats.value = _filter(jenis_diklat, function(o){
  			return !n.includes(o.id.toString()) && o.id != 1
  		});
  	}else{
		jenis_diklats.value = _filter(jenis_diklat, function(o){
  			return o.id == jdiklat
  		});
  	}

  	rumpuns.value = rumpun
  	diklats.value = diklat  	

  	if(refJnsDiklat){
  		refJnsDiklat.value.reinit()
  	}

	if(refRumpun){
		refRumpun.value.reinit()
	}

	if(refDiklat){
		refDiklat.value.reinit()
	}
});

const { pending, data, refresh} = await useAsyncData('getdataDiklat', async() => {
	console.log("CariData")
	if(snip){
		let nip = $decodeBase64(snip)
		let tmulai = $decodeBase64(stmulai)
		let jdiklat = $decodeBase64(sjdiklat)

		var result = await $fetch(`/api/gets/riwayat_diklat/${nip}/${jdiklat}/${tmulai}`);
		if(result.nip == nip){
			method.value = "Update"

			dataDiklat.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataDiklat.value), key);
			})

			dataRefDiklat.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataRefDiklat.value), key);
			})
		}else{
			dataDiklat.value.jdiklat = jdiklat
			dataRefDiklat.value.jdiklat = jdiklat
		}
	}
})

const jdiklat = computed({
	get(){
		return dataDiklat.value.jdiklat
	},
	set(val){
		dataDiklat.value.jdiklat = val
	}
})

const ndiklat = computed({
	get(){
		return dataDiklat.value.kdiklat
	},
	set(val){
		dataDiklat.value.kdiklat = val

		var result = _find(diklats.value, {id: val})

		dataDiklat.value.ndiklat = result.nama
	}
})

const tmulai = computed({
	get(){
		return dayjs(dataDiklat.value.tmulai).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataDiklat.value.tmulai = val
	}
})

const takhir = computed({
	get(){
		return dayjs(dataDiklat.value.takhir).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataDiklat.value.takhir = val
	}
})

const tsttpp = computed({
	get(){
		return dayjs(dataDiklat.value.tsttpp).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataDiklat.value.tsttpp = val
	}
})

const akhir = computed({
	get(){
		return dataDiklat.value.akhir == 0 ? false : true
	},
	set(val){
		dataDiklat.value.akhir = val
	}
})

const aakhir = computed({
	get(){
		return dataDiklat.value.akhir == 0 ? 'Tidak' : 'Ya'
	},
})

onMounted(() => {
	//refreshNuxtData(["getDataRef", "getdataDiklatAkhir"])
	refreshNuxtData(["getDataRef", "getdataDiklat"])
	modalUploadDoc.$onAction(callback, false)
})

const showSpinner = ref(false)
const toast = useToast()

async function simpanData() {
	showSpinner.value = true

	var compacted = _pickBy(dataDiklat.value);
	compacted.tsttpp = compacted.tsttpp ? dayjs(dayjs(compacted.tsttpp, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tmulai = compacted.tmulai ? dayjs(dayjs(compacted.tmulai, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.takhir = compacted.takhir ? dayjs(dayjs(compacted.takhir, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.jdiklat = compacted.jdiklat *1
	compacted.kdiklat = compacted.kdiklat *1
	compacted.jam = compacted.jam *1
	compacted.akhir = compacted.akhir == true ? 1 : 0

	console.log(compacted)
	
	var dr = dataRefDiklat.value

	if(method.value == 'Update'){
		var body = {
			"refdata" : dr,
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_diklat', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_diklat',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_diklat',
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		//compacted.akhir = 1
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_diklat/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_diklat',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_diklat',
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
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataDiklat.value.filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen ${njdiklat.value}`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen ${njdiklat.value}`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		formData.append("tmulai", tmulai.value);
		formData.append("jdiklat", jdiklat);
		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_diklat_"+jdiklat+ "_" +tmulai.value);
		formData.append("path", 'dokumen/'+nip);

		var {data, error} = await useFetch(`/api/uploads/upload/diklat/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(error.value){
			toast.add({
		    	id: 'error_post_upload_diklat',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_diklat',
		    	description: `Dokumen Diklat berhasil di Upload`,
		    	timeout: 6000
		  	}) 	
		}

		refreshNuxtData(["getDataRef", "getdataDiklat"])

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
const nsttpl = ref([
	"STTPP", //0
	"STTPP", //1
	"STTPP", //2
	"STTPP", //3
	"Piagam", //4
	"Piagam",
	"Piagam",
	"STTPP",
	"Piagam",
	"STTPP",
	"Sertifikat",
])

const ntmulai = ref([
	"Tanggal Mulai", //0
	"Tanggal Mulai", //1
	"Tanggal Mulai", //2
	"Tanggal Mulai", //3
	"Tanggal Mulai", //4
	"Tanggal Mulai", //5
	"Tanggal Mulai", //6
	"Bulan /Tahun Dilaksanakan", //7
	"Bulan /Tahun Dilaksanakan", //8
	"Bulan /Tahun dibuat", //9
	"Tanggal Ujian", //10
])
</script>
<template>
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
						<h2 class="text-xl font-bold text-blue-600">Diklat</h2>
						<div class="flex gap-x-2 w-[60%]">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP" :value="dataDiklat.nip">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="Nama">
						</div>
					</div>
				</div>
				
				<form>
					<!-- Grid -->
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jenis</label>
						</div>
						<div class="sm:col-span-9">
							<SearchSelect2 ref="refJnsDiklat" id="jenis_diklats" :options="jenis_diklats" keyList="id" namaList="nama" v-model="jdiklat"/>
						</div>
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nama</label>
						</div>
						<!-- End Col -->
						<template v-if="jdiklat == '1'">
							<div class="sm:col-span-9">
								<SearchSelect2 ref="refDiklat" id="diklats" :options="diklats" keyList="id" namaList="nama" v-model="ndiklat"/>
							</div>
							<!-- End Col -->
						</template>
						<template v-else>
							<div class="sm:col-span-9">
								<textarea class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataDiklat.ndiklat"></textarea>
							</div>
							<!-- End Col -->
						</template>

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Tempat {{jdiklat == '9' ? 'Pembuatan' : (jdiklat == '10' ? 'Ujian' : 'Pelaksanaan')}}</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataDiklat.tempat">
						</div>
						<!-- End Col -->

						<template v-if="jdiklat != '9'">
							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">{{jdiklat == '8' ? 'Peran Sebagai' : 'Panitia/Penyelenggara'}}</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-9">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataDiklat.pan">
							</div>
							<!-- End Col -->
						</template>

						<template v-if="[1,2,3,7].includes(jdiklat*1)">
							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Angkatan</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataDiklat.angkatan">
							</div>
							<!-- End Col -->
						</template>

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">{{ntmulai[jdiklat]}}</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<div class="sm:flex gap-2">
								<input type="date" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tmulai">
							</div>
						</div>
						<!-- End Col -->
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">{{jdiklat == '10' ? 'Berlaku Sampai' : 'Tanggal Selesai'}}</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<div class="sm:flex gap-2">
								<input type="date" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="takhir">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jumlah Jam</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataDiklat.jam" @keypress="onlyNumber">
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">{{nsttpl[jdiklat]}}</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="test" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataDiklat.nsttpp">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tsttpp">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Rumpun Diklat</label>
						</div>
						<!-- End Col -->
						<div class="sm:col-span-9">
							<SearchSelect2 ref="refRumpun" id="rumpuns" :options="rumpuns" keyList="id" namaList="nama" v-model="dataDiklat.rumpun_id"/>
						</div>
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
	</LayoutRiwayatPendidikan>
</template>