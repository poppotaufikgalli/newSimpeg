<script setup>
const route = useRoute()
console.log(route.query.akhir)

const pejabats = ref([])
const pangkats = ref([])
const jenis_kps = ref([])

const { pending: pendingRef, data, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
  	const [pejabat, pangkat, jenis_kp] = await Promise.all([
    	$fetch('/api/gets/pejabat'),
    	$fetch('/api/gets/pangkat'),
    	$fetch('/api/gets/jenis_kp'),
  	])

  	pejabats.value = pejabat
  	pangkats.value = pangkat
  	jenis_kps.value = jenis_kp
});

const kpej = ref(null)
const kgolru = ref(null)

onMounted(() => {
	refreshNuxtData(["getDataRef"])
})
</script>
<template>
	<LayoutDataInduk>
		<div class="mx-auto">
			<!-- Card -->
			<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
				<div class="mb-8">
					<div class="inline-flex justify-between items-center w-full gap-x-2">
						<h2 class="text-xl font-bold text-blue-600">Pangkat {{route.query.akhir ? 'Terakhir' : ''}}</h2>
						<div class="flex gap-x-2 w-[60%]">
							<input type="text" name="hs-input-with-add-on-url" id="hs-input-with-add-on-url" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP">
							<input type="text" name="hs-input-with-add-on-url" id="hs-input-with-add-on-url" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="Nama">
						</div>
					</div>
				</div>

				<form>
					<!-- Grid -->
					<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
						<!--<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Foto</label>
						</div>

						<div class="sm:col-span-9">
							<div class="flex items-center gap-5">
								<img class="inline-block size-16 rounded-full ring-2 ring-white" src="https://preline.co/assets/img/160x160/img1.jpg" alt="Image Description">
								<div class="flex gap-x-2">
									<div>
										<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none">
											<svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/><polyline points="17 8 12 3 7 8"/><line x1="12" x2="12" y1="3" y2="15"/></svg>
											Upload Foto
										</button>
									</div>
								</div>
							</div>
						</div>-->
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Nota BKN</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nip" placeholder="xx-xxxxxxxxxxxx">
								<input id="af-account-full-name" type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nip">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Pejabat Penetap</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<SearchSelect id0="pejabats" :dataList="pejabats" keyList="kpej" namaList="npej" v-model="kpej"/>
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">SK Pangkat</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nip">
								<input id="af-account-full-name" type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nip">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">TMT Pangkat</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input id="af-account-full-name" type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nip">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Pangkat / Golongan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<SearchSelect id0="pangkats" :dataList="pangkats" v-model="kgolru" keyList="id" namaList="pangkat_gol"/>
							</div>
						</div>

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Masa Kerja</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-[20%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nip" placeholder="thn">
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-[20%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nip" placeholder="bln">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Jenis Kenaikan Pangkat</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<SearchSelect id0="jenis_kps" :dataList="jenis_kps" v-model="jns_kp" keyList="id" namaList="nama"/>
							</div>
						</div>

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Angka Kredit</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<div class="sm:flex gap-2">
								<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nip">
							</div>
						</div>
						<!-- End Col -->

					</div>
					<div class="mt-5 grid sm:grid-cols-12 gap-x-2">
						<div class="sm:col-span-6 sm:col-start-4">
							<div class="sm:flex gap-2">
								<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none">
					  				Batal
								</button>
								<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none">
					  				Simpan Data
								</button>	
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3 flex justify-end">
							<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-green-600 text-white hover:bg-green-700 disabled:opacity-50 disabled:pointer-events-none">
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