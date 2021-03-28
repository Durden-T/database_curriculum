package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ExportDBtoCSV(c *gin.Context) {
	var r model.ImportExportTableinfo
	c.ShouldBindJSON(&r)
	err := service.ExportDBtoCSV(r)
	if err != nil {
		global.GVA_LOG.Error("导出失败", zap.Any("err", err))
		response.FailWithMessage("导出失败", c)
		return
	}
	response.OkWithMessage("导出成功", c)
}
