package request

import (
	"github.com/lhxlnsy/pap-go-server/model/common/request"
	"github.com/lhxlnsy/pap-go-server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
