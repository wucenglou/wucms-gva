import service from '@/utils/request'

// @Tags Pkg
// @Summary 创建Pkg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pkg true "创建Pkg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /reg/createPkg [post]
export const createReg = (data) => {
    return service({
        url: '/reg/createReg',
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
// @Router /reg/deletePkg [delete]
export const deleteReg = (data) => {
    return service({
        url: '/reg/deleteReg',
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
// @Router /reg/deletePkg [delete]
export const deleteRegByIds = (data) => {
    return service({
        url: '/reg/deleteRegByIds',
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
// @Router /reg/updatePkg [put]
export const updateReg = (data) => {
    return service({
        url: '/reg/updateReg',
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
// @Router /reg/findPkg [get]
export const findReg = (params) => {
    return service({
        url: '/reg/findReg',
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
// @Router /reg/getPkgList [get]
export const getRegList = (params) => {
    return service({
        url: '/reg/getRegList',
        method: 'get',
        params
    })
}
