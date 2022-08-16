package gmonitor

import (
	"net/smtp"
	"strings"
)

// src/github.com/shirou/gopsutil/disk.go
type SendToMailViewModel struct {
	User     string   `json:"user"`
	Password string   `json:"password"`
	Host     string   `json:"host"`
	Subject  string   `json:"subject"`
	Body     string   `json:"body"`
	Mailtype string   `json:"mailtype"`
	Send_to  []string `json:"send_to"`
}

///发送成功的场景：content_type 必须是 text/html 且不能有html标签
/// 无附件、无抄送
func SendToMailV2(vm SendToMailViewModel) error {
	hp := strings.Split(vm.Host, ":")
	auth := smtp.PlainAuth("", vm.User, vm.Password, hp[0])
	var content_type string
	if vm.Mailtype == "html" {
		content_type = "Content-Type: text/" + vm.Mailtype + "; charset=UTF-8" //发送成功的场景：content_type 必须是 text/html 且不能有html标签
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	to_address := strings.Join(vm.Send_to, ";")
	msg := []byte("To: " + to_address + "\r\nFrom: " + vm.User + "\r\nSubject: " + vm.Subject + "\r\n " + "\r\n" + content_type + "\r\n\r\n" + vm.Body)
	err := smtp.SendMail(vm.Host, auth, vm.User, vm.Send_to, msg)
	return err
}

/// 无附件、无抄送
func SendToMail(user, password, host, subject, body, mailtype string, send_to []string) error {
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

//func SendToMail(user, password, host, subject, body, mailtype, replyToAddress string, to, cc, bcc []string) error {
//	hp := strings.Split(host, ":")
//	auth := smtp.PlainAuth("", user, password, hp[0])
//	var content_type string
//	if mailtype == "html" {
//		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
//	} else {
//		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
//	}
//	cc_address := strings.Join(cc, ";")
//	bcc_address := strings.Join(bcc, ";")
//	to_address := strings.Join(to, ";")
//	msg := []byte("To: " + to_address + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\nReply-To: " + replyToAddress + "\r\nCc: " + cc_address + "\r\nBcc: " + bcc_address + "\r\n" + content_type + "\r\n\r\n" + body)
//	send_to := MergeSlice(to, cc)
//	send_to = MergeSlice(send_to, bcc)
//	err := smtp.SendMail(host, auth, user, send_to, msg)
//	return err
//}

//func MergeSlice(s1 []string, s2 []string) []string {
//	slice := make([]string, len(s1)+len(s2))
//	copy(slice, s1)
//	copy(slice[len(s1):], s2)
//	return slice
//}
