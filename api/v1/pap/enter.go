package pap

import "github.com/lhxlnsy/pap-go-server/service"

type ApiGroup struct {
	PAPApi
	SiteInformationApi
}

var (
	papService             = service.ServiceGroupApp.PAPServiceGroup.PAPService
	siteInformationService = service.ServiceGroupApp.PAPServiceGroup.SiteInformationService
)
