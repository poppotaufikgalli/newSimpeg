<script setup>
import { useAuthStore } from '~/store/auth'; // import the auth store we just created
const { authenticated } = storeToRefs(useAuthStore()); // make authenticated state reactive with storeToRefs
</script>
<template>
	<header>
      	<p class="mb-2 text-sm font-semibold text-blue-600">Starter Pages & Examples</p>
      	<h1 class="block text-2xl font-bold text-gray-800 sm:text-3xl dark:text-white">Application Layout: Sidebar Dark using Tailwind CSS</h1>
      	<p class="mt-2 text-lg text-gray-800 dark:text-gray-400">This is a simple application layout with sidebar and header examples using Tailwind CSS.</p>
      	<div class="mt-5 flex flex-col items-center gap-2 sm:flex-row sm:gap-3">
        	<a class="w-full sm:w-auto py-3 px-4 inline-flex justify-center items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none dark:focus:outline-none dark:focus:ring-1 dark:focus:ring-gray-600" href="https://github.com/htmlstreamofficial/preline/tree/main/examples/html" target="_blank">
          		<svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
            		<path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.012 8.012 0 0 0 16 8c0-4.42-3.58-8-8-8z"/>
          		</svg>
          		Get the source code
        	</a>
        	<a class="w-full sm:w-auto py-3 px-4 inline-flex justify-center items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent text-blue-600 hover:text-blue-800 disabled:opacity-50 disabled:pointer-events-none dark:text-blue-500 dark:hover:text-blue-400 dark:focus:outline-none dark:focus:ring-1 dark:focus:ring-gray-600" href="../examples.html">
          		<svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m15 18-6-6 6-6"/></svg>
          		Back to examples
        	</a>
      	</div>
    </header>
	<header id="mainNavbar" class="navbar navbar-expand-lg bg-white sticky-top flex-md-nowrap p-0 shadow">
		<div class="container-fluid">
		  	<a class="navbar-brand w-100 px-3 fs-6 text-white" href="#">Sistem Informasi Kepegawaian</a>
		  	<div class="collapse navbar-collapse end pe-3">
		  		<ul class="navbar-nav flex-row">
			    	<li class="nav-item dropdown">
			      		<button class="btn btn-sm btn-dark" type="button" data-bs-toggle="dropdown" aria-expanded="false">
			        		<i class="bi bi-bell"></i>
					  	</button>
					  	<ul class="dropdown-menu dropdown-menu-end">
					    	<li><button class="dropdown-item" type="button">Action</button></li>
					    	<li><button class="dropdown-item" type="button">Another action</button></li>
					    	<li><button class="dropdown-item" type="button">Something else here</button></li>
					  	</ul>
			    	</li>
			    	<li class="nav-item dropdown">
			          	<button class="btn btn-sm btn-dark dropdown-toggle" type="button" data-bs-toggle="dropdown" aria-expanded="false">
					    	Hi, User ...
					  	</button>
					  	<ul class="dropdown-menu dropdown-menu-end">
					  		<li v-if="!authenticated">
          						<NuxtLink class="dropdown-item" to="/login">Login</NuxtLink>
        					</li>
        					<li v-if="authenticated">
          						<NuxtLink class="dropdown-item" @click="logout">Logout</NuxtLink>
        					</li>
					    	<li><button class="dropdown-item" type="button">Action</button></li>
					    	<li><button class="dropdown-item" type="button">Another action</button></li>
					    	<li><button class="dropdown-item" type="button">Something else here</button></li>
					  	</ul>
			        </li>
			  	</ul>
		  	</div>
		</div>

	  	<ul class="navbar-nav flex-row d-md-none">
	    	<li class="nav-item text-nowrap">
	      		<button class="nav-link px-3 text-white" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSearch" aria-controls="navbarSearch" aria-expanded="false" aria-label="Toggle search">
	        		<i class="bi bi-search"></i>
	      		</button>
	    	</li>
	    	<li class="nav-item text-nowrap">
	      		<button class="nav-link px-3 text-white" type="button" data-bs-toggle="offcanvas" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
	        		<i class="bi bi-list"></i>
	      		</button>
	    	</li>
	  	</ul>

	  	<div id="navbarSearch" class="navbar-search w-100 collapse">
	    	<input class="form-control w-100 rounded-0 border-0" type="text" placeholder="Search" aria-label="Search">
	  	</div>
	</header>
</template>