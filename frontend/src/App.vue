<script setup>
import { onMounted } from "vue";
import { useI18n } from "vue-i18n";
import { useConfigStore } from "./stores/config";
import { toast } from "./stores/toast";
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";

const { locale } = useI18n();
const configStore = useConfigStore();

onMounted(async () => {
  await configStore.load();
  locale.value = configStore.locale;
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
