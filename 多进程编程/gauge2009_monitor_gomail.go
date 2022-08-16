package main

import (
	"crypto/tls"
	"fmt"
	//"crypto/tls"
	"gopkg.in/gomail.v2"
)

/// gomail 正文中不能出现html标签，否则有的smtp会拦截？
func main() {

	m := gomail.NewMessage()
	//m.SetHeader("From", "richter.zhang@hrlink.com.cn")// gomail: could not send email 1: 501 mail from address must be same as authorization user
	m.SetHeader("From", "hrlinktest@hrlink.com.cn")
	m.SetHeader("To", "richter.zhang@hrlink.com.cn")
	m.SetAddressHeader("Cc", "promvc@live.com", "PROMVC")
	m.SetHeader("Subject", "email test by gopkg.in gomail.v2")
	//m.SetBody("text/html", "email test by  <b>gopkg.in/gomail.v2</b> and <i>gopkg.in/gomail.v2</i>!")
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
