<script setup>
import { useI18n } from "vue-i18n";
import { WindowSetTitle } from "../../wailsjs/runtime";
const { t, availableLocales: languages, locale } = useI18n();
console.log(localStorage.getItem('locale'))
WindowSetTitle(t('g.title'))
const onclickLanguageHandle = (item) => {
    item !== locale.value ? (locale.value = item) : false;
    localStorage.setItem('locale', item)
    WindowSetTitle(t('g.title'))
};
onclickLanguageHandle(localStorage.getItem('locale') ? localStorage.getItem('locale') : 'zh-Hans')

</script>
<script>
</script>
<template>
    <div class="bg-slate-500 flex w-full p-[12px] content-center justify-between items-center select-none text-black">
        <div class="font-extrabold ">
            {{ t('g.title') }}
        </div>
        <div class=" flex fw-4 gap-[6px] cursor-pointer text-slate-900">
            <div v-for="item in [...languages]" :key="item"
                :class="{ 'text-black font-bold cursor-not-allowed': item === locale }"
                @click="onclickLanguageHandle(item)" class="lang-item">
                {{ t("languages." + item) }}
            </div>
        </div>
    </div>
</template>
<style lang="scss" scoped>
.lang-item:not(:first-child)::before {
    content: '|';
    margin-right: 6px;
    color: black;
    font-weight: 700;
}
</style>