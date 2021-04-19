package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Tbkpi
// @Summary 获取tbkpi图表属性
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Chart true "获取tbkpi图表属性"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /businessQuery/getTbkpiChartValue [get]
func GetTbkpiChartValue(c *gin.Context) {
	var r request.Chart
	_ = c.ShouldBindQuery(&r)
	if err := utils.Verify(r, utils.TbkpiChartVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, res := service.GetTbkpiChartValue(r); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// @Tags Tbkpi
// @Summary 获取所有小区名称
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /businessQuery/getAllSectorName [get]
func GetAllSectorName(c *gin.Context) {
	if err, res := service.GetAllSectorName(); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}
