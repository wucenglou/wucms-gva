import service from '@/utils/request'
// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const upload = (data) => {
    return service({
        url: '/uploadimgs',
        method: 'post',
        data: data
    })
}