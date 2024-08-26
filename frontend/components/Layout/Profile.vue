<script setup>
	import { useModalUpload } from '~/store/modalUpload';
	const modalUpload = useModalUpload();
	const { fileBlob } = storeToRefs(modalUpload)

	const { $decodeBase64, $encodeBase64 } = useNuxtApp()
	const route = useRoute()
	const {snip} = route.params

	const photo = ref("https://upload.wikimedia.org/wikipedia/commons/6/65/No-Image-Placeholder.svg")

	const dataInduk = ref({
		nip: '',
		nama: '',
		file_bmp: '',
	})

	const { pending, data, refresh} = await useAsyncData('getDataIdentitasPhoto', async() => {
		console.log("getDataIdentitasPhoto")
		if(snip){
			let nip = $decodeBase64(snip)
			var result = await $fetch('/api/gets/pegawai/'+nip);

			if(result.nip == nip){
				dataInduk.value = _pickBy(result, function(val, key) {
					return _includes(_keys(dataInduk.value), key);
				})

				if(result.file_bmp !== ""){
					var file_bmp = await $fetch(`/api/fileserver/static/photo/${result.file_bmp}`)

					const blobUrl = URL.createObjectURL(file_bmp)
					photo.value = blobUrl
				}
			}
		}
	})
	
	const disabled = computed({
		get(){
			if(snip == 'null'){
				return true
			}
			return false
		}
	})

	const showProfilePic = () => {
		modalUpload.showModal("Upload Photo", "image/jpg,image/png")
	}

	onMounted(() => {
		refreshNuxtData(["getDataIdentitasPhoto"])
		modalUpload.$onAction(callback, true)
	})

	const callback = async(e) => {
		const {name} = e

		if(name == 'doAction'){
			let nip = $decodeBase64(snip)
			let formData = new FormData();

			//append your file or image
			formData.append("file", fileBlob.value[0]);
			formData.append("filename", nip);
			formData.append("path", 'photo');

			await useFetch(`/api/uploads/upload/pegawai/${nip}`, {
				method: 'POST',
				body: formData,
			}).catch((error) => {
		        //assign response error data to state "errors"
		        errors.value = error.data
		    });

		    refreshNuxtData(["getDataIdentitasPhoto"])
		}
	}
</script>
<style scoped>
.router-link-active {
	@apply bg-blue-800 text-gray-200;
}
.router-link-active:hover {
	@apply font-medium;
}
</style>
<template>
	<div class="mx-auto">
		<!-- Card -->
		<div class="bg-white rounded-xl shadow p-4 mb-6">
			<TopBarPegawai />
			<div class="mt-5">
				<div class="grid sm:grid-cols-12 gap-4">
					<div class="sm:col-span-2" v-if="!disabled">
						<div class="flex flex-wrap">
							<div class="mb-5">
								<img :src="photo" @click="showProfilePic" :alt="dataInduk.file_bmp">
							</div>
							<div class="w-full">
								<nav class="-me-0.5 flex flex-col border divide-y rounded-md overflow-hidden">
									<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/DataInduk/identitas`">
										Identitas Pegawai
									</NuxtLink>
									<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/DataInduk/cpns`">
										CPNS 
									</NuxtLink>
									<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="`/pegawai/${snip}/DataInduk/pns`">
										PNS 
									</NuxtLink>
									<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="{ path: '/pegawai/'+snip+'/DataInduk/pangkat', query: { akhir: true}}" >
										Pangkat Terakhir
									</NuxtLink>
									<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="{ path: '/pegawai/'+snip+'/DataInduk/jabatan', query: { akhir: true}}" >
										Jabatan Terakhir 
									</NuxtLink>
									<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="{ path: '/pegawai/'+snip+'/DataInduk/tugas-tambahan', query: { akhir: true}}" >
										Tugas Tambahan
									</NuxtLink>
									<NuxtLink class="p-2 inline-flex items-center gap-2 text-sm font-medium text-gray hover:outline-none hover:text-blue-600 hover:bg-blue-200" :to="{ path: '/pegawai/'+snip+'/DataInduk/gaji-berkala', query: { akhir: true}}" >
										Gaji Berkala Terakhir
									</NuxtLink>
								</nav>
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