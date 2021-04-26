package main

import "fmt"

func list() {
	var balance = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(balance)

	balanceStr := [...]string{"abc", "def", "ghi"}
	fmt.Println(balanceStr[2])

	for i := 0; i < len(balance); i++ {
		fmt.Println(balance[i])
	}

	// 内建函数，append 给切片的尾部增加元素
	var balanceS []int
	balanceS = append(balanceS, 1)
	fmt.Println(balanceS)
	balanceS = append(balanceS, 2, 3)
	fmt.Println(balanceS)

	s1 := balance[2:]
	fmt.Println(s1)

	s2 := balance[:1]
	fmt.Println(s2)
}

func main() {
	list()
}
