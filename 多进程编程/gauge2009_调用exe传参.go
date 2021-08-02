package main

import (
	"fmt"
	"os/exec"
)


func main(){

	callEXE1()
}


func callEXE1() {
	arg := []string{}
	arg = append(arg,"inspect,60004")

	//cmd1 := exec.Command("cd C:\\src\\Saas-TaskTrack\\框架\\TMW\\ActorStage\\HRLink.ExecutorStage\\bin\\Debug" )
	//output1, err := cmd1.CombinedOutput()
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//fmt.Printf("callEXE2结果:\n%v\n\n%v\n\n%v", string(output1), cmd1.Stdout, cmd1.Stderr)

	cmd := exec.Command("C:\\src\\Saas-TaskTrack\\框架\\TMW\\ActorStage\\HRLink.ExecutorStage\\bin\\Debug\\HRLink.Actor.inspect.ric", arg...)
	//会向 cmd.Stdout和cmd.Stderr写入信息,其实cmd.Stdout==cmd.Stderr,具体可见源码
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("callEXE2结果:\n%v\n\n%v\n\n%v", string(output), cmd.Stdout, cmd.Stderr)
}

func callEXE2() {
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