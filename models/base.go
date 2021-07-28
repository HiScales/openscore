package models

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var x *xorm.Engine

func init() {
	var err error
	user, _ := beego.AppConfig.String("mysql::MYSQL_USER")
	password, _ := beego.AppConfig.String("mysql::MYSQL_PASSWORD")
	host, _ := beego.AppConfig.String("mysql::MYSQL_HOST")
	port, _ := beego.AppConfig.String("mysql::MYSQL_PORT")
	database, _ := beego.AppConfig.String("mysql::MYSQL_DATABASE")
	mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, database)
	fmt.Println(mysqlDSN)
	x, err = xorm.NewEngine("mysql", mysqlDSN)
	if err != nil {
		fmt.Println("Fail to create xorm engine")
		panic(err)
	}
	x.ShowSQL(true)
	initMarkingModels()

}
