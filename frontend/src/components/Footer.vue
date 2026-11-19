<script setup>
import { useI18n } from "vue-i18n";
import { BrowserOpenURL } from "../../wailsjs/runtime";
const { t } = useI18n();
</script>

<template>
  <footer
    class="flex w-full items-center justify-center border-t border-base-300 bg-base-200 p-3 select-none">
    <div
      @click="BrowserOpenURL('https://wiiiiill.github.io/')"
      class="cursor-pointer text-sm opacity-70 hover:opacity-100">
      {{ t("footer.by") }}
      <span class="underline font-bold">github.com/Wiiiiill</span>
      {{ t("footer.dev") }}
    </div>
  </footer>
</template>
