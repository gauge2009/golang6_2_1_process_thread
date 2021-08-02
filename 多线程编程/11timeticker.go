package main

import (
	"time"
	"fmt"
)

func main(){
	intChan:=make(chan int,1)
	ticker:=time.NewTicker(time.Second*2)
	go func(){
		for _=range ticker.C{
			select {
			case intChan<-1:
			case intChan<-2:
			case intChan<-3:
			case intChan<-4:
			}
		}
		fmt.Println("end sender")
	}()

	var sum int
	for e:=range intChan{
		fmt.Println("reve",e)
		sum+=e
		fmt.Println("sum",sum)
		if sum>30{
			fmt.Println("got",sum)
			ticker.Stop()
			break
		}
	}
	fmt.Println("end receiver")





}
