package request

import "gin-vue-admin/model"

type ActivitiesSearch struct{
    model.Activities
    PageInfo
}