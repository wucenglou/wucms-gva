import service from '@/utils/request'

export const createCmsModel = (data) => {
    return service({
        url: '/catModel',
        method: 'post',
        data: data
    })
}

export const updateCmsModel = (data) => {
    return service({
        url: '/catModel',
        method: 'put',
        data: data
    })
}

export const deleteCmsModel = (data) => {
    return service({
        url: '/catModel',
        method: 'delete',
        data: data
    })
}

export const deleteCmsModelByIds = (data) => {
    return service({
        url: '/catModel/ByIds',
        method: 'delete',
        data: data
    })
}

export const findCmsModel = (params) => {
    return service({
        url: '/catModel/ById',
        method: 'get',
        params
    })
}

export const getCmsModelList = (params) => {
    console.log(params)
    return service({
        url: '/catModel',
        method: 'get',
        params
    })
}