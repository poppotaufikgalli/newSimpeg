<script lang="ts" setup>
import { storeToRefs } from 'pinia'; // import storeToRefs helper hook from pinia
import { useModalStore } from '~/store/modalStore'; // import the auth store we just created

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
    <ShowModal />
    <ShowModalUpload />
    <ModalUploadDoc />
  </Body>
</template>