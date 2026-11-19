<script setup>
import { useI18n } from "vue-i18n";
import { WindowMinimise, WindowToggleMaximise, Quit } from "../../wailsjs/runtime";
import { useThemeStore } from "../stores/theme";

const { t, availableLocales: languages, locale } = useI18n();
const themeStore = useThemeStore();

// 初始化语言：优先读取本地存储，否则默认中文
const savedLocale = localStorage.getItem("locale");
locale.value = savedLocale && languages.includes(savedLocale) ? savedLocale : "zh-Hans";

const onclickLanguageHandle = (item) => {
  if (item !== locale.value) {
    locale.value = item;
    localStorage.setItem("locale", item);
  }
};

const minimise = () => WindowMinimise();
const toggleMaximise = () => WindowToggleMaximise();
const close = () => Quit();
</script>

<template>
  <header
    class="draggable flex h-12 items-center justify-between border-b border-base-300 bg-base-200 select-none">
    <!-- 标题 -->
    <div class="px-3 font-bold">{{ t('g.title') }}</div>

    <!-- 语言 + 主题切换 -->
    <div class="no-drag flex flex-1 items-center justify-center gap-4">
      <div class="flex items-center gap-1 text-sm">
        <template v-for="(item, index) in [...languages]" :key="item">
          <span
            :class="[
              item === locale
                ? 'font-bold text-primary'
                : 'cursor-pointer opacity-60 hover:opacity-100',
            ]"
            @click="onclickLanguageHandle(item)">
            {{ t('languages.' + item) }}
          </span>
          <span v-if="index < [...languages].length - 1" class="text-base-300">|</span>
        </template>
      </div>

      <!-- 主题切换开关 -->
      <label class="swap swap-rotate" :title="t('topbar.toggleTheme')">
        <input
          type="checkbox"
          :checked="themeStore.theme === 'dark'"
          @change="themeStore.toggle()" />
        <!-- 太阳（暗色时显示，点击切回亮色） -->
        <svg
          class="swap-on h-5 w-5"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round">
          <circle cx="12" cy="12" r="4" />
          <path d="M12 2v2" />
          <path d="M12 20v2" />
          <path d="m4.93 4.93 1.41 1.41" />
          <path d="m17.66 17.66 1.41 1.41" />
          <path d="M2 12h2" />
          <path d="M20 12h2" />
          <path d="m6.34 17.66-1.41 1.41" />
          <path d="m19.07 4.93-1.41 1.41" />
        </svg>
        <!-- 月亮（亮色时显示，点击切到暗色） -->
        <svg
          class="swap-off h-5 w-5"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round">
          <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z" />
        </svg>
      </label>
    </div>

    <!-- 窗口控制 -->
    <div class="no-drag flex h-full items-stretch">
      <button
        class="flex w-11 items-center justify-center hover:bg-base-300"
        :title="t('topbar.minimise')"
        @click="minimise">
        <svg
          class="h-4 w-4"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round">
          <path d="M5 12h14" />
        </svg>
      </button>
      <button
        class="flex w-11 items-center justify-center hover:bg-base-300"
        :title="t('topbar.maximise')"
        @click="toggleMaximise">
        <svg
          class="h-4 w-4"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round">
          <rect width="18" height="18" x="3" y="3" rx="2" />
        </svg>
      </button>
      <button
        class="flex w-11 items-center justify-center hover:bg-red-600 hover:text-white"
        :title="t('topbar.close')"
        @click="close">
        <svg
          class="h-4 w-4"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round">
          <path d="M18 6 6 18" />
          <path d="m6 6 12 12" />
        </svg>
      </button>
    </div>
  </header>
</template>
