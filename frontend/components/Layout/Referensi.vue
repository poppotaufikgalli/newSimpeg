<script setup>
	const route = useRoute()
</script>
<style scoped>
#leftNav > .router-link-active {
  @apply border-e-2 border-blue-600 text-blue-600
}
.router-link-active:hover {
  @apply font-medium;
}
</style>
<template>
	<div class="mx-auto">
		<!-- Card -->
		<div class="bg-white min-h-[calc(100vh-140px)] rounded-xl shadow p-4 mb-6">
			<div class="mt-5">
				<div class="grid sm:grid-cols-12 gap-4">
					<div class="sm:col-span-2">
						<div class="flex flex-wrap">
					  	<div class="w-full">
					  		<nav class="-me-0.5 flex flex-col" id="leftNav">
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/referensi/opd`">
						      		OPD
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/referensi/pengguna`">
						      		Pengguna
						    	</NuxtLink>
						    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/referensi/singkronisasi`">
						      		Singkronisasi BKN
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