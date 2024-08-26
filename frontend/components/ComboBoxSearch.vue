<script setup>
const emit = defineEmits(['update:modelValue']);
const props = defineProps(['modelValue', 'options', 'keyList', 'namaList', 'id', 'multiple', 'disabled', 'apiUrl']);
const SearchString = ref(null)
const selected = ref('')

onMounted(() => {
	//HSComboBox.autoInit()
	const selectEl = document.querySelector('#'+props.id);
	var eli = new HSComboBox(selectEl, {
		apiUrl : props.apiUrl,
	})

	console.log(eli)

	var attributes = null

	eli.output.addEventListener("click", (val) => {
		//loadDataFromApi(eli, eli.search.value)		
		console.log(eli)
		//emit('update:modelValue', val)
	})
});


</script>
<template>
	<div :id="id" class="relative" data-hs-combo-box='{
		"apiQuery": "limit=10",
		"apiSearchQuery": "searchNama",
		"outputItemTemplate": "<div class=\"cursor-pointer py-2 px-4 w-full text-sm text-gray-800 hover:bg-gray-100 rounded-lg focus:outline-none focus:bg-gray-100\" data-hs-combo-box-output-item><div class=\"flex justify-between items-center bg-white w-full\"><div><div data-hs-combo-box-output-item-field=\"nama\" data-hs-combo-box-search-text data-hs-combo-box-value></div></div><span class=\"hidden hs-combo-box-selected:block\"><svg class=\"flex-shrink-0 size-3.5 text-blue-600\" xmlns=\"http:.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><polyline points=\"20 6 9 17 4 12\"></polyline></svg></span></div></div>",
		"outputEmptyTemplate": "<div class=\"py-2 px-4 w-full text-sm text-gray-800 bg-white rounded-lg\">Nothing found...</div>",
		"outputLoaderTemplate": "<div class=\"flex justify-center items-center py-2 px-4 text-sm text-gray-800 rounded-lg bg-white\"><div class=\"animate-spin inline-block size-6 border-[3px] border-current border-t-transparent text-blue-600 rounded-full bg-white\" role=\"status\" aria-label=\"loading\"><span class=\"sr-only\">Loading...</span></div></div>"
	}'>
		<div class="relative">
			<input class="py-3 px-4 bg-white block w-full border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none" type="text" v-bind:value="SearchString" data-hs-combo-box-input>
			<div class="absolute top-1/2 end-3 -translate-y-1/2" data-hs-combo-box-toggle="">
				<svg class="flex-shrink-0 size-3.5 text-gray-500" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
					<path d="m7 15 5 5 5-5"></path>
					<path d="m7 9 5-5 5 5"></path>
				</svg>
			</div>
		</div>
		<div class="absolute z-50 w-full max-h-72 p-1 bg-white border border-gray-200 rounded-lg overflow-hidden overflow-y-auto [&::-webkit-scrollbar]:w-2 [&::-webkit-scrollbar-thumb]:rounded-full [&::-webkit-scrollbar-track]:bg-gray-100 [&::-webkit-scrollbar-thumb]:bg-gray-300" style="display: none;" data-hs-combo-box-output=""></div>
	</div>
</template>