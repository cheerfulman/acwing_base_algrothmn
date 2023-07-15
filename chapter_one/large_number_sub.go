package main

import (
	"fmt"
)

func main() {
	var (
		a, b    string
		A, B, c []int
	)
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)
	for i := len(a) - 1; i >= 0; i-- { // A=[6,5,4,3,2,1]
		A = append(A, int(a[i]-'0')) // 存储数字到数组中
	}
	for i := len(b) - 1; i >= 0; i-- {
		B = append(B, int(b[i]-'0')) // 存储数字到数组中
	}

	if !cmp(A, B) {
		c = sub(B, A)
		fmt.Print("-")
	} else {
		c = sub(A, B)
	}
	for i := len(c) - 1; i >= 0; i-- {
		fmt.Printf("%d", c[i])
	}
}

// 大数加法倒序
func sub(A, B []int) []int {
	var c []int
	t := 0
	for i := 0; i < len(A); i++ {
		t = A[i] - t
		if i < len(B) {
			t -= B[i]
		}
		c = append(c, (t+10)%10)
		if t < 0 {
			t = 1
		} else {
			t = 0
		}
	}
	// 去掉前导0, 比如 1111 - 1111 = 0000
	for len(c) > 1 && c[len(c)-1] == 0 {
		c = c[:len(c)-1]
	}
	return c
}

func cmp(A, B []int) bool {
	if len(A) != len(B) {
		return len(A) > len(B)
	}
	for i := len(A); i >= 0; i-- {
		if A[i] != B[i] {
			return A[i] > B[i]
		}
	}
	return true
}
