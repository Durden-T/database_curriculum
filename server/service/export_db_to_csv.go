package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"io"
	"os"
)

const (
	exportDataSQL = `select * from %s into outfile '%s'
    CHARACTER SET gbk
    fields terminated by ','
    optionally enclosed BY '"'
    escaped by '\q' # 相当于删除转义
`
	getColumnSQL = `select group_concat(COLUMN_NAME order by ORDINAL_POSITION Separator ',') from information_schema.COLUMNS
where table_name = '%s' and table_schema = '%s';
`
)

func ExportDBtoCSV(r model.ImportExportTableinfo) error {
	tmpFileName := r.CSVpath + ".tmp"
	if err := global.GVA_DB.Exec(fmt.Sprintf(exportDataSQL, r.TableName, tmpFileName)).Error; err != nil {
		return err
	}
	dbName := global.GVA_CONFIG.Mysql.Dbname
	var column string
	if err := global.GVA_DB.Raw(fmt.Sprintf(getColumnSQL, r.TableName, dbName)).Scan(&column).Error; err != nil {
		return err
	}

	reader, err := os.Open(tmpFileName)
	if err != nil {
		return err
	}
	defer reader.Close()

	writer, err := os.Create(r.CSVpath)
	if err != nil {
		return err
	}
	defer writer.Close()
	_, err = writer.WriteString(column + "\n")
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, reader)

	reader.Close()
	os.Remove(tmpFileName)

	return err
}
