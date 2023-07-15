package main

import (
	"fmt"
)

func main() {
	var (
		a string
		b int
		A []int
	)
	fmt.Scanf("%s", &a)
	fmt.Scanf("%d", &b)
	for i := len(a) - 1; i >= 0; i-- { // A=[6,5,4,3,2,1]
		A = append(A, int(a[i]-'0')) // 存储数字到数组中
	}
	c, r := div(A, b)
	for i := len(c) - 1; i >= 0; i-- {
		fmt.Printf("%d", c[i])
	}
	fmt.Printf("\n%d\n", r)
}

// 大数除法A高精度，B整数, A是逆序
func div(A []int, b int) ([]int, int) {
	var r int
	var c []int
	for i := len(A) - 1; i >= 0; i-- {
		r = r*10 + A[i]
		c = append(c, r/b)
		r %= b
	}
	reverseInt(c)
	for len(c) > 1 && c[len(c)-1] == 0 {
		c = c[:len(c)-1]
	}
	return c, r
}

func reverseInt(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
