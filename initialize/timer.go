package initialize

import (
	"fmt"

	"github.com/lhxlnsy/pap-go-server/config"
	"github.com/lhxlnsy/pap-go-server/global"
	"github.com/lhxlnsy/pap-go-server/utils"
)

func Timer() {
	if global.GVA_CONFIG.Timer.Start {
		for i := range global.GVA_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.GVA_Timer.AddTaskByFunc("ClearDB", global.GVA_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.GVA_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(global.GVA_CONFIG.Timer.Detail[i])
		}
	}
}
