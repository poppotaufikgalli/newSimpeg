<script setup lang="ts">
import { useAuthStore } from '~/store/auth'; // import the auth store we just created
const { authenticated } = storeToRefs(useAuthStore());
const { userInfo, logUserOut } = useAuthStore(); // use authenticateUser action from  auth store

const snip = ref('');

const items = [
  [{
    label: 'Profil',
    icon: 'i-heroicons-user-circle',
    to: '/profil'
  }], [{
    label: 'Keluar',
    icon: 'i-heroicons-arrow-left-on-rectangle',
    click: () => {
      logUserOut()      
    },
    to: "/login"
  }]
]
</script>
<template>
	<div>
		<!-- ========== HEADER ========== -->
		<header class="flex flex-wrap md:justify-start md:flex-nowrap z-50 w-full bg-white text-sm py-3 md:py-0">
		  <nav class="max-w-[85rem] w-full mx-auto px-4 md:px-6 lg:px-8" aria-label="Global">
		    <div class="relative md:flex md:items-center md:justify-between">
		      <div class="flex items-center justify-between">
		        <a class="inline-flex my-2" href="#" aria-label="SIMPEG">
		        	<AppLogo />
		        </a>
		        <div class="md:hidden">
		          <button type="button" class="hs-collapse-toggle flex justify-center items-center size-9 text-sm font-semibold rounded-lg border border-gray-200 text-gray-800 hover:bg-gray-100 disabled:opacity-50 disabled:pointer-events-none" data-hs-collapse="#navbar-collapse-with-animation" aria-controls="navbar-collapse-with-animation" aria-label="Toggle navigation">
		            <svg class="hs-collapse-open:hidden flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="3" x2="21" y1="6" y2="6"/><line x1="3" x2="21" y1="12" y2="12"/><line x1="3" x2="21" y1="18" y2="18"/></svg>
		            <svg class="hs-collapse-open:block hidden flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
		          </button>
		        </div>
		      </div>

		      <div id="navbar-collapse-with-animation" class="hs-collapse hidden overflow-hidden transition-all duration-300 basis-full grow md:block">
		        <div class="overflow-hidden overflow-y-auto max-h-[75vh] [&::-webkit-scrollbar]:w-2 [&::-webkit-scrollbar-thumb]:rounded-full [&::-webkit-scrollbar-track]:bg-gray-100 [&::-webkit-scrollbar-thumb]:bg-gray-300">
		          <div class="flex flex-col gap-x-0 mt-5 divide-y divide-dashed divide-gray-200 md:flex-row md:items-center md:justify-end md:gap-x-7 md:mt-0 md:ps-7 md:divide-y-0 md:divide-solid">
		            <NuxtLink class="font-medium hover:text-gray-400 transition py-3 md:py-6" activeClass="text-blue-600" to="/">Beranda</NuxtLink>
		            <NuxtLink class="font-medium hover:text-gray-400 transition py-3 md:py-6" activeClass="text-blue-600" to="/pegawai">Pegawai</NuxtLink>
		            <NuxtLink class="font-medium hover:text-gray-400 transition py-3 md:py-6" activeClass="text-blue-600" to="/konfigurasi/group">Konfigurasi</NuxtLink>
		            <NuxtLink class="font-medium hover:text-gray-400 transition py-3 md:py-6" activeClass="text-blue-600" to="/">Approval</NuxtLink>
		            <div class="hs-dropdown [--strategy:static] md:[--strategy:absolute] [--adaptive:none] md:[--trigger:hover] py-3 md:py-6">
								  <button type="button" class="flex items-center w-full text-gray-800 hover:text-gray-600 font-medium">
								    Resources
								    <svg class="flex-shrink-0 ms-2 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m6 9 6 6 6-6"/></svg>
								  </button>

								  <div class="hs-dropdown-menu transition-[opacity,margin] duration-[0.1ms] md:duration-[150ms] hs-dropdown-open:opacity-100 opacity-0 md:w-[705px] lg:w-[750px] hidden z-10 top-full end-0 overflow-hidden bg-white md:shadow-2xl rounded-lg before:absolute before:-top-5 before:start-0 before:w-full before:h-5">
								    <div class="grid grid-cols-2 md:grid-cols-10">
								      <div class="md:col-span-3">
								        <div class="flex flex-col py-6 px-3 md:px-6">
								          <div class="space-y-4">
								            <span class="mb-2 text-xs font-semibold uppercase text-gray-800">About us</span>

								            <a class="flex gap-x-4 text-gray-800 hover:text-blue-600" href="#">
								              <svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z"/><path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z"/></svg>
								              <div class="grow">
								                <p>Support Docs</p>
								              </div>
								            </a>

								            <a class="flex gap-x-4 text-gray-800 hover:text-blue-600" href="#">
								              <svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="7" height="7" x="14" y="3" rx="1"/><path d="M10 21V8a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1H3"/></svg>
								              <div class="grow">
								                <p>Integrations</p>
								              </div>
								            </a>

								            <a class="flex gap-x-4 text-gray-800 hover:text-blue-600" href="#">
								              <svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><polyline points="14 2 14 8 20 8"/><line x1="16" x2="8" y1="13" y2="13"/><line x1="16" x2="8" y1="17" y2="17"/><line x1="10" x2="8" y1="9" y2="9"/></svg>
								              <div class="grow">
								                <p>Guides</p>
								              </div>
								            </a>

								            <a class="flex gap-x-4 text-gray-800 hover:text-blue-600" href="#">
								              <svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m7 11 2-2-2-2"/><path d="M11 13h4"/><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/></svg>
								              <div class="grow">
								                <p>API Reference</p>
								              </div>
								            </a>

								            <a class="flex gap-x-4 text-gray-800 hover:text-blue-600" href="#">
								              <svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 12h-4l-3 9L9 3l-3 9H2"/></svg>
								              <div class="grow">
								                <p>API Status</p>
								              </div>
								            </a>
								          </div>
								        </div>
								      </div>

								      <div class="md:col-span-3">
								        <div class="flex flex-col py-6 px-3 md:px-6">
								          <div class="space-y-4">
								            <span class="mb-2 text-xs font-semibold uppercase text-gray-800">Support</span>

								            <a class="flex gap-x-4 text-gray-800 hover:text-blue-600" href="#">
								              <svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"/><path d="M12 17h.01"/></svg>
								              <div class="grow">
								                <p>Help Center</p>
								              </div>
								            </a>

								            <a class="flex gap-x-4 text-gray-800 hover:text-blue-600" href="#">
								              <svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="4"/><path d="M16 8v5a3 3 0 0 0 6 0v-1a10 10 0 1 0-4 8"/></svg>
								              <div class="grow">
								                <p>Developer Hub</p>
								              </div>
								            </a>

								            <a class="flex gap-x-4 text-gray-800 hover:text-blue-600" href="#">
								              <svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" ><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M22 21v-2a4 4 0 0 0-3-3.87"/><path d="M16 3.13a4 4 0 0 1 0 7.75"/></svg>
								              <div class="grow">
								                <p>Community Forum</p>
								              </div>
								            </a>
								          </div>

								          <div class="mt-7 space-y-4">
								            <span class="mb-2 text-xs font-semibold uppercase text-gray-800">Partners</span>

								            <a class="flex gap-x-4 text-gray-800 hover:text-blue-600" href="#">
								              <svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 7V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2h-6"/><polyline points="14 2 14 8 20 8"/><path d="M5 17a3 3 0 1 0 0-6 3 3 0 0 0 0 6Z"/><path d="M7 16.5 8 22l-3-1-3 1 1-5.5"/></svg>
								              <div class="grow">
								                <p>Become a Partner</p>
								              </div>
								            </a>

								            <a class="flex gap-x-4 text-gray-800 hover:text-blue-600" href="#">
								              <svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m15 5 4 4"/><path d="M13 7 8.7 2.7a2.41 2.41 0 0 0-3.4 0L2.7 5.3a2.41 2.41 0 0 0 0 3.4L7 13"/><path d="m8 6 2-2"/><path d="m2 22 5.5-1.5L21.17 6.83a2.82 2.82 0 0 0-4-4L3.5 16.5Z"/><path d="m18 16 2-2"/><path d="m17 11 4.3 4.3c.94.94.94 2.46 0 3.4l-2.6 2.6c-.94.94-2.46.94-3.4 0L11 17"/></svg>
								              <div class="grow">
								                <p>Build on Preline</p>
								              </div>
								            </a>
								          </div>
								        </div>
								      </div>

								      <div class="col-span-full md:col-span-4">
								        <div class="flex flex-col bg-gray-50 p-6">
								          <span class="text-xs font-semibold uppercase text-gray-800">Customer stories</span>

								          <a class="mt-4 group" href="#">
								            <div class="aspect-w-16 aspect-h-9">
								              <img class="w-full object-cover rounded-lg" src="https://images.unsplash.com/photo-1661956602116-aa6865609028?ixlib=rb-4.0.3&ixid=MnwxMjA3fDF8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1064&q=80" alt="Image Description">
								            </div>
								            <div class="mt-2">
								              <p class="text-gray-800">Preline Projects has proved to be most efficient cloud based project tracking and bug tracking tool.</p>
								              <p class="mt-3 inline-flex items-center gap-x-1 text-sm text-blue-600 decoration-2 hover:underline font-medium">
								                Learn more
								                <svg class="flex-shrink-0 size-4 transition ease-in-out group-hover:translate-x-1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m9 18 6-6-6-6"/></svg>
								              </p>
								            </div>
								          </a>
								        </div>
								      </div>
								    </div>
								  </div>
								</div>

		            <a class="font-medium text-gray-500 hover:text-gray-400 py-3 md:py-6" href="#">
		              Join us <span class="py-0.5 px-1.5 rounded-full text-xs font-medium bg-blue-50 border border-blue-200 text-blue-600">4</span>
		            </a>

		            <div class="pt-3 md:pt-0">
		              <UDropdown :items="items" :ui="{ item: { disabled: 'cursor-text select-text' } }" :popper="{ placement: 'bottom-end' }">
		              	<div class="flex items-center">
		              		<label>{{userInfo?.nama}}</label>
						    			<UAvatar src="https://avatars.githubusercontent.com/u/739984?v=4" />
						    		</div>

								    <template #account="{ item }">
								      	<div class="text-left">
								        	<p class="text-xs text-gray-900 dark:text-white">
								          		{{ item.label }}
								        	</p>
								      	</div>
								    </template>

								    <template #item="{ item }">
								      	<span class="truncate">{{ item.label }}</span>

								      	<UIcon :name="item.icon" class="flex-shrink-0 h-4 w-4 text-gray-400 ms-auto" />
								    </template>
									</UDropdown>
		            </div>
		          </div>
		        </div>
		      </div>
		    </div>
		  </nav>
		</header>
		<!-- ========== END HEADER ========== -->
  	<div class="bg-gradient-to-r from-cyan-500 to-blue-500 h-36"></div>
  	<GlobalSpinner>
			<div class="w-full pt-5 px-4 sm:px-6 md:px-8">
				<main class="p-md-4 max-w-[85rem] mx-auto bg">
					<slot />
				</main>
			</div>
		</GlobalSpinner>
	</div>
</template>