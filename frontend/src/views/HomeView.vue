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
function setUpAttr(){
  ReplaceFilenames(path.value,"_1.","_属性.")
}

</script>
<template>
  <div class="flex h-full flex-col justify-center p-5">
    <div class="mb-3 flex items-center gap-2">
      <h1 class="text-base font-semibold">
        {{ t('main.title') }}
      </h1>
      <span class="text-xs opacity-60"><span class="text-red-600">*</span>{{ t('main.support-formats') }}</span>
    </div>

    <!-- 第一行：目录 + 操作按钮 -->
    <div class="flex items-end gap-2">
      <label class="form-control w-full max-w-xs">
        <div class="label py-1">
          <span class="label-text">{{ t('main.directory') }}:</span>
        </div>
        <input type="text" v-model="path" :placeholder="t('main.folderTip')"
          class="input input-bordered input-sm w-full max-w-xs" />
      </label>
      <button class="btn btn-sm" @click="openFolder">{{ t('main.FolderBtn') }}</button>
      <button class="btn btn-sm" :disabled="!path" @click="decodeFilenames">{{ t('main.decodeBtn') }}</button>
      <button class="btn btn-sm" :disabled="!path" @click="copyFilesBasedOnPrefix">{{ t('main.fileBtn') }}</button>
      <button class="btn btn-sm" :disabled="!path" @click="setUpAttr">{{ t('main.setUpAttrBtn') }}</button>
    </div>

    <!-- 第二行：替换文件名 -->
    <div class="mt-3 flex items-end gap-2">
      <label class="form-control w-full max-w-xs">
        <div class="label py-1">
          <span class="label-text">{{ t('main.ofn') }}:</span>
        </div>
        <input type="text" v-model="originFilename" :placeholder="t('main.ofnTip')"
          class="input input-bordered input-sm w-full max-w-xs" />
      </label>
      <label class="form-control w-full max-w-xs">
        <div class="label py-1">
          <span class="label-text">{{ t('main.nfn') }}:</span>
        </div>
        <input type="text" v-model="replaceFilename" :placeholder="t('main.nfnTip')"
          class="input input-bordered input-sm w-full max-w-xs" />
      </label>
      <button @click="replaceFilenames" class="btn btn-sm" :disabled="!path || !originFilename || !replaceFilename">{{ t('main.replaceBtn') }}</button>
    </div>
  </div>
</template>
