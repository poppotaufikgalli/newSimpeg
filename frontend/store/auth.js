import { defineStore } from "pinia"

export const useAuthStore = defineStore("auth", {
  state: () => ({
    authenticated: false,
    message: '',
    loading: false,
    userInfo: null,
    gid: null,
    nama: '',
  }),
  actions: {
    async authenticateUser({ username, password }) {
      console.log('do authenticateUser')
      const {data, token, error, pending, refresh, status} = await $fetch("/api/login", {
        method: "POST",
        body: {
          username: username,
          password: password,
        },
      });

      this.loading = pending;
      this.message = data.message || "Username atau Password Anda Salah";

      if (status == 200) {
        const token = useCookie("token") // useCookie new hook in nuxt 3
        token.value = data?.token // set token to cookie
        console.log(token.value)
        this.userInfo = data?.datauser
        this.gid = data?.datauser?.gid
        this.authenticated = true // set authenticated  state value to true  
      }
    },
    async islogin(){
      console.log('do islogin')
      const {data, error} = await useFetch("/api/islogin");
      //console.log(data.value.datauser)
      console.log(error.value)
      if(error.value){
        this.authenticated = false
      }else{
        this.userInfo = data.value.datauser
        this.gid = data.value.datauser?.gid
        this.authenticated = true
      }

      /*const {datauser, statusCode} = await $fetch("/api/islogin");
      
      if(statusCode == 200){
        this.userInfo = datauser
        this.gid = datauser?.gid
        this.authenticated = true // set authenticated  state value to true  
      }else{
        this.authenticated = false
      }*/
    },
    async authenticateFakeUser({username, password}) {
      const {data, token, error, pending, refresh, status} = await $fetch("api/fakelogin", {
        method: "POST",
        body: {
          username: username,
          password: password,
        },
      });

      this.loading = pending;
      this.message = data.message || "Username atau Password Anda Salah";

      if (status == 200) {
        const token = useCookie("token") // useCookie new hook in nuxt 3
        token.value = data?.token // set token to cookie
        this.userInfo = data?.datauser
        this.gid = data?.datauser.gid
        this.authenticated = true // set authenticated  state value to true  
      }
    },
    fakeauth(){
      const token = useCookie("token") // useCookie new hook in nuxt 3
      token.value = 'asdasdasasdsadsa' // set token to cookie
      this.authenticated = true // set authenticated  state value to true  
      this.userInfo = {
        nip: '1',
        nama: 'fake_admin',
      }
    },
    logUserOut() {
      console.log('do LogOut')
      const token = useCookie("token") // useCookie new hook in nuxt 3
      this.authenticated = false // set authenticated  state value to false
      this.gid = null
      token.value = null // clear the token cookie
    }
  }
})
