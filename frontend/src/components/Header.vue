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
  // 选择后关闭下拉菜单
  if (document.activeElement) {
    document.activeElement.blur();
  }
};

const minimise = () => WindowMinimise();
const toggleMaximise = () => WindowToggleMaximise();
const close = () => Quit();
</script>

<template>
  <header
    class="draggable relative flex h-12 items-center border-b border-base-300 bg-base-200 select-none">
    <!-- 标题（左） -->
    <div class="px-3 font-bold">{{ t('g.title') }}</div>

    <!-- 语言 + 主题切换（绝对居中，位置固定） -->
    <div
      class="no-drag absolute left-1/2 top-1/2 flex -translate-x-1/2 -translate-y-1/2 items-center gap-1">
      <!-- 语言切换（地球图标 + 下拉菜单） -->
      <div class="dropdown dropdown-bottom dropdown-center">
        <button
          tabindex="0"
          class="btn btn-ghost btn-circle btn-sm"
          :title="t('topbar.language')">
          <svg
            class="h-5 w-5"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round">
            <circle cx="12" cy="12" r="10" />
            <path d="M12 2a14.5 14.5 0 0 0 0 20 14.5 14.5 0 0 0 0-20" />
            <path d="M2 12h20" />
          </svg>
        </button>
        <ul
          tabindex="0"
          class="dropdown-content menu z-50 w-32 gap-1 rounded-box bg-base-200 p-2 shadow">
          <li v-for="item in [...languages]" :key="item">
            <button
              :class="item === locale ? 'active' : ''"
              @click="onclickLanguageHandle(item)">
              {{ t('languages.' + item) }}
            </button>
          </li>
        </ul>
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

    <!-- 窗口控制（右） -->
    <div class="no-drag ml-auto flex h-full items-stretch">
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
