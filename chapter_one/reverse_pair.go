package main

import "fmt"

var ans int = 0

func main() {
	var (
		n int
	)

	fmt.Scanf("%d", &n)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	tmp := make([]int, n)
	merge_sort(q, tmp, 0, n-1)
	fmt.Printf("%d", ans)
}

// 归并模板
func merge_sort(q, tmp []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	merge_sort(q, tmp, l, mid)
	merge_sort(q, tmp, mid+1, r)
	k, i, j := 0, l, mid+1
	for i <= mid && j <= r {
		if q[i] <= q[j] {
			tmp[k] = q[i]
			k, i = k+1, i+1
		} else {
			ans += mid - i + 1
			tmp[k] = q[j]
			k, j = k+1, j+1
		}
	}

	for ; i <= mid; k, i = k+1, i+1 {
		tmp[k] = q[i]
	}
	for ; j <= r; k, j = k+1, j+1 {
		tmp[k] = q[j]
	}

	for i, j = 0, l; j <= r; i, j = i+1, j+1 {
		q[j] = tmp[i]
	}
}

func reverse_pair(q, tmp []int, l, r int) int {
	if l >= r {
		return 0
	}
	mid := (l + r) / 2
	ans := reverse_pair(q, tmp, l, mid) + reverse_pair(q, tmp, mid+1, r)

	k, i, j := 0, l, mid+1
	for i <= mid && j <= r {
		if q[i] <= q[j] {
			tmp[k] = q[i]
			k, i = k+1, i+1
		} else {
			ans += mid - i + 1
			tmp[k] = q[j]
			k, j = k+1, j+1
		}
	}

	for ; i <= mid; k, i = k+1, i+1 {
		tmp[k] = q[i]
	}
	for ; j <= r; k, j = k+1, j+1 {
		tmp[k] = q[j]
	}

	for i, j = 0, l; j <= r; i, j = i+1, j+1 {
		q[j] = tmp[i]
	}
	return ans
}
