package request

import "gin-vue-admin/model"

type ActivitiesManagementSearch struct{
    model.ActivitiesManagement
    PageInfo
}