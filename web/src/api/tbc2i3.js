import service from '@/utils/request'



// @Tags Tbprb
// @Summary 获取tbprb图表属性
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Chart true "获取tbprb图表属性"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /businessQuery/getTbprbChartValue [post]
export const sectorTripleAnalysis = (params) => {
    return service({
        url: "/dataAnalysis/sectorTripleAnalysis",
        method: 'get',
        params
    })
}
