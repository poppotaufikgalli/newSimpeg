import { useAuthStore } from '~/store/auth';

export default defineNuxtRouteMiddleware((to, from) => {
    if(!process.client){
        //console.log("AAA")
        return
    }

    const { authenticated } = storeToRefs(useAuthStore()); // make authenticated state reactive
    
    if (authenticated.value && to?.name === 'login') {
        return navigateTo('/');
    }

      // if token doesn't exist redirect to log in
    if (!authenticated.value && to?.name !== 'login') {
        abortNavigation();
        return navigateTo('/login');
    }

    // if token exists and url is /login redirect to homepage
    if (!authenticated.value && to?.name === 'api/realm/tpi') {
        return navigateTo('/test');
    }
})