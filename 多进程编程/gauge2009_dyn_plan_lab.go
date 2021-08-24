package main

import "fmt"

func main() {
	/// 带缓存的递归法
	fmt.Printf("n=0 : %v\n", numWays(0))
	fmt.Printf("n=1 : %v\n", numWays(1))
	fmt.Printf("n=2 : %v\n", numWays(2))

	/// 自底向上的动态规划法
	fmt.Printf("numWays_dp| n=0 : %v\n", numWays_dp(0))
	fmt.Printf("numWays_dp| n=1 : %v\n", numWays_dp(1))
	fmt.Printf("numWays_dp| n=2 : %v\n", numWays_dp(2))

}

var tempDic = make(map[int]int)

func numWays(n int) int {

	// n = 0 也算1种
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	//先判断有没计算过，即看看备忘录有没有
	if _, ok := tempDic[n]; ok {
		//存在
		//备忘录有，即计算过，直接返回
		return tempDic[n]
	} else {
		// 备忘录没有，即没有计算过，执行递归计算,并且把结果保存到备忘录map中，对1000000007取余（这个是leetcode题目规定的）
		tempDic[n] = (numWays(n-1) + numWays(n-2)) % 1000000007
		return tempDic[n]

	}

}

func numWays_dp(n int) int {
	// n = 0 也算是一种
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	lastoflast := 1
	last := 2
	temp := 0
	for l := 3; l <= n; l++ {
		temp = (lastoflast + last) % 1000000007
		lastoflast = last
		last = temp
	}
	return temp
}
