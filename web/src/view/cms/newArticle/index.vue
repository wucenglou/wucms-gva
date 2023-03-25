<template>
    <div>
        <div class="gva-search-box">

            <el-form ref="Form" :model="form" :rules="rules" label-width="4em">


                <el-form-item label="标题" prop="title">
                    <el-input v-model="form.title"></el-input>
                </el-form-item>
                <el-form-item label="关键词" prop="keywords">
                    <el-input v-model="form.keywords"></el-input>
                </el-form-item>
                <el-form-item label="摘要" prop="description">
                    <el-input v-model="form.description" type="textarea"></el-input>
                </el-form-item>
                <!-- <el-form-item label="权重">
                    <el-input-number v-model="form.sort" :precision="2" :step="0.01" :max="10" />
                </el-form-item>
                <el-form-item label="状态" prop="status">
                    <el-radio-group v-model="form.status">
                        <el-radio label="1">已发布</el-radio>
                        <el-radio label="2">待审核</el-radio>
                        <el-radio label="3">未通过</el-radio>
                        <el-radio label="4">已废弃</el-radio>
                    </el-radio-group>
                </el-form-item> -->
                <el-row :gutter="24">
                    <el-col :xs="24" :sm="12">
                        <el-form-item label="权重">
                            <el-input-number v-model="form.sort" :precision="2" :step="0.01" :max="10" />
                        </el-form-item>

                    </el-col>
                    <el-col :xs="24" :sm="12">
                        <el-form-item label="状态" prop="status">
                            <el-radio-group v-model="form.status">
                                <el-radio label="1">已发布</el-radio>
                                <el-radio label="2">待审核</el-radio>
                                <el-radio label="3">未通过</el-radio>
                                <el-radio label="4">已废弃</el-radio>
                            </el-radio-group>
                        </el-form-item>

                    </el-col>
                </el-row>
                <!-- 
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="权重">
                            <el-input-number v-model="form.sort" :precision="2" :step="0.01" :max="10" />
                        </el-form-item>

                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="状态" prop="status">
                            <el-radio-group v-model="form.status">
                                <el-radio label="1">已发布</el-radio>
                                <el-radio label="2">待审核</el-radio>
                                <el-radio label="3">未通过</el-radio>
                                <el-radio label="4">已废弃</el-radio>
                            </el-radio-group>
                        </el-form-item>
                    </el-col>
                </el-row> -->
                <!-- <el-form-item> -->
                <div id="text" style="width: 100%">
                    <editor v-model="form.content" />
                </div>
                <!-- </el-form-item> -->

                <div style="margin: 2rem;">
                    <el-button>取消</el-button>
                    <el-button type="primary" @click="onSubmit('create')">提交</el-button>
                </div>
            </el-form>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Edit',
}
</script>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import editor from "@/components/Editor/index.vue"
const form = ref({
    catId: '',
    user_id: '',
    // modeId: '',
    title: '标题',
    status: '1',
    keywords: '',
    description: '',
    content: "",
    // flag: [],
    // type: [],
    picList: [],
    sort: 0,
    // page_views: '',
})
</script>