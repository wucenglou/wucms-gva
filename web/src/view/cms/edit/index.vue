<template>
    <div>
        <div class="gva-search-box">

            <el-form ref="Form" :model="formData" :rules="rules" label-width="5em">

                <el-form-item label="栏目" prop="term_id">
                    <el-cascader v-model="formData.term_id" placeholder="无" :disabled="dialogType == 'add'"
                        :options="catOption"
                        :props="{ multiple: false, checkStrictly: true, label: 'name', value: 'TermTaxonomyId', disabled: 'disabled', emitPath: false }"
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
                <el-row :gutter="24">
                    <el-col :xs="24" :sm="12">
                        <el-form-item label="权重">
                            <el-input-number v-model="formData.menu_order" :precision="2" :step="0.01" :max="10" />
                        </el-form-item>
                        <el-form-item label="状态" prop="post_status">
                            <el-radio-group v-model="formData.post_status">
                                <el-radio label="publish">已发布</el-radio>
                                <el-radio label="Pending">待审核</el-radio>
                                <el-radio label="reject">未通过</el-radio>
                                <el-radio label="Draft">草稿</el-radio>
                            </el-radio-group>
                        </el-form-item>

                    </el-col>
                    <el-col :xs="24" :sm="12">
                        <el-form-item label="文章配图" label-width="80px">
                            <div style="display:inline-block" @click="openHeaderChange">
                                <img v-if="formData.post_img" alt="配图" class="header-img-box"
                                    :src="(formData.post_img && formData.post_img.slice(0, 4) !== 'http') ? path + formData.post_img : formData.post_img">
                                <div v-else class="header-img-box">从媒体库选择</div>
                            </div>
                        </el-form-item>

                    </el-col>
                </el-row>
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
        <ChooseImg ref="chooseImg" :target="formData" :target-key="`post_img`" />
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
    createPost,
    updatePost,
    deletePost,
    deletePostByIds,
    findPost,
    getPostList
} from '@/api/cmsPost'
import {
    getCmsCatList
} from '@/api/cms'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useRoute } from 'vue-router';
import ChooseImg from '@/components/chooseImg/index.vue'

const path = ref(import.meta.env.VITE_BASE_API + '/')

const formData = ref({
    term_id: [],
    user_id: '',
    post_title: '',
    status: '1',
    post_excerpt: '',
    post_img: '',
    post_content: "",
    post_status: "publish",
    menu_order: 0,
    comment_status: "open"
})
const initForm = () => {
    formData.value = {
        term_id: [0],
        user_id: '',
        post_title: '',
        status: '1',
        post_excerpt: '',
        post_img: '',
        post_content: "",
        post_status: "publish",
        menu_order: 0,
        comment_status: "open"
    }
}

const type = ref('create')
// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const route = useRoute()
const model = ref('cat')
const post_id = ref(0)
if (route.params.value) {
    // 如果有参数通过params传递model，则为创建模式，
    // 否则是通过query传参，传过来model和post_id，则为修改模式
    type.value = "create"
    model.value = route.params.model
} else {
    if (route.query.post_id) {
        type.value = "update"
        post_id.value = route.query.post_id
    } else {
        type.value = "create"
    }
    model.value = route.query.model
}

// 查询
const catOption = ref([])
const getModel = async () => {
    console.log("调用一次")
    searchInfo.value.model = model
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
            name: '未分类',
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

const getPost = async () => {
    const post = await findPost({ "id": post_id.value })
    if (post.code === 0) {
        formData.value = post.data.post
        formData.value.term_id = post.data.post.termtaxonomy[0].term_id
    }
}
if (post_id.value) {
    console.log("hava")
    getPost()
} else {
    console.log("no")
}


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
    console.log(type.value)
    let res
    switch (type.value) {
        case 'create':
            console.log(formData.value)
            console.log(typeof formData.value.term_id)
            if (typeof formData.value.term_id == "number") {
                formData.value.term_id = [formData.value.term_id]
            }
            // formData.value.term_taxonomy.taxonomy = route.params.model
            res = await createPost(formData.value)
            console.log(res)
            break
        case 'update':
            res = await updatePost(formData.value)
            break
        default:
            console.log(formData.value)
            // res = await createTermStruct(formData.value)
            break
    }
    if (res.code === 0) {
        console.log("----------------------")
        ElMessage({
            type: 'success',
            message: '创建/更改成功'
        })

        initForm()
    }
}

// 图片上传
const chooseImg = ref(null)
const openHeaderChange = () => {
    chooseImg.value.open()
}

</script>

<style>
 .header-img-box {
  width: 300px;
  height: 200px;
  border: 1px dashed #ccc;
  border-radius: 20px;
  text-align: center;
  line-height: 200px;
  cursor: pointer;
}
</style>