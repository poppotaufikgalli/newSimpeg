<script setup>
</script>
<template>
	<div class="mx-auto">
		<!-- Card -->
		<div class="bg-white rounded-xl shadow p-4">
			<nav class="relative z-0 flex border rounded-xl divide-x divide-gray-200 overflow-hidden">
				<NuxtLink class="group relative min-w-0 flex-1 bg-white py-4 px-4 border-b-2 border-b-blue-600 text-gray-800 text-sm font-medium text-center overflow-hidden focus:z-10 focus:outline-none focus:text-blue-600" aria-current="page" to="/Pegawai/form/DataInduk/identitas">
					Data Induk Pegawai
				</NuxtLink>
				<NuxtLink class="group relative min-w-0 flex-1 bg-white py-4 px-4 text-gray-500 hover:text-gray-700 text-sm font-medium text-center overflow-hidden hover:bg-gray-50 focus:z-10 focus:outline-none focus:text-blue-600" href="/Pegawai/form/RiwayatPegawai/Pangkat">
					Riwayat Pegawai
				</NuxtLink>
				<NuxtLink class="group relative min-w-0 flex-1 bg-white py-4 px-4 text-gray-500 hover:text-gray-700 text-sm font-medium text-center overflow-hidden hover:bg-gray-50 focus:z-10 focus:outline-none focus:text-blue-600" href="#">
					Riwayat Pendidikan
				</NuxtLink>
				<NuxtLink class="group relative min-w-0 flex-1 bg-white py-4 px-4 text-gray-500 hover:text-gray-700 text-sm font-medium text-center overflow-hidden hover:bg-gray-50 focus:z-10 focus:outline-none focus:text-blue-600" href="#">
					Data Keluarga
				</NuxtLink>
			</nav>
			<LayoutDataInduk>
				<slot />
			</LayoutDataInduk>
		</div>
	</div>
</template>