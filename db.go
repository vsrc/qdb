package qdb

import (
	"log"

	"github.com/qapi/gorm"
)

type DBConf struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBName     string
}

type Impl struct {
	DB *gorm.DB
}

func (i *Impl) InitDB(dbConf *DBConf) {

	var err error
	i.DB, err = gorm.Open("mysql", dbConf.DBUsername+":"+dbConf.DBPassword+"@"+dbConf.DBHost+"/"+dbConf.DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	i.DB.LogMode(true)

}
