package main

import (
	"time"
	"fmt"
)

func main(){
	sendtime:=time.Second
	recetime:=time.Second*3

	//非缓冲管道,谁先操作谁先阻塞，谁后操作谁先完成
	intchan:=make(chan int ,0)
	go func() {
		var ts0,ts1 int64
		for i:=0;i<=5;i++{
			intchan<-i
			ts1=time.Now().Unix()
			if ts0==0{
				fmt.Println("send",i)
			}else{
				fmt.Println("send !",i,ts1-ts0)
			}
			ts0=time.Now().Unix()
			time.Sleep(sendtime)
		}
		close(intchan)
	}()
	var ts0,ts1 int64
	Loop:
		for{
			select {
			case v,ok:=<-intchan:
				if !ok{
					break Loop
				}
				ts1=time.Now().Unix()
				if ts0==0{
					fmt.Println("rece",v)
				}else{
					fmt.Println("rece !",v,ts1-ts0)
				}

			}
			ts0=time.Now().Unix()
			time.Sleep(recetime)
		}

	fmt.Println("end")



}
