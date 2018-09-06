package lang

import (
	"github.com/astaxie/beego"
)

type LangOut struct {
	En string
	Ch string
}

/**
* @todo 获取提示
* @param code 代码 int 类型 
* @return tip 提示语句
*/
func Tips(code int) (tip string){
	typ := beego.AppConfig.String("lang") //获取系统配置的语言
	if typ != "en"{
		typ = "ch"
	}
	tip = getLang(code , typ)
	return 
}

/**
* @todo 转换代码对应的语句
* @param key 代码 int
* param typ 系统语言类型 en | ch
*/
func getLang(key int ,typ string) (tip string) {
	tip = ""
	langMap := map[int]LangOut{ //语言提示
		-1 : LangOut{ En :"Error: token is required" ,Ch : "错误: token 为必填项" },
		-2 : LangOut{ En :"Error: key is required" ,Ch : "错误: key 为必填项" },
		-3 : LangOut{ En :"Error: request failed" ,Ch : "错误: 请求失败" },
		-4 :  LangOut{ En :"Error: you have not such order" ,Ch : "错误: 你没有该订单"},
		-5 : LangOut{ En :"Error: System busy" ,Ch:"错误: 系统繁忙"},
		-6 : LangOut{En:"Error: usercode is required",Ch:"错误: usercode 为必填项"},
		-7 : LangOut{En: "Error: you have not such order" ,Ch:"错误: 你没有此订单"},
		-8 : LangOut{En:"Tip : Order Condition is not change",Ch:"提示: 订单信息没有变化"},
		-9 : LangOut{En:"Error: ordercode is required",Ch:"错误: ordercode 为必填项"},
	}

	
	if _,ok := langMap[key]; ok{ //查看代码是否配置
		if typ == "en"{
			tip = langMap[key].En
		}else{
			tip = langMap[key].Ch
		}
		
	}
	return 
}



