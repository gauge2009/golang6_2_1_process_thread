package main

import (
	"crypto/tls"
	"fmt"
	//"crypto/tls"
	"gopkg.in/gomail.v2"
)

/// █ █ █ █22-816 gomail发送不成功，1）先检查正文中不能出现html标签，否则有的smtp会拦截？  2）再检查配置中安全传输的 tls 是否开启
/// █ █ █ █22-816  emali目前唯一成功的程序片段如下
func main() {
	subject := "email test by gopkg.in gomail.v2"
	msg := "email test by gopkg.in gomail.v2"
	AlertByMail(subject, msg)
}

func AlertByMail(subject, message string) {
	m := gomail.NewMessage()
	///////m.SetHeader("From", "richter.zhang@hrlink.com.cn")// gomail: could not send email 1: 501 mail from address must be same as authorization user
	//m.SetHeader("From", "hrlinktest@hrlink.com.cn")
	//m.SetHeader("To", "richter.zhang@hrlink.com.cn")
	//m.SetAddressHeader("Cc", "promvc@live.com", "PROMVC")
	//m.SetHeader("Subject", subject)
	//m.SetBody("text/html", message)
	////////m.Attach("/home/Alex/lolcat.jpg")
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress("hrlinktest@hrlink.com.cn", "GrystalBeacon")},
		"To":      {"richter.zhang@hrlink.com.cn", "promvc@live.com"},
		"Cc":      {"promvc@live.com"},
		"Subject": {subject},
	})
	m.SetBody("text/html", message)
	d := gomail.NewDialer("smtp.exmail.qq.com", 587, "hrlinktest@hrlink.com.cn", "!QA2ws3ed")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Email Sent Successfully!")
}

func AlertByMailV0(subject, message string) {
	m := gomail.NewMessage()
	//m.SetHeader("From", "richter.zhang@hrlink.com.cn")// gomail: could not send email 1: 501 mail from address must be same as authorization user
	m.SetHeader("From", "hrlinktest@hrlink.com.cn")
	//m.SetHeader("To", "richter.zhang@hrlink.com.cn")
	m.SetHeader("To", "richter.zhang@hrlink.com.cn")
	m.SetAddressHeader("Cc", "promvc@live.com", "PROMVC")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.exmail.qq.com", 587, "hrlinktest@hrlink.com.cn", "!QA2ws3ed")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Email Sent Successfully!")
}
