import { useAuthStore } from '~/store/auth';

export default defineNuxtRouteMiddleware((to, from) => {
    const { authenticated } = storeToRefs(useAuthStore()); // make authenticated state reactive
    const { userInfo } = storeToRefs(useAuthStore());
    const token = useCookie('token'); // get token from cookies

    console.log("middleware auth check")
    console.log(authenticated.value)
    console.log(userInfo.value)

    if (token.value) {
        // check if value exists
        authenticated.value = true; // update the state to authenticated
    }
  
    // if token exists and url is /login redirect to homepage
    if (token.value && to?.name === 'login') {
        return navigateTo('/');
    }

      // if token doesn't exist redirect to log in
    if (!token.value && to?.name !== 'login') {
        abortNavigation();
        return navigateTo('/login');
    }

    if(authenticated.value == false){
        abortNavigation();
        return navigateTo('/login');
    }
})