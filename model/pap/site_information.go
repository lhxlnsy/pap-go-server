// 自动生成模板SiteInformation
package pap

import "time"

// SiteInformation 结构体
// 如果含有time.Time 请自行import time包
type SiteInformation struct {
	ID           uint   `json:"id" form:"id" gorm:"column:id;primarykey"` // 主键ID
	SiteCode     string `json:"siteCode" form:"siteCode" gorm:"column:site_code;comment:;"`
	SiteName     string `json:"siteName" form:"siteName" gorm:"column:site_name;comment:;"`
	SiteLocation string `json:"siteLocation" form:"siteLocation" gorm:"column:site_location;comment:;"`
	SiteLat      string `json:"siteLat" form:"siteLat" gorm:"column:site_lat;comment:;"`
	SiteLng      string `json:"siteLng" form:"siteLng" gorm:"column:site_lng;comment:;"`
}

type ErrorStieInformation struct {
	TOTALERROR      int    `json:"totalError" form:"totalError" gorm:"column:total_error;comment:;"`
	LATESTTIMESTAMP string `json:"latestTimeStamp" form:"latestTimeStamp" gorm:"column:latest_timestamp;comment:;"`
	SiteName        string `json:"siteName" form:"siteName" gorm:"column:site_name;comment:;"`
	ErrorSite       string `json:"errorSite" form:"errorSite" gorm:"column:error_site;comment:;"`
	ErrorSource     string `json:"errorsource" form:"errorSource" gorm:"column:error_source;comment:'"`
	SiteLat         string `json:"siteLat" form:"siteLat" gorm:"column:site_lat;comment:;"`
	SiteLng         string `json:"siteLng" form:"siteLng" gorm:"column:site_lng;comment:;"`
}

type AlarmInformation struct {
	ID             uint   `json:"id" form:"id" gorm:"column:id;primarykey"` // 主键ID
	ERRORTIMESTAMP string `json:"errorTimeStamp" form:"errorTimeStamp" gorm:"column:error_timestamp;comment:;"`
	ERRORSITE      string `json:"errorSite" form:"errorSite" gorm:"column:error_site;comment:;"`
	ERRORCODE      string `json:"errorCode" form:"errorCode" gorm:"column:error_code;comment:;"`
	ERRORDESC      string `json:"errorDesc" form:"errorDesc" gorm:"column:error_desc;comment:;"`
	ERRORSTATUS    bool   `json:"errorStatus" form:"errorStatus" gorm:"column:error_status;comment:;"`
	SITENAME       string `json:"siteName" form:"siteName" gorm:"column:site_name;comment:;"`
	SITELOCATION   string `json:"siteLocation" form:"siteLocation" gorm:"column:site_location;comment:;"`
	TOTALERROR     int64  `json:"totalError" form:"totalError" gorm:"column:total_error;comment:;"`
}

type SiteInfo struct {
	SiteCode string `form:"siteCode"`
}

type SiteAlarmDataStatus struct {
	TOTALERROR int64  `json:"totalError" form:"totalError" gorm:"column:total_error;comment:;"`
	ERRORSITE  string `json:"errorSite" form:"errorSite" gorm:"column:error_site;comment:;"`
	SITENUMBER string `json:"siteNumber" form:"siteNumber" gorm:"column:site_number;comment:;"`
}

type SiteSummary struct {
	DateAndTime  time.Time `json:"Data and Time" gorm:"colmun:date_and_time"`
	SiteCode     string    `json:"Site Code" gorm:"column:site_code"`
	DeviceName   string    `json:"Device Name" gorm:"column:device_name"`
	InstACurrnt  float32   `json:"Inst A-Phase Current" gorm:"column:inst_a_current"`
	InstAVoltage float32   `json:"Inst A-Phase Voltage" gorm:"column:int_a_voltage"`
}

// TableName SiteInformation 表名
func (SiteInformation) TableName() string {
	return "site_information"
}

type RedisSiteInformation struct {
	SiteID     string `json:"siteID" form:"siteID"`
	SiteAlarm  []RedisSiteAlarm
	SiteValue  []RedisSiteValue
	SiteStatus []RedisSiteStatus
	StieCMD    []RedisSiteCMD
	SiteIPC    []RedisStieIPC
	SiteSA     []RedisSiteSA
	SiteSolar  []RedisSiteSolar
}
type RedisSiteValue struct {
	TimeStamp  string `json:"timeStamp" form:"timeStamp"`
	Value      string `json:"value" form:"value"`
	DeviceName string `json:"deviceName" form:"deviceName"`
}
type RedisSiteAlarm struct {
	TimeStamp string `json:"timeStamp" form:"timeStamp"`
	Value     string `json:"value" form:"value"`
	Device    string `json:"device" form:"device"`
}
type RedisSiteStatus struct {
	TimeStamp string `json:"timeStamp" form:"timeStamp"`
}
type RedisSiteCMD struct {
	TimeStamp string `json:"timeStamp" form:"timeStamp"`
	Value     string `json:"value" form:"value"`
}
type RedisStieIPC struct {
	TimeStamp string `json:"timeStamp" form:"timeStamp"`
	Value     string `json:"value" form:"value"`
	Device    string `json:"device" form:"device"`
}
type RedisSiteSA struct {
	TimeStamp    string    `json:"timeStamp" form:"timeStamp"`
	Value        []float32 `json:"value" form:"value"`
	IP           string    `json:"ip" form:"ip"`
	Register     string    `json:"register" form:"register"`
	DeviceTypeCD string    `json:"deviceTypeCD" form:"deviceTypeCD"`
}
type RedisSiteSolar struct {
	TimeStamp string    `json:"timeStamp" form:"timeStamp"`
	Value     []float32 `json:"value" form:"value"`
	Device    string    `json:"device" form:"device"`
}
