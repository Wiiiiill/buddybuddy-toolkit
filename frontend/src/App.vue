<script setup>
import { onMounted } from "vue";
import { useThemeStore } from "./stores/theme";
import { useConfigStore } from "./stores/config";
import { toast } from "./stores/toast";
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";

const themeStore = useThemeStore();
const configStore = useConfigStore();

onMounted(async () => {
  themeStore.init();
  await configStore.load();
});
</script>

<template>
  <div class="flex h-full flex-col bg-base-100 text-base-content">
    <Header />
    <main class="flex-1 overflow-hidden">
      <router-view />
    </main>
    <Footer />
    <div class="toast toast-end">
      <div v-if="toast.visible" :class="['alert', toast.type === 'success' ? 'alert-success' : 'alert-error']">
        <span>{{ toast.message }}</span>
      </div>
    </div>
  </div>
</template>
