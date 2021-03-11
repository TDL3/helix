package request

import "gin-vue-admin/model"

type LostItemsSearch struct{
    model.LostItems
    PageInfo
}