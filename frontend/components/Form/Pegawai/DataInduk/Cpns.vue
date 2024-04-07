<script setup>
const props = defineProps({
  	pangkats: {
    	type: Object,
    	default: [],
  	},
  	pejabats: {
    	type: Object,
    	default: [],
  	},
});

//const {data} = props

console.log(props)

onMounted(() => {
	//refreshNuxtData(["getDataRef"])
	console.log("onMounted CPNS")
})
</script>
<template>
	<div class="mx-auto">
		<!-- Card -->
		<div class="bg-white rounded-xl shadow py-4 px-7">
			<div class="mb-8">
				<h2 class="text-xl font-bold text-gray-800">Data CPNS</h2>
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
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">NIP</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nip">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nama">
						</div>
					</div>
					<!-- End Col -->
					<div class="sm:col-span-12">
						<hr>
					</div>
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
							<select class="py-2.5 px-2 block w-[60%] border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kagama">
							  	<option selected>Pilih Pejabat Penetap</option>
							  	<template  v-for="(item, idx) in pejabats" :key="idx">
								  	<option :value="item.kpej">{{item.kpej}} - {{item.npej}}</option>
								  </template>
							</select>
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">SK CPNS</label>
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
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">TMT CPNS</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input id="af-account-full-name" type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nip">
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Golongan CPNS</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<select class="py-2.5 px-2 block w-[30%] border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kagama">
							  	<option selected>Pilih Golongan CPNS</option>
							  	<template  v-for="(item, idx) in pangkats" :key="idx">
								  	<option :value="item.id">{{item.id}} - {{item.nama}}</option>
								  </template>
							</select>
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Nomor STTPP</label>
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
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Tanggal Melaksanakan Tugas</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input id="af-account-full-name" type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nip">
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Masa Kerja Sebelum CPNS</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-[20%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nip" placeholder="thn">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-[20%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nip" placeholder="bln">
						</div>
					</div>
					<!-- End Col -->

				</div>
				<div class="mt-5 flex justify-end gap-x-2">
					<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none">
		  				Batal
					</button>
					<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none">
		  				Simpan Data
					</button>
	  			</div>
			</form>
  		</div>
  		<!-- End Card -->
	</div>
<!-- End Card Section -->
</template>