package request

import "github.com/helix/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}