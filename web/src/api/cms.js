import service from '@/utils/request'

export const getCatList = (data) => {
    return service({
        url: '/catList',
        method: 'get',
        data: data
    })
}

export const postCatMenu = (data) => {
    return service({
        url: '/catmenu',
        method: 'post',
        data: data
    })
}

export const editCatMenu = (data) => {
    return service({
        url: '/catmenu',
        method: 'put',
        data: data
    })
}

export const deleteCatMenu = (data) => {
    return service({
        url: '/catmenu',
        method: 'delete',
        data: data
    })
}