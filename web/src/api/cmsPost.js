import service from '@/utils/request'

export const createPost = (data) => {
    return service({
        url: '/post',
        method: 'post',
        data: data
    })
}

export const updatePost = (data) => {
    return service({
        url: '/post',
        method: 'put',
        data: data
    })
}

export const deletePost = (data) => {
    return service({
        url: '/post',
        method: 'delete',
        data: data
    })
}

export const deletePostByIds = (data) => {
    return service({
        url: '/post/ByIds',
        method: 'delete',
        data: data
    })
}

export const findPost = (params) => {
    return service({
        url: '/post/ById',
        method: 'get',
        params
    })
}

export const getPostList = (params) => {
    return service({
        url: '/post',
        method: 'get',
        params
    })
}