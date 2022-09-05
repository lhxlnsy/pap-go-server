package example

import "github.com/lhxlnsy/pap-go-server/model/system"

type ExcelInfo struct {
	FileName string               `json:"fileName"` // 文件名
	InfoList []system.SysBaseMenu `json:"infoList"`
}
