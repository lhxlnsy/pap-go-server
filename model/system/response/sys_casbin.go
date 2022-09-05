package response

import (
	"github.com/lhxlnsy/pap-go-server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
