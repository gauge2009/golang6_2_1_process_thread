package main

import (
	"crypto/tls"
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"golang.org/x/sys/windows"
	"gopkg.in/gomail.v2"
	"log"
	"strconv"

	//"errors"
	"unsafe"
)

/// https://golangbyexample.com/
/// ██ ██ ██  检查磁盘 + 预警邮件  22-817
func main() {

	/// ██ ██ ██  检查磁盘
	path := "C:"
	var ceiling float64 = 91
	var ceilingLine int = int(ceiling)

	perc := CheckMDiskUsage(path)
	fmt.Printf("the disk [%+v] with usage [%+v]\n", path, perc)
	if perc > ceiling {

		message := "the disk [" + path + "] with usage [" + strconv.FormatInt(int64(int(perc)), 10) + "%],exceeds the warning ceiling [" + strconv.FormatInt(int64(ceilingLine), 10) + "%] "
		fmt.Printf("%+v \n", message)

		/// ██ ██ ██  预警邮件
		//subject := message
		//msg := message
		subject := "subject12"
		msg := "message23"
		AlertByMail_ByG(subject, msg)

	}
	//CheckMemory()
}

func AlertByMail_ByG(subject, message string) {
	m := gomail.NewMessage()
	///////m.SetHeader("From", "richter.zhang@hrlink.com.cn")// gomail: could not send email 1: 501 mail from address must be same as authorization user
	//m.SetHeader("From", "hrlinktest@hrlink.com.cn")
	//m.SetHeader("To", "richter.zhang@hrlink.com.cn")
	//m.SetAddressHeader("Cc", "promvc@live.com", "PROMVC")
	//m.SetHeader("Subject", subject)
	//m.SetBody("text/html", message)
	////////m.Attach("/home/Alex/lolcat.jpg")
	m.SetHeaders(map[string][]string{
		"From": {m.FormatAddress("hrlinktest@hrlink.com.cn", "GrystalBeaconsdf")},
		"To":   {"richter.zhang@hrlink.com.cn"},
		//"To":      {"richter.zhang@hrlink.com.cn", "promvc@live.com"},
		//"Cc":      {"promvc@live.com"},
		"Subject": {subject},
	})
	m.SetBody("text/html", message)
	d := gomail.NewDialer("smtp.exmail.qq.com", 587, "hrlinktest@hrlink.com.cn", "!QA2ws3ed")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Email Sent Successfully!")
}

func CheckMemory() {
	v, _ := mem.VirtualMemory()

	fmt.Printf("Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Available, v.UsedPercent)

	fmt.Println(v)
}
func CheckMDiskUsage(path string) float64 {
	v, _ := disk.Usage(path)

	percent := v.UsedPercent // float64
	//fmt.Printf("Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Used, v.UsedPercent)
	//fmt.Println(v)
	return percent
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
