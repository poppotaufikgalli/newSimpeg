<script setup>
const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const toast = useToast()
const {snip} = route.params
const showSpinner = ref(false)

onMounted(() => {
	refreshNuxtData(["getdataPendum"])
})

const dataSimpeg = ref([])
const { pending, data, refresh} = await useLazyAsyncData('getdataPendum', async() => {
	console.log("CariData pendum")

	let nip = $decodeBase64(snip)
	var body = JSON.stringify({
		nip: nip,
	});
	var result =  await $fetch('/api/posts/riwayat_pendum', {
		method: 'POST',
		body: body,
	});

	dataSimpeg.value = result

	return result
})

const { pending : pending_bkn, data : result_bkn, error} = await useLazyAsyncData('getdataPendumBkn', async() => {
	console.log("CariData pendum Bkn")
	let nip = $decodeBase64(snip)

	var result = await $fetch(`/api/gets/singkronisasi_bkn/getDataBKN/rw-pendidikan/${nip}`)
	return result.data
});

const editData = (item) => {
	let sktpu = item ? $encodeBase64(item.ktpu) : $encodeBase64(null)
	let skjur = item ? $encodeBase64(item.kjur) : $encodeBase64(null)
	navigateTo({ path: `/pegawai/${snip}/RiwayatPendidikan/pendum/${sktpu}-${skjur}` })
}

const hapusData = async(item) => {
	showSpinner.value = true
	var body = JSON.stringify({
		nip: item.nip,
		ktpu: item.ktpu,
		kjur: item.kjur,
	});

	var {data, error} = await useFetch('/api/delete/riwayat_pendum', {
		method: 'DELETE',
		body: body,
	});

	if(error.value){
		const {title, message, color} = error.value.data.data
		toast.add({
	    	id: 'error_delete_riwayat_pendum',
	    	icon: "i-heroicons-x-circle",
	    	title: title,
	    	description: message,
	    	description: error.value.data.data,
	    	timeout: 6000,
	    	color: 'red',
	  	}) 	
	}else{
		toast.add({
	    	id: 'success_delete_riwayat_pendum',
	    	icon: "i-heroicons-check-circle",
	    	title: "Data Riwayat Pendidikan Umum",
	    	description: "Data Pendidikan Umum berhasil di Hapus",
	    	timeout: 6000
	  	}) 	
	}
	showSpinner.value = false
	refreshNuxtData(["getdataPendum"])
}

function checkPath(path){
	if(path != null){
		//if(path["872"] != undefined){
			return true
		//}
		//console.log(path["872"])
		
	}
	return false
}

function checkIdSync(item){
	//console.log(item.id)
	/*var idSync = item.id

	return _find(dataSimpeg.value, function(o){
		if (o.idSync == item.id) {
			return true
		}
		return false
	})*/

	return true
}

/*function getDataBKN2(item){
	let sktpu = $encodeBase64(null)
	let skjur = $encodeBase64(null)
	let sdataBkn = $encodeBase64(JSON.stringify(item))
	navigateTo({ path: `/pegawai/${snip}/RiwayatPendidikan/pendum/${sktpu}-${skjur}`, query: {
		sdataBkn : sdataBkn
	} })
}*/

async function getDataBKN(item){
	//let tmulaiBkn = item.tahunMulaiPenailan+"-"+item.bulanMulaiPenailan+"-1"
	//tmulaiBkn = dayjs(tmulaiBkn).format("YYYY-MM-DD").toString()

	var dd = _find(dataSimpeg.value, function(o){
		let tmulaiSimpeg = dayjs(o.tmulai).format("YYYY-MM-DD").toString()
		return o.idSync == item.id && o.idSync == item.id
	})

	let nip = $decodeBase64(snip)

	//let tselesaiBkn = item.tahunSelesaiPenailan+"-"+item.bulanSelesaiPenailan+"-1"
	//tselesaiBkn = dayjs(tselesaiBkn).format("YYYY-MM-DD").toString()

	//let tanggalSkBkn = dayjs(item.tanggalSk).format("YYYY-MM-DD").toString()

	var compacted = {};
	compacted.idSync = item.id
	compacted.id = 0
	compacted.nip = nip


	//{ "gelarBelakang": null, "gelarDepan": null, "id": "B583300866735D21E040640A02023F06", "idPns": "A8ACA7CD3C7F3912E040640A040269BB", "isPendidikanPertama": null, "namaSekolah": null, "nipBaru": "198602162008031001", "nipLama": "P20006592", "nomorIjasah": null, "path": null, "pendidikanId": "A5EB03E2149FF6A0E040640A040252AD", "pendidikanNama": "D-III KOMPUTER", "tahunLulus": "2006", "tglLulus": "", "tkPendidikanId": "30", "tkPendidikanNama": "Diploma III/Sarjana Muda" }

	compacted.ktpu = parseInt(item.tkPendidikanId, 10).toString()
	compacted.kjur = item.pendidikanId
	compacted.kjur_bkn = item.pendidikanId
	compacted.nsek = item.namaSekolah
	compacted.nsttb = item.nomorIjasah
	compacted.npdum = item.pendidikanNama
	console.log(item)

	/*compacted.tmulai = dayjs(dayjs(tmulaiBkn, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]")
	compacted.tselesai = dayjs(dayjs(tmulaiBkn, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]")
	compacted.utama = item.kreditUtamaBaru *1
	compacted.tambahan = item.kreditPenunjangBaru *1
	compacted.total = item.kreditBaruTotal *1
	compacted.nsk = item.nomorSk
	compacted.tsk = dayjs(dayjs(tanggalSkBkn, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]")*/
	
	if(dd){
		console.log("ketemu")
		var body = {
			"refdata" : {
				nip: nip,
				ktpu: dd.ktpu,
				kjur: dd.kjur,
			},
			"data" : compacted
		}

		console.log(body)

		var {data, error} = await useFetch('/api/puts/riwayat_pendum', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			toast.add({
		    	id: 'error_put_pendum',
		    	icon: "i-heroicons-x-circle",
		    	title: "Tambah Pendidikan Umum dari Data BKN",
		    	description: error.value.data.data,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_put_pendum',
		    	icon: "i-heroicons-check-circle",
		    	title: "Tambah Pendidikan Umum dari Data BKN",
		    	description: "Data berhasil di Update",
		    	timeout: 6000
		  	})

		  	editData(data.value)

		  	//console.log(data.value)
		}
	}else{
		console.log("tidak ketemu")
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_pendum/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			toast.add({
		    	id: 'error_post_pendum',
		    	icon: "i-heroicons-x-circle",
		    	title: "Tambah data Pendidikan Umum dari Data BKN",
		    	description: error.value.data.data,
		    	timeout: 6000,
		    	color: 'red',
		  	}) 	
		}else{
			toast.add({
		    	id: 'success_post_pendum',
		    	icon: "i-heroicons-check-circle",
		    	title: "Tambah data Pendidikan Umum dari Data BKN",
		    	description: "Data berhasil di Ditambahkan",
		    	timeout: 6000
		  	}) 	
			editData(data.value)
		  	//console.log(data.value)
		}
	}

	refreshNuxtData(["getdataPendum"])
}
</script>
<template>
	<AppLoadingSpinner :show="showSpinner" />
    <LayoutRiwayatPendidikan>
        <div class="mx-auto">
            <!-- Card -->
            <div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
                <div class="flex justify-end mb-2">
                    <button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="editData()">Tambah</button>
                </div>
                <div class="mb-8">
                    <div class="overflow-x-auto border">
                        <div class="min-w-full inline-block align-middle">
                            <!-- Table -->
                            <table class="min-w-full divide-y divide-gray-200">
                                <thead class="bg-gray-50 h-12">
                                    <tr>
                                        <th scope="col" class="px-3 text-center">
                                            <span class="text-xs font-semibold uppercase text-gray-800">
                                                No
                                            </span>
                                        </th>
                                        <th scope="col" class="px-3 text-center">
                                            <span class="text-xs font-semibold uppercase text-gray-800">
                                                Tingkat Pendidikan
                                            </span>
                                        </th>
                                        <th scope="col" class="px-3 text-center">
                                            <span class="text-xs font-semibold uppercase text-gray-800">
                                                Jurusan
                                            </span>
                                        </th>
                                        <th scope="col" class="px-3 text-center">
                                            <span class="text-xs font-semibold uppercase text-gray-800">
                                                Nama Sekolah
                                            </span>
                                        </th>
                                        <th scope="col" class="px-3 text-center">
                                            <span class="text-xs font-semibold uppercase text-gray-800">
                                                Tempat
                                            </span>
                                        </th>
                                        <th scope="col" class="px-3 text-center">
                                            <span class="text-xs font-semibold uppercase text-gray-800">
                                                Akhir
                                            </span>
                                        </th>
                                        <th scope="col" class="px-3 text-center"></th>
                                    </tr>
                                </thead>
                                <tbody class="divide-y divide-gray-200">
                                    <template v-if="pending">
                                        <tr>
                                            <td colspan="7" class="size-px p-1" style="vertical-align: top; min-width: 100px;">
                                                <div class="flex justify-center animate-pulse">
                                                    <span class="text-sm italic">Loading</span>
                                                </div>
                                            </td>
                                        </tr>
                                    </template>
                                    <template v-else-if="data" v-for="(item, idx) in data" :key="idx">
                                        <tr class="odd:bg-white even:bg-gray-100">
                                            <td class="size-px px-1" style="vertical-align: top;">
                                                <button type="button" class="block" data-hs-overlay="#hs-ai-invoice-modal">
                                                    <span class="flex justify-center text-xs px-3 py-0">{{idx+1}}</span>
                                                </button>
                                            </td>
                                            <td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
                                                <span class="flex justify-start text-xs py-0">{{item.TingkatPendidikan?.nama}}</span>
                                            </td>
                                            <td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
                                                <span class="flex justify-start text-xs py-0">{{item.Pendidikan?.nama}}</span>
                                            </td>
                                            <td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
                                                <span class="flex justify-start text-xs py-0">{{item.nsek}}</span>
                                            </td>
                                            <td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
                                                <span class="flex justify-start text-xs py-0">{{item.tempat}}</span>
                                            </td>
                                            <td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
                                                <span class="flex justify-center text-xs py-0">{{item.akhir == 1 ? 'Ya' : ''}}</span>
                                            </td>
                                            <td class="size-px px-1" style="vertical-align: top;">
                                                <div class="sm:flex gap-x-1">
                                                    <button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="editData(item)">
                                                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="flex-shrink-0 size-3">
                                                            <path d="M13.488 2.513a1.75 1.75 0 0 0-2.475 0L6.75 6.774a2.75 2.75 0 0 0-.596.892l-.848 2.047a.75.75 0 0 0 .98.98l2.047-.848a2.75 2.75 0 0 0 .892-.596l4.261-4.262a1.75 1.75 0 0 0 0-2.474Z" />
                                                            <path d="M4.75 3.5c-.69 0-1.25.56-1.25 1.25v6.5c0 .69.56 1.25 1.25 1.25h6.5c.69 0 1.25-.56 1.25-1.25V9A.75.75 0 0 1 14 9v2.25A2.75 2.75 0 0 1 11.25 14h-6.5A2.75 2.75 0 0 1 2 11.25v-6.5A2.75 2.75 0 0 1 4.75 2H7a.75.75 0 0 1 0 1.5H4.75Z" />
                                                        </svg>
                                                    </button>
                                                    <button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-red-600 text-white hover:bg-red-700 disabled:opacity-50 disabled:pointer-events-none" @click="hapusData(item)">
                                                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="flex-shrink-0 size-3">
                                                            <path fill-rule="evenodd" d="M5 3.25V4H2.75a.75.75 0 0 0 0 1.5h.3l.815 8.15A1.5 1.5 0 0 0 5.357 15h5.285a1.5 1.5 0 0 0 1.493-1.35l.815-8.15h.3a.75.75 0 0 0 0-1.5H11v-.75A2.25 2.25 0 0 0 8.75 1h-1.5A2.25 2.25 0 0 0 5 3.25Zm2.25-.75a.75.75 0 0 0-.75.75V4h3v-.75a.75.75 0 0 0-.75-.75h-1.5ZM6.05 6a.75.75 0 0 1 .787.713l.275 5.5a.75.75 0 0 1-1.498.075l-.275-5.5A.75.75 0 0 1 6.05 6Zm3.9 0a.75.75 0 0 1 .712.787l-.275 5.5a.75.75 0 0 1-1.498-.075l.275-5.5a.75.75 0 0 1 .786-.711Z" clip-rule="evenodd" />
                                                        </svg>
                                                    </button>
                                                </div>
                                            </td>
                                        </tr>
                                    </template>
                                </tbody>
                            </table>
                            <!-- End Table -->
                        </div>
                    </div>
                </div>
                <div class="mb-8">
                    <div class="overflow-x-auto border">
                        <h2 class="font-bold text-teal-600">Data BKN</h2>
                        <div class="min-w-full inline-block align-middle">
                            <!-- Table -->
                            <table class="min-w-full divide-y divide-gray-200">
                                <thead class="bg-gray-50 h-12">
                                    <tr>
                                        <th scope="col" class="px-3 text-center">
                                            <span class="text-xs font-semibold uppercase text-gray-800">
                                                No
                                            </span>
                                        </th>
                                        <th scope="col" class="px-3 text-center"></th>
                                        <th scope="col" class="px-3 text-center">
                                            <span class="text-xs font-semibold uppercase text-gray-800">
                                                Tingkat Pendidikan
                                            </span>
                                        </th>
                                        <th scope="col" class="px-3 text-center">
                                            <span class="text-xs font-semibold uppercase text-gray-800">
                                                Jurusan
                                            </span>
                                        </th>
                                        <th scope="col" class="px-3 text-center">
                                            <span class="text-xs font-semibold uppercase text-gray-800">
                                                Nama Sekolah
                                            </span>
                                        </th>
                                        <th scope="col" class="px-3 text-center">
                                            <span class="text-xs font-semibold uppercase text-gray-800">
                                                Nomor Ijazah
                                            </span>
                                        </th>
                                        <th scope="col" class="px-3 text-center">
                                            <span class="text-xs font-semibold uppercase text-gray-800">
                                                Pendidikan Pertama?
                                            </span>
                                        </th>
                                        <th scope="col" class="px-3 text-center"></th>
                                    </tr>
                                </thead>
                                <tbody class="divide-y divide-gray-200">
                                    <template v-if="pending_bkn">
                                        <tr>
                                            <td colspan="8" align="center">
                                                <div class="flex justify-center animate-pulse">
                                                    <span class="text-sm italic">Loading</span>
                                                </div>
                                            </td>
                                        </tr>
                                    </template>
                                    <template v-else-if="error">
                                        <tr>
                                            <td colspan="8" align="center">
                                                <span class="flex justify-center text-xs py-1 text-red-500">Get Data BKN Gagal :<b> {{error.data.data.message || error.data.data.errors}}</b></span>
                                            </td>
                                        </tr>
                                    </template>
                                    <template v-else-if="result_bkn" v-for="(item, idx) in result_bkn" :key="idx">
                                        <tr class="odd:bg-white even:bg-gray-100">
                                            <td class="size-px px-1" style="vertical-align: top;">
                                                <button type="button" class="block" data-hs-overlay="#hs-ai-invoice-modal">
                                                    <span class="flex justify-center text-xs px-3 py-0">{{idx+1}}</span>
                                                </button>
                                            </td>
                                            <td class="size-px px-1" style="vertical-align: top; width: 5%">
                                                <div class="sm:flex justify-center">
                                                    <template v-if="checkPath(item.path)">
                                                        <button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="downloadDoc(item)">
                                                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4">
                                                                <path fill-rule="evenodd" d="M2 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4Zm10.5 5.707a.5.5 0 0 0-.146-.353l-1-1a.5.5 0 0 0-.708 0L9.354 9.646a.5.5 0 0 1-.708 0L6.354 7.354a.5.5 0 0 0-.708 0l-2 2a.5.5 0 0 0-.146.353V12a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V9.707ZM12 5a1 1 0 1 1-2 0 1 1 0 0 1 2 0Z" clip-rule="evenodd" />
                                                            </svg>
                                                        </button>
                                                    </template>
                                                </div>
                                            </td>
                                            <td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
                                                <span class="flex justify-start text-xs py-0">{{item.tkPendidikanNama}}</span>
                                            </td>
                                            <td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
                                                <span class="flex justify-start text-xs py-0">{{item.pendidikanNama}}</span>
                                            </td>
                                            <td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
                                                <span class="flex justify-start text-xs py-0">{{item.namaSekolah}}</span>
                                            </td>
                                            <td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
                                                <span class="flex justify-start text-xs py-0">{{item.nomorIjasah}}</span>
                                            </td>
                                            <td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
                                                <span class="flex justify-center text-xs py-0">{{item.isPendidikanPertama == 1 ? 'Ya' : ''}}</span>
                                            </td>
                                            <td class="size-px px-1" style="vertical-align: top;">
                                                <template v-if="checkIdSync(item)">
                                                    <button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-green-600 text-white hover:bg-green-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="getDataBKN(item)">
                                                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="flex-shrink-0 size-3">
                                                            <path fill-rule="evenodd" d="M4.5 13a3.5 3.5 0 0 1-1.41-6.705A3.5 3.5 0 0 1 9.72 4.124a2.5 2.5 0 0 1 3.197 3.018A3.001 3.001 0 0 1 12 13H4.5Zm.72-5.03a.75.75 0 0 0 1.06 1.06l.97-.97v2.69a.75.75 0 0 0 1.5 0V8.06l.97.97a.75.75 0 1 0 1.06-1.06L8.53 5.72a.75.75 0 0 0-1.06 0L5.22 7.97Z" clip-rule="evenodd" />
                                                        </svg>
                                                    </button>
                                                </template>
                                            </td>
                                        </tr>
                                    </template>
                                </tbody>
                            </table>
                            <!-- End Table -->
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </LayoutRiwayatPendidikan>
</template>