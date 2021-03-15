package request

import "gin-vue-admin/model"

type ActivitySearch struct{
    model.Activity
    PageInfo
}