// 自动生成模板Tbprb
package model

const NewTableSuffix = "_new"

// 如果含有time.Time 请自行import time包
type Tbprb struct {
	SectorName        string `json:"sectorName" form:"sectorName" gorm:"column:sector_name;comment:;type:varchar(255);size:255;"`
	StartTime         string `json:"startTime" form:"startTime" gorm:"column:start_time;comment:;type:varchar(255);size:255;"`
	EnodebName        string `json:"enodebName" form:"enodebName" gorm:"column:enodeb_name;comment:;type:varchar(255);size:255;"`
	SectorDescription string `json:"sectorDescription" form:"sectorDescription" gorm:"column:sector_description;comment:;type:varchar(255);size:255;"`

	Prb0  int `json:"prb0" form:"prb0" gorm:"column:prb0;comment:;type:int;size:10;"`
	Prb1  int `json:"prb1" form:"prb1" gorm:"column:prb1;comment:;type:int;size:10;"`
	Prb10 int `json:"prb10" form:"prb10" gorm:"column:prb10;comment:;type:int;size:10;"`
	Prb11 int `json:"prb11" form:"prb11" gorm:"column:prb11;comment:;type:int;size:10;"`
	Prb12 int `json:"prb12" form:"prb12" gorm:"column:prb12;comment:;type:int;size:10;"`
	Prb13 int `json:"prb13" form:"prb13" gorm:"column:prb13;comment:;type:int;size:10;"`
	Prb14 int `json:"prb14" form:"prb14" gorm:"column:prb14;comment:;type:int;size:10;"`
	Prb15 int `json:"prb15" form:"prb15" gorm:"column:prb15;comment:;type:int;size:10;"`
	Prb16 int `json:"prb16" form:"prb16" gorm:"column:prb16;comment:;type:int;size:10;"`
	Prb17 int `json:"prb17" form:"prb17" gorm:"column:prb17;comment:;type:int;size:10;"`
	Prb18 int `json:"prb18" form:"prb18" gorm:"column:prb18;comment:;type:int;size:10;"`
	Prb19 int `json:"prb19" form:"prb19" gorm:"column:prb19;comment:;type:int;size:10;"`
	Prb2  int `json:"prb2" form:"prb2" gorm:"column:prb2;comment:;type:int;size:10;"`
	Prb20 int `json:"prb20" form:"prb20" gorm:"column:prb20;comment:;type:int;size:10;"`
	Prb21 int `json:"prb21" form:"prb21" gorm:"column:prb21;comment:;type:int;size:10;"`
	Prb22 int `json:"prb22" form:"prb22" gorm:"column:prb22;comment:;type:int;size:10;"`
	Prb23 int `json:"prb23" form:"prb23" gorm:"column:prb23;comment:;type:int;size:10;"`
	Prb24 int `json:"prb24" form:"prb24" gorm:"column:prb24;comment:;type:int;size:10;"`
	Prb25 int `json:"prb25" form:"prb25" gorm:"column:prb25;comment:;type:int;size:10;"`
	Prb26 int `json:"prb26" form:"prb26" gorm:"column:prb26;comment:;type:int;size:10;"`
	Prb27 int `json:"prb27" form:"prb27" gorm:"column:prb27;comment:;type:int;size:10;"`
	Prb28 int `json:"prb28" form:"prb28" gorm:"column:prb28;comment:;type:int;size:10;"`
	Prb29 int `json:"prb29" form:"prb29" gorm:"column:prb29;comment:;type:int;size:10;"`
	Prb3  int `json:"prb3" form:"prb3" gorm:"column:prb3;comment:;type:int;size:10;"`
	Prb30 int `json:"prb30" form:"prb30" gorm:"column:prb30;comment:;type:int;size:10;"`
	Prb31 int `json:"prb31" form:"prb31" gorm:"column:prb31;comment:;type:int;size:10;"`
	Prb32 int `json:"prb32" form:"prb32" gorm:"column:prb32;comment:;type:int;size:10;"`
	Prb33 int `json:"prb33" form:"prb33" gorm:"column:prb33;comment:;type:int;size:10;"`
	Prb34 int `json:"prb34" form:"prb34" gorm:"column:prb34;comment:;type:int;size:10;"`
	Prb35 int `json:"prb35" form:"prb35" gorm:"column:prb35;comment:;type:int;size:10;"`
	Prb36 int `json:"prb36" form:"prb36" gorm:"column:prb36;comment:;type:int;size:10;"`
	Prb37 int `json:"prb37" form:"prb37" gorm:"column:prb37;comment:;type:int;size:10;"`
	Prb38 int `json:"prb38" form:"prb38" gorm:"column:prb38;comment:;type:int;size:10;"`
	Prb39 int `json:"prb39" form:"prb39" gorm:"column:prb39;comment:;type:int;size:10;"`
	Prb4  int `json:"prb4" form:"prb4" gorm:"column:prb4;comment:;type:int;size:10;"`
	Prb40 int `json:"prb40" form:"prb40" gorm:"column:prb40;comment:;type:int;size:10;"`
	Prb41 int `json:"prb41" form:"prb41" gorm:"column:prb41;comment:;type:int;size:10;"`
	Prb42 int `json:"prb42" form:"prb42" gorm:"column:prb42;comment:;type:int;size:10;"`
	Prb43 int `json:"prb43" form:"prb43" gorm:"column:prb43;comment:;type:int;size:10;"`
	Prb44 int `json:"prb44" form:"prb44" gorm:"column:prb44;comment:;type:int;size:10;"`
	Prb45 int `json:"prb45" form:"prb45" gorm:"column:prb45;comment:;type:int;size:10;"`
	Prb46 int `json:"prb46" form:"prb46" gorm:"column:prb46;comment:;type:int;size:10;"`
	Prb47 int `json:"prb47" form:"prb47" gorm:"column:prb47;comment:;type:int;size:10;"`
	Prb48 int `json:"prb48" form:"prb48" gorm:"column:prb48;comment:;type:int;size:10;"`
	Prb49 int `json:"prb49" form:"prb49" gorm:"column:prb49;comment:;type:int;size:10;"`
	Prb5  int `json:"prb5" form:"prb5" gorm:"column:prb5;comment:;type:int;size:10;"`
	Prb50 int `json:"prb50" form:"prb50" gorm:"column:prb50;comment:;type:int;size:10;"`
	Prb51 int `json:"prb51" form:"prb51" gorm:"column:prb51;comment:;type:int;size:10;"`
	Prb52 int `json:"prb52" form:"prb52" gorm:"column:prb52;comment:;type:int;size:10;"`
	Prb53 int `json:"prb53" form:"prb53" gorm:"column:prb53;comment:;type:int;size:10;"`
	Prb54 int `json:"prb54" form:"prb54" gorm:"column:prb54;comment:;type:int;size:10;"`
	Prb55 int `json:"prb55" form:"prb55" gorm:"column:prb55;comment:;type:int;size:10;"`
	Prb56 int `json:"prb56" form:"prb56" gorm:"column:prb56;comment:;type:int;size:10;"`
	Prb57 int `json:"prb57" form:"prb57" gorm:"column:prb57;comment:;type:int;size:10;"`
	Prb58 int `json:"prb58" form:"prb58" gorm:"column:prb58;comment:;type:int;size:10;"`
	Prb59 int `json:"prb59" form:"prb59" gorm:"column:prb59;comment:;type:int;size:10;"`
	Prb6  int `json:"prb6" form:"prb6" gorm:"column:prb6;comment:;type:int;size:10;"`
	Prb60 int `json:"prb60" form:"prb60" gorm:"column:prb60;comment:;type:int;size:10;"`
	Prb61 int `json:"prb61" form:"prb61" gorm:"column:prb61;comment:;type:int;size:10;"`
	Prb62 int `json:"prb62" form:"prb62" gorm:"column:prb62;comment:;type:int;size:10;"`
	Prb63 int `json:"prb63" form:"prb63" gorm:"column:prb63;comment:;type:int;size:10;"`
	Prb64 int `json:"prb64" form:"prb64" gorm:"column:prb64;comment:;type:int;size:10;"`
	Prb65 int `json:"prb65" form:"prb65" gorm:"column:prb65;comment:;type:int;size:10;"`
	Prb66 int `json:"prb66" form:"prb66" gorm:"column:prb66;comment:;type:int;size:10;"`
	Prb67 int `json:"prb67" form:"prb67" gorm:"column:prb67;comment:;type:int;size:10;"`
	Prb68 int `json:"prb68" form:"prb68" gorm:"column:prb68;comment:;type:int;size:10;"`
	Prb69 int `json:"prb69" form:"prb69" gorm:"column:prb69;comment:;type:int;size:10;"`
	Prb7  int `json:"prb7" form:"prb7" gorm:"column:prb7;comment:;type:int;size:10;"`
	Prb70 int `json:"prb70" form:"prb70" gorm:"column:prb70;comment:;type:int;size:10;"`
	Prb71 int `json:"prb71" form:"prb71" gorm:"column:prb71;comment:;type:int;size:10;"`
	Prb72 int `json:"prb72" form:"prb72" gorm:"column:prb72;comment:;type:int;size:10;"`
	Prb73 int `json:"prb73" form:"prb73" gorm:"column:prb73;comment:;type:int;size:10;"`
	Prb74 int `json:"prb74" form:"prb74" gorm:"column:prb74;comment:;type:int;size:10;"`
	Prb75 int `json:"prb75" form:"prb75" gorm:"column:prb75;comment:;type:int;size:10;"`
	Prb76 int `json:"prb76" form:"prb76" gorm:"column:prb76;comment:;type:int;size:10;"`
	Prb77 int `json:"prb77" form:"prb77" gorm:"column:prb77;comment:;type:int;size:10;"`
	Prb78 int `json:"prb78" form:"prb78" gorm:"column:prb78;comment:;type:int;size:10;"`
	Prb79 int `json:"prb79" form:"prb79" gorm:"column:prb79;comment:;type:int;size:10;"`
	Prb8  int `json:"prb8" form:"prb8" gorm:"column:prb8;comment:;type:int;size:10;"`
	Prb80 int `json:"prb80" form:"prb80" gorm:"column:prb80;comment:;type:int;size:10;"`
	Prb81 int `json:"prb81" form:"prb81" gorm:"column:prb81;comment:;type:int;size:10;"`
	Prb82 int `json:"prb82" form:"prb82" gorm:"column:prb82;comment:;type:int;size:10;"`
	Prb83 int `json:"prb83" form:"prb83" gorm:"column:prb83;comment:;type:int;size:10;"`
	Prb84 int `json:"prb84" form:"prb84" gorm:"column:prb84;comment:;type:int;size:10;"`
	Prb85 int `json:"prb85" form:"prb85" gorm:"column:prb85;comment:;type:int;size:10;"`
	Prb86 int `json:"prb86" form:"prb86" gorm:"column:prb86;comment:;type:int;size:10;"`
	Prb87 int `json:"prb87" form:"prb87" gorm:"column:prb87;comment:;type:int;size:10;"`
	Prb88 int `json:"prb88" form:"prb88" gorm:"column:prb88;comment:;type:int;size:10;"`
	Prb89 int `json:"prb89" form:"prb89" gorm:"column:prb89;comment:;type:int;size:10;"`
	Prb9  int `json:"prb9" form:"prb9" gorm:"column:prb9;comment:;type:int;size:10;"`
	Prb90 int `json:"prb90" form:"prb90" gorm:"column:prb90;comment:;type:int;size:10;"`
	Prb91 int `json:"prb91" form:"prb91" gorm:"column:prb91;comment:;type:int;size:10;"`
	Prb92 int `json:"prb92" form:"prb92" gorm:"column:prb92;comment:;type:int;size:10;"`
	Prb93 int `json:"prb93" form:"prb93" gorm:"column:prb93;comment:;type:int;size:10;"`
	Prb94 int `json:"prb94" form:"prb94" gorm:"column:prb94;comment:;type:int;size:10;"`
	Prb95 int `json:"prb95" form:"prb95" gorm:"column:prb95;comment:;type:int;size:10;"`
	Prb96 int `json:"prb96" form:"prb96" gorm:"column:prb96;comment:;type:int;size:10;"`
	Prb97 int `json:"prb97" form:"prb97" gorm:"column:prb97;comment:;type:int;size:10;"`
	Prb98 int `json:"prb98" form:"prb98" gorm:"column:prb98;comment:;type:int;size:10;"`
	Prb99 int `json:"prb99" form:"prb99" gorm:"column:prb99;comment:;type:int;size:10;"`
}

func (Tbprb) TableName() string {
	return "tbprb"
}

func (Tbprb) NewTableName() string {
	return "tbprb_new"
}
