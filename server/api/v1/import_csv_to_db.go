package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ImportCSVtoDB(c *gin.Context) {
	var r model.ImportExportTableinfo
	c.ShouldBindJSON(&r)

	if err := service.ImportCSVtoDB(r); err != nil {
		global.GVA_LOG.Error("导入失败", zap.Any("err", err))
		response.FailWithMessage("导入失败", c)
		return
	}
	response.OkWithMessage("导入成功", c)
}
