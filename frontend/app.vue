<script lang="ts" setup>
//const { awesome } = useAppConfig()
import { storeToRefs } from 'pinia'; // import storeToRefs helper hook from pinia
import { useAuthStore } from '~/store/auth'; // import the auth store we just created
const { islogin } = useAuthStore(); // use authenticateUser action from  auth store

onMounted(async() => {
  //await islogin();
})
</script>

<template>
  <Body class="antialiased duration-300 transition-colors text-gray-800 bg-gray-100">
    <NuxtLoadingIndicator />
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
    <UNotifications />
  </Body>
</template>