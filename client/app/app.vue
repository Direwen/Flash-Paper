<template>
  <div class="dark">
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </div>
</template>

<script lang="ts" setup>
import { useAuthStore } from './stores/auth';


const authStore = useAuthStore()
onMounted(() => {
  authStore.initAuth()
})
</script>