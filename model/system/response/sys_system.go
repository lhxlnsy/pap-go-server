package response

import "github.com/lhxlnsy/pap-go-server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
