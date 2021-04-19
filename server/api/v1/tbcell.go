package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Tbcell
// @Summary 分页获取Tbcell列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TbcellSearch true "分页获取Tbcell列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /businessQuery/getTbcellList [get]
func GetTbcellList(c *gin.Context) {
	var pageInfo request.TbcellSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTbcellInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
