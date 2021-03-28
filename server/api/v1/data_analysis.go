package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func SectorTripleAnalysis(c *gin.Context) {
	thresholdStr, _ := c.GetQuery("threshold")
	threshold, err := strconv.ParseFloat(thresholdStr, 64)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.SectorTripleAnalysis(threshold); err != nil {
		global.GVA_LOG.Error("分析失败", zap.Any("err", err))
		response.FailWithMessage("分析失败", c)
		return
	}
	response.OkWithMessage("分析成功", c)
}
