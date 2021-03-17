package response

type UserInfo struct {
	UUID        string    `json:"uuid"`
	NickName    string    `json:"nick_name"`
	AuthorityID string    `json:"authority_id"`
}