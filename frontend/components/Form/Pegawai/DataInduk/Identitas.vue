<script setup>
const agamas = ref([])
const statuss = ref([])
const jeniss = ref([])
const duduks = ref([])
const kawins = ref([])
const goldars = ref([])
const provs = ref([])
const kabs = ref([])
const kecs = ref([])
const kels = ref([])
const kpes = ref([])

const { pending: pendingRef, data, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
  	const [agama, dataPangkat, dataJenisJabatan, dataEselon, status, jenis, duduk, kawin, goldar, prov, kpe] = await Promise.all([
    	$fetch('/api/gets/agama'),
    	$fetch('/api/gets/pangkat'),
    	$fetch('/api/gets/jenis_jabatan'),
    	$fetch('/api/gets/eselon'),
    	$fetch('/api/gets/status_pegawai'),
    	$fetch('/api/gets/jenis_pegawai'),
    	$fetch('/api/gets/kedudukan_pegawai'),
    	$fetch('/api/gets/jenis_kawin'),
    	$fetch('/api/gets/jenis_goldar'),
    	$fetch('/api/posts/wilayah', {
    		method: 'POST',
    		body: JSON.stringify({
    			'twil': "1"
    		})
    	}),
    	$fetch('/api/gets/kpe'),
  	])

  	agamas.value = agama
  	statuss.value = status
  	jeniss.value = jenis
  	duduks.value = duduk
  	kawins.value = kawin
  	goldars.value = goldar
  	provs.value = prov
  	kpes.value = kpe
});

onMounted(() => {
	refreshNuxtData(["getDataRef"])
	console.log("onMounted Identitas")
})

const kprov = ref(0)

watch(kprov, async(newX) => {
  	var a = newX.slice(0,2)
  	console.log(a)

  	var result = await $fetch('/api/posts/wilayah', {
		method: 'POST',
		body: JSON.stringify({
			'twil': "2",
			'kprov': a
		})
	})

	kabs.value = result
	kecs.value = 0
	kels.value = 0
})

const kkab = ref(0)

watch(kkab, async(newX) => {
  	var a = newX.slice(0,2)
  	var b = newX.slice(2,4)

  	console.log(b)

  	var result = await $fetch('/api/posts/wilayah', {
		method: 'POST',
		body: JSON.stringify({
			'twil': "3",
			'kprov': a,
			'kkab': b,
		})
	})

	kecs.value = result
	kels.value = 0
})

const kkec = ref(0)

watch(kkec, async(newX) => {
  	var a = newX.slice(0,2)
  	var b = newX.slice(2,4)
  	var c = newX.slice(4,6)

  	var result = await $fetch('/api/posts/wilayah', {
		method: 'POST',
		body: JSON.stringify({
			'twil': "4",
			'kprov': a,
			'kkab': b,
			'kkec': c,
		})
	})

	kels.value = result
})

const kkel = ref(0)


</script>
<template>
	<div class="mx-auto">
		<!-- Card -->
		<div class="bg-white rounded-xl shadow py-4 px-7">
			<div class="mb-8">
				<h2 class="text-xl font-bold text-gray-800">Identitas Pegawai</h2>
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

					<div class="sm:col-span-9 sm:col-start-4">
						<div class="flex">
							<input type="text" name="hs-input-with-add-on-url" id="hs-input-with-add-on-url" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-s-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="ID PNS">
							<button type="button" class="py-3 px-4 inline-flex justify-center items-center gap-x-2 text-xs font-semibold rounded-e-md border border-transparent bg-gray-600 text-white hover:bg-gray-700 disabled:opacity-50 disabled:pointer-events-none whitespace-nowrap">
								<svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
  									<path fill-rule="evenodd" d="M15.312 11.424a5.5 5.5 0 0 1-9.201 2.466l-.312-.311h2.433a.75.75 0 0 0 0-1.5H3.989a.75.75 0 0 0-.75.75v4.242a.75.75 0 0 0 1.5 0v-2.43l.31.31a7 7 0 0 0 11.712-3.138.75.75 0 0 0-1.449-.39Zm1.23-3.723a.75.75 0 0 0 .219-.53V2.929a.75.75 0 0 0-1.5 0V5.36l-.31-.31A7 7 0 0 0 3.239 8.188a.75.75 0 1 0 1.448.389A5.5 5.5 0 0 1 13.89 6.11l.311.31h-2.432a.75.75 0 0 0 0 1.5h4.243a.75.75 0 0 0 .53-.219Z" clip-rule="evenodd" />
								</svg>
								ID PNS
							</button>
						</div>
					</div>

					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">NIP</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nip">
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">NIK/NPWP</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="NIK" v-model="nik">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="NPWP" v-model="npwp">
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Nama</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-[35%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="gelar depan" v-model="gldepan">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nama">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-[35%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="Gelar Belakang" v-model="glblk">
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
		  				<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Tempat / Tgl Lahir</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
		  				<div class="sm:flex gap-2">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="ktlahir">
							<input id="af-account-full-name" type="date" class="py-2 px-3 block w-[50%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tlahir">
		  				</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Jenis Kelamin</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<div class="sm:flex">
							<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kjkel">
							  	<option selected>Pilih Jenis Kelamin</option>
							  	<option value="1">1 - Laki-laki</option>
							  	<option value="2">2 - Perempuan</option>
							</select>
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-2 sm:col-start-8">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Agama</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<div class="sm:flex">
					  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kagama">
							  	<option selected>Pilih Agama</option>
							  	<template  v-for="(item, idx) in agamas" :key="idx">
								  	<option :value="item.id">{{item.id}} - {{item.nama}}</option>
								  </template>
							</select>
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Status Kepegawaian</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
					  	<div class="sm:flex">
					  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kstatus">
							  	<option selected>Pilih Status Kepegawaian</option>
							  	<template  v-for="(item, idx) in statuss" :key="idx">
								  	<option :value="item.kstatus">{{item.kstatus}} - {{item.nama}}</option>
								  </template>
							</select>
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Jenis Kepegawaian</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
					  	<div class="sm:flex">
					  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kjpeg">
							  	<option selected>Pilih Jenis Kepegawaian</option>
							  	<template  v-for="(item, idx) in jeniss" :key="idx">
								  	<option :value="item.id">{{item.id}} - {{item.nama}}</option>
								  </template>
							</select>
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Kedudukan Kepegawaian</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
					  	<div class="sm:flex">
					  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kduduk">
							  	<option selected>Pilih Kedudukan Kepegawaian</option>
							  	<template  v-for="(item, idx) in duduks" :key="idx">
								  	<option :value="item.id">{{item.id}} - {{item.nama}}</option>
								  </template>
							</select>
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Status Perkawinan</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<div class="sm:flex">
					  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kskawin">
							  	<option selected>Pilih Status Perkawinan</option>
							  	<template  v-for="(item, idx) in kawins" :key="idx">
								  	<option :value="item.id">{{item.id}} - {{item.nama}}</option>
								  </template>
							</select>
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-2 sm:col-start-8">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Golongan Darah</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<div class="sm:flex">
					  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kgoldar">
							  	<option selected>Pilih Golongan Darah</option>
							  	<template  v-for="(item, idx) in goldars" :key="idx">
								  	<option :value="item.id">{{item.id}} - {{item.nama}}</option>
								  </template>
							</select>
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Alamat</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<textarea id="af-account-bio" class="py-2 px-3 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" rows="3" placeholder="Alamat..." v-model="aljalan"></textarea>
						</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">RT/RW</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<div class="sm:flex gap-2">
					  		<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="RT" v-model="alrt">
							<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="RW" v-model="alrw">
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-2 sm:col-start-8">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Kode Pos</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<div class="sm:flex">
					  		<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kpos">
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Nomor Telpon/ WA</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
					  	<div class="sm:flex gap-2">
					  		<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="Nomor HP/Telpon" v-model="altelp">
					  		<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="Nomor Whatsapp" v-model="altelpwa">
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Provinsi</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
					  	<div class="sm:flex">
					  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kprov">
							  	<option selected>Pilih Provinsi</option>
							  	<template  v-for="(item, idx) in provs" :key="idx">
								  	<option :value="item.kwil">{{item.kwil}} - {{item.nwil}}</option>
								  </template>
							</select>
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Kabupaten/Kota</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
					  	<div class="sm:flex">
					  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kkab">
							  	<option selected>Pilih Kabupaten/Kota</option>
							  	<template  v-for="(item, idx) in kabs" :key="idx">
								  	<option :value="item.kwil">{{item.kwil}} - {{item.nwil}}</option>
								  </template>
							</select>
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Kecamatan</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
					  	<div class="sm:flex">
					  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kkec">
							  	<option selected>Pilih Kecamatan</option>
							  	<template  v-for="(item, idx) in kecs" :key="idx">
								  	<option :value="item.kwil">{{item.kwil}} - {{item.nwil}}</option>
								  </template>
							</select>
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Desa/Kelurahan</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
					  	<div class="sm:flex">
					  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="kkel">
							  	<option selected>Pilih Desa/Kelurahan</option>
							  	<template  v-for="(item, idx) in kels" :key="idx">
								  	<option :value="item.kwil">{{item.kwil}} - {{item.nwil}}</option>
								  </template>
							</select>
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Nomor Karpeg</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<div class="sm:flex">
					  		<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nkarpeg">
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-2 sm:col-start-8">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">No Karis/Karsu</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<div class="sm:flex">
					  		<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="nkaris_su">
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Nomor Taspen</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
					  	<div class="sm:flex">
					  		<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg sm:mt-0 sm:first:ms-0 text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="ntaspen">
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Status Kepemilikan KPE</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<div class="sm:flex">
					  		<select class="py-2.5 px-2 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="stat_kpe">
							  	<option selected>Pilih Status KPE</option>
							  	<template  v-for="(item, idx) in kpes" :key="idx">
								  	<option :value="item.kkpe">{{item.kkpe}} - {{item.nkpe}}</option>
								  </template>
							</select>
					  	</div>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-2 sm:col-start-8">
					  	<label for="af-account-gender-checkbox" class="inline-block text-sm text-gray-800 mt-2.5">Tanggal KPE</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
					  	<div class="sm:flex">
					  		<input id="af-account-full-name" type="date" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tgl_kpe">
					  	</div>
					</div>
					<!-- End Col -->
	 			</div>
	  			<!-- End Grid -->

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