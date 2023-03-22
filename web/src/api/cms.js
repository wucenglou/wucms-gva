import service from '@/utils/request'

export const createCmsCat = (data) => {
    return service({
        url: '/cat',
        method: 'post',
        data: data
    })
}

export const updateCmsCat = (data) => {
    return service({
        url: '/cat',
        method: 'put',
        data: data
    })
}

export const deleteCmsCat = (data) => {
    return service({
        url: '/cat',
        method: 'delete',
        data: data
    })
}

export const deleteCmsCatByIds = (data) => {
    return service({
        url: '/cat/ByIds',
        method: 'delete',
        data: data
    })
}

export const findCmsCat = (params) => {
    return service({
        url: '/cat/ById',
        method: 'get',
        params
    })
}

export const getCmsCatList = (params) => {
    console.log(params)
    return service({
        url: '/cat',
        method: 'get',
        params
    })
}