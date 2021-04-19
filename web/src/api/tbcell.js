import service from '@/utils/request'

// @Tags Tbcell
// @Summary 分页获取Tbcell列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param params body request.PageInfo true "分页获取Tbcell列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /businessQuery/getTbcellList [get]
 export const getTbcellList = (params) => {
     return service({
         url: "/businessQuery/getTbcellList",
         method: 'get', 
         params
     })
 }