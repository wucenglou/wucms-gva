<template>
    <div>
        <!-- <div class="gva-search-box">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
                <el-form-item label="创建时间">
                    <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
                    —
                    <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
                </el-form-item>
                <el-form-item>
                    <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
                    <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
                </el-form-item>
            </el-form>
        </div> -->
        <div class="gva-table-box">
            <div class="gva-btn-list">
                <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
                <el-popover v-model:visible="deleteVisible" placement="top" width="160">
                    <p>确定要删除吗？</p>
                    <div style="text-align: right; margin-top: 8px;">
                        <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
                        <el-button size="small" type="primary" @click="onDelete">确定</el-button>
                    </div>
                    <template #reference>
                        <el-button icon="delete" size="small" style="margin-left: 10px;"
                            :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
                    </template>
                </el-popover>
            </div>
            <el-table style="width: 100%" tooltip-effect="dark" default-expand-all :data="tableData" row-key="term_id"
                @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55" />
                <el-table-column align="left" label="ID" prop="Term.term_id" width="120" />
                <el-table-column align="left" label="名称" prop="Term.name" width="120" />
                <el-table-column align="left" label="别名" prop="Term.slug" width="120" />
                <el-table-column align="left" label="描述" prop="description" show-overflow-tooltip width="320" />
                <el-table-column align="left" label="总数" prop="count" width="120" />
                <el-table-column align="left" label="对象分组" prop="Term.term_group" width="120" />
                <el-table-column align="left" label="按钮组" width="200">
                    <template #default="scope">
                        <el-button type="primary" link icon="edit" size="small" class="table-button"
                            @click="updateTermStructFunc(scope.row)">变更</el-button>
                        <el-button type="primary" link icon="delete" size="small"
                            @click="deleteRow(scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <!-- <div class="gva-pagination">
                <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
                    :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
                    @size-change="handleSizeChange" />
            </div> -->
        </div>
        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="添加分类">
            <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">

                <el-form-item label="上级栏目">
                    <el-cascader v-model="formData.term_taxonomy.parent_id" placeholder="无" :disabled="dialogType == 'add'"
                        :options="catOption"
                        :props="{ checkStrictly: true, label: 'name', value: 'TermTaxonomyId', disabled: 'disabled', emitPath: false }"
                        :show-all-levels="false" filterable />
                </el-form-item>
                <el-form-item label="名称:" prop="name">
                    <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
                </el-form-item>
                <el-form-item label="别名:" prop="slug">
                    <el-input v-model="formData.slug" :clearable="true" placeholder="请输入" />
                </el-form-item>
                <el-form-item label="描述:" prop="description">
                    <el-input v-model="formData.term_taxonomy.description" type="textarea" :clearable="true"
                        placeholder="请输入" />
                </el-form-item>
                <el-form-item label="图标" label-width="80px">
                    <div style="display:inline-block" @click="openHeaderChange1">
                        <img v-if="formData.term_taxonomy.headerImg" alt="图标" class="header-img-box"
                            :src="(formData.term_taxonomy.headerImg && formData.term_taxonomy.headerImg.slice(0, 4) !== 'http') ? path + formData.term_taxonomy.headerImg : formData.term_taxonomy.headerImg">
                        <div v-else class="header-img-box">从媒体库选择</div>
                    </div>
                </el-form-item>
                <el-form-item label="栏目配图" label-width="80px">
                    <div style="display:inline-block" @click="openHeaderChange2">
                        <img v-if="formData.term_taxonomy.descImg" alt="栏目配图" class="header-img-box2"
                            :src="(formData.term_taxonomy.descImg && formData.term_taxonomy.descImg.slice(0, 4) !== 'http') ? path + formData.term_taxonomy.descImg : formData.term_taxonomy.descImg">
                        <div v-else class="header-img-box">从媒体库选择</div>
                    </div>
                </el-form-item>

                <el-form-item label="对象分组:" prop="term_group">
                    <el-input v-model.number="formData.term_group" :clearable="true" placeholder="请输入" />
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button size="small" @click="closeDialog">取 消</el-button>
                    <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
                </div>
            </template>
        </el-dialog>
        <ChooseImg ref="chooseImg1" :target="formData.term_taxonomy" :target-key="`headerImg`" />
        <ChooseImg ref="chooseImg2" :target="formData.term_taxonomy" :target-key="`descImg`" />
    </div>
</template>

<script>
export default {
    name: 'cmsCategory'
}
</script>

<script setup>
import {
    createCmsCat,
    updateCmsCat,
    deleteCmsCat,
    deleteCmsCatByIds,
    findCmsCat,
    getCmsCatList
} from '@/api/cms'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, watch, onActivated, onDeactivated } from 'vue'
import { useRoute } from 'vue-router';
import ChooseImg from '@/components/chooseImg/index.vue'

const path = ref(import.meta.env.VITE_BASE_API + '/')

const route = useRoute()

console.log("99999999999999999999")
// 自动化生成的字典（可能为空）以及字段
const formData = ref({
    name: '',
    slug: '',
    term_group: 0,
    term_taxonomy: {
        "taxonomy": "cat",
        "description": "",
        "parent_id": 0,
        "headerImg": '',
        "descImg": '',
    },
})

// 验证规则
const rule = reactive({
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
    searchInfo.value = {}
    getTableData()
}

// 搜索
const onSubmit = () => {
    page.value = 1
    pageSize.value = 10
    getTableData()
}

// 分页
const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
    page.value = val
    // getTableData()
}

watch(
    () => route.params,
    (a, n) => {
        if (a.model) {
            console.log("此为模块")
            getTableData()
        }
    }
)
watch(
    () => route.path,
    (a, n) => {
        if (a.model) {
            console.log("此为模块")
            getTableData()
        }
    }
)

const chooseImg1 = ref(null)
const chooseImg2 = ref(null)
const openHeaderChange1 = () => {
    chooseImg1.value.open()
}
const openHeaderChange2 = () => {
    chooseImg2.value.open()
}

// 查询
const catOption = ref([])
const getTableData = async () => {
    console.log("调用一次")
    searchInfo.value.model = route.params.model
    const table = await getCmsCatList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        handleName(table.data.list)
        tableData.value = table.data.list.slice()
        catOption.value = table.data.list.slice()
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
getTableData()

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

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        deleteTermStructFunc(row)
    })
}


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async () => {
    const ids = []
    if (multipleSelection.value.length === 0) {
        ElMessage({
            type: 'warning',
            message: '请选择要删除的数据'
        })
        return
    }
    multipleSelection.value &&
        multipleSelection.value.map(item => {
            ids.push(item.term_id)
        })
    console.log("-------------")
    console.log(ids)
    const res = await deleteCmsCatByIds({ ids })
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
            page.value--
        }
        deleteVisible.value = false
        getTableData()
    }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateTermStructFunc = async (row) => {
    console.log("----------------")
    console.log(row)
    const res = await findCmsCat({ id: row.term_id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.term
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteTermStructFunc = async (row) => {
    const res = await deleteCmsCat({ ID: row.term_id })
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: '删除成功'
        })
        if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        name: '',
        slug: '',
        term_group: 0,
        term_taxonomy: {
            "taxonomy": "",
            "description": "",
            "parent_id": 0,
            "headerImg": '',
            "descImg": '',
        },
    }
}
// 弹窗确定
const enterDialog = async () => {
    elFormRef.value?.validate(async (valid) => {
        if (!valid) return
        let res
        switch (type.value) {
            case 'create':
                console.log(formData.value)
                formData.value.term_taxonomy.taxonomy = route.params.model
                res = await createCmsCat(formData.value)
                break
            case 'update':
                res = await updateCmsCat(formData.value)
                break
            default:
                res = await createTermStruct(formData.value)
                break
        }
        console.log(res)
        if (res.code === 0) {
            console.log("----------------------")
            ElMessage({
                type: 'success',
                message: '创建/更改成功'
            })
            closeDialog()
            getTableData()
        }
    })
}

</script>

<style>

 .header-img-box {
  width: 160px;
  height: 160px;
  border: 1px dashed #ccc;
  border-radius: 20px;
  text-align: center;
  line-height: 200px;
  cursor: pointer;
}
 .header-img-box2 {
  width: 300px;
  height: 200px;
  border: 1px dashed #ccc;
  border-radius: 20px;
  text-align: center;
  line-height: 200px;
  cursor: pointer;
}
</style>
