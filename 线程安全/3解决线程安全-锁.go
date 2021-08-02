package main

import (
	"fmt"
	"time"
	"sync"
)

var money int
var mutex sync.Mutex

func Add(p* int){
	mutex.Lock()
	for i:=0;i<100000;i++{
		*p++
	}
	mutex.Unlock()
}

func main(){
	money=0
	p:=&money
	for i:=0;i<100;i++{
		go Add(p)
	}
	time.Sleep(time.Second*10)
	fmt.Println(p,*p)
}