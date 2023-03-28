import service from '@/utils/request'

export const createpost = (data) => {
    return service({
        url: '/post',
        method: 'post',
        data: data
    })
}

export const updatepost = (data) => {
    return service({
        url: '/post',
        method: 'put',
        data: data
    })
}

export const deletepost = (data) => {
    return service({
        url: '/post',
        method: 'delete',
        data: data
    })
}

export const deletepostByIds = (data) => {
    return service({
        url: '/post/ByIds',
        method: 'delete',
        data: data
    })
}

export const findpost = (params) => {
    return service({
        url: '/post/ById',
        method: 'get',
        params
    })
}

export const getpostList = (params) => {
    console.log(params)
    return service({
        url: '/post',
        method: 'get',
        params
    })
}