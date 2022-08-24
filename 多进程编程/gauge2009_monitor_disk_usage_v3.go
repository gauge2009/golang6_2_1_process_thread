package main

import (
	//"crypto/tls"
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"os/exec"
	//"github.com/shirou/gopsutil/v3/mem"
	//"golang.org/x/sys/windows"
	//"gopkg.in/gomail.v2"
	//"log"
	"strconv"
	//"errors"
	//"unsafe"
)

/// https://golangbyexample.com/
/// ██ ██ ██  检查磁盘 + 预警邮件  22-817
func main() {

	/// ██ ██ ██  检查磁盘
	path := "C:"
	var ceiling float64 = 91
	var ceilingLine int = int(ceiling)

	perc := CheckMDiskUsageV3(path)
	fmt.Printf("the disk [%+v] with usage [%+v]\n", path, perc)
	if perc > ceiling {

		message := "the disk [" + path + "] with usage [" + strconv.FormatInt(int64(int(perc)), 10) + "%],exceeds the warning ceiling [" + strconv.FormatInt(int64(ceilingLine), 10) + "%] "
		fmt.Printf("%+v \n", message)

		/// ██ ██ ██  预警邮件
		subject := message
		msg := message
		//AlertByMail_ByG(subject, msg)
		AlertByMail_ByC(subject, msg)
	}
	//CheckMemory()
}

func CheckMDiskUsageV3(path string) float64 {
	v, _ := disk.Usage(path)

	percent := v.UsedPercent // float64
	//fmt.Printf("Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Used, v.UsedPercent)
	//fmt.Println(v)
	return percent
}

func callEXEV3(subject, message string) {
	arg := []string{}
	arg = append(arg, subject+"|"+message)

	/// C:\src\Saas-TaskTrack\框架\TMW\CrystalBeacon\GrystalBeacon.Callee\bin\Debug\GrystalBeacon.Callee.exe
	cmd := exec.Command("C:\\src\\Saas-TaskTrack\\框架\\TMW\\CrystalBeacon\\GrystalBeacon.Callee\\bin\\Debug\\GrystalBeacon.Callee", arg...)
	//C:\src\Saas-TaskTrack\框架\TMW\CrystalBeacon\GrystalBeacon.Callee\bin\Debug\GrystalBeacon.Callee.exe
	//会向 cmd.Stdout和cmd.Stderr写入信息,其实cmd.Stdout==cmd.Stderr,具体可见源码
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("callEXE2结果:\n%v\n\n%v\n\n%v", string(output), cmd.Stdout, cmd.Stderr)
}
func AlertByMail_ByC(subject, message string) {
	callEXEV3(subject, message)

	fmt.Println("Email Sent Successfully!")
}
