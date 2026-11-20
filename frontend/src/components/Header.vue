<script setup>
import { ref, onMounted, onBeforeUnmount } from "vue";
import { useI18n } from "vue-i18n";
import { useRouter } from "vue-router";
import { WindowMinimise, WindowToggleMaximise, Quit } from "../../wailsjs/runtime";
import { useConfigStore, THEMES } from "../stores/config";

const { t, availableLocales: languages, locale } = useI18n();
const router = useRouter();
const configStore = useConfigStore();

const localeCodes = { "en": "EN", "zh-Hans": "ZH" };

const themeOpen = ref(false);
const langOpen = ref(false);

const toggleThemeMenu = () => {
  themeOpen.value = !themeOpen.value;
  langOpen.value = false;
};
const toggleLangMenu = () => {
  langOpen.value = !langOpen.value;
  themeOpen.value = false;
};
const closeMenus = () => {
  themeOpen.value = false;
  langOpen.value = false;
};

const onclickLanguageHandle = (item) => {
  if (item !== locale.value) {
    locale.value = item;
    configStore.setLocale(item);
  }
  langOpen.value = false;
};

const selectTheme = (th) => {
  configStore.setTheme(th);
  themeOpen.value = false;
};

onMounted(() => document.addEventListener("click", closeMenus));
onBeforeUnmount(() => document.removeEventListener("click", closeMenus));

const minimise = () => WindowMinimise();
const toggleMaximise = () => WindowToggleMaximise();
const close = () => Quit();
</script>

<template>
  <header
    class="draggable relative z-50 flex h-12 items-center border-b border-base-300 bg-base-200 select-none">
    <!-- 标题（左） -->
    <div class="px-3 font-bold">{{ t('g.title') }}</div>

    <!-- 语言 + 主题 + 设置（绝对居中，位置固定） -->
    <div
      class="no-drag absolute left-1/2 top-1/2 flex -translate-x-1/2 -translate-y-1/2 items-center gap-1">
      <!-- 语言切换 -->
      <div class="relative flex h-8 items-center" @click.stop>
        <button
          class="btn btn-ghost btn-sm h-8 gap-1.5"
          :title="t('topbar.language')"
          @click="toggleLangMenu">
          <svg
            class="h-5 w-5 opacity-70"
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
          <svg
            class="h-3 w-3 opacity-60"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round">
            <path d="m6 9 6 6 6-6" />
          </svg>
        </button>
        <div
          v-if="langOpen"
          class="absolute left-1/2 top-full z-50 mt-3 w-40 -translate-x-1/2 rounded-box bg-base-200 p-2 shadow-xl">
          <button
            v-for="item in [...languages]"
            :key="item"
            class="flex w-full items-center gap-3 rounded-lg px-2 py-1.5 text-left text-sm hover:bg-base-300"
            @click="onclickLanguageHandle(item)">
            <span class="rounded-md bg-base-100 p-1 font-mono text-xs font-bold opacity-60 shadow-sm">{{ localeCodes[item] }}</span>
            <span class="flex-1">{{ t('languages.' + item) }}</span>
            <svg
              v-if="item === locale"
              class="h-3 w-3"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round">
              <path d="M20 6 9 17l-5-5" />
            </svg>
          </button>
        </div>
      </div>

      <!-- 主题切换 -->
      <div class="relative flex h-8 items-center" @click.stop>
        <button
          class="btn btn-ghost btn-sm h-8 gap-1.5"
          :title="t('topbar.toggleTheme')"
          @click="toggleThemeMenu">
          <div class="grid h-5 w-5 shrink-0 grid-cols-2 gap-0.5 rounded-md border border-base-content/10 p-1">
            <div class="bg-base-content h-1 w-1 rounded-full"></div>
            <div class="bg-primary h-1 w-1 rounded-full"></div>
            <div class="bg-secondary h-1 w-1 rounded-full"></div>
            <div class="bg-accent h-1 w-1 rounded-full"></div>
          </div>
          <svg
            class="h-3 w-3 opacity-60"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round">
            <path d="m6 9 6 6 6-6" />
          </svg>
        </button>
        <div
          v-if="themeOpen"
          class="absolute left-1/2 top-full z-50 mt-3 max-h-[60vh] w-60 -translate-x-1/2 overflow-y-auto rounded-box bg-base-200 p-2 shadow-xl">
          <button
            v-for="th in THEMES"
            :key="th"
            class="flex w-full items-center gap-3 rounded-lg px-2 py-1.5 text-left text-sm hover:bg-base-300"
            @click="selectTheme(th)">
            <div :data-theme="th" class="bg-base-100 grid grid-cols-2 gap-0.5 rounded-md p-1 shadow-sm">
              <div class="bg-base-content h-1 w-1 rounded-full"></div>
              <div class="bg-primary h-1 w-1 rounded-full"></div>
              <div class="bg-secondary h-1 w-1 rounded-full"></div>
              <div class="bg-accent h-1 w-1 rounded-full"></div>
            </div>
            <span class="flex-1">{{ t('themes.' + th) }}</span>
            <svg
              v-if="configStore.theme === th"
              class="h-3 w-3"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round">
              <path d="M20 6 9 17l-5-5" />
            </svg>
          </button>
        </div>
      </div>

      <!-- 设置 -->
      <button
        class="btn btn-ghost btn-circle btn-sm"
        :title="t('topbar.settings')"
        @click="router.push('/settings')">
        <svg
          class="h-5 w-5"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round">
          <path d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z" />
          <circle cx="12" cy="12" r="3" />
        </svg>
      </button>
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
