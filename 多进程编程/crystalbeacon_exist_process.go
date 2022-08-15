package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"runtime"
	"time"
)

type ProcessStat struct {
	Name       string
	Staus_code string
}

/// 22.2.8
/// 注意 ： 监听nginx 有很多途径， 基于结果的可以走 http 响应检测； 基于进程的最为简单
func JudgeProcessExist(pname *string) *ProcessStat {
	p_stat := ProcessStat{Name: "nginx.exe", Staus_code: "404"}
	//基于进程的最为简单

	fmt.Printf("Process Name: %v | Status: %v | time: %v", p_stat.Name, p_stat.Staus_code, time.Now().Format("2006-01-02 00:00:00"))
	return &p_stat
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
		// fmt.Println(p.name) //?? gopsutil为何把 name属性设计为私有
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

///  反射获取当前函数名
func funcName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func main() {
	target := "nginx.exe"
	MoniorProcess()
	JudgeProcessExist(&target)
}
