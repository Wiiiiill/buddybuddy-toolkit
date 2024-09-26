<script setup>
import { useI18n } from "vue-i18n";
import { BrowserOpenURL } from "../../wailsjs/runtime";
const { t } = useI18n();
</script>
<template>
    <div
        class="bg-slate-400 select-none flex absolute bottom-0 w-full p-[12px] items-center justify-center content-center">
        <div @click="BrowserOpenURL('https://wiiiiill.github.io/')">
            {{ t("footer.by") }}
            <span class="underline cursor-pointer font-bold">github.com/Wiiiiill</span>
            {{ t("footer.dev") }}
        </div>
    </div>
</template>
<style lang="scss" scoped></style>