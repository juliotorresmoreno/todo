package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var orm *xorm.Engine

func init() {
	var charset = "?charset=utf8&parseTime=true"
	var dsn = "root:paramore@tcp(127.0.0.1:3306)/crud" + charset
	var err error
	orm, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
}

//GetEngine Obtiene la conexion a la base de datos
func GetEngine() *xorm.Engine {
	return orm
}
