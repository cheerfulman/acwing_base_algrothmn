## 数的范围

### 左边界

```go
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
```

### 右边界

```go
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
```

## 浮点数模板

```go
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
```

