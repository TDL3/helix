package response

import (
	"gin-vue-admin/global"
)

type UserScoreInfo struct {
	global.GVA_MODEL
	HeaderImg   string       `json:"headerImg"`
	Score       int          `json:"score"`
	Username    string       `json:"userName"`
	NickName    string       `json:"nickName"`
}