<script setup>
import { onMounted } from "vue";
import { useThemeStore } from "./stores/theme";
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";

const themeStore = useThemeStore();
onMounted(() => themeStore.init());
</script>

<template>
  <div class="flex h-full flex-col bg-base-100 text-base-content">
    <Header />
    <main class="flex-1 overflow-hidden">
      <router-view />
    </main>
    <Footer />
  </div>
</template>
