package pap

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pap"
	papReq "github.com/flipped-aurora/gin-vue-admin/server/model/pap/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SiteInformationApi struct {
}

// CreateSiteInformation 创建SiteInformation
// @Tags SiteInformation
// @Summary 创建SiteInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pap.SiteInformation true "创建SiteInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /siteInformation/createSiteInformation [post]
func (siteInformationApi *SiteInformationApi) CreateSiteInformation(c *gin.Context) {
	var siteInformation pap.SiteInformation
	_ = c.ShouldBindJSON(&siteInformation)
	if err := siteInformationService.CreateSiteInformation(siteInformation); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// CreateChartSaved 创建User Saved Chart
// @Tags ChartSaved
// @Summary 创建UserSavedChart
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pap.ChartSaved true "创建UserSavedChart"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /siteInformation/createChartSaved [post]
func (siteInformationApi *SiteInformationApi) CreateChartSaved(c *gin.Context) {
	var chartInformation pap.ChartSaved
	_ = c.ShouldBindJSON(&chartInformation)
	if err := siteInformationService.CreateChartSaved(chartInformation); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (siteInformationApi *SiteInformationApi) UpdateChartSaved(c *gin.Context) {
	var chartInformation pap.ChartSaved
	_ = c.ShouldBindJSON(&chartInformation)
	if err := siteInformationService.UpdateChartSaved(chartInformation); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSiteInformation 用id查询Saved Chart
// @Tags ChartSaved
// @Summary 用id查询ChartSaved
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pap.ChartSaved true "用id查询ChartSaved"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /siteInformation/getChartSavedByUserId [get]
func (siteInformationApi *SiteInformationApi) GetChartSavedByUserId(c *gin.Context) {
	var chartInformation pap.ChartSaved
	if c.ShouldBindQuery(&chartInformation) == nil {
		fmt.Println(chartInformation)
		if chartInformation.SiteID == "" {
			response.FailWithMessage("查询失败", c)
		} else {
			if err, rechartInformation := siteInformationService.GetChartSavedByUserId(chartInformation); err != nil {
				global.GVA_LOG.Error("查询失败!", zap.Error(err))
				response.FailWithMessage("查询失败", c)
			} else {
				response.OkWithData(gin.H{"rechartInformation": rechartInformation}, c)
			}
		}
	} else {
		fmt.Println(c.ShouldBindQuery(&chartInformation))
		fmt.Println(chartInformation.UserName)
		fmt.Println(chartInformation.SiteID)
		response.FailWithMessage("查询失败", c)
	}
}

// DeleteSiteInformation 删除SiteInformation
// @Tags SiteInformation
// @Summary 删除SiteInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pap.SiteInformation true "删除SiteInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /siteInformation/deleteSiteInformation [delete]
func (siteInformationApi *SiteInformationApi) DeleteChartSavedByIds(c *gin.Context) {
	var chartSaved pap.ChartSaved
	_ = c.ShouldBindQuery(&chartSaved)
	fmt.Println(chartSaved)
	if err := siteInformationService.DeleteChartSavedByIds(chartSaved.ChartUUID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSiteInformation 删除SiteInformation
// @Tags SiteInformation
// @Summary 删除SiteInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pap.SiteInformation true "删除SiteInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /siteInformation/deleteSiteInformation [delete]
func (siteInformationApi *SiteInformationApi) DeleteSiteInformation(c *gin.Context) {
	var siteInformation pap.SiteInformation
	_ = c.ShouldBindJSON(&siteInformation)
	if err := siteInformationService.DeleteSiteInformation(siteInformation); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSiteInformationByIds 批量删除SiteInformation
// @Tags SiteInformation
// @Summary 批量删除SiteInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SiteInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /siteInformation/deleteSiteInformationByIds [delete]
func (siteInformationApi *SiteInformationApi) DeleteSiteInformationByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := siteInformationService.DeleteSiteInformationByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSiteInformation 更新SiteInformation
// @Tags SiteInformation
// @Summary 更新SiteInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pap.SiteInformation true "更新SiteInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /siteInformation/updateSiteInformation [put]
func (siteInformationApi *SiteInformationApi) UpdateSiteInformation(c *gin.Context) {
	var siteInformation pap.SiteInformation
	_ = c.ShouldBindJSON(&siteInformation)
	if err := siteInformationService.UpdateSiteInformation(siteInformation); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSiteInformation 用id查询SiteInformation
// @Tags SiteInformation
// @Summary 用id查询SiteInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pap.SiteInformation true "用id查询SiteInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /siteInformation/findSiteInformation [get]
func (siteInformationApi *SiteInformationApi) FindSiteInformation(c *gin.Context) {
	var siteInformation pap.SiteInformation
	_ = c.ShouldBindQuery(&siteInformation)
	if err, resiteInformation := siteInformationService.GetSiteInformation(siteInformation.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resiteInformation": resiteInformation}, c)
	}
}

// GetSiteInformationList 分页获取SiteInformation列表
// @Tags SiteInformation
// @Summary 分页获取SiteInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query papReq.ErrorInformation true "分页获取SiteInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /siteInformation/getSiteInformationList [get]
func (siteInformationApi *SiteInformationApi) GetErrorInformation(c *gin.Context) {
	var pageInfo pap.ErrorStieInformation
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list := siteInformationService.GetErrorInformation(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.ErrorResult{
			List: list,
		}, "获取成功", c)
	}
}

func (siteInformationApi *SiteInformationApi) GetSiteDailySummary(c *gin.Context) {
	var query pap.SiteInfo
	if c.ShouldBindQuery(&query) == nil {
		if query.SiteCode == "" {
			response.FailWithMessage("查询失败", c)
		} else {
			if err, list := siteInformationService.GetSiteDailySummary(query.SiteCode); err != nil {
				global.GVA_LOG.Error("获取失败", zap.Error(err))
				response.FailWithMessage("获取失败", c)
			} else {
				response.OkWithDetailed(response.PageResult{
					List: list,
				}, "获取成功", c)
			}
		}
	} else {
		response.FailWithMessage("查询失败", c)
	}
}

// GetErrorInformation 获取最新的错误信息以及对应的site信息
// @Tags SiteInformation
// @Summary 获取最新的错误信息以及对应的site信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query papReq.SiteInformationSearch true "分页获取SiteInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /siteInformation/getSiteInformationList [get]
func (siteInformationApi *SiteInformationApi) GetSiteInformationList(c *gin.Context) {
	var pageInfo papReq.SiteInformationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := siteInformationService.GetSiteInformationInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (siteInformationApi *SiteInformationApi) SearchSiteAlarmDetail(c *gin.Context) {
	var query papReq.AlarmInformationSearch
	if c.ShouldBindQuery(&query) == nil {
		if query.SITEID == "" {
			response.FailWithMessage("查询失败", c)
		} else {
			if err, list, total := siteInformationService.SearchSiteAlarmDetail(query); err != nil {
				global.GVA_LOG.Error("获取失败", zap.Error(err))
				response.FailWithMessage("获取失败", c)
			} else {
				response.OkWithDetailed(response.PageResult{
					List:     list,
					Total:    total,
					Page:     query.Page,
					PageSize: query.PageSize,
				}, "获取成功", c)
			}
		}
	} else {
		response.FailWithMessage("查询失败", c)
	}
}

func (siteInformationApi *SiteInformationApi) GetSiteAlarmDataStatus(c *gin.Context) {
	var query pap.SiteInfo
	if c.ShouldBindQuery(&query) == nil {
		if query.SiteCode == "" {
			response.FailWithMessage("查询失败", c)
		} else {
			if err, list := siteInformationService.GetSiteAlarmDataStatus(query.SiteCode); err != nil {
				global.GVA_LOG.Error("获取失败", zap.Error(err))
				response.FailWithMessage("获取失败", c)
			} else {
				response.OkWithDetailed(response.PageResult{
					List: list,
				}, "获取成功", c)
			}
		}
	} else {
		response.FailWithMessage("查询失败", c)
	}
}

func (SiteInformationApi *SiteInformationApi) GetSiteRedisData(c *gin.Context) {
	var query pap.SiteInfo
	if c.ShouldBindQuery(&query) == nil {
		if query.SiteCode == "" {
			response.FailWithMessage("查询失败", c)
		} else {
			if err, list := siteInformationService.GetSiteRedisData(query.SiteCode); err != nil {
				global.GVA_LOG.Error("获取失败", zap.Error(err))
				response.FailWithMessage("获取失败", c)
			} else {
				response.OkWithDetailed(response.PageResult{
					List: list,
				}, "获取成功", c)
			}
		}
	} else {
		response.FailWithMessage("查询失败", c)
	}
}
func (SiteInformationApi *SiteInformationApi) GetRealTimeDataSearch(c *gin.Context) {
	var query papReq.HistoricalDataSearchRequest
	if c.ShouldBindQuery(&query) == nil {
		if query.JsonBody == "" {
			response.FailWithMessage("查询失败", c)
		} else {
			var queryJson papReq.HistoricalDataSearch
			json.Unmarshal([]byte(query.JsonBody), &queryJson)
			if err, list := siteInformationService.GetRealTimeDataSearch(queryJson); err != nil {
				global.GVA_LOG.Error("获取失败", zap.Error(err))
				response.FailWithMessage("获取失败", c)
			} else {
				response.OkWithDetailed(response.PageResult{
					List: list,
				}, "获取成功", c)
			}
		}
	} else {
		response.FailWithMessage("查询失败", c)
	}
}
func (SiteInformationApi *SiteInformationApi) GetHistoricalDataSearch(c *gin.Context) {
	var query papReq.HistoricalDataSearchRequest
	if c.ShouldBindQuery(&query) == nil {
		if query.JsonBody == "" {
			response.FailWithMessage("查询失败", c)
		} else {
			var queryJson papReq.HistoricalDataSearch
			json.Unmarshal([]byte(query.JsonBody), &queryJson)
			if err, list := siteInformationService.GetHistoricalDataSearch(queryJson); err != nil {
				global.GVA_LOG.Error("获取失败", zap.Error(err))
				response.FailWithMessage("获取失败", c)
			} else {
				response.OkWithDetailed(response.PageResult{
					List:     list,
					Page:     queryJson.Page,
					PageSize: queryJson.PageSize,
				}, "获取成功", c)
			}
		}
	} else {
		response.FailWithMessage("查询失败", c)
	}
}
func (SiteInformationApi *SiteInformationApi) GetAllSites(c *gin.Context) {
	var query papReq.SiteInformationSearch
	_ = c.ShouldBindQuery(&query)
	if err, list, total := siteInformationService.GetAllSites(query); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Page:     query.Page,
			PageSize: query.PageSize,
			Total:    total,
		}, "获取成功", c)
	}
}
