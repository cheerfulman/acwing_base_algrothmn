package main

import "fmt"

func main() {
	var (
		n, q, t int
	)

	fmt.Scanf("%d%d", &n, &q)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	for i := 0; i < q; i++ {
		fmt.Scanf("%d", &t)
		fmt.Printf("%d %d\n", search_min(arr, t), search_max(arr, t))
	}
}

func search_min(arr []int, target int) int {
	i, j := 0, len(arr)-1
	for i < j {
		mid := (i + j) / 2
		if arr[mid] >= target {
			j = mid
		} else {
			i = mid + 1
		}
	}
	if arr[i] != target {
		return -1
	}
	return i
}

func search_max(arr []int, target int) int {
	i, j := 0, len(arr)-1
	for i < j {
		mid := (i + j + 1) / 2
		if arr[mid] <= target {
			i = mid
		} else {
			j = mid - 1
		}
	}
	if arr[i] != target {
		return -1
	}
	return i
}
