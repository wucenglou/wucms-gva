import service from '@/utils/request'

// @Tags Pkg
// @Summary 创建Pkg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pkg true "创建Pkg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /doctor/createPkg [post]
export const createDoctor = (data) => {
    return service({
        url: '/doctor/createDoctor',
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
// @Router /doctor/deletePkg [delete]
export const deleteDoctor = (data) => {
    return service({
        url: '/doctor/deleteDoctor',
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
// @Router /doctor/deletePkg [delete]
export const deleteDoctorByIds = (data) => {
    return service({
        url: '/doctor/deleteDoctorByIds',
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
// @Router /doctor/updatePkg [put]
export const updateDoctor = (data) => {
    return service({
        url: '/doctor/updateDoctor',
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
// @Router /doctor/findPkg [get]
export const findDoctor = (params) => {
    return service({
        url: '/doctor/findDoctor',
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
// @Router /doctor/getPkgList [get]
export const getDoctorList = (params) => {
    return service({
        url: '/doctor/getDoctorList',
        method: 'get',
        params
    })
}
