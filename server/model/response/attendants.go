package response

import "gin-vue-admin/model"

type Attendants struct {
	User    model.SysUser `json:"user"`
	Present *bool         `json:"present"`
}
