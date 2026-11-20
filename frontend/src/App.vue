<script setup>
import { onMounted, onBeforeUnmount, nextTick, ref, watch } from "vue";
import { useI18n } from "vue-i18n";
import { useConfigStore } from "./stores/config";
import { toast } from "./stores/toast";
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";
import { EventsOn, WindowSetBackgroundColour } from "../wailsjs/runtime";
import { ResolveClose } from "../wailsjs/go/main/App";

const { locale, t } = useI18n();
const configStore = useConfigStore();

// 关闭确认框：Go 端 beforeClose 触发 close:ask 事件后由前端展示
const showClosePrompt = ref(false);
let offCloseAsk = null;

const resolveClose = (action) => {
  showClosePrompt.value = false;
  ResolveClose(action).catch(() => {});
};

// 让原生窗口底色跟随当前主题：
// Windows 的无边框窗口在 CSS 圆角之外会露出窗口类背景色，
// macOS 则需要窗口背景保持透明，四角裁切才能透出桌面。
const syncWindowBackground = () => {
  if (typeof window === "undefined" || !window.runtime?.WindowSetBackgroundColour) {
    return;
  }

  // macOS 的窗口透明完全由原生层处理；再设置底色会在 CSS 圆角外露出矩形色块
  if (document.documentElement.dataset.platform === "darwin") {
    return;
  }

  const shell = document.querySelector(".app-shell");
  // Windows 的窗口底色只会从四角露出，四角紧邻 Header/Footer（base-200），
  // 取它们的底色可以避免在系统圆角与 CSS 圆角之间出现一圈异色
  const source = document.querySelector(".app-shell > header") || shell;
  const canvas = document.createElement("canvas");
  const ctx = canvas.getContext("2d");
  if (!source || !ctx) return;

  // daisyUI 4 主题色是 oklch，借助 canvas 转成 sRGB 整数；
  // 先填白色兜底，避免旧内核不认识 oklch 时把原生底色刷成黑色
  ctx.fillStyle = "#ffffff";
  ctx.fillStyle = getComputedStyle(source).backgroundColor;
  ctx.fillRect(0, 0, 1, 1);
  const [r, g, b, a] = ctx.getImageData(0, 0, 1, 1).data;
  if (a === 0) return;

  // 原生背景取主题底色但保持全透明
  WindowSetBackgroundColour(r, g, b, 0);
};

const syncAfterThemeChange = async () => {
  await nextTick();
  syncWindowBackground();
};

onMounted(async () => {
  await configStore.load();
  locale.value = configStore.locale;
  syncAfterThemeChange();
  offCloseAsk = EventsOn("close:ask", () => {
    showClosePrompt.value = true;
  });
});

onBeforeUnmount(() => {
  if (offCloseAsk) offCloseAsk();
});

watch(() => configStore.theme, syncAfterThemeChange);
</script>

<template>
  <div class="app-shell flex h-full flex-col bg-base-100 text-base-content">
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

    <!-- 关闭确认框 -->
    <div v-if="showClosePrompt" class="fixed inset-0 z-[100] flex items-center justify-center bg-black/40">
      <div class="w-80 rounded-box bg-base-100 p-5 shadow-xl">
        <h3 class="mb-2 text-base font-bold">{{ t('closePrompt.title') }}</h3>
        <p class="mb-4 text-sm opacity-70">{{ t('closePrompt.message') }}</p>
        <div class="flex justify-end gap-2">
          <button class="btn btn-ghost btn-sm" @click="showClosePrompt = false">{{ t('closePrompt.cancel') }}</button>
          <button class="btn btn-outline btn-sm" @click="resolveClose('quit')">{{ t('closePrompt.quit') }}</button>
          <button class="btn btn-primary btn-sm" @click="resolveClose('tray')">{{ t('closePrompt.tray') }}</button>
        </div>
      </div>
    </div>
  </div>
</template>
