<script setup>
	const route = useRoute()
	const {snip} = route.params

	const disabled = computed({
		get(){
			//console.log(snip == 'null')
			if(snip == 'null'){
				return 'pointer-events-none italic'
			}
			return ''
		}
	})
</script>
<template>
	<nav class="relative z-0 flex border rounded-xl divide-x overflow-hidden">
		<NuxtLink class="group relative min-w-0 flex-1 bg-white py-4 px-4 text-sm font-medium text-center overflow-hidden focus:z-10 focus:outline-none focus:text-blue-600" :class="[$route.name.match('DataInduk') ? 'border-b-2 border-b-blue-600 text-gray-800' :'text-gray-500 hover:text-gray-700 hover:bg-gray-50', disabled]" :to="`/pegawai/${snip}/DataInduk/identitas`">
			Data Induk Pegawai
		</NuxtLink>
		<NuxtLink class="group relative min-w-0 flex-1 bg-white py-4 px-4 text-sm font-medium text-center overflow-hidden focus:z-10 focus:outline-none focus:text-blue-600" :class="[$route.name.match('RiwayatPegawai') ? 'border-b-2 border-b-blue-600 text-gray-800' :'text-gray-500 hover:text-gray-700 hover:bg-gray-50', disabled]" :to="`/pegawai/${snip}/RiwayatPegawai/pangkat`">
			Riwayat Pegawai
		</NuxtLink>
		<NuxtLink class="group relative min-w-0 flex-1 bg-white py-4 px-4 text-sm font-medium text-center overflow-hidden focus:z-10 focus:outline-none focus:text-blue-600" :class="[$route.name.match('RiwayatPendidikan') ? 'border-b-2 border-b-blue-600 text-gray-800' :'text-gray-500 hover:text-gray-700 hover:bg-gray-50', disabled]" :to="`/pegawai/${snip}/RiwayatPendidikan/pendum`">
			Riwayat Pendidikan
		</NuxtLink>
		<NuxtLink class="group relative min-w-0 flex-1 bg-white py-4 px-4 text-sm font-medium text-center overflow-hidden focus:z-10 focus:outline-none focus:text-blue-600" :class="[$route.name.match('DataKeluarga') ? 'border-b-2 border-b-blue-600 text-gray-800' :'text-gray-500 hover:text-gray-700 hover:bg-gray-50', disabled]" :to="`/pegawai/${snip}/DataKeluarga/1`">
			Data Keluarga
		</NuxtLink>
		<NuxtLink class="group relative min-w-0 flex-1 bg-white py-4 px-4 text-sm font-medium text-center overflow-hidden focus:z-10 focus:outline-none focus:text-blue-600" :class="[$route.name.match('Singkronisasi') ? 'border-b-2 border-b-blue-600 text-gray-800' :'text-gray-500 hover:text-gray-700 hover:bg-gray-50', disabled]" :to="`/pegawai/${snip}/Singkronisasi/token`">
			Singkronisasi
		</NuxtLink>
	</nav>
</template>