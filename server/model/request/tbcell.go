package request

import "gin-vue-admin/model"

type TbcellSearch struct {
	model.Tbcell
	PageInfo
}
