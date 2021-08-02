package main

import (
	"sync"
	"fmt"
	"runtime"
	"time"
)

type sever struct{
	clients uint64
	cond *sync.Cond//条件变量
}
//创建服务器
func Newsever()* sever{
	s:=&sever{}
	s.cond=sync.NewCond(&sync.Mutex{})
	return s
}
func (s*sever)IOLoop(){
	for{
		s.cond.L.Lock()
		for s.clients==3{
			fmt.Println("等待消息，有消息继续执行",s.clients)
			s.cond.Wait()
			fmt.Println("后续执行",s.clients)
		}

		s.cond.L.Unlock()
		s.clients++
		fmt.Println("IOLoop()",s.clients)
	}
}

func (s*sever)Release(){
	s.cond.L.Lock()
	s.clients--
	fmt.Println("Release()  up",s.clients)
	s.cond.Signal()//发出信号
	fmt.Println("Release()  down",s.clients)
	s.cond.L.Unlock()
}




func main(){
	runtime.GOMAXPROCS(4)
	s:=Newsever()
	go s.IOLoop()//并发循环等待

	time.Sleep(time.Second*3)
	go func() {
		s.Release()
	}()
	go func() {
		s.Release()
	}()
	time.Sleep(time.Second*1)
	s.Release()
	time.Sleep(time.Second*1)
	s.Release()
	time.Sleep(time.Second*1)
	s.Release()
	fmt.Println("game over")
}
