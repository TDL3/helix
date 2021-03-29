package response

import "github.com/helix/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
