package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2"
	"log"
)

var Database *gorose.Engin

func init() {
	var err error
	Database, err = gorose.Open(DbConfig())
	if err != nil {
		log.Panic(err)
	}
}

func DbConfig() *gorose.Config {
	var conf gorose.Config
	conf.Driver = "mysql"
	conf.SetMaxIdleConns = 32
	conf.SetMaxOpenConns = 96
	conf.Prefix = ""
	conf.Dsn = dsn()
	return &conf
}

func dsn() string {
	dbname := "mirai"
	dbuser := "mirai"
	dbpass := "qwerty123"
	dbhost := "10.0.0.170"
	conntype := "tcp"
	dbport := "3306"
	charset := "utf8mb4"
	return dbuser + ":" + dbpass + "@" + conntype + "(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=" + charset + "&parseTime=true"

}
