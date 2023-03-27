<template>
    <div>
        <div class="gva-search-box">

            <el-form ref="Form" :model="formData" :rules="rules" label-width="5em">

                <el-form-item label="栏目" prop="term_id">
                    <el-cascader v-model="formData.term_id" placeholder="无" :disabled="dialogType == 'add'"
                        :options="catOption"
                        :props="{ checkStrictly: true, label: 'name', value: 'TermTaxonomyId', disabled: 'disabled', emitPath: false }"
                        :show-all-levels="false" filterable />
                </el-form-item>

                <el-form-item label="标题" prop="post_title">
                    <el-input v-model="formData.post_title"></el-input>
                </el-form-item>
                <el-form-item label="关键词" prop="keywords">
                    <el-input v-model="formData.keywords"></el-input>
                </el-form-item>
                <el-form-item label="摘要" prop="post_excerpt">
                    <el-input v-model="formData.post_excerpt" type="textarea"></el-input>
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
                            <el-input-number v-model="formData.menu_order" :precision="2" :step="0.01" :max="10" />
                        </el-form-item>

                    </el-col>
                    <el-col :xs="24" :sm="12">
                        <el-form-item label="状态" prop="post_status">
                            <el-radio-group v-model="formData.post_status">
                                <el-radio label="publish">已发布</el-radio>
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
                <div id="text" style="width: 100%;border-radius: 1px;">
                    <editor v-model="formData.post_content" />
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
import {
    createCmsCat,
    updateCmsCat,
    deleteCmsCat,
    deleteCmsCatByIds,
    findCmsCat,
    getCmsCatList
} from '@/api/cms'
import { useRoute } from 'vue-router';

const route = useRoute()


const formData = ref({
    term_id: '',
    user_id: '',
    post_title: '标题',
    status: '1',
    post_excerpt: '',
    post_content: "",
    post_status: "",
    menu_order: 0,
    comment_status: "open"
})
// const model
const type = ref('create')
// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

console.log("777777777")
console.log(route)

// 查询
const catOption = ref([])
const getModel = async () => {
    console.log("调用一次")
    searchInfo.value.model = "model"
    const table = await getCmsCatList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        handleName(table.data.list)
        // table.data.list.forEach(res=>{
        //     if(res.Term.slug != "model"){
        //         catOption.value.push(res)
        //     }
        // })
        tableData.value = table.data.list.slice()
        catOption.value = table.data.list.slice()
        console.log("+++++++++++")
        console.log(catOption.value)
        console.log(table.data.list)

        catOption.value.unshift({
            term_id: 0,
            TermTaxonomyId: 0,
            name: '根目录',
            term_group: 0,
            term_taxonomy: {
                "taxonomy": "cat",
                "description": "P描述",
                "parent_id": 0
            }
        })
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
    }
}
getModel()
const handleName = (list) => {
    let arr = []
    list.forEach(item => {
        item.name = item.Term.name
        if (item.children) {
            handleName(item.children)
        }
    })
    return arr
}

const onSubmit = async () => {
    switch (type.value) {
        case 'create':
            console.log(formData.value)
            // formData.value.term_taxonomy.taxonomy = route.params.model
            // res = await createCmsCat(formData.value)
            break
        case 'update':
            console.log(formData.value)
            // res = await updateCmsCat(formData.value)
            break
        default:
            console.log(formData.value)
            // res = await createTermStruct(formData.value)
            break
    }
}

</script>