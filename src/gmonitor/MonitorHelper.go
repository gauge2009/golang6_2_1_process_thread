package gmonitor

import (
	"encoding/json"
	"fmt"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
	"runtime"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/v3/mem"
	// "github.com/shirou/gopsutil/mem"  // to use v2
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/shirou/gopsutil/v3/winservices"
) //cpu：CPU 相关；
//disk：磁盘相关；
//docker：docker 相关；
//host：主机相关；
//mem：内存相关；
//net：网络相关；
//process：进程相关；
//winservices：Windows 服务相关。

// src/github.com/shirou/gopsutil/cpu/cpu.go
type TimesStat struct {
	CPU       string  `json:"cpu"`
	User      float64 `json:"user"`
	System    float64 `json:"system"`
	Idle      float64 `json:"idle"`
	Nice      float64 `json:"nice"`
	Iowait    float64 `json:"iowait"`
	Irq       float64 `json:"irq"`
	Softirq   float64 `json:"softirq"`
	Steal     float64 `json:"steal"`
	Guest     float64 `json:"guest"`
	GuestNice float64 `json:"guestNice"`
}

// src/github.com/shirou/gopsutil/disk.go
type UsageStat struct {
	Path              string  `json:"path"`
	Fstype            string  `json:"fstype"`
	Total             uint64  `json:"total"`
	Free              uint64  `json:"free"`
	Used              uint64  `json:"used"`
	UsedPercent       float64 `json:"usedPercent"`
	InodesTotal       uint64  `json:"inodesTotal"`
	InodesUsed        uint64  `json:"inodesUsed"`
	InodesFree        uint64  `json:"inodesFree"`
	InodesUsedPercent float64 `json:"inodesUsedPercent"`
}

//host.Users()返回终端连接上来的用户信息，每个用户一个UserStat结构：
// src/github.com/shirou/gopsutil/host/host.go
type UserStat struct {
	User     string `json:"user"`
	Terminal string `json:"terminal"`
	Host     string `json:"host"`
	Started  int    `json:"started"`
}

func MoniorMemory() {
	fmt.Printf("██ ██ ██ %v\n", funcName())
	v, _ := mem.VirtualMemory()

	// 单位为字节
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	fmt.Println(v)
}

/// 5秒采样一次，存储在nosql，每次采样一个时间戳，过期时间1小时
/// 每次采样，顺便统计各个时间窗口的平均使用率
func MoniorCPU() {
	fmt.Printf("██ ██ ██ %v\n", funcName())
	// 获取 CPU 的详细信息
	infos, _ := cpu.Info()
	for _, info := range infos {
		data, _ := json.MarshalIndent(info, "", " ")
		fmt.Print(string(data))
	}

	physicalCnt, _ := cpu.Counts(false)
	logicalCnt, _ := cpu.Counts(true)
	fmt.Printf("physical count:%d logical count:%d\n", physicalCnt, logicalCnt)

	totalPercent, _ := cpu.Percent(3*time.Second, false)
	perPercents, _ := cpu.Percent(3*time.Second, true)
	fmt.Printf("总使用率:%v 每核心利用率%v", totalPercent, perPercents)

	//从开机算起，总 CPU 和 每个单独的 CPU 时间占用情况
	infos_during, _ := cpu.Times(true)
	i := 0
	for _, info := range infos_during {
		if i <= 2 { // 只显示前三个内核
			data, _ := json.MarshalIndent(info, "", " ")
			fmt.Print(string(data))
		}
		i++
	}

}

func MoniorDisk() {
	fmt.Printf("██ ██ ██ %v\n", funcName())
	fmt.Printf("██ ██ ██ %v\n", "磁盘读写")
	mapStat, _ := disk.IOCounters()
	for name, stat := range mapStat {
		fmt.Println(name)
		data, _ := json.MarshalIndent(stat, "", "  ")
		fmt.Println(string(data))
	}
	fmt.Printf("██ ██ ██ %v\n", "磁盘容量")
	info, _ := disk.Usage("D:\\T2_21") //由于返回的是磁盘的使用情况，所以路径D:/code/golang和D:返回同样的结果，只是结构中的Path字段不同而已
	data, _ := json.MarshalIndent(info, "", "  ")
	fmt.Println(string(data))

	usageStat := UsageStat{}
	err := json.Unmarshal([]byte(data), &usageStat)
	if err != nil {
		fmt.Println("反序列化失败", err)
	}
	fmt.Printf("%+v\n", usageStat)
	fmt.Println(usageStat)
	uPercent, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", usageStat.UsedPercent), 64)
	fmt.Printf("%+v盘使用率：%+v\n", usageStat.Path, uPercent)
}

func MoniorHost() {
	fmt.Printf("██ ██ ██ %v\n", funcName())
	fmt.Printf("██ ██ ██ %v\n", "Host")
	//返回主机开机时间的时间戳
	timestamp, _ := host.BootTime()
	t := time.Unix(int64(timestamp), 0)
	fmt.Println(t.Local().Format("2006-01-02 15:04:05"))

	//内核版本和平台信息
	version, _ := host.KernelVersion()
	fmt.Println(version)

	platform, family, version, _ := host.PlatformInformation()
	fmt.Println("platform:", platform)
	fmt.Println("family:", family)
	fmt.Println("version:", version)

	///终端用户
	users, _ := host.Users()
	for _, user := range users {
		data, _ := json.MarshalIndent(user, "", " ")
		fmt.Println(string(data))
	}

}

func MoniorProcess() {
	fmt.Printf("██ ██ ██ %v\n", funcName())
	fmt.Printf("██ ██ ██ %v\n", "Processes")
	var rootProcess *process.Process
	processes, _ := process.Processes()
	for _, p := range processes {
		if p.Pid == 9120 {
			rootProcess = p
			break
		}
		//if p.Name() == "HRLink.TaskScheduler" {
		//	rootProcess = p
		//	break
		//}
	}

	fmt.Println(rootProcess)

	fmt.Println("children:")
	children, _ := rootProcess.Children()
	for _, p := range children {
		fmt.Println(p)
	}

}

//winservices子包可以获取 Windows 系统中的服务信息，内部使用了golang.org/x/sys包。在winservices中，一个服务对应一个Service结构：
// src/github.com/shirou/gopsutil/winservices/winservices.go
type Service struct {
	Name   string
	Config mgr.Config
	Status ServiceStatus
	// contains filtered or unexported fields
}

//mgr.Config为包golang.org/x/sys中的结构，该结构详细记录了服务类型、启动类型（自动/手动）、二进制文件路径等信息：
// src/golang.org/x/sys/windows/svc/mgr/config.go
type Config struct {
	ServiceType      uint32
	StartType        uint32
	ErrorControl     uint32
	BinaryPathName   string
	LoadOrderGroup   string
	TagId            uint32
	Dependencies     []string
	ServiceStartName string
	DisplayName      string
	Password         string
	Description      string
	SidType          uint32
	DelayedAutoStart bool
}

//ServiceStatus结构记录了服务的状态：
// src/github.com/shirou/gopsutil/winservices/winservices.go
type ServiceStatus struct {
	State         svc.State
	Accepts       svc.Accepted
	Pid           uint32
	Win32ExitCode uint32
}

//State：为服务状态，有已停止、运行、暂停等；
//Accepts：表示服务接收哪些操作，有暂停、继续、会话切换等；
//Pid：进程 ID；
//Win32ExitCode：应用程序退出状态码。
func MoniorWindowsService() {
	fmt.Printf("██ ██ ██ %v\n", funcName())
	fmt.Printf("██ ██ ██ %v\n", "Processes")
	services, _ := winservices.ListServices()

	for _, service := range services {
		if service.Name == "HRLink.TaskScheduler" {
			newservice, _ := winservices.NewService(service.Name)
			err := newservice.GetServiceDetail()
			if err != nil {
				println("获取服务信息失败，%+v\n", err)
			}
			fmt.Println("Name:", newservice.Name, "Binary Path:", newservice.Config.BinaryPathName, "State: ", newservice.Status.State)
			fmt.Println("Binary Path:", newservice.Config.BinaryPathName)
			fmt.Println("State: ", newservice.Status.State)

		}
	}

}

///  反射获取当前函数名
func funcName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
