package main

import (
	"time"
	"fmt"
)

func main(){
	intChan:=make(chan int,1)
	go func(){
		for i:=0;i<20;i++{
			time.Sleep(time.Second*2)
			intChan<-i
		}
	}()

	timeout:=time.Second*3
	var timer *time.Timer//定时器
	timeAfterC:=<-time.After(100*time.Millisecond)
	fmt.Println("timeAfterC",timeAfterC)
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("2秒以后执行")
	})
	for{
		if timer==nil{
			timer=time.NewTimer(timeout)//新构建
		}else{
			timer.Reset(timeout)//重置
		}
		select {
		case e,ok:=<-intChan:
			if!ok{
				fmt.Println("received end")
				return
			}else{
				fmt.Println("received ",e)
			}
		case c:=<-timer.C:
			fmt.Println("超时",c)
		}
	}







}
