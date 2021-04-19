import service from '@/utils/request'

// @Tags Tbkpi
// @Summary 获取tbkpi图表属性
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TbkpiChart true "获取tbkpi图表属性"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /businessQuery/getTbkpiChartValue [post]
export const getTbkpiChartValue = (params) => {
     return service({
         url: "/businessQuery/getTbkpiChartValue",
         method: 'get',
         params
     })
 }


// @Tags Tbkpi
// @Summary 获取所有小区名称
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /businessQuery/getAllSectorName [get]
export const getAllSectorName = () => {
    return service({
        url: "/businessQuery/getAllSectorName",
        method: 'get'
    })
}
