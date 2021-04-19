// 自动生成模板Tbcell
package model

// 如果含有time.Time 请自行import time包
type Tbcell struct {
	City       string  `json:"city" form:"city" gorm:"column:city;comment:;type:varchar(255);size:255;" csv:"CITY"`
	SectorId   string  `json:"sectorId" form:"sectorId" gorm:"column:sector_id;comment:;type:varchar(255);size:255; csv:"SECTOR_ID"`
	SectorName string  `json:"sectorName" form:"sectorName" gorm:"column:sector_name;comment:;type:varchar(255);size:255;" csv:"SECTOR_NAME"`
	EnodebId   int     `json:"enodebId" form:"enodebId" gorm:"column:enodeb_id;comment:;type:int;size:10;" csv:"ENODEB_ID"`
	EnodebName string  `json:"enodebName" form:"enodebName" gorm:"column:enodeb_name;comment:;type:varchar(255);size:255;" csv:"ENODEB_NAME"`
	Earfcn     int     `json:"earfcn" form:"earfcn" gorm:"column:earfcn;comment:;type:int;size:10;" csv:"EARFCN"`
	Pci        int     `json:"pci" form:"pci" gorm:"column:pci;comment:;type:int;size:10;" csv:"PCI"`
	Pss        int     `json:"pss" form:"pss" gorm:"column:pss;comment:;type:int;size:10;" csv:"PSS"`
	Sss        int     `json:"sss" form:"sss" gorm:"column:sss;comment:;type:int;size:10;" csv:"SSS"`
	Tac        int     `json:"tac" form:"tac" gorm:"column:tac;comment:;type:int;size:10;" csv:"TAC"`
	Vendor     string  `json:"vendor" form:"vendor" gorm:"column:vendor;comment:;type:varchar(255);size:255;" csv:"VENDOR"`
	Longitude  float64 `json:"longitude" form:"longitude" gorm:"column:longitude;comment:;type:float;" csv:"LONGITUDE"`
	Latitude   float64 `json:"latitude" form:"latitude" gorm:"column:latitude;comment:;type:float;" csv:"LATITUDE"`
	Style      string  `json:"style" form:"style" gorm:"column:style;comment:;type:varchar(255);size:255;" csv:"STYLE"`
	Azimuth    float64 `json:"azimuth" form:"azimuth" gorm:"column:azimuth;comment:;type:float;" csv:"AZIMUTH"`
	Height     float64 `json:"height" form:"height" gorm:"column:height;comment:;type:float;" csv:"HEIGHT"`
	Electtilt  float64 `json:"electtilt" form:"electtilt" gorm:"column:electtilt;comment:;type:float;" csv:"ELECTTILT"`
	Mechtilt   float64 `json:"mechtilt" form:"mechtilt" gorm:"column:mechtilt;comment:;type:float;" csv:"MECHTILT"`
	Totletilt  float64 `json:"totletilt" form:"totletilt" gorm:"column:totletilt;comment:;type:float;" csv:"TOTLETILT"`
}

func (Tbcell) TableName() string {
	return "tbcell"
}
