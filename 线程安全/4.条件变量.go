package main

import (
	"sync"
	"runtime"
	"time"
	"fmt"
)

func main(){
	runtime.GOMAXPROCS(4)


	cond:=sync.NewCond(&sync.Mutex{})
	condition:=false
	go func(){
		time.Sleep(time.Second*30)
		cond.L.Lock()
		fmt.Println("子线程","状态改变，发出信号")
		condition=true
		cond.Signal()
		fmt.Println("子线程","等到消息，继续干活")
		cond.L.Unlock()
	}()

	cond.L.Lock()//等待
	fmt.Println("主线程正在运行",1)
	for !condition{
		fmt.Println("主线程正在运行",2)
		cond.Wait()
		fmt.Println("主线程正在运行",3)
	}
	fmt.Println("主线程正在运行",4)

	cond.L.Unlock()//解锁
	fmt.Println("game over")

}
