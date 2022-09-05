package v1

import (
	"github.com/lhxlnsy/pap-go-server/api/v1/autocode"
	"github.com/lhxlnsy/pap-go-server/api/v1/example"
	"github.com/lhxlnsy/pap-go-server/api/v1/pap"
	"github.com/lhxlnsy/pap-go-server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
	PAPApiGroup      pap.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
