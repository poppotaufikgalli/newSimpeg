<script setup>
const route = useRoute()
const {snip} = route.params

const { pending, data: jnsDiklat, refresh} = await useLazyAsyncData('getJenisDiklat', async() => {
		console.log("CariData Jenis Diklat")

		return await $fetch('/api/gets/jenis_diklat');
});

onMounted(() => {
		refreshNuxtData(["getJenisDiklat"])
})
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
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPendidikan/pendum`">
							      		Pendidikan Umum
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/RiwayatPendidikan/tubel`">
							      		Tugas Belajar
							    	</NuxtLink>
							    	<template v-if="jnsDiklat" v-for="(item, idx) in jnsDiklat" :key="idx">
							    		<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" 
							    		:to="{ path: `/pegawai/${snip}/RiwayatPendidikan/diklat/${item.id}`}">
							    	  	{{item.nama}}
							    		</NuxtLink>	
							    	</template>
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