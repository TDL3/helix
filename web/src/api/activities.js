import service from '@/utils/request'

// @Tags Activities
// @Summary 创建Activities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Activities true "创建Activities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /acti/createActivities [post]
export const createActivities = (data) => {
     return service({
         url: "/acti/createActivities",
         method: 'post',
         data
     })
 }


// @Tags Activities
// @Summary 删除Activities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Activities true "删除Activities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /acti/deleteActivities [delete]
 export const deleteActivities = (data) => {
     return service({
         url: "/acti/deleteActivities",
         method: 'delete',
         data
     })
 }

// @Tags Activities
// @Summary 删除Activities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Activities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /acti/deleteActivities [delete]
 export const deleteActivitiesByIds = (data) => {
     return service({
         url: "/acti/deleteActivitiesByIds",
         method: 'delete',
         data
     })
 }

// @Tags Activities
// @Summary 更新Activities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Activities true "更新Activities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /acti/updateActivities [put]
 export const updateActivities = (data) => {
     return service({
         url: "/acti/updateActivities",
         method: 'put',
         data
     })
 }


// @Tags Activities
// @Summary 用id查询Activities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Activities true "用id查询Activities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /acti/findActivities [get]
 export const findActivities = (params) => {
     return service({
         url: "/acti/findActivities",
         method: 'get',
         params
     })
 }


// @Tags Activities
// @Summary 分页获取Activities列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Activities列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /acti/getActivitiesList [get]
 export const getActivitiesList = (params) => {
     return service({
         url: "/acti/getActivitiesList",
         method: 'get',
         params
     })
 }