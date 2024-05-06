<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)

const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const {snip, stmtPi} = route.params

const dataRefPi = ref({
	nip: '',
	tmtPi: '',
})

const dataPi = ref({
	nip: '',
	tmtPi:'',
	jenisPi: '',
	instansiIndukLama: '',
	instansiIndukBaru: '',
	skBknNomor: '',
	skBknTanggal: '',
	skAsalProvNomor: '',
	skAsalProvTanggal: '',
	skAsalTanggal: '',
	skAsalNomor:'',
	skTujuanNomor:'',
	skTujuanTanggal: '',
	filename: '',
	kpknBaru:'',
})

const instansis = ref([])
const jnsPis = ref([
	{id: 'DPK/DPB', nama: 'DPK/DPB'},
	{id: 'Penugasan', nama: 'Penugasan'},
	{id: 'Pengalihan', nama: 'Pengalihan'},
	{id: 'Pindah Instansi', nama: 'Pindah Instansi'},
])
const kpkns = ref([])

const refJnsPi = ref(null)
const refInstansiTujuan = ref(null)
const refInstansiAsal = ref(null)
const refKpkn = ref(null)

const method = ref('Simpan')

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
	
  	const [instansi, kpkn] = await Promise.all([
    	$fetch('/api/gets/instansi'),
    	$fetch('/api/gets/kppn'),
  	])

  	//jnsPis.value = jenisPi
  	//jenisJenjangs.value = jenisJenjang
  	instansis.value = instansi
  	kpkns.value = kpkn
  	
  	if(refInstansiAsal){
		refInstansiAsal.value.reinit()
	}

	if(refInstansiTujuan){
		refInstansiTujuan.value.reinit()
	}

	if(refKpkn){
		refKpkn.value.reinit()
	}
});

const { pending, data, refresh} = await useAsyncData('getdataPi', async() => {
	console.log("CariData")
	if(snip){
		let nip = $decodeBase64(snip)
		let tmtPi = $decodeBase64(stmtPi)
		var result = await $fetch(`/api/gets/riwayat_pindah_instansi/${nip}/${tmtPi}`);

		if(result.nip == nip){
			method.value = "Update"

			dataPi.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataPi.value), key);
			})

			dataRefPi.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataRefPi.value), key);
			})
		}
	}
})

const skBknTanggal = computed({
	get(){
		return dayjs(dataPi.value.skBknTanggal).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataPi.value.skBknTanggal = val
	}
})

const skAsalProvTanggal = computed({
	get(){
		return dayjs(dataPi.value.skAsalProvTanggal).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataPi.value.skAsalProvTanggal = val
	}
})

const skAsalTanggal = computed({
	get(){
		return dayjs(dataPi.value.skAsalTanggal).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataPi.value.skAsalTanggal = val
	}
})

const skTujuanTanggal = computed({
	get(){
		return dayjs(dataPi.value.skTujuanTanggal).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataPi.value.skTujuanTanggal = val
	}
})

const tmtPi = computed({
	get(){
		return dayjs(dataPi.value.tmtPi).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataPi.value.tmtPi = val
	}
})

onMounted(() => {
	//refreshNuxtData(["getDataRef", "getdataPiAkhir"])
	refreshNuxtData(["getDataRef", "getdataPi"])
	modalUploadDoc.$onAction(callback, false)
})

const showSpinner = ref(false)
const toast = useToast()

async function simpanData() {
	showSpinner.value = true

	var compacted = _pickBy(dataPi.value);
	compacted.skBknTanggal = compacted.skBknTanggal ? dayjs(dayjs(compacted.skBknTanggal, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.skAsalProvTanggal = compacted.skAsalProvTanggal ? dayjs(dayjs(compacted.skAsalProvTanggal, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.skAsalTanggal = compacted.skAsalTanggal ? dayjs(dayjs(compacted.skAsalTanggal, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.skTujuanTanggal = compacted.skTujuanTanggal ? dayjs(dayjs(compacted.skTujuanTanggal, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tmtPi = compacted.tmtPi ? dayjs(dayjs(compacted.tmtPi, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null

	var dr = dataRefPi.value

	if(method.value == 'Update'){
		var body = {
			"refdata" : dr,
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_pindah_instansi', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_pi',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_pi',
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		//compacted.akhir = 1
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_pindah_instansi/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_pi',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_pi',
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
		}
	}
  	showSpinner.value = false
  	refreshNuxtData(["getDataRef", "getdataPi"])
}

const doUploadDoc = async() => {
	let nip = $decodeBase64(snip)
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataPi.value.filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen Pindah Instansi`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen Pindah Instansi`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		formData.append("tmtPi", tmtPi.value);
		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_pindahInstansi_"+tmtPi.value);
		formData.append("path", 'dokumen/'+nip);

		var {data, error} = await useFetch(`/api/uploads/upload/pindah_instansi/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(error.value){
			toast.add({
		    	id: 'error_post_upload_pindah_instansi',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_pindah_instansi',
		    	description: `Dokumen Pindah Instansi berhasil di Upload`,
		    	timeout: 6000
		  	}) 	
		}

		refreshNuxtData(["getDataRef", "getdataPi"])

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
						<h2 class="text-xl font-bold text-blue-600">Pindah Instansi</h2>
						<div class="flex gap-x-2 w-[60%]">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP" :value="dataPi.nip">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="Nama">
						</div>
					</div>
				</div>
				
				<form>
					<!-- Grid -->
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jenis Pindah Instansi</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refJnsPi" id="jnsPis" :options="jnsPis" keyList="id" namaList="nama" v-model="dataPi.jenisPi"/>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Instansi Tujuan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refInstansiTujuan" id="instansiTujuan" :options="instansis" keyList="id" namaList="nama" v-model="dataPi.instansiIndukLama"/>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Instansi Asal</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refInstansiAsal" id="instansiAsal" :options="instansis" keyList="id" namaList="nama" v-model="dataPi.instansiIndukBaru"/>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">SK BKN</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPi.skBknNomor">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="skBknTanggal">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">SK Provinsi</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPi.skAsalProvNomor">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="skAsalProvTanggal">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Surat/Rekomendasi Melepas (Instansi Asal)</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPi.skAsalNomor">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="skAsalTanggal">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Surat/Rekomendasi Menerima (Instansi Tujuan)</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPi.skTujuanNomor">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="skTujuanTanggal">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">TMT Pindah Instansi</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<div class="sm:flex gap-2">
								<input type="date" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tmtPi">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">KPKN</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refKpkn" id="kpkns" :options="kpkns" keyList="id" namaList="nama" v-model="dataPi.kpknBaru"/>
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