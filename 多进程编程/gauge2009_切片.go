package main

import (
	"fmt"

)


func main(){

	//spliceDemo_append()

	 spliceDemo_delete()
}


func spliceDemo_append() {
	arg := []string{}// 切片
	arg = append(arg,"inspect,60004")// 切片后加一个元素
	arg = append(arg,"inspect,60005","inspect,60006")// 切片后加两个元素
	arg = append(arg,[]string{"sidecar,26666","sidecar,27777"}...)// 切片后加多元素【传入的临时切片被打散为多元素】
	arg = append([]string{"salary,61001","salary,61002"}, arg...)// 切片前加多元素【在传入的临时切片后加原切片，原切片需要被打散为多元素才能传入】

	/*在切片第i个元素开始插入切片（多元素）*/
	i :=  2
	//arg = append([]string{"payroll,64001","payroll,64002"}, arg[i:]...)// 切片第i个元素前加多元素【在传入的临时切片后加原切片，原切片第i个元素之后需要被打散为多元素才能传入】
	arg = append(arg[:i],append([]string{"payroll,64001","payroll,64002"}, arg[i:]...)...)// 在切片第i个元素开始插入切片（多元素）
	// OR:
	ext := []string{}// 切片
	ext = append(arg,"ext,00000","ext,00001")
	arg = append(arg,ext...)
	copy(arg[i+len(ext):],arg[i:])// 相当于 i之后的元素统一向后移动len(ext)个位置
    copy(arg[i:],ext)//n
    for _, e:= range arg{
		//fmt.Printf("spliceDemo_append 结果:\n", e )
		fmt.Printf("spliceDemo_append 结果:%v\n", e)
	}
	fmt.Printf("spliceDemo_append  ok" )

}


func spliceDemo_delete() {
	arg := []string{}// 切片
	arg = append(arg,"inspect,60004")// 切片后加一个元素
	arg = append(arg,"inspect,60005","inspect,60006")// 切片后加两个元素
	arg = append(arg,[]string{"sidecar,26666","sidecar,27777"}...)// 切片后加多元素【传入的临时切片被打散为多元素】
	arg = append([]string{"salary,61001","salary,61002"}, arg...)// 切片前加多元素【在传入的临时切片后加原切片，原切片需要被打散为多元素才能传入】
	for _, e:= range arg{
		fmt.Printf("spliceDemo_delete:%v\n", e )
	}
	/*删除尾部元素*/
	fmt.Printf("删除尾部元素 ：" )
     arg = arg[:len(arg)-1]
	for _, e:= range arg{
		fmt.Printf(" 结果:%v\n", e )
	}
	/*删除头部元素*/
	fmt.Printf("删除头部N个元素 ：" )
	N:= 2
	arg = arg[N:]
	for _, e:= range arg{
		fmt.Printf(" 结果:%v\n", e )
	}



}
