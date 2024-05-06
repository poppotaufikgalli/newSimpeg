<script setup>
import { storeToRefs } from 'pinia'; // import storeToRefs helper hook from pinia
import { useAuthStore } from '~/store/auth'; // import the auth store we just created

const { authenticateUser, fakeauth, authenticateFakeUser } = useAuthStore(); // use authenticateUser action from  auth store
const { authenticated, message } = storeToRefs(useAuthStore()); // make authenticated state reactive with storeToRefs

definePageMeta({
  	layout: 'login'
})

const toast = useToast()
const showSpinner = ref(false)


const state = ref({
  username: '',
  password: '',
});

const validate = (state) => {
  const errors = []
  if (!state.username) errors.push({ path: 'username', message: 'Required' })
  if (!state.password) errors.push({ path: 'password', message: 'Required' })
  return errors
}

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

const fakeLogin = async() => {
	showSpinner.value = true
  await authenticateFakeUser(state.value); // call authenticateUser and pass the user object
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

          <div class="mt-5 text-center">
            <h1 class="block text-2xl font-bold text-gray-800">Form Login</h1>
          </div>

          <div class="mt-5">
            <!-- Form -->
            <form @submit.prevent="login">
              <div class="grid gap-y-4">
                <!-- Form Group -->
                <div>
                  <label for="username" class="block text-sm mb-1">Username/NIP</label>
                  <div class="relative">
                    <input type="text" id="username" v-model="state.username" class="py-3 px-4 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white shadow" required aria-describedby="email-error">
                    <div class="hidden absolute inset-y-0 end-0 pointer-events-none pe-3">
                      <svg class="size-5 text-red-500" width="16" height="16" fill="currentColor" viewBox="0 0 16 16" aria-hidden="true">
                        <path d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0zM8 4a.905.905 0 0 0-.9.995l.35 3.507a.552.552 0 0 0 1.1 0l.35-3.507A.905.905 0 0 0 8 4zm.002 6a1 1 0 1 0 0 2 1 1 0 0 0 0-2z"/>
                      </svg>
                    </div>
                  </div>
                  <p class="hidden text-xs text-red-600 mt-2" id="email-error">Please include a valid email address so we can get back to you</p>
                </div>
                <!-- End Form Group -->

                <!-- Form Group -->
                <div>
                  <label for="password" class="block text-sm mb-1">Password</label>
                  <div class="relative">
                    <input type="password" id="password" v-model="state.password" class="py-3 px-4 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white shadow" required aria-describedby="password-error">
                    <div class="hidden absolute inset-y-0 end-0 pointer-events-none pe-3">
                      <svg class="size-5 text-red-500" width="16" height="16" fill="currentColor" viewBox="0 0 16 16" aria-hidden="true">
                        <path d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0zM8 4a.905.905 0 0 0-.9.995l.35 3.507a.552.552 0 0 0 1.1 0l.35-3.507A.905.905 0 0 0 8 4zm.002 6a1 1 0 1 0 0 2 1 1 0 0 0 0-2z"/>
                      </svg>
                    </div>
                  </div>
                  <p class="hidden text-xs text-red-600 mt-2" id="password-error">8+ characters required</p>
                </div>
                <!-- End Form Group -->
                <div class="mt-5">
                  <button type="submit" class="w-full py-3 px-4 inline-flex justify-center items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none">Log in</button>
                </div>
              </div>
            </form>
            <!-- End Form -->
          </div>
        </div>
      </div>
    </main>
  </div>
</template>