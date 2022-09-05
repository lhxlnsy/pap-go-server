package request

import (
	"github.com/lhxlnsy/pap-go-server/model/common/request"
	"github.com/lhxlnsy/pap-go-server/model/pap"
)

type SiteInformationSearch struct {
	pap.SiteInformation
	request.PageInfo
}

type AlarmInformationSearch struct {
	SITEID string `json:"siteId" form:"siteId"`
	pap.AlarmInformation
	request.PageInfo
}

type HistroicalDataItem struct {
	MeterName     string `json:"meterName" form:"meterName"`
	MeterFunction string `json:"meterFunction" form:"meterFunction"`
	MeterSun      string `json:"meterSun" form:"meterSun"`
}

type HistoricalDataSearchRequest struct {
	JsonBody string `json:"jsonBody" form:"jsonBody"`
}

type HistoricalDataSearch struct {
	request.PageInfo
	SiteCode   string               `json:"siteCode" form:"siteCode"`
	TimeStart  string               `json:"timeStart" form:"timeStart"`
	TimeEnd    string               `json:"timeEnd" form:"timeEnd"`
	TimeBucket string               `json:"timeBucket" form:"timeBucket"`
	Device     string               `json:"device" form:"device"`
	MeterList  []HistroicalDataItem `json:"meterList" form:"meterList"`
}
