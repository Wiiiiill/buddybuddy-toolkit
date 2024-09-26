<script setup>
import { useI18n } from "vue-i18n";
import { ref } from 'vue'
import { CopyFilesBasedOnPrefix, DecodeFilenames, OpenFolder, ReplaceFilenames } from '../../wailsjs/go/main/App'
const path = ref('')
const originFilename = ref('')
const replaceFilename = ref('')
const { t } = useI18n();
function openFolder() {
  OpenFolder().then(result => {
    path.value = result
  })
}
function decodeFilenames(){
  DecodeFilenames(path.value)
}
function copyFilesBasedOnPrefix(){
  CopyFilesBasedOnPrefix(path.value)
}
function replaceFilenames(){
  ReplaceFilenames(path.value,originFilename.value,replaceFilename.value)
}

</script>
<template>
  <div class="p-[24px]">
    <div>
      {{ t('main.title') }}
    </div>
    <div class="flex items-center gap-[6px]">
      <label class="form-control w-full max-w-xs">
        <div class="label">
          <span class="label-text">{{ t('main.directory') }}:</span>
        </div>
        <input type="text" v-model="path" :placeholder="t('main.folderTip')"
          class="input input-bordered w-full max-w-xs" />
        <div class="label">
          <span class="label-text-alt"><span class="text-red-600">*</span>{{ t('main.support-formats') }}</span>
        </div>
      </label>
      <button class="btn " @click="openFolder">{{ t('main.FolderBtn') }}</button>
      <button class="btn" :disabled="!path" @click="decodeFilenames">{{ t('main.decodeBtn') }}</button>
      <button class="btn" :disabled="!path" @click="copyFilesBasedOnPrefix">{{ t('main.fileBtn') }}</button>
    </div>
    <div class="flex items-center gap-[6px]">
      <label class="form-control w-full max-w-xs">
        <div class="label">
          <span class="label-text">{{ t('main.ofn') }}:</span>
        </div>
        <input type="text" v-model="originFilename" :placeholder="t('main.ofnTip')"
          class="input input-bordered w-full max-w-xs" />
        <div class="label">
          <span class="label-text-alt">&nbsp;</span>
        </div>
      </label>
      <label class="form-control w-full max-w-xs">
        <div class="label">
          <span class="label-text">{{ t('main.nfn') }}:</span>
        </div>
        <input type="text" v-model="replaceFilename" :placeholder="t('main.nfnTip')"
          class="input input-bordered w-full max-w-xs" />
        <div class="label">
          <span class="label-text-alt">&nbsp;</span>
        </div>
      </label>
      <button @click="replaceFilenames" class="btn" :disabled="!path || !originFilename || !replaceFilename">{{ t('main.replaceBtn') }}</button>
    </div>
  </div>
</template>