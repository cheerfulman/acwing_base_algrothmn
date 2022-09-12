package main

import "fmt"

func main() {
	var (
		n, k int
	)

	fmt.Scanf("%d%d", &n, &k)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	// k作为下标传入
	fmt.Printf("%d ", q_sort(q, 0, n-1, k-1))
}

func q_sort(q []int, l, r, k int) int {
	if l >= r {
		return q[r]
	}
	i, j, x := l-1, r+1, q[(l+r)/2]
	for i < j {
		for {
			i++
			if q[i] >= x {
				break
			}
		}

		for {
			j--
			if q[j] <= x {
				break
			}
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}
	if k <= j {
		return q_sort(q, l, j, k)
	}
	return q_sort(q, j+1, r, k)
}
