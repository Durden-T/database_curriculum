package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [Durden-T](https://github.com/Durden-T)
//@function: GetTbcellInfoList
//@description: 分页获取Tbcell记录
//@param: info request.TbcellSearch
//@return: err error, list interface{}, total int64

func GetTbcellInfoList(info request.TbcellSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Tbcell{})
	var tbcells []model.Tbcell
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.SectorId != "" {
		db = db.Where("`sector_id` LIKE ?", "%"+info.SectorId+"%")
	}
	if info.SectorName != "" {
		db = db.Where("`sector_name` LIKE ?", "%"+info.SectorName+"%")
	}
	if info.EnodebId != 0 {
		db = db.Where("`enodeb_id` = ?", info.EnodebId)
	}
	if info.EnodebName != "" {
		db = db.Where("`enodeb_name` LIKE ?", "%"+info.EnodebName+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&tbcells).Error
	return err, tbcells, total
}
