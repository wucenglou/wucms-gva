import service from '@/utils/request'

// @Tags Pkg
// @Summary 创建Pkg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pkg true "创建Pkg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /patient/createPkg [post]
export const createPatient = (data) => {
    return service({
        url: '/patient/createPatient',
        method: 'post',
        data
    })
}

// @Tags Pkg
// @Summary 删除Pkg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pkg true "删除Pkg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /patient/deletePkg [delete]
export const deletePatient = (data) => {
    return service({
        url: '/patient/deletePatient',
        method: 'delete',
        data
    })
}

// @Tags Pkg
// @Summary 删除Pkg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Pkg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /patient/deletePkg [delete]
export const deletePatientByIds = (data) => {
    return service({
        url: '/patient/deletePatientByIds',
        method: 'delete',
        data
    })
}

// @Tags Pkg
// @Summary 更新Pkg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pkg true "更新Pkg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /patient/updatePkg [put]
export const updatePatient = (data) => {
    return service({
        url: '/patient/updatePatient',
        method: 'put',
        data
    })
}

// @Tags Pkg
// @Summary 用id查询Pkg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Pkg true "用id查询Pkg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /patient/findPkg [get]
export const findPatient = (params) => {
    return service({
        url: '/patient/findPatient',
        method: 'get',
        params
    })
}

// @Tags Pkg
// @Summary 分页获取Pkg列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Pkg列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /patient/getPkgList [get]
export const getPatientList = (params) => {
    return service({
        url: '/patient/getPatientList',
        method: 'get',
        params
    })
}
