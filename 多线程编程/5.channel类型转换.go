package main

import "fmt"

func main(){
	var ok bool
	ch:=make(chan int ,1)
	_,ok=interface{}(ch).(<-chan int )
	fmt.Println(ok)
	_,ok=interface{}(ch).(chan<- int )
	fmt.Println(ok)
	_,ok=interface{}(ch).(chan int )
	fmt.Println(ok)

	sch:=make(chan <-int,1)//单向发送
	_,ok=interface{}(sch).(<-chan int )
	fmt.Println(ok)
	_,ok=interface{}(sch).(chan<- int )
	fmt.Println(ok)
	_,ok=interface{}(sch).(chan int )
	fmt.Println(ok)


	rch:=make( <- chan int,1)//单向接收
	_,ok=interface{}(rch).(<-chan int )
	fmt.Println(ok)
	_,ok=interface{}(rch).(chan<- int )
	fmt.Println(ok)
	_,ok=interface{}(rch).(chan int )
	fmt.Println(ok)


}
