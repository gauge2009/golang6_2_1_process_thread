package main

import (
	"fmt"
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

func main() {
	target := "nginx.exe"
	JudgeProcessExist(&target)
}
