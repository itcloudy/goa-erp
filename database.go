package main

import (
	"goa-erp/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetDb get database conection
func GetDb() (db *gorm.DB, err error) {
	dbInfo := GetConfig().Database
	switch dbInfo.DbType {
	case "mysql":
		db, err = gorm.Open(dbInfo.DbType, utils.StringsJoin(
			dbInfo.DbUser,
			":",
			dbInfo.DbPwd,
			"@/",
			dbInfo.DbName,
			"?charset=",
			dbInfo.DbChart,
			"&parseTime=True&loc=Local"))
		defer db.Close()
	}
	return db, err
}
