import service from '@/utils/request'

// This will get current looged in user's info, for now ony uuid is returned from server
export const getUserInfo = (params) => {
    return service({
        url: "/user/getUserInfo",
        method: 'get',
        params
    })
}

// This will only get current user's item list determined by uuid
export const getCurrentUserItemsList = (params) => {
    return service({
        url: "/itm/getCurrentUserItemsList",
        method: 'get',
        params
    })
}



// @Tags Items
// @Summary 创建Items
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Items true "创建Items"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /itm/createItems [post]
export const createItems = (data) => {
     return service({
         url: "/itm/createItems",
         method: 'post',
         data
     })
 }


// @Tags Items
// @Summary 删除Items
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Items true "删除Items"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /itm/deleteItems [delete]
 export const deleteItems = (data) => {
     return service({
         url: "/itm/deleteItems",
         method: 'delete',
         data
     })
 }

// @Tags Items
// @Summary 删除Items
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Items"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /itm/deleteItems [delete]
 export const deleteItemsByIds = (data) => {
     return service({
         url: "/itm/deleteItemsByIds",
         method: 'delete',
         data
     })
 }

// @Tags Items
// @Summary 更新Items
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Items true "更新Items"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /itm/updateItems [put]
 export const updateItems = (data) => {
     return service({
         url: "/itm/updateItems",
         method: 'put',
         data
     })
 }


// @Tags Items
// @Summary 用id查询Items
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Items true "用id查询Items"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /itm/findItems [get]
 export const findItems = (params) => {
     return service({
         url: "/itm/findItems",
         method: 'get',
         params
     })
 }


// @Tags Items
// @Summary 分页获取Items列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Items列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /itm/getItemsList [get]
 export const getItemsList = (params) => {
     return service({
         url: "/itm/getItemsList",
         method: 'get',
         params
     })
 }
