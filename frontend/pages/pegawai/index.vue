<script setup>
const search = ref('');
const page = ref(1);

const from = ref(0)
const perPage = ref(10)

const dataPegawai = await useFetch('/api/pegawai')
console.log(dataPegawai)

const dataOPD = await useFetch('/api/opd')

const selOpd = ref([])
const searchOpd = ref(null)
//const selIdOpd = ref('1')


const selSubOpd = ref([])
const searchSubOpd = ref(null)
//const selIdSubOpd = ref(null)


const searchNama = ref(null)

const pegawais = computed({
	get(){
		//const {data} = dataPegawai
		//return data.value.pegawai
		return dataPegawai
	}
})

const opds_utama = computed({
	get(){
		const {data} = dataOPD
		//console.log(_filter)  
		var result = (data.value.opd).filter((opd) => {
			if(searchOpd.value != null && searchOpd.value != ""){
				return opd.nama.toLowerCase().includes(searchOpd.value.toLowerCase()) && opd.sfilter == 1
			}else{
				return opd.sfilter == 1
			}
		});

		return result
	}
})

const sub_opds_utama = computed({
	get(){
		const {data} = dataOPD
		//console.log(_filter)  
		console.log(data.value)
		var result = (data.value.opd).filter((opd) => {
			if(searchSubOpd.value != null && searchSubOpd.value != ""){
				return opd.nama.toLowerCase().includes(searchSubOpd.value.toLowerCase()) && opd.parent_opd == selOpd.value.id
			}else{
				return opd.parent_opd == selOpd.value.id
			}
		});

		return result
	}
})

const paginate_pegawais = computed({
	get(){
		return pegawais.value.slice(from.value, from.value + perPage.value) || null
	}
})

const prevPage = () => {
	if(from.value - perPage.value >= 0){
		from.value = from.value - perPage.value 
	}
}

const nextPage = () => {
	if(from.value + perPage.value <= pegawais.value.length){
		from.value = from.value + perPage.value
	}
}

const setSearchOPD = (opd) => {
	selOpd.value = opd
	//selIdOpd.value = opd.id

	//console.log(selIdOpd.value, opd)

	HSDropdown.close('#hs-dropdown-1')
}

const setSearchSubOPD = (opd) => {
	selSubOpd.value = opd

	HSDropdown.close('#hs-dropdown-2')
}

const ResetData = () => {
	selOpd.value = []
	selSubOpd.value = []
	searchNama.value = null
	//selIdOpd.value = '1'
	//searchOpd.value = null

	//selSubOpd.value = null
	//setSearchSubOPD.value = null
}

const CariData = async() => {
	console.log("cari data")

	const sdataPeg = await $fetch('/api/pegawai', {
		method: 'POST',
		data: {
			parent_opd: selSubOpd.value != [] ? selSubOpd.value.id : selOpd.value.id,
			nama: searchNama.value,
		}
	})

	console.log(sdataPeg)
	
}
//HSDropdown.open('#hs-dropdown')
</script>

<template>
	<!-- Table Section -->
	<div class="">
		<!-- Card -->
		<div class="flex flex-col">
			<div class="-m-1.5 overflow-x-auto">
				<div class="p-1.5 min-w-full inline-block align-middle">
					<div class="bg-white border border-gray-200 rounded-xl shadow-sm overflow-hidden">
						<!-- Header -->
						<div class="px-6 py-4 grid gap-3 md:flex md:justify-between md:items-center border-b border-gray-200">
							<div>
								<h2 class="text-xl font-semibold text-gray-800">
									Data Pegawai
								</h2>
								<p class="text-sm text-gray-600">
									Jumlah Pegawai: {{pegawais.length || 0}}
								</p>
							</div>

							<div>
								<div class="inline-flex gap-x-2">
									<a class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none" href="#">
										Lihat Semua
									</a>

									<a class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" href="#">
										<svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="M12 5v14"/></svg>
										Tambah
									</a>
								</div>
							</div>
						</div>
						<!-- End Header -->

						<!-- Accordion -->
						<div class="border-b border-gray-200 bg-gray-50">
							<div class="py-4 px-6 w-full flex items-center gap-2 font-semibold text-gray-800">

								<div class="grid grid-cols-1 lg:grid-cols-2 gap-y-2 gap-x-4">
									<div class="flex flex-col gap-2">
										<!--Unit Kerja -->
										<div class="sm:inline-flex sm:items-center space-y-2 sm:space-y-0 sm:space-x-3 w-full">
											<label for="inline-input-label-with-helper-text" class="block text-sm font-medium w-24">Unit Kerja</label>
											<div id="hs-dropdown-1" class="mt-60 mx-1 sm:mt-1 hs-dropdown relative sm:inline-flex [--auto-close:inside]">
												<button id="hs-dropdown" type="button" class="hs-dropdown-toggle py-2 px-4 inline-flex items-center justify-between gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none w-80">
													{{selOpd.nama || "Semua"}}
													<svg class="hs-dropdown-open:rotate-180 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m6 9 6 6 6-6"/></svg>
												</button>

												<div class="hs-dropdown-menu w-100 transition-[opacity,margin] duration hs-dropdown-open:opacity-100 opacity-0 hidden z-10 bg-white shadow-md rounded-lg p-2" aria-labelledby="hs-dropdown">
													<input type="search" class="py-2 px-4 mb-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="Cari Data Unor" v-model="searchOpd">
													<div class="max-h-[400px] overflow-y-auto
													  [&::-webkit-scrollbar]:w-2
													  [&::-webkit-scrollbar-track]:bg-gray-100
													  [&::-webkit-scrollbar-thumb]:bg-gray-300
													  dark:[&::-webkit-scrollbar-track]:bg-slate-700
													  dark:[&::-webkit-scrollbar-thumb]:bg-slate-500">
														<template v-if="opds_utama" v-for="(opd, idx) in opds_utama" :key="opd.id">
															<span class="flex items-center gap-x-3.5 py-2 px-3 rounded-lg text-sm text-gray-800 hover:bg-gray-100 focus:outline-none focus:bg-gray-100" @click="setSearchOPD(opd)">
																{{opd.nama}}
															</span> 
														</template>
													</div>
												</div>
											</div>
										</div>

										<!--Sub Unit Kerja -->
										<div class="sm:inline-flex sm:items-center space-y-2 sm:space-y-0 sm:space-x-3 w-full">
											<label for="inline-input-label-with-helper-text" class="block text-sm font-medium w-24">Sub Unit Kerja</label>
											<div id="hs-dropdown-2" class="mt-60 mx-1 sm:mt-1 hs-dropdown relative sm:inline-flex [--auto-close:inside]">
												<button id="hs-dropdown" type="button" class="hs-dropdown-toggle py-2 px-4 inline-flex items-center justify-between gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none w-80">
													{{selSubOpd.nama || "Semua"}}
													<svg class="hs-dropdown-open:rotate-180 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m6 9 6 6 6-6"/></svg>
												</button>

												<div class="hs-dropdown-menu w-100 transition-[opacity,margin] duration hs-dropdown-open:opacity-100 opacity-0 hidden z-10 bg-white shadow-md rounded-lg p-2" aria-labelledby="hs-dropdown">
													<input type="search" class="py-2 px-4 mb-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="Cari Data Sub Unor" v-model="searchSubOpd">
													<div class="max-h-[400px] overflow-y-auto
													  [&::-webkit-scrollbar]:w-2
													  [&::-webkit-scrollbar-track]:bg-gray-100
													  [&::-webkit-scrollbar-thumb]:bg-gray-300
													  dark:[&::-webkit-scrollbar-track]:bg-slate-700
													  dark:[&::-webkit-scrollbar-thumb]:bg-slate-500">
														<template v-if="sub_opds_utama" v-for="(opd, idx) in sub_opds_utama" :key="opd.id">
															<span class="flex items-center gap-x-3.5 py-2 px-3 rounded-lg text-sm text-gray-800 hover:bg-gray-100 focus:outline-none focus:bg-gray-100" @click="setSearchSubOPD(opd)">
																{{opd.nama}}
															</span> 
														</template>
													</div>
												</div>
											</div>
										</div>

										<div class="sm:inline-flex sm:items-center space-y-2 sm:space-y-0 sm:space-x-3 w-full">
											<label for="inline-input-label-with-helper-text" class="block text-sm font-medium w-24">Nama</label>
											<div class="w-80">
												<input type="text" class="py-2 px-4 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="searchNama">
											</div>
										</div>

										<div class="inline-flex items-center justify-end w-full gap-2">
											<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-red-600 text-white hover:bg-red-700 disabled:opacity-50 disabled:pointer-events-none" @click="ResetData()">
											  Reset
											</button>
											<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="CariData()">
											  Cari
											</button>
										</div>
									</div>
								</div>
							</div>
							<!--<button type="button" class="hs-collapse-toggle py-4 px-6 w-full flex items-center gap-2 font-semibold text-gray-800" id="hs-basic-collapse" data-hs-collapse="#hs-as-table-collapse">
								<svg class="hs-collapse-open:rotate-90 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m9 18 6-6-6-6"/></svg>
								Filter Data
							</button>
							<div id="hs-as-table-collapse" class="hs-collapse hidden w-full overflow-hidden transition-[height] duration-300" aria-labelledby="hs-basic-collapse">
								<div class="pb-4 px-6">
									<div class="flex items-center space-x-2">
										<span class="size-5 flex justify-center items-center rounded-full bg-blue-600 text-white">
											<svg class="flex-shrink-0 size-3.5" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"/></svg>
										</span>
										<span class="text-sm text-gray-800">
											There are no insights for this period.
										</span>
									</div>
								</div>
							</div>-->
						</div>
						<!-- End Accordion -->

						<!-- Table -->
						<table class="min-w-full divide-y divide-gray-200">
							<thead class="bg-gray-50">
								<tr>
									<th scope="col" class="p-3 text-center">
										<span class="text-xs font-semibold uppercase tracking-wide text-gray-800">
											No
										</span>
									</th>

									<th scope="col" class="p-3 w-[15%] text-center">
										<span class="text-xs font-semibold uppercase tracking-wide text-gray-800">
											NIP
										</span>
									</th>

									<th scope="col" class="p-3 w-[20%] text-center">
										<span class="text-xs font-semibold uppercase tracking-wide text-gray-800">
											Nama
										</span>
									</th>

									<th scope="col" class="p-3 w-[10%] text-center">
										<span class="text-xs font-semibold uppercase tracking-wide text-gray-800">
											Pangkat
										</span>
									</th>

									<th scope="col" class="p-3 w-[25%] text-center">
										<span class="text-xs font-semibold uppercase tracking-wide text-gray-800">
											Jabatan
										</span>
									</th>

									<th scope="col" class="p-3 w-[10%] text-center">
										<span class="text-xs font-semibold uppercase tracking-wide text-gray-800">
											Masa Kerja
										</span>
									</th>

									<th scope="col" class="p-3 w-[10%] text-center">
										<span class="text-xs font-semibold uppercase tracking-wide text-gray-800">
											Pendidikan
										</span>
									</th>

									<th scope="col" class="p-3 w-[10%] text-center">
										<span class="text-xs font-semibold uppercase tracking-wide text-gray-800">
											Usia
										</span>
									</th>

									<th scope="col" class="p-3 text-center"></th>
								</tr>
							</thead>

							<tbody class="divide-y divide-gray-200">
								<template v-if="pegawais" v-for="(pegawai, idx) in paginate_pegawais" :key="idx">
									<tr class="bg-white hover:bg-gray-50">
										<td class="size-px">
											<button type="button" class="block" data-hs-overlay="#hs-ai-invoice-modal">
												<span class="flex justify-center text-sm px-3 py-0">{{idx+from+1}}</span>
											</button>
										</td>
										<td class="size-px">
											<span class="flex justify-center text-center text-sm px-3 py-0">{{pegawai.nip}}</span>
										</td>
										<td class="size-px">
											<span class="flex justify-start text-sm px-3 py-0">{{pegawai.nama}}</span>
										</td>
										<td class="size-px">
											<span class="flex justify-center text-center text-sm px-3 py-0">{{pegawai.golongan_ruang}} {{pegawai.nama_pangkat ? "("+pegawai.nama_pangkat+")" : ""}}</span>
										</td>
										<td class="size-px">
											<span class="flex justify-start text-sm px-3 py-0">{{pegawai.njab}}</span>
										</td>
										<td class="size-px">
											<span class="flex justify-start text-sm px-3 py-0">28 Dec, 12:12</span>
										</td>
										<td class="size-px">
											<span class="flex justify-start text-sm px-3 py-0">28 Dec, 12:12</span>
										</td>
										<td class="size-px">
											<span class="flex justify-start text-sm px-3 py-0">28 Dec, 12:12</span>
										</td>
										<td class="size-px">
											<button type="button" class="block" data-hs-overlay="#hs-ai-invoice-modal">
												<span class="px-3 py-1.5">
													<span class="py-1 px-2 inline-flex justify-center items-center gap-2 rounded-lg border font-medium bg-white text-gray-700 shadow-sm align-middle hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-white focus:ring-blue-600 transition-all text-sm">
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
						</table>
						<!-- End Table -->

						<!-- Footer -->
						<div class="px-6 py-4 grid gap-3 md:flex md:justify-between md:items-center border-t border-gray-200">
							<div>
								<p class="text-sm text-gray-600">
									<span class="font-semibold text-gray-800">9</span> results
								</p>
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
