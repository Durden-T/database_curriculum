package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"strings"
)

const timeSuffix = " 00:00:00"

type chartValue struct {
	Xaxis []interface{}
	Yaxis []interface{}
}

func newChartValue(size int) chartValue {
	return chartValue{
		Xaxis: make([]interface{}, 0, size),
		Yaxis: make([]interface{}, 0, size),
	}
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: GetTbkpiChartValue
//@description: 获取tbkpi图表属性
//@param: r request.Chart
//@return: err error, list interface{}

func GetTbkpiChartValue(r request.Chart) (err error, list interface{}) {
	// 格式化时间
	r.StartTime += timeSuffix
	r.EndTime += timeSuffix
	// 创建db
	db := global.GVA_DB.Model(&model.Tbkpi{})
	var tbkpis []map[string]interface{}
	err = db.Debug().Select("start_time", r.Type).Where("`sector_name` = ?", r.Name).
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
		startTime, ok := kpi["start_time"].(string)
		if !ok {
			continue
		}
		startTime = strings.Split(startTime, " ")[0]
		res.Xaxis = append(res.Xaxis, startTime)
		res.Yaxis = append(res.Yaxis, kpi[r.Type])
	}

	return err, res
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: getAllSectorName
//@description: 获取所有小区名称
//@return: err error, list interface{}

func GetAllSectorName() (err error, list interface{}) {
	// 创建db
	db := global.GVA_DB.Model(&model.Tbkpi{})
	var names []string
	err = db.Distinct().Pluck("sector_name", &names).Error
	if err != nil {
		return
	}

	return err, names
}
