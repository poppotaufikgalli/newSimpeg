import { useAuthStore } from '~/store/auth';

export default defineNuxtPlugin(async(nuxtApp) => {
  console.log("mengecek token den..[plugins-start.client.ts]")
  //const config = useRuntimeConfig()

  const { islogin } = useAuthStore(); 
  //const { authenticated } = storeToRefs(useAuthStore()); // make authenticated state reactive
  //const token = useCookie('token');
  await islogin()  
  
  //if(token.value){
    //if(authenticated.value){
      
      //console.log("token", token.value)
      //console.log("authenticated", authenticated.value)
    //}
  //}
})