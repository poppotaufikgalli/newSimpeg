<script setup>
    import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)

const route = useRoute()
const dayjs = useDayjs()
const toast = useToast()
const { $decodeBase64, $encodeBase64 } = useNuxtApp()

const {snip} = route.params

const showSpinner = ref(false)
const method = ref('Simpan')

const dataInduk = ref({
	nip: '',
	nama: '',
	gldepan: '',
	glblk: '',
	ktlahir: '',
	tlahir: '',
	kjkel: null,
	kagama: null,
	kstatus: 2,
	kjpeg: 15,
	kduduk: 1,
	kskawin: '',
	tinggi_badan:null,
	berat_badan:null,
	kgoldar: '',
	aljalan:'',
	alrt: '',
	alrw: '',
	altelp: '',
	altelpwa: '',
	alkoprop: "21",
	alkokab: "72",
	alkokec: '',
	alkodes: '',
	kpos:'',
	nkarpeg: '',
	naskes: '',
	ntaspen:'',
	nkaris_su:'',
	npwp: '',
	nik:'',
	aktif: '',
	jjiwa: '',
	marga: '',
	suku: '',
	stat_kpe: '',
	tgl_kpe: '',
	thn_pendataan: '',
	no_ref_bkn: '',
	email:'',
	email_pemerintahan: '',
	kartu_asn:'',
	bpjs:'',
	niplama:'',
	tapera:'',
});

const refDataBkn = ref([
	["nip","nipBaru",null, "NIP"],
	["nama","nama",null, "Nama"],
	["gldepan","gelarDepan",null, "Gelar Depan"],
	["glblk","gelarBelakang",null, "Gelar Belakang"],
	["ktlahir","tempatLahir",null, "Tempat Lahir"],
	["tlahir","tglLahir",null, "Tanggal Lahir"],
	["niplama","nipLama", null, "Nip Lama"],
	["nik","nik",null, "NIK"],	
    ["no_ref_bkn","id","pns_orang_id", "PNS ID"],
    ["kagama","agamaId","agama_id", "Agama"],
    ["aljalan","alamat","alamat", "Alamat"],
    ["bpjs","bpjs","nomor_bpjs", "BPJS"],
    ["email","email","email", "Email"],
    ["email_pemerintahan","emailGov","email_gov", "Email Pemerintahan"],
    ["kskawin","jenisKawinId",null, "Kawin"],
    ["kjkel","jenisKelamin",null, "Jenis Kelamin"],
    ["kartu_asn","kartuAsn",null, "Kartu ASN"],
    ["kduduk","kedudukanPnsId",null, "Jenis Kedudukan PNS"],
    ["nkaris_su","karis_karsu","karis_karsu", "Karis Karsu"],
    ["kpos","kodePos",null, "Kode Pos"],
    ["naskes","noAskes",null, "Askes"],
    ["altelp","noTelp","nomor_telpon", "Nomor Telpon"],
    ["altelpwa","noHp","nomor_hp", "Nomor HP"],
    ["npwp","noNpwp","npwp_nomor", "NPWP"],
    ["nkarpeg","noSeriKarpeg",null, "Karpeg"],
    ["ntaspen","noTaspen","taspen_nomor", "Taspen"],
    ["tapera","tapera","tapera_nomor", "Tapera"],
    ["kjpeg","jenisPegawaiId",null, "Jenis Pegawai"],
])

const refUpdateBkn = ref({
	"agama_id": "kagama",
	"alamat": "aljalan",
	"email": "email",
	"email_gov": "email_pemerintahan",
	"karis_karsu": "nkaris_su",
	"nomor_bpjs": "bpjs",
	"nomor_telpon": "altelp",
	"nomor_hp": "altelpwa",
	"npwp_nomor": "npwp",
	"pns_orang_id": "no_ref_bkn",
	"taspen_nomor": "ntaspen",
	"tapera_nomor": "tapera",
})

const dataBkn = ref({
	agamaId: '', 
	alamat: '', 
	bpjs: '', 
	email: '', 
	emailGov: '', 
	gelarBelakang: '', 
	gelarDepan: '', 
	id: '', 
	jenisKawinId: '', 
	jenisKelamin: '', 
	jenisPegawaiId: '', 
	karis_karsu: '', 
	kartuAsn: '', 
	kedudukanPnsId: '', 
	kodePos: '', 
	nama: '', 
	nik: '', 
	nipBaru: '', 
	nipLama: '', 
	noAskes: '', 
	noHp: '', 
	noNpwp: '', 
	noSeriKarpeg: '', 
	noTaspen: '', 
	noTelp: '', 
	tempatLahir: '', 
	tglLahir: ''
})
//const datax = ref({ "agamaId": "1", "alamat": "JL. ADI SUCIPTO KM. 11 PERUM BHUMI ANGGREK RESIDANCE BLOK G/3 TANJUNGPINANG", "bpjs": "0000161112159", "email": "taufik@live.com", "emailGov": "taufik@tanjungpinangkota.go.id", "gelarBelakang": "ST", "gelarDepan": "", "id": "A8ACA7CD3C7F3912E040640A040269BB", "jenisKawinId": "1", "jenisKelamin": "M", "jenisPegawaiId": "15", "karis_karsu": "", "kartuAsn": "A200900283492", "kedudukanPnsId": "01", "kodePos": "", "nama": "MUHAMMAD TAUFIK HIDAYAT", "nik": "2101071602860001", "nipBaru": "198602162008031001", "nipLama": "P20006592", "noAskes": "", "noHp": "085272220227", "noNpwp": "14.831.626.8-224.000", "noSeriKarpeg": "", "noTaspen": "198602162008031001", "noTelp": "085272220227", "tempatLahir": "TANJUNG UBAN", "tglLahir": "16-02-1986" })

const dataDiff = ref([])

const { pending, data, refresh} = await useAsyncData('getDataIdentitas', async() => {
	console.log("getDataIdentitasIdentitas")
	if(snip){
		showSpinner.value = true
		let nip = $decodeBase64(snip)

		const [data_simpeg, wilayah] = await Promise.all([
	    	$fetch('/api/gets/pegawai/'+nip),
	    	$fetch('/api/gets/wilayah'),
	  	])

	  	wilayahs.value = wilayah

	  	if(data_simpeg.nip == nip){
	  		method.value = "Update"

	  		dataInduk.value = _pickBy(data_simpeg, function(val, key) {
				return _includes(_keys(dataInduk.value), key);
			})
	  	}

	  	getDataBKN()
	  	
	  	showSpinner.value = false
	}
})

async function getDataBKN(){
	if(snip){
		let nip = $decodeBase64(snip)
		var result = await $fetch('/api/gets/singkronisasi_bkn/getDataBKN/data-utama/'+nip)
		
		//dataBkn.value = result.data
		dataBkn.value = _pickBy(result.data, function(val, key) {
			return _includes(_keys(dataBkn.value), key);
		})

		console.log(dataBkn.value)

	  	dataDiff.value = _reduce(refDataBkn.value, function(result, value, key) {
		  (result[value[3]] || (result[value[3]] = [])).push(
		  	_find(dataBkn.value, function(val,key){
		  		return key == value[1]
		  	}),
		  	_find(dataInduk.value, function(val,key){
		  		return key == value[0]
		  	})
		  );
		  return result;
		}, {});

		//console.log(dataDiff.value)
	}
}

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

const selUploadDoc = ref(null)

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useAsyncData('getDataRef', async () => {
  	const [agama, status, jenis, duduk, kawin, goldar, kpe] = await Promise.all([
    	$fetch('/api/gets/agama'),
    	$fetch('/api/gets/status_pegawai'),
    	$fetch('/api/gets/jenis_pegawai'),
    	$fetch('/api/gets/kedudukan_pegawai'),
    	$fetch('/api/gets/jenis_kawin'),
    	$fetch('/api/gets/jenis_goldar'),
    	$fetch('/api/gets/kpe'),
  	])

  	agamas.value = agama
  	statuss.value = status
  	jeniss.value = jenis
  	duduks.value = duduk
  	kawins.value = kawin
  	goldars.value = goldar
  	kpes.value = kpe
});

onMounted(() => {
	refreshNuxtData(["getDataIdentitas", "getDataRef"])
	modalUploadDoc.$onAction(callback, false)
})

const provs = computed({
	get(){
		var result = _filter(wilayahs.value, function(o){
			return o.twil == 1
		})
		return result
	}
})

const alkoprop = computed({
	get(){
		return dataInduk.value.alkoprop
	},
	set(val){
		dataInduk.value.alkoprop = val

		if(refKkab){
			refKkab.value.reinit()
		}

		alkokab.value = null
	}
})

const kabs = computed({
	get(){
		var result = _filter(wilayahs.value, function(o){
			return o.twil == 2 && o.kprov == alkoprop.value
		})
		return result
	}
})

const alkokab = computed({
	get(){
		return dataInduk.value.alkokab
	},
	set(val){
		dataInduk.value.alkokab = val

		if(refKkec){
			refKkec.value.reinit()
		}

		alkokec.value = null
	}
})

const kecs = computed({
	get(){
		var result = _filter(wilayahs.value, function(o){
			return o.twil == 3 && o.kprov == alkoprop.value && o.kkab == alkokab.value
		})
		return result
	}
})

const alkokec = computed({
	get(){
		return dataInduk.value.alkokec
	},
	set(val){
		dataInduk.value.alkokec = val

		if(refKkel){
			refKkel.value.reinit()
		}

		dataInduk.value.alkodes = null
	}
})

const kels = computed({
	get(){
		var result = _filter(wilayahs.value, function(o){
			return o.twil == 4 && o.kprov == alkoprop.value && o.kkab == alkokab.value && o.kkec == alkokec.value
		})
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

const tlahir2 = computed({
	get(){
		return dayjs(dataInduk.value.tlahir).format("DD-MM-YYYY").toString()
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

const kjkel = computed({
	get(){
		return dataInduk.value.kjkel
	},
	set(val){
		dataInduk.value.kjkel = val *1
	}
})

const kjkel2 = computed({
	get(){
		return dataInduk.value.kjkel == 1 ? 'M' : 'F'
	}
})

const kagama = computed({
	get(){
		return dataInduk.value.kagama
	},
	set(val){
		dataInduk.value.kagama = val *1
	}
})

const kskawin = computed({
	get(){
		return dataInduk.value.kskawin
	},
	set(val){
		dataInduk.value.kskawin = val *1
	}
})

const kgoldar = computed({
	get(){
		return dataInduk.value.kgoldar
	},
	set(val){
		dataInduk.value.kgoldar = val *1
	}
})

const stat_kpe = computed({
	get(){
		return dataInduk.value.stat_kpe
	},
	set(val){
		dataInduk.value.stat_kpe = val *1
	}
})

async function simpanData(lvl) {

	//1 operator opd
	//2 admin simpeg
	//3 update bkn

	showSpinner.value = true

	var compacted = _pickBy(dataInduk.value);
	compacted.tlahir = dayjs(dayjs(compacted.tlahir, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]")
	compacted.tgl_kpe = compacted.tgl_kpe ? dayjs(dayjs(compacted.tgl_kpe, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null

	var body = JSON.stringify(compacted)

	if(method.value == 'Update'){
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
	}else{
		var {data, error} = await useFetch('/api/pegawai/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_identitas',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_identitas',
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
		  	const snip = $encodeBase64(compacted.nip)
			navigateTo({ path: '/pegawai/'+snip+'/DataInduk/identitas' })
		}
	}

	if(lvl == 2){
		//====================update to BKN====================//

		var sendDataToBKN = _pickBy(_mapValues(refUpdateBkn.value, function(val, key){
			return _find(compacted, function(val1, key1){
				return val == key1
			})?.toString()
		}))
		
		var {data, error} = await useFetch(`/api/puts/singkronisasi_bkn/putsDataBKN/data-utama`, {
			method: 'PUT',
			body: JSON.stringify(sendDataToBKN),
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_identitas_bkn',
		    	description: error.value.data.data,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_identitas_bkn',
		    	description: data.value.code == "1" ? "Data BKN Update Success |" : "Data BKN Update Failed |"  +data.value.message,
		    	timeout: 6000,
		    	color:  data.value.code == "1" ? 'green' : 'red',
		  	}) 	
		}

		//==================end update to BKN==================//
	}
  	showSpinner.value = false
}

function onlyNumber ($event) {
   	let keyCode = ($event.keyCode ? $event.keyCode : $event.which);
   	if ((keyCode < 48 || keyCode > 57) && keyCode !== 46) { // 46 is dot
    	$event.preventDefault();
   	}
}

const doUploadDoc = async(caption, doc) => {
	let nip = $decodeBase64(snip)
	selUploadDoc.value = doc
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${nip}_${doc}.pdf`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen ${caption}`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen ${caption}`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e
	console.log(e)

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_"+selUploadDoc.value);
		formData.append("path", 'dokumen/'+nip);

		var {data, error} = await useFetch(`/api/uploads/upload/pegawai/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(error.value){
			toast.add({
		    	id: 'error_post_upload_'+selUploadDoc.value,
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_'+selUploadDoc.value,
		    	description: `Dokumen ${selUploadDoc.value} berhasil di Upload`,
		    	timeout: 6000
		  	}) 	
		}

		showSpinner.value = false
	}
}
function checkBeda(key, item) {
	var a = item[0]
	var b = item[1]
	if(key == 'Jenis Kelamin'){
		a = a == 'M' ? '1' : '2'
	}
	if(key == 'Tanggal Lahir'){
		b = dayjs(dayjs(b, "YYYY-MM-DD")).format("DD-MM-YYYY")
	}
	if(a != b){
		return true
	}else{
		return false
	}
}
</script>
<style scoped>
    .match {
  	@apply border-red-200;
}
</style>
<template>
    <AppLoadingSpinner :show="showSpinner" />
    <LayoutDataInduk>
        <div class="mx-auto">
            <!-- Card -->
            <div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
                <div class="mb-4">
                    <div class="flex justify-between gap-x-2">
                        <h2 class="text-xl font-bold text-blue-600">Identitas Pegawai</h2>
                        <button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="getDataBKN">
                            Get Data BKN
                        </button>
                    </div>
                </div>
                <div class="bg-red-50 border border-red-200 text-sm text-red-800 rounded-lg p-4 mb-4" role="alert">
                    <div class="flex">
                        <div class="flex-shrink-0">
                            <svg class="flex-shrink-0 size-4 mt-0.5" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                <circle cx="12" cy="12" r="10"></circle>
                                <path d="m15 9-6 6"></path>
                                <path d="m9 9 6 6"></path>
                            </svg>
                        </div>
                        <div class="ms-4">
                            <h3 class="text-sm font-semibold">
                                Perbedaan Data BKN
                            </h3>
                            <div class="mt-2 text-sm text-red-700">
                                <ul class="list-disc space-y-1 ps-5" v-if="dataDiff" v-for="(item, key) in dataDiff">
                                    <li v-if="checkBeda(key, item)">
                                        {{key}} - Berbeda (data BKN : {{item[0]}})
                                    </li>
                                </ul>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="grid sm:grid-cols-12 gap-2 gap-2.5">
                    <div class="sm:col-span-3">
                        <label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">NIP</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-7">
                        <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" maxlength="18" v-model="dataInduk.nip" @keypress="onlyNumber" :class="dataInduk.nip != dataBkn.nipBaru ? 'border-red-200' : 'border-gray-200'">
                    </div>
                    <div class="sm:col-span-2">
                        <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.niplama" :class="dataInduk.niplama != dataBkn.nipLama ? 'border-red-200' : 'border-gray-200'">
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">ID PNS BKN</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-9">
                        <div class="sm:flex gap-2">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="NIK" v-model="dataInduk.no_ref_bkn" :class="dataInduk.no_ref_bkn != dataBkn.id ? 'border-red-200' : 'border-gray-200'">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Nama</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-9">
                        <input type="text" class="py-2 px-3 block w-full border shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white uppercase" :class="dataInduk.nama != dataBkn.nama ? 'border-red-200' : 'border-gray-200'" v-model="dataInduk.nama">
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Gelar</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-9">
                        <div class="sm:flex gap-2">
                            <input type="text" class="py-2 px-3 block w-[35%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="gelar depan" v-model="dataInduk.gldepan" :class="dataInduk.gldepan != dataBkn.gelarDepan ? 'border-red-200' : 'border-gray-200'">
                            <input type="text" class="py-2 px-3 block w-[35%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="Gelar Belakang" v-model="dataInduk.glblk" :class="dataInduk.glblk != dataBkn.gelarBelakang ? 'border-red-200' : 'border-gray-200'">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Tempat / Tgl Lahir</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-9">
                        <div class="sm:flex gap-2">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white uppercase" v-model="dataInduk.ktlahir" :class="dataInduk.ktlahir != dataBkn.tempatLahir ? 'border-red-200' : 'border-gray-200'">
                            <input type="date" class="py-2 px-3 block w-[50%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tlahir" :class="tlahir2 != dataBkn.tglLahir ? 'border-red-200' : 'border-gray-200'">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Jenis Kelamin</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <div class="sm:flex">
                            <select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kjkel" :class="kjkel2 != dataBkn.jenisKelamin ? 'border-red-200' : 'border-gray-200'">
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
                            <select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kagama" :class="dataInduk.kagama != dataBkn.agamaId ? 'border-red-200' : 'border-gray-200'">
                                <option selected>Pilih Agama</option>
                                <template v-for="(item, idx) in agamas" :key="idx">
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
                            <select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kskawin" :class="dataInduk.kskawin != dataBkn.jenisKawinId ? 'border-red-200' : 'border-gray-200'">
                                <option selected>Pilih Status Perkawinan</option>
                                <template v-for="(item, idx) in kawins" :key="idx">
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
                            <select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kgoldar">
                                <option selected>Pilih Golongan Darah</option>
                                <template v-for="(item, idx) in goldars" :key="idx">
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
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.tinggi_badan" @keypress="onlyNumber">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-2 sm:col-start-8">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Berat Badan</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <div class="sm:flex">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.berat_badan" @keypress="onlyNumber">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>
                    <div class="sm:col-span-3">
                        <label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Alamat</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-9">
                        <textarea class="py-2 px-3 block w-full border rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" rows="3" placeholder="Alamat..." v-model="dataInduk.aljalan" :class="dataInduk.aljalan != dataBkn.alamat ? 'border-red-200' : 'border-gray-200'"></textarea>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">RT/RW</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <div class="sm:flex gap-2">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="RT" v-model="dataInduk.alrt" @keypress="onlyNumber" maxlength="3">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="RW" v-model="dataInduk.alrw" @keypress="onlyNumber" maxlength="3">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-2 sm:col-start-8">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Kode Pos</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <div class="sm:flex">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.kpos" @keypress="onlyNumber" :class="dataInduk.kpos != dataBkn.kodepos ? 'border-red-200' : 'border-gray-200'">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Provinsi</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-9">
                        <SearchSelect2 v-if="!pending" ref="refKprov" id="provs" :options="provs" keyList="kprov" namaList="nwil" v-model="alkoprop" />
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Kabupaten/Kota</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-9">
                        <SearchSelect2 v-if="!pending" ref="refKkab" id="kabs" :options="kabs" keyList="kkab" namaList="nwil" v-model="alkokab" />
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Kecamatan</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-9">
                        <SearchSelect2 v-if="!pending" ref="refKkec" id="kecs" :options="kecs" keyList="kkec" namaList="nwil" v-model="alkokec" />
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Desa/Kelurahan</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-9">
                        <SearchSelect2 v-if="!pending" ref="refKkel" id="kels" :options="kels" keyList="tk_kdesa" namaList="nwil" v-model="dataInduk.alkodes" />
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>
                    <div class="sm:col-span-3">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Nomor Telpon/ WA</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-9">
                        <div class="sm:flex gap-2">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="Nomor HP/Telpon" v-model="dataInduk.altelp" @keypress="onlyNumber" :class="dataInduk.altelp != dataBkn.noTelp ? 'border-red-200' : 'border-gray-200'">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="Nomor Whatsapp" v-model="dataInduk.altelpwa" @keypress="onlyNumber" :class="dataInduk.altelpwa != dataBkn.noHp ? 'border-red-200' : 'border-gray-200'">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Email</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-9">
                        <div class="sm:flex gap-2">
                            <input type="email" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="Email Pribadi" v-model="dataInduk.email" :class="dataInduk.email != dataBkn.email ? 'border-red-200' : 'border-gray-200'">
                            <input type="email" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="Email Pemerintahan" v-model="dataInduk.email_pemerintahan" :class="dataInduk.email_pemerintahan != dataBkn.emailGov ? 'border-red-200' : 'border-gray-200'">
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
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Jenis Kepegawaian</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-9">
                        <SearchSelect2 v-if="!pending && !pendingRef" id="jeniss" :options="jeniss" keyList="id" namaList="nama" v-model="dataInduk.kjpeg" />
                    </div>
                    <div class="sm:col-span-3">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Kedudukan Kepegawaian</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-9">
                        <SearchSelect2 v-if="!pending && !pendingRef" id="duduks" :options="duduks" keyList="id" namaList="nama" v-model="dataInduk.kduduk" :class="dataInduk.kduduk != dataBkn.kedudukanPnsId ? 'border-red-200' : 'border-gray-200'" />
                    </div>
                    <div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>
                    <div class="sm:col-span-3">
                        <label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">NIK/NPWP</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <div class="sm:flex gap-2">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="NIK" v-model="dataInduk.nik" @keypress="onlyNumber" :class="dataInduk.nik != dataBkn.nik ? 'border-red-200' : 'border-gray-200'">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-2 sm:col-start-8">
                        <label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">NPWP 16 Digit</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <div class="sm:flex gap-2">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="NPWP" v-model="dataInduk.npwp" @keypress="onlyNumber" :class="dataInduk.npwp != dataBkn.noNpwp ? 'border-red-200' : 'border-gray-200'">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Nomor Karpeg</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <div class="sm:flex">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.nkarpeg" :class="dataInduk.nkarpeg != dataBkn.noSeriKarpeg ? 'border-red-200' : 'border-gray-200'">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-2 sm:col-start-8">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">No Karis/Karsu</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <div class="sm:flex">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.nkaris_su" :class="dataInduk.nkaris_su != dataBkn.karis_karsu ? 'border-red-200' : 'border-gray-200'">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Nomor Taspen</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <div class="sm:flex">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.ntaspen" :class="dataInduk.ntaspen != dataBkn.noTaspen ? 'border-red-200' : 'border-gray-200'">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-2 sm:col-start-8">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Tapera</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <div class="sm:flex">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.tapera">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Status Kepemilikan KPE</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <div class="sm:flex">
                            <select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="stat_kpe">
                                <option selected>Pilih Status KPE</option>
                                <template v-for="(item, idx) in kpes" :key="idx">
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
                            <input type="date" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tgl_kpe">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">ASKES</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <div class="sm:flex">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.naskes" :class="dataInduk.naskes != dataBkn.noAskes ? 'border-red-200' : 'border-gray-200'">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-2 sm:col-start-8">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">BPJS</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <div class="sm:flex">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.bpjs" :class="dataInduk.bpjs != dataBkn.bpjs ? 'border-red-200' : 'border-gray-200'">
                        </div>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Kartu ASN</label>
                    </div>
                    <!-- End Col -->
                    <div class="sm:col-span-3">
                        <div class="sm:flex">
                            <input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataInduk.kartu_asn" :class="dataInduk.kartu_asn != dataBkn.kartuAsn ? 'border-red-200' : 'border-gray-200'">
                        </div>
                    </div>
                    <!-- End Col -->
                </div>
                <!-- End Grid -->
                <div class="mt-5 flex justify-end gap-x-2">
                    <button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none" @click="refresh">
                        Batal
                    </button>
                    <button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-500 text-white hover:bg-blue-600 disabled:opacity-50 disabled:pointer-events-none" @click="simpanData(0)">
                        {{method}} Data [0]
                    </button>
                    <button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="simpanData(1)">
                        {{method}} Data [1]
                    </button>
                    <button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-700 text-white hover:bg-blue-800 disabled:opacity-50 disabled:pointer-events-none" @click="simpanData(2)">
                        {{method}} Data [2]
                    </button>
                </div>
            </div>
            <!-- End Card -->
            <div class="bg-white rounded-xl shadow py-4 px-6 border-t-2 mt-6" v-if="method == 'Update'">
                <div class="mb-8">
                    <h2 class="text-xl font-bold text-blue-600">Dokumen Pegawai</h2>
                </div>
                <div class="grid sm:grid-cols-4 gap-2 gap-2.5">
                    <button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc('KTP','ktp')">
                        KTP
                    </button>
                    <button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc('NPWP','npwp')">
                        NPWP
                    </button>
                    <button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc('TASPEN','taspen')">
                        TASPEN
                    </button>
                    <button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc('Perbaikan Nama', 'perbaikan_nama')">
                        Perbaikan Nama
                    </button>
                    <button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc('BPJS','akses')">
                        BPJS
                    </button>
                    <button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc('KARPEG','karpeg')">
                        KARPEG
                    </button>
                    <button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc('KPE','kpe')">
                        KPE
                    </button>
                    <button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc('KONVERSI_NIP','konversi')">
                        Konversi NIP
                    </button>
                </div>
            </div>
        </div>
        <!-- End Card Section -->
    </LayoutDataInduk>
</template>