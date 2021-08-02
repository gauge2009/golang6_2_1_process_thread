package main

import (
	"fmt"
	"time"
)

var mapChan=make(chan map[string]int,1)//引用类型

func main(){
	syncChan:=make(chan struct{},2)
	go func() {
		for {
			if elem ,ok:=<-mapChan;ok{
				elem["count"]++
			}else{
				break
			}
		}
		fmt.Println("stop 接收")

		syncChan<- struct{}{}
	}()

	go func() {
		countmap:=make(map[string]int)
		for i:=0;i<5;i++{
			mapChan<-countmap//压入数据
			time.Sleep(time.Microsecond)
			fmt.Printf("countmap %v\n",countmap)
		}
		close(mapChan)
		syncChan<- struct{}{}
	}()
	<-syncChan
	<-syncChan



}
