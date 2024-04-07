<script setup>
	const showSpinner = ref(false)

	const sh = () => {
		showSpinner.value = true
	}
</script>
<template>
	<AppLoadingSpinner :show="showSpinner" />
	<button @click="sh()">test</button>
</template>