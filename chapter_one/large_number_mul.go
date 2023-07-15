package main

import (
	"fmt"
)

func main() {
	var (
		a, b string
		A, B []int
	)
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)
	for i := len(a) - 1; i >= 0; i-- { // A=[6,5,4,3,2,1]
		A = append(A, int(a[i]-'0')) // 存储数字到数组中
	}
	for i := len(b) - 1; i >= 0; i-- {
		B = append(B, int(b[i]-'0')) // 存储数字到数组中
	}
	c := mul(A, B)
	for i := len(c) - 1; i >= 0; i-- {
		fmt.Printf("%d", c[i])
	}
}

// 大数乘法倒序
// 1. 将每个数统一相乘
// 2. 统一进位
// 3. 去掉前导0  -- 比如100 * 0 = 000， 100 * 10 = 01000
func mul(A, B []int) []int {
	c := make([]int, len(A)+len(B))
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			c[i+j] += A[i] * B[j]
		}
	}

	for i, t := 0, 0; i < len(c) || t > 0; i++ {
		t += c[i]
		c[i] = t % 10
		t /= 10
	}

	for len(c) > 1 && c[len(c)-1] == 0 {
		c = c[:len(c)-1]
	}
	return c
}
