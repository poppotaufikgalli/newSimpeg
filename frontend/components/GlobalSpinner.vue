<script setup>
//const loading = ref(false)

const nuxtApp = useNuxtApp();
const loading = ref(false);
nuxtApp.hook("page:start", () => {
  loading.value = true;
});
nuxtApp.hook("page:finish", () => {
  loading.value = false;
});
</script>
<template>
  <div class="relative">
    <div class="flex -mt-36 pt-10">
      <slot />
    </div>

    <template v-if="loading">
      <div class="absolute top-0 start-0 size-full bg-gray-900/[.5] rounded-lg"></div>

      <div class="absolute top-1/2 start-1/2 transform -translate-x-1/2 -translate-y-1/2">
        <div class="animate-spin inline-block size-6 border-[3px] border-current border-t-transparent text-blue-600 rounded-full" role="status" aria-label="loading">
          <span class="sr-only">Loading...</span>
        </div>
      </div>
    </template>
  </div>
</template>