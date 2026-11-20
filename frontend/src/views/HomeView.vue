<script setup>
import { useI18n } from "vue-i18n";
import { ref } from 'vue'
import { CopyFilesBasedOnPrefix, DecodeFilenames, OpenFolder, ReplaceFilenames } from '../../wailsjs/go/main/App'
import { useConfigStore } from '../stores/config'
import { showToast } from '../stores/toast'

const path = ref('')
const originFilename = ref('')
const replaceFilename = ref('')
const { t } = useI18n();
const configStore = useConfigStore();

function openFolder() {
  OpenFolder().then(result => {
    path.value = result
  })
}
function decodeFilenames(){
  DecodeFilenames(path.value)
    .then((count) => showToast(t('main.doneCount', { count })))
    .catch((err) => showToast(String(err), 'error'))
}
function copyFilesBasedOnPrefix(){
  CopyFilesBasedOnPrefix(path.value)
    .then((count) => showToast(t('main.doneCount', { count })))
    .catch((err) => showToast(String(err), 'error'))
}
function replaceFilenames(){
  ReplaceFilenames(path.value, originFilename.value, replaceFilename.value)
    .then((count) => showToast(t('main.doneCount', { count })))
    .catch((err) => showToast(String(err), 'error'))
}
function fillPreset(p) {
  originFilename.value = p.before
  replaceFilename.value = p.after
}
</script>
<template>
  <div class="flex h-full flex-col justify-center gap-3 p-5">
    <!-- ① 目录（前提） -->
    <div class="rounded-box border border-base-300 p-4">
      <div class="mb-3 text-xs font-bold uppercase tracking-wide opacity-60">📁 {{ t('main.directory') }}</div>
      <div class="flex items-center gap-2">
        <input type="text" v-model="path" :placeholder="t('main.folderTip')"
          class="input input-bordered input-sm flex-1" />
        <button class="btn btn-sm" @click="openFolder">{{ t('main.FolderBtn') }}</button>
      </div>
      <div class="mt-2 text-xs opacity-60">{{ t('main.support-formats') }}: {{ configStore.supportedFormats.join(' / ') }}</div>
    </div>

    <!-- 未选目录提示 -->
    <div v-if="!path" class="-my-1 text-center text-xs opacity-50">{{ t('main.needDir') }}</div>

    <!-- ② 操作区（依赖目录，未选目录时置灰禁用） -->
    <div :class="['flex flex-col gap-3', !path && 'pointer-events-none opacity-40']">
      <div class="rounded-box border border-base-300 p-4">
        <div class="mb-3 text-xs font-bold uppercase tracking-wide opacity-60">🛠 {{ t('main.fileOps') }}</div>
        <div class="flex items-center gap-2">
          <button class="btn btn-sm" @click="decodeFilenames">{{ t('main.decodeBtn') }}</button>
          <button class="btn btn-sm" @click="copyFilesBasedOnPrefix">{{ t('main.fileBtn') }}</button>
        </div>
      </div>

      <div class="rounded-box border border-base-300 p-4">
        <div class="mb-3 text-xs font-bold uppercase tracking-wide opacity-60">✏️ {{ t('main.replaceTitle') }}</div>
        <div v-if="configStore.presets.length" class="mb-3 flex flex-wrap items-center gap-2">
          <span class="text-xs opacity-60">{{ t('main.quick') }}</span>
          <button v-for="p in configStore.presets" :key="p.name + p.before + p.after"
            class="btn btn-xs rounded-full" @click="fillPreset(p)">
            {{ p.name }}: {{ p.before }} → {{ p.after }}
          </button>
        </div>
        <div class="flex items-end gap-2">
          <label class="form-control flex-1">
            <span class="label-text mb-1 text-xs">{{ t('main.ofn') }}</span>
            <input type="text" v-model="originFilename" :placeholder="t('main.ofnTip')"
              class="input input-bordered input-sm w-full" />
          </label>
          <label class="form-control flex-1">
            <span class="label-text mb-1 text-xs">{{ t('main.nfn') }}</span>
            <input type="text" v-model="replaceFilename" :placeholder="t('main.nfnTip')"
              class="input input-bordered input-sm w-full" />
          </label>
          <button class="btn btn-primary btn-sm" @click="replaceFilenames">{{ t('main.replaceBtn') }}</button>
        </div>
      </div>
    </div>
  </div>
</template>
