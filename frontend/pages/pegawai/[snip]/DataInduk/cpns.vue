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
	nlgas:'',
	tmtlgas: '',
	mskerja: 0,
	blnkerja: 0,
	srtsehatno: '',
	srtsehattgl: '',
	nsttpp: '',
	tsttpp: '',
	filename: '',
	filename_spmt: '',
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
		//var result = await $fetch('/api/gets/cpns/'+nip);

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

async function simpanData() {
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

		console.log(data.value)
		console.log(error.value.data.data)

		if(error.value){
			const {title, description} = error.value.data.data
			toast.add({
				id: 'error_put_cpns',
				icon: "i-heroicons-x-circle",
				title: title,
				description: description,
				timeout: 6000,
				color: 'red',
			}) 	
		}else{
			toast.add({
				id: 'success_put_cpns',
				icon: "i-heroicons-check-circle",
				title: "Update Data Berhasil",
				description: "Data telah di Update",
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
			const {title, description} = error.value.data.data
			toast.add({
				id: 'error_post_cpns',
				icon: "i-heroicons-x-circle",
				title: title,
				description: description,
				timeout: 6000,
				color: 'red',
			}) 	
		}else{
			toast.add({
				id: 'success_post_cpns',
				icon: "i-heroicons-check-circle",
				description: "Data berhasil di Ditambahkan",
				timeout: 6000
			}) 	
		}
	}
  	showSpinner.value = false

  	/*if(lvl == 2){
  		//save cpns bkn
	    let nip = $decodeBase64(snip)
	    var {data, error} = await useFetch('/api/puts/singkronisasi_bkn/putsDataBKNCpnsPns/'+nip, {
	        method: 'PUT',
	    })

	    console.log(data.value)
	    console.log(error.value)

	    if(error.value){
	        toast.add({
	            id: 'error_put_singkronisasi_ak',
	            description: error,
	            timeout: 6000,
	            color: 'red',
	        })  
	    }else{
	        toast.add({
	            id: 'success_post_singkronisasi_ak',
	            description: data.value.code == 1 ? "Data BKN Update Success |" : "Data BKN Update Failed |"  +data.value.data,
		    	timeout: 6000,
		    	color:  data.value.code == 1 ? 'green' : 'red',
	        })  
	    }
  		//end save cpns bkn
  	}*/

  	refreshNuxtData(["getDataRef", "getDataCPNS"])
}

const selUploadDoc = ref(null)

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

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		formData.append("tmtpang", tmtpang.value);
		formData.append("kgolru", dataCpns.value.kgolru);
		formData.append("knpang", 211);

		formData.append("file", fileBlob.value[0]);
		formData.append("updateField", selUploadDoc.value == 'spmt' ? "filename_spmt" : "filename");
		formData.append("filename", nip+"_"+selUploadDoc.value);
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
		    	description: `Dokumen ${selUploadDoc.value} berhasil di Upload`,
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
							<label class="inline-block text-sm text-gray-800 mt-2.5">Surat Pernyataan Melaksanakan Tugas (SPMT)</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataCpns.nlgas">
								<input type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tmtlgas">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">STTPP (Prajab/Latsar)</label>
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
							<label class="inline-block text-sm text-gray-800 mt-2.5">Surat Pemeriksaan Kesehatan (Dokter)</label>
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

	  		<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2 mt-6" v-if="method == 'Update'">
				<div class="mb-8">
					<h2 class="text-xl font-bold text-blue-600">Dokumen CPNS</h2>
				</div>
				<div class="grid sm:grid-cols-4 gap-2 gap-2.5">
					<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc('CPNS','cpns')">
		  				SK CPNS
					</button>
					<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc('SPMT','spmt')">
		  				SPMT
					</button>
				</div>
			</div>
		</div>
		<!-- End Card Section -->
	</LayoutDataInduk>
</template>