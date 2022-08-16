package main

import (
	"fmt"
	"gmonitor"
)

func main() {
	vm := gmonitor.SendToMailViewModel{}
	vm.User = "hrlinktest@hrlink.com.cn"
	vm.Password = "!QA2ws3ed"
	vm.Host = "smtp.exmail.qq.com:587"
	vm.Send_to = []string{"richter.zhang@hrlink.com.cn"}
	vm.Subject = "email test by net / smtp"
	vm.Mailtype = "html"
	vm.Body = "email test by net/smtp"
	fmt.Println("send email")
	err := gmonitor.SendToMailV2(vm)
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}
