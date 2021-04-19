import service from '@/utils/request'


// @Tags Tbprb
// @Summary 获取tbprb图表属性
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Chart true "获取tbprb图表属性"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /businessQuery/getTbprbChartValue [post]
export const getTbprbChartValue = (params) => {
    return service({
        url: "/businessQuery/getTbprbChartValue",
        method: 'get',
        params
    })
}


// @Tags Tbprb
// @Summary 获取所有基站名称
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /businessQuery/getAllEnodebName [get]
export const getAllEnodebName = () => {
   return service({
       url: "/businessQuery/getAllEnodebName",
       method: 'get'
   })
}
