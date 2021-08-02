package main

import "sync"
import (
	"time"
	"fmt"
)
//高频读写，用内存池，提高速度
var bytetools=sync.Pool{New: func() interface{} {
	b:=make([]byte ,4096)
	return &b
}}

func main1(){
	nowtime:=time.Now()
	for i:=0;i<100000000;i++{
		obj:=bytetools.Get().(*[]byte)
		_=obj
		bytetools.Put(obj)
	}
	fmt.Println(time.Since(nowtime))
}
func main(){
	nowtime:=time.Now()
	for i:=0;i<100000000;i++{
		obj:=make([]byte,1024)
		_=obj

	}
	fmt.Println(time.Since(nowtime))
}