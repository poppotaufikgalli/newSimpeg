<script setup>
const route = useRoute()
console.log(route.query.akhir)

const pejabats = ref([])
const jenis_jabatans = ref([])
const eselons = ref([])
const opds = ref([])

const kpej = ref(null)
const kgolru = ref(null)
const jns_jab = ref(null)
const kjab = ref(null)
const fjab = ref(null)
const keselon = ref(null)
const kunker = ref(null)

const { pending: pendingRef, data, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
	//loading.value = true
  	const [opd, jenis_jabatan, eselon] = await Promise.all([
    	$fetch('/api/gets/opd'),
    	$fetch('/api/gets/jenis_jabatan'),
    	$fetch('/api/gets/eselon'),
  	])

  	opds.value = opd
  	jenis_jabatans.value = jenis_jabatan
  	eselons.value = eselon
});

const proxyEselons = computed({
	get(){
		var sjns_jab = jns_jab.value == 1 ? jns_jab.value : 4
		return _filter(eselons.value, function(o) {
			return o.id_jenis_jabatan == sjns_jab
		})
	}
})

const formasiJabatans = computed({
	get(){
		if(opds.value != null){
	  		var result = _find(opds.value, function(o) {
	  			if(kunker.value != null){
	  				return o.id == kunker.value
	  			}
	  		})

	  		if(result){
	  			return result.formasi_jabatan
	  		}
	  	}
	  	return null
	}
})

watch(fjab, (value) => {
	if(value){
		var result = _find(formasiJabatans.value, function(o){
			return o.id == value
		})
		
		const {id_eselon, id_jenis_jabatan} = result
		console.log(id_eselon)
		jns_jab.value = id_jenis_jabatan
		keselon.value = id_eselon.toString()
		
	}
})

const { pending, data: jabatans, refresh: refreshDataJabatan } = await useAsyncData('getDataJabatan', () => CariJabatan())

const CariJabatan = async() => {
	var body = JSON.stringify({
		'id_opd' : kunker.value,
		'id_jenis_jabatan' : jns_jab.value,
		'id_eselon' : keselon.value
	});

	console.log("CariJabatan")

	body = body == "{}" ? null : body

	return await $fetch('/api/posts/formasi_jabatan', {
		method: 'POST',
		body: body,
	});
}

watch(jns_jab, (value)=> {
	/*if(value != 1){
		keselon.value = '99'
	}else{
		keselon.value = null
	}*/
})

watch(keselon, (value) => {
	//refreshNuxtData(["getDataJabatan"])
})

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
						<h2 class="text-xl font-bold text-blue-600">Tugas Tambahan {{route.query.akhir ? 'Terakhir' : ''}}</h2>
						<div class="flex gap-x-2 w-[60%]">
							<input type="text" name="hs-input-with-add-on-url" id="hs-input-with-add-on-url" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP">
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
							<div class="sm:flex gap-2">
								<SearchSelect id0="pejabats" :dataList="pejabats" keyList="kpej" namaList="npej" v-model="kpej"/>
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">SK Jabatan</label>
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
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">TMT Jabatan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<input id="af-account-full-name" type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nip">
							</div>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Unit Kerja</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<SearchSelect id0="opds" :dataList="opds" keyList="id" namaList="nama" v-model="kunker"/>
							</div>
						</div>

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Tugas Tambahan Sebagai</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<SearchSelect id0="jenis_jabatans" :dataList="jenis_jabatans" keyList="id" namaList="nama" v-model="jns_jab"/>
							</div>
						</div>

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Jabatan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<SearchSelect id0="jabatans" :dataList="jabatans" keyList="id" namaList="nama" v-model="kjab"/>
							</div>
						</div>

						<div class="sm:col-span-3">
							<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Nama Tugas Tambahan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<textarea id="af-account-bio" class="py-2 px-3 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" rows="3" placeholder="Nama Jabatan" v-model="njab"></textarea>
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