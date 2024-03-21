<script lang="ts" setup>
const { awesome } = useAppConfig()

useHead({
    //title: awesome.name,
    //titleTemplate: `%s - ${awesome.name}`,
})
</script>

<template>
  <Body class="antialiased duration-300 transition-colors text-gray-800 bg-gray-300">
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
    <UNotifications />
  </Body>
</template>