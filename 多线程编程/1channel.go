package main

import (
	"fmt"
	"time"
)

func main(){
	//1G文件，10份，搜索yincheng,结果归并
	var strChan=make(chan string,3)
	syncChannel1:=make(chan struct{},1)
	syncChannel2:=make(chan struct{},2)
	//读取
	go func() {
		<-syncChannel1
		fmt.Println("syncChannel1收到了信号")
		time.Sleep(time.Second)
		for{
			if elem,ok:=<-strChan;ok{
				fmt.Println("收到",elem,"recv")
			}else{
				break
			}
		}

        fmt.Println("停止接收")
		syncChannel2<- struct{}{}
	}()
	//写入
	go func() {
		for _,elem:=range[]string{"a","b","c","d"}{
			strChan<-elem //压入数据到
			fmt.Println("Send",elem,"sender")
			if elem=="c"{
				syncChannel1<- struct{}{}
				fmt.Println("send signal sender")
			}
		}
		fmt.Println("wait 2 s")
		time.Sleep(time.Second*2)
		close(strChan)//关闭
		syncChannel2<- struct{}{}

	}()
	<-syncChannel2
	<-syncChannel2

}
