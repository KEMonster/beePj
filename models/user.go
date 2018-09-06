package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	"errors"
	//"strconv"
)


type TbUserToken struct {
	Id int `orm:"column(iId);pk"`
	SToken string `orm:"column(sToken);size(50);null"`
	SKey string `orm:"column(sKey);size(128);null"`
	SUserCode string `orm:"column(sUserCode);size(20);null"`
}

//注册表TbUserToken
func init() {
	orm.RegisterModel(new(TbUserToken))
}

//验证token和key信息
func CheckUser(m *TbUserToken) (usercode string,err error) {
	usercode = ""
	o := orm.NewOrm()
	//user_token := TbUserToken{SToken: token, SKey:key}
	err = o.Read(m ,"sToken" ,"sKey")
	if err != nil {
		return 	
	}else if m.SUserCode != ""{
		usercode = m.SUserCode
		//r = true
	}else{
		err = errors.New("Error: No such token")
	}

	return
	
}

