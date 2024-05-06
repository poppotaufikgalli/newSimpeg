<script setup>
	//const { element } = HSOverlay.getInstance('#loadingSpinner', true);
	//const modal = new HSOverlay(document.querySelector('#loadingSpinner'));
	const props = defineProps({
		show: {
      type: Boolean,
      default: false,
    },
	})

	const loadingSpinner = ref(null);
	//const a = ref(null)

	watch(
    () => props.show,
    () => {
  		console.log(props.show)
      if (props.show) {
        HSOverlay.open('#loadingSpinner')
      		//element.open();
      	//a.value.open()
      } else {
        HSOverlay.close('#loadingSpinner')
      	//element.close();
      }
    }
	);
</script>

<template>
	<button type="button" class="hidden" data-hs-overlay="#loadingSpinner">
  	Open modal
	</button>

	<!--<div id="loadingSpinner" ref="loadingSpinner" class="hs-overlay size-full fixed top-0 start-0 z-[80] overflow-x-hidden overflow-y-auto pointer-events-none">
	  <div class="hs-overlay-open:mt-7 hs-overlay-open:opacity-100 hs-overlay-open:duration-500 mt-0 opacity-0 ease-out transition-all sm:max-w-lg sm:w-full m-3 sm:mx-auto min-h-[calc(100%-3.5rem)] flex items-center">
	    <div class="w-full flex flex-col pointer-events-auto">
	      <div class="p-4 overflow-y-auto mx-auto">
	        <div class="animate-spin inline-block size-8 border-[3px] border-current border-t-transparent text-blue-600 rounded-full dark:text-blue-500" role="status" aria-label="loading">
					  <span class="sr-only">Loading...</span>
					</div>
	      </div>
	    </div>
	  </div>
	</div>-->
	<div id="loadingSpinner" ref="loadingSpinner" class="hs-overlay hidden size-full fixed top-0 start-0 z-[60] overflow-x-hidden overflow-y-auto pointer-events-none">
	  <div class="hs-overlay-open:opacity-100 hs-overlay-open:duration-500 opacity-0 transition-all sm:max-w-lg sm:w-full m-3 sm:mx-auto">
	    <div class="pointer-events-auto">&nbsp;</div>
	  </div>
	</div>
</template>