package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"golang.org/x/sys/windows"
	"log"
	"net/smtp"

	//"errors"
	"unsafe"
)

func main() {
	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "richter.zhang@hrlink.com.cn"
	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{"richter.zhang@hrlink.com.cn"}
	// 设置主题
	em.Subject = "email test by jordan-wright /email"
	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte("Hi， email test by jordan-wright /email\"")
	//设置服务器相关的配置
	err := em.Send("smtp.exmail.qq.com:587", smtp.PlainAuth("", "hrlinktest@hrlink.com.cn", "!QA2ws3ed", "smtp.exmail.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("send successfully ... ")
}
func main1() {

	CheckMemory()
	CheckMDiskUsage("C:")
	CheckMDiskUsage("D:")
	CheckMDiskUsage("R:")
	//DiskUsage("D:")
}

func CheckMemory() {
	v, _ := mem.VirtualMemory()

	fmt.Printf("Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Available, v.UsedPercent)

	fmt.Println(v)
}
func CheckMDiskUsage(path string) {
	v, _ := disk.Usage(path)

	fmt.Printf("Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Used, v.UsedPercent)

	fmt.Println(v)
}

type DiskStatus struct {
	All  uint64
	Used uint64
	Free uint64
}

func DiskUsage(path string) (disk DiskStatus) {
	h := windows.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")
	lpFreeBytesAvailable := uint64(0)
	lpTotalNumberOfBytes := uint64(0)
	lpTotalNumberOfFreeBytes := uint64(0)
	r1, r2, err := c.Call(uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("C:"))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))
	//failOnError(err,"Failed to check FreeBytesAvailable")
	if err != nil {
		log.Fatalf("channel.open: %s", err)
	}
	fmt.Println(r1)
	fmt.Println(r2)
	disk.All = lpTotalNumberOfBytes
	disk.Free = lpTotalNumberOfFreeBytes
	disk.Used = lpFreeBytesAvailable
	return
}
