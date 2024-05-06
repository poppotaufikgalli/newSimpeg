<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)

const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const toast = useToast()
const {snip, stanggalAwal} = route.params

const showSpinner = ref(false)

const dataRefPMK = ref({
	nip: '',
	tanggalAwal: '',
})

const dataPMK = ref({
	nip: '',
	dinilai: '',
	tanggalAwal: '',
	tanggalAkhir: '',
	masaKerjaBulan: 0,
	masaKerjaTahun: 0,
	ptetap: '',
	nomorSk:'',
	tanggalSk: '',
	ninstansi: '',
	pengalaman: '',
	nomorBkn: '',
	tanggalBkn: '',
	filename: '',
})

const jabfungs = ref([])
const jnsInstansis = ref([
	{id: 100, nama: 'Pemerintah'},
	{id: 50, nama: 'Swasta'},
])
const eselons = ref([])
const pejabats = ref([])
const unitKerja = ref([])

const refPejabat = ref(null)
const refJnsInstansi = ref(null)
const method = ref('Simpan')

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
	//loading.value = true
	let nip = $decodeBase64(snip)
  	const [pejabat, jabfung, jenisJenjang] = await Promise.all([
  		$fetch('/api/gets/pejabat'),
  	])

  	pejabats.value = pejabat
  	
  	if(refPejabat){
		refPejabat.value.reinit()
	}

  	if(refJnsInstansi){
		refJnsInstansi.value.reinit()
	}
});

const { pending, data, refresh} = await useAsyncData('getdataPMK', async() => {
	if(snip){
		let nip = $decodeBase64(snip)
		let tanggalAwal = $decodeBase64(stanggalAwal)
		var result = await $fetch(`/api/gets/riwayat_pmk/${nip}/${tanggalAwal}`);
		console.log(result)

		if(result.nip == nip){
			method.value = "Update"

			dataPMK.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataPMK.value), key);
			})

			dataRefPMK.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataRefPMK.value), key);
			})
		}
	}
})

const tanggalSk = computed({
	get(){
		return dayjs(dataPMK.value.tanggalSk).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataPMK.value.tanggalSk = val
	}
})

const tanggalBkn = computed({
	get(){
		return dayjs(dataPMK.value.tanggalBkn).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataPMK.value.tanggalBkn = val
	}
})

const tanggalAwal = computed({
	get(){
		return dayjs(dataPMK.value.tanggalAwal).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataPMK.value.tanggalAwal = val
	}
})

const tanggalAkhir = computed({
	get(){
		return dayjs(dataPMK.value.tanggalAkhir).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataPMK.value.tanggalAkhir = val
	}
})

const masaKerjaTahun = computed({
	get(){
		const date1 = dayjs(dataPMK.value.tanggalAkhir)
		const date2 = dayjs(dataPMK.value.tanggalAwal)
		var result_1 = date1.diff(date2, 'month')
		var nilai = dataPMK.value.dinilai == 50 ? 0.5 :1
		var result = Math.floor((result_1 * nilai) / 12)
		dataPMK.value.masaKerjaTahun = result 
		return result
	}
})

const masaKerjaBulan = computed({
	get(){
		const date1 = dayjs(dataPMK.value.tanggalAkhir)
		const date2 = dayjs(dataPMK.value.tanggalAwal)
		var result_1 = date1.diff(date2, 'month')
		var nilai = dataPMK.value.dinilai == 50 ? 0.5 :1
		var result = Math.floor((result_1 * nilai) % 12)
		dataPMK.value.masaKerjaBulan = result 
		return result
	}
})

onMounted(() => {
	//refreshNuxtData(["getDataRef", "getdataPMK"])
	refreshNuxtData(["getDataRef", "getdataPMK"])
	modalUploadDoc.$onAction(callback, false)
})

async function simpanData() {
	showSpinner.value = true

	var compacted = _pickBy(dataPMK.value);
	compacted.tanggalAwal = compacted.tanggalAwal ? dayjs(dayjs(compacted.tanggalAwal, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tanggalAkhir = compacted.tanggalAkhir ? dayjs(dayjs(compacted.tanggalAkhir, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tanggalSk = compacted.tanggalSk ? dayjs(dayjs(compacted.tanggalSk, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tanggalBkn = compacted.tanggalBkn ? dayjs(dayjs(compacted.tanggalBkn, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null

	compacted.dinilai = compacted.dinilai *1
	//compacted.tambahan = compacted.tambahan *1
	//compacted.total = compacted.total *1

	var dr = dataRefPMK.value
	//dr.kgolru = dr.kgolru *1

	if(method.value == 'Update'){
		var body = {
			"refdata" : dr,
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_pmk', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_ak',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_ak',
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		//compacted.akhir = 1
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_pmk/new', {
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
  	refreshNuxtData(["getDataRef"])
}

const doUploadDoc = async() => {
	let nip = $decodeBase64(snip)
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataPMK.value.filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen Penyesuaian Masa Kerja`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen Penyesuaian Masa Kerja`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		formData.append("tanggalAwal", tanggalAwal.value);
		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_pmk"+tanggalAwal.value);
		formData.append("path", 'dokumen/'+nip);

		var {data, error} = await useFetch(`/api/uploads/upload/pmk/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(error.value){
			toast.add({
		    	id: 'error_post_upload_pmk',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_pmk',
		    	description: `Dokumen Penyesuaian Masa Kerja berhasil di Upload`,
		    	timeout: 6000
		  	}) 	
		}

		refreshNuxtData(["getDataRef", "getdataPMK"])

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
						<h2 class="text-xl font-bold text-blue-600">Penyesuaian Masa Kerja</h2>
						<div class="flex gap-x-2 w-[60%]">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP" :value="dataPMK.nip">
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
							<SearchSelect2 ref="refPejabat" id="pejabats" :options="pejabats" keyList="kpej" namaList="npej" v-model="dataPMK.kpej"/>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">SK PMK</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPMK.nomorSk">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tanggalSk">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Pertek BKN</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPMK.nomorBkn">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tanggalBkn">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jenis Instansi</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refJnsInstansi" id="jnsInstansis" :options="jnsInstansis" keyList="id" namaList="nama" v-model="dataPMK.dinilai"/>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nama Instansi</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPMK.ninstansi">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Rentang Tanggal Kerja</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="date" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tanggalAwal">
								<label class="inline-block text-sm text-gray-800 mt-2.5">s/d</label>
								<input type="date" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tanggalAkhir">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Masa Kerja</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-4">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="masaKerjaTahun" disabled>
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="masaKerjaBulan" disabled>
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