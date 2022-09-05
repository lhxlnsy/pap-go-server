package request

import (
	"github.com/lhxlnsy/pap-go-server/model/autocode"
	"github.com/lhxlnsy/pap-go-server/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}