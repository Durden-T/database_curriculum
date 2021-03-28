package model

type TbC2Inew struct {
	ServingSector     string  `json:"servingSector" form:"servingSector" gorm:"column:serving_sector;comment:;type:varchar(255);size:255;"`
	InterferingSector string  `json:"interferingSector" form:"interferingSector" gorm:"column:interfering_sector;comment:;type:varchar(255);size:255;"`
	C2iMean           float64 `json:"c2iMean" form:"c2iMean" gorm:"column:c2i_mean;comment:;type:float;"`
	C2iStd            float64 `json:"c2iStd" form:"c2iStd" gorm:"column:c2i_std;comment:;type:float;"`
	PrbC2i9           float64 `json:"prbC2i9" form:"prbC2i9" gorm:"column:prb_c2i9;comment:;type:float;"`
	PrbAbs6           float64 `json:"prbAbs6" form:"prbAbs6" gorm:"column:prb_abs6;comment:;type:float;"`
}

func (TbC2Inew) TableName() string {
	return "tbc2inew"
}
