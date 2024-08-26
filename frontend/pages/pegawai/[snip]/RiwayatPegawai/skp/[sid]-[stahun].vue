<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)

const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const {snip, stahun, sid} = route.params

const ntahun = ref('')

const dataRefSkp = ref({
	id:'',
	nip: '',
	tahun:'',
})

const dataSkp = ref({
	id:'',
	nip: '',
	tahun: '',
	jenisJabatan: '',
	kerjasama: '',
	integritas: '',
	komitmen: '',
	disiplin: '',
	orientasiPelayanan: '',
	kepemimpinan : '',
	nilaiSkp: '',
	jumlah: '',
	nilairatarata: '',
	nilaiPerilakuKerja: '',
	nilaiPrestasiKerja: '',
	jenisPeraturanKinerjaKd: '',
	inisiatifKerja: '',
	konversiNilai: '',
	nilaiIntegrasi: '',
	statusPenilai: '',
	penilaiNipNrp: '',
	penilaiNama: '',
	penilaiGolongan: '',
	penilaiTmtGolongan: '',
	penilaiUnorNama:'',
	penilaiJabatan: '',
	statusAtasanPenilai:'',
	atasanPenilaiNipNrp:'',
	atasanPenilaiNama: '',
	atasanPenilaiGolongan: '',
	atasanPenilaiTmtGolongan: '',
	atasanPenilaiJabatan: '',
	atasanPenilaiUnorNama: '',
	idSync: '',
	filename: '',
})

const tahuns = ref([
	{id: '2017', nama: '2017'},
	{id: '2018', nama: '2018'},
	{id: '2019', nama: '2019'},
	{id: '2020', nama: '2020'},
	{id: '2021', nama: '2021'},
])

const jenisJabatans = ref([
	{id: 1, nama: 'Struktural'},
	{id: 2, nama: 'Non Struktural'},
])

const peraturans = ref([
	{id: 'PP30', nama: 'PP30'},
	{id: 'PP46', nama: 'PP46'},
])

const pangkats = ref([])

const refTahun = ref(null)
const refPeraturan = ref(null)
const refKuadranKinerja = ref(null)
const refPangkat = ref(null)
const refPangkat2 = ref(null)

const method = ref('Simpan')
const dtSkp = ref([])

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
	
  	const [pangkat] = await Promise.all([
    	$fetch('/api/gets/pangkat'),
  	])

  	pangkats.value = pangkat
  	
  	
  	if(refPangkat){
		refPangkat.value.reinit()
	}

	if(refPangkat){
		refPangkat2.value.reinit()
	}
});

const { pending, data, refresh} = await useAsyncData('getdataSkpp', async() => {
	console.log("CariData Skp")
	if(snip){
		let nip = $decodeBase64(snip)
		let id = $decodeBase64(sid)
		let tahun = $decodeBase64(stahun)

		//ntahun.value = tahun

		var result = await $fetch(`/api/gets/riwayat_skp/${nip}/${id}/${tahun}`);
		dtSkp.value = result

		if(result.nip == nip){
			console.log(result)
			method.value = "Update"

			dataSkp.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataSkp.value), key);
			})

			dataRefSkp.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataRefSkp.value), key);
			})

			if(refTahun){
				refTahun.value.reinit()
			}
		}
	}
})

const nilaiPerilakuKerja = computed({
	get(){
		//var result = dataSkp.value.nilaiSkp * 1
		/*var result = dataSkp.value.kerjasama + dataSkp.value.integritas + dataSkp.value.komitmen + dataSkp.value.disiplin + dataSkp.value.orientasiPelayanan
		if(dataSkp.value.jenisJabatan == 1){
			result = (result + dataSkp.value.kepemimpinan) / 6
		}else{
			result = result / 5
		}
		return result.toFixed(3)*/
		return dataSkp.value.nilaiPerilakuKerja
	},
	set(){
		dataSkp.value.nilaiPerilakuKerja = val
	}
})

const nilaiSkp60 = computed({
	get(){
		var result = dataSkp.value.nilaiSkp * 1
		return (result * 0.6).toFixed(3)
	}
})

const nilaiPerilakuKerja40 = computed({
	get(){
		var result = dataSkp.value.nilairatarata * 1
		return (result * 0.4).toFixed(3)
	}
})

const jnsPejabatPenilai = computed({
	get(){
		return dataSkp.value.statusPenilai == 'ASN' ? true : false
	},
	set(val){
		dataSkp.value.statusPenilai = val == true ? 'ASN' : 'NON ASN'
	}
})

const jnsAtasanPejabatPenilai = computed({
	get(){
		return dataSkp.value.statusAtasanPenilai == 'ASN' ? true : false
	},
	set(val){
		dataSkp.value.statusAtasanPenilai = val == true ? 'ASN' : 'NON ASN'
	}
})

onMounted(() => {
	//refreshNuxtData(["getDataRef", "getdataSkpAkhir"])
	refreshNuxtData(["getDataRef", "getdataSkpp"])
	modalUploadDoc.$onAction(callback, false)
})

const showSpinner = ref(false)
const toast = useToast()

async function simpanData() {
	showSpinner.value = true
	let canUpdateBkn = false

	var compacted = _pickBy(dataSkp.value);
	/*compacted.tmulai = compacted.tmulai ? dayjs(dayjs(compacted.tmulai, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tselesai = compacted.tselesai ? dayjs(dayjs(compacted.tselesai, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.thnilai = compacted.thnilai *1
	compacted.hasil_kinerja = compacted.hasil_kinerja *1
	compacted.perilaku_kerja = compacted.perilaku_kerja *1*/

	var dr = dataRefSkp.value

	if(method.value == 'Update'){
		var body = {
			"refdata" : dr,
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_skp', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			const {title, description} = error.value.data.data
			toast.add({
		    	id: 'error_put_skp',
		    	icon: "i-heroicons-x-circle",
		    	title: title,
		    	description: description,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_skp',
		    	icon: "i-heroicons-check-circle",
		    	title: "Update Data SKP",
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		  	canUpdateBkn = false
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		//compacted.akhir = 1
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_skp/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_skp',
		    	icon: "i-heroicons-x-circle",
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_skp',
		    	icon: "i-heroicons-check-circle",
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
		  	compacted.id = data.value.id
		  	canUpdateBkn = false
		}
	}

	//update BKN
	//compacted.nip = $decodeBase64(snip)
	//compacted.akhir = 1
	if(canUpdateBkn == true){
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_skp/bkn', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			const {title, description} = error.value.data.data
			toast.add({
		    	id: 'error_post_skp',
		    	icon: "i-heroicons-x-circle",
		    	title: title,
		    	description: description,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_kinerja',
		    	icon: "i-heroicons-check-circle",
		    	title: "Update Data SKP BKN",
		    	description: "Data SKP BKN berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
	}
	

	//end Update BKN

  	showSpinner.value = false
  	refreshNuxtData(["getDataRef", "getdataSkp"])
}

const doUploadDoc = async() => {
	let nip = $decodeBase64(snip)
	//console.log(dataSkp.value.filename)
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

		formData.append("tahun", dataSkp.value.tahun);
		formData.append("id", dataSkp.value.id);
		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_skp_"+dataSkp.value.tahun+"_"+dataSkp.value.id);
		formData.append("path", 'dokumen/'+nip);

		formData.append("id_ref_dokumen", 891)
		formData.append("idSync", dataSkp.value.idSync)

		var {data, error} = await useFetch(`/api/uploads/upload/skp/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(data.value.status != "200"){
			const {title, description} = data.value.data
			toast.add({
		    	id: 'error_post_upload_skp',
		    	icon: "i-heroicons-x-circle",
				title: title,
				description: description,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_skp',
		    	icon: "i-heroicons-check-circle",
	    		title: 'Upload Dokumen SKP',
		    	description: `Dokumen SKP berhasil di Upload`,
		    	timeout: 6000,
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
	let nip = dataSkp.value.penilaiNipNrp
	console.log(nip)

	let anip = $decodeBase64(snip)

	if(anip == nip){
		toast.add({
	    	id: 'error_search_penilai',
	    	icon: "i-heroicons-x-circle",
	    	title: 'Cari Pejabat Penilai',
	    	description: "Tidak Boleh Menilai Diri Sendiri",
	    	timeout: 6000,
	    	color: 'red',
	  	}) 
	}else{
		if(dataSkp.value.statusPenilai == 'ASN'){
			var _result = await $fetch(`/api/posts/pegawai/`, {
				method: "POST",
				body: JSON.stringify({
					nip: nip,
				})
			});

			var result = _result[0]

			if(result.nip == nip){
				dataSkp.value.penilaiNama = result.namapeg
				dataSkp.value.penilaiGolongan = result.pangkat_akhir?.pangkat?.nama
				if(refPangkat){
					refPangkat.value.reinit()
				}
				dataSkp.value.penilaiJabatan = result?.jabatan_akhir?.njab
				dataSkp.value.penilaiUnorNama = result?.jabatan_akhir?.nunker
			}
		}
	}
}

const cariDataPNS2 = async() => {
	let nip = dataSkp.value.atasanPenilaiNipNrp
	console.log(nip)

	let anip = $decodeBase64(snip)

	if(anip == nip){
		toast.add({
	    	id: 'error_search_penilai',
	    	icon: "i-heroicons-x-circle",
	    	title: 'Cari Atasam Pejabat Penilai',
	    	description: "Tidak Boleh Menilai Diri Sendiri",
	    	timeout: 6000,
	    	color: 'red',
	  	}) 
	}else{
		if(dataSkp.value.statusPenilai == 'ASN'){
			var _result = await $fetch(`/api/posts/pegawai/`, {
				method: "POST",
				body: JSON.stringify({
					nip: nip,
				})
			});

			var result = _result[0]

			if(result.nip == nip){
				console.log(result)
				dataSkp.value.atasanPenilaiNama = result.namapeg
				dataSkp.value.atasanPenilaiGolongan = result.pangkat_akhir?.pangkat?.nama
				if(refPangkat2){
					refPangkat2.value.reinit()
				}
				dataSkp.value.atasanPenilaiJabatan = result?.jabatan_akhir?.njab
				dataSkp.value.atasanPenilaiUnorNama = result?.jabatan_akhir?.nunker
			}
		}
	}
}

function checkMax(){
	var thn = dataSkp.value.tahun
	console.log(thn)
	if(thn > 2021){
		dataSkp.value.tahun = 2021
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
						<h2 class="text-xl font-bold text-blue-600">SKP {{ntahun}}</h2>
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
							<input type="number" step="1" max="2021" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" v-model="dataSkp.tahun" @keyup="checkMax">
							<!--<SearchSelect2 ref="refTahun" id="tahuns" :options="tahuns" keyList="id" namaList="nama" v-model="dataSkp.tahun"/>-->
							<!--<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.thnilai" @keypress="onlyNumber">
							</div>-->
						</div>
						<!-- End Col -->

						<template v-if="dataSkp.tahun == '2021'">
							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Peraturan</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<!--<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.jenisPeraturanKinerjaKd">-->
								<SearchSelect2 ref="refPeraturan" id="peraturans" :options="peraturans" keyList="id" namaList="nama" v-model="dataSkp.jenisPeraturanKinerjaKd"/>
							</div>	
						</template>

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jenis Jabatan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refJenisJabatan" id="jenisJabatans" :options="jenisJabatans" keyList="id" namaList="nama" v-model="dataSkp.jenisJabatan"/>
						</div>
						<!-- End Col -->
						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nilai SKP</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.nilaiSkp" @keypress="onlyNumber">
						</div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nilai SKP (60%)</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nilaiSkp60" @keypress="onlyNumber">
						</div>

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Orientasi Pelayanan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.orientasiPelayanan" @keypress="onlyNumber">
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Integritas</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.integritas" @keypress="onlyNumber">
						</div>
						<!-- End Col -->
						<template v-if="dataSkp.tahun == '2021'">
							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Inisiatif Kerja</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.inisiatifKerja" @keypress="onlyNumber">
							</div>
							<!-- End Col -->
						</template>

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Komitmen</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.komitmen" @keypress="onlyNumber">
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Disiplin</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.disiplin" @keypress="onlyNumber">
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Kerjasama</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.kerjasama" @keypress="onlyNumber">
						</div>
						<!-- End Col -->
						
						<template v-if="dataSkp.jenisJabatan == 1">
							<div class="sm:col-span-3 sm:col-start-1">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Kepemimpinan</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.kepemimpinan" @keypress="onlyNumber">
							</div>
							<!-- End Col -->
						</template>

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nilai Perilaku Kerja Rata-rata</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.nilairatarata" @keypress="onlyNumber">
						</div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nilai Perilaku Kerja (40%)</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nilaiPerilakuKerja40" @keypress="onlyNumber">
						</div>

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nilai Prestasi Kerja</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.nilaiPrestasiKerja" @keypress="onlyNumber">
						</div>

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>
						<template v-if="dataSkp.tahun == '2021'">
							<div class="sm:col-span-3 sm:col-start-1">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Konversi Nilai Prestasi Kerja</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.konversiNilai" @keypress="onlyNumber">
							</div>

							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Nilai Integrasi Kinerja</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.nilaiIntegrasi" @keypress="onlyNumber">
							</div>

							<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>
						</template>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jenis Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<div class="flex mt-2">
								  	<input type="checkbox" class="shrink-0 mt-0.5 border-gray-200 rounded text-blue-600 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none" id="hs-default-checkbox1" v-model="jnsPejabatPenilai">
								  	<label for="hs-default-checkbox1" class="text-sm text-gray-500 ms-3">{{dataSkp.statusPenilai}}</label>
								</div>
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">{{jnsPejabatPenilai ==1 ? "NIP" : "NIK"}} Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.penilaiNipNrp">
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
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.penilaiNama">
						</div>
						<!-- End Col -->
						<template v-if="jnsPejabatPenilai ==1">
							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Golongan Pejabat Penilai</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<SearchSelect2 ref="refPangkat" id="pangkats" :options="pangkats" keyList="nama" namaList="pangkat_gol" v-model="dataSkp.penilaiGolongan"/>
							</div>
							<!-- End Col -->
						</template>
						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Unit Kerja Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.penilaiUnorNama">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jabatan Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.penilaiJabatan">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jenis Atasan Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<div class="flex mt-2">
								  	<input type="checkbox" class="shrink-0 mt-0.5 border-gray-200 rounded text-blue-600 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none" id="hs-default-checkbox" v-model="jnsAtasanPejabatPenilai">
								  	<label for="hs-default-checkbox" class="text-sm text-gray-500 ms-3">{{dataSkp.statusAtasanPenilai}}</label>
								</div>
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">{{jnsAtasanPejabatPenilai ==1 ? "NIP" : "NIK"}} Atasan Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.atasanPenilaiNipNrp">
						</div>
						<div class="sm:col-span-3">
							<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="cariDataPNS2">
				  				Cari NIP
							</button>	
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nama Atasan Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.atasanPenilaiNama">
						</div>
						<!-- End Col -->
						<template v-if="jnsAtasanPejabatPenilai ==1">
							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Golongan Atasan Pejabat Penilai</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<SearchSelect2 ref="refPangkat2" id="pangkats2" :options="pangkats" keyList="nama" namaList="pangkat_gol" v-model="dataSkp.atasanPenilaiGolongan"/>
							</div>
							<!-- End Col -->
						</template>
						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Unit Kerja Atasan Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.atasanPenilaiUnorNama">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jabatan Atasan Pejabat Penilai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataSkp.atasanPenilaiJabatan">
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

						<!--<div class="sm:col-span-3 flex justify-end" v-if="method == 'Update'">
							<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-green-600 text-white hover:bg-green-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="doUploadDoc()">
				  				Upload File
							</button>
						</div>-->
						<!-- End Col -->
		  			</div>
				</form>
	  		</div>
	  		<!-- End Card -->
		</div>
		<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2 mt-6" v-if="method == 'Update'">
			<div class="mb-8">
				<h2 class="text-xl font-bold text-blue-600">Dokumen SKP</h2>
			</div>
			<div class="grid sm:grid-cols-4 gap-2 gap-2.5">
				<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc()">
					Dokumen Realisasi SKP
				</button>
			</div>
		</div>
		<!-- End Card Section -->
	</LayoutRiwayatPegawai>
</template>