package main

import (
	"time"
	"fmt"
)

func main1(){
	timer:=time.NewTimer(10*time.Second)
	fmt.Println(time.Now())
	etime:=<-timer.C //插入时间，10秒以后在执行。
	fmt.Println(etime)
	fmt.Println(timer.Stop())


}

//管道读取超时
func main2(){
	intChan :=make(chan int ,1)
	go func() {
		time.Sleep(time.Millisecond*500)
		intChan<-1
	}()
	select {
	case e:=<-intChan:
		fmt.Println(e)
	case <-time.NewTimer(time.Millisecond*400).C:
		fmt.Println("time out")
	}
}