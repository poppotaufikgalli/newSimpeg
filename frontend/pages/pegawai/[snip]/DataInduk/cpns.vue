<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)
const { $decodeBase64, $encodeBase64 } = useNuxtApp()

const route = useRoute()
const dayjs = useDayjs()
const {snip} = route.params

const dataRefCpns = ref({
	nip: '',
	knpang: 211,
	tmtpang: '',
	kgolru: '',
})

const dataCpns = ref({
	nip: '',
	nntbakn: '',
	tntbakn: '',
	ptetap: '',
	nskpang:'',
	tskpang: '',
	tmtpang: '',
	kgolru: '',
	tmtlgas: '',
	mskerja: 0,
	blnkerja: 0,
	srtsehatno: '',
	srtsehattgl: '',
	nsttpp: '',
	tsttpp: '',
	filename: '',
})

const pejabats = ref([])
const pangkats = ref([])
const method = ref('Simpan')

const refPejabat = ref(null)
const refPangkat = ref(null)

const { pending, data, refresh} = await useAsyncData('getDataCPNS', async() => {
	if(snip){
		let nip = $decodeBase64(snip)
		//dataCpns.value = await $fetch('/api/gets/riwayat_pangkat/'+nip+'/cpns');
		var result = await $fetch('/api/gets/riwayat_pangkat/'+nip+'/cpns');

		if(result.nip == nip){
			method.value = "Update"

			//var cpnskey = _keys(dataCpns.value)
			dataCpns.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataCpns.value), key);
			})

			//var Refkey = _keys(dataRefCpns.value)
			dataRefCpns.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataRefCpns.value), key);
			})
		}
	}
})

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
  	const [pejabat, pangkat] = await Promise.all([
    	$fetch('/api/gets/pejabat'),
    	$fetch('/api/gets/pangkat'),
  	])

  	pejabats.value = pejabat
  	pangkats.value = pangkat

  	if(refPejabat){
		refPejabat.value.reinit()
	}

	if(refPangkat){
		refPangkat.value.reinit()
	}
});

onMounted(() => {
	refreshNuxtData(["getDataCPNS","getDataRef"])
	modalUploadDoc.$onAction(callback, false)
})

const tntbakn = computed({
	get(){
		return dayjs(dataCpns.value.tntbakn).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataCpns.value.tntbakn = val
	}
})

const tskpang = computed({
	get(){
		return dayjs(dataCpns.value.tskpang).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataCpns.value.tskpang = val
	}
})

const tmtpang = computed({
	get(){
		return dayjs(dataCpns.value.tmtpang).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataCpns.value.tmtpang = val
	}
})

const tmtlgas = computed({
	get(){
		return dayjs(dataCpns.value.tmtlgas).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataCpns.value.tmtlgas = val
	}
})

const srtsehattgl = computed({
	get(){
		return dayjs(dataCpns.value.srtsehattgl).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataCpns.value.srtsehattgl = val
	}
})

const tsttpp = computed({
	get(){
		return dayjs(dataCpns.value.tsttpp).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataCpns.value.tsttpp = val
	}
})

const showSpinner = ref(false)
const toast = useToast()

async function simpanData(lvl) {
	showSpinner.value = true

	var compacted = dataCpns.value;
	compacted.tntbakn = compacted.tntbakn ? dayjs(dayjs(compacted.tntbakn, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tskpang = compacted.tskpang ? dayjs(dayjs(compacted.tskpang, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tmtpang = compacted.tmtpang ? dayjs(dayjs(compacted.tmtpang, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tmtlgas = compacted.tmtlgas ? dayjs(dayjs(compacted.tmtlgas, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.srtsehattgl = compacted.srtsehattgl ? dayjs(dayjs(compacted.srtsehattgl, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tsttpp = compacted.tsttpp ? dayjs(dayjs(compacted.tsttpp, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.kgolru = compacted.kgolru *1
	compacted.ptetap = compacted.ptetap *1
	compacted.mskerja = compacted.mskerja *1
	compacted.blnkerja = compacted.blnkerja *1
	
	compacted.knpang = 211 //knpang cpns = 211

	var dr = dataRefCpns.value
	dr.kgolru = dr.kgolru *1

	if(method.value == 'Update'){
		var body = {
			"refdata" : dr,
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_pangkat', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		console.log(error)

		if(error.value){
			toast.add({
		    	id: 'error_put_cpns',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_cpns',
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_pangkat/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_cpns',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_cpns',
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
		}
	}
  	showSpinner.value = false

  	if(lvl == 2){
  		//save cpns bkn
  		let dataCpnsBkn = {
  			nomor_sk_cpns : dataCpns.value.nskpang,
			nomor_spmt	: dataCpns.value.nomor_spmt,
			nomor_sttpl	: dataCpns.value.nsttpp,
			pertek_cpns_pns_l2th_nomor	: dataCpns.value.nntbakn,
			pertek_cpns_pns_l2th_tanggal	: dataCpns.value.tntbakn,
			//pns_orang_id	:dataCpns.value.no_ref_bkn,
			status_cpns_pns	: "cpns",
			tgl_sk_cpns	: dataCpns.value.tskpang,
			tgl_sttpl	: dataCpns.value.tsttpp,
  		}

	    let nip = $decodeBase64(snip)
	    var {data, error} = await useFetch('/api/puts/singkronisasi_bkn/putsDataBKNCpns/'+nip, {
	        method: 'POST',
	        body: JSON.stringify(dataCpnsBkn),
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
  		//end save cpns bkn
  	}

  	refreshNuxtData(["getDataRef", "getDataCPNS"])
}

const doUploadDoc = async() => {
	let nip = $decodeBase64(snip)
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataCpns.value.filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen CPNS`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen CPNS`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		formData.append("tmtpang", tmtpang.value);
		formData.append("kgolru", dataCpns.value.kgolru);
		formData.append("knpang", 211);

		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_cpns");
		formData.append("path", 'dokumen/'+nip);

		var {data, error} = await useFetch(`/api/uploads/upload/pangkat/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(error.value){
			toast.add({
		    	id: 'error_post_upload_cpns',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_cpns',
		    	description: `Dokumen CPNS berhasil di Upload`,
		    	timeout: 6000
		  	}) 	
		}

		refreshNuxtData(["getDataRef", "getDataCPNS"])

		showSpinner.value = false
	}
}
</script>
<template>
	<AppLoadingSpinner :show="showSpinner" />
	<LayoutDataInduk>
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="mb-8">
					<div class="inline-flex justify-between items-center w-full gap-x-2">
						<h2 class="text-xl font-bold text-blue-600">CPNS</h2>
						<div class="flex gap-x-2 w-[60%]">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP" :value="dataCpns.nip">
							<input type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="Nama">
						</div>
					</div>
				</div>
				<form>
					<!-- Grid -->
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nota BKN</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataCpns.nntbakn" placeholder="xx-xxxxxxxxxxxx">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tntbakn">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Pejabat Penetap</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 v-if="!pendingRef" ref="refPejabat" id="pejabats" :options="pejabats" keyList="kpej" namaList="npej" v-model="dataCpns.ptetap" />
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">SK CPNS</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataCpns.nskpang">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tskpang">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">TMT CPNS</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tmtpang">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Golongan CPNS</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 v-if="!pendingRef" ref="refPangkat" id="pangkats" :options="pangkats" keyList="id" namaList="nama" v-model="dataCpns.kgolru" />
						</div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Tanggal Melaksanakan Tugas</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tmtlgas">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Masa Kerja CPNS</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-[20%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataCpns.mskerja" placeholder="thn">
								<input type="text" class="py-2 px-3 block w-[20%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataCpns.blnkerja" placeholder="bln">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Nomor STTPP</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataCpns.nsttpp">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tsttpp">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Surat Kesehatan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataCpns.srtsehatno">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="srtsehattgl">
							</div>
						</div>
						<!-- End Col -->

					</div>
					<div class="mt-5 grid sm:grid-cols-12 gap-x-2">
						<div class="sm:col-span-7 sm:col-start-4">
							<div class="sm:flex gap-2">
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
	</LayoutDataInduk>
</template>