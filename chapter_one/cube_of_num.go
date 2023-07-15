package main

import (
	"fmt"
)

func main() {
	var (
		n float64
	)
	fmt.Scanf("%f", &n)
	fmt.Printf("%.6f\n", cube_num(n))
}

func cube_num(x float64) float64 {
	var l, r float64 = -100, 100
	for r-l > 1e-8 {
		var mid float64 = (l + r) / 2
		if mid*mid*mid >= x {
			r = mid
		} else {
			l = mid
		}

	}
	return l
}
