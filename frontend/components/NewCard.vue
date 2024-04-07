<script setup>
	const nuxtApp = useNuxtApp();
	const loading = ref(false);
	nuxtApp.hook("page:start", () => {
	    loading.value = true;
	});
	nuxtApp.hook("page:finish", () => {
	    loading.value = false;
	});
	/*const props = defineProps({
		show: {
      		type: Boolean,
      		default: false,
    	},
	})

	const loadingSpinner = ref(null);

	watch(
    	() => props.show,
    	() => {
  			console.log(props.show)
      		if (props.show) {
        		HSOverlay.open('#loadingSpinner')
      		} else {
        		HSOverlay.close('#loadingSpinner')
      		}
    	}
	);*/
</script>
<template>
	<div class="relative" id="loadingSpinner">
		<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
			<slot/>
		</div>  
		<template v-if="loading">
		  	<div class="absolute top-0 start-0 size-full bg-white/[.5] rounded-lg dark:bg-gray-800/[.4]"></div>

		  	<div class="absolute top-1/2 start-1/2 transform -translate-x-1/2 -translate-y-1/2">
		    	<div class="animate-spin inline-block size-6 border-[3px] border-current border-t-transparent text-blue-600 rounded-full dark:text-blue-500" role="status" aria-label="loading">
		      		<span class="sr-only">Loading...</span>
		    	</div>
		  	</div>
		</template>
	</div>
	
</template>