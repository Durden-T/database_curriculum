package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDataAnalysisRouter(Router *gin.RouterGroup) {
	dataAnalysisRouter := Router.Group("dataAnalysis").Use(middleware.OperationRecord())
	{
		dataAnalysisRouter.GET("sectorTripleAnalysis", v1.SectorTripleAnalysis) // 重叠覆盖干扰小区三元组分析
	}
}
