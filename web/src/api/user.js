import service from '@/utils/request'

// This will get current looged in user's info, for now ony uuid is returned from server
export const getUserInfo = (params) => {
    return service({
        url: "/user/getUserInfo",
        method: 'get',
        params
    })
}

// This will get all students score list
export const getStudentUserScoreList = (data) => {
    return service({
        url: "/user/getStudentUserScoreList",
        method: 'post',
        data: data
    })
}

// This will get all students score list
export const getUnionScoreList = (data) => {
    return service({
        url: "/user/getUnionScoreList",
        method: 'post',
        data: data
    })
}


// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const login = (data) => {
    return service({
        url: "/base/login",
        method: 'post',
        data: data
    })
}

// @Summary 获取验证码
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/captcha [post]
export const captcha = (data) => {
    return service({
        url: "/base/captcha",
        method: 'post',
        data: data
    })
}

// @Summary 用户注册
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/resige [post]
export const register = (data) => {
        return service({
            url: "/user/register",
            method: 'post',
            data: data
        })
    }
    // @Summary 修改密码
    // @Produce  application/json
    // @Param data body {username:"string",password:"string",newPassword:"string"}
    // @Router /user/changePassword [post]
export const changePassword = (data) => {
        return service({
            url: "/user/changePassword",
            method: 'post',
            data: data
        })
    }
    // @Tags User
    // @Summary 分页获取用户列表
    // @Security ApiKeyAuth
    // @accept application/json
    // @Produce application/json
    // @Param data body modelInterface.PageInfo true "分页获取用户列表"
    // @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
    // @Router /user/getUserList [post]
export const getUserList = (data) => {
    return service({
        url: "/user/getUserList",
        method: 'post',
        data: data
    })
}


// @Tags User
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.SetUserAuth true "设置用户权限"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
export const setUserAuthority = (data) => {
    return service({
        url: "/user/setUserAuthority",
        method: 'post',
        data: data
    })
}


// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/deleteUser [delete]
export const deleteUser = (data) => {
    return service({
        url: "/user/deleteUser",
        method: 'delete',
        data: data
    })
}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "设置用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserInfo [put]
export const setUserInfo = (data) => {
    return service({
        url: "/user/setUserInfo",
        method: 'put',
        data: data
    })
}
