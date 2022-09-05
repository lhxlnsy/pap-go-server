package pap

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SiteInformationRouter struct {
}

// InitSiteInformationRouter 初始化 SiteInformation 路由信息
func (s *SiteInformationRouter) InitSiteInformationRouter(Router *gin.RouterGroup) {
	siteInformationRouter := Router.Group("siteInformation").Use(middleware.OperationRecord())
	siteInformationRouterWithoutRecord := Router.Group("siteInformation")
	var siteInformationApi = v1.ApiGroupApp.PAPApiGroup.SiteInformationApi
	{
		siteInformationRouter.POST("createSiteInformation", siteInformationApi.CreateSiteInformation)             // 新建SiteInformation
		siteInformationRouter.DELETE("deleteSiteInformation", siteInformationApi.DeleteSiteInformation)           // 删除SiteInformation
		siteInformationRouter.DELETE("deleteSiteInformationByIds", siteInformationApi.DeleteSiteInformationByIds) // 批量删除SiteInformation
		siteInformationRouter.PUT("updateSiteInformation", siteInformationApi.UpdateSiteInformation)              // 更新SiteInformation
	}
	{
		siteInformationRouter.POST("createChartSaved", siteInformationApi.CreateChartSaved)
		siteInformationRouter.DELETE("deleteChartSavedByIds", siteInformationApi.DeleteChartSavedByIds)
		siteInformationRouter.PUT("updateChartSavedByIds", siteInformationApi.UpdateChartSaved)
	}
	{
		siteInformationRouterWithoutRecord.GET("getRealTimeDataSearch", siteInformationApi.GetRealTimeDataSearch)
		siteInformationRouterWithoutRecord.GET("getChartSavedByUserId", siteInformationApi.GetChartSavedByUserId)
		siteInformationRouterWithoutRecord.GET("findSiteInformation", siteInformationApi.FindSiteInformation)       // 根据ID获取SiteInformation
		siteInformationRouterWithoutRecord.GET("getSiteInformationList", siteInformationApi.GetSiteInformationList) // 获取SiteInformation列表
		siteInformationRouterWithoutRecord.GET("getErrorInformationList", siteInformationApi.GetErrorInformation)   // 获取错误信息列表
		siteInformationRouterWithoutRecord.GET("getSiteDailySummary", siteInformationApi.GetSiteDailySummary)
		siteInformationRouterWithoutRecord.GET("searchSiteAlarmDetail", siteInformationApi.SearchSiteAlarmDetail)
		siteInformationRouterWithoutRecord.GET("getSiteAlarmDataStatus", siteInformationApi.GetSiteAlarmDataStatus)
		siteInformationRouterWithoutRecord.GET("getSiteRedisData", siteInformationApi.GetSiteRedisData)
		siteInformationRouterWithoutRecord.GET("getHistoricalDataSearch", siteInformationApi.GetHistoricalDataSearch)
		siteInformationRouterWithoutRecord.GET("getAllSites", siteInformationApi.GetAllSites)
	}
}
