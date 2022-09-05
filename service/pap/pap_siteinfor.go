package pap

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/lhxlnsy/pap-go-server/global"
	"github.com/lhxlnsy/pap-go-server/model/common/request"
	"github.com/lhxlnsy/pap-go-server/model/pap"
	papReq "github.com/lhxlnsy/pap-go-server/model/pap/request"
)

type SiteInformationService struct {
}

func (siteInformationService *SiteInformationService) CreateSiteInformation(siteInformation pap.SiteInformation) (err error) {
	err = global.GVA_DB.Create(&siteInformation).Error
	return err
}

func (SiteInformationService *SiteInformationService) CreateChartSaved(chartInformation pap.ChartSaved) (err error) {
	err = global.GVA_DB.Create(&chartInformation).Error
	return err
}

func (siteInformationService *SiteInformationService) DeleteChartSavedByIds(id string) (err error) {
	err = global.GVA_DB.Delete(&pap.ChartSaved{}, "chart_uuid = ?", id).Error
	return err
}

// GetSiteInformation 根据id获取SiteInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteInformationService *SiteInformationService) GetChartSavedByUserId(chartInformation pap.ChartSaved) (err error, chartsSaved []pap.ChartSaved) {
	fmt.Println(chartInformation.UserName)
	fmt.Println(chartInformation.SiteID)
	fmt.Println(chartInformation)
	fmt.Println("GET CHARTS")
	err = global.GVA_DB.Where("user_name = ?", chartInformation.UserName).Where("site_id=?", chartInformation.SiteID).Find(&chartsSaved).Error
	return
}

// DeleteSiteInformation 删除SiteInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteInformationService *SiteInformationService) DeleteSiteInformation(siteInformation pap.SiteInformation) (err error) {
	err = global.GVA_DB.Delete(&siteInformation).Error
	return err
}

// DeleteSiteInformationByIds 批量删除SiteInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteInformationService *SiteInformationService) DeleteSiteInformationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]pap.SiteInformation{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSiteInformation 更新SiteInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteInformationService *SiteInformationService) UpdateSiteInformation(siteInformation pap.SiteInformation) (err error) {
	err = global.GVA_DB.Save(&siteInformation).Error
	return err
}

func (siteInformationService *SiteInformationService) UpdateChartSaved(chartInformation pap.ChartSaved) (err error) {
	var model pap.ChartSaved
	uuid := chartInformation.ChartUUID
	err = global.GVA_DB.Where("chart_uuid = ?", uuid).First(&model).Update("chart_style_option", chartInformation.ChartStyleOption).Error
	return err
}

// GetSiteInformation 根据id获取SiteInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteInformationService *SiteInformationService) GetSiteInformation(id uint) (err error, siteInformation pap.SiteInformation) {
	err = global.GVA_DB.Where("id = ?", id).First(&siteInformation).Error
	return
}
func (siteInformationService *SiteInformationService) GetAllSites(query papReq.SiteInformationSearch) (err error, list interface{}, total int64) {
	redis := global.GVA_REDIS
	ctx := context.Background()
	redis.Do(ctx, "SELECT", "0")
	item, err := redis.Get(ctx, "ModelRegisterDict").Result()
	result := map[string]map[string]interface{}{}
	json.Unmarshal([]byte(item), &result)
	limit := query.PageSize
	offset := query.PageSize * (query.Page - 1)
	// 创建db
	db := global.GetGlobalDBByDBName("planetarkpower").Model(&pap.SiteInformation{})
	var siteInformations []pap.SiteInformation
	var siteCodeInDatabase []string
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&siteInformations).Error
	for index, _ := range siteInformations {
		siteCodeInDatabase = append(siteCodeInDatabase, "EEMS_"+siteInformations[index].SiteCode)
	}
	for key := range siteCodeInDatabase {
		println(siteCodeInDatabase[key])
	}
	sort.Strings(siteCodeInDatabase)
	currentLen := len(siteInformations)
	for key := range result {
		i := sort.SearchStrings(siteCodeInDatabase, key)
		if !(i < len(siteCodeInDatabase) && siteCodeInDatabase[i] == key) {
			var tempresult pap.SiteInformation
			currentLen = currentLen + 1
			tempresult.ID = uint(currentLen)
			tempresult.SiteCode = key[5:]
			tempresult.SiteLat = "Unknown"
			tempresult.SiteLng = "Unknown"
			tempresult.SiteLocation = "Unknown"
			tempresult.SiteName = "Unknown"
			siteInformations = append(siteInformations, tempresult)
			total = total + 1
		}
	}
	sort.Slice(siteInformations, func(i, j int) bool {
		return siteInformations[i].SiteCode < siteInformations[j].SiteCode
	})
	// for key, value := range result {
	// 	for subkey, subvalue := range value {
	// 	}
	// }
	return err, siteInformations, total
}

// GetErrorInformation 获取最新的错误信息以及对应的site信息
// Author [JeffreyLi]
func (SiteInformationService *SiteInformationService) GetErrorInformation() (err error, list interface{}) {
	var item []pap.ErrorStieInformation
	err = global.GetGlobalDBByDBName("planetarkpower").Raw(`
	SELECT count(*) as total_error, max(error_timestamp) as latest_timestamp,site_name, error_site, site_lat,site_lng FROM error_code_status
	LEFT JOIN site_information
		ON site_information.site_code = TRIM('EEMS_' from error_site)
	where error_status = false
	GROUP BY error_code_status.error_site,site_name,site_lat,site_lng
	`).Scan(&item).Error
	return err, item
}

func (SiteInformationService *SiteInformationService) GetSiteAlarmDataStatus(id string) (err error, list interface{}) {
	var item []pap.SiteAlarmDataStatus
	sqlquery := fmt.Sprintf(`
	select
		count(*) as total_error,
		error_site
	from error_code_status
	where error_site='EEMS_%s' and error_code <> '0'
	group by error_site
	`, id)
	err = global.GetGlobalDBByDBName("planetarkpower").Raw(sqlquery).Scan(&item).Error
	return err, item
}

func (SiteInformationService *SiteInformationService) GetSiteDailySummary(id string) (err error, list interface{}) {
	redis := global.GVA_REDIS
	ctx := context.Background()
	redis.Do(ctx, "SELECT", "0")
	item, err := redis.Get(ctx, "ModelRegisterDict").Result()
	result := map[string]interface{}{}
	json.Unmarshal([]byte(item), &result)
	return err, Scan(result, "EEMS_"+id)
}

// GetSiteInformationInfoList 分页获取SiteInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteInformationService *SiteInformationService) GetSiteInformationInfoList(info papReq.SiteInformationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GetGlobalDBByDBName("planetarkpower").Model(&pap.SiteInformation{})
	var siteInformations []pap.SiteInformation
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&siteInformations).Error
	return err, siteInformations, total
}

func (SiteInformationService *SiteInformationService) SearchSiteAlarmDetail(info papReq.AlarmInformationSearch) (err error, list interface{}, total int64) {
	var item []pap.AlarmInformation
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	id := info.SITEID
	fmt.Println(limit)
	fmt.Println(offset)
	sqlquery := fmt.Sprintf(`
	SELECT 
		count(*) over() as total_error,
		error_code_status.id,
		error_code_status.error_timestamp, 
		error_code_status.error_site, 
		error_code_status.error_code,
		error_code_status.error_desc,
		error_code_status.error_status
	FROM error_code_status
	where error_code_status.error_site = 'EEMS_%s' and error_code_status.error_code <> '0'
	limit %d offset %d
	`, id, limit, offset)
	err = global.GetGlobalDBByDBName("planetarkpower").Raw(sqlquery).Scan(&item).Error
	total = item[0].TOTALERROR
	return err, item, total
}

func (SiteInformationService *SiteInformationService) GetSiteRedisData(id string) (err error, list interface{}) {
	redis := global.GVA_REDIS
	ctx := context.Background()
	redis.Do(ctx, "SELECT", "0")
	item, err := redis.Get(ctx, "ModelRegisterDict").Result()
	result := map[string]interface{}{}
	json.Unmarshal([]byte(item), &result)
	return err, Scan(result, "EEMS_"+id)
}

func Scan(m map[string]interface{}, target string) (v interface{}) {
	for k, v := range m {
		if k == target {
			return v
		}
	}
	return nil
}
func (SiteInformationService *SiteInformationService) GetRealTimeDataSearch(query papReq.HistoricalDataSearch) (err error, list interface{}) {
	redis := global.GVA_REDIS
	ctx := context.Background()
	redis.Do(ctx, "SELECT", "0")
	item, err := redis.Get(ctx, "MQTTModel").Result()
	result := map[string]interface{}{}
	json.Unmarshal([]byte(item), &result)
	siteresult := Scan(result, query.SiteCode)
	switch v := siteresult.(type) {
	case map[string]interface{}:
		for k, v := range v {
			if k == query.Device {
				fmt.Println(v)
				return err, v
			}
		}
	}
	return err, nil
}
func (SiteInformationService *SiteInformationService) GetHistoricalDataSearch(query papReq.HistoricalDataSearch) (err error, list interface{}) {
	/*
			select
		    time_bucket('7 days', pap."TOD") as "Date and Time",
		    pap."DeviceID" as "Device",
		    pap."DeviceDesc" as "Device Description",
		    max(pap."Inst A-Phase Current") as "Inst A-Phase Current",
		    max(pap."Inst B-Phase Current") as "Inst B-Phase Current",
		    max(pap."Inst C-Phase Current") as "Inst C-Phase Current",
		    max(pap."Inst A-Phase Voltage") as "Inst A-Phase Voltage",
		    max(pap."Inst B-Phase Voltage") as "Inst B-Phase Voltage",
		    max(pap."Inst C-Phase Voltage") as "Inst C-Phase Voltage"
			from
					"PAP_EEMS_XXXX" pap
			where
					pap."TOD" > '2020-01-01'
					and
					pap."DeviceID" = 'XXXXX'
			group by
					"Date and Time",
					"Device Description",
					"Device"
			order by
					"Date and Time" asc
	*/
	var siteID = query.SiteCode
	var startTime = query.TimeStart
	var endTime = query.TimeEnd
	var timeBucket = query.TimeBucket
	var limit = query.PageSize
	var offset = query.PageSize * (query.Page - 1)
	var device = query.Device
	var meterList = query.MeterList
	var selectQuery []string
	var whereQuery []string
	// redis := global.GVA_REDIS
	// ctx := context.Background()
	// redis.Do(ctx, "SELECT", "0")
	// item, err := redis.Get(ctx, "ModelRegisterDict").Result()
	// redisresult := map[string]interface{}{}
	// json.Unmarshal([]byte(item), &redisresult)
	// targetSite := Scan(redisresult, siteID)
	// targetMap := targetSite.(map[string]interface{})
	// registers := map[string]string{}
	// for _, value := range targetMap {
	// 	subitem := value.(map[string]interface{})
	// 	if subitem["DeviceID"] == device {
	// 		registers = subitem["Register"].(map[string]string)
	// 	}
	// }
	selectQuery = append(selectQuery, "count(*) OVER() AS total_count")
	selectQuery = append(selectQuery, fmt.Sprintf(`
		time_bucket('%s', pap."TOD") as "Date and Time"
	`, timeBucket))
	selectQuery = append(selectQuery, "pap.\"DeviceID\" as \"Device\"")
	selectQuery = append(selectQuery, "pap.\"DeviceDesc\" as \"Device Description\"")
	for _, v := range meterList {
		var meter = v
		columnname := strings.ReplaceAll(meter.MeterName, " ", "")
		columnname = strings.ReplaceAll(columnname, ":", "_")
		columnname = strings.ReplaceAll(columnname, "%", "_")
		var tempselectQuery = fmt.Sprintf(`
		%s("%s") as "%s:%s"
		`, meter.MeterFunction, columnname, meter.MeterFunction, meter.MeterName)
		selectQuery = append(selectQuery, tempselectQuery)
	}
	if startTime != "" {
		whereQuery = append(whereQuery, fmt.Sprintf(`
		pap."TOD" >= '%s'
		`, startTime))
	}
	if endTime != "" {
		whereQuery = append(whereQuery, fmt.Sprintf(`
		pap."TOD" <= '%s'
		`, endTime))
	}
	whereQuery = append(whereQuery, fmt.Sprintf(`
		pap."DeviceID" = '%s'
		`, device))
	selectStatement := strings.Join(selectQuery, ",")
	whereStatement := strings.Join(whereQuery, ` and `)
	queryStatement := fmt.Sprintf(`
		SELECT
		%s
		FROM
		"PAPSolar_%s" pap
		WHERE
		%s
		GROUP BY
		"Date and Time",
		"DeviceID",
		"DeviceDesc"
		ORDER BY
		"Date and Time" asc
		LIMIT %d offset %d
	`, selectStatement, siteID, whereStatement, limit, offset)
	var result []map[string]interface{}
	println(queryStatement)
	err = global.GetGlobalDBByDBName("planetarkpower").Raw(queryStatement).Scan(&result).Error
	println(result)
	return err, result
}
