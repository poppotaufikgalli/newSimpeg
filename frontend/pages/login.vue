<script setup>
import { storeToRefs } from 'pinia'; // import storeToRefs helper hook from pinia
import { useAuthStore } from '~/store/auth'; // import the auth store we just created

const { authenticateUser, fakeauth, authenticateFakeUser } = useAuthStore(); // use authenticateUser action from  auth store
const { authenticated, message } = storeToRefs(useAuthStore()); // make authenticated state reactive with storeToRefs
const route = useRoute();

definePageMeta({
	layout: 'login'
})

const toast = useToast()
const showSpinner = ref(false)

const state = ref({
	username: '',
	password: '',
});

const router = useRouter();

const login = async() => {
	showSpinner.value = true
	await authenticateUser(state.value); // call authenticateUser and pass the user object
	showSpinner.value = false
	if (authenticated.value) {
		router.push('/');
	}else{
		errorToast(message)
	}
}

const errorToast = (message) => {
	toast.add({
		id: 'error_login',
		description: message,
		timeout: 6000
	}) 
}
</script>
<template>
	<AppLoadingSpinner :show="showSpinner" />
	<div class="flex h-screen bg-gray-200">
		<main class="w-full max-w-md m-auto">
			<div class="bg-white border border-gray-200 rounded-xl shadow-sm">
				<div class="p-4 sm:p-7">
					<div class="flex justify-center items-center">
						<AppLogo />
					</div>
				
					<div class="mt-5">
						<!-- Form -->
						<form @submit.prevent="login">
							<div class="grid gap-y-4">
								<!-- Form Group -->
								<div>
									<label for="username" class="block text-sm mb-1">Username/NIP</label>
									<input type="text" id="username" v-model="state.username" class="py-3 px-4 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white shadow" required>
								</div>
								<!-- End Form Group -->

								<!-- Form Group -->
								<div class="max-w-sm">
									<label class="block text-sm mb-2">Password</label>
									<div class="relative">
										<input id="hs-toggle-password" type="password" v-model="state.password" class="py-3 ps-4 pe-10 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white shadow" placeholder="Enter password"  required>
										<button type="button" data-hs-toggle-password='{
											"target": "#hs-toggle-password"
										}' class="absolute inset-y-0 end-0 flex items-center z-20 px-3 cursor-pointer text-gray-400 rounded-e-md focus:outline-none focus:text-blue-600">
											<svg class="shrink-0 size-3.5" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
												<path class="hs-password-active:hidden" d="M9.88 9.88a3 3 0 1 0 4.24 4.24"></path>
												<path class="hs-password-active:hidden" d="M10.73 5.08A10.43 10.43 0 0 1 12 5c7 0 10 7 10 7a13.16 13.16 0 0 1-1.67 2.68"></path>
												<path class="hs-password-active:hidden" d="M6.61 6.61A13.526 13.526 0 0 0 2 12s3 7 10 7a9.74 9.74 0 0 0 5.39-1.61"></path>
												<line class="hs-password-active:hidden" x1="2" x2="22" y1="2" y2="22"></line>
												<path class="hidden hs-password-active:block" d="M2 12s3-7 10-7 10 7 10 7-3 7-10 7-10-7-10-7Z"></path>
												<circle class="hidden hs-password-active:block" cx="12" cy="12" r="3"></circle>
									  		</svg>
										</button>
									</div>
								</div>
								<!-- End Form Group -->
								<!-- End Form Group -->
								<div class="mt-5">
									<button type="submit" class="w-full py-3 px-4 inline-flex justify-center items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none">Log in</button>
								</div>
							</div>
						</form>
						<!-- End Form -->
						<div class="mt-2">
							<button class="w-full py-3 px-4 inline-flex justify-center items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-green-600 text-white hover:bg-green-700 disabled:opacity-50 disabled:pointer-events-none" @click="SSO()">Log in SSO</button>
						</div>
					</div>
				</div>
			</div>
		</main>
	</div>
</template>