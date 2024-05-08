<script setup>
	const route = useRoute()
	const {snip} = route.params
</script>
<style scoped>
.router-link-active {
  @apply bg-blue-800 text-gray-200;
}
.router-link-active:hover {
  @apply font-medium;
}
</style>
<template>
	<div class="mx-auto">
		<!-- Card -->
		<div class="bg-white rounded-xl shadow p-4 mb-6">
			<TopBarPegawai />
			<div class="mt-5">
				<div class="grid sm:grid-cols-12 gap-4">
					<div class="sm:col-span-2">
						<div class="flex flex-wrap">
					  	<div class="w-full">
					  		<nav class="-me-0.5 flex flex-col border divide-y rounded-md overflow-hidden">
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPegawai/pangkat`">
						      		Pangkat
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPegawai/angka-kredit`">
						      		Angka Kredit
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPegawai/stlud`">
						    	  	Riwayat STLUD
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPegawai/jabatan`">
						    	  	Jabatan
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPegawai/tugas-tambahan`">
						    	  	Tugas Tambahan
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPegawai/gaji-berkala`">
						    	  	Gaji Berkala
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPegawai/pmk`">
						    	  	Penyesuaian Masa Kerja
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPegawai/pindah-instansi`">
						    	  	Pindah Instansi
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPegawai/skp`">
						    	  	SKP
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPegawai/penghargaan`">
						    	  	Penghargaan / Tanda Jasa
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPegawai/hukdis`">
						    	  	Hukuman Disiplin
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPegawai/cuti`">
						    	  	Cuti
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPegawai/organisasi`">
						    	  	Organisasi
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPegawai/tugas-ln`">
						    	  	Tugas Luar Negeri
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPegawai/bahasa`">
						    	  	Penguasaan Bahasa
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPegawai/pemberhentian-jft`">
						    	  	Pemberhentian dari Jabatan Fungsional
						    	</NuxtLink>
						  	</nav>
					  	</div>
						</div>
					</div>
					<div class="sm:col-span-10">
						<slot />
					</div>
				</div>
			</div>
		</div>
	</div>
</template>