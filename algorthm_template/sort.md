## 快排

```go
// 快排模板
func q_sort(q []int, l, r int) {
	if l >= r {
		return
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
	// i >= j
	// q[(l, i - 1)] <= x q[i] >= x      q[(j + 1, r)] >= x, q[j] <= x
	q_sort(q, l, j)
	q_sort(q, j+1, r)
}

```

## 归并排序

```go
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

	for i, j := l, 0; i <= r; i, j = i+1, j+1 {
		q[i] = tmp[j]
	}
}
```

