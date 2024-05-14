<script setup>
const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const toast = useToast()
import customParseFormat from "dayjs/plugin/customParseFormat";
dayjs.extend(customParseFormat)

const {snip} = route.params
const result = ref(null)
const cb = ref([])
const cbSimpeg = ref([])
const cbBKN = ref([])
const showSpinner = ref(false)

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
	nakses: '',
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
})

const refDataBkn = ref({
	"nip": "nipBaru", 
	"nama": "nama", 
	"glblk": "gelarBelakang", 
	"gldepan": "gelarDepan", 
	"ktlahir": "tempatLahir",
	"tlahir": "tglLahir", 
	"nik": "nik", 
	"no_ref_bkn": "id", 
	"niplama": "nipLama", 
	"kagama": "agamaId", 
	"aljalan": "alamat", 
	"bpjs": "bpjs", 
	"email": "email", 
	"email_pemerintahan": "emailGov", 
	"kskawin": "jenisKawinId", 
	"kjkel": "jenisKelamin", 
	"kartu_asn": "kartuAsn", 
	"kduduk": "kedudukanPnsId", 
	"nkaris_su": "karis_karsu", 
	"kpos": "kodePos", 
	"naskes": "noAskes", 
	"altelp": "noTelp", 
	"altelpwa": "noHp",
	"npwp": "noNpwp", 
	"nkarpeg": "noSeriKarpeg", 
	"ntaspen": "noTaspen", 
	"tapera": "tapera",
})

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

const getDataBKN = async() => {
	if(snip){
		let nip = $decodeBase64(snip)
		var result_simpeg = await $fetch('/api/gets/pegawai/'+nip);
		var result_bkn = await $fetch('/api/gets/singkronisasi_bkn/getDataBKN/data-utama/'+nip)

		if(result_simpeg.nip == nip){
			var dataBkn = result_bkn.data
			console.log(dataBkn)
			//var dataBkn = dummyData.value.data
			result_simpeg.tlahir = dayjs(result_simpeg.tlahir).format("DD-MM-YYYY").toString()
			result_simpeg.kjkel = result_simpeg.kjkel == 1 ? '1/M' : '2/F'
			dataBkn.jenisKelamin = dataBkn.jenisKelamin == 'M' ? '1/M' : '2/F'
			result.value = _mapValues(refDataBkn.value, function(val,key){
				var res1 = _find(result_simpeg, function(val1,key1){
					return key1 == key
				})
				var res2 = _find(dataBkn, function(val1,key1){
					return key1 == val
				})

				return [res1, res2]
			})
		}
	}
}

const doUpdateBKN = async() => {
	showSpinner.value = true
	
	var compacted = _invert(cbBKN.value)
	compacted = _mapValues(compacted, function(val, key){
		return _find(result.value, function(val1, key1){
			return key1 == key
		})[0]
	})

	compacted = _mapValues(refUpdateBkn.value, function(val, key){
		return _find(compacted, function(val1, key1){
			return val == key1
		})
	})

	if(_includes(compacted, "pns_orang_id") == false){
		compacted.pns_orang_id = result.value['no_ref_bkn'][1]
	}
	
	var body = JSON.stringify(compacted)

	var {data, error} = await useFetch(`/api/puts/singkronisasi_bkn/putsDataBKN/data-utama`, {
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
  	showSpinner.value = false

  	getDataBKN()
}

const doUpdateSIMPEG = async() => {
	showSpinner.value = true
	var compacted = _invert(cbSimpeg.value)
	compacted = _mapValues(compacted, function(val, key){
		return _find(result.value, function(val1, key1){
			return key1 == key
		})[1]
	})
	
	compacted.nip = result.value['nip'][1]
	if(_includes(compacted, "nama") == false){
		compacted.nama = result.value['nama'][1]
	}

	//console.log(compacted)
	if(_includes(compacted, "tlahir") == true){
		compacted.tlahir = dayjs(dayjs(compacted.tlahir, "DD-MM-YYYY")).format("YYYY-MM-DD[T]HH:mm:ss[Z]")
	}else{
		compacted.tlahir = dayjs(dayjs(result.value['tlahir'][1], "DD-MM-YYYY")).format("YYYY-MM-DD[T]HH:mm:ss[Z]")
	}

	if(_includes(compacted, "kjkel") == true){
		compacted.kjkel = compacted.kjkel == '1/M' ? 1 : 2
	}

	//console.log(compacted)

	var body = JSON.stringify(compacted)

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
  	showSpinner.value = false

  	getDataBKN()
}

function canUpdateBKN(key){
	var nkey = refDataBkn.value[key]
	var result = _includes(_values(refUpdateBkn.value), key)
	return result
}

</script>
<template>
	<AppLoadingSpinner :show="showSpinner" />
	<LayoutSingkronisasi>
		<TokenInfo />
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="flex justify-between mb-6">
					<h2 class="text-xl font-bold text-blue-600">Identitas Pegawai</h2>
					<button class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="getDataBKN()">Get Data BKN</button>
				</div>

				<div class="mb-8">
					<div class="overflow-x-auto border rounded-lg">
						<div class="min-w-full inline-block align-middle">
							<!-- Table -->
							<table class="min-w-full divide-y divide-gray-200 text-sm">
								<thead class="bg-gray-50 h-12">
									<tr class="divide-x divide-gray-200">
										<!--<th scope="col" class="px-3 text-center" width="5%">
											<span class="text-xs font-semibold uppercase text-gray-800">
												&nbsp;
											</span>
										</th>-->
										<th scope="col" class="px-3 text-center" width="40%">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Item
											</span>
										</th>
										<th scope="col" class="px-3 text-center" width="30%">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Data Simpeg
											</span>
										</th>
										<th scope="col" class="px-3 text-center" width="30%">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Data BKN
											</span>
										</th>
									</tr>
								</thead>
								<tbody v-for="(item,key) in refDataBkn" v-if="result">
									<tr class="divide-x divide-gray-200" :class="result[key][0] == result[key][1]? 'bg-green-100' : 'bg-red-100'">
										<!--<td class="size-px px-1" style="vertical-align: top;">
											<div class="flex justify-center items-center h-5" v-if="result[key][0] != result[key][1]">
							                  	<input :id="`cb-${key}`" type="checkbox" class="border-gray-200 rounded text-blue-600 focus:ring-blue-500" v-model="cb" :value="key">
							                </div>
										</td>-->
										<td class="size-px px-1" style="vertical-align: top;">
											<label :for="`cb-${key}`">{{key}}</label>
										</td>
										<td class="size-px px-1" :class="canUpdateBKN(key) ? '':'bg-gray-100 italic'" style="vertical-align: top;">
											<div class="flex gap-2">
												<template v-if="result[key][0] != result[key][1] && result[key][0] != '' && result[key][0] != null && canUpdateBKN(key)">
													<input :id="`cb-${item}`" type="checkbox" class="border-gray-200 rounded text-blue-600 focus:ring-blue-500" v-model="cbBKN" :value="key">	
												</template>
												
												<label :for="`cb-${item}`">{{result[key][0]}}</label>
											</div>
										</td>
										<td class="size-px px-1" style="vertical-align: top;">
											<div class="flex gap-2">
												<template v-if="result[key][0] != result[key][1] && result[key][1] != '' && result[key][1] != null">
													<input :id="`cb-${key}`" type="checkbox" class="border-gray-200 rounded text-blue-600 focus:ring-blue-500" v-model="cbSimpeg" :value="key">
												</template>
												<label :for="`cb-${key}`">{{result[key][1]}}</label>
											</div>
										</td>
									</tr>
								</tbody>
								<tfoot v-if="result">
									<tr class="divide-x divide-gray-200">
										<td class="size-px px-1" style="vertical-align: top;"></td>
										<td class="size-px px-1 py-2 text-center" style="vertical-align: top;">
											<button type="button" class="p-1 inline-flex justify-center items-center gap-x-2 text-xs font-semibold rounded-md border border-transparent bg-gray-600 text-white hover:bg-gray-700 disabled:opacity-50 disabled:pointer-events-none whitespace-nowrap" @click="doUpdateBKN()">
												<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4">
												  	<path fill-rule="evenodd" d="M8 14a.75.75 0 0 0 .75-.75V4.56l1.22 1.22a.75.75 0 1 0 1.06-1.06l-2.5-2.5a.75.75 0 0 0-1.06 0l-2.5 2.5a.75.75 0 0 0 1.06 1.06l1.22-1.22v8.69c0 .414.336.75.75.75Z" clip-rule="evenodd" />
												</svg>
												Update to BKN
											</button>
										</td>
										<td class="size-px px-1 py-2 text-center" style="vertical-align: top;">
											<button type="button" class="p-1 inline-flex justify-center items-center gap-x-2 text-xs font-semibold rounded-md border border-transparent bg-gray-600 text-white hover:bg-gray-700 disabled:opacity-50 disabled:pointer-events-none whitespace-nowrap" @click="doUpdateSIMPEG()">
												<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4">
												  	<path fill-rule="evenodd" d="M8 2a.75.75 0 0 1 .75.75v8.69l1.22-1.22a.75.75 0 1 1 1.06 1.06l-2.5 2.5a.75.75 0 0 1-1.06 0l-2.5-2.5a.75.75 0 1 1 1.06-1.06l1.22 1.22V2.75A.75.75 0 0 1 8 2Z" clip-rule="evenodd" />
												</svg>
												Update to Simpeg
											</button>
										</td>
									</tr>
								</tfoot>
							</table>
							<!-- End Table -->	
						</div>
					</div>
				</div>
			</div>
		</div>
	</LayoutSingkronisasi>
</template>