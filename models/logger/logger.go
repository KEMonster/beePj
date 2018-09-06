package logger

import (
	"time"
	"github.com/astaxie/beego/logs"
)

/**
* @todo 记录日志到对应到文本
* @param typ 日志类型 logs|error|warning|notice|debug|info|critical|alert|emergency ...
* @param msg 日志内容
*/
func Note(typ string ,msg string){
	Nowdate := time.Now().String()
	Nowdate = Nowdate[0:10] //日期
	filename := "logs/"+typ+"_"+Nowdate+".log" //文件名：日志类型
	logs_config := `{"filename":"`+ filename+ `"}` //设置日志属性
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterFile, logs_config) 
	log.Info(msg)//写入日志
}