package pap

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	PAPApi
	SiteInformationApi
}

var (
	papService             = service.ServiceGroupApp.PAPServiceGroup.PAPService
	siteInformationService = service.ServiceGroupApp.PAPServiceGroup.SiteInformationService
)
