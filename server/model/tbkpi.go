// 自动生成模板Tbkpi
package model

type Tbkpi struct {
	StartTime         string  `json:"startTime" form:"startTime" gorm:"column:start_time;comment:;type:varchar(255);size:255;" csv:"起始时间"`
	EnodebName        string  `json:"enodebName" form:"enodebName" gorm:"column:enodeb_name;comment:;type:varchar(255);size:255;" csv:"网元/基站名称"`
	SectorDescription string  `json:"sectorDescription" form:"sectorDescription" gorm:"column:sector_description;comment:;type:varchar(255);size:255;" csv:"小区"`
	SectorName        string  `json:"sectorName" form:"sectorName" gorm:"column:sector_name;comment:;type:varchar(255);size:255;" csv:"小区名称"`
	RccConnSucc       int     `json:"rccConnSucc" form:"rccConnSucc" gorm:"column:rcc_conn_succ;comment:;type:int;size:10;" csv:"RRC连接建立完成次数 (无)"`
	RccConnAtt        int     `json:"rccConnAtt" form:"rccConnAtt" gorm:"column:rcc_conn_att;comment:;type:int;size:10;" csv:"RRC连接请求次数（包括重发） (无)"`
	RccConnRate       float64 `json:"rccConnRate" form:"rccConnRate" gorm:"column:rcc_conn_rate;comment:;type:decimal;size:7,4;" csv:"RRC建立成功率qf (%)"`
}

func (Tbkpi) TableName() string {
	return "tbkpi"
}
