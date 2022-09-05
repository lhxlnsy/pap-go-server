package pap

import (
	"github.com/lhxlnsy/pap-go-server/global"
)

type ChartSaved struct {
	global.GVA_MODEL
	UserName         string `json:"userName" form:"userName" gorm:"commnet:用户名"`
	ChartUUID        string `json:"chartUUID" form:"chartUUID" gorm:"not null;unique;primary_key;commnet: Chart UUID"`
	SiteID           string `json:"siteID" form:"siteID" gorm:"comment:Site ID"`
	TimeBucket       string `json:"timeBucket" gorm:"commnet: Time Bucket"`
	Device           string `json:"device" gorm:"comment: Device"`
	MeterList        string `json:"meterList" gorm:"commnet: Serialized Meter List Array"`
	DateRange        string `json:"dataRange" gorm:"comment: Date Range"`
	ChartStyleOption string `json:"chartStyleOption" gorm:"commnet Serialized Style Option"`
}
