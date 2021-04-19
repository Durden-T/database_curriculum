package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitBusinessQueryRouter(Router *gin.RouterGroup) {
	businessQueryRouter := Router.Group("businessQuery").Use(middleware.OperationRecord())
	{
		businessQueryRouter.GET("getTbcellList", v1.GetTbcellList)           // 获取Tbcell列表
		businessQueryRouter.GET("getTbkpiChartValue", v1.GetTbkpiChartValue) // 获取tbkpi图表属性
		businessQueryRouter.GET("getAllSectorName", v1.GetAllSectorName)     // 获取所有小区名称
		businessQueryRouter.GET("getTbprbChartValue", v1.GetTbprbChartValue) // 获取tbkpi图表属性
		businessQueryRouter.GET("getAllEnodebName", v1.GetAllEnodebName)     // 获取所有小区名称
	}
}
