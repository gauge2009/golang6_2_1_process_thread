package main

import (
	"bytes"
	"fmt"
	ps "github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
	"github.com/bhendo/go-powershell/middleware"
	"io/ioutil"
	"log"
	"os/exec"
	//"Common"
	"path"
	"runtime"
)

//ReviveTMW_x86.ps1

func main() {
	//InvokePowershell()
	//RemotePowershell()
	//InvokePowershellScript()
	//ExceuteScript()
	//InvokePowershellComplexScript()
	//InvokePowershellMultiRowFile("ps_multi_row.ps1") // 成功版 （ASCII）
	InvokePowershellMultiRowFile("ReviveTMW.ps1") // 成功版 （带有BOM的UTF-8）
}

//调用多行脚本文件
func InvokePowershellMultiRowFile(filename string) {
	var out bytes.Buffer
	path := FetchCurrentAbPathByCaller() + "/" + filename
	fmt.Println("path = ", path)
	//cmd := exec.Command("powershell",  `D:\......\多进程编程\ps_multi_row.ps1` )
	cmd := exec.Command("powershell", path)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s", out.String())
}

/// 调用脚本
func InvokePowershellComplexScript() {
	// choose a backend
	back := &backend.Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	//cmd:="Get-WmiObject -Class Win32_Processor"
	//path_root := Common.GetCurrentAbPathByCaller() // // go run 与 go build  统一使用go run 制指定的
	path_root := FetchCurrentAbPathByCaller()
	fmt.Println("GetCurrentAbPathByCaller = ", path_root)
	cmd := ReadScriprFile(path_root + "/ps_multi_row.ps1")

	// ... and interact with it
	//stdout, stderr, err := shell.Execute(path_root+"/ps_multi_row.ps1")
	stdout, stderr, err := shell.Execute(cmd)
	if stderr != "" {
		fmt.Println(stderr)
	}
	if err != nil {
		panic(err)
	}

	fmt.Println(stdout)
}

/// 调用脚本
func ExceuteScript() {
	posh := New()

	//path_root := Common.GetCurrentAbPathByCaller() // // go run 与 go build  统一使用go run 制指定的
	path_root := FetchCurrentAbPathByCaller()
	fmt.Println("GetCurrentAbPathByCaller = ", path_root)
	cmd := ReadScriprFile(path_root + "/psdemo.ps1")
	stdout, stderr, err := posh.Execute("$OutputEncoding = [Console]::OutputEncoding; " + cmd + " ")

	fmt.Println(stdout)
	fmt.Println(stderr)

	if err != nil {
		fmt.Println(err)
	}
}

/// Golang invoking PowerShell.exe always returns ASCII characters
type PowerShell struct {
	powerShell string
}

///Golang invoking PowerShell.exe always returns ASCII characters
func New() *PowerShell {
	ps, _ := exec.LookPath("powershell.exe")
	return &PowerShell{
		powerShell: ps,
	}
}

///Golang invoking PowerShell.exe always returns ASCII characters
func (p *PowerShell) Execute(args ...string) (stdOut string, stdErr string, err error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(p.powerShell, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}

// 获取当前执行文件绝对路径（go run）
func FetchCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func ReadScriprFile(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read fail", err)
	}
	return string(f)
}

/// 调用脚本
func InvokePowershellScript() {
	// choose a backend
	back := &backend.Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	//cmd:="Get-WmiObject -Class Win32_Processor"
	//path_root := Common.GetCurrentAbPathByCaller() // // go run 与 go build  统一使用go run 制指定的
	path_root := FetchCurrentAbPathByCaller()
	fmt.Println("GetCurrentAbPathByCaller = ", path_root)
	cmd := ReadScriprFile(path_root + "/psdemo.ps1")

	// ... and interact with it
	stdout, stderr, err := shell.Execute(cmd)
	if stderr != "" {
		fmt.Println(stderr)
	}
	if err != nil {
		panic(err)
	}

	fmt.Println(stdout)
}

/// 远程执行
func RemotePowershell() {
	// choose a backend
	back := &backend.Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}

	// prepare remote session configuration
	config := middleware.NewSessionConfig()
	config.ComputerName = "devops2020"

	// create a new shell by wrapping the existing one in the session middleware
	session, err := middleware.NewSession(shell, config)
	if err != nil {
		panic(err)
	}
	defer session.Exit() // will also close the underlying ps shell!

	// everything run via the session is run on the remote machine
	stdout, stderr, err := session.Execute("Get-WmiObject -Class Win32_Processor")
	if err != nil {
		panic(err)
	}
	if stderr != "" {
		fmt.Println(stderr)
	}
	fmt.Println(stdout)
}

/// 调用命令
func InvokePowershell() {
	// choose a backend
	back := &backend.Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	// ... and interact with it
	stdout, stderr, err := shell.Execute("Get-WmiObject -Class Win32_Processor")
	if stderr != "" {
		fmt.Println(stderr)
	}
	if err != nil {
		panic(err)
	}

	fmt.Println(stdout)
}
