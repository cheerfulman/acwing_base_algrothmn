## 大数

### 大数加法

```go
// 大数加法倒序
func add(A, B []int) []int {
	var c []int
	t := 0
	for i := 0; i < len(A) || i < len(B); i++ {
		if i < len(A) {
			t += A[i]
		}
		if i < len(B) {
			t += B[i]
		}
		c = append(c, t%10)
		t /= 10
	}
	if t != 0 {
		c = append(c, 1)
	}
	return c
}
```

### 大数减法

```go
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
```

## 大数乘法

```go
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
```

## 大数除法

```go
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
```

