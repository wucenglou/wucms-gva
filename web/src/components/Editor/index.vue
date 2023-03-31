<script>
export default {
    name: 'Editor',
}
</script>
<script setup>
import { ref, reactive, onMounted, computed, watch } from 'vue'
import tinymce from 'tinymce/tinymce'
import TinymceEditor from '@tinymce/tinymce-vue'
import 'tinymce/themes/silver/theme'
import 'tinymce/icons/default/icons'
// import 'tinymce/models/dom'
// import "tinymce-plugin"

import 'tinymce/plugins/advlist'
import 'tinymce/plugins/anchor'
import 'tinymce/plugins/autolink'
import 'tinymce/plugins/autoresize'
import 'tinymce/plugins/autosave'
import 'tinymce/plugins/charmap'
import 'tinymce/plugins/code'
import 'tinymce/plugins/codesample'
import 'tinymce/plugins/directionality'
// import 'tinymce/plugins/emoticons'
import 'tinymce/plugins/fullscreen'
import 'tinymce/plugins/help'
import 'tinymce/plugins/image'
import 'tinymce/plugins/importcss'
import 'tinymce/plugins/insertdatetime'
import 'tinymce/plugins/link'
import 'tinymce/plugins/lists'
import 'tinymce/plugins/media'
import 'tinymce/plugins/nonbreaking'
import 'tinymce/plugins/pagebreak'
import 'tinymce/plugins/preview'
import 'tinymce/plugins/quickbars'
import 'tinymce/plugins/save'
import 'tinymce/plugins/searchreplace'
import 'tinymce/plugins/table'
import 'tinymce/plugins/visualblocks'
import 'tinymce/plugins/visualchars'
import 'tinymce/plugins/wordcount'
import 'tinymce/plugins/print'
import 'tinymce/plugins/template'
import 'tinymce/plugins/hr'
import 'tinymce/plugins/imagetools'
import 'tinymce/plugins/textpattern'


import { upload } from '@/api/upload'
// import useSettingsStore from '@/store/modules/settings'

// const settingsStore = useSettingsStore()

const props = defineProps({
    modelValue: {
        type: String,
        default: ''
    },
    setting: {
        type: Object,
        default: () => { }
    },
    disabled: {
        type: Boolean,
        default: false
    }
})

const emit = defineEmits(['update:modelValue'])

const defaultSetting = ref({
    language_url: 'tinymce/langs/zh-Hans.js',
    language: 'zh-Hans',
    skin_url: 'tinymce/skins/ui/oxide',
    content_css: 'tinymce/skins/content/default/content.min.css',
    min_height: 650,
    max_height: 800,
    selector: '#text textarea',
    external_plugins: {
        'importword': 'tinymce/plugins/importword/plugin.min.js',
        'indent2em': 'tinymce/plugins/indent2em/plugin.min.js',
        'emoticons': 'tinymce/plugins/emoticons/plugin.min.js',
        'bdmap': 'tinymce/plugins/bdmap/plugin.min.js',
        'layout': 'tinymce/plugins/layout/plugin.min.js'
    },
    font_formats: '微软雅黑=Microsoft YaHei,Helvetica Neue,PingFang SC,sans-serif;苹果苹方=PingFang SC,Microsoft YaHei,sans-serif;宋体=simsun,serif',
    plugins: 'paste edit code layout importword print preview searchreplace autolink directionality visualblocks visualchars fullscreen image link media template codesample table charmap hr pagebreak nonbreaking anchor insertdatetime advlist lists wordcount imagetools textpattern help emoticons autosave bdmap autoresize importword',
    toolbar: 'undo redo code preview layout| forecolor backcolor bold italic underline strikethrough link anchor | alignleft aligncenter alignright indent2em lineheight | \
                    formatselect fontselect fontsizeselect | bullist numlist | blockquote subscript superscript removeformat | \
                    table image media bdmap emoticons charmap hr pagebreak insertdatetime importword| code fullscreen ',
    branding: false,
    // menubar: 'file edit insert view format table tools help',
    toolbar_mode: 'sliding',
    promotion: false,
    insertdatetime_formats: [
        '%Y年%m月%d日',
        '%H点%M分%S秒',
        '%Y-%m-%d',
        '%H:%M:%S'
    ],
    paste_data_images: true,
    image_prepend_url: '/api/',
    images_upload_handler: (blobInfo, success, failure) => {
        var formData;
        formData = new FormData();
        formData.append("file", blobInfo.blob(), blobInfo.filename());
        upload(formData).then(res => {
            console.log(res.data)
            success(res.data.file.url)
        }).catch(err => {
            failure(err)
        })
    },
})
// console.log(blobInfo)



const myValue = ref(props.modelValue)

const completeSetting = computed(() => {
    return Object.assign(defaultSetting.value, props.setting)
})

watch(() => myValue.value, newValue => {
    emit('update:modelValue', newValue)
})

watch(() => props.modelValue, newValue => {
    myValue.value = newValue
})

onMounted(() => {
    tinymce.init({})
})
</script>

<template>
    <div class="editor">
        <TinymceEditor v-model="myValue" :init="completeSetting" :disabled="disabled" />
    </div>
</template>

<style lang="scss" scoped>
:deep(.tox-tinymce) {
    border-radius: 10px;
}
</style>
