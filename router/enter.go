package router

import (
	"github.com/lhxlnsy/pap-go-server/router/autocode"
	"github.com/lhxlnsy/pap-go-server/router/example"
	"github.com/lhxlnsy/pap-go-server/router/pap"
	"github.com/lhxlnsy/pap-go-server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
	PAP      pap.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
