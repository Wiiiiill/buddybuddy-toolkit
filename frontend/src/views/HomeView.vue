<script setup>
import { useI18n } from "vue-i18n";
import { computed, ref } from 'vue'
import { CopyFilesBasedOnPrefix, DecodeFilenames, OpenFolder, ReplaceFilenames } from '../../wailsjs/go/main/App'
import { useConfigStore } from '../stores/config'
import { showToast } from '../stores/toast'
import { compressLargeImages } from '../lib/batchCompress'

const path = ref('')
const originFilename = ref('')
const replaceFilename = ref('')
const { t } = useI18n();
const configStore = useConfigStore();

const compressing = ref(false)
const progress = ref(null)
const summary = ref(null)

const targetMB = computed(() => {
  const bytes = configStore.compressTargetBytes || 2 * 1024 * 1024
  const mb = bytes / 1024 / 1024
  return Number.isInteger(mb) ? String(mb) : mb.toFixed(1)
})

function formatBytes(n) {
  if (n >= 1024 * 1024) return (n / 1024 / 1024).toFixed(1) + ' MB'
  if (n >= 1024) return (n / 1024).toFixed(0) + ' KB'
  return n + ' B'
}

async function compressImages() {
  if (!path.value || compressing.value) return
  compressing.value = true
  summary.value = null
  progress.value = null
  try {
    const result = await compressLargeImages(
      path.value,
      configStore.compressTargetBytes,
      (p) => { progress.value = p }
    )
    summary.value = result
    if (result.total === 0) {
      showToast(t('main.compressNoFiles'))
      return
    }
    let message = t('main.compressDone', {
      compressed: result.compressed,
      skipped: result.skipped,
    })
    const details = []
    if (result.belowTarget) details.push(t('main.compressBelowTarget', { count: result.belowTarget }))
    if (result.failed) details.push(t('main.compressFailed', { count: result.failed }))
    if (details.length) message += '（' + details.join('，') + '）'
    showToast(message, result.compressed > 0 ? 'success' : 'error')
  } catch (err) {
    showToast(String(err), 'error')
  } finally {
    compressing.value = false
  }
}

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
  <div class="flex h-full flex-col overflow-y-auto">
    <div class="my-auto flex flex-col gap-3 p-5">
    <!-- ① 目录（前提） -->
    <div class="rounded-box border border-base-content/20 p-4">
      <div class="mb-3 text-xs font-bold uppercase tracking-wide opacity-60">{{ t('main.directory') }}</div>
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
      <div class="rounded-box border border-base-content/20 p-4">
        <div class="mb-3 flex items-center justify-between gap-2">
          <div class="text-xs font-bold uppercase tracking-wide opacity-60">{{ t('main.fileOps') }}</div>
          <span class="shrink-0 text-xs opacity-60">{{ t('main.compressTarget', { size: targetMB }) }}</span>
        </div>

        <div class="flex flex-wrap items-center gap-2">
          <button class="btn btn-sm" @click="decodeFilenames">{{ t('main.decodeBtn') }}</button>
          <button class="btn btn-sm" @click="copyFilesBasedOnPrefix">{{ t('main.fileBtn') }}</button>
          <button class="btn btn-primary btn-sm" :disabled="compressing" @click="compressImages">
            {{ compressing ? t('main.compressing') : t('main.compressBtn') }}
          </button>
        </div>

        <div v-if="compressing && progress" class="mt-3">
          <progress class="progress progress-primary h-1.5 w-full" :value="progress.done" :max="progress.total"></progress>
          <div class="mt-1 flex items-center justify-between gap-3 text-xs opacity-60">
            <span class="truncate">{{ progress.current || t('main.compressing') }}</span>
            <span class="shrink-0">{{ progress.done }}/{{ progress.total }}</span>
          </div>
        </div>
        <div v-else-if="summary" class="mt-2 text-xs opacity-70">
          {{ t('main.compressSummary', { compressed: summary.compressed, skipped: summary.skipped }) }}
          <template v-if="summary.compressed > 0">
            · {{ formatBytes(summary.beforeBytes) }} → {{ formatBytes(summary.afterBytes) }}
          </template>
        </div>

        <div class="my-3 border-t border-base-content/10"></div>

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
  </div>
</template>
