package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

///发送成功的场景：content_type 必须是 text/html 且不嗯呢该有html标签
func SendToMailWithoutCC(user, password, host, subject, body, mailtype string, send_to []string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	to_address := strings.Join(send_to, ";")
	msg := []byte("To: " + to_address + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n " + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
	user := "hrlinktest@hrlink.com.cn"
	password := "!QA2ws3ed"
	host := "smtp.exmail.qq.com:587"
	to := []string{"richter.zhang@hrlink.com.cn"}
	subject := "email test by net/smtp"
	//mailtype :="html"
	mailtype := "html"
	body := "email test by net/smtp"
	fmt.Println("send email")
	err := SendToMailWithoutCC(user, password, host, subject, body, mailtype, to)
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}

//func main() {
//	user := "hrlinktest@hrlink.com.cn"
//	password := "!QA2ws3ed"
//	host := "smtp.exmail.qq.com:587"
//	to := []string{"richter.zhang@hrlink.com.cn" }
//	subject := "email test by jordan-wright /email"
//	//mailtype :="html"
//	mailtype:="";
//    body:= "email test by jordan-wright /email"
//	fmt.Println("send email")
//	err := gmonitor.SendToMail(user, password, host, subject, body, mailtype, to)
//	if err != nil {
//		fmt.Println("Send mail error!")
//		fmt.Println(err)
//	} else {
//		fmt.Println("Send mail success!")
//	}
//}
//

//func main() {
//	vm:= gmonitor.SendToMailViewModel{}
//	vm.User = "hrlinktest@hrlink.com.cn"
//	vm.Password = "!QA2ws3ed"
//	vm.Host = "smtp.exmail.qq.com:587"
//	vm.Send_to = []string{"richter.zhang@hrlink.com.cn" }
//	vm.Subject = "email test by jordan-wright /email"
//	vm.Mailtype ="";//"html"
//	vm.Body = "email test by jordan-wright /email"
//	fmt.Println("send email")
//	err := gmonitor.SendToMailV2(vm)
//	if err != nil {
//		fmt.Println("Send mail error!")
//		fmt.Println(err)
//	} else {
//		fmt.Println("Send mail success!")
//	}
//}
