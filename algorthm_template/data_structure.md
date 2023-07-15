# 简单数据结构模板

## trim树 结构体版本

```go
type trimTree struct {
	son   [26]*trimTree
	flags bool
	cnt   int
}

func NewTrimTree() *trimTree {
	return &trimTree{
		flags: false,
		cnt:   0,
	}
}

func (t *trimTree) insert(str string) {
	cur := t
	for i := 0; i < len(str); i++ {
		u := str[i] - 'a'
		if cur.son[u] == nil {
			cur.son[u] = NewTrimTree()
		}
		cur = cur.son[u]
	}
	cur.flags = true
	cur.cnt++
}

func (t *trimTree) query(str string) int {
	cur := t
	for i := 0; i < len(str); i++ {
		u := str[i] - 'a'
		if cur.son[u] == nil {
			return 0
		}
		cur = cur.son[u]
	}
	if cur.flags {
		return cur.cnt
	}
	return 0
}
```

