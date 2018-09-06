package controllers

import (
	"github.com/astaxie/beego"
	"beePj/models"
	"beePj/models/vaildate"
	"beePj/models/lang"
	"beePj/models/logger"
	"fmt"
	"strconv"
)

type OrderController struct {
	beego.Controller
}

/**
* @todo 验证token,user的有效性
* @return
*  user_code 用户名
*  status 状态码 0为成功 其他为失败 
*/
func BaseCheck(token string, key string) (user_code string, status int){
	user_code = ""
	if token == "" {
		status = -1
	}else if key == "" {
		status = -2
	}else{
		var tb_user_token models.TbUserToken
		tb_user_token.SToken = token
		tb_user_token.SKey = key
		userCode,err := models.CheckUser(&tb_user_token)
		user_code = userCode
		if err != nil {
			errMsg := fmt.Sprintf("user model error SQL ERROR: s%",err.Error())
			logger.Note("error", errMsg)
			status = -3
		}
	}
	return 
}

/**
* @todao 新增订单
* @method post
*/
func (this *OrderController) AddOrder() {
	logger.Note("logs",this.Ctx.Request.RequestURI) //记录请求地址
	var vaildate_order vaildate.Order
	var result models.ResponseData
	var user_code string
	key := this.GetString("key")
	token := this.GetString("token")
	userCode,status := BaseCheck(token, key) //验证user token
	user_code = userCode
	if status != 0{
		result.Status = status
		result.Msg = lang.Tips(status) //获取抛错提示
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	
	method := this.GetString("method")
	zip_code := this.GetString("zipcode")
	send_name := this.GetString("sendname")
	send_address := this.GetString("sendaddress")
	recipient_name := this.GetString("recipientname")
	recipient_address := this.GetString("recipientaddress")
	vaildate_order.Method = method
	vaildate_order.ZipCode = zip_code
	vaildate_order.SendName = send_name
	vaildate_order.SendAddress = send_address
	vaildate_order.RecipientName = recipient_name
	vaildate_order.RecipientAddress = recipient_address
	status,errMsg := vaildate_order.Vaildate() //校验参数 status 0为成功，其它为失败
	if status != 0 {
		result.Status = status
		result.Msg = errMsg
		this.Data["json"] = result
		this.ServeJSON()
		return
	}else{
		var tb_order models.TbOrder
		tb_order.SMethod = vaildate_order.Method
		tb_order.SZipCode = vaildate_order.ZipCode
		tb_order.SSendName = vaildate_order.SendName
		tb_order.SSendAddress = vaildate_order.SendAddress
		tb_order.SRecipientName = vaildate_order.RecipientName
		tb_order.SRecipientAddress = vaildate_order.RecipientAddress
		tb_order.SUserCode = user_code
		order_code,err := models.AddOrder(&tb_order) //新增订单
		if err != nil{
			errMsg := fmt.Sprintf("order model error SQL ERROR: s%",err.Error())
			logger.Note("error", errMsg) //错误记录日志
			result.Status = -5
			result.Msg = lang.Tips(-5)
			this.Data["json"] = result
			this.ServeJSON()
		}else{
			result.Status = status
			result.Msg = models.OrederResponse{OrderCode:order_code}
			this.Data["json"] = result
			this.ServeJSON()
		}
		return
	}

}

/**
* @todo 查找订单
* @method get
*/
func (this *OrderController) SearchOrder() {
	logger.Note("logs",this.Ctx.Request.RequestURI) //记录请求地址
	var result models.ResponseData
	key := this.GetString("key")
	token := this.GetString("token")
	order_code := this.GetString("ordercode")
	user_code := this.GetString("usercode")
	userCode,status := BaseCheck(token, key) //校验user token
	if status != 0{
		result.Status = status
		result.Msg = lang.Tips(status)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	if user_code == ""{
		result.Status = -6
		result.Msg = lang.Tips(-6)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	
	if user_code != userCode{
		result.Status = -4
		result.Msg = lang.Tips(-4)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	orderCode, err := strconv.Atoi(order_code) //将order_code 转化为int类型
	search_order,err := models.GetOrderByOrderCode(orderCode) //查找订单，若成功返回订单的数据结构
	if err != nil{
		errMsg := fmt.Sprintf("user model error SQL ERROR: s%",err.Error())
		logger.Note("error", errMsg)
		result.Status = -4
		result.Msg = lang.Tips(-4)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}else{
		result.Status = 0
		result.Msg = search_order
		this.Data["json"] = result
		this.ServeJSON()
	}
		
	
}

/**
* @todo更新订单
* @method get
*/
func (this *OrderController) UpdateOrder() {
	logger.Note("logs",this.Ctx.Request.RequestURI)
	var vaildate_order vaildate.Order
	var result models.ResponseData
	key := this.GetString("key")
	token := this.GetString("token")
	order_code := this.GetString("ordercode")
	user_code := this.GetString("usercode")
	userCode,status := BaseCheck(token, key)
	if status != 0{
		result.Status = status
		result.Msg = lang.Tips(status)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	if order_code == ""{
		result.Status = -9
		result.Msg = lang.Tips(-9)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	if user_code == ""{
		result.Status = -6
		result.Msg = lang.Tips(-6)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	if user_code != userCode{
		result.Status = -7
		result.Msg = lang.Tips(-7)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	method := this.GetString("method")
	zip_code := this.GetString("zipcode")
	send_name := this.GetString("sendname")
	send_address := this.GetString("sendaddress")
	recipient_name := this.GetString("recipientname")
	recipient_address := this.GetString("recipientaddress")
	vaildate_order.Method = method
	vaildate_order.ZipCode = zip_code
	vaildate_order.SendName = send_name
	vaildate_order.SendAddress = send_address
	vaildate_order.RecipientName = recipient_name
	vaildate_order.RecipientAddress = recipient_address
	status,errMsg := vaildate_order.Vaildate()
	if status != 0 {
		result.Status = status
		result.Msg = errMsg
		this.Data["json"] = result
		this.ServeJSON()
		return
	}else{
		var tb_order models.TbOrder
		orderCode, err := strconv.Atoi(order_code)
		tb_order.Id = orderCode
		tb_order.SMethod = vaildate_order.Method
		tb_order.SZipCode = vaildate_order.ZipCode
		tb_order.SSendName = vaildate_order.SendName
		tb_order.SSendAddress = vaildate_order.SendAddress
		tb_order.SRecipientName = vaildate_order.RecipientName
		tb_order.SRecipientAddress = vaildate_order.RecipientAddress
		tb_order.SUserCode = user_code
		affect_row,err := models.UpdateOrder(&tb_order)
		if err != nil{
			errMsg := fmt.Sprintf("order model error SQL ERROR: s%",err.Error())
			logger.Note("error", errMsg)
			result.Status = -5
			result.Msg = lang.Tips(-5)
			this.Data["json"] = result
			this.ServeJSON()
		}else{
			if affect_row < 1{
				result.Status = -8
				result.Msg = lang.Tips(-8)
			}else{
				result.Status = 0
				result.Msg = "Update Success"
			}
			this.Data["json"] = result
			this.ServeJSON()
		}
		
		return
	}
}

/**
* @todo 删除订单
* @method delete
*/
func (this *OrderController) DeleteOrder() {
	logger.Note("logs",this.Ctx.Request.RequestURI)
	var result models.ResponseData
	key := this.GetString("key")
	token := this.GetString("token")
	order_code := this.GetString("ordercode")
	user_code := this.GetString("usercode")

	userCode,status := BaseCheck(token, key)

	if status != 0{
		result.Status = status
		result.Msg = lang.Tips(status)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	if user_code == ""{
		result.Status = -6
		result.Msg = lang.Tips(-6)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	if user_code != userCode{
		result.Status = -4
		result.Msg = lang.Tips(-4)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	orderCode, err := strconv.Atoi(order_code)
	affect_row,err := models.DeleteOrder(orderCode)
	if err != nil{
		errMsg := fmt.Sprintf("user model error SQL ERROR: s%",err.Error())
		logger.Note("error", errMsg)
		result.Status = -3
		result.Msg = lang.Tips(-3)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}else{
		if affect_row < 1{
			result.Status = -4
			result.Msg = lang.Tips(-4)
		}else{
			result.Status = 0
			result.Msg = "Delete Success"
		}
		this.Data["json"] = result
		this.ServeJSON()
	}
		
	

}