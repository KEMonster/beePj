package vaildate

type Order struct{
	Method string
	ZipCode string
	SendName string
	SendAddress string
	RecipientName string
	RecipientAddress string
	UserCode string
}

/**
* @todo 验证工厂类
*/
func (this *Order) Vaildate() (status int, err string) {
	switch this.Method { //运输方式
		case "SZUPS":
			status,err = this.SzupsVaildate()
		default:
			status,err = this.CommonVaildate()
	}
	return
}