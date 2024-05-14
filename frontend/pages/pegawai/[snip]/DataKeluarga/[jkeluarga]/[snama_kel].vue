<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)

const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const {snip, jkeluarga, snama_kel} = route.params

const dataRefKeluarga = ref({
	nip: '',
	jkeluarga:jkeluarga,
	nama_kel: '',
})

const dataKeluarga = ref({
	nip: '',
	jkeluarga: jkeluarga,
	nama_kel:'',
	nama_ortu: '',
	ktlahir: '',
	tlahir: '',
	tijazah: '',
	tkawin: '',
	stunj: '',
	kjkel:'',
	kkerja: '',
	instansi:'',
	nip_kel: '',
	hubkel: '',
	akhir:'',
	aljalan:'',
	alrt:'',
	alrw:'',
	kodepos:'',
	notelp:'',
	wil:'',
	filename: '',
})

const njkeluarga = ref('')
const ktpus = ref('')
const tkawins = ref([
	{id:"Menikah", nama: "Menikah"},
	{id:"Cerai Hidup", nama: "Cerai Hidup"},
	{id:"Cerai Mati", nama: "Cerai Mati"},
])
const kjkels = ref([
	{id:1, nama: "Laki-laki"},
	{id:2, nama: "Perempuan"},
])

const stunjs = ref([
	{id:'D', nama: "Dapat"},
	{id:'T', nama: "Tidak Dapat"},
])
const ortus = ref([])
const kerjas = ref([])
const refKtpu = ref(null)
const refKerja = ref(null)
const refKjkel = ref(null)
const refOrtu = ref(null)
const refTunj = ref(null)
const method = ref('Simpan')

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
	let wnip = $decodeBase64(snip)
  	const [ktpu, jenis_keluarga, kerja, ortu] = await Promise.all([
  		$fetch('/api/gets/tingkat_pendidikan'),
  		$fetch('/api/gets/jenis_keluarga/'+jkeluarga),
  		$fetch('/api/gets/pekerjaan'),
  		$fetch('/api/posts/riwayat_keluarga', {
  			method: "POST",
  			body: JSON.stringify({
  				nip: wnip,
  				jkeluarga: 1,
  			})
  		}),
  	])

  	var aktpu = _sortBy(ktpu, [function(o){
  		return o.id *1
  	}])
  	
  	ktpus.value = aktpu
  	njkeluarga.value = jenis_keluarga.nama
  	kerjas.value = kerja
  	ortus.value = ortu

  	if(refKtpu){
  		refKtpu.value.reinit()
  	}

  	if(refKerja){
  		refKerja.value.reinit()
  	}
  	
  	if(refOrtu){
  		refOrtu.value.reinit()
  	}
});

const { pending, data, refresh} = await useAsyncData('getdataKeluarga', async() => {
	console.log("CariData")
	if(snip){
		let nip = $decodeBase64(snip)
		let nama_kel = $decodeBase64(snama_kel)
		var result = await $fetch(`/api/gets/riwayat_keluarga/${nip}/${jkeluarga}/${nama_kel}`);

		console.log(result)

		if(result.nip == nip){
			method.value = "Update"

			dataKeluarga.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataKeluarga.value), key);
			})

			dataRefKeluarga.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataRefKeluarga.value), key);
			})
		}
	}
})

const tlahir = computed({
	get(){
		return dayjs(dataKeluarga.value.tlahir).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataKeluarga.value.tlahir = val
	}
})

const tkawin = computed({
	get(){
		return dayjs(dataKeluarga.value.tkawin).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataKeluarga.value.tkawin = val
	}
})

const akhir = computed({
	get(){
		return dataKeluarga.value.akhir == 0 ? false : true
	},
	set(val){
		dataKeluarga.value.akhir = val
	}
})

const aakhir = computed({
	get(){
		return dataKeluarga.value.akhir == 0 ? 'Tidak' : 'Ya'
	},
})

onMounted(() => {
	//refreshNuxtData(["getDataRef", "getdataKeluargaAkhir"])
	refreshNuxtData(["getDataRef", "getdataKeluarga"])
	modalUploadDoc.$onAction(callback, false)
})

const showSpinner = ref(false)
const toast = useToast()

async function simpanData() {
	showSpinner.value = true

	var compacted = _pickBy(dataKeluarga.value);
	compacted.tlahir = compacted.tlahir ? dayjs(dayjs(compacted.tlahir, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tkawin = compacted.tkawin ? dayjs(dayjs(compacted.tkawin, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.jkeluarga = compacted.jkeluarga *1
	compacted.kkerja = compacted.kkerja *1
	compacted.kjkel = compacted.kjkel *1
	compacted.akhir = compacted.akhir == true ? 1 : 0
	
	var dr = dataRefKeluarga.value

	if(method.value == 'Update'){
		var body = {
			"refdata" : dr,
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_keluarga', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_keluarga',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_keluarga',
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		//compacted.akhir = 1
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_keluarga/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_keluarga',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_keluarga',
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
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataKeluarga.value.filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen ${njkeluarga.value}`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen ${njkeluarga.value}`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		formData.append("nama_kel", dataKeluarga.value.nama_kel);
		formData.append("jkeluarga", jkeluarga);
		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_keluarga_"+jkeluarga+ "_" +dataKeluarga.value.nama_kel);
		formData.append("path", 'dokumen/'+nip);

		var {data, error} = await useFetch(`/api/uploads/upload/keluarga/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(error.value){
			toast.add({
		    	id: 'error_post_upload_keluarga',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_keluarga',
		    	description: `Dokumen Keluarga berhasil di Upload`,
		    	timeout: 6000
		  	}) 	
		}

		refreshNuxtData(["getDataRef", "getdataKeluarga"])

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
	<LayoutRiwayatKeluarga>
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
						<h2 class="text-xl font-bold text-blue-600">{{njkeluarga}}</h2>
						<div class="flex gap-x-2 w-[60%]">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP" :value="dataKeluarga.nip">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="Nama">
						</div>
					</div>
				</div>
				
				<form>
					<!-- Grid -->
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nama {{njkeluarga}} {{jkeluarga}}</label>
						</div>
						<!-- End Col -->
						<div class="sm:col-span-9">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKeluarga.nama_kel">
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Tempat Tanggal Lahir</label>
						</div>
						<!-- End Col -->
						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKeluarga.ktlahir">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tlahir">
							</div>
						</div>
						<!-- End Col -->

						<template v-if="[1].includes(jkeluarga*1)">
							<div class="sm:col-span-3 sm:col-start-1">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Status Perkawinan</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<SearchSelect2 ref="refKawin" id="tkawins" :options="tkawins" keyList="id" namaList="nama" v-model="dataKeluarga.hubkel"/>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Tanggal Menikah/Cerai</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<input type="date" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tkawin">
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">{{njkeluarga}} Terakhir</label>
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

							<div class="sm:col-span-3 sm:col-start-1">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Tingkat Pendidikan</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-9">
								<SearchSelect2 ref="refKtpu" id="ktpus" :options="ktpus" keyList="id" namaList="nama" v-model="dataKeluarga.tijazah"/>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Status Tunjangan</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-9">
								<SearchSelect2 ref="refTunj" id="stunjs" :options="stunjs" keyList="id" namaList="nama" v-model="dataKeluarga.stunj"/>
							</div>
						</template>

						<template v-else-if="[2].includes(jkeluarga*1)">
							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Jenis Kelamin</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<SearchSelect2 ref="refKjkel" id="kjkels" :options="kjkels" keyList="id" namaList="nama" v-model="dataKeluarga.hubkel"/>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3 sm:col-start-1">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Tingkat Pendidikan</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-9">
								<SearchSelect2 ref="refKtpu" id="ktpus" :options="ktpus" keyList="id" namaList="nama" v-model="dataKeluarga.tijazah"/>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Status Tunjangan</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-9">
								<SearchSelect2 ref="refTunj" id="stunjs" :options="stunjs" keyList="id" namaList="nama" v-model="dataKeluarga.stunj"/>
							</div>

							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Nama Orang Tua</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-9">
								<SearchSelect2 ref="refOrtu" id="ortus" :options="ortus" keyList="nama_kel" namaList="nama_kel" v-model="dataKeluarga.nama_ortu"/>
							</div>
							<!-- End Col -->
						</template>

						<template v-else>
							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Alamat</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-9">
								<textarea class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKeluarga.aljalan"></textarea>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">RT</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKeluarga.alrt" @keypress="onlyNumber">
							</div>
							<!-- End Col -->
							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">RW</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKeluarga.alrw" @keypress="onlyNumber">
							</div>
							<!-- End Col -->
							
							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Kelurahan/ Kecamatan/ Provinsi</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-4">
								<textarea class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKeluarga.wil"></textarea>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-2">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Kodepos</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKeluarga.kodepos" @keypress="onlyNumber">
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Nomor Telpon</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-9">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKeluarga.notelp" @keypress="onlyNumber">
							</div>
						</template>

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Pekerjaan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refKerja" id="kerjas" :options="kerjas" keyList="id" namaList="nama" v-model="dataKeluarga.kkerja"/>
						</div>
						<!-- End Col -->

						<template v-if="[1].includes(jkeluarga*1)">

							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Instansi Bekerja</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-9">
								<textarea class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKeluarga.instansi"></textarea>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">NIP <i>(*jika PNS)</i></label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-9">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKeluarga.nip_kel">
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
	</LayoutRiwayatKeluarga>
</template>