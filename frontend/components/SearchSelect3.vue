<script setup>
const emit = defineEmits(['update:modelValue']);
const props = defineProps(['modelValue', 'options', 'keyList', 'namaList', 'id', 'multiple', 'disabled', "apiUrl"]);
const searchString = ref('')

onMounted(async() => {
	HSStaticMethods.autoInit();

	const selectEl = document.querySelector('#'+props.id);
	selectEl.value = props.modelValue
	const eli = HSSelect.getInstance(selectEl);
	//eli.value = props.modelValue

	if(props.multiple == 'multiple'){
		Object.values(props.modelValue).map(val => {
			eli.value.push(val.toString())

			const i = _findKey(eli.el.options, function(o){
				return o.value == val
			})

			eli.el.options[i].selected = true
		})
		reinit()
	}else{
		console.log("onMounted")
		await loadDataFromApi(eli, props.modelValue)
		//console.log(eli, props.modelValue)
		if(eli){
			eli.el.value = props.modelValue	
		}

		reinit()
	}

	//console.log(eli)
});

async function loadDataFromApi(select, searchNama){
	if(searchNama != undefined){
		var data = await $fetch(props.apiUrl+"?searchNama="+searchNama, {
			method: 'GET',
		})
		var title = props.namaList ?? "nama"
		var val = props.keyList ?? "id"

		if(data.length > 0){
			var result = data.map(function(o) {
			  	return _extend({}, o, {val: o[val], title: o[title]})
			});

			if(result){
				select.addOption(result)
			}
		}
	}
}

const optionsInit = computed({
	get(){
		return _slice(props.options, 0, 200)
	}
})

const reinit = () => {
	//console.log(props.options)
	HSStaticMethods.autoInit();
	const el = HSSelect.getInstance('#'+props.id);

	if(el){
		el.destroy();	

		HSStaticMethods.autoInit();
		const selectEl = document.querySelector('#'+props.id);
		
		setTimeout(() => {
			var eli = new HSSelect(selectEl)

			eli.on('change', (val) => {
				emit('update:modelValue', val)
				//console.log(eli)
			})

			if(props.apiUrl != ""){
				eli.search.addEventListener("keypress", (val) => {
					loadDataFromApi(eli, eli.search.value)		
				})
			}
		}, 100);
	}
}

const readOptions = (val) => {
	const eli = HSSelect.getInstance('#'+props.id);
	return eli.el.selectedOptions[0].text
}

defineExpose({reinit, readOptions});

function truncating(text){
	if(text.length > 5){
		return text.slice(0,5)+'...'
	}
	return text
}

</script>
<template>
	<div>
		<template v-if="multiple == 'multiple' ">
			<select multiple="" :id="id" data-hs-select='{
				  "placeholder": "Pilih Beberapa Data...",
				  "toggleTag": "<button type=\"button\"></button>",
				  "toggleClasses": "hs-select-disabled:pointer-events-none hs-select-disabled:opacity-50 relative py-2 px-3 pe-9 flex text-nowrap w-full cursor-pointer bg-white border border-gray-200 rounded-lg text-start text-sm focus:border-blue-500 focus:ring-blue-500 before:absolute before:inset-0 before:z-[1]",
				  "dropdownClasses": "mt-2 z-50 w-full max-h-72 p-1 space-y-0.5 bg-white border border-gray-200 rounded-lg overflow-hidden overflow-y-auto [&::-webkit-scrollbar]:w-2 [&::-webkit-scrollbar-thumb]:rounded-full [&::-webkit-scrollbar-track]:bg-gray-100 [&::-webkit-scrollbar-thumb]:bg-gray-300",
				  "optionClasses": "py-2 px-4 w-full text-sm text-gray-800 cursor-pointer hover:bg-gray-100 rounded-lg focus:outline-none focus:bg-gray-100",
				  "optionTemplate": "<div class=\"flex justify-between items-center w-full\"><span data-title></span><span data-description></span><span class=\"hidden hs-selected:block\"><svg class=\"flex-shrink-0 size-3.5 text-blue-600\" xmlns=\"http:.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><polyline points=\"20 6 9 17 4 12\"/></svg></span></div>",
				  "extraMarkup": "<div class=\"absolute top-1/2 end-3 -translate-y-1/2\"><svg class=\"flex-shrink-0 size-3.5 text-gray-500\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"m7 15 5 5 5-5\"/><path d=\"m7 9 5-5 5 5\"/></svg></div>"
				}' class="hidden" :value="modelValue" :disabled="disabled">
			</select>
		</template>
		<template v-else>
			<select :id="id" data-hs-select='{
		  		"hasSearch": true,
		  		"searchPlaceholder": "Cari...",
		  		"searchClasses": "block w-full text-sm bg-white border border-gray-200 rounded-lg focus:border-blue-500 focus:ring-blue-500 before:absolute before:inset-0 before:z-[1] py-2 px-3",
		  		"searchWrapperClasses": "bg-white p-2 -mx-1 sticky top-0",
		  		"placeholder": "Pilih Data...",
		  		"toggleTag": "<button type=\"button\"><span class=\"me-2\" data-icon></span><span class=\"text-gray-800\" data-title></span></button>",
		  		"toggleClasses": "hs-select-disabled:pointer-events-none hs-select-disabled:opacity-50 relative py-2 px-3 pe-9 flex text-nowrap w-full cursor-pointer bg-white border border-gray-200 rounded-lg text-start text-sm focus:border-blue-500 focus:ring-blue-500 before:absolute before:inset-0 before:z-[1]",
		  		"dropdownClasses": "mt-2 max-h-72 pb-1 px-1 space-y-0.5 z-20 w-full bg-white border border-gray-200 rounded-lg overflow-hidden overflow-y-auto [&::-webkit-scrollbar]:w-2 [&::-webkit-scrollbar-thumb]:rounded-full [&::-webkit-scrollbar-track]:bg-gray-100 [&::-webkit-scrollbar-thumb]:bg-gray-300",
		  		"optionClasses": "py-2 px-4 w-full text-sm text-gray-800 cursor-pointer hover:bg-gray-100 rounded-lg focus:outline-none focus:bg-gray-100",
		  		"optionTemplate": "<div><div class=\"flex items-center\"><div class=\"me-2\" data-icon></div><div class=\"text-gray-800\" data-title></div></div><div class=\"me-2\"><i data-description></i></div></div>",
		  		"extraMarkup": "<div class=\"absolute top-1/2 end-3 -translate-y-1/2\"><svg class=\"flex-shrink-0 size-3.5 text-gray-500\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"m7 15 5 5 5-5\"/><path d=\"m7 9 5-5 5 5\"/></svg></div>"
				}' class="hidden" :value="modelValue" :disabled="disabled">
		  			<option value="">Choose</option>
			</select>
		</template>
	</div>
</template>
