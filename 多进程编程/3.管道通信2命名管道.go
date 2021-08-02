package main

import (
	"io"
	"fmt"
	"time"
	"os"
)
func SystemPiPE(){
	reader,writer,err:=os.Pipe()
	if err!=nil{
		fmt.Println(err)
		return
	}


	go func() {
		output:=make([]byte,100)
		n,err:=reader.Read(output)//读取100个字节
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println("read ",n)
	}()
	input:=make([]byte,16)//开辟数组
	for i:=65;i<=80;i++{
		input[i-65]=byte(i) //压入字母
	}
	n,err:=writer.Write(input)//写入
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("write",n)
	time.Sleep(time.Second*3)

}


func MEMPIPE(){
	reader,writer:=io.Pipe()//管道,内存
	go func() {
		output:=make([]byte,100)
		n,err:=reader.Read(output)//读取100个字节
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println("read ",n)
	}()
	input:=make([]byte,26)//开辟数组
	for i:=65;i<=90;i++{
		input[i-65]=byte(i) //压入字母
	}
	n,err:=writer.Write(input)//写入
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("write",n)
	time.Sleep(time.Second*3)

}






func main(){

	//MEMPIPE()
	SystemPiPE()



}
