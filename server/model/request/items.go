package request

import "github.com/helix/model"

type ItemsSearch struct{
    model.Items
    PageInfo
}