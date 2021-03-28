package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDataManageRouter(Router *gin.RouterGroup) {
	dataManageRouter := Router.Group("dataManage").Use(middleware.OperationRecord())
	{
		dataManageRouter.PUT("importCSVtoDB", v1.ImportCSVtoDB)  // 导入csv
		dataManageRouter.POST("exportDBtoCSV", v1.ExportDBtoCSV) // 导出csv
	}
}
