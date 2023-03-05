package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/go-sql-driver/mysql"
)

func NewDB() *gorm.DB {
	DBMS := "mysql"
	mysqlconfig := &mysql.Config{
		User:				config.C.Database.User,
		Passwd:				config.C.Database.Passwd,
		Net:				config.C.Database.Net,
		Address:			config.C.Database.Address,
		DBName: 			config.C.Database.DBName,
		AllowNativePasswd	config.C.Database.AllowNativePasswd,	
		Params: map[string]string {
			"paramsTime": config.C.Database.Params.paramsTime,
		},
	}
	db, err := gorm.Open(DBMS, mysqlconfig.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}
	return db
}