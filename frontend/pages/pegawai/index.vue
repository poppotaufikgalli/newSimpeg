<script setup>
const { $encodeBase64 } = useNuxtApp()
const emptyNip = null
const search = ref('');
const page = ref(1);

const from = ref(0)
const perPage = ref(10)

const opds = ref([])
const lsPangkats = ref([])
const jenisjabatans = ref([])
const eselons = ref([])
const statuspegawais = ref([])
const jenispegawais = ref([])
const loading = ref(true)

const { pending: pendingRef, data, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
  const [dataOPD, dataPangkat, dataJenisJabatan, dataEselon, dataStatusPegawai, dataJenisPegawai] = await Promise.all([
    $fetch('/api/gets/opd'),
    $fetch('/api/gets/pangkat'),
    $fetch('/api/gets/jenis_jabatan'),
    $fetch('/api/gets/eselon'),
    $fetch('/api/gets/status_pegawai'),
    $fetch('/api/gets/jenis_pegawai')
  ])

  //console.log("getDataRef")

  opds.value = dataOPD
  lsPangkats.value = dataPangkat
  jenisjabatans.value = dataJenisJabatan
  eselons.value = dataEselon
  statuspegawais.value = dataStatusPegawai
  jenispegawais.value = dataJenisPegawai

  //loading.value = false
});

const { pending, data: dataPegawai, refresh: refreshDataPegawai } = await useAsyncData('getDataPegawai', () => CariData())

const searchOpd = ref(null)
const searchJJ = ref(null)
const searchEselon = ref(null)
const searchSubOpd = ref(null)
const searchNama = ref(null)
const searchNip = ref(null)
const searchStatusPegawai = ref([1,2,11])
const searchJenisPegawai = ref([10,15,20])

const refSubOpd = ref(null)
const refOpd = ref(null)
const refJenisJabatan = ref(null)
const refEselon = ref(null)
const refStatusPegawai = ref(null)
const refJenisPegawai = ref(null)

const opds_utama = computed({
	get(){
		return opds.value
	}
})

const sub_opds_utama = computed({
	get(){
		if(searchOpd.value != null){
	  	var result = _find(opds.value, function(o){
	  		return o.id == searchOpd.value
	  	})
	  	if(result){
	  		if(refSubOpd){
	  			refSubOpd.value.reinit()	
	  		}
	  	}
	  	return result.sub_opd
	  }
	  //searchSubOpd.value = null
	  return searchOpd.value
	},
})

const paginate_pegawais = computed({
	get(){
		return dataPegawai.value.slice(from.value, from.value + perPage.value) || null
	}
})

const prevPage = () => {
	if(from.value - perPage.value >= 0){
		from.value = from.value - perPage.value 
	}
}

const nextPage = () => {
	if(from.value + perPage.value <= dataPegawai.value.length){
		from.value = from.value + perPage.value
	}
}

const setStatusPegawai = (item) => {
	let id = item.kstatus;
	selStatusPegawai.value[id] = !selStatusPegawai.value[id]
}

const setJenisPegawai = (item) => {
	let id = item.id;
	selJenisPegawai.value[id] = !selJenisPegawai.value[id]
}

const ResetData = () => {
	searchOpd.value = null
	searchSubOpd.value = null
	searchNama.value = null
	searchJJ.value = null
	searchEselon.value = null
	searchNip.value = null
	searchStatusPegawai.value = [1,2,11]
	searchJenisPegawai.value = [10,15,20]

	refOpd.value.reinit()
	refSubOpd.value.reinit()
	refJenisJabatan.value.reinit()
	refEselon.value.reinit()
	refStatusPegawai.value.reinit()
	refJenisPegawai.value.reinit()

	refreshNuxtData("getDataPegawai")
}

const CariData2 = async() => {
	refreshNuxtData("getDataPegawai")
}

const CariData = async() => {
	loading.value = true

	var body = JSON.stringify({
		nip: searchNip.value ?? undefined,
		nama: searchNama.value ?? undefined,
		id_opd: searchSubOpd.value ? searchSubOpd.value : searchOpd.value,
		jnsjab: searchJJ.value ? searchJJ.value *1 : undefined,
		keselon: searchEselon.value ? searchEselon.value *1 : undefined,
		kstatus: searchStatusPegawai.value.map(Number),
		kjpeg: searchJenisPegawai.value.map(Number),
	})

	body = body == "{}" ? null : body

	var result = await $fetch('/api/pegawai', {
		method: 'POST',
		body: body,
	})

	loading.value = false

	return result
}

onMounted(() => {
	console.log("onMounted ==== pegawai")
	refreshNuxtData(["getDataPegawai", "getDataRef"])
	//console.log(searchJenisPegawai.value)
})

function nama_pangkat(kgolru){
	var n = _find(lsPangkats.value, {id: kgolru})
	return n?.nama
}

function formatNip(nip) {
	var nnip = nip.slice(0,8) +" "+ nip.slice(8,14)+" "+ nip.slice(14,15)+" "+ nip.slice(15,18)
	return nnip
}

const router = useRouter()

function EditData(nip){
	const snip = $encodeBase64(nip)
	navigateTo({ path: '/pegawai/'+snip+'/DataInduk/identitas' })
}
</script>

<template>
	<!-- Table Section -->
	<div class="">
		<!-- Card -->
		<div class="flex flex-col mb-6">
			<div class="bg-white border border-gray-200 rounded-xl shadow-sm overflow-hidden">
				<!-- Header -->
				<div class="px-6 py-4 grid gap-3 md:flex md:justify-between md:items-center">
					<div>
						<h2 class="text-xl font-semibold text-gray-800">
							Data Pegawai
						</h2>
						<p class="text-sm text-gray-600">
							Jumlah Pegawai: {{dataPegawai ? dataPegawai.length : 0}}
						</p>
					</div>

					<div>
						<div class="inline-flex gap-x-2">
							<NuxtLink class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" :to="`/pegawai/${emptyNip}/DataInduk/identitas`">
								<svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="M12 5v14"/></svg>
								Tambah
							</NuxtLink>
						</div>
					</div>
				</div>
				<!-- End Header -->

				<div class="hs-accordion-group">
				  <div class="hs-accordion bg-gray-100 border border-x-0 -mt-px" id="hs-bordered-heading-two">
    				<button class="hs-accordion-toggle hs-accordion-active:text-blue-600 inline-flex items-center gap-x-3 w-full font-semibold text-start text-gray-800 py-4 px-5 hover:text-gray-500 disabled:opacity-50 disabled:pointer-events-none" aria-controls="hs-basic-bordered-collapse-two">
      				<svg class="hs-accordion-active:block hidden size-3.5" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        				<path d="M5 12h14"></path>
      				</svg>
      				<svg class="hs-accordion-active:hidden block size-3.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor">
							  <path d="M14 2a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v2.172a2 2 0 0 0 .586 1.414l2.828 2.828A2 2 0 0 1 6 9.828v4.363a.5.5 0 0 0 .724.447l2.17-1.085A2 2 0 0 0 10 11.763V9.829a2 2 0 0 1 .586-1.414l2.828-2.828A2 2 0 0 0 14 4.172V2Z" />
							</svg>
      				Filter Lainnya
    				</button>
    				<div id="hs-basic-bordered-collapse-two" class="hs-accordion-content hidden w-full transition-[height] duration-300" aria-labelledby="hs-bordered-heading-two">
      				<div class="pb-4 px-5">
        				<div class="gap-2 text-gray-800">
									<div class="flex flex-col gap-2">
										<!--Unit Kerja -->
										<div class="grid sm:grid-cols-12 gap-1 pt-8 pb-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200">
							        <label class="inline-block sm:col-span-1 text-sm font-medium mt-2.5">
						            Unit Kerja
						          </label>
							        <!-- End Col -->
							        <div class="sm:col-span-11 px-1 my-auto">
							        	<SearchSelect2 v-if="!pendingRef" ref="refOpd" id="opd_utama" :options="opds_utama" keyList="id" namaList="nama" v-model="searchOpd" />
							        </div>
									     
											<template v-if="opds_utama">
								        <label class="inline-block sm:col-span-1 text-sm font-medium mt-2.5">
						            	Sub Unit Kerja
						          	</label>
								        <!-- End Col -->
								        <div class="sm:col-span-11 px-1 my-auto">
							        		<SearchSelect2 v-if="!pendingRef" ref="refSubOpd" id="sub_opd_utama" :options="sub_opds_utama" keyList="id" namaList="nama" v-model="searchSubOpd" />
							        	</div>
											</template>

											<label class="inline-block sm:col-span-1 text-sm font-medium mt-2.5">
					            	Jenis Jabatan
					          	</label>
							        <!-- End Col -->
							        <div class="sm:col-span-5 px-1 my-auto">
						        		<SearchSelect2 v-if="!pendingRef" ref="refJenisJabatan" id="jenisjabatans" :options="jenisjabatans" keyList="id" namaList="nama" v-model="searchJJ" />
						        	</div>

											<label class="inline-block sm:col-span-1 text-sm font-medium mt-2.5">
					            	Eselon
					          	</label>
									    <!-- End Col -->
									    <div class="sm:col-span-5 px-1 my-auto">
						        		<SearchSelect2 v-if="!pendingRef" ref="refEselon" id="eselons" :options="eselons" keyList="id" namaList="nama" v-model="searchEselon" />
						        	</div>

											<label class="inline-block sm:col-span-1 text-sm font-medium mt-2.5">
					            	Status Pegawai 
					          	</label>
									    <!-- End Col -->
									    <div class="sm:col-span-5 px-1 my-auto">
						        		<SearchSelect2 v-if="!loading" ref="refStatusPegawai" id="statuspegawais" :options="statuspegawais" keyList="kstatus" namaList="nama" v-model="searchStatusPegawai" multiple="multiple" />
						        	</div>

											<label class="inline-block sm:col-span-1 text-sm font-medium mt-2.5">
						            Jenis Pegawai
						          </label>
							        <!-- End Col -->
							        <div class="sm:col-span-5 px-1 my-auto">
						        		<SearchSelect2 v-if="!loading" ref="refJenisPegawai" id="jenispegawais" :options="jenispegawais" keyList="id" namaList="nama" v-model="searchJenisPegawai" multiple="multiple" />
						        	</div>
										</div>
									</div>
								</div>
      				</div>
    				</div>
  				</div>
				</div>

				<div class="bg-gray-50">
					<div class="p-4 gap-2 text-gray-800">
						<div class="flex flex-col gap-2">
							<div class="grid sm:grid-cols-12 gap-1 pt-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200">
								<label class="inline-block sm:col-span-1 whitespace-nowrap text-sm font-medium mt-2.5">
		            	NIP
		          	</label>
				        <!-- End Col -->
				        <div class="mx-1 sm:col-span-5 sm:mt-1 relative sm:inline-flex">
									<input type="search" class="py-2 px-4 mb-2 block w-full border border-gray-200 rounded-lg text-sm disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="searchNip">
								</div>

								<label class="inline-block sm:col-span-1 whitespace-nowrap text-sm font-medium mt-2.5">
		            	Nama
		          	</label>
				        <!-- End Col -->
				        <div class="mx-1 sm:col-span-5 sm:mt-1 relative sm:inline-flex">
									<input type="search" class="py-2 px-4 mb-2 block w-full border border-gray-200 rounded-lg text-sm disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="searchNama" v-on:keyup.enter="CariData2">
								</div>

								<div class="mx-1 sm:col-span-5 sm:col-start-2 relative sm:inline-flex gap-2">
									<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-red-600 text-white hover:bg-red-700 disabled:opacity-50 disabled:pointer-events-none" @click="ResetData()">
									  Reset
									</button>
									<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="CariData2()">
									  Cari
									</button>	
								</div>
							</div>
						</div>
					</div>
				</div>
				<!-- End Accordion -->

				<div class="overflow-x-auto border border-t-2">
					<div class="min-w-full inline-block align-middle">
						<!-- Table -->
						<table class="min-w-full divide-y divide-gray-200">
							<thead class="bg-gray-50">
								<tr>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											No
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											NIP
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											Nama
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											NIK
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											NPWP
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											TTL
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											Usia
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											Alamat
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											Golongan Ruang
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											TMT Pangkat
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											Status Kepegawaian
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											Jabatan
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
											Eselon
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											Nama Unit Kerja
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											Nama Unit Organisasi
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											Pangkat / TMT CPNS
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											Pangkat /TMT PNS
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											Tingkat Pendidikan
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											Pendidikan Umum
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											Nama Sekolah
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											Jenis Kelamin
										</span>
									</th>
									<th scope="col" class="px-3 text-center">
										<span class="text-xs font-semibold uppercase text-gray-800">
											Peringkat
										</span>
									</th>
									<th scope="col" class="px-3 text-center"></th>
								</tr>
							</thead>
							<template v-if="pending">
								<td colspan="25" align="center">
									<i>Loading</i>
								</td>
							</template>
							<template v-else>
								{{data}}
								<tbody class="divide-y divide-gray-200">
									<template v-if="dataPegawai" v-for="(pegawai, idx) in paginate_pegawais" :key="idx">
										<tr class="odd:bg-white even:bg-gray-100" @click="EditData(pegawai.nip)">
											<td class="size-px p-1" style="vertical-align: top;">
												<span class="flex justify-center text-xs px-3 py-0">{{idx+from+1}}</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 150px;">
												<span class="flex justify-center text-center text-xs py-0">{{formatNip(pegawai.nip)}}</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 150px;">
												<span class="flex justify-start text-xs py-0">
													{{pegawai.gldepan ? pegawai.gldepan + ". ": ""}}
													{{pegawai.nama}}{{pegawai.glblk ? ", "+pegawai.glblk: ""}}
												</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0">{{pegawai.nik}}</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 150px;">
												<span class="flex justify-center text-center text-xs py-0">{{pegawai.npwp}}</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-center text-xs py-0">{{pegawai.ktlahir}}<br>{{$dayjs(pegawai.tlahir).format("DD-MM-YYYY").toString()}}</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 50px;">
												<span class="flex justify-start text-center text-xs py-0">{{pegawai.umur}}</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 200px;">
												<span class="flex justify-start text-xs py-0">{{pegawai.aljalan}}</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">{{nama_pangkat(pegawai.kgolru)}}</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{pegawai.tmtpang ? $dayjs(pegawai.tmtpang).format("DD-MM-YYYY").toString() : ""}}
												</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">{{pegawai.kstatus}}-{{pegawai.nstatus}}</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 200px;">
												<span class="flex justify-start text-xs py-0">{{pegawai.njab}}</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{pegawai.tmtjab ? $dayjs(pegawai.tmtjab).format("DD-MM-YYYY").toString() : ""}}
												</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{pegawai.nskjabat}}
												</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{pegawai.neselon}}
												</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 200px;">
												<span class="flex justify-start text-xs py-0">
													{{pegawai.nunker}}
												</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 200px;">
												<span class="flex justify-start text-xs py-0">
													{{pegawai.nopd}}
												</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; text-align: center;min-width: 100px;">
												<span class="flex justify-center align-center text-xs py-0">
													{{nama_pangkat(pegawai.kgolru_cpns)}}<br>{{pegawai.tmtcpns ? $dayjs(pegawai.tmtcpns).format("DD-MM-YYYY").toString() : ""}}
												</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; text-align: center; min-width: 100px;">
												<span class="flex justify-center text-xs py-0">
													{{nama_pangkat(pegawai.kgolru_pns)}}<br>{{pegawai.tmtpns ? $dayjs(pegawai.tmtpns).format("DD-MM-YYYY").toString() : ""}}
												</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-start text-xs py-0">
													{{pegawai.ntpu}}
												</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-start text-xs py-0">
													{{pegawai.npdum}}
												</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-start text-xs py-0">
													{{pegawai.nsek}}
												</span>
											</td>
											<td class="size-px p-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-start text-xs py-0">
													{{pegawai.kjkel == 1 ? "Laki - laki" : "Perempuan"}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top; min-width: 100px;">
												<span class="flex justify-start text-xs py-0">
													{{pegawai.nsek}}
												</span>
											</td>
											<td class="size-px px-1" style="vertical-align: top;">
												<button type="button" class="block" data-hs-overlay="#hs-ai-invoice-modal">
													<span class="px-3 py-1.5">
														<span class="py-1 px-2 inline-flex justify-center items-center gap-2 rounded-lg border font-medium bg-white text-gray-700 shadow-sm align-middle hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-white focus:ring-blue-600 transition-all text-xs">
															<svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
																<path d="M1.92.506a.5.5 0 0 1 .434.14L3 1.293l.646-.647a.5.5 0 0 1 .708 0L5 1.293l.646-.647a.5.5 0 0 1 .708 0L7 1.293l.646-.647a.5.5 0 0 1 .708 0L9 1.293l.646-.647a.5.5 0 0 1 .708 0l.646.647.646-.647a.5.5 0 0 1 .708 0l.646.647.646-.647a.5.5 0 0 1 .801.13l.5 1A.5.5 0 0 1 15 2v12a.5.5 0 0 1-.053.224l-.5 1a.5.5 0 0 1-.8.13L13 14.707l-.646.647a.5.5 0 0 1-.708 0L11 14.707l-.646.647a.5.5 0 0 1-.708 0L9 14.707l-.646.647a.5.5 0 0 1-.708 0L7 14.707l-.646.647a.5.5 0 0 1-.708 0L5 14.707l-.646.647a.5.5 0 0 1-.708 0L3 14.707l-.646.647a.5.5 0 0 1-.801-.13l-.5-1A.5.5 0 0 1 1 14V2a.5.5 0 0 1 .053-.224l.5-1a.5.5 0 0 1 .367-.27zm.217 1.338L2 2.118v11.764l.137.274.51-.51a.5.5 0 0 1 .707 0l.646.647.646-.646a.5.5 0 0 1 .708 0l.646.646.646-.646a.5.5 0 0 1 .708 0l.646.646.646-.646a.5.5 0 0 1 .708 0l.646.646.646-.646a.5.5 0 0 1 .708 0l.646.646.646-.646a.5.5 0 0 1 .708 0l.509.509.137-.274V2.118l-.137-.274-.51.51a.5.5 0 0 1-.707 0L12 1.707l-.646.647a.5.5 0 0 1-.708 0L10 1.707l-.646.647a.5.5 0 0 1-.708 0L8 1.707l-.646.647a.5.5 0 0 1-.708 0L6 1.707l-.646.647a.5.5 0 0 1-.708 0L4 1.707l-.646.647a.5.5 0 0 1-.708 0l-.509-.51z"/>
																<path d="M3 4.5a.5.5 0 0 1 .5-.5h6a.5.5 0 1 1 0 1h-6a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h6a.5.5 0 1 1 0 1h-6a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h6a.5.5 0 1 1 0 1h-6a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 0 1h-6a.5.5 0 0 1-.5-.5zm8-6a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 0 1h-1a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 0 1h-1a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 0 1h-1a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 0 1h-1a.5.5 0 0 1-.5-.5z"/>
															</svg>
														</span>
													</span>
												</button>
											</td>
										</tr>
									</template>
								</tbody>
							</template>
							
						</table>
						<!-- End Table -->	
					</div>
				</div>

				<!-- Footer -->
				<div class="px-6 py-4 grid gap-3 md:flex md:justify-between md:items-center border-t border-gray-200">
					<div>
						<!--<p class="text-sm text-gray-600">
							<span class="font-semibold text-gray-800">9</span> results
						</p>-->
						<select class="py-2 px-3 pe-9 block w-full border border-gray-200 rounded-lg text-sm disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="perPage">
							<option value="10">10</option>
							<option value="20">20</option>
							<option value="50">50</option>
							<option value="100">100</option>
						</select>
					</div>

					<div>
						<div class="inline-flex gap-x-2">
							<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none" @click="prevPage">
								<svg class="size-3" width="16" height="16" viewBox="0 0 16 15" fill="none" xmlns="http://www.w3.org/2000/svg">
									<path d="M10.506 1.64001L4.85953 7.28646C4.66427 7.48172 4.66427 7.79831 4.85953 7.99357L10.506 13.64" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
								</svg>
								Prev
							</button>

							<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none" @click="nextPage">
								Next
								<svg class="size-3" width="16" height="16" viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
									<path d="M4.50598 2L10.1524 7.64645C10.3477 7.84171 10.3477 8.15829 10.1524 8.35355L4.50598 14" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
								</svg>
							</button>
						</div>
					</div>
				</div>
				<!-- End Footer -->
			</div>
		</div>
		<!-- End Card -->
	</div>
	<!-- End Table Section -->

	<!-- Modal -->
	<div id="hs-ai-invoice-modal" class="hs-overlay hidden size-full fixed top-0 start-0 z-[80] overflow-x-hidden overflow-y-auto pointer-events-none">
		<div class="hs-overlay-open:mt-7 hs-overlay-open:opacity-100 hs-overlay-open:duration-500 mt-0 opacity-0 ease-out transition-all sm:max-w-lg sm:w-full m-3 sm:mx-auto">
			<div class="relative flex flex-col bg-white shadow-lg rounded-xl pointer-events-auto">
				<div class="relative overflow-hidden min-h-32 bg-gray-900 text-center rounded-t-xl">
					<!-- Close Button -->
					<div class="absolute top-2 end-2">
						<button type="button" class="inline-flex flex-shrink-0 justify-center items-center size-8 rounded-lg text-gray-500 hover:text-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-offset-2 focus:ring-offset-gray-900 transition-all text-sm" data-hs-overlay="#hs-ai-invoice-modal">
							<span class="sr-only">Close</span>
							<svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
						</button>
					</div>
					<!-- End Close Button -->

					<!-- SVG Background Element -->
					<figure class="absolute inset-x-0 bottom-0">
						<svg preserveAspectRatio="none" xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" viewBox="0 0 1920 100.1">
							<path fill="currentColor" class="fill-white" d="M0,0c0,0,934.4,93.4,1920,0v100.1H0L0,0z"></path>
						</svg>
					</figure>
					<!-- End SVG Background Element -->
				</div>

				<div class="relative z-10 -mt-12">
					<!-- Icon -->
					<span class="mx-auto flex justify-center items-center size-[62px] rounded-full border border-gray-200 bg-white text-gray-700 shadow-sm">
						<svg class="size-6" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
							<path d="M1.92.506a.5.5 0 0 1 .434.14L3 1.293l.646-.647a.5.5 0 0 1 .708 0L5 1.293l.646-.647a.5.5 0 0 1 .708 0L7 1.293l.646-.647a.5.5 0 0 1 .708 0L9 1.293l.646-.647a.5.5 0 0 1 .708 0l.646.647.646-.647a.5.5 0 0 1 .708 0l.646.647.646-.647a.5.5 0 0 1 .801.13l.5 1A.5.5 0 0 1 15 2v12a.5.5 0 0 1-.053.224l-.5 1a.5.5 0 0 1-.8.13L13 14.707l-.646.647a.5.5 0 0 1-.708 0L11 14.707l-.646.647a.5.5 0 0 1-.708 0L9 14.707l-.646.647a.5.5 0 0 1-.708 0L7 14.707l-.646.647a.5.5 0 0 1-.708 0L5 14.707l-.646.647a.5.5 0 0 1-.708 0L3 14.707l-.646.647a.5.5 0 0 1-.801-.13l-.5-1A.5.5 0 0 1 1 14V2a.5.5 0 0 1 .053-.224l.5-1a.5.5 0 0 1 .367-.27zm.217 1.338L2 2.118v11.764l.137.274.51-.51a.5.5 0 0 1 .707 0l.646.647.646-.646a.5.5 0 0 1 .708 0l.646.646.646-.646a.5.5 0 0 1 .708 0l.646.646.646-.646a.5.5 0 0 1 .708 0l.646.646.646-.646a.5.5 0 0 1 .708 0l.646.646.646-.646a.5.5 0 0 1 .708 0l.509.509.137-.274V2.118l-.137-.274-.51.51a.5.5 0 0 1-.707 0L12 1.707l-.646.647a.5.5 0 0 1-.708 0L10 1.707l-.646.647a.5.5 0 0 1-.708 0L8 1.707l-.646.647a.5.5 0 0 1-.708 0L6 1.707l-.646.647a.5.5 0 0 1-.708 0L4 1.707l-.646.647a.5.5 0 0 1-.708 0l-.509-.51z"/>
							<path d="M3 4.5a.5.5 0 0 1 .5-.5h6a.5.5 0 1 1 0 1h-6a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h6a.5.5 0 1 1 0 1h-6a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h6a.5.5 0 1 1 0 1h-6a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 0 1h-6a.5.5 0 0 1-.5-.5zm8-6a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 0 1h-1a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 0 1h-1a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 0 1h-1a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 0 1h-1a.5.5 0 0 1-.5-.5z"/>
						</svg>
					</span>
					<!-- End Icon -->
				</div>

				<!-- Body -->
				<div class="p-4 sm:p-7 overflow-y-auto">
					<div class="text-center">
						<h3 class="text-lg font-semibold text-gray-800">
							Invoice from Preline
						</h3>
						<p class="text-sm text-gray-500">
							Invoice #3682303
						</p>
					</div>

					<!-- Grid -->
					<div class="mt-5 sm:mt-10 grid grid-cols-2 sm:grid-cols-3 gap-5">
						<div>
							<span class="block text-xs uppercase text-gray-500">Amount paid:</span>
							<span class="block text-sm font-medium text-gray-800">$316.8</span>
						</div>
						<!-- End Col -->

						<div>
							<span class="block text-xs uppercase text-gray-500">Date paid:</span>
							<span class="block text-sm font-medium text-gray-800">April 22, 2020</span>
						</div>
						<!-- End Col -->

						<div>
							<span class="block text-xs uppercase text-gray-500">Payment method:</span>
							<div class="flex items-center gap-x-2">
								<svg class="size-5" width="400" height="248" viewBox="0 0 400 248" fill="none" xmlns="http://www.w3.org/2000/svg">
									<g clip-path="url(#clip0)">
									<path d="M254 220.8H146V26.4H254V220.8Z" fill="#FF5F00"/>
									<path d="M152.8 123.6C152.8 84.2 171.2 49 200 26.4C178.2 9.2 151.4 0 123.6 0C55.4 0 0 55.4 0 123.6C0 191.8 55.4 247.2 123.6 247.2C151.4 247.2 178.2 238 200 220.8C171.2 198.2 152.8 163 152.8 123.6Z" fill="#EB001B"/>
									<path d="M400 123.6C400 191.8 344.6 247.2 276.4 247.2C248.6 247.2 221.8 238 200 220.8C228.8 198.2 247.2 163 247.2 123.6C247.2 84.2 228.8 49 200 26.4C221.8 9.2 248.6 0 276.4 0C344.6 0 400 55.4 400 123.6Z" fill="#F79E1B"/>
									</g>
									<defs>
									<clipPath id="clip0">
									<rect width="400" height="247.2" fill="white"/>
									</clipPath>
									</defs>
								</svg>
								<span class="block text-sm font-medium text-gray-800">•••• 4242</span>
							</div>
						</div>
						<!-- End Col -->
					</div>
					<!-- End Grid -->

					<div class="mt-5 sm:mt-10">
						<h4 class="text-xs font-semibold uppercase text-gray-800">Summary</h4>

						<ul class="mt-3 flex flex-col">
							<li class="inline-flex items-center gap-x-2 py-3 px-4 text-sm border text-gray-800 -mt-px first:rounded-t-lg first:mt-0 last:rounded-b-lg">
								<div class="flex items-center justify-between w-full">
									<span>Payment to Front</span>
									<span>$264.00</span>
								</div>
							</li>
							<li class="inline-flex items-center gap-x-2 py-3 px-4 text-sm border text-gray-800 -mt-px first:rounded-t-lg first:mt-0 last:rounded-b-lg">
								<div class="flex items-center justify-between w-full">
									<span>Tax fee</span>
									<span>$52.8</span>
								</div>
							</li>
							<li class="inline-flex items-center gap-x-2 py-3 px-4 text-sm font-semibold bg-gray-50 border text-gray-800 -mt-px first:rounded-t-lg first:mt-0 last:rounded-b-lg">
								<div class="flex items-center justify-between w-full">
									<span>Amount paid</span>
									<span>$316.8</span>
								</div>
							</li>
						</ul>
					</div>

					<!-- Button -->
					<div class="mt-5 flex justify-end gap-x-2">
						<a class="py-2 px-3 inline-flex justify-center items-center gap-2 rounded-lg border font-medium bg-white text-gray-700 shadow-sm align-middle hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-white focus:ring-blue-600 transition-all text-sm" href="#">
							<svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/><polyline points="7 10 12 15 17 10"/><line x1="12" x2="12" y1="15" y2="3"/></svg>
							Invoice PDF
						</a>
						<a class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" href="#">
							<svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="6 9 6 2 18 2 18 9"/><path d="M6 18H4a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-2"/><rect width="12" height="8" x="6" y="14"/></svg>
							Print
						</a>
					</div>
					<!-- End Buttons -->

					<div class="mt-5 sm:mt-10">
						<p class="text-sm text-gray-500">If you have any questions, please contact us at <a class="inline-flex items-center gap-x-1.5 text-blue-600 decoration-2 hover:underline font-medium" href="#">example@site.com</a> or call at <a class="inline-flex items-center gap-x-1.5 text-blue-600 decoration-2 hover:underline font-medium" href="tel:+1898345492">+1 898-34-5492</a></p>
					</div>
				</div>
				<!-- End Body -->
			</div>
		</div>
	</div>
	<!-- End Modal -->
</template>
