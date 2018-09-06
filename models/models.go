package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type ResponseData struct {//抛出的数据结构
	Status int
	Msg interface{}
}

type Order struct{//订单的数据结构
	Method string
	ZipCode string
	SendName string
	SendAddress string
	RecipientName string
	RecipientAddress string
	UserCode string
}

type OrederResponse struct {//订单请求返回的数据结构
	OrderCode int64
}


func init() {
	//orm.Debug = true //是否打印orm操作日志
	//数据库连接配置
	mysql_db := beego.AppConfig.String("mysql_db")
	mysql_user := beego.AppConfig.String("mysql_user")
	mysql_pwd := beego.AppConfig.String("mysql_pwd")
	mysql_addr := beego.AppConfig.String("mysql_addr")
	mysql_port := beego.AppConfig.String("mysql_port")
	mysql_charset := beego.AppConfig.String("mysql_charset")
	mysql_init := mysql_user + ":" + mysql_pwd + "@tcp(" + mysql_addr + ":" + mysql_port + ")/" + mysql_db + "?charset=" + mysql_charset
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysql_init, 30)
	
}