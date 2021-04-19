package service

import "gin-vue-admin/global"

func SectorTripleAnalysis(threshold float64) error {
	return global.GVA_DB.Exec("call triple_analysis(?)", threshold).Error
}
