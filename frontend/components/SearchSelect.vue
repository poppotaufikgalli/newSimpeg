<script setup>
const emit = defineEmits(['update:modelValue']);
const props = defineProps(['modelValue', 'dataList', 'keyList', 'namaList', 'id0']);

const caption = ref(null)
const search = ref(null)

onMounted(() => {
	//console.log("onMounted")
	var mv = props.modelValue
	var list = props.dataList

	var item = _find(list, function(o){
		return o[props.keyList] == mv
	})

	if(item){
		caption.value = item[props.namaList]	
	}
})


watch(() => props.modelValue, (baru, lama) => {
	//console.log("watch")
  	//console.log(`Watch ${props.id0} function called with args: dari ${lama} ke ${baru}`);
  	if(baru == null){
  		caption.value = null
  		//search.value = null
  	}else{
  		var item = _find(props.dataList, function(o){
  			return o[props.keyList] == baru
  		})
  		if(item){
  			caption.value = item[props.namaList]
  			//search.value = item[props.namaList]	
  		}
  	}
});

function setValue(val){
	//console.log("setValue")
	var id = val ? val[props.keyList] : null;
	var nama = val ? val[props.namaList] : null;

	caption.value = nama
	search.value = null
	emit('update:modelValue', id)
	var ids2 = "#"+props.id0
	HSDropdown.close(ids2);
}

const proxyData = computed({
	get(){
		//console.log("proxyData")
		const {dataList} = props
		if(dataList){
			return dataList.filter((dt) => {
				if(search.value != null){
					return dt[props.namaList].toLowerCase().includes(search.value.toLowerCase())
				}
				return dt
			});
		}
		return null
	}
})
</script>
<template>
	<div :id="id0" class="hs-dropdown w-full relative sm:col-span-11 sm:mt-1 sm:inline-flex [--auto-close:inside]">
		<button id="hs-dropdown-11" type="button" class="hs-dropdown-toggle py-2 px-4 inline-flex items-center justify-between gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none w-[100%]">
			<span class="text-start truncate">{{caption || "Semua"}}</span>
			<svg class="hs-dropdown-open:rotate-180 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m6 9 6 6 6-6"/></svg>
		</button>
		<div class="hs-dropdown-menu transition-[opacity,margin] duration hs-dropdown-open:opacity-100 opacity-0 z-10 bg-gray-400 shadow-md rounded-lg p-1 hidden" aria-labelledby="hs-dropdown-11">
			<input type="search" class="p-2 px-4 mb-2 block w-full min-w-[200px] border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" placeholder="Cari Data Unor" v-model="search">
			<div class="max-h-[400px] min-w-screen/2 overflow-y-auto [&::-webkit-scrollbar]:w-2 [&::-webkit-scrollbar-track]:bg-gray-100 [&::-webkit-scrollbar-thumb]:bg-gray-300">
				<template v-if="proxyData" v-for="(item, idx) in proxyData" :key="item[keyList]">
					<span class="flex items-center gap-x-3.5 py-2 px-3 rounded-lg text-xs text-gray-800 hover:bg-gray-100 focus:outline-none focus:bg-gray-100" @click="setValue(item)">
						{{item[keyList]}} - {{item[namaList]}}
					</span> 
				</template>
				<template v-else>
					<span class="flex justify-center items-center gap-x-3.5 py-2 px-3 rounded-lg text-xs text-gray-800 focus:outline-none focus:bg-gray-100">
						<i>Data Tidak / Belum Tersedia</i>
					</span>
				</template>
			</div>
		</div>
	</div>
</template>
<style scoped>
	[type="search"]::-webkit-search-cancel-button {
  		-webkit-appearance: none;
  		appearance: none;
  		height: 10px;
  		width: 10px;
  		background-image: url(data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iaXNvLTg4NTktMSI/Pg0KPCEtLSBHZW5lcmF0b3I6IEFkb2JlIElsbHVzdHJhdG9yIDE2LjAuMCwgU1ZHIEV4cG9ydCBQbHVnLUluIC4gU1ZHIFZlcnNpb246IDYuMDAgQnVpbGQgMCkgIC0tPg0KPCFET0NUWVBFIHN2ZyBQVUJMSUMgIi0vL1czQy8vRFREIFNWRyAxLjEvL0VOIiAiaHR0cDovL3d3dy53My5vcmcvR3JhcGhpY3MvU1ZHLzEuMS9EVEQvc3ZnMTEuZHRkIj4NCjxzdmcgdmVyc2lvbj0iMS4xIiBpZD0iQ2FwYV8xIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB4PSIwcHgiIHk9IjBweCINCgkgd2lkdGg9IjEyMy4wNXB4IiBoZWlnaHQ9IjEyMy4wNXB4IiB2aWV3Qm94PSIwIDAgMTIzLjA1IDEyMy4wNSIgc3R5bGU9ImVuYWJsZS1iYWNrZ3JvdW5kOm5ldyAwIDAgMTIzLjA1IDEyMy4wNTsiDQoJIHhtbDpzcGFjZT0icHJlc2VydmUiPg0KPGc+DQoJPHBhdGggZD0iTTEyMS4zMjUsMTAuOTI1bC04LjUtOC4zOTljLTIuMy0yLjMtNi4xLTIuMy04LjUsMGwtNDIuNCw0Mi4zOTlMMTguNzI2LDEuNzI2Yy0yLjMwMS0yLjMwMS02LjEwMS0yLjMwMS04LjUsMGwtOC41LDguNQ0KCQljLTIuMzAxLDIuMy0yLjMwMSw2LjEsMCw4LjVsNDMuMSw0My4xbC00Mi4zLDQyLjVjLTIuMywyLjMtMi4zLDYuMSwwLDguNWw4LjUsOC41YzIuMywyLjMsNi4xLDIuMyw4LjUsMGw0Mi4zOTktNDIuNGw0Mi40LDQyLjQNCgkJYzIuMywyLjMsNi4xLDIuMyw4LjUsMGw4LjUtOC41YzIuMy0yLjMsMi4zLTYuMSwwLTguNWwtNDIuNS00Mi40bDQyLjQtNDIuMzk5QzEyMy42MjUsMTcuMTI1LDEyMy42MjUsMTMuMzI1LDEyMS4zMjUsMTAuOTI1eiIvPg0KPC9nPg0KPGc+DQo8L2c+DQo8Zz4NCjwvZz4NCjxnPg0KPC9nPg0KPGc+DQo8L2c+DQo8Zz4NCjwvZz4NCjxnPg0KPC9nPg0KPGc+DQo8L2c+DQo8Zz4NCjwvZz4NCjxnPg0KPC9nPg0KPGc+DQo8L2c+DQo8Zz4NCjwvZz4NCjxnPg0KPC9nPg0KPGc+DQo8L2c+DQo8Zz4NCjwvZz4NCjxnPg0KPC9nPg0KPC9zdmc+DQo=);
  		background-size: 10px 10px;
	}
</style>
