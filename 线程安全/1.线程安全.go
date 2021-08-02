package main

import (
	"fmt"
	"time"
)

var money int

func Add(p* int){
	for i:=0;i<100000;i++{
		*p++
	}
}

func main(){
	money=0
	p:=&money
	for i:=0;i<100;i++{
		  Add(p)
	}
	time.Sleep(time.Second*10)
	fmt.Println(p,*p)
}
