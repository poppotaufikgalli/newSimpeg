<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)

const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const {snip, sthnilai, sid, jns} = route.params

const dataRefKinerja = ref({
	id:'',
	nip: '',
	thnilai:'',
})

const dataKinerja = ref({
	id: '',
	nip: '',
	thnilai:'',
	tmulai: '',
	tselesai: '',
	jns: jns,
	hasil_kinerja: '',
	perilaku_kerja: '',
	kuadran_kinerja: '',
	periodik_id : '',
	jenjang:'',
	koefisien_id: '',
	nilai_koefisien: '',
	pej_pns: 'ASN',
	pej_nik_nip: '',
	pej_nama: '',
	pej_kgolru:'',
	pej_nunker:'',
	pej_njab: '',
	idSync: '',
	filename: '',
})

const ratings = ref([
	{id: '1', nama: 'Diatas Ekspektasi'},
	{id: '2', nama: 'Sesuai Ekspektasi'},
	{id: '3', nama: 'Dibawah Ekspektasi'},
])
const kuadrankinerjas = ref([
	{id: '1', nama: 'Sangat Baik', persentase: 150},
	{id: '2', nama: 'Baik', persentase: 100},
	{id: '3', nama: 'Butuh Perbaikan', persentase: 75},
	{id: '4', nama: 'Kurang', persentase: 50},
	{id: '5', nama: 'Sangat Kurang', persentase: 25},
])

const periodik = ref([
	{id: 1, nama: 'Bulanan'},
	{id: 2, nama: 'Triwulan'},
])

const pangkats = ref([])
const jenjang_jabfungs = ref([])

const refKuadranKinerja = ref(null)
const refPangkat = ref(null)
const refPeriodik = ref(null)
const refJenjangJabfung = ref(null)

const method = ref('Simpan')

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
	
  	const [pangkat, jenjang_jabfung] = await Promise.all([
    	$fetch('/api/gets/pangkat'),
    	$fetch('/api/gets/jenjang_jabfung'),
  	])

  	pangkats.value = pangkat
  	jenjang_jabfungs.value = jenjang_jabfung
  	console.log(jenjang_jabfungs.value)
  	
  	if(refPangkat){
		refPangkat.value.reinit()
	}

	if(refJenjangJabfung){
		refJenjangJabfung.value.reinit()
	}
});

const { pending, data, refresh} = await useAsyncData('getdataKinerja', async() => {
	console.log("CariData")
	if(snip){
		let nip = $decodeBase64(snip)
		let thnilai = $decodeBase64(sthnilai)
		let id = $decodeBase64(sid)

		var result = await $fetch(`/api/gets/riwayat_kinerja/${nip}/${id}/${thnilai}`);

		if(result.nip == nip){
			method.value = "Update"

			dataKinerja.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataKinerja.value), key);
			})

			dataRefKinerja.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataRefKinerja.value), key);
			})

			CalKuadranKinerja()
		}
	}
})

const tmulai = computed({
	get(){
		return dayjs(dataKinerja.value.tmulai).format("YYYY-MM").toString()
	},
	set(val){
		dataKinerja.value.tmulai = val
	}
})

const tselesai = computed({
	get(){
		return dayjs(dataKinerja.value.tselesai).format("YYYY-MM").toString()
	},
	set(val){
		dataKinerja.value.tselesai = val
	}
})

const hasil_kinerja = computed({
	get(){
		return dataKinerja.value.hasil_kinerja
	},
	set(val){
		dataKinerja.value.hasil_kinerja = val
		CalKuadranKinerja()
	}
})

const perilaku_kerja = computed({
	get(){
		return dataKinerja.value.perilaku_kerja
	},
	set(val){
		dataKinerja.value.perilaku_kerja = val
		CalKuadranKinerja()
	}
})

var resKuadran = ref([
    [0,1,2,3],
    [1,1,2,4],
    [2,2,2,4],
    [3,3,3,5]
]);

const persentase = ref(0)

function CalKuadranKinerja(){
	if(dataKinerja.value.hasil_kinerja != "" && dataKinerja.value.perilaku_kerja != ""){
		dataKinerja.value.kuadran_kinerja = resKuadran.value[dataKinerja.value.hasil_kinerja][dataKinerja.value.perilaku_kerja]
		//console.log(resKuadran.value[dataKinerja.value.hasil_kinerja])

		if(refKuadranKinerja){
			refKuadranKinerja.value.reinit()
		}

		var a = _find(kuadrankinerjas.value, function(o){
			return o.id == dataKinerja.value.kuadran_kinerja
		})

		persentase.value = a.persentase
	}
}

const jenjang = computed({
	get(){
		return dataKinerja.value.jenjang
	},
	set(val){
		dataKinerja.value.jenjang = val

		var e = _find(jenjang_jabfungs.value, function(o){
			return o.id == val
		})
		dataKinerja.value.koefisien_id = e.koefisien_id
		dataKinerja.value.nilai_koefisien = e.nilai_koefisien
	}
})

const jnsPejabatPenilai = computed({
	get(){
		return dataKinerja.value.pej_pns == 'ASN' ? true : false
	},
	set(val){
		dataKinerja.value.pej_pns = val == true ? 'ASN' : 'NON ASN'
	}
})

onMounted(() => {
	//refreshNuxtData(["getDataRef", "getdataKinerjaAkhir"])
	refreshNuxtData(["getDataRef", "getdataKinerja"])
	modalUploadDoc.$onAction(callback, false)
})

const showSpinner = ref(false)
const toast = useToast()

async function simpanData() {
	showSpinner.value = true
	let canUpdateBkn = false

	var compacted = _pickBy(dataKinerja.value);
	compacted.tmulai = compacted.tmulai ? dayjs(dayjs(compacted.tmulai, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tselesai = compacted.tselesai ? dayjs(dayjs(compacted.tselesai, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.thnilai = compacted.thnilai *1
	compacted.hasil_kinerja = compacted.hasil_kinerja *1
	compacted.perilaku_kerja = compacted.perilaku_kerja *1
	compacted.nilai_koefisien = compacted.nilai_koefisien *1
	compacted.periodik_id = compacted.periodik_id *1
	compacted.koefisien_id = (compacted.koefisien_id).toString()

	var dr = dataRefKinerja.value

	if(method.value == 'Update'){
		var body = {
			"refdata" : dr,
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_kinerja', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			const {title, message, color} = error.value.data.data
			toast.add({
		    	id: 'error_put_kinerja',
		    	icon: "i-heroicons-x-circle",
		    	title: title ?? "Update Riwayat Kinerja Simpeg",
		    	description: message ?? error.value.data.data.errors,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_kinerja',
		    	icon: "i-heroicons-check-circle",
	    		title: "Data Riwayat Kinerja",
		    	description: "Data berhasil di Update",
		    	timeout: 6000,
		  	}) 	
		  	canUpdateBkn = true
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		//compacted.akhir = 1
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_kinerja/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			const {title, message, color} = error.value.data.data
			toast.add({
		    	id: 'error_post_kinerja',
		    	icon: "i-heroicons-x-circle",
		    	title: title ?? "Tambah Riwayat Kinerja Simpeg",
		    	description: message,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_kinerja',
		    	icon: "i-heroicons-check-circle",
	    		title: "Data Riwayat Kinerja",
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
		  	compacted.id = data.value.id
		  	canUpdateBkn = true
		}
	}

	//update bkn
	if(canUpdateBkn == true){
		var body = JSON.stringify(compacted)
		var {data, error} = await useFetch('/api/posts/riwayat_kinerja/bkn', {
			method: 'POST',
			body: JSON.stringify(body),
		})

		if(error.value){
			const {title, message, color} = error.value.data.data
			toast.add({
		    	id: 'error_put_kinerja_bkn',
		    	icon: "i-heroicons-x-circle",
		    	title: title ?? 'Data Riwayat Kinerja BKN',
		    	description: message,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_kinerja_bkn',
		    	icon: "i-heroicons-check-circle",
	    		title: "Data Riwayat Kinerja BKN",
		    	description: "Data berhasil di Update",
		    	timeout: 6000,
		  	}) 	
		  	canUpdateBkn = true
		}
	}
	
	//end update bkn

  	showSpinner.value = false
  	refreshNuxtData(["getDataRef", "getdataKinerja"])
}

const doUploadDoc = async() => {
	let nip = $decodeBase64(snip)
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataKinerja.value.filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen Kinerja`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen Kinerja`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		formData.append("thnilai", dataKinerja.value.thnilai);
		formData.append("id", dataKinerja.value.id);
		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_kinerja_"+dataKinerja.value.thnilai+"_"+dataKinerja.value.id);
		formData.append("path", 'dokumen/'+nip);

		formData.append("id_ref_dokumen", 891)
		formData.append("idSync", dataKinerja.value.idSync)

		var {data, error} = await useFetch(`/api/uploads/upload/kinerja/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(data.value.status != "200"){
			const {title, description} = data.value.data
			toast.add({
		    	id: 'error_post_upload_kinerja',
		    	icon: "i-heroicons-x-circle",
				title: title,
				description: description,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_kinerja',
		    	icon: "i-heroicons-check-circle",
	    		title: 'Upload Dokumen Kinerja',
		    	description: `Dokumen Kinerja berhasil di Upload`,
		    	timeout: 6000,
		  	}) 	
		}

		refreshNuxtData(["getDataRef", "getdataKinerja"])

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
	let nip = dataKinerja.value.pej_nik_nip

	let anip = $decodeBase64(snip)

	if(anip == nip){
		toast.add({
	    	id: 'error_search_penilai',
	    	icon: "i-heroicons-x-circle",
	    	title: 'Cari Pejabat Penilai',
	    	description: "Tidak Boleh Menilai Diri Sendiri",
	    	timeout: 6000
	  	}) 
	}else{
		if(dataKinerja.value.pej_pns == 'ASN'){
			var _result = await $fetch(`/api/posts/pegawai/`, {
				method: "POST",
				body: JSON.stringify({
					nip: nip,
				})
			});

			var result = _result[0]

			if(result.nip == nip){
				dataKinerja.value.pej_nama = result.namapeg
				console.log(result)
				dataKinerja.value.pej_kgolru = result.pangkat_akhir?.kgolru
				if(refPangkat){
					refPangkat.value.reinit()
				}
				dataKinerja.value.pej_njab = result?.jabatan_akhir?.njab
				dataKinerja.value.pej_nunker = result?.jabatan_akhir?.nunker
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
						<h2 class="text-xl font-bold text-blue-600">Kinerja {{jns == 'P' ? 'Periodik' : ''}}</h2>
						<div class="flex gap-x-2 w-[60%]">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP" :value="dataKinerja.nip">
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
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKinerja.thnilai" @keypress="onlyNumber">
							</div>
						</div>
						<!-- End Col -->

						<template v-if="jns== 'P'">
							<div class="sm:col-span-3 sm:col-start-1">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Jenis Periodik</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<SearchSelect2 ref="refPeriodik" id="periodik" :options="periodik" keyList="id" namaList="nama" v-model="dataKinerja.periodik_id"/>
							</div>
							<!-- End Col -->

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
						</template>

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
							<SearchSelect2 ref="refKuadranKinerja" id="kuadrankinerjas" :options="kuadrankinerjas" keyList="id" namaList="nama" v-model="dataKinerja.kuadran_kinerja" disabled="true"/>
						</div>
						<!-- End Col -->
						<template v-if="jns== 'P'">
							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Persentase Kinerja</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<div class="sm:flex gap-2">
									<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="persentase" disabled>
								</div>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Jenjang Jabatan</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-9">
								<SearchSelect2 ref="refJenjangJabfung" id="jenjang_jabfungs" :options="jenjang_jabfungs" keyList="id" namaList="nama" v-model="jenjang"/>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Koefisien {{dataKinerja.priodik_id}}</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<div class="sm:flex gap-2">
									<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKinerja.nilai_koefisien" disabled>
								</div>
							</div>
							<!-- End Col -->
						</template>

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jenis Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<div class="flex mt-2">
								  	<input type="checkbox" class="shrink-0 mt-0.5 border-gray-200 rounded text-blue-600 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none" id="hs-default-checkbox" v-model="jnsPejabatPenilai">
								  	<label for="hs-default-checkbox" class="text-sm text-gray-500 ms-3">{{dataKinerja.pej_pns}}</label>
								</div>
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">{{dataKinerja.pej_pns == 'ASN' ? "NIP" : "NIK"}} Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKinerja.pej_nik_nip">
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
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKinerja.pej_nama">
						</div>
						<!-- End Col -->
						<template v-if="dataKinerja.pej_pns == 'ASN'">
							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Golongan Pejabat Penilai</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<SearchSelect2 ref="refPangkat" id="pangkats" :options="pangkats" keyList="id" namaList="pangkat_gol" v-model="dataKinerja.pej_kgolru"/>
							</div>
							<!-- End Col -->
						</template>
						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Unit Kerja Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKinerja.pej_nunker">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jabatan Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataKinerja.pej_njab">
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
				<h2 class="text-xl font-bold text-blue-600">Dokumen Kinerja</h2>
			</div>
			<div class="grid sm:grid-cols-4 gap-2 gap-2.5">
				<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc()">
					Dokumen Realisasi Kinerja
				</button>
			</div>
		</div>
		<!-- End Card Section -->
	</LayoutRiwayatPegawai>
</template>