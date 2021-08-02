package main

import "fmt"

func main(){
	chanCap:=5
	intChan:=make(chan int ,chanCap)
	for i:=0;i<chanCap;i++{
		select {
		case intChan<-1:
			fmt.Println("s1",i)
		case intChan<-2:
			fmt.Println("s2",i)
		case intChan<-3:
			fmt.Println("s3",i)
		}

	}

	for i:=00;i<chanCap;i++{
		fmt.Printf("%d\n",<-intChan)
	}
}
