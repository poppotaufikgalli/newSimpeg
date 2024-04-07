<script setup>
	const route = useRoute()
	const {snip} = route.params
	//console.log(snip)
	
	const disabled = computed({
		get(){
			if(snip == 'null'){
				return true
			}
			return false
		}
	})
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
					<div class="sm:col-span-2" v-if="!disabled">
						<div class="flex flex-wrap">
							<div class="mb-5">
				    			<img src="https://upload.wikimedia.org/wikipedia/commons/6/65/No-Image-Placeholder.svg">
				    		</div>
						  	<div class="w-full">
						  		<nav class="-me-0.5 flex flex-col">
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/DataInduk/identitas`">
							      		Identitas Pegawai
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/DataInduk/cpns`">
							      		CPNS
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/DataInduk/pns`">
							    	  	PNS
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="{ path: '/pegawai/'+snip+'/DataInduk/pangkat', query: { akhir: true}}" >
							    	  	Pangkat Terakhir
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="{ path: '/pegawai/'+snip+'/DataInduk/jabatan', query: { akhir: true}}" >
							    	  	Jabatan Terakhir
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="{ path: '/pegawai/'+snip+'/DataInduk/tugas-tambahan', query: { akhir: true}}" >
							    	  	Tugas Tambahan
							    	</NuxtLink>
							    	<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray-500 hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="{ path: '/pegawai/'+snip+'/DataInduk/gaji-berkala', query: { akhir: true}}" >
							    	  	Gaji Berkala Terakhir
							    	</NuxtLink>
							  	</nav>
							  	
						    	<!--<nav class="flex flex-col space-t-1" aria-label="Tabs" role="tablist" data-hs-tabs-vertical="true">
						      		<button type="button" class="hs-tab-active:border-blue-500 hs-tab-active:text-blue-600 hs-tab-active:bg-blue-100 p-2 rounded-e-lg pe-4 inline-flex items-center gap-x-2 border-e-2 border-transparent text-sm whitespace-nowrap text-gray-500 hover:text-blue-600 focus:outline-none focus:text-blue-600 disabled:opacity-50 disabled:pointer-events-none active" id="vertical-tab-with-border-item-1" data-hs-tab="#vertical-tab-with-border-1" aria-controls="vertical-tab-with-border-1" role="tab">
						        		Identitas Pegawai
						      		</button>
						      		<button type="button" class="hs-tab-active:border-blue-500 hs-tab-active:text-blue-600 hs-tab-active:bg-blue-100 p-2 rounded-e-lg pe-4 inline-flex items-center gap-x-2 border-e-2 border-transparent text-sm whitespace-nowrap text-gray-500 hover:text-blue-600 focus:outline-none focus:text-blue-600 disabled:opacity-50 disabled:pointer-events-none" id="vertical-tab-with-border-item-2" data-hs-tab="#vertical-tab-with-border-2" aria-controls="vertical-tab-with-border-2" role="tab">
						        		CPNS
						      		</button>
						      		<button type="button" class="hs-tab-active:border-blue-500 hs-tab-active:text-blue-600 hs-tab-active:bg-blue-100 p-2 rounded-e-lg pe-4 inline-flex items-center gap-x-2 border-e-2 border-transparent text-sm whitespace-nowrap text-gray-500 hover:text-blue-600 focus:outline-none focus:text-blue-600 disabled:opacity-50 disabled:pointer-events-none" id="vertical-tab-with-border-item-3" data-hs-tab="#vertical-tab-with-border-3" aria-controls="vertical-tab-with-border-3" role="tab">
						        		PNS
						      		</button>
						      		<button type="button" class="hs-tab-active:border-blue-500 hs-tab-active:text-blue-600 hs-tab-active:bg-blue-100 p-2 rounded-e-lg pe-4 inline-flex items-center gap-x-2 border-e-2 border-transparent text-sm whitespace-nowrap text-gray-500 hover:text-blue-600 focus:outline-none focus:text-blue-600 disabled:opacity-50 disabled:pointer-events-none" id="vertical-tab-with-border-item-4" data-hs-tab="#vertical-tab-with-border-4" aria-controls="vertical-tab-with-border-4" role="tab">
						        		Pangkat Terakhir
						      		</button>
						      		<button type="button" class="hs-tab-active:border-blue-500 hs-tab-active:text-blue-600 hs-tab-active:bg-blue-100 p-2 rounded-e-lg pe-4 inline-flex items-center gap-x-2 border-e-2 border-transparent text-sm whitespace-nowrap text-gray-500 hover:text-blue-600 focus:outline-none focus:text-blue-600 disabled:opacity-50 disabled:pointer-events-none" id="vertical-tab-with-border-item-5" data-hs-tab="#vertical-tab-with-border-5" aria-controls="vertical-tab-with-border-5" role="tab">
						        		Jabatan Terakhir
						      		</button>
						      		<button type="button" class="hs-tab-active:border-blue-500 hs-tab-active:text-blue-600 hs-tab-active:bg-blue-100 p-2 rounded-e-lg pe-4 inline-flex items-center gap-x-2 border-e-2 border-transparent text-sm whitespace-nowrap text-gray-500 hover:text-blue-600 focus:outline-none focus:text-blue-600 disabled:opacity-50 disabled:pointer-events-none" id="vertical-tab-with-border-item-6" data-hs-tab="#vertical-tab-with-border-6" aria-controls="vertical-tab-with-border-6" role="tab">
						        		Tugas Tambahan
						      		</button>
						      		<button type="button" class="hs-tab-active:border-blue-500 hs-tab-active:text-blue-600 hs-tab-active:bg-blue-100 p-2 rounded-e-lg pe-4 inline-flex items-center gap-x-2 border-e-2 border-transparent text-sm whitespace-nowrap text-gray-500 hover:text-blue-600 focus:outline-none focus:text-blue-600 disabled:opacity-50 disabled:pointer-events-none" id="vertical-tab-with-border-item-7" data-hs-tab="#vertical-tab-with-border-7" aria-controls="vertical-tab-with-border-7" role="tab">
						        		Gaji Berkala Terakhir
						      		</button>

						      		<a type="button" class="hs-tab-active:border-blue-500 hs-tab-active:text-blue-600 hs-tab-active:bg-blue-100 p-2 rounded-e-lg pe-4 inline-flex items-center gap-x-2 border-e-2 border-transparent text-sm whitespace-nowrap text-gray-500 hover:text-blue-600 focus:outline-none focus:text-blue-600 disabled:opacity-50 disabled:pointer-events-none" id="vertical-tab-with-border-item-7" data-hs-tab="#vertical-tab-with-border-7" aria-controls="vertical-tab-with-border-7" role="tab">
						        		Gaji Berkala Terakhir
						      		</a>
						    	</nav>-->
						  	</div>
						</div>
					</div>
					<div :class="!disabled ? 'sm:col-span-10' : 'sm:col-span-12'">
						<slot />
					</div>
				</div>
			</div>
		</div>
	</div>
</template>