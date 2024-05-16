<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)

const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const showSpinner = ref(false)
const toast = useToast()
const {snip, stmulai} = route.params

const dataRefAngkaKredit = ref({
	nip: '',
	tmulai: '',
})

const dataAngkaKredit = ref({
	nip: '',
	jns: '',
	nsk:'',
	tsk: '',
	thn: '',
	kjab: '',
	tmulai: '',
	tselesai: '',
	utama: 0,
	tambahan: 0,
	total: 0,
	filename: '',
	idSync:'',
})

const jabfungs = ref([])
const jnsAks = ref([
	{id: 'Pertama', nama: 'Pertama'},
	{id: 'Lanjutan', nama: 'Lanjutan'},
	{id: 'Integrasi', nama: 'Integrasi'},
	{id: 'Konversi', nama: 'Konversi'},
])
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

const refJabfung = ref(null)
const refJnsAK = ref(null)
const method = ref('Simpan')

const refDataBkn = ref({
	"id": "idSync",
    "kreditBaruTotal": "total",
    "kreditPenunjangBaru": "tambahan",
    "kreditUtamaBaru": "utama",
    "nomorSk": "nsk",
    "tanggalSk": "tsk",
})

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
	//loading.value = true
	let nip = $decodeBase64(snip)
  	const [jabfung] = await Promise.all([
  		$fetch(`/api/posts/riwayat_jabatan/`, {
  			method: 'POST',
  			body: JSON.stringify({
  				nip: nip,
  				jnsjab: "2",
  			})
  		}),
  	])

  	jabfungs.value = jabfung
  	//jenisJenjangs.value = jenisJenjang
  	
  	if(refJabfung){
		refJabfung.value.reinit()
	}

	if(refJnsAK){
		refJnsAK.value.reinit()
	}
});

const { pending, data, refresh} = await useAsyncData('getdataAngkaKreditAkhir', async() => {
	if(snip){
		let nip = $decodeBase64(snip)
		let tmulai = $decodeBase64(stmulai)
		var result = await $fetch(`/api/gets/riwayat_angka_kredit/${nip}/${tmulai}`);
		console.log(result)

		if(result.nip == nip){
			method.value = "Update"

			dataAngkaKredit.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataAngkaKredit.value), key);
			})

			dataRefAngkaKredit.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataRefAngkaKredit.value), key);
			})
		}
	}
})

const tsk = computed({
	get(){
		return dayjs(dataAngkaKredit.value.tsk).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataAngkaKredit.value.tsk = val
	}
})

const tmulai = computed({
	get(){
		return dayjs(dataAngkaKredit.value.tmulai).format("YYYY-MM").toString()
	},
	set(val){
		dataAngkaKredit.value.tmulai = val
	}
})

const tselesai = computed({
	get(){
		return dayjs(dataAngkaKredit.value.tselesai).format("YYYY-MM").toString()
	},
	set(val){
		dataAngkaKredit.value.tselesai = val
	}
})

const total = computed({
	get(){
		dataAngkaKredit.value.total = (dataAngkaKredit.value.utama *1) + (dataAngkaKredit.value.tambahan *1)
		return dataAngkaKredit.value.total;
	}
})

onMounted(() => {
	//refreshNuxtData(["getDataRef", "getdataAngkaKreditAkhir"])
	refreshNuxtData(["getDataRef", "getdataAngkaKreditAkhir"])
	modalUploadDoc.$onAction(callback, false)
})

function resetJab() {
	//refreshNuxtData(["getDataRef", "getdataAngkaKreditAkhir", "getListdataAngkaKredit"])
}

async function simpanData(lvl) {
	showSpinner.value = true

	var compacted = _pickBy(dataAngkaKredit.value);
	compacted.tsk = compacted.tsk ? dayjs(dayjs(compacted.tsk, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tmulai = compacted.tmulai ? dayjs(dayjs(compacted.tmulai, "YYYY-MM")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tselesai = compacted.tselesai ? dayjs(dayjs(compacted.tselesai, "YYYY-MM")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null

	compacted.utama = compacted.utama ? compacted.utama *1 : 0
	compacted.tambahan = compacted.tambahan ? compacted.utama *1 : 0
	compacted.total = compacted.total ? compacted.total *1 : 0

	if(lvl == 2){
		let body = compacted

	    let nip = $decodeBase64(snip)
	    var {data, error} = await useFetch('/api/puts/singkronisasi_bkn/putsDataBKNAK/'+nip, {
	        method: 'POST',
	        body: body,
	    })

	    console.log(data.value)
	    console.log(error.value)

	    if(data.value.success){
	    	compacted.idSync = data.value?.mapData?.rwAngkaKreditId;
	    }

	    if(error.value){
	        toast.add({
	            id: 'error_put_singkronisasi_ak',
	            description: error.value.data.data,
	            timeout: 6000,
	            color: 'red',
	        })  
	    }else{
	        toast.add({
	            id: 'success_post_singkronisasi_ak',
	            description: data.value.success == true ? "Data BKN Update Success |" : "Data BKN Update Failed |"  +data.value.message,
		    	timeout: 6000,
		    	color:  data.value.success == "1" ? 'green' : 'red',
	        })  
	    }
	}

	var dr = dataRefAngkaKredit.value
	//dr.kgolru = dr.kgolru *1

	if(method.value == 'Update'){
		var body = {
			"refdata" : dr,
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_angka_kredit', {
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

		var {data, error} = await useFetch('/api/posts/riwayat_angka_kredit/new', {
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
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataAngkaKredit.value.filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen Angka Kredit`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen Angka Kredit`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		let xmulai = $decodeBase64(stmulai)
		formData.append("tmulai", xmulai);
		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_angka_kredit"+tmulai.value);
		formData.append("path", 'dokumen/'+nip);

		var {data, error} = await useFetch(`/api/uploads/upload/ak/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(error.value){
			toast.add({
		    	id: 'error_post_upload_ak',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_ak',
		    	description: `Dokumen Angka Kredit berhasil di Upload`,
		    	timeout: 6000
		  	}) 	
		}

		//upload to BKN

		if(dataAngkaKredit.value.idSync != ""){
			let formDataBkn = new FormData();
			formDataBkn.append("file", fileBlob.value[0]);
			formDataBkn.append("idSync", dataAngkaKredit.value.idSync);
			formDataBkn.append("id_ref_dokumen", 879);

			var {data, error} = await useFetch(`/api/uploads/upload/dokRwBkn/${nip}`, {
				method: 'POST',
				body: formDataBkn,
			});	

			console.log(error.value)
			console.log(data.value)

			if(error.value){
		        toast.add({
		            id: 'error_post_upload_ak_bkn',
		            description: error.value.data.data,
		            timeout: 6000,
		            color: 'red',
		        })  
		    }else{
		        toast.add({
		            id: 'success_post_upload_ak_bkn',
		            description: data.value.data.code == 1 ? "Dokumen Angka Kredit BKN Upload Success |" +data.value.data.message : "Dokumen Angka Kredit BKN Upload Failed |"  +data.value.data.message,
			    	timeout: 6000,
			    	color:  data.value.data.code == 1 ? 'green' : 'red',
		        })  
		    }
		}

		//end upload to BKN

		showSpinner.value = false
		refreshNuxtData(["getDataRef", "getdataAngkaKreditAkhir"])
	}
}

function onlyNumber ($event) {
   	//console.log($event.keyCode); //keyCodes value
   	let keyCode = ($event.keyCode ? $event.keyCode : $event.which);
   	if ((keyCode < 48 || keyCode > 57) && keyCode !== 46) { // 46 is dot
    	$event.preventDefault();
   	}
}

const doRefresh = () => {
	refreshNuxtData(["getDataRef", "getdataAngkaKreditAkhir"])
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
						<h2 class="text-xl font-bold text-blue-600">Angka Kredit</h2>
						<div class="flex gap-x-2 w-[60%]">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP" :value="dataAngkaKredit.nip">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="Nama">
						</div>
					</div>
				</div>
				
				<form>
					<!-- Grid -->
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jabatan Fungsional</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refJabfung" id="jabfungs" :options="jabfungs" keyList="kjab" namaList="njab" v-model="dataAngkaKredit.kjab"/>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">SK Angka Kredit</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataAngkaKredit.nsk">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tsk">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jenis Angka Kredit</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-7">
							<SearchSelect2 ref="refJnsAK" id="jnsAks" :options="jnsAks" keyList="id" namaList="nama" v-model="dataAngkaKredit.jns"/>
						</div>
						<template v-if="dataAngkaKredit.jns == 'Konversi'">
							<div class="sm:col-span-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataAngkaKredit.thn" placeholder="Tahun Konversi" required>
							</div>
						</template>
						<!-- End Col -->

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Priode Penilaian</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="month" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tmulai">
								<label class="inline-block text-sm text-gray-800 mt-2.5">s/d</label>
								<input type="month" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tselesai">
							</div>
						</div>
						<!-- End Col -->

						<template v-if="(_includes(['Pertama', 'Lanjutan'], dataAngkaKredit.jns))">
							<div class="sm:col-span-3 sm:col-start-1">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Angka Kredit Utama</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<div class="sm:flex gap-2">
									<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataAngkaKredit.utama" @keypress="onlyNumber">
								</div>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Angka Kredit Penunjang</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<div class="sm:flex gap-2">
									<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataAngkaKredit.tambahan" @keypress="onlyNumber">
								</div>
							</div>
							<!-- End Col -->
						</template>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Total Angka Kredit</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="total">
							</div>
						</div>
						<!-- End Col -->
					</div>
					<div class="mt-5 grid sm:grid-cols-12 gap-x-2">
						<div class="sm:col-span-7 sm:col-start-4">
							<div class="sm:flex gap-2">
								<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none" @click="doRefresh">
					  				Batal
								</button>
								<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="simpanData(0)">
					  				{{method}} Data [0]
								</button>	
								<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="simpanData(1)">
					  				{{method}} Data [1]
								</button>	
								<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="simpanData(2)">
					  				{{method}} Data [2]
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
	</LayoutRiwayatPegawai>
</template>