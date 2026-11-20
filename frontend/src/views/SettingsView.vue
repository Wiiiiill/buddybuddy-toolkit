<script setup>
import { ref } from "vue";
import { useI18n } from "vue-i18n";
import { useRouter } from "vue-router";
import { useConfigStore, THEMES } from "../stores/config";
import { showToast } from "../stores/toast";
import { BrowserOpenURL } from "../../wailsjs/runtime";

const { t, locale, availableLocales: languages } = useI18n();
const router = useRouter();
const configStore = useConfigStore();

const localeCodes = { "en": "EN", "zh-Hans": "ZH" };

const activeTab = ref("settings");

// 本地编辑副本（进入页面时从 store 拷贝）
const formats = ref([...configStore.supportedFormats]);
const presets = ref(configStore.presets.map((p) => ({ ...p })));
const compressMB = ref(configStore.compressTargetBytes / 1024 / 1024);
const closeBehavior = ref(configStore.closeBehavior || "ask");
const saved = ref(false);

const addFormat = () => formats.value.push("");
const removeFormat = (i) => formats.value.splice(i, 1);

const addPreset = () => presets.value.push({ name: "", before: "", after: "" });
const removePreset = (i) => presets.value.splice(i, 1);

const save = async () => {
  configStore.supportedFormats = formats.value
    .map((f) => f.trim())
    .filter((f) => f !== "");
  configStore.presets = presets.value
    .filter((p) => p.before.trim() !== "" && p.after.trim() !== "")
    .map((p) => ({ name: p.name, before: p.before, after: p.after }));
  configStore.compressTargetBytes = Math.max(
    1,
    Math.round(compressMB.value * 1024 * 1024)
  );
  configStore.closeBehavior = closeBehavior.value;
  await configStore.save();
  saved.value = true;
  setTimeout(() => (saved.value = false), 2000);
};

const resetToDefault = async () => {
  await configStore.reset();
  formats.value = [...configStore.supportedFormats];
  presets.value = configStore.presets.map((p) => ({ ...p }));
  compressMB.value = configStore.compressTargetBytes / 1024 / 1024;
  closeBehavior.value = configStore.closeBehavior || "ask";
  locale.value = configStore.locale;
  showToast(t('settings.resetDone'));
};

const setLocale = (item) => {
  locale.value = item;
  configStore.setLocale(item);
};

// 左侧导航点击 → 滚动到对应区域
const scrollTo = (id) => {
  activeTab.value = id;
  const el = document.getElementById(id + "-section");
  if (el) el.scrollIntoView({ behavior: "smooth", block: "start" });
};
</script>

<template>
  <div class="flex h-full">
    <!-- 左侧导航 -->
    <aside class="flex w-44 shrink-0 flex-col gap-1 border-r border-base-300 bg-base-200 p-3">
      <button class="btn btn-ghost btn-sm justify-start" @click="router.push('/')">
        ← {{ t('settings.back') }}
      </button>
      <div class="my-2 border-t border-base-300"></div>
      <button
        :class="['btn btn-sm justify-start', activeTab === 'settings' ? 'btn-primary' : 'btn-ghost']"
        @click="scrollTo('settings')">
        {{ t('settings.title') }}
      </button>
      <button
        :class="['btn btn-sm justify-start', activeTab === 'theme' ? 'btn-primary' : 'btn-ghost']"
        @click="scrollTo('theme')">
        {{ t('settings.theme') }}
      </button>
      <button
        :class="['btn btn-sm justify-start', activeTab === 'about' ? 'btn-primary' : 'btn-ghost']"
        @click="scrollTo('about')">
        {{ t('about.title') }}
      </button>
    </aside>

    <!-- 右侧长页面：设置 → 主题 → 关于 -->
    <div class="flex-1 overflow-auto px-6 py-8 scroll-pt-8">
      <section id="settings-section" class="mb-16">
        <h2 class="mb-4 text-lg font-bold">{{ t('settings.title') }}</h2>

        <div class="mb-5">
          <div class="mb-2 text-sm font-semibold">{{ t('settings.language') }}</div>
          <div class="flex gap-2">
            <button
              v-for="item in [...languages]"
              :key="item"
              :class="['btn btn-sm gap-2', locale === item ? 'btn-primary' : 'btn-outline']"
              @click="setLocale(item)">
              <span class="font-mono text-xs font-bold opacity-60">{{ localeCodes[item] }}</span>
              <span>{{ t('languages.' + item) }}</span>
            </button>
          </div>
        </div>

        <div class="mb-5">
          <div class="mb-2 text-sm font-semibold">{{ t('settings.formats') }}</div>
          <div v-for="(f, i) in formats" :key="i" class="mb-2 flex items-center gap-2">
            <input v-model="formats[i]" class="input input-bordered input-sm flex-1" placeholder=".jpg" />
            <button class="btn btn-ghost btn-sm" @click="removeFormat(i)">✕</button>
          </div>
          <button class="btn btn-outline btn-sm" @click="addFormat">+ {{ t('settings.addFormat') }}</button>
        </div>

        <div class="mb-5">
          <div class="mb-2 text-sm font-semibold">{{ t('settings.presets') }}</div>
          <div v-for="(p, i) in presets" :key="i" class="mb-2 flex items-center gap-2">
            <input v-model="presets[i].name" class="input input-bordered input-sm w-24" :placeholder="t('settings.presetName')" />
            <input v-model="presets[i].before" class="input input-bordered input-sm flex-1" :placeholder="t('settings.find')" />
            <input v-model="presets[i].after" class="input input-bordered input-sm flex-1" :placeholder="t('settings.replaceTo')" />
            <button class="btn btn-ghost btn-sm" @click="removePreset(i)">✕</button>
          </div>
          <button class="btn btn-outline btn-sm" @click="addPreset">+ {{ t('settings.addPreset') }}</button>
        </div>

        <div class="mb-5">
          <div class="mb-2 text-sm font-semibold">{{ t('settings.compressTarget') }}</div>
          <div class="flex items-center gap-2">
            <input type="number" min="0.1" step="0.1" v-model.number="compressMB"
              class="input input-bordered input-sm w-28" />
            <span class="text-sm opacity-60">MB</span>
          </div>
          <div class="mt-1 text-xs opacity-50">{{ t('settings.compressTargetTip') }}</div>
        </div>

        <div class="mb-5">
          <div class="mb-2 text-sm font-semibold">{{ t('settings.closeBehavior') }}</div>
          <select v-model="closeBehavior" class="select select-bordered select-sm w-56">
            <option value="ask">{{ t('settings.closeAsk') }}</option>
            <option value="tray">{{ t('settings.closeTray') }}</option>
            <option value="quit">{{ t('settings.closeQuit') }}</option>
          </select>
          <div class="mt-1 text-xs opacity-50">{{ t('settings.closeTip') }}</div>
        </div>

        <div class="flex items-center gap-2">
          <button class="btn btn-primary btn-sm" @click="save">
            {{ saved ? t('settings.saved') : t('settings.save') }}
          </button>
          <button class="btn btn-ghost btn-sm" @click="resetToDefault">
            {{ t('settings.resetDefault') }}
          </button>
        </div>
      </section>

      <section id="theme-section" class="mb-16">
        <h2 class="mb-4 text-lg font-bold">{{ t('settings.theme') }}</h2>
        <div class="grid grid-cols-2 gap-3 sm:grid-cols-3 md:grid-cols-4">
          <button
            v-for="th in THEMES"
            :key="th"
            class="overflow-hidden rounded-lg border text-left transition-colors hover:border-base-content/40"
            :class="configStore.theme === th ? 'border-base-content outline outline-2 outline-offset-2 outline-base-content' : 'border-base-content/20'"
            @click="configStore.setTheme(th)">
            <div :data-theme="th" class="bg-base-100 text-base-content w-full font-sans">
              <div class="grid grid-cols-5 grid-rows-3">
                <div class="bg-base-200 col-start-1 row-span-2 row-start-1"></div>
                <div class="bg-base-300 col-start-1 row-start-3"></div>
                <div class="col-span-4 col-start-2 row-span-3 row-start-1 flex flex-col gap-1 p-2">
                  <div class="font-bold">{{ t('themes.' + th) }}</div>
                  <div class="flex flex-wrap gap-1">
                    <div class="bg-primary flex aspect-square w-5 items-center justify-center rounded">
                      <div class="text-primary-content text-sm font-bold">A</div>
                    </div>
                    <div class="bg-secondary flex aspect-square w-5 items-center justify-center rounded">
                      <div class="text-secondary-content text-sm font-bold">A</div>
                    </div>
                    <div class="bg-accent flex aspect-square w-5 items-center justify-center rounded">
                      <div class="text-accent-content text-sm font-bold">A</div>
                    </div>
                    <div class="bg-neutral flex aspect-square w-5 items-center justify-center rounded">
                      <div class="text-neutral-content text-sm font-bold">A</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </button>
        </div>
      </section>

      <section id="about-section" class="pb-[60vh]">
        <h2 class="mb-4 text-lg font-bold">{{ t('about.title') }}</h2>
        <p class="mb-2 text-sm opacity-80">{{ t('about.intro') }}</p>
        <p class="mb-4 text-sm">
          <span class="opacity-60">{{ t('about.author') }}:</span>
          <a
            class="link cursor-pointer font-semibold"
            @click="BrowserOpenURL('https://wiiiiill.github.io/')">Wiiiiill</a>
        </p>
        <p class="text-sm">
          <span class="opacity-60">{{ t('about.repo') }}:</span>
          <a
            class="link cursor-pointer"
            @click="BrowserOpenURL('https://github.com/Wiiiiill/buddybuddy-toolkit')">buddybuddy-toolkit</a>
        </p>
      </section>
    </div>
  </div>
</template>
