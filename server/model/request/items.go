package request

import "gin-vue-admin/model"

type ItemsSearch struct{
    model.Items
    PageInfo
}