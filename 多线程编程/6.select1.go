package main

import "fmt"

var intChan1 chan int
var intChan2 chan int
var channels=[]chan int{intChan1, intChan2 }

var nums=[]int{10,2,3,4,5}

func main(){

	select {
	case getchan(0)<-getnumber(0):
		fmt.Println("1")
	case getchan(1)<-getnumber(1):
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

}
func getchan(i int)chan int{
	fmt.Println("channels",i)
	return channels[i]
}
func getnumber(i int)int{
	fmt.Println("nums",i)
	return nums[i]
}
