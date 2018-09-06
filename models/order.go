package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	//"beePj/models/vaildate"
	"errors"
)

type TbOrder struct {
	Id int `orm:"column(sOrderCode);pk"`
	SMethod string `orm:"column(sMethod);size(20);"`
	SZipCode string `orm:"column(sZipCode);size(10);"`
	SSendName string `orm:"column(sSendName);size(50);"`
	SSendAddress string `orm:"column(sSendAddress);size(128);"`
	SRecipientName string `orm:"column(sRecipientName);size(50);"`
	SRecipientAddress string `orm:"column(sRecipientAddress);size(128);"`
	SUserCode string `orm:"column(sUserCode);size(20);"`
}

//注册表TbOrder
func init() {
	orm.RegisterModel(new(TbOrder))
}

//增加订单
func AddOrder (m *TbOrder)(id int64 , err error){
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//获取订单详情
func GetOrderByOrderCode(ordercode int)(m *TbOrder ,err error){
	o := orm.NewOrm();
	v := &TbOrder{Id : ordercode}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//更新订单信息
func UpdateOrder(m *TbOrder) (num int64,err error){
	o := orm.NewOrm();
	update_order := TbOrder{Id : m.Id}
	if o.Read(&update_order) == nil {
		num, err = o.Update(m)
	}else{
		err = errors.New("Error: No such order")
	}
	
	return 
}

//删除订单
func DeleteOrder(ordercode int) (num int64,err error) {
	o := orm.NewOrm()
	v := &TbOrder{Id : ordercode}
	num, err = o.Delete(v)
	return
}
