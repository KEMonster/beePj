package vaildate

/**
* @todo SZUPS验证
*/
func (this *Order) SzupsVaildate() (status int, err string){
	if this.Method == ""{
		status = 1
		err = "method is required"
	}else if this.ZipCode == "" {
		status = 2
		err = "zipcode is required"
	}else if this.SendName == "" {
		status = 3
		err = "sendname is required"
	}else if this.SendAddress == "" {
		status = 4
		err = "sendaddress is required"
	}else if this.RecipientName == "" {
		status = 5
		err = "recipientname is required"
	}else if this.RecipientAddress == "" {
		status = 6
		err = "recipientaddress is required"
	}else{
		status = 0
	}
	return
}