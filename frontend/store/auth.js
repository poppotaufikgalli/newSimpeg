import { defineStore } from "pinia"

export const useAuthStore = defineStore("auth", {
  state: () => ({
    authenticated: false,
    message: '',
    loading: false,
    userInfo: null,
    nama: '',
  }),
  actions: {
    async authenticateUser({ username, password }) {
      const {data, token, error, pending, refresh, status} = await $fetch("api/login", {
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
        this.nama = data?.datauser.nama
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
      token.value = null // clear the token cookie
    }
  }
})
