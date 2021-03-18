import service from '@/utils/request'

// @Tags ActivitiesManagement
// @Summary 创建ActivitiesManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActivitiesManagement true "创建ActivitiesManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /acm/createActivitiesManagement [post]
export const createActivitiesManagement = (data) => {
     return service({
         url: "/acm/createActivitiesManagement",
         method: 'post',
         data
     })
 }


// @Tags ActivitiesManagement
// @Summary 删除ActivitiesManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActivitiesManagement true "删除ActivitiesManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /acm/deleteActivitiesManagement [delete]
 export const deleteActivitiesManagement = (data) => {
     return service({
         url: "/acm/deleteActivitiesManagement",
         method: 'delete',
         data
     })
 }

// @Tags ActivitiesManagement
// @Summary 删除ActivitiesManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ActivitiesManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /acm/deleteActivitiesManagement [delete]
 export const deleteActivitiesManagementByIds = (data) => {
     return service({
         url: "/acm/deleteActivitiesManagementByIds",
         method: 'delete',
         data
     })
 }

// @Tags ActivitiesManagement
// @Summary 更新ActivitiesManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActivitiesManagement true "更新ActivitiesManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /acm/updateActivitiesManagement [put]
 export const updateActivitiesManagement = (data) => {
     return service({
         url: "/acm/updateActivitiesManagement",
         method: 'put',
         data
     })
 }


// @Tags ActivitiesManagement
// @Summary 用id查询ActivitiesManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActivitiesManagement true "用id查询ActivitiesManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /acm/findActivitiesManagement [get]
 export const findActivitiesManagement = (params) => {
     return service({
         url: "/acm/findActivitiesManagement",
         method: 'get',
         params
     })
 }


// @Tags ActivitiesManagement
// @Summary 分页获取ActivitiesManagement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取ActivitiesManagement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /acm/getActivitiesManagementList [get]
 export const getActivitiesManagementList = (params) => {
     return service({
         url: "/acm/getActivitiesManagementList",
         method: 'get',
         params
     })
 }




export const getAttendedActivitiesInfoList = (params) => {
    return service({
        url: "/acm/getAttendedActivitiesInfoList",
        method: 'get',
        params
    })
}
