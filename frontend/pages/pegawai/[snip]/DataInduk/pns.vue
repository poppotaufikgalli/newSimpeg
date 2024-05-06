<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)

const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const {snip} = route.params

const dataRefPns = ref({
	nip: '',
	knpang: 212,
	tmtpang: '',
	kgolru: '',
})

const dataPns = ref({
	nip: '',
	ptetap: '',
	nskpang:'',
	tskpang: '',
	tmtpang: '',
	kgolru: '',
	kjpns: '',
	filename: '',
})

const pejabats = ref([])
const pangkats = ref([])
const method = ref('Simpan')

const refPejabat = ref(null)
const refPangkat = ref(null)

const { pending, data, refresh} = await useAsyncData('getDataPNS', async() => {
	//console.log("CariData Identitas")

	if(snip){
		let nip = $decodeBase64(snip)
		var result = await $fetch('/api/gets/riwayat_pangkat/'+nip+'/pns');

		if(result.nip == nip){
			method.value = "Update"

			dataPns.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataPns.value), key);
			})

			dataRefPns.value = _pickBy(result, function(val, key) {
				return _includes(_keys(dataRefPns.value), key);
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

const ssumpah = computed({
	get(){
		return dataPns.value.kjpns == 0 ? 'Belum' : 'Sudah'
	},
})

onMounted(() => {
	refreshNuxtData(["getDataPNS", "getDataRef"])
	modalUploadDoc.$onAction(callback, false)
})

const tskpang = computed({
	get(){
		return dayjs(dataPns.value.tskpang).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataPns.value.tskpang = val
	}
})

const tmtpang = computed({
	get(){
		return dayjs(dataPns.value.tmtpang).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataPns.value.tmtpang = val
	}
})

const kjpns = computed({
	get(){
		return dataPns.value.kjpns == 0 ? false : true
	},
	set(val){
		dataPns.value.kjpns = val
	}
})

const showSpinner = ref(false)
const toast = useToast()

async function simpanData() {
	showSpinner.value = true

	//var compacted = _pickBy(dataCpns.value);
	var compacted = dataPns.value;
	compacted.tskpang = compacted.tskpang ? dayjs(dayjs(compacted.tskpang, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tmtpang = compacted.tmtpang ? dayjs(dayjs(compacted.tmtpang, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.kgolru = compacted.kgolru *1
	compacted.kjpns = compacted.kjpns == true ? 1 : 0
	compacted.ptetap = compacted.ptetap *1

	compacted.knpang = 212 //knpang pns = 212

	var dr = dataRefPns.value
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

		if(error.value){
			toast.add({
		    	id: 'error_put_pns',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_pns',
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	}) 	
		}
		//console.log(error.value.data.data)
		
	}else{
		compacted.nip = $decodeBase64(snip)
		var body = JSON.stringify(compacted)
		var {data, error} = await useFetch('/api/posts/riwayat_pangkat/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_pns',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_pns',
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
		}
	}
  	showSpinner.value = false
  	refreshNuxtData(["getDataRef", "getDataPNS"])
}

const doUploadDoc = async() => {
	let nip = $decodeBase64(snip)
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataPns.value.filename}`)
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen PNS`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen PNS`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)

		formData.append("tmtpang", tmtpang.value);
		formData.append("kgolru", dataPns.value.kgolru);
		formData.append("knpang", 212);

		formData.append("file", fileBlob.value[0]);
		formData.append("filename", nip+"_pns");
		formData.append("path", 'dokumen/'+nip);

		var {data, error} = await useFetch(`/api/uploads/upload/pangkat/${nip}`, {
			method: 'POST',
			body: formData,
		});

		if(error.value){
			toast.add({
		    	id: 'error_post_upload_pns',
		    	description: error.value.data.data,
		    	timeout: 6000
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_upload_pns',
		    	description: `Dokumen PNS berhasil di Upload`,
		    	timeout: 6000
		  	}) 	
		}

		refreshNuxtData(["getDataRef", "getDataPNS"])

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
						<h2 class="text-xl font-bold text-blue-600">PNS</h2>
						<div class="flex gap-x-2 w-[60%]">
							<input type="text" name="hs-input-with-add-on-url" id="hs-input-with-add-on-url" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP" :value="dataPns.nip">
							<input type="text" name="hs-input-with-add-on-url" id="hs-input-with-add-on-url" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="Nama">
						</div>
					</div>
				</div>
				<form>
					<!-- Grid -->
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Pejabat Penetap</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refPejabat" id="pejabats" :options="pejabats" keyList="kpej" namaList="npej" v-model="dataPns.ptetap"/>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">SK PNS</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataPns.nskpang">
								<input id="af-account-full-name" type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tskpang">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">TMT PNS</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input id="af-account-full-name" type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tmtpang">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Pangkat/Golongan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refPangkat" id="pangkats" :options="pangkats" keyList="id" namaList="pangkat_gol" v-model="dataPns.kgolru"/>
						</div>

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Sudah Sumpah PNS</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<div class="flex mt-2">
								  	<input type="checkbox" class="shrink-0 mt-0.5 border-gray-200 rounded text-blue-600 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none" id="hs-default-checkbox" v-model="kjpns">
								  	<label for="hs-default-checkbox" class="text-sm text-gray-500 ms-3">{{ssumpah}}</label>
								</div>
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
	</LayoutDataInduk>
</template>