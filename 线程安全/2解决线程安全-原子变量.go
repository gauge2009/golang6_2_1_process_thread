package main

import (
"fmt"
"time"
"sync/atomic"
)

var money int32

func Addx(p* int32){
	for i:=0;i<100000;i++{
		atomic.AddInt32(p,1)
	}
}

func main(){
	money=0
	p:=&money
	for i:=0;i<100;i++{
		go  Addx(p)
	}
	time.Sleep(time.Second*10)
	fmt.Println(p,*p)
}

