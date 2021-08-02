package main

import (
	"fmt"
	"time"
	"strconv"
)

func main(){
	intChan:=make(chan int,10)
	stringChan:=make(chan string,10)
	go func() {
		for i:=0;i<10;i++{
			intChan<-i
			time.Sleep(time.Second)
		}
		close(intChan)
	}()


	go func() {
		for i:=0;i<10;i++{
			stringChan<- "abc"+strconv.Itoa(i)
			time.Sleep(time.Second*2)
		}
		close(intChan)
	}()
	syncChan:=make(chan struct{},1)
	go func() {
		Loop:
			for{
				select {
				case e,ok:=<-intChan:
					if !ok{
						fmt.Println("end")
						break Loop
					}
					fmt.Printf("收到%v\n",e)
				case e,ok:=<-stringChan:
					if !ok{
						fmt.Println("stringend")
						break Loop
					}
					fmt.Printf("收到%v\n",e)
				}

			}

		syncChan<- struct{}{}

	}()

	<-syncChan
}
