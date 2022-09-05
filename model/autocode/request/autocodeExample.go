// 自动生成模板SysDictionaryDetail
package request

import (
	"github.com/lhxlnsy/pap-go-server/model/autocode"
	"github.com/lhxlnsy/pap-go-server/model/common/request"
)

// 如果含有time.Time 请自行import time包
type AutoCodeExampleSearch struct {
	autocode.AutoCodeExample
	request.PageInfo
}
