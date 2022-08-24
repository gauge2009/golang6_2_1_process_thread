package main

import (
	"fmt"
	"os/exec"
)

func main() {

	callEXE01()
}

func callEXE01() {
	arg := []string{}
	arg = append(arg, "主题,正文")

	//cmd1 := exec.Command("cd C:\\src\\Saas-TaskTrack\\框架\\TMW\\ActorStage\\HRLink.ExecutorStage\\bin\\Debug" )
	//output1, err := cmd1.CombinedOutput()
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//fmt.Printf("callEXE2结果:\n%v\n\n%v\n\n%v", string(output1), cmd1.Stdout, cmd1.Stderr)

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

func callEXE02() {
	arg := []string{}
	cmd := exec.Command("test", arg...)
	//会向 cmd.Stdout写入信息
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("callEXE2结果:\n%v\n%v\n%v", string(output), cmd.Stdout, cmd.Stderr)
}
