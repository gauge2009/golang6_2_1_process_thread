package main

import (
	"os/exec"
	"fmt"
	"bytes"
	"io"
	"bufio"
)

//文件缓冲区
func RunCMD(){
	cmd0:=exec.Command("tasklist")
	//cmd0:=exec.Command("ping","www.qq.com")
	stdout0,err:=cmd0.StdoutPipe()//输出
	if err!=nil{
		fmt.Println(err)
		return
	}
	if err:=cmd0.Start();err!=nil{ //开始执行
		fmt.Println(err)
		return
	}
	useBufferIO:=false
	if !useBufferIO{
		var outputBuf0 bytes.Buffer//二进制保存
		for{
			tempoutput:=make([]byte, 256)
			n,err:=stdout0.Read(tempoutput)//读取二进制
			if err!=nil{
				if err==io.EOF{
					break //跳出循环
				}else{
					fmt.Println(err)
					return
				}
			}
			if n>0{
				outputBuf0.Write(tempoutput[:n])//读取到的写入
			}
		}

		fmt.Println(outputBuf0.String() )


	}else{
		outputbuf0:=bufio.NewReader(stdout0)
		output0,_,err:=outputbuf0.ReadLine()
		if err!=nil{
			return
		}
		fmt.Println(string(output0))//读取行


	}




}
func RunCMDPipe(){
	cmd1:=exec.Command("tasklist")
	var ouputbuf1 bytes.Buffer//输出
	cmd1.Stdout=&ouputbuf1//设置输入
	if err:=cmd1.Start();err!=nil{
		fmt.Println(err)
		return
	}
	if err:=cmd1.Wait();err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("%s",ouputbuf1.Bytes())


}


func main(){
	go  RunCMDPipe()
}
