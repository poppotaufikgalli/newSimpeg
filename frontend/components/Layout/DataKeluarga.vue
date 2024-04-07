<script setup>
const route = useRoute()

const { pending, data: jnsDiklat, refresh} = await useLazyAsyncData('getJenisKeluarga', () => CariData())

const CariData = async() => {
		console.log("CariData Jenis Keluarga")

		return await $fetch('/api/gets/jenis_keluarga');
}
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
							    	<template v-if="jnsDiklat" v-for="(item, idx) in jnsDiklat" :key="idx">
							    		<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" 
							    		:to="{ path: '/pegawai/form/DataKeluarga/'+item.id}">
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