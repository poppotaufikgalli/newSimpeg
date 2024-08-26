<script setup>
const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
const toast = useToast()
const {snip} = route.params
const showSpinner = ref(false)
import customParseFormat from "dayjs/plugin/customParseFormat";
dayjs.extend(customParseFormat)

onMounted(() => {
	//refreshNuxtData(["getData"])
})

const result_simpeg = ref([])
const result_bkn = ref([])

async function getDataRiwJabatanBKN() {
    showSpinner.value = true
	console.log("CariData")
	let nip = $decodeBase64(snip)
	var body = JSON.stringify({
		nip: nip,
	});

	result_simpeg.value = await $fetch('/api/posts/riwayat_jabatan', {
		method: 'POST',
		body: body,
	});

	var bkn = await $fetch('/api/gets/singkronisasi_bkn/getDataBKN/rw-jabatan/'+nip)
	result_bkn.value = bkn.data

    //result_bkn.value = result.value

    if(result_simpeg.value){
        for (var i = 0; i < result_simpeg.value.length; i++) {
            var data_bkn = _find(result_bkn.value, {
                tmtJabatan: (dayjs(dayjs(result_simpeg.value[i].tmtjab, "YYYY-MM-DD")).format("DD-MM-YYYY")).toString(),
            })
            //console.log(data_bkn)
            result_simpeg.value[i]['data_bkn'] = JSON.stringify(data_bkn)
        }
    }

    if(result_bkn.value){
        for (var i = 0; i < result_bkn.value.length; i++) {
            var data_simpeg = _find(result_simpeg.value, {
                tmtjab: dayjs(dayjs(result_bkn.value[i].tmtJabatan, 'DD-MM-YYYY')).format("YYYY-MM-DD[T]HH:mm:ss+07:00"),
            })
            //console.log(data_simpeg)
            result_bkn.value[i]['data_simpeg'] = JSON.stringify(data_simpeg)
        }
    }

    showSpinner.value = false
}

/*const editData = (item) => {
	let stmtjab = item ? $encodeBase64(item.tmtjab) : $encodeBase64(null)
	navigateTo({ path: `/pegawai/${snip}/RiwayatPegawai/jabatan/${stmtjab}` })
}

const hapusData = async(item) => {
	var body = JSON.stringify({
		nip: item.nip,
		tmtjab: item.tmtjab,
	});

	var tmtjab = dayjs(item.tmtjab).format("DD-MM-YYYY").toString()

	var {data, error} = await useFetch('/api/delete/riwayat_jabatan', {
		method: 'DELETE',
		body: body,
	});

	if(error.value){
		toast.add({
	    	id: 'error_delete_riwayat_jabatan',
	    	description: error.value.data.data,
	    	timeout: 6000
	  	}) 	
	}else{
		toast.add({
	    	id: 'success_delete_riwayat_jabatan',
	    	description: `Data jabatan ${item.njab} / TMT ${tmtjab} berhasil di Hapus`,
	    	timeout: 6000
	  	}) 	
	}

	//refreshNuxtData(["getData"])
}*/

function chekJenisJab(item){
    if(item.jenisJabatan == "1") return item.namaJabatan + " PADA " + item.unorNama + " " + item.unorIndukNama
    else if(item.jenisJabatan == "2") return item.jabatanFungsionalNama + " PADA " + item.unorNama + " " + item.unorIndukNama
    else if(item.jenisJabatan == "4") return item.jabatanFungsionalUmumNama + " PADA " + item.unorNama + " " + item.unorIndukNama
    else return item.jabatanFungsionalUmumNama + " PADA " + item.unorNama + " " + item.unorIndukNama
}

const refDataBkn = ref({
    "eselonId": "keselon",
    "id": "idSync",
    "instansiKerjaId": "id_instansi",
    "instansiKerjaNama": "nama_instansi",
    //"jabatanFungsionalId": "",
    //"jabatanFungsionalNama": "",
    //"jabatanFungsionalUmumId": "8ae482893e8315e6013e9bc3760f0af7",
    //"jabatanFungsionalUmumNama": "PELAKSANA",
    "jenisJabatan": "jnsjab",
    //"jenisMutasiId": "",
    //"jenisPenugasanId": "",
    //"namaJabatan": null,
    //"namaUnor": "BADAN KEPEGAWAIAN DAERAH KAB. BINTAN",
    "nipBaru": "nip",
    //"nipLama": "P20006592",
    "nomorSk": "nskjabat",
    //"path": null,
    //"satuanKerjaId": "A5EB03E24320F6A0E040640A040252AD",
    //"satuanKerjaNama": "Pemerintah Kab. Bintan",
    "subJabatanId": "sub_jab",
    "tanggalSk": "tskjabat",
    "tmtJabatan": "tmtjab",
    "tmtMutasi": "tmt_mutasi",
    "tmtPelantikan": "tlantik",
    "unorId": "id_sub_opd",
    //"unorIndukId": "",
    //"unorIndukNama": "",
    //"unorNama": "BADAN KEPEGAWAIAN DAERAH "
})

async function simpanData(item){
    //showSpinner.value = true
    let nip = $decodeBase64(snip)
    var nMethod = item.data_simpeg ? 'Update' : 'Simpan'; 

    //= _pickBy(item);
    var aitem = _pickBy(item, function(val, key) {
        return _includes(_keys(refDataBkn.value), key);
    })

    var compacted = _mapKeys(aitem, function(val, key){
        return _find(refDataBkn.value, function(val1, key1){
            return key == key1
        })
    })
    console.log(item)

    compacted.tmtjab = compacted.tmtjab ? dayjs(dayjs(compacted.tmtjab, "DD-MM-YYYY")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
    compacted.tmt_mutasi = compacted.tmt_mutasi ? dayjs(dayjs(compacted.tmt_mutasi, "DD-MM-YYYY")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : compacted.tmtjab
    compacted.tlantik = compacted.tlantik ? dayjs(dayjs(compacted.tlantik, "DD-MM-YYYY")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
    compacted.tskjabat = compacted.tskjabat ? dayjs(dayjs(compacted.tskjabat, "DD-MM-YYYY")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
    compacted.nunker = item.unorNama + " " + item.unorIndukNama


    if(compacted.jnsjab == "1"){
        compacted.jns_tugas_mutasi = item.jenisPenugasanId
    }else{
        compacted.jns_tugas_mutasi = item.jenisMutasiId

        if(compacted.jnsjab == "2"){
            compacted.kjab_bkn = item.jabatanFungsionalId
            compacted.njab0 = item.jabatanFungsionalNama.toUpperCase(); 
        }else{
            compacted.kjab_bkn = item.jabatanFungsionalUmumId
            compacted.njab0 = item.jabatanFungsionalUmumNama.toUpperCase(); 
        }

        compacted.kjab = compacted.kjab_bkn
        compacted.njab = compacted.njab0 + " PADA " + compacted.nunker
        compacted.keselon = 99
    }

    var result = await $fetch("/api/posts/opd", {
        method: 'POST',
        body: JSON.stringify({
            id_unor_bkn: item.unorId
        })
    })
    
    if(result && result[0]){
        compacted.id_sub_opd = result[0].id
        compacted.id_opd = (compacted.id_sub_opd).substring(0,3)
    }

    console.log(compacted)
    
    var dr = {
        nip: nip,
        tmtjab: compacted.tmtjab,
        tmt_mutasi: compacted.tmt_mutasi,
    }

    if(nMethod == 'Update'){
        var body = {
            "refdata" : dr,
            "data" : compacted
        }

        var {data, error} = await useFetch('/api/puts/riwayat_jabatan', {
            method: 'PUT',
            body: JSON.stringify(body),
        })

        if(error.value){
            toast.add({
                id: 'error_put_jabatan',
                description: error.value.data.data,
                timeout: 6000
            })  
        }else{
            toast.add({
                id: 'success_put_jabatan',
                description: "Data berhasil di Update",
                timeout: 6000
            })  
        }
    }else{
        compacted.nip = $decodeBase64(snip)
        //compacted.akhir = 1
        var body = JSON.stringify(compacted)

        var {data, error} = await useFetch('/api/posts/riwayat_jabatan/new', {
            method: 'POST',
            body: body,
        })

        if(error.value){
            toast.add({
                id: 'error_post_jabatan',
                description: error.value.data.data,
                timeout: 6000
            })  
        }else{
            toast.add({
                id: 'success_post_jabatan',
                description: "Data berhasil di Ditambahkan",
                timeout: 6000
            })  
        }
    }
    showSpinner.value = false
    getDataRiwJabatanBKN()
}
</script>
<template>
	<AppLoadingSpinner :show="showSpinner" />
	<LayoutSingkronisasi>
		<TokenInfo />
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="flex justify-end mb-2">
					<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="getDataRiwJabatanBKN()">Get Data BKN</button>
				</div>
				<div class="mb-8">
					<h2 class="font-bold text-teal-600">Data Simpeg</h2>
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
												Nama
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												TMT Jabatan
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Nomor SK
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Tanggal
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Eselon
											</span>
										</th>
										<th scope="col" class="px-3 text-center"></th>
									</tr>
								</thead>
								<tbody class="divide-y divide-gray-200">
									<template v-if="result_simpeg" v-for="(item, idx) in result_simpeg" :key="idx">
										<tr class="odd:bg-white even:bg-gray-100">
											<td class="size-px px-1" style="vertical-align: top;">
												<button type="button" class="block" data-hs-overlay="#hs-ai-invoice-modal">
													<span class="flex justify-center text-xs px-3 py-0">{{idx+1}}</span>
												</button>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 200px;">
												<span class="flex justify-start text-xs py-0">{{item.njab}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 80px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.tmtjab ? $dayjs(item.tmtjab).format("DD-MM-YYYY").toString() : '-'}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 80px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.nskjabat}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 80px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.tskjabat ? $dayjs(item.tskjabat).format("DD-MM-YYYY").toString() : '-'}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 80px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.eselon?.nama}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top;">
											     &nbsp;	
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
					<h2 class="font-bold text-teal-600">Data BKN</h2>
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
												Nama
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												TMT Jabatan
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Nomor SK
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Tanggal
											</span>
										</th>
										<th scope="col" class="px-3 text-center">
											<span class="text-xs font-semibold uppercase text-gray-800">
												Eselon
											</span>
										</th>
										<th scope="col" class="px-3 text-center"></th>
									</tr>
								</thead>
								<tbody class="divide-y divide-gray-200">
									<template v-if="result_bkn" v-for="(item, idx) in result_bkn" :key="idx">
										<tr class="odd:bg-white even:bg-gray-100">
											<td class="size-px px-1" style="vertical-align: top;">
												<button type="button" class="block" data-hs-overlay="#hs-ai-invoice-modal">
													<span class="flex justify-center text-xs px-3 py-0">{{idx+1}}</span>
												</button>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 200px;">
												<span class="flex justify-start text-xs py-0">{{chekJenisJab(item)}}</span>
                                                <span class="flex justify-start text-xs py-0 italic">{{item.id}}</span>
                                                <span class="flex justify-start text-xs py-0 italic">{{item.tmtMutasi}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 80px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.tmtJabatan}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 80px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.nomorSk}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 80px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.tanggalSk}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 80px;">
												<span class="flex justify-center text-center text-xs py-0">{{item.eselon}}</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top;">
												<div class="sm:flex gap-x-1">
                                                    <template v-if="item.data_simpeg">
                                                        <button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-green-600 text-white hover:bg-green-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="simpanData(item)">
                                                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4">
                                                                <path fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14Zm3.844-8.791a.75.75 0 0 0-1.188-.918l-3.7 4.79-1.649-1.833a.75.75 0 1 0-1.114 1.004l2.25 2.5a.75.75 0 0 0 1.15-.043l4.25-5.5Z" clip-rule="evenodd" />
                                                            </svg>
                                                        </button>
                                                    </template>
                                                    <template v-else>
                                                        <button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-yellow-600 text-white hover:bg-yellow-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="simpanData(item)">
                                                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4">
                                                                <path fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14Zm.75-10.25a.75.75 0 0 0-1.5 0v4.69L6.03 8.22a.75.75 0 0 0-1.06 1.06l2.5 2.5a.75.75 0 0 0 1.06 0l2.5-2.5a.75.75 0 1 0-1.06-1.06L8.75 9.44V4.75Z" clip-rule="evenodd" />
                                                            </svg>
                                                        </button>
                                                    </template>
                                                    <template v-if="item.path">
                                                        <button type="button" class="flex flex-shrink-0 justify-center items-center gap-2 size-[24px] text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" v-on:click="doSyncDoc(item)">
                                                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4">
                                                                <path fill-rule="evenodd" d="M2 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4Zm10.5 5.707a.5.5 0 0 0-.146-.353l-1-1a.5.5 0 0 0-.708 0L9.354 9.646a.5.5 0 0 1-.708 0L6.354 7.354a.5.5 0 0 0-.708 0l-2 2a.5.5 0 0 0-.146.353V12a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V9.707ZM12 5a1 1 0 1 1-2 0 1 1 0 0 1 2 0Z" clip-rule="evenodd" />
                                                            </svg>
                                                        </button>
                                                    </template>
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
			</div>
		</div>
	</LayoutSingkronisasi>
</template>