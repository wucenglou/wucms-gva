<template>
    <div>
        <div class="gva-search-box">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
                <el-form-item label="创建时间">
                    <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
                    —
                    <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
                </el-form-item>
                <el-form-item label="医生">
                    <el-input v-model="searchInfo.keyword" placeholder="搜索条件" />

                </el-form-item>
                <el-form-item>
                    <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
                    <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
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
            <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
                @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55" />

                <!-- <el-table-column align="left" label="头像" min-width="75">
                    <template #default="scope">
                        <CustomPic style="margin-top:8px;align-items: center;" :pic-src="scope.row.profile_img" />
                    </template>
                </el-table-column> -->

                <el-table-column align="left" label="姓名" prop="name" width="120" />
                <el-table-column align="left" label="头衔" prop="title" width="200" />
                <el-table-column align="left" label="专长" show-overflow-tooltip prop="specialty" width="250" />
                <el-table-column align="left" label="个人介绍" show-overflow-tooltip prop="desc" width="300" />
                <el-table-column align="left" label="手机号" prop="phone" width="120" />
                <el-table-column align="left" label="日期" width="180">
                    <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
                </el-table-column>
                <el-table-column align="left" label="按钮组" width="200">
                    <template #default="scope">
                        <el-button type="primary" link icon="edit" size="small" class="table-button"
                            @click="updateDoctorFunc(scope.row)">变更</el-button>
                        <el-button type="primary" link icon="delete" size="small"
                            @click="deleteRow(scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="gva-pagination">
                <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
                    :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
                    @size-change="handleSizeChange" />
            </div>
        </div>
        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
            <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
                <el-form-item label="医生:" prop="name">
                    <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
                </el-form-item>
                <el-form-item label="头衔:" prop="title">
                    <el-input v-model="formData.title" :clearable="true" placeholder="请输入" />
                </el-form-item>
                <el-form-item label="专长:" prop="specialty">
                    <el-input v-model="formData.specialty" type="textarea" :clearable="true" placeholder="请输入" />
                </el-form-item>
                <el-form-item label="个人介绍:" prop="desc">
                    <el-input v-model="formData.desc" type="textarea" :clearable="true" placeholder="请输入" />
                </el-form-item>
                <el-form-item label="手机号:" prop="phone">
                    <el-input v-model.number="formData.phone" :clearable="true" placeholder="请输入" />
                </el-form-item>
                <el-form-item label="医生头像" label-width="80px">
                    <div style="display:inline-block" @click="openHeaderChange1">
                        <img v-if="formData.profile_img" alt="图标" class="header-img-box"
                            :src="(formData.profile_img && formData.profile_img.slice(0, 4) !== 'http') ? path + formData.profile_img : formData.profile_img">
                        <div v-else class="header-img-box">从媒体库选择</div>
                    </div>
                </el-form-item>
                <el-form-item label="医生配图" label-width="80px">
                    <div style="display:inline-block" @click="openHeaderChange2">
                        <img v-if="formData.desc_img" alt="栏目配图" class="header-img-box2"
                            :src="(formData.desc_img && formData.desc_img.slice(0, 4) !== 'http') ? path + formData.desc_img : formData.desc_img">
                        <div v-else class="header-img-box">从媒体库选择</div>
                    </div>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button size="small" @click="closeDialog">取 消</el-button>
                    <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
                </div>
            </template>
        </el-dialog>
        <ChooseImg ref="chooseImg1" :target="formData" :target-key="`profile_img`" />
        <ChooseImg ref="chooseImg2" :target="formData" :target-key="`desc_img`" />
    </div>
</template>

<script>
export default {
    name: 'Doctor'
}
</script>

<script setup>
import {
    createDoctor,
    deleteDoctor,
    deleteDoctorByIds,
    updateDoctor,
    findDoctor,
    getDoctorList
} from '@/api/doctor'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import CustomPic from '@/components/customPic/index.vue'

const path = ref(import.meta.env.VITE_BASE_API + '/')

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
    name: '',
    phone: 0,
    desc: '',
    profile_img: '',
    desc_img: '',
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
    getTableData()
}

// 查询
const getTableData = async () => {
    const table = await getDoctorList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
    }
}

getTableData()

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
        deletePkgFunc(row)
    })
}

const chooseImg1 = ref(null)
const chooseImg2 = ref(null)
const openHeaderChange1 = () => {
    chooseImg1.value.open()
}
const openHeaderChange2 = () => {
    chooseImg2.value.open()
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
            ids.push(item.ID)
        })
    const res = await deleteDoctorByIds({ ids })
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
const updateDoctorFunc = async (row) => {
    const res = await findDoctor({ id: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.Doctor
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteDoctorFunc = async (row) => {
    const res = await deleteDoctor({ id: row.ID })
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
        phone: 0,
        desc: '',
    }
}
// 弹窗确定
const enterDialog = async () => {
    elFormRef.value?.validate(async (valid) => {
        if (!valid) return
        let res
        switch (type.value) {
            case 'create':
                res = await createDoctor(formData.value)
                break
            case 'update':
                res = await updateDoctor(formData.value)
                break
            default:
                res = await createDoctor(formData.value)
                break
        }
        if (res.code === 0) {
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

<style lang="scss">
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
.user-dialog {
    .header-img-box {
        width: 200px;
        height: 200px;
        border: 1px dashed #ccc;
        border-radius: 20px;
        text-align: center;
        line-height: 200px;
        cursor: pointer;
    }

    .avatar-uploader .el-upload:hover {
        border-color: #409eff;
    }

    .avatar-uploader-icon {
        border: 1px dashed #d9d9d9 !important;
        border-radius: 6px;
        font-size: 28px;
        color: #8c939d;
        width: 178px;
        height: 178px;
        line-height: 178px;
        text-align: center;
    }

    .avatar {
        width: 178px;
        height: 178px;
        display: block;
    }
}

.nickName {
    display: flex;
    justify-content: flex-start;
    align-items: center;
}

.pointer {
    cursor: pointer;
    font-size: 16px;
    margin-left: 2px;
}
</style>