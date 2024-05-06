<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';
const modalUploadDoc = useModalUploadDoc();
const fileInput = ref(null)

const { displayModal, modalTitle, modalFile, modalFileType, fileBlob } = storeToRefs(modalUploadDoc)

watch(displayModal, (val) => {
	if(val){
		HSOverlay.open('#modalUploadDoc')
	}else{
		// /fileInput.value = null
		HSOverlay.close('#modalUploadDoc')
	}
})

const handleFileChange = (e) => {
	var files = e.target.files || e.dataTransfer.files;
    if (!files.length)
        return;
    fileBlob.value = files
}

const doAction = () => {
	if(fileInput.value.files.length > 0){
		modalUploadDoc.doAction()	
	}
}

</script>
<template>
	<div id="modalUploadDoc" class="hs-overlay hidden size-full fixed top-0 start-0 z-[80] overflow-x-hidden overflow-y-auto">
		<div class="hs-overlay-open:mt-7 hs-overlay-open:opacity-100 hs-overlay-open:duration-500 mt-0 opacity-0 ease-out transition-all md:max-w-4xl md:w-full m-3 md:mx-auto">
			<div class="relative flex flex-col bg-white border shadow-sm rounded-xl overflow-hidden">
				<div class="absolute top-2 end-2">
					<button type="button" class="flex justify-center items-center size-7 text-sm font-semibold rounded-lg border border-transparent text-gray-800 hover:bg-gray-100 disabled:opacity-50 disabled:pointer-events-none" data-hs-overlay="#modalUploadDoc" @click="modalUploadDoc.closeModal()">
						<span class="sr-only">Close</span>
						<svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
					</button>
				</div>

				<div class="p-4 sm:p-10 overflow-y-auto">
					<div class="flex gap-x-4 md:gap-x-7">
						<div class="grow">
							<h3 class="mb-2 text-xl font-bold text-gray-800">
								{{modalTitle}}
							</h3>
							<div class="max-w">
							  	<div class="flex justify-between">
							    	<label class="block">
							      		<span class="sr-only">Choose profile photo</span>
							      		<input ref="fileInput" type="file" class="block w-full text-sm text-gray-500 file:me-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:text-sm file:font-semibold file:bg-blue-600 file:text-white hover:file:bg-blue-700 file:disabled:opacity-50 file:disabled:pointer-events-none" :accept="modalFileType" v-on:change="handleFileChange($event)">
							    	</label>
							    	<button class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-green-500 text-white hover:bg-green-600 disabled:opacity-50 disabled:pointer-events-none" @click="doAction">
										Upload
									</button>
							  	</div>
							</div>
							<div class="flex justify-center mt-2" v-if="displayModal">
								<!--<img :src="modalFile" class="w-full rounded">-->
								<embed :src="modalFile" class="w-full h-screen" />
							</div>
						</div>
					</div>
				</div>

				<div class="flex justify-end items-center gap-x-2 py-3 px-4 bg-gray-50 border-t">
					<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none" data-hs-overlay="#modalUploadDoc" @click="modalUploadDoc.closeModal()">
						Tutup
					</button>
				</div>
			</div>
		</div>
	</div>
</template>