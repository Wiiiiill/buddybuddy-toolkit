<script setup>
import { ref } from "vue";
import { useI18n } from "vue-i18n";
import { useRouter } from "vue-router";
import { useConfigStore } from "../stores/config";

const { t } = useI18n();
const router = useRouter();
const configStore = useConfigStore();

const activeTab = ref("settings");

// 本地编辑副本（进入页面时从 store 拷贝）
const formats = ref([...configStore.supportedFormats]);
const presets = ref(configStore.presets.map((p) => ({ ...p })));
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
  await configStore.save();
  saved.value = true;
  setTimeout(() => (saved.value = false), 2000);
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
        ⚙️ {{ t('settings.title') }}
      </button>
      <button
        :class="['btn btn-sm justify-start', activeTab === 'about' ? 'btn-primary' : 'btn-ghost']"
        @click="scrollTo('about')">
        ℹ️ {{ t('about.title') }}
      </button>
    </aside>

    <!-- 右侧长页面：设置在上，关于在下 -->
    <div class="flex-1 overflow-auto px-6 py-8 scroll-pt-8">
      <section id="settings-section" class="mb-16">
        <h2 class="mb-4 text-lg font-bold">{{ t('settings.title') }}</h2>

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

        <button class="btn btn-primary btn-sm" @click="save">
          {{ saved ? t('settings.saved') : t('settings.save') }}
        </button>
      </section>

      <section id="about-section" class="pb-[60vh]">
        <h2 class="mb-4 text-lg font-bold">{{ t('about.title') }}</h2>
        <p class="mb-2 text-sm opacity-80">{{ t('about.intro') }}</p>
        <p class="mb-4 text-sm">
          <span class="opacity-60">{{ t('about.author') }}:</span>
          <span class="font-semibold">github.com/Wiiiiill</span>
        </p>
        <p class="text-sm">
          <span class="opacity-60">{{ t('about.repo') }}:</span>
          <a class="link" href="https://github.com/Wiiiiill" target="_blank">github.com/Wiiiiill</a>
        </p>
      </section>
    </div>
  </div>
</template>
