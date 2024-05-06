<script setup>
import { useModalUpload } from '~/store/modalUpload';
const modalUpload = useModalUpload();
const { fileBlob } = storeToRefs(modalUpload)
const file = ref(null)

const openModal = () => {
	modalUpload.showModal("Upload Gambar", "image/jpg,image/png")
}

onMounted(() => {
	modalUpload.$onAction(callback, true)
})

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		let formData = new FormData();

		//append your file or image
		formData.append("file", fileBlob.value[0]);
		formData.append("nip", '123');
		formData.append("path", 'Gambar');

		await useFetch('/api/uploads/pegawai/upload', {
			method: 'POST',
			body: formData,
		}).catch((error) => {
	        //assign response error data to state "errors"
	        errors.value = error.data
	    });
	}
}

</script>
<template>
	<p>
		<button v-on:click="openModal">test</button>
		<img :src="file" />
	</p>
</template>