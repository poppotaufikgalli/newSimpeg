<script setup>
const pangkats = ref([])
const pejabats = ref([])
const { pending, data, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
  	const [pangkat, pejabat] = await Promise.all([
    	$fetch('/api/gets/pangkat'),
    	$fetch('/api/gets/pejabat')
  	]);

  	pangkats.value = pangkat
  	pejabats.value = pejabat
  	//return {pangkat, pejabat}

  	//agamas.value = agama
  	//statuss.value = status
  	//jeniss.value = jenis
  	//duduks.value = duduk
  	//kawins.value = kawin
  	//goldars.value = goldar
  	//provs.value = prov
  	//kpes.value = kpe
});

onMounted(() => {
	refreshNuxtData(["getDataRef"])

	//console.log(data)
})
</script>
<template>
	<div class="mx-auto">
		<!-- Card -->
		<div class="">
			<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
				<div class="sm:col-span-2">
					<div class="flex flex-wrap">
						<div class="mb-12">
			    			<img src="https://upload.wikimedia.org/wikipedia/commons/6/65/No-Image-Placeholder.svg">
			    		</div>
					  	<div class="w-full">
					    	<nav class="flex flex-col space-y-2" aria-label="Tabs" role="tablist" data-hs-tabs-vertical="true">
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
					    	</nav>
					  	</div>
					</div>
				</div>
				<div class="sm:col-span-10">
					<div class="ms-3">
				    	<div id="vertical-tab-with-border-1" role="tabpanel" aria-labelledby="vertical-tab-with-border-item-1">
				      		<p class="text-gray-500 dark:text-gray-400">
				        		<FormPegawaiDataIndukIdentitas />
				      		</p>
				    	</div>
				    	<div id="vertical-tab-with-border-2" class="hidden" role="tabpanel" aria-labelledby="vertical-tab-with-border-item-2">
				      		<p class="text-gray-500 dark:text-gray-400">
				        		<FormPegawaiDataIndukCpns :pangkats="pangkats" :pejabats="pejabats"/>
				      		</p>
				    	</div>
				    	<div id="vertical-tab-with-border-3" class="hidden" role="tabpanel" aria-labelledby="vertical-tab-with-border-item-3">
				      		<p class="text-gray-500 dark:text-gray-400">
				        		This is the <em class="font-semibold text-gray-800 dark:text-gray-200">third</em> item's tab body.
				      		</p>
				    	</div>
				  	</div>
				</div>
			</div>
		</div>
	</div>
</template>