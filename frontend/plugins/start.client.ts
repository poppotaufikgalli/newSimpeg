export default defineNuxtPlugin(async(nuxtApp) => {
  console.log("mengecek token den..")
    
  const config = useRuntimeConfig()
  const token = useCookie('token');
  
  if(token.value){
    const data = await fetch('/api/islogin').then((response) => 
      response.json().then(data => ({status: response.status, data: data}))
    )

    if(data.status !== 200){
      token.value = ""
    }
  }
})