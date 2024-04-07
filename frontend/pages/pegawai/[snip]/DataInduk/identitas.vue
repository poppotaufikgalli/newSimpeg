<script setup>
const route = useRoute()
const { $decodeBase64 } = useNuxtApp()
const dayjs = useDayjs()
const {snip} = route.params

const method = ref('Simpan')

const dataInduk = ref({
	nip: '',
	nama: '',
	gldepan: '',
	glblk: '',
	ktlahir: '',
	tlahir: '',
	kjkel: '',
	kagama: '',
	kstatus: '',
	kjpeg: '',
	kduduk: '',
	kskawin: '',
	tinggi_badan:'',
	berat_badan:'',
	kgoldar: '',
	aljalan:'',
	alrt: '',
	alrw: '',
	altelp: '',
	altelpwa: '',
	alkoprop: '',
	alkokab: '',
	alkokec: '',
	alkodes: '',
	kpos:'',
	kaparpol: '',
	npap_g: '',
	nkarpeg: '',
	nakses: '',
	ntaspen:'',
	nkaris_su:'',
	npwp: '',
	nik:'',
	file_bmp: '',
	aktif: '',
	jjiwa: '',
	marga: '',
	suku: '',
	tgl_peg: '',
	niplama: '',
	tgl_reg: '',
	stat_kpe: '',
	no_pinpeg: '',
	alkoproplahir:'',
	alkokablahir: '',
	tgl_kpe: '',
	thn_pendataan: '',
	no_ref_bkn: '',
	email:'',
	email_pemerintah: '',
})

const { pending, data, refresh} = await useAsyncData('getDataIdentitas', async() => {
	console.log("CariData Identitas")

	if(snip){
		let nip = $decodeBase64(snip)
		dataInduk.value = await $fetch('/api/gets/pegawai/'+nip);

		if(dataInduk.value.nip == nip){
			method.value = "Update"
		}
	}
})

const agamas = ref([])
const statuss = ref([])
const jeniss = ref([])
const duduks = ref([])
const kawins = ref([])
const goldars = ref([])
const wilayahs = ref([])
const kpes = ref([])

const refKprov = ref(null)
const refKkab = ref(null)
const refKkec = ref(null)
const refKkel = ref(null)

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useAsyncData('getDataRef', async () => {
  	const [agama, status, jenis, duduk, kawin, goldar, kpe, wilayah] = await Promise.all([
    	$fetch('/api/gets/agama'),
    	$fetch('/api/gets/status_pegawai'),
    	$fetch('/api/gets/jenis_pegawai'),
    	$fetch('/api/gets/kedudukan_pegawai'),
    	$fetch('/api/gets/jenis_kawin'),
    	$fetch('/api/gets/jenis_goldar'),
    	$fetch('/api/gets/kpe'),
    	$fetch('/api/gets/wilayah'),
  	])

  	agamas.value = agama
  	statuss.value = status
  	jeniss.value = jenis
  	duduks.value = duduk
  	kawins.value = kawin
  	goldars.value = goldar
  	kpes.value = kpe
  	wilayahs.value = wilayah
});

onMounted(() => {
	refreshNuxtData(["getDataIdentitas", "getDataRef"])
})

const provs = computed({
	get(){
		var result = _filter(wilayahs.value, function(o){
			return o.twil == 1
		})

		return result
	}
})

const kabs = computed({
	get(){
		var result = _filter(wilayahs.value, function(o){
			return o.twil == 2 && o.kprov == dataInduk.value.alkoprop
		})

		//console.log(result)
		//if(result.length > 0){
			nextTick(() => {
				if(refKkab.value){
					refKkab.value.reinit()		
				}
			})
		//}
		
		return result
	}
})

const kecs = computed({
	get(){
		var result = _filter(wilayahs.value, function(o){
			return o.twil == 3 && o.kprov == dataInduk.value.alkoprop && o.kkab == dataInduk.value.alkokab
		})

		//if(result.length > 0){
			nextTick(() => {
				if(refKkec.value){
					refKkec.value.reinit()	
				}
			})
		//}

		return result
	}
})

const kels = computed({
	get(){
		var result = _filter(wilayahs.value, function(o){
			return o.twil == 4 && o.kprov == dataInduk.value.alkoprop && o.kkab == dataInduk.value.alkokab && o.kkec == dataInduk.value.alkokec
		})

		//if(result.length > 0){
			//refKkel.value.reinit()
			nextTick(() => {
				if(refKkel.value){
					refKkel.value.reinit()	
				}
			})
		//}

		return result
	}
})

const tlahir = computed({
	get(){
		return dayjs(dataInduk.value.tlahir).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataInduk.value.tlahir = val
	}
})

const tgl_kpe = computed({
	get(){
		return dayjs(dataInduk.value.tgl_kpe).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataInduk.value.tgl_kpe = val
	}
})

const showSpinner = ref(false)
const toast = useToast()

async function simpanData() {
	//console.log(dataInduk.value)

	showSpinner.value = true

	var body = JSON.stringify(dataInduk.value)

	var url = '/api/pegawai/new'

	if(method.value == 'Update'){
		//url = '/api/pegawai'
		var {data, error} = await useFetch('/api/pegawai', {
			method: 'PUT',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_identitas',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_identitas',
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
		//console.log(error.value.data.data)
		
	}else{
		var result = await $fetch('/api/pegawai/new', {
			method: 'POST',
			body: body,
		})

		console.log(result)
	}

  	showSpinner.value = false
}

</script>
<template>
	<AppLoadingSpinner :show="showSpinner" />
	<LayoutDataInduk>
		<div class="mx-auto">
			<!-- Card -->
			<NewCard>
				<div class="mb-8">
					<div class="inline-flex justify-between items-center w-full gap-x-2">
						<h2 class="text-xl font-bold text-blue-600">Identitas Pegawai</h2>
						<div class="flex w-[60%]">
							<input type="text" name="hs-input-with-add-on-url" id="hs-input-with-add-on-url" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-s-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" disabled placeholder="ID PNS" :value="dataInduk.no_ref_bkn">
							<button type="button" class="py-3 px-4 inline-flex justify-center items-center gap-x-2 text-xs font-semibold rounded-e-md border border-transparent bg-gray-600 text-white hover:bg-gray-700 disabled:opacity-50 disabled:pointer-events-none whitespace-nowrap">
								<svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
										<path fill-rule="evenodd" d="M15.312 11.424a5.5 5.5 0 0 1-9.201 2.466l-.312-.311h2.433a.75.75 0 0 0 0-1.5H3.989a.75.75 0 0 0-.75.75v4.242a.75.75 0 0 0 1.5 0v-2.43l.31.31a7 7 0 0 0 11.712-3.138.75.75 0 0 0-1.449-.39Zm1.23-3.723a.75.75 0 0 0 .219-.53V2.929a.75.75 0 0 0-1.5 0V5.36l-.31-.31A7 7 0 0 0 3.239 8.188a.75.75 0 1 0 1.448.389A5.5 5.5 0 0 1 13.89 6.11l.311.31h-2.432a.75.75 0 0 0 0 1.5h4.243a.75.75 0 0 0 .53-.219Z" clip-rule="evenodd" />
								</svg>
								ID PNS
							</button>
						</div>
					</div>
				</div>

				<form>
					<!-- Grid -->
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">NIP</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.nip">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">NIK/NPWP</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="NIK" v-model="dataInduk.nik">
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="NPWP" v-model="dataInduk.npwp">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Nama</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex">
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.nama">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Gelar</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-[35%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="gelar depan" v-model="dataInduk.gldepan">
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-[35%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="Gelar Belakang" v-model="dataInduk.glblk">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
			  				<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Tempat / Tgl Lahir</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
			  				<div class="sm:flex gap-2">
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.ktlahir">
								<input id="af-account-full-name" type="date" class="py-2 px-3 block w-[50%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tlahir">
			  				</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Jenis Kelamin</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<div class="sm:flex">
								<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.kjkel">
								  	<option selected>Pilih Jenis Kelamin</option>
								  	<option value="1">1 - Laki-laki</option>
								  	<option value="2">2 - Perempuan</option>
								</select>
						  	</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-2 sm:col-start-8">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Agama</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<div class="sm:flex">
						  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.kagama">
								  	<option selected>Pilih Agama</option>
								  	<template  v-for="(item, idx) in agamas" :key="idx">
									  	<option :value="item.id">{{item.id}} - {{item.nama}}</option>
									</template>
								</select>
						  	</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Status Perkawinan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<div class="sm:flex">
						  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.kskawin">
								  	<option selected>Pilih Status Perkawinan</option>
								  	<template  v-for="(item, idx) in kawins" :key="idx">
									  	<option :value="item.id">{{item.id}} - {{item.nama}}</option>
									  </template>
								</select>
						  	</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-2 sm:col-start-8">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Golongan Darah</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<div class="sm:flex">
						  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.kgoldar">
								  	<option selected>Pilih Golongan Darah</option>
								  	<template  v-for="(item, idx) in goldars" :key="idx">
									  	<option :value="item.id">{{item.id}} - {{item.nama}}</option>
									  </template>
								</select>
						  	</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Tinggi Badan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<div class="sm:flex">
						  		<input id="af-account-full-name" type="number" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.tinggi_badan">
						  	</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-2 sm:col-start-8">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Berat Badan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<div class="sm:flex">
						  		<input id="af-account-full-name" type="number" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.berat_badan">
						  	</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Alamat</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<textarea id="af-account-bio" class="py-2 px-3 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" rows="3" placeholder="Alamat..." v-model="dataInduk.aljalan"></textarea>
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">RT/RW</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<div class="sm:flex gap-2">
						  		<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="RT" v-model="dataInduk.alrt">
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="RW" v-model="dataInduk.alrw">
						  	</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-2 sm:col-start-8">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Kode Pos</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<div class="sm:flex">
						  		<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.kpos">
						  	</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Provinsi</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 v-if="!pending && !pendingRef" ref="refKprov" id="provs" :options="provs" keyList="kprov" namaList="nwil" v-model="dataInduk.alkoprop" />
						  	<!--<div class="sm:flex">
						  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.alkoprop">
								  	<option selected>Pilih Provinsi</option>
								  	<template v-for="(item, idx) in provs" :key="idx">
									  	<option :value="item.kprov">{{item.kprov}} - {{item.nwil}}</option>
									  </template>
								</select>
						  	</div>-->
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Kabupaten/Kota</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 v-if="!pending && !pendingRef" ref="refKkab" id="kabs" :options="kabs" keyList="kkab" namaList="nwil" v-model="dataInduk.alkokab" />
						  	<!--<div class="sm:flex">
						  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.alkokab">
								  	<option selected>Pilih Kabupaten/Kota</option>
								  	<template v-for="(item, idx) in kabs" :key="idx">
									  	<option :value="item.kkab">{{item.kkab}} - {{item.nwil}}</option>
									  </template>
								</select>
						  	</div>-->
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Kecamatan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 v-if="!pending && !pendingRef" ref="refKkec" id="kecs" :options="kecs" keyList="kkec" namaList="nwil" v-model="dataInduk.alkokec" />
						  	<!--<div class="sm:flex">
						  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.alkokec">
								  	<option selected>Pilih Kecamatan</option>
								  	<template  v-for="(item, idx) in kecs" :key="idx">
									  	<option :value="item.kkec">{{item.kkec}} - {{item.nwil}}</option>
									  </template>
								</select>
						  	</div>-->
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Desa/Kelurahan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 v-if="!pending && !pendingRef" ref="refKkel" id="kels" :options="kels" keyList="tk_kdesa" namaList="nwil" v-model="dataInduk.alkodes" />
						  	<!--<div class="sm:flex">
						  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.alkodes">
								  	<option selected>Pilih Desa/Kelurahan</option>
								  	<template  v-for="(item, idx) in kels" :key="idx">
									  	<option :value="item.tkdesa+item.kdesa">{{item.tkdesa+item.kdesa}} - {{item.nwil}}</option>
									  </template>
								</select>
						  	</div>-->
						</div>
						<!-- End Col -->

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Nomor Telpon/ WA</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
						  	<div class="sm:flex gap-2">
						  		<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="Nomor HP/Telpon" v-model="dataInduk.altelp">
						  		<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="Nomor Whatsapp" v-model="dataInduk.altelpwa">
						  	</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Email</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
						  	<div class="sm:flex gap-2">
						  		<input id="af-account-full-name" type="email" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="Email Pribadi" v-model="dataInduk.email">
						  		<input id="af-account-full-name" type="email" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="Email Pemerintahan" v-model="dataInduk.email_pemerintahan">
						  	</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Status Kepegawaian</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
			        		<SearchSelect2 v-if="!pending && !pendingRef" id="statuss" :options="statuss" keyList="kstatus" namaList="nama" v-model="dataInduk.kstatus" />
			        	</div>
						<!--<div class="sm:col-span-9">
						  	<div class="sm:flex">
						  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kstatus">
								  	<option selected>Pilih Status Kepegawaian</option>
								  	<template  v-for="(item, idx) in statuss" :key="idx">
									  	<option :value="item.kstatus">{{item.kstatus}} - {{item.nama}}</option>
									  </template>
								</select>
						  	</div>
						</div>-->
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Jenis Kepegawaian</label>
						</div>
						<!-- End Col -->
						<div class="sm:col-span-9">
			        		<SearchSelect2 v-if="!pending && !pendingRef" id="jeniss" :options="jeniss" keyList="id" namaList="nama" v-model="dataInduk.kjpeg" />
			        	</div>
						<!--<div class="sm:col-span-9">
						  	<div class="sm:flex">
						  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kjpeg">
								  	<option selected>Pilih Jenis Kepegawaian</option>
								  	<template  v-for="(item, idx) in jeniss" :key="idx">
									  	<option :value="item.id">{{item.id}} - {{item.nama}}</option>
									  </template>
								</select>
						  	</div>
						</div>-->
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Kedudukan Kepegawaian</label>
						</div>
						<!-- End Col -->
						<div class="sm:col-span-9">
			        		<SearchSelect2 v-if="!pending && !pendingRef" id="duduks" :options="duduks" keyList="id" namaList="nama" v-model="dataInduk.kduduk" />
			        	</div>
						<!--<div class="sm:col-span-9">
						  	<div class="sm:flex">
						  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kduduk">
								  	<option selected>Pilih Kedudukan Kepegawaian</option>
								  	<template  v-for="(item, idx) in duduks" :key="idx">
									  	<option :value="item.id">{{item.id}} - {{item.nama}}</option>
									  </template>
								</select>
						  	</div>
						</div>-->
						<!-- End Col -->

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Nomor Karpeg</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<div class="sm:flex">
						  		<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.nkarpeg">
						  	</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-2 sm:col-start-8">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">No Karis/Karsu</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<div class="sm:flex">
						  		<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.nkaris_su">
						  	</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Nomor Taspen</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
						  	<div class="sm:flex">
						  		<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.ntaspen">
						  	</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Status Kepemilikan KPE</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<div class="sm:flex">
						  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.stat_kpe">
								  	<option selected>Pilih Status KPE</option>
								  	<template  v-for="(item, idx) in kpes" :key="idx">
									  	<option :value="item.kkpe">{{item.kkpe}} - {{item.nkpe}}</option>
									  </template>
								</select>
						  	</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-2 sm:col-start-8">
						  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Tanggal KPE</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
						  	<div class="sm:flex">
						  		<input id="af-account-full-name" type="date" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tgl_kpe">
						  	</div>
						</div>
						<!-- End Col -->
		 			</div>
		  			<!-- End Grid -->

		  			<div class="mt-5 flex justify-end gap-x-2">
						<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none" @click="refresh">
			  				Batal
						</button>
						<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="simpanData">
			  				{{method}} Data
						</button>
		  			</div>
				</form>
	  		</NewCard>
	  		<!-- End Card -->
		</div>
		<!-- End Card Section -->
	</LayoutDataInduk>
</template>