<script setup>
const route = useRoute()
const {snip, skgolru, stmtpang, sknpang} = route.params
</script>
<template>
	<LayoutDataInduk>
		<FormPangkat :snip="snip" :skgolru="skgolru" :stmtpang="stmtpang" :sknpang="sknpang"/>
	</LayoutDataInduk>
</template>