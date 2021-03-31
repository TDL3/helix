package response

import (
	"gin-vue-admin/global"
)

type UnionScoreInfo struct {
	global.GVA_MODEL
	Score         int          `json:"score"`
	ReqUnion      int          `json:"reqUnion"`
}