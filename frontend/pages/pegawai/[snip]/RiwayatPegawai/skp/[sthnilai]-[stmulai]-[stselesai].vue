<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)

const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const {snip, sthnilai, stmulai} = route.params

const dataRefSkp = ref({
	nip: '',
	thnilai:'',
	tmulai: '',
	tselesai: '',
})

const dataSkp = ref({
	nip: '',
	thnilai:'',
	tmulai: '',
	tselesai: '',
	hasil_kinerja: '',
	perilaku_kerja: '',
	kuadran_kinerja: '',
	pej_pns: 1,
	pej_nik_nip: '',
	pej_nama: '',
	pej_kgolru:'',
	pej_nunker:'',
	pej_njab: '',
	filename: '',
})

const ratings = ref([
	{id: '1', nama: 'Diatas Ekspektasi'},
	{id: '2', nama: 'Sesuai Ekspektasi'},
	{id: '3', nama: 'Dibawah Ekspektasi'},
])
const kuadrankinerjas = ref([
	{id: '1', nama: 'Sangat Baik'},
	{id: '2', nama: 'Baik'},
	{id: '3', nama: 'Butuh Perbaikan'},
	{id: '4', nama: 'Kurang'},
	{id: '5', nama: 'Sangat Kurang'},
])
const pangkats = ref([])

const refKuadranKinerja = ref(null)
const refPangkat = ref(null)

const method = ref('Simpan')

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
	
  	const [pangkat] = await Promise.all([
    	$fetch('/api/gets/pangkat'),
  	])

  	pangkats.value = pangkat
  	
  	
  	if(refPangkat){
		refPangkat.value.reinit()
	}
});

const { pending, data, refresh} = await useAsyncData('getdataSkp', async() => {
	console.log("CariData")
	if(snip){
		let nip = $decodeBase64(snip)
		let thnilai = $decodeBase64(sthnilai)
		let tmulai = $decodeBase64(stmulai)

		var result = await $fetch(`/api/gets/riwayat_p2kp/${nip}/${thnilai}/${tmulai}`);

		if(result.nip == nip){
			method.value = "Update"

			dataSkp.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataSkp.value), key);
			})

			dataRefSkp.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataRefSkp.value), key);
			})
		}
	}
})

const tmulai = computed({
	get(){
		return dayjs(dataSkp.value.tmulai).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataSkp.value.tmulai = val
	}
})

const tselesai = computed({
	get(){
		return dayjs(dataSkp.value.tselesai).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataSkp.value.tselesai = val
	}
})

const hasil_kinerja = computed({
	get(){
		return dataSkp.value.hasil_kinerja
	},
	set(val){
		dataSkp.value.hasil_kinerja = val
		CalKuadranKinerja()
	}
})

const perilaku_kerja = computed({
	get(){
		return dataSkp.value.perilaku_kerja
	},
	set(val){
		dataSkp.value.perilaku_kerja = val
		CalKuadranKinerja()
	}
})

function CalKuadranKinerja(){
	if(dataSkp.value.hasil_kinerja != "" && dataSkp.value.perilaku_kerja != ""){
		dataSkp.value.kuadran_kinerja = resKuadran.value[dataSkp.value.hasil_kinerja][dataSkp.value.perilaku_kerja]

		if(refKuadranKinerja){
			refKuadranKinerja.value.reinit()
		}
	}
}

var resKuadran = ref([
    [0,1,2,3],
    [1,1,2,4],
    [2,2,2,4],
    [3,3,3,5]
]);

const jnsPejabatPenilai = computed({
	get(){
		return dataSkp.value.pej_pns == 0 ? false : true
	},
	set(val){
		dataSkp.value.pej_pns = val
	}
})

const ajnsPejabatPenilai = computed({
	get(){
		return dataSkp.value.pej_pns == 0 ? 'Non' : ''
	}
})

onMounted(() => {
	//refreshNuxtData(["getDataRef", "getdataSkpAkhir"])
	refreshNuxtData(["getDataRef", "getdataSkp"])
	modalUploadDoc.$onAction(callback, false)
})

const showSpinner = ref(false)
const toast = useToast()

async function simpanData() {
	showSpinner.value = true

	var compacted = _pickBy(dataSkp.value);
	compacted.tmulai = compacted.tmulai ? dayjs(dayjs(compacted.tmulai, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tselesai = compacted.tselesai ? dayjs(dayjs(compacted.tselesai, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.thnilai = compacted.thnilai *1
	compacted.hasil_kinerja = compacted.hasil_kinerja *1
	compacted.perilaku_kerja = compacted.perilaku_kerja *1

	var dr = dataRefSkp.value

	if(method.value == 'Update'){
		var body = {
			"refdata" : dr,
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_p2kp', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_p2kp',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_p2kp',
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		//compacted.akhir = 1
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_p2kp/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_p2kp',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_p2kp',
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
		}
	}
  	showSpinner.value = false
  	refreshNuxtData(["getDataRef", "getdataSkp"])
}

const doUploadDoc = async() => {
	let nip = $decodeBase64(snip)
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataSkp.value.filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen SKP`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen SKP`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		formData.append("thnilai", dataSkp.value.thnilai);
		formData.append("tmulai", tmulai.value);
		formData.append("tselesai", tselesai.value);
		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_p2kp_"+dataSkp.value.thnilai+"_"+tmulai.value);
		formData.append("path", 'dokumen/'+nip);

		var {data, error} = await useFetch(`/api/uploads/upload/p2kp/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(error.value){
			toast.add({
		    	id: 'error_post_upload_p2kp',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_p2kp',
		    	description: `Dokumen SKP berhasil di Upload`,
		    	timeout: 6000
		  	}) 	
		}

		refreshNuxtData(["getDataRef", "getdataSkp"])

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

const cariDataPNS = async() => {
	let nip = dataSkp.value.pej_nik_nip

	let anip = $decodeBase64(snip)

	if(anip == nip){
		toast.add({
	    	id: 'error_search_penilai',
	    	description: "Tidak Boleh Menilai Diri Sendiri",
	    	timeout: 6000
	  	}) 
	}else{
		if(dataSkp.value.pej_pns == 1){
			var _result = await $fetch(`/api/posts/pegawai/`, {
				method: "POST",
				body: JSON.stringify({
					nip: nip,
				})
			});

			var result = _result[0]

			if(result.nip == nip){
				dataSkp.value.pej_nama = result.namapeg
				console.log(result)
				dataSkp.value.pej_kgolru = result.pangkat_akhir?.kgolru
				if(refPangkat){
					refPangkat.value.reinit()
				}
				dataSkp.value.pej_njab = result?.jabatan_akhir?.njab
				dataSkp.value.pej_nunker = result?.jabatan_akhir?.nunker
			}
		}
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
						<h2 class="text-xl font-bold text-blue-600">SKP</h2>
						<div class="flex gap-x-2 w-[60%]">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP" :value="dataSkp.nip">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="Nama">
						</div>
					</div>
				</div>
				
				<form>
					<!-- Grid -->
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Tahun Penilaian</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.thnilai" @keypress="onlyNumber">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Priode Penilaian</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="date" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tmulai">
								<label class="inline-block text-sm text-gray-800 mt-2.5">s/d</label>
								<input type="date" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tselesai">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Hasil Kinerja</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<SearchSelect2 ref="refHasilKinerja" id="hasilKinerja" :options="ratings" keyList="id" namaList="nama" v-model="hasil_kinerja"/>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Perilaku Kerja</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<SearchSelect2 ref="refPerilakuKerja" id="perilakuKerja" :options="ratings" keyList="id" namaList="nama" v-model="perilaku_kerja"/>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Kuadran Kinerja</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refKuadranKinerja" id="kuadrankinerjas" :options="kuadrankinerjas" keyList="id" namaList="nama" v-model="dataSkp.kuadran_kinerja" disabled="true"/>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jenis Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<div class="flex mt-2">
								  	<input type="checkbox" class="shrink-0 mt-0.5 border-gray-200 rounded text-blue-600 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none" id="hs-default-checkbox" v-model="jnsPejabatPenilai">
								  	<label for="hs-default-checkbox" class="text-sm text-gray-500 ms-3">{{ajnsPejabatPenilai}} ASN</label>
								</div>
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">{{jnsPejabatPenilai ==1 ? "NIP" : "NIK"}} Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.pej_nik_nip">
						</div>
						<div class="sm:col-span-3">
							<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="cariDataPNS">
				  				Cari NIP
							</button>	
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nama Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.pej_nama">
						</div>
						<!-- End Col -->
						<template v-if="jnsPejabatPenilai ==1">
							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Golongan Pejabat Penilai</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<SearchSelect2 ref="refPangkat" id="pangkats" :options="pangkats" keyList="id" namaList="pangkat_gol" v-model="dataSkp.pej_kgolru"/>
							</div>
							<!-- End Col -->
						</template>
						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Unit Kerja Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.pej_nunker">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jabatan Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.pej_njab">
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