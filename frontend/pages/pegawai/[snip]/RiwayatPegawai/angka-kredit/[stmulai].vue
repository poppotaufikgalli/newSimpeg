<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';
import { useAuthStore } from '~/store/auth'; // import the auth store we just created

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)
const { userInfo } = useAuthStore();
const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
import customParseFormat from "dayjs/plugin/customParseFormat";
dayjs.extend(customParseFormat)

const showSpinner = ref(false)
const toast = useToast()
const {snip, stmulai} = route.params
const {sbkn} = route.query

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
	filename_dok_pak: '',
	filename_sk_pak: '',
	idSync:'',
	rw_jabatan_id:'',
	rw_jabatan_id_bkn:'',
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

		if(sbkn != ""){
			let sid = $decodeBase64(sbkn)
			var result = await $fetch(`/api/gets/singkronisasi_bkn/getDataBKNById/angkakredit/${sid}`)
			var item = result.data
			let tsk = dayjs(dayjs(item.tanggalSk, "DD-MM-YYYY")).format("YYYY-MM-DD[T]HH:mm:ss[Z]")
			let tmulai = dayjs(item.tahunMulaiPenailan+"-"+item.bulanMulaiPenailan+"-1").format("YYYY-MM-DD[T]HH:mm:ss[Z]")
			let tselesai = dayjs(item.tahunSelesaiPenailan+"-"+item.bulanSelesaiPenailan+"-1").format("YYYY-MM-DD[T]HH:mm:ss[Z]")
			let jns = ""
			if(item.isAngkaKreditPertama == "1") {jns = "Pertama"}
			else if(item.isIntegrasi == "1") {jns ="Integrasi"}
			else if(item.isKonversi == "1") {jns ="Konversi"}
			else {jns = "Lanjutan"}

			dataAngkaKredit.value = {
				nip: nip,
				jns: jns,
				nsk: result.data.nomorSk,
				tsk: tsk,
				kjab: '',
				tmulai: tmulai,
				tselesai: tselesai,
				utama: result.data.kreditUtamaBaru,
				tambahan: result.data.kreditPenunjangBaru,
				total: result.data.kreditBaruTotal,
				idSync: sid,
				rw_jabatan_id:'',
				rw_jabatan_id_bkn:result.data.rwJabatan,
			}
			
		}else{
			let tmulai = $decodeBase64(stmulai)
			tmulai = dayjs(tmulai, 'YYYY-MM-DD').format("YYYY-MM-DD")
			var result = await $fetch(`/api/gets/riwayat_angka_kredit/${nip}/${tmulai}`);

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

const utama = computed({
	get(){
		return dataAngkaKredit.value.utama;
	},
	set(val){
		dataAngkaKredit.value.utama = val
		dataAngkaKredit.value.total = (val*1) + (dataAngkaKredit.value.tambahan *1)
	}
})

const tambahan = computed({
	get(){
		return dataAngkaKredit.value.tambahan;
	},
	set(val){
		dataAngkaKredit.value.tambahan = val
		dataAngkaKredit.value.total = (dataAngkaKredit.value.utama *1) + (val*1)
	}
})

const total = computed({
	get(){
		//dataAngkaKredit.value.total = (dataAngkaKredit.value.utama *1) + (dataAngkaKredit.value.tambahan *1)
		return dataAngkaKredit.value.total;
	},
	set(val){
		dataAngkaKredit.value.total = val
	}
})

const selkjab = computed({
	get(){
		return dataAngkaKredit.value.kjab
	},
	set(val){
		dataAngkaKredit.value.kjab = val

		var selJab = _find(jabfungs.value, function(o){
			return o.kjab == val
		})

		dataAngkaKredit.value.rw_jabatan_id = selJab.id
		dataAngkaKredit.value.rw_jabatan_id_bkn = selJab.idSync
	}
})

onMounted(() => {
	refreshNuxtData(["getDataRef", "getdataAngkaKreditAkhir"])
	modalUploadDoc.$onAction(callback, false)
})

async function simpanData() {
	showSpinner.value = true

	var compacted = _pickBy(dataAngkaKredit.value);
	
	compacted.tsk = compacted.tsk ? dayjs(dayjs(compacted.tsk, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tmulai = compacted.tmulai ? dayjs(dayjs(compacted.tmulai, "YYYY-MM")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tselesai = compacted.tselesai ? dayjs(dayjs(compacted.tselesai, "YYYY-MM")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.utama = compacted.utama ? compacted.utama *1 : 0
	compacted.tambahan = compacted.tambahan ? compacted.utama *1 : 0
	compacted.total = compacted.total ? compacted.total *1 : 0

	var dr = dataRefAngkaKredit.value

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
			const {title, description} = error.value.data.data
			toast.add({
				id: 'error_put_angka_kredit',
				icon: "i-heroicons-x-circle",
				title: title,
				description: description,
				timeout: 6000,
				color: 'red',
			}) 	
		}else{
			toast.add({
				id: 'success_put_angka_kredit',
				icon: "i-heroicons-check-circle",
				title: "Update Data Berhasil",
				description: "Data telah di Update",
				timeout: 6000
			}) 	

			sendDataBKN(JSON.stringify(compacted))
			navigateTo({ path: '/pegawai/'+snip+'/RiwayatPegawai/angka-kredit/'})
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		var body = JSON.stringify(compacted)


		var {data, error} = await useFetch('/api/posts/riwayat_angka_kredit/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_angka_kredit',
		    	icon: "i-heroicons-x-circle",
		    	title: error.value.data.data.title ?? "Gagal Menambahkan Data Angka Kredit",
		    	description: error.value.data.data.error,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_angka_kredit',
		    	icon: "i-heroicons-check-circle",
		    	title: "Menambahkan Data Angka Kredit",
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	})

		  	sendDataBKN(body)
		  	navigateTo({ path: '/pegawai/'+snip+'/RiwayatPegawai/angka-kredit/'})
		}
	}

  	showSpinner.value = false
}

async function sendDataBKN(body){
	if(userInfo.gid == 1){
		var {data, error} = await useFetch('/api/posts/riwayat_angka_kredit/bkn', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_angka_kredit_bkn',
		    	icon: "i-heroicons-x-circle",
		    	title: "Data Angka Kredit BKN",
		    	description: error.value.data.data.error,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_angka_kredit_bkn',
		    	icon: "i-heroicons-check-circle",
		    	title: "Angka Kredit BKN",
		    	description: "Data berhasil di Dikirim",
		    	timeout: 6000
		  	})
		}	
	}
}

const selUploadDoc = ref("")
const selUploadDocJudul= ref("")
const selUploadDocBkn = ref("")

const doUploadDoc = async(judul, fileIdx, idxBkn="") => {
	let nip = $decodeBase64(snip)
	let idx = "filename_"+fileIdx;

	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataAngkaKredit.value[idx]}`)
	selUploadDoc.value = fileIdx
	selUploadDocBkn.value = idxBkn
	selUploadDocJudul.value = judul
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen Angka Kredit ${judul}`, "application/pdf", blobUrl, 'upload')	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen Angka Kredit ${judul}`, "application/pdf", null, 'upload')	
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
		formData.append("field_name", "filename_"+selUploadDoc.value)
		formData.append("filename", nip+"_"+selUploadDoc.value+tmulai.value);
		formData.append("path", 'dokumen/'+nip);

		formData.append("id_ref_dokumen", selUploadDocBkn.value)
		formData.append("idSync", dataAngkaKredit.value.idSync)

		var {data, error} = await useFetch(`/api/uploads/upload/ak/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(data.value.status != "200"){
			const {title, description} = data.value.data
			toast.add({
				id: 'error_post_upload_angka_kredit',
				icon: "i-heroicons-x-circle",
				title: title,
				description: description,
				timeout: 6000,
				color: 'red'
			}) 	
		}else{
			toast.add({
				id: 'success_post_upload_angka_kredit',
				icon: "i-heroicons-check-circle",
				title: `Berhasil Upload Dokumen ${selUploadDocJudul.value}`,
				description: `Dokumen berhasil di Upload`,
				timeout: 6000
			}) 	
		}

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
	<AppLoadingSpinner :show="showSpinner" />
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
							<SearchSelect2 ref="refJabfung" id="jabfungs" :options="jabfungs" keyList="kjab" namaList="njab" v-model="selkjab"/>
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
									<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="utama" @keypress="onlyNumber">
								</div>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<label class="inline-block text-sm text-gray-800 mt-2.5">Angka Kredit Penunjang</label>
							</div>
							<!-- End Col -->

							<div class="sm:col-span-3">
								<div class="sm:flex gap-2">
									<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tambahan" @keypress="onlyNumber">
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
								<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="simpanData()">
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
				<h2 class="text-xl font-bold text-blue-600">Dokumen Angka Kredit</h2>
			</div>
			<div class="grid sm:grid-cols-4 gap-2 gap-2.5">
				<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc('SK PAK','sk_pak', '879')">
					SK PAK
				</button>
				<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc('Dokumen PAK','dok_pak', '880')">
					Dok PAK
				</button>
			</div>
		</div>
		<!-- End Card Section -->
	</LayoutRiwayatPegawai>
</template>