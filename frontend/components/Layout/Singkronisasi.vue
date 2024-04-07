<script setup>
	const route = useRoute()
</script>
<style scoped>
.router-link-active {
  @apply border-e-2 border-blue-600 text-blue-600
}
.router-link-active:hover {
  @apply font-medium;
}
</style>
<template>
	<div class="mx-auto">
		<!-- Card -->
		<div class="bg-white rounded-xl shadow p-4">
			<TopBarPegawai />
			<div class="mt-5">
				<div class="grid sm:grid-cols-12 gap-4">
					<div class="sm:col-span-2">
						<div class="flex flex-wrap">
						  	<div class="w-full">
						  		<nav class="-me-0.5 flex flex-col">
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/token">
							      		Token Info
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/DataInduk/identitas">
							      		Data Utama
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/angka-kredit">
							      		Angka Kredit
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/riwayat-stlud">
							    	  	Riwayat STLUD
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/jabatan">
							    	  	Jabatan
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/tugas-tambahan">
							    	  	Tugas Tambahan
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/gaji-berkala">
							    	  	Gaji Berkala
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/pmk">
							    	  	Penyesuaian Masa Kerja
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/pindah-instansi">
							    	  	Pindah Instansi
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/skp">
							    	  	SKP
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/penghargaan">
							    	  	Penghargaan / Tanda Jasa
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/hukdis">
							    	  	Hukuman Disiplin
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/cuti">
							    	  	Cuti
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/organisasi">
							    	  	Organisasi
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/tugas-ln">
							    	  	Tugas Luar Negeri
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/bahasa">
							    	  	Penguasaan Bahasa
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/pemberhentian-jft">
							    	  	Pemberhentian dari Jabatan Fungsional
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" to="/pegawai/form/Singkronisasi/tubel">
							    	  	Tugas Belajar
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