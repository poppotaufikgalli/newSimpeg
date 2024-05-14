<script setup>
const dayjs = useDayjs()

const { pending, data, refresh} = await useFetch('/api/gets/singkronisasi/bkn');
</script>
<template>
	<div class="mx-auto mb-6">
		<!-- Card -->
		<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
			<!-- Grid -->
			<div class="grid sm:grid-cols-12 gap-2 gap-2.5" v-if="!pending">
				<div class="sm:col-span-2">
					<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">Token SSO Expired</label>
				</div>
				<!-- End Col -->

				<div class="sm:col-span-3">
					<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 disabled:opacity-50 disabled:pointer-events-none bg-white" :value="data.token_sso_expired" v-if="data" :class="$dayjs().isBefore($dayjs(data.token_sso_expired)) ? 'border-gray-200 focus:border-blue-500 focus:ring-blue-500' : 'border-red-500 focus:border-red-500 focus:ring-red-500'" disabled>
				</div>
				<!-- End Col -->

				<div class="sm:col-span-2 sm:col-start-8">
					<label for="af-account-full-name" class="inline-block text-sm text-gray-800 mt-2.5">API Manager Expired</label>
				</div>
				<!-- End Col -->

				<div class="sm:col-span-3">
					<input id="af-account-full-name" type="text" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" :value="data.token_apimanager_expired" v-if="data" :class="$dayjs().isBefore($dayjs(data.token_apimanager_expired)) ? 'border-gray-200 focus:border-blue-500 focus:ring-blue-500' : 'border-red-500 focus:border-red-500 focus:ring-red-500'" disabled>
				</div>
				<!-- End Col -->
			</div>
		</div>
	</div>
</template>