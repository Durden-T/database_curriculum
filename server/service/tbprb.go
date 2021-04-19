package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gorm.io/gorm"
)

//@author: [Durden-T](https://github.com/Durden-T)
//@function: GetTbprbChartValue
//@description: 获取tbprb图表属性
//@param: r request.Chart
//@return: err error, list interface{}

func GetTbprbChartValue(r request.Chart) (err error, list interface{}) {
	// 创建db
	var db *gorm.DB
	switch r.Interval {
	case request.MinuteInterval:
		db = global.GVA_DB.Model(&model.Tbprb{})
	case request.HourInterval:
		db = global.GVA_DB.Table(model.Tbprb{}.NewTableName())
	}
	var tbkpis []map[string]interface{}
	err = db.Select("start_time", r.Type).Where("`enodeb_name` = ?", r.Name).
		Where("start_time between ? and ?", r.StartTime, r.EndTime).Order("start_time asc").Find(&tbkpis).Error
	if err != nil {
		return
	}

	if len(tbkpis) == 0 {
		err = errors.New("无对应数据")
		return
	}
	res := newChartValue(len(tbkpis))
	for _, kpi := range tbkpis {
		res.Xaxis = append(res.Xaxis, kpi["start_time"])
		res.Yaxis = append(res.Yaxis, kpi[r.Type])
	}

	return err, res
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: GetAllEnodebName
//@description: 获取所有小区名称
//@return: err error, list interface{}

func GetAllEnodebName() (err error, list interface{}) {
	db := global.GVA_DB.Model(&model.Tbkpi{})
	var names []string
	err = db.Distinct().Pluck("enodeb_name", &names).Error
	if err != nil {
		return
	}

	return err, names
}
