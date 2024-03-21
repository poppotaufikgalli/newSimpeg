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
	<div id="loadingSpinner" ref="loadingSpinner" class="hs-overlay hidden size-full fixed top-0 start-0 z-[80] overflow-x-hidden overflow-y-auto">
	  <div class="hs-overlay-open:mt-7 hs-overlay-open:opacity-100 hs-overlay-open:duration-500 mt-0 opacity-0 ease-out transition-all sm:max-w-lg sm:w-full m-3 sm:mx-auto">
	    <div class="relative flex flex-col bg-white shadow-lg rounded-xl">
	      <div class="p-4 sm:p-10 text-center overflow-y-auto">
	        <!-- Icon -->
	        <span class="mb-4 inline-flex justify-center items-center size-[46px] rounded-full border-4 border-green-50 bg-green-100 text-green-500">
	          <svg class="flex-shrink-0 size-5" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
	            <path d="M11.251.068a.5.5 0 0 1 .227.58L9.677 6.5H13a.5.5 0 0 1 .364.843l-8 8.5a.5.5 0 0 1-.842-.49L6.323 9.5H3a.5.5 0 0 1-.364-.843l8-8.5a.5.5 0 0 1 .615-.09z"/>
	          </svg>
	        </span>
	        <!-- End Icon -->

	        <h3 class="mb-2 text-xl font-bold text-gray-800">
	          Mohon Tunggu Sebentar
	        </h3>
	      </div>
	    </div>
	  </div>
	</div>
</template>