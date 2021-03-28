package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"go.uber.org/zap"
	"math"
	"sync"
)

const importDataSQL = `
load data infile '%s' ignore into table %s
    fields terminated by ',' optionally enclosed by '"'
    lines terminated by '\r\n'
    ignore 1 lines
`

const (
	smallBatch    = 4
	writeBigBatch = 512
)

var specialTable = map[string]func() error{
	"tbprb":     calculateTbprb,
	"tbmrodata": calculateC2Iprb,
}

func ImportCSVtoDB(r model.ImportExportTableinfo) error {
	err := global.GVA_DB.Exec(fmt.Sprintf(importDataSQL, r.CSVpath, r.TableName)).Error
	if err != nil {
		return err
	}

	if fun := specialTable[r.TableName]; fun != nil {
		return fun()
	}
	return nil
}

func calculateTbprb() error {
	resultChan := make(chan *model.Tbprb, writeBigBatch)
	resultBuffer := make([]*model.Tbprb, 0, writeBigBatch)
	go func() {
		out := global.GVA_DB.Table(model.Tbprb{}.NewTableName())
		for prb := range resultChan {
			resultBuffer = append(resultBuffer, prb)
			if len(resultBuffer) >= writeBigBatch {
				if err := out.Create(&resultBuffer).Error; err != nil {
					global.GVA_LOG.Error("err", zap.Error(err))
				}
				resultBuffer = resultBuffer[:0]
			}
		}
		if err := out.Create(&resultBuffer).Error; err != nil {
			global.GVA_LOG.Error("err", zap.Error(err))
		}
	}()

	in := global.GVA_DB.Model(&model.Tbprb{})
	var buffer []*model.Tbprb
	in.Find(&buffer)
	var wg sync.WaitGroup
	for i := 0; i < len(buffer); i += smallBatch {
		wg.Add(1)
		go func(buf []*model.Tbprb) {
			defer wg.Done()
			tmp := &model.Tbprb{}
			for _, tbprb := range buf {
				tmp.StartTime = tbprb.StartTime
				tmp.EnodebName = tbprb.EnodebName
				tmp.SectorDescription = tbprb.SectorDescription
				tmp.SectorName = tbprb.SectorName
				tbprbSum(tmp, tbprb)
			}

			tbprbAvg(tmp)
			resultChan <- tmp
		}(buffer[i : i+smallBatch])
	}
	wg.Wait()

	close(resultChan)
	return nil
}

func calculateC2Iprb() error {
	db := global.GVA_DB.Model(&model.TbC2Inew{})
	if err := db.Exec("call gen_tbc2i_new(10)").Error; err != nil {
		return err
	}

	var result []*model.TbC2Inew
	if err := db.Find(&result).Error; err != nil {
		return err
	}

	for _, c2i := range result {
		c2i.PrbC2i9 = cdf(c2i.C2iMean, c2i.C2iStd, 9)
		c2i.PrbAbs6 = cdf(c2i.C2iMean, c2i.C2iStd, 6) - cdf(c2i.C2iMean, c2i.C2iStd, -6)
	}
	return db.Save(&result).Error
}

func cdf(mean, std, x float64) float64 {
	return 0.5 * math.Erfc(-(x-mean)/(std*math.Sqrt2))
}

func tbprbSum(tmp, tbprb *model.Tbprb) {
	tmp.Prb0 += tbprb.Prb0
	tmp.Prb1 += tbprb.Prb1
	tmp.Prb2 += tbprb.Prb2
	tmp.Prb3 += tbprb.Prb3
	tmp.Prb4 += tbprb.Prb4
	tmp.Prb5 += tbprb.Prb5
	tmp.Prb6 += tbprb.Prb6
	tmp.Prb7 += tbprb.Prb7
	tmp.Prb8 += tbprb.Prb8
	tmp.Prb9 += tbprb.Prb9
	tmp.Prb10 += tbprb.Prb10
	tmp.Prb11 += tbprb.Prb11
	tmp.Prb12 += tbprb.Prb12
	tmp.Prb13 += tbprb.Prb13
	tmp.Prb14 += tbprb.Prb14
	tmp.Prb15 += tbprb.Prb15
	tmp.Prb16 += tbprb.Prb16
	tmp.Prb17 += tbprb.Prb17
	tmp.Prb18 += tbprb.Prb18
	tmp.Prb19 += tbprb.Prb19
	tmp.Prb20 += tbprb.Prb20
	tmp.Prb21 += tbprb.Prb21
	tmp.Prb22 += tbprb.Prb22
	tmp.Prb23 += tbprb.Prb23
	tmp.Prb24 += tbprb.Prb24
	tmp.Prb25 += tbprb.Prb25
	tmp.Prb26 += tbprb.Prb26
	tmp.Prb27 += tbprb.Prb27
	tmp.Prb28 += tbprb.Prb28
	tmp.Prb29 += tbprb.Prb29
	tmp.Prb30 += tbprb.Prb30
	tmp.Prb31 += tbprb.Prb31
	tmp.Prb32 += tbprb.Prb32
	tmp.Prb33 += tbprb.Prb33
	tmp.Prb34 += tbprb.Prb34
	tmp.Prb35 += tbprb.Prb35
	tmp.Prb36 += tbprb.Prb36
	tmp.Prb37 += tbprb.Prb37
	tmp.Prb38 += tbprb.Prb38
	tmp.Prb39 += tbprb.Prb39
	tmp.Prb40 += tbprb.Prb40
	tmp.Prb41 += tbprb.Prb41
	tmp.Prb42 += tbprb.Prb42
	tmp.Prb43 += tbprb.Prb43
	tmp.Prb44 += tbprb.Prb44
	tmp.Prb45 += tbprb.Prb45
	tmp.Prb46 += tbprb.Prb46
	tmp.Prb47 += tbprb.Prb47
	tmp.Prb48 += tbprb.Prb48
	tmp.Prb49 += tbprb.Prb49
	tmp.Prb50 += tbprb.Prb50
	tmp.Prb51 += tbprb.Prb51
	tmp.Prb52 += tbprb.Prb52
	tmp.Prb53 += tbprb.Prb53
	tmp.Prb54 += tbprb.Prb54
	tmp.Prb55 += tbprb.Prb55
	tmp.Prb56 += tbprb.Prb56
	tmp.Prb57 += tbprb.Prb57
	tmp.Prb58 += tbprb.Prb58
	tmp.Prb59 += tbprb.Prb59
	tmp.Prb60 += tbprb.Prb60
	tmp.Prb61 += tbprb.Prb61
	tmp.Prb62 += tbprb.Prb62
	tmp.Prb63 += tbprb.Prb63
	tmp.Prb64 += tbprb.Prb64
	tmp.Prb65 += tbprb.Prb65
	tmp.Prb66 += tbprb.Prb66
	tmp.Prb67 += tbprb.Prb67
	tmp.Prb68 += tbprb.Prb68
	tmp.Prb69 += tbprb.Prb69
	tmp.Prb70 += tbprb.Prb70
	tmp.Prb71 += tbprb.Prb71
	tmp.Prb72 += tbprb.Prb72
	tmp.Prb73 += tbprb.Prb73
	tmp.Prb74 += tbprb.Prb74
	tmp.Prb75 += tbprb.Prb75
	tmp.Prb76 += tbprb.Prb76
	tmp.Prb77 += tbprb.Prb77
	tmp.Prb78 += tbprb.Prb78
	tmp.Prb79 += tbprb.Prb79
	tmp.Prb80 += tbprb.Prb80
	tmp.Prb81 += tbprb.Prb81
	tmp.Prb82 += tbprb.Prb82
	tmp.Prb83 += tbprb.Prb83
	tmp.Prb84 += tbprb.Prb84
	tmp.Prb85 += tbprb.Prb85
	tmp.Prb86 += tbprb.Prb86
	tmp.Prb87 += tbprb.Prb87
	tmp.Prb88 += tbprb.Prb88
	tmp.Prb89 += tbprb.Prb89
	tmp.Prb90 += tbprb.Prb90
	tmp.Prb91 += tbprb.Prb91
	tmp.Prb92 += tbprb.Prb92
	tmp.Prb93 += tbprb.Prb93
	tmp.Prb94 += tbprb.Prb94
	tmp.Prb95 += tbprb.Prb95
	tmp.Prb96 += tbprb.Prb96
	tmp.Prb97 += tbprb.Prb97
	tmp.Prb98 += tbprb.Prb98
	tmp.Prb99 += tbprb.Prb99
}

func tbprbAvg(tmp *model.Tbprb) {
	tmp.Prb0 /= smallBatch
	tmp.Prb1 /= smallBatch
	tmp.Prb2 /= smallBatch
	tmp.Prb3 /= smallBatch
	tmp.Prb4 /= smallBatch
	tmp.Prb5 /= smallBatch
	tmp.Prb6 /= smallBatch
	tmp.Prb7 /= smallBatch
	tmp.Prb8 /= smallBatch
	tmp.Prb9 /= smallBatch
	tmp.Prb10 /= smallBatch
	tmp.Prb11 /= smallBatch
	tmp.Prb12 /= smallBatch
	tmp.Prb13 /= smallBatch
	tmp.Prb14 /= smallBatch
	tmp.Prb15 /= smallBatch
	tmp.Prb16 /= smallBatch
	tmp.Prb17 /= smallBatch
	tmp.Prb18 /= smallBatch
	tmp.Prb19 /= smallBatch
	tmp.Prb20 /= smallBatch
	tmp.Prb21 /= smallBatch
	tmp.Prb22 /= smallBatch
	tmp.Prb23 /= smallBatch
	tmp.Prb24 /= smallBatch
	tmp.Prb25 /= smallBatch
	tmp.Prb26 /= smallBatch
	tmp.Prb27 /= smallBatch
	tmp.Prb28 /= smallBatch
	tmp.Prb29 /= smallBatch
	tmp.Prb30 /= smallBatch
	tmp.Prb31 /= smallBatch
	tmp.Prb32 /= smallBatch
	tmp.Prb33 /= smallBatch
	tmp.Prb34 /= smallBatch
	tmp.Prb35 /= smallBatch
	tmp.Prb36 /= smallBatch
	tmp.Prb37 /= smallBatch
	tmp.Prb38 /= smallBatch
	tmp.Prb39 /= smallBatch
	tmp.Prb40 /= smallBatch
	tmp.Prb41 /= smallBatch
	tmp.Prb42 /= smallBatch
	tmp.Prb43 /= smallBatch
	tmp.Prb44 /= smallBatch
	tmp.Prb45 /= smallBatch
	tmp.Prb46 /= smallBatch
	tmp.Prb47 /= smallBatch
	tmp.Prb48 /= smallBatch
	tmp.Prb49 /= smallBatch
	tmp.Prb50 /= smallBatch
	tmp.Prb51 /= smallBatch
	tmp.Prb52 /= smallBatch
	tmp.Prb53 /= smallBatch
	tmp.Prb54 /= smallBatch
	tmp.Prb55 /= smallBatch
	tmp.Prb56 /= smallBatch
	tmp.Prb57 /= smallBatch
	tmp.Prb58 /= smallBatch
	tmp.Prb59 /= smallBatch
	tmp.Prb60 /= smallBatch
	tmp.Prb61 /= smallBatch
	tmp.Prb62 /= smallBatch
	tmp.Prb63 /= smallBatch
	tmp.Prb64 /= smallBatch
	tmp.Prb65 /= smallBatch
	tmp.Prb66 /= smallBatch
	tmp.Prb67 /= smallBatch
	tmp.Prb68 /= smallBatch
	tmp.Prb69 /= smallBatch
	tmp.Prb70 /= smallBatch
	tmp.Prb71 /= smallBatch
	tmp.Prb72 /= smallBatch
	tmp.Prb73 /= smallBatch
	tmp.Prb74 /= smallBatch
	tmp.Prb75 /= smallBatch
	tmp.Prb76 /= smallBatch
	tmp.Prb77 /= smallBatch
	tmp.Prb78 /= smallBatch
	tmp.Prb79 /= smallBatch
	tmp.Prb80 /= smallBatch
	tmp.Prb81 /= smallBatch
	tmp.Prb82 /= smallBatch
	tmp.Prb83 /= smallBatch
	tmp.Prb84 /= smallBatch
	tmp.Prb85 /= smallBatch
	tmp.Prb86 /= smallBatch
	tmp.Prb87 /= smallBatch
	tmp.Prb88 /= smallBatch
	tmp.Prb89 /= smallBatch
	tmp.Prb90 /= smallBatch
	tmp.Prb91 /= smallBatch
	tmp.Prb92 /= smallBatch
	tmp.Prb93 /= smallBatch
	tmp.Prb94 /= smallBatch
	tmp.Prb95 /= smallBatch
	tmp.Prb96 /= smallBatch
	tmp.Prb97 /= smallBatch
	tmp.Prb98 /= smallBatch
	tmp.Prb99 /= smallBatch
}
