package main

import (
"fmt"
"time"
)

var strChan=make(chan string,3)

func recive(strChan <-chan string ,	syncChannel1<-chan struct{},
									syncChannel2 chan<- struct{}){
	<-syncChannel1
	fmt.Println("syncChannel1收到了信号")
	time.Sleep(time.Second)
	/*
	for{
		if elem,ok:=<-strChan;ok{
			fmt.Println("收到",elem,"recv")
		}else{
			break
		}
	}*/
	for elem:=range strChan{
		fmt.Println("收到",elem,"recv")
	}

	fmt.Println("停止接收")
	syncChannel2<- struct{}{}
}

func send(strChan chan<- string ,	syncChannel1 chan<- struct{},
	syncChannel2 chan<- struct{}){

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
	close(strChan)
	syncChannel2<- struct{}{}
}

func main(){
	//1G文件，10份，搜索yincheng,结果归并

	syncChannel1:=make(chan struct{},1)
	syncChannel2:=make(chan struct{},2)
	go recive(strChan,syncChannel1,syncChannel2)
	go send(strChan,syncChannel1,syncChannel2)
	<-syncChannel2
	<-syncChannel2

}

