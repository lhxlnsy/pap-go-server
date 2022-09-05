package service

import (
	"github.com/lhxlnsy/pap-go-server/service/autocode"
	"github.com/lhxlnsy/pap-go-server/service/example"
	"github.com/lhxlnsy/pap-go-server/service/pap"
	"github.com/lhxlnsy/pap-go-server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
	PAPServiceGroup      pap.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
